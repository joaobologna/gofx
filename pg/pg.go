package pg

import "context"

type Gateway struct {
	storage []Message
}

func NewGateway() Gateway {
	return Gateway{storage: make([]Message, 100)}
}

type Message struct {
	message string
	author  string
}

func (g Gateway) AddKudos(_ context.Context, kudos string, author string) error {
	g.storage = append(g.storage, Message{message: kudos, author: author})
	return nil
}

func (g Gateway) AddAnonymousKudos(_ context.Context, kudos string) error {
	g.storage = append(g.storage, Message{message: kudos})
	return nil
}

func (g Gateway) AddReport(_ context.Context, report string, author string) error {
	g.storage = append(g.storage, Message{message: report, author: author})

	return nil
}

func (g Gateway) AddAnonymousReport(_ context.Context, report string) error {
	g.storage = append(g.storage, Message{message: report})
	return nil
}
