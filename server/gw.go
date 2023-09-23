package server

import (
	"context"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strings"
)

func NewHTTP(registers []Registerer, log *zap.Logger, grpcMux *grpc.Server) *http.Server {
	address := "localhost:8080"
	ctx := context.Background()

	// Always dial to localhost!
	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dail grpc server", zap.Error(err))
	}

	httpMux := gwruntime.NewServeMux()
	for _, register := range registers {
		err := register.RegisterHandler(ctx, httpMux, conn)
		if err != nil {
			log.Fatal("failed to create gw", zap.Error(err))
		}
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If request is of type gRPC, handle to grpc server.
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcMux.ServeHTTP(w, r)
			return
		}

		// If request it not of type gRPC, then handle to the http server.
		httpMux.ServeHTTP(w, r)
	})

	return &http.Server{
		Addr: address,
		// h2c allows both HTTP/1 and HTTP/2 requests to be multiplexed by this handler.
		Handler: h2c.NewHandler(handler, &http2.Server{}),
	}
}

// RunHTTP starts an HTTP server and blocks while running if successful.
// The server will be shutdown when "ctx" is canceled.
func RunHTTP(lc fx.Lifecycle, log *zap.Logger, srv *http.Server) {
	address := "localhost:8080"

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting server", zap.String("address", address))
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			<-ctx.Done()
			return srv.Shutdown(ctx)
		},
	})
}
