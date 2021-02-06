package grpckit

// A generic integration with protoc-gen-validate that generates helper code
// for validation of protobuf messages. The internal interface of this package
// are emulating the validation results and is returning validation errors as
// status errors.

import (
	"strings"

	"github.com/clouway/go-genproto/clouwayapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidateRequest validates the incoming request and returns error object
// that encodes the error details. The returned result is nil of no error is
// found.
func ValidateRequest(req interface{}) error {
	f, ok := req.(validator)
	if !ok {
		return nil
	}
	err := f.Validate()
	if err == nil {
		return nil
	}

	verr, ok := err.(validationError)
	if !ok {
		return nil
	}

	st := status.New(codes.InvalidArgument, verr.Reason())
	errorDetails := &errdetails.BadRequest{
		Message: st.Message(),
		Errors: []*errdetails.BadRequest_FieldViolation{
			{Reason: verr.Reason(), Field: lowerFirst(verr.Field())},
		},
	}
	st, _ = st.WithDetails(errorDetails)
	return st.Err()
}

func lowerFirst(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

type validator interface {
	Validate() error
}

type validationError interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
}
