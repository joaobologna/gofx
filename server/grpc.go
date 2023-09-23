package server

import (
	"google.golang.org/grpc"
)

func NewGRPC(registers []Registerer) *grpc.Server {
	server := grpc.NewServer()
	for _, register := range registers {
		register.Register(server)
	}
	return server
}
