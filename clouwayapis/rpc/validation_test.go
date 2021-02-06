package rpc_test

import (
	"reflect"
	"testing"

	"github.com/clouway/go-genproto/clouwayapis/rpc"
	"github.com/clouway/go-genproto/clouwayapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestValidation(t *testing.T) {
	err := rpc.ValidateRequest(&sampleMessage{Name: ""})
	if err == nil {
		t.Errorf("expected BadRequest error from rpc.ValidateRequest, but got nil")
	}
	st, _ := status.FromError(err)

	if st.Code() != codes.InvalidArgument {
		t.Errorf("unexpected BadRequest code:\n- want: %v\n-  got: %v", codes.InvalidArgument, st.Code())
	}
	got := st.Details()[0].(*errdetails.BadRequest)
	want := []*errdetails.BadRequest_FieldViolation{
		{
			Reason: "name should be between 5 and 10",
			Field:  "name",
		},
	}

	if expMsg := "name should be between 5 and 10"; expMsg != st.Message() {
		t.Errorf("unexpected error message from rpc.ValidateRequest:\n- want: %v\n-  got: %v", expMsg, st.Message())
		return
	}

	if !reflect.DeepEqual(want, got.Errors) {
		t.Errorf("unexpected error details from rpc.ValidateRequest:\n- want: %v\n-  got: %v", want, got.Errors)
		return
	}
}

type sampleMessage struct {
	Name string
}

func (s *sampleMessage) Validate() error {
	if len(s.Name) == 0 {
		return &fieldError{
			field: "Name", reason: "name should be between 5 and 10",
		}
	}
	return nil
}

type fieldError struct {
	field  string
	reason string
	key    string
}

func (f *fieldError) Error() string {
	return f.reason
}

func (f *fieldError) Field() string {
	return f.field
}
func (f *fieldError) Reason() string {
	return f.reason
}

func (f *fieldError) Key() bool {
	return true
}

func (f *fieldError) Cause() error {
	return nil
}

func (f *fieldError) ErrorName() string {
	return ""
}
