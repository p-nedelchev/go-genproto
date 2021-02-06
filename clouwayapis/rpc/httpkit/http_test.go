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

func TestHeadersToContextExcluding(t *testing.T) {
	ctx := context.Background()
	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", "Bearer token")
	req.Header.Add("Content-Type", "application/json")

	ctx = httpkit.HeadersToContextExcluding(ctx, req, []string{"Content-Type"})

	got := ctx.Value(request.ContextKey("content-type"))

	if got != nil {
		t.Errorf("unexpected value of exluded field:\n- want: %v\n-  got: %v", nil, got)
	}
}
