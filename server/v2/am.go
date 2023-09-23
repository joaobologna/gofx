package v2

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	am "github.com/joaobologna/gofx/protogen/v2"
	"github.com/joaobologna/gofx/ucs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log/slog"
)

type AssessmentManagerServer struct {
	// Embed unimplemented struct to avoid compilation errors
	// once new methods are added to AM proto.
	am.UnimplementedAssessmentManagerServer
	log        *zap.Logger
	positiveUC ucs.PositiveUC
	negativeUC ucs.NegativeUC
}

func NewAssessmentManagerServer(log *zap.Logger, positiveUC ucs.PositiveUC, negativeUC ucs.NegativeUC) *AssessmentManagerServer {
	return &AssessmentManagerServer{
		log:                                  log,
		positiveUC:                           positiveUC,
		negativeUC:                           negativeUC,
		UnimplementedAssessmentManagerServer: am.UnimplementedAssessmentManagerServer{},
	}
}

func (srv *AssessmentManagerServer) HC(context.Context, *am.HCRequest) (*am.HCResponse, error) {
	slog.Info("hc invoked")
	return new(am.HCResponse), nil
}

func (srv *AssessmentManagerServer) KudosAnonymous(ctx context.Context, in *am.KudosAnonymousRequest) (*am.KudosAnonymousResponse, error) {
	return new(am.KudosAnonymousResponse), srv.positiveUC.KudosAnonymous(ctx, in.Message)
}

func (srv *AssessmentManagerServer) ReportAnonymous(ctx context.Context, in *am.ReportAnonymousRequest) (*am.ReportAnonymousResponse, error) {
	return new(am.ReportAnonymousResponse), srv.negativeUC.TalkToUsAnonymous(ctx, in.Message)
}

func (srv *AssessmentManagerServer) Register(server *grpc.Server) {
	am.RegisterAssessmentManagerServer(server, srv)
}

func (srv *AssessmentManagerServer) RegisterHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return am.RegisterAssessmentManagerHandler(ctx, mux, conn)
}
