package httpkit_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/clouway/go-genproto/clouwayapis/rpc/errdetails"
	"github.com/clouway/go-genproto/clouwayapis/rpc/httpkit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type input struct {
	code    codes.Code
	message string
	details *errdetails.BadRequest
}
type want struct {
	headers map[string]string
	status  int
	body    string
}

type test struct {
	input input
	want  want
}

func TestEncodeError(t *testing.T) {
	tests := []test{
		{
			input: input{code: codes.NotFound, message: "not found", details: nil},
			want: want{
				headers: map[string]string{"Content-Type": "application/json; charset=utf-8"},
				status:  http.StatusNotFound,
				body:    `{"message":"not found"}`,
			},
		},
		{
			input: input{code: codes.AlreadyExists, message: "already exists", details: nil},
			want: want{
				headers: map[string]string{"Content-Type": "application/json; charset=utf-8"},
				status:  http.StatusConflict,
				body:    `{"message":"already exists"}`,
			},
		},
		{
			input: input{code: codes.AlreadyExists, details: &errdetails.BadRequest{Message: "item already added"}},
			want: want{
				headers: map[string]string{"Content-Type": "application/json; charset=utf-8"},
				status:  http.StatusConflict,
				body:    `{"message":"item already added"}`,
			},
		},
		{
			input: input{
				code: codes.AlreadyExists,
				details: &errdetails.BadRequest{
					Message: "item already added",
					Errors: []*errdetails.BadRequest_FieldViolation{
						{Reason: "Item with id '123' already exists", Field: "itemId"},
					},
				},
			},
			want: want{
				headers: map[string]string{"Content-Type": "application/json; charset=utf-8"},
				status:  http.StatusConflict,
				body:    `{"message":"item already added","errors":[{"reason":"Item with id '123' already exists","field":"itemId"}]}`,
			},
		},
	}
	for _, test := range tests {
		st := status.New(test.input.code, test.input.message)
		if test.input.details != nil {
			st, _ = st.WithDetails(test.input.details)
		}

		rec := httptest.NewRecorder()
		httpkit.EncodeError(context.Background(), st.Err(), rec)

		for whn, whv := range test.want.headers {
			headerValue := rec.Header().Get(whn)
			if headerValue != whv {
				t.Errorf("unexpected header returned:\n- want: %v\n-  got: %v", whv, headerValue)
				return
			}
		}
		if rec.Result().StatusCode != test.want.status {
			t.Errorf("unexpected status code:\n- want: %v\n-  got: %v", test.want.status, rec.Result().StatusCode)
			return
		}
		body, _ := ioutil.ReadAll(rec.Body)
		sbody := string(body)
		if sbody != test.want.body {
			t.Errorf("unexpected body:\n- want: %v\n-  got: %v", test.want.body, sbody)
		}
	}

}
