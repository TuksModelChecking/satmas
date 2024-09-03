// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: mra/algorithm.proto

package proto

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

type SynthesisAlgorithm int32

const (
	SynthesisAlgorithm_COLLECTIVE             SynthesisAlgorithm = 0
	SynthesisAlgorithm_NASHEQUILIBRIUM        SynthesisAlgorithm = 1
	SynthesisAlgorithm_EPSILONNASHEQUILIBRIUM SynthesisAlgorithm = 2
)

// Enum value maps for SynthesisAlgorithm.
var (
	SynthesisAlgorithm_name = map[int32]string{
		0: "COLLECTIVE",
		1: "NASHEQUILIBRIUM",
		2: "EPSILONNASHEQUILIBRIUM",
	}
	SynthesisAlgorithm_value = map[string]int32{
		"COLLECTIVE":             0,
		"NASHEQUILIBRIUM":        1,
		"EPSILONNASHEQUILIBRIUM": 2,
	}
)

func (x SynthesisAlgorithm) Enum() *SynthesisAlgorithm {
	p := new(SynthesisAlgorithm)
	*p = x
	return p
}

func (x SynthesisAlgorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SynthesisAlgorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_mra_algorithm_proto_enumTypes[0].Descriptor()
}

func (SynthesisAlgorithm) Type() protoreflect.EnumType {
	return &file_mra_algorithm_proto_enumTypes[0]
}

func (x SynthesisAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SynthesisAlgorithm.Descriptor instead.
func (SynthesisAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_mra_algorithm_proto_rawDescGZIP(), []int{0}
}

var File_mra_algorithm_proto protoreflect.FileDescriptor

var file_mra_algorithm_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x72, 0x61, 0x2f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x72, 0x61, 0x2a, 0x55, 0x0a, 0x12, 0x53, 0x79,
	0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x73, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x00,
	0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x41, 0x53, 0x48, 0x45, 0x51, 0x55, 0x49, 0x4c, 0x49, 0x42, 0x52,
	0x49, 0x55, 0x4d, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x50, 0x53, 0x49, 0x4c, 0x4f, 0x4e,
	0x4e, 0x41, 0x53, 0x48, 0x45, 0x51, 0x55, 0x49, 0x4c, 0x49, 0x42, 0x52, 0x49, 0x55, 0x4d, 0x10,
	0x02, 0x42, 0x0a, 0x5a, 0x08, 0x75, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mra_algorithm_proto_rawDescOnce sync.Once
	file_mra_algorithm_proto_rawDescData = file_mra_algorithm_proto_rawDesc
)

func file_mra_algorithm_proto_rawDescGZIP() []byte {
	file_mra_algorithm_proto_rawDescOnce.Do(func() {
		file_mra_algorithm_proto_rawDescData = protoimpl.X.CompressGZIP(file_mra_algorithm_proto_rawDescData)
	})
	return file_mra_algorithm_proto_rawDescData
}

var file_mra_algorithm_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mra_algorithm_proto_goTypes = []any{
	(SynthesisAlgorithm)(0), // 0: mra.SynthesisAlgorithm
}
var file_mra_algorithm_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mra_algorithm_proto_init() }
func file_mra_algorithm_proto_init() {
	if File_mra_algorithm_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mra_algorithm_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mra_algorithm_proto_goTypes,
		DependencyIndexes: file_mra_algorithm_proto_depIdxs,
		EnumInfos:         file_mra_algorithm_proto_enumTypes,
	}.Build()
	File_mra_algorithm_proto = out.File
	file_mra_algorithm_proto_rawDesc = nil
	file_mra_algorithm_proto_goTypes = nil
	file_mra_algorithm_proto_depIdxs = nil
}