package grpckit

import (
	"context"

	"github.com/clouway/go-genproto/clouwayapis/rpc/request"
	"google.golang.org/grpc/metadata"
)

// MetadataToContext converts metadata of gRPC to context variables.
func MetadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added in metadata key format (k).
			ctx = context.WithValue(ctx, request.ContextKey(k), v[0])
		}
	}

	return ctx
}
