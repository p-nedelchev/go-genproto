package rpc

import (
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

// ServiceBinder is a binder interface used to bind gRPC and HTTP services.
type ServiceBinder interface {
	HTTPBinder
	GRPCBinder
}

// HTTPBinder is a binder for an HTTP service.
type HTTPBinder interface {
	BindHTTP(bind func(*mux.Router))
}

// GRPCBinder is a biner for an GRPC service.
type GRPCBinder interface {
	BindGRPC(bind func(grpc.ServiceRegistrar))
}
