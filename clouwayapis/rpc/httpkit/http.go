package httpkit

import (
	"context"
	"net/http"
	"strings"

	"github.com/clouway/go-genproto/clouwayapis/rpc/request"
)

// HeadersToContext adds all HTTP header values into the passed context.Context. The keys
// are added with request.ContextKey as and lookups should be performed by using the same
// type.
func HeadersToContext(ctx context.Context, r *http.Request) context.Context {
	for k := range r.Header {
		// The key is added in strings.ToLower which is the grpc metadata format of the key so
		// that it can be accessed in either format
		key := request.ContextKey(strings.ToLower(k))
		ctx = context.WithValue(ctx, key, r.Header.Get(k))
	}

	// Tune specific change.
	// also add the request url
	ctx = context.WithValue(ctx, request.ContextKey("request-url"), r.URL.Path)
	ctx = context.WithValue(ctx, request.ContextKey("transport"), "HTTPJSON")

	return ctx
}
