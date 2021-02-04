// Copyright 2021 clouWay eood.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: clouway/rpc/errdetails/error_details.proto

package errdetails

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Describes how a quota check failed.
//
// For example if a daily limit was exceeded for the calling project,
// a service could respond with a QuotaFailure detail containing the tenant
// id and the description of the quota limit that was exceeded.  If the
// calling tenant hasn't enabled the service in the console, then
// a service could respond with the tenant and set `service_disabled`
// to true.
type QuotaFailure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Describes all quota violations.
	Violations []*QuotaFailure_Violation `protobuf:"bytes,1,rep,name=violations,proto3" json:"violations,omitempty"`
}

func (x *QuotaFailure) Reset() {
	*x = QuotaFailure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clouway_rpc_errdetails_error_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuotaFailure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuotaFailure) ProtoMessage() {}

func (x *QuotaFailure) ProtoReflect() protoreflect.Message {
	mi := &file_clouway_rpc_errdetails_error_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuotaFailure.ProtoReflect.Descriptor instead.
func (*QuotaFailure) Descriptor() ([]byte, []int) {
	return file_clouway_rpc_errdetails_error_details_proto_rawDescGZIP(), []int{0}
}

func (x *QuotaFailure) GetViolations() []*QuotaFailure_Violation {
	if x != nil {
		return x.Violations
	}
	return nil
}

// BadRequest is an error used to indicate bad requests.
type BadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Describes all violations in a client request.
	FieldViolations []*BadRequest_FieldViolation `protobuf:"bytes,1,rep,name=field_violations,json=fieldViolations,proto3" json:"field_violations,omitempty"`
}

func (x *BadRequest) Reset() {
	*x = BadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clouway_rpc_errdetails_error_details_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadRequest) ProtoMessage() {}

func (x *BadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clouway_rpc_errdetails_error_details_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadRequest.ProtoReflect.Descriptor instead.
func (*BadRequest) Descriptor() ([]byte, []int) {
	return file_clouway_rpc_errdetails_error_details_proto_rawDescGZIP(), []int{1}
}

func (x *BadRequest) GetFieldViolations() []*BadRequest_FieldViolation {
	if x != nil {
		return x.FieldViolations
	}
	return nil
}

// Describes the cause of the error with structured details.
//
// Example of an error when contacting the "pubsub.googleapis.com" API when it
// is not enabled:
//
//     { "reason": "API_DISABLED"
//       "domain": "api.politis.com"
//       "metadata": {
//         "resource": "tenant/12312312",
//         "service": "api.politis.com"
//       }
//     }
//
// This response indicates that the api.politis.com API is not enabled.
type ErrorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The reason of the error. This is a constant value that identifies the
	// proximate cause of the error. Error reasons are unique within a particular
	// domain of errors. This should be at most 63 characters and match
	// /[A-Z0-9_]+/.
	Reason string `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
	// The logical grouping to which the "reason" belongs. The error domain
	// is typically the registered service name of the tool or product that
	// generates the error. Example: "pubsub.googleapis.com". If the error is
	// generated by some common infrastructure, the error domain must be a
	// globally unique value that identifies the infrastructure. For Google API
	// infrastructure, the error domain is "googleapis.com".
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// Additional structured details about this error.
	//
	// Keys should match /[a-zA-Z0-9-_]/ and be limited to 64 characters in
	// length. When identifying the current value of an exceeded limit, the units
	// should be contained in the key, not the value.  For example, rather than
	// {"instanceLimit": "100/request"}, should be returned as,
	// {"instanceLimitPerRequest": "100"}, if the client exceeds the number of
	// instances that can be created in a single (batch) request.
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ErrorInfo) Reset() {
	*x = ErrorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clouway_rpc_errdetails_error_details_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorInfo) ProtoMessage() {}

func (x *ErrorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_clouway_rpc_errdetails_error_details_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorInfo.ProtoReflect.Descriptor instead.
func (*ErrorInfo) Descriptor() ([]byte, []int) {
	return file_clouway_rpc_errdetails_error_details_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorInfo) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *ErrorInfo) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ErrorInfo) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// A message type used to describe a single quota violation.  For example, a
// daily quota or a custom quota that was exceeded.
type QuotaFailure_Violation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The subject on which the quota check failed.
	// For example, "clientip:<ip address of client>" or "tenant:<Your tenant>".
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	// A description of how the quota check failed. Clients can use this
	// description to find more about the quota configuration in the service's
	// public documentation, or find the relevant quota limit to adjust through
	// developer console.
	//
	// For example: "Service disabled" or "Daily Limit for read operations
	// exceeded".
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *QuotaFailure_Violation) Reset() {
	*x = QuotaFailure_Violation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clouway_rpc_errdetails_error_details_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuotaFailure_Violation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuotaFailure_Violation) ProtoMessage() {}

func (x *QuotaFailure_Violation) ProtoReflect() protoreflect.Message {
	mi := &file_clouway_rpc_errdetails_error_details_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuotaFailure_Violation.ProtoReflect.Descriptor instead.
func (*QuotaFailure_Violation) Descriptor() ([]byte, []int) {
	return file_clouway_rpc_errdetails_error_details_proto_rawDescGZIP(), []int{0, 0}
}

func (x *QuotaFailure_Violation) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *QuotaFailure_Violation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// A message type used to describe a single bad request field.
type BadRequest_FieldViolation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The reason of the violation
	Reason string `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
	// The field that is affected
	Field string `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	// The code of the error, e.g business errors like "eik_not_found"
	Code string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	// The message of the error
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BadRequest_FieldViolation) Reset() {
	*x = BadRequest_FieldViolation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clouway_rpc_errdetails_error_details_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BadRequest_FieldViolation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadRequest_FieldViolation) ProtoMessage() {}

func (x *BadRequest_FieldViolation) ProtoReflect() protoreflect.Message {
	mi := &file_clouway_rpc_errdetails_error_details_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadRequest_FieldViolation.ProtoReflect.Descriptor instead.
func (*BadRequest_FieldViolation) Descriptor() ([]byte, []int) {
	return file_clouway_rpc_errdetails_error_details_proto_rawDescGZIP(), []int{1, 0}
}

func (x *BadRequest_FieldViolation) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *BadRequest_FieldViolation) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *BadRequest_FieldViolation) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *BadRequest_FieldViolation) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_clouway_rpc_errdetails_error_details_proto protoreflect.FileDescriptor

var file_clouway_rpc_errdetails_error_details_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6c, 0x6f, 0x75, 0x77, 0x61, 0x79, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x72,
	0x72, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a,
	0x0c, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x37, 0x0a,
	0x0a, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x2e, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x76, 0x69, 0x6f, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x47, 0x0a, 0x09, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xc1, 0x01, 0x0a, 0x0a, 0x42, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45,
	0x0a, 0x10, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x42, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x69, 0x6f, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x69, 0x6f, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x6c, 0x0a, 0x0e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x69,
	0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0xae, 0x01, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x34, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x84, 0x01, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x72, 0x72, 0x72, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x42, 0x14, 0x52, 0x70, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x63, 0x6c, 0x6f, 0x75, 0x77, 0x61,
	0x79, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x72, 0x72, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x3b, 0x65, 0x72, 0x72, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_clouway_rpc_errdetails_error_details_proto_rawDescOnce sync.Once
	file_clouway_rpc_errdetails_error_details_proto_rawDescData = file_clouway_rpc_errdetails_error_details_proto_rawDesc
)

func file_clouway_rpc_errdetails_error_details_proto_rawDescGZIP() []byte {
	file_clouway_rpc_errdetails_error_details_proto_rawDescOnce.Do(func() {
		file_clouway_rpc_errdetails_error_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_clouway_rpc_errdetails_error_details_proto_rawDescData)
	})
	return file_clouway_rpc_errdetails_error_details_proto_rawDescData
}

var file_clouway_rpc_errdetails_error_details_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_clouway_rpc_errdetails_error_details_proto_goTypes = []interface{}{
	(*QuotaFailure)(nil),              // 0: QuotaFailure
	(*BadRequest)(nil),                // 1: BadRequest
	(*ErrorInfo)(nil),                 // 2: ErrorInfo
	(*QuotaFailure_Violation)(nil),    // 3: QuotaFailure.Violation
	(*BadRequest_FieldViolation)(nil), // 4: BadRequest.FieldViolation
	nil,                               // 5: ErrorInfo.MetadataEntry
}
var file_clouway_rpc_errdetails_error_details_proto_depIdxs = []int32{
	3, // 0: QuotaFailure.violations:type_name -> QuotaFailure.Violation
	4, // 1: BadRequest.field_violations:type_name -> BadRequest.FieldViolation
	5, // 2: ErrorInfo.metadata:type_name -> ErrorInfo.MetadataEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_clouway_rpc_errdetails_error_details_proto_init() }
func file_clouway_rpc_errdetails_error_details_proto_init() {
	if File_clouway_rpc_errdetails_error_details_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clouway_rpc_errdetails_error_details_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuotaFailure); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clouway_rpc_errdetails_error_details_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BadRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clouway_rpc_errdetails_error_details_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clouway_rpc_errdetails_error_details_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuotaFailure_Violation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clouway_rpc_errdetails_error_details_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BadRequest_FieldViolation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_clouway_rpc_errdetails_error_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_clouway_rpc_errdetails_error_details_proto_goTypes,
		DependencyIndexes: file_clouway_rpc_errdetails_error_details_proto_depIdxs,
		MessageInfos:      file_clouway_rpc_errdetails_error_details_proto_msgTypes,
	}.Build()
	File_clouway_rpc_errdetails_error_details_proto = out.File
	file_clouway_rpc_errdetails_error_details_proto_rawDesc = nil
	file_clouway_rpc_errdetails_error_details_proto_goTypes = nil
	file_clouway_rpc_errdetails_error_details_proto_depIdxs = nil
}
