package ucs

import "context"

type PositiveUC struct {
	pg PositivePGGateway
}

func NewPositiveUC(pg PositivePGGateway) PositiveUC {
	return PositiveUC{pg: pg}
}

type PositivePGGateway interface {
	AddKudos(ctx context.Context, kudos string, author string) error
	AddAnonymousKudos(ctx context.Context, kudos string) error
}

func (uc PositiveUC) Kudos(ctx context.Context, kudos string, author string) error {
	return uc.pg.AddKudos(ctx, kudos, author)
}

func (uc PositiveUC) KudosAnonymous(ctx context.Context, kudos string) error {
	return uc.pg.AddAnonymousKudos(ctx, kudos)
}
