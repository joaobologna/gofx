package ucs

import (
	"github.com/joaobologna/gofx/pg"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			pg.NewGateway,
			fx.As(new(PositivePGGateway)),
			fx.ResultTags(`name:"pgp"`),
		),
		fx.Annotate(
			pg.NewGateway,
			fx.As(new(NegativePGGateway)),
			fx.ResultTags(`name:"pgn"`),
		),
	),
	fx.Provide(
		fx.Annotate(NewPositiveUC, fx.ParamTags(`name:"pgp"`)),
		fx.Annotate(NewNegativeUC, fx.ParamTags(`name:"pgn"`)),
	),
)
