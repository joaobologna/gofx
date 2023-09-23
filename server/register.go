package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type Registerer interface {
	Register(*grpc.Server)
	RegisterHandler(context.Context, *runtime.ServeMux, *grpc.ClientConn) error
}

// AsRegisterer annotates the given constructor to state that
// it provides a server to the "servers" group.
func AsRegisterer(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Registerer)),
		fx.ResultTags(`group:"registers"`),
	)
}
