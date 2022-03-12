// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: admin_error.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type AdminErrorReason int32

const (
	AdminErrorReason_NOT_LOGGED_IN         AdminErrorReason = 0  // 401
	AdminErrorReason_ACCESS_FORBIDDEN      AdminErrorReason = 1  // 403
	AdminErrorReason_RESOURCE_NOT_FOUND    AdminErrorReason = 2  // 404
	AdminErrorReason_METHOD_NOT_ALLOWED    AdminErrorReason = 3  // 405
	AdminErrorReason_REQUEST_TIMEOUT       AdminErrorReason = 4  // 408
	AdminErrorReason_INTERNAL_SERVER_ERROR AdminErrorReason = 5  // 500
	AdminErrorReason_NOT_IMPLEMENTED       AdminErrorReason = 6  // 501
	AdminErrorReason_NETWORK_ERROR         AdminErrorReason = 7  // 502
	AdminErrorReason_SERVICE_UNAVAILABLE   AdminErrorReason = 8  // 503
	AdminErrorReason_NETWORK_TIMEOUT       AdminErrorReason = 9  // 504
	AdminErrorReason_REQUEST_NOT_SUPPORT   AdminErrorReason = 10 // 505
	AdminErrorReason_USER_NOT_FOUND        AdminErrorReason = 11
)

// Enum value maps for AdminErrorReason.
var (
	AdminErrorReason_name = map[int32]string{
		0:  "NOT_LOGGED_IN",
		1:  "ACCESS_FORBIDDEN",
		2:  "RESOURCE_NOT_FOUND",
		3:  "METHOD_NOT_ALLOWED",
		4:  "REQUEST_TIMEOUT",
		5:  "INTERNAL_SERVER_ERROR",
		6:  "NOT_IMPLEMENTED",
		7:  "NETWORK_ERROR",
		8:  "SERVICE_UNAVAILABLE",
		9:  "NETWORK_TIMEOUT",
		10: "REQUEST_NOT_SUPPORT",
		11: "USER_NOT_FOUND",
	}
	AdminErrorReason_value = map[string]int32{
		"NOT_LOGGED_IN":         0,
		"ACCESS_FORBIDDEN":      1,
		"RESOURCE_NOT_FOUND":    2,
		"METHOD_NOT_ALLOWED":    3,
		"REQUEST_TIMEOUT":       4,
		"INTERNAL_SERVER_ERROR": 5,
		"NOT_IMPLEMENTED":       6,
		"NETWORK_ERROR":         7,
		"SERVICE_UNAVAILABLE":   8,
		"NETWORK_TIMEOUT":       9,
		"REQUEST_NOT_SUPPORT":   10,
		"USER_NOT_FOUND":        11,
	}
)

func (x AdminErrorReason) Enum() *AdminErrorReason {
	p := new(AdminErrorReason)
	*p = x
	return p
}

func (x AdminErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdminErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_admin_error_proto_enumTypes[0].Descriptor()
}

func (AdminErrorReason) Type() protoreflect.EnumType {
	return &file_admin_error_proto_enumTypes[0]
}

func (x AdminErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdminErrorReason.Descriptor instead.
func (AdminErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_admin_error_proto_rawDescGZIP(), []int{0}
}

var File_admin_error_proto protoreflect.FileDescriptor

var file_admin_error_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0xec, 0x02, 0x0a, 0x10, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x4c,
	0x4f, 0x47, 0x47, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03,
	0x12, 0x1a, 0x0a, 0x10, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x46, 0x4f, 0x52, 0x42, 0x49,
	0x44, 0x44, 0x45, 0x4e, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x1c, 0x0a, 0x12,
	0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1c, 0x0a, 0x12, 0x4d, 0x45,
	0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44,
	0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x95, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x04, 0x1a, 0x04, 0xa8,
	0x45, 0x98, 0x03, 0x12, 0x1f, 0x0a, 0x15, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x1a, 0x04,
	0xa8, 0x45, 0xf4, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4d, 0x50, 0x4c,
	0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0xf5, 0x03, 0x12,
	0x17, 0x0a, 0x0d, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x07, 0x1a, 0x04, 0xa8, 0x45, 0xf6, 0x03, 0x12, 0x1d, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x08, 0x1a, 0x04, 0xa8, 0x45, 0xf7, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x4e, 0x45, 0x54, 0x57, 0x4f,
	0x52, 0x4b, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x09, 0x1a, 0x04, 0xa8, 0x45,
	0xf8, 0x03, 0x12, 0x1d, 0x0a, 0x13, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x0a, 0x1a, 0x04, 0xa8, 0x45, 0xf9,
	0x03, 0x12, 0x18, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x0b, 0x1a, 0x04, 0xa8, 0x45, 0xd8, 0x04, 0x1a, 0x04, 0xa0, 0x45, 0xf4,
	0x03, 0x42, 0x0f, 0x50, 0x01, 0x5a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_error_proto_rawDescOnce sync.Once
	file_admin_error_proto_rawDescData = file_admin_error_proto_rawDesc
)

func file_admin_error_proto_rawDescGZIP() []byte {
	file_admin_error_proto_rawDescOnce.Do(func() {
		file_admin_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_error_proto_rawDescData)
	})
	return file_admin_error_proto_rawDescData
}

var file_admin_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_admin_error_proto_goTypes = []interface{}{
	(AdminErrorReason)(0), // 0: admin.v1.AdminErrorReason
}
var file_admin_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_error_proto_init() }
func file_admin_error_proto_init() {
	if File_admin_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_error_proto_goTypes,
		DependencyIndexes: file_admin_error_proto_depIdxs,
		EnumInfos:         file_admin_error_proto_enumTypes,
	}.Build()
	File_admin_error_proto = out.File
	file_admin_error_proto_rawDesc = nil
	file_admin_error_proto_goTypes = nil
	file_admin_error_proto_depIdxs = nil
}
