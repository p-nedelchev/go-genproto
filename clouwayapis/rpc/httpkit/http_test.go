package httpkit_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/clouway/go-genproto/clouwayapis/rpc/errdetails"
	"github.com/clouway/go-genproto/clouwayapis/rpc/httpkit"
	"github.com/clouway/go-genproto/clouwayapis/rpc/request"
)

func TestEncodeHTTPGenericResponse(t *testing.T) {
	protoResponse := &errdetails.ErrorInfo{Reason: "Test Reason"}
	w := httptest.NewRecorder()
	httpkit.EncodeHTTPGenericResponse(context.Background(), w, protoResponse)
	b, _ := ioutil.ReadAll(w.Result().Body)
	body := string(b)
	want := `{"reason":"Test Reason","domain":"","metadata":{}}`

	if body != want {
		t.Errorf("unexpected response of EncodeHTTPGenericResponse:\n- want: %v\n-  got: %v", want, body)
	}
}

func TestEncodeHTTPGenericResponseWithEmptySlice(t *testing.T) {
	protoResponse := &errdetails.BadRequest{Errors: []*errdetails.BadRequest_FieldViolation{}}
	w := httptest.NewRecorder()
	httpkit.EncodeHTTPGenericResponse(context.Background(), w, protoResponse)

	b, _ := ioutil.ReadAll(w.Result().Body)
	body := string(b)
	want := `{"message":"","errors":[]}`

	if body != want {
		t.Errorf("unexpected response of EncodeHTTPGenericResponse:\n- want: %v\n-  got: %v", want, body)
	}
}

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
