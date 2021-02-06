package httpkit_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/clouway/go-genproto/clouwayapis/rpc/httpkit"
	"github.com/clouway/go-genproto/clouwayapis/rpc/request"
)

func TestHeadersToContext(t *testing.T) {
	ctx := context.Background()
	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", "Bearer token")
	ctx = httpkit.HeadersToContext(ctx, req)

	got := ctx.Value(request.ContextKey("authorization")).(string)
	want := "Bearer token"
	if want != got {
		t.Errorf("unexpected context value:\n- want: %v\n-  got: %v", want, got)
	}
}
