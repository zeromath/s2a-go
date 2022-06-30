// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: internal/proto/v2/common.proto

package common_go_proto

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

// The ciphersuites for the TLS handshake supported by S2A. Note that handshake
// ciphersuites are NOT configurable in TLS 1.3.
type HandshakeCiphersuite int32

const (
	HandshakeCiphersuite_HANDSHAKE_CIPHERSUITE_UNSPECIFIED HandshakeCiphersuite = 0
	HandshakeCiphersuite_HANDSHAKE_CIPHERSUITE_ECDHE       HandshakeCiphersuite = 1
)

// Enum value maps for HandshakeCiphersuite.
var (
	HandshakeCiphersuite_name = map[int32]string{
		0: "HANDSHAKE_CIPHERSUITE_UNSPECIFIED",
		1: "HANDSHAKE_CIPHERSUITE_ECDHE",
	}
	HandshakeCiphersuite_value = map[string]int32{
		"HANDSHAKE_CIPHERSUITE_UNSPECIFIED": 0,
		"HANDSHAKE_CIPHERSUITE_ECDHE":       1,
	}
)

func (x HandshakeCiphersuite) Enum() *HandshakeCiphersuite {
	p := new(HandshakeCiphersuite)
	*p = x
	return p
}

func (x HandshakeCiphersuite) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HandshakeCiphersuite) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_proto_v2_common_proto_enumTypes[0].Descriptor()
}

func (HandshakeCiphersuite) Type() protoreflect.EnumType {
	return &file_internal_proto_v2_common_proto_enumTypes[0]
}

func (x HandshakeCiphersuite) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HandshakeCiphersuite.Descriptor instead.
func (HandshakeCiphersuite) EnumDescriptor() ([]byte, []int) {
	return file_internal_proto_v2_common_proto_rawDescGZIP(), []int{0}
}

// The ciphersuites for the record protocol supported by S2A. The name
// determines the confidentiality, and authentication ciphers as well as the
// hash algorithm used for PRF in TLS 1.2 or HKDF in TLS 1.3. Thus, the
// components of the name are:
//   - AEAD -- for encryption and authentication, e.g., AES_128_GCM.
//   - Hash algorithm -- used in PRF or HKDF, e.g., SHA256.
type RecordCiphersuite int32

const (
	RecordCiphersuite_RECORD_CIPHERSUITE_UNSPECIFIED              RecordCiphersuite = 0
	RecordCiphersuite_RECORD_CIPHERSUITE_AES_128_GCM_SHA256       RecordCiphersuite = 1
	RecordCiphersuite_RECORD_CIPHERSUITE_AES_256_GCM_SHA384       RecordCiphersuite = 2
	RecordCiphersuite_RECORD_CIPHERSUITE_CHACHA20_POLY1305_SHA256 RecordCiphersuite = 3
)

// Enum value maps for RecordCiphersuite.
var (
	RecordCiphersuite_name = map[int32]string{
		0: "RECORD_CIPHERSUITE_UNSPECIFIED",
		1: "RECORD_CIPHERSUITE_AES_128_GCM_SHA256",
		2: "RECORD_CIPHERSUITE_AES_256_GCM_SHA384",
		3: "RECORD_CIPHERSUITE_CHACHA20_POLY1305_SHA256",
	}
	RecordCiphersuite_value = map[string]int32{
		"RECORD_CIPHERSUITE_UNSPECIFIED":              0,
		"RECORD_CIPHERSUITE_AES_128_GCM_SHA256":       1,
		"RECORD_CIPHERSUITE_AES_256_GCM_SHA384":       2,
		"RECORD_CIPHERSUITE_CHACHA20_POLY1305_SHA256": 3,
	}
)

func (x RecordCiphersuite) Enum() *RecordCiphersuite {
	p := new(RecordCiphersuite)
	*p = x
	return p
}

func (x RecordCiphersuite) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecordCiphersuite) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_proto_v2_common_proto_enumTypes[1].Descriptor()
}

func (RecordCiphersuite) Type() protoreflect.EnumType {
	return &file_internal_proto_v2_common_proto_enumTypes[1]
}

func (x RecordCiphersuite) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecordCiphersuite.Descriptor instead.
func (RecordCiphersuite) EnumDescriptor() ([]byte, []int) {
	return file_internal_proto_v2_common_proto_rawDescGZIP(), []int{1}
}

// The TLS versions supported by S2A's handshaker module.
type TLSVersion int32

const (
	TLSVersion_TLS_VERSION_UNSPECIFIED TLSVersion = 0
	TLSVersion_TLS_VERSION_1_0         TLSVersion = 1
	TLSVersion_TLS_VERSION_1_1         TLSVersion = 2
	TLSVersion_TLS_VERSION_1_2         TLSVersion = 3
	TLSVersion_TLS_VERSION_1_3         TLSVersion = 4
)

// Enum value maps for TLSVersion.
var (
	TLSVersion_name = map[int32]string{
		0: "TLS_VERSION_UNSPECIFIED",
		1: "TLS_VERSION_1_0",
		2: "TLS_VERSION_1_1",
		3: "TLS_VERSION_1_2",
		4: "TLS_VERSION_1_3",
	}
	TLSVersion_value = map[string]int32{
		"TLS_VERSION_UNSPECIFIED": 0,
		"TLS_VERSION_1_0":         1,
		"TLS_VERSION_1_1":         2,
		"TLS_VERSION_1_2":         3,
		"TLS_VERSION_1_3":         4,
	}
)

func (x TLSVersion) Enum() *TLSVersion {
	p := new(TLSVersion)
	*p = x
	return p
}

func (x TLSVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TLSVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_proto_v2_common_proto_enumTypes[2].Descriptor()
}

func (TLSVersion) Type() protoreflect.EnumType {
	return &file_internal_proto_v2_common_proto_enumTypes[2]
}

func (x TLSVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TLSVersion.Descriptor instead.
func (TLSVersion) EnumDescriptor() ([]byte, []int) {
	return file_internal_proto_v2_common_proto_rawDescGZIP(), []int{2}
}

// The side in the TLS connection.
type ConnectionSide int32

const (
	ConnectionSide_CONNECTION_SIDE_UNSPECIFIED ConnectionSide = 0
	ConnectionSide_CONNECTION_SIDE_CLIENT      ConnectionSide = 1
	ConnectionSide_CONNECTION_SIDE_SERVER      ConnectionSide = 2
)

// Enum value maps for ConnectionSide.
var (
	ConnectionSide_name = map[int32]string{
		0: "CONNECTION_SIDE_UNSPECIFIED",
		1: "CONNECTION_SIDE_CLIENT",
		2: "CONNECTION_SIDE_SERVER",
	}
	ConnectionSide_value = map[string]int32{
		"CONNECTION_SIDE_UNSPECIFIED": 0,
		"CONNECTION_SIDE_CLIENT":      1,
		"CONNECTION_SIDE_SERVER":      2,
	}
)

func (x ConnectionSide) Enum() *ConnectionSide {
	p := new(ConnectionSide)
	*p = x
	return p
}

func (x ConnectionSide) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionSide) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_proto_v2_common_proto_enumTypes[3].Descriptor()
}

func (ConnectionSide) Type() protoreflect.EnumType {
	return &file_internal_proto_v2_common_proto_enumTypes[3]
}

func (x ConnectionSide) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionSide.Descriptor instead.
func (ConnectionSide) EnumDescriptor() ([]byte, []int) {
	return file_internal_proto_v2_common_proto_rawDescGZIP(), []int{3}
}

var File_internal_proto_v2_common_proto protoreflect.FileDescriptor

var file_internal_proto_v2_common_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x73, 0x32, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x32, 0x2a, 0x5e,
	0x0a, 0x14, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x43, 0x69, 0x70, 0x68, 0x65,
	0x72, 0x73, 0x75, 0x69, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x21, 0x48, 0x41, 0x4e, 0x44, 0x53, 0x48,
	0x41, 0x4b, 0x45, 0x5f, 0x43, 0x49, 0x50, 0x48, 0x45, 0x52, 0x53, 0x55, 0x49, 0x54, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a,
	0x1b, 0x48, 0x41, 0x4e, 0x44, 0x53, 0x48, 0x41, 0x4b, 0x45, 0x5f, 0x43, 0x49, 0x50, 0x48, 0x45,
	0x52, 0x53, 0x55, 0x49, 0x54, 0x45, 0x5f, 0x45, 0x43, 0x44, 0x48, 0x45, 0x10, 0x01, 0x2a, 0xbe,
	0x01, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x73,
	0x75, 0x69, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x1e, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x43,
	0x49, 0x50, 0x48, 0x45, 0x52, 0x53, 0x55, 0x49, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x29, 0x0a, 0x25, 0x52, 0x45, 0x43, 0x4f,
	0x52, 0x44, 0x5f, 0x43, 0x49, 0x50, 0x48, 0x45, 0x52, 0x53, 0x55, 0x49, 0x54, 0x45, 0x5f, 0x41,
	0x45, 0x53, 0x5f, 0x31, 0x32, 0x38, 0x5f, 0x47, 0x43, 0x4d, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35,
	0x36, 0x10, 0x01, 0x12, 0x29, 0x0a, 0x25, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x43, 0x49,
	0x50, 0x48, 0x45, 0x52, 0x53, 0x55, 0x49, 0x54, 0x45, 0x5f, 0x41, 0x45, 0x53, 0x5f, 0x32, 0x35,
	0x36, 0x5f, 0x47, 0x43, 0x4d, 0x5f, 0x53, 0x48, 0x41, 0x33, 0x38, 0x34, 0x10, 0x02, 0x12, 0x2f,
	0x0a, 0x2b, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x43, 0x49, 0x50, 0x48, 0x45, 0x52, 0x53,
	0x55, 0x49, 0x54, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x43, 0x48, 0x41, 0x32, 0x30, 0x5f, 0x50, 0x4f,
	0x4c, 0x59, 0x31, 0x33, 0x30, 0x35, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x03, 0x2a,
	0x7d, 0x0a, 0x0a, 0x54, 0x4c, 0x53, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x17, 0x54, 0x4c, 0x53, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x4c,
	0x53, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x31, 0x5f, 0x30, 0x10, 0x01, 0x12,
	0x13, 0x0a, 0x0f, 0x54, 0x4c, 0x53, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x31,
	0x5f, 0x31, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x4c, 0x53, 0x5f, 0x56, 0x45, 0x52, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x31, 0x5f, 0x32, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x4c, 0x53,
	0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x31, 0x5f, 0x33, 0x10, 0x04, 0x2a, 0x69,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x64, 0x65,
	0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x49, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x49, 0x44, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x1a, 0x0a,
	0x16, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x44, 0x45,
	0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x02, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73,
	0x32, 0x61, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x67, 0x6f, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_v2_common_proto_rawDescOnce sync.Once
	file_internal_proto_v2_common_proto_rawDescData = file_internal_proto_v2_common_proto_rawDesc
)

func file_internal_proto_v2_common_proto_rawDescGZIP() []byte {
	file_internal_proto_v2_common_proto_rawDescOnce.Do(func() {
		file_internal_proto_v2_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_v2_common_proto_rawDescData)
	})
	return file_internal_proto_v2_common_proto_rawDescData
}

var file_internal_proto_v2_common_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_internal_proto_v2_common_proto_goTypes = []interface{}{
	(HandshakeCiphersuite)(0), // 0: s2a.proto.v2.HandshakeCiphersuite
	(RecordCiphersuite)(0),    // 1: s2a.proto.v2.RecordCiphersuite
	(TLSVersion)(0),           // 2: s2a.proto.v2.TLSVersion
	(ConnectionSide)(0),       // 3: s2a.proto.v2.ConnectionSide
}
var file_internal_proto_v2_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_proto_v2_common_proto_init() }
func file_internal_proto_v2_common_proto_init() {
	if File_internal_proto_v2_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_proto_v2_common_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_proto_v2_common_proto_goTypes,
		DependencyIndexes: file_internal_proto_v2_common_proto_depIdxs,
		EnumInfos:         file_internal_proto_v2_common_proto_enumTypes,
	}.Build()
	File_internal_proto_v2_common_proto = out.File
	file_internal_proto_v2_common_proto_rawDesc = nil
	file_internal_proto_v2_common_proto_goTypes = nil
	file_internal_proto_v2_common_proto_depIdxs = nil
}
