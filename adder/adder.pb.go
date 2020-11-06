// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: adder/adder.proto

package adder

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type NumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *NumberRequest) Reset() {
	*x = NumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adder_adder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberRequest) ProtoMessage() {}

func (x *NumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adder_adder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberRequest.ProtoReflect.Descriptor instead.
func (*NumberRequest) Descriptor() ([]byte, []int) {
	return file_adder_adder_proto_rawDescGZIP(), []int{0}
}

func (x *NumberRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type NumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num    int32  `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *NumberResponse) Reset() {
	*x = NumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adder_adder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberResponse) ProtoMessage() {}

func (x *NumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_adder_adder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberResponse.ProtoReflect.Descriptor instead.
func (*NumberResponse) Descriptor() ([]byte, []int) {
	return file_adder_adder_proto_rawDescGZIP(), []int{1}
}

func (x *NumberResponse) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *NumberResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_adder_adder_proto protoreflect.FileDescriptor

var file_adder_adder_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x64, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x64, 0x65, 0x72, 0x22, 0x21, 0x0a, 0x0d, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e,
	0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x3a, 0x0a,
	0x0e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75,
	0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x43, 0x0a, 0x05, 0x41, 0x64, 0x64,
	0x65, 0x72, 0x12, 0x3a, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x14, 0x2e, 0x61, 0x64, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x64, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x24,
	0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6a, 0x61,
	0x66, 0x61, 0x72, 0x69, 0x39, 0x38, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61,
	0x64, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_adder_adder_proto_rawDescOnce sync.Once
	file_adder_adder_proto_rawDescData = file_adder_adder_proto_rawDesc
)

func file_adder_adder_proto_rawDescGZIP() []byte {
	file_adder_adder_proto_rawDescOnce.Do(func() {
		file_adder_adder_proto_rawDescData = protoimpl.X.CompressGZIP(file_adder_adder_proto_rawDescData)
	})
	return file_adder_adder_proto_rawDescData
}

var file_adder_adder_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_adder_adder_proto_goTypes = []interface{}{
	(*NumberRequest)(nil),  // 0: adder.NumberRequest
	(*NumberResponse)(nil), // 1: adder.NumberResponse
}
var file_adder_adder_proto_depIdxs = []int32{
	0, // 0: adder.Adder.AddNumber:input_type -> adder.NumberRequest
	1, // 1: adder.Adder.AddNumber:output_type -> adder.NumberResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_adder_adder_proto_init() }
func file_adder_adder_proto_init() {
	if File_adder_adder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_adder_adder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumberRequest); i {
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
		file_adder_adder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumberResponse); i {
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
			RawDescriptor: file_adder_adder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_adder_adder_proto_goTypes,
		DependencyIndexes: file_adder_adder_proto_depIdxs,
		MessageInfos:      file_adder_adder_proto_msgTypes,
	}.Build()
	File_adder_adder_proto = out.File
	file_adder_adder_proto_rawDesc = nil
	file_adder_adder_proto_goTypes = nil
	file_adder_adder_proto_depIdxs = nil
}
