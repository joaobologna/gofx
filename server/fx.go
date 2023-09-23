package server

import (
	"github.com/joaobologna/gofx/server/v1"
	"github.com/joaobologna/gofx/server/v2"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		AsRegisterer(v1.NewAssessmentManagerServer),
		AsRegisterer(v2.NewAssessmentManagerServer),
		fx.Annotate(
			NewGRPC,
			fx.ParamTags(`group:"registers"`),
		),
		fx.Annotate(
			NewHTTP,
			fx.ParamTags(`group:"registers"`),
		),
		// NewInternalHTTP,
	),
	fx.Invoke(
		RunHTTP,
	),
)
