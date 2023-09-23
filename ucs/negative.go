package ucs

import "context"

type NegativeUC struct {
	pg NegativePGGateway
}

func NewNegativeUC(pg NegativePGGateway) NegativeUC {
	return NegativeUC{pg: pg}
}

type NegativePGGateway interface {
	AddReport(ctx context.Context, report string, author string) error
	AddAnonymousReport(ctx context.Context, report string) error
}

func (uc NegativeUC) TalkToUs(ctx context.Context, report string, author string) error {
	return uc.pg.AddReport(ctx, report, author)
}

func (uc NegativeUC) TalkToUsAnonymous(ctx context.Context, report string) error {
	return uc.pg.AddAnonymousReport(ctx, report)
}
