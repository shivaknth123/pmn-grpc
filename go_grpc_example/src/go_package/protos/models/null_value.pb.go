//
//Nudm_SDM
//
//Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
//
//The version of the OpenAPI document: 2.1.7
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.6.1
// source: null_value.proto

package models

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

type NullValue int32

const (
	NullValue_NullValue_NULL NullValue = 0
)

// Enum value maps for NullValue.
var (
	NullValue_name = map[int32]string{
		0: "NullValue_NULL",
	}
	NullValue_value = map[string]int32{
		"NullValue_NULL": 0,
	}
)

func (x NullValue) Enum() *NullValue {
	p := new(NullValue)
	*p = x
	return p
}

func (x NullValue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NullValue) Descriptor() protoreflect.EnumDescriptor {
	return file_null_value_proto_enumTypes[0].Descriptor()
}

func (NullValue) Type() protoreflect.EnumType {
	return &file_null_value_proto_enumTypes[0]
}

func (x NullValue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NullValue.Descriptor instead.
func (NullValue) EnumDescriptor() ([]byte, []int) {
	return file_null_value_proto_rawDescGZIP(), []int{0}
}

var File_null_value_proto protoreflect.FileDescriptor

var file_null_value_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x67, 0x6f, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2a, 0x1f, 0x0a, 0x09, 0x4e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x6f, 0x5f, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_null_value_proto_rawDescOnce sync.Once
	file_null_value_proto_rawDescData = file_null_value_proto_rawDesc
)

func file_null_value_proto_rawDescGZIP() []byte {
	file_null_value_proto_rawDescOnce.Do(func() {
		file_null_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_null_value_proto_rawDescData)
	})
	return file_null_value_proto_rawDescData
}

var file_null_value_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_null_value_proto_goTypes = []interface{}{
	(NullValue)(0), // 0: go_package.protos.NullValue
}
var file_null_value_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_null_value_proto_init() }
func file_null_value_proto_init() {
	if File_null_value_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_null_value_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_null_value_proto_goTypes,
		DependencyIndexes: file_null_value_proto_depIdxs,
		EnumInfos:         file_null_value_proto_enumTypes,
	}.Build()
	File_null_value_proto = out.File
	file_null_value_proto_rawDesc = nil
	file_null_value_proto_goTypes = nil
	file_null_value_proto_depIdxs = nil
}
