// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: proto/second.proto

package ch205_protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SecondMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Number  *wrapperspb.Int32Value  `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	Address *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Cell    *wrapperspb.BytesValue  `protobuf:"bytes,5,opt,name=cell,proto3" json:"cell,omitempty"`
}

func (x *SecondMsg) Reset() {
	*x = SecondMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_second_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecondMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecondMsg) ProtoMessage() {}

func (x *SecondMsg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_second_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecondMsg.ProtoReflect.Descriptor instead.
func (*SecondMsg) Descriptor() ([]byte, []int) {
	return file_proto_second_proto_rawDescGZIP(), []int{0}
}

func (x *SecondMsg) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SecondMsg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecondMsg) GetNumber() *wrapperspb.Int32Value {
	if x != nil {
		return x.Number
	}
	return nil
}

func (x *SecondMsg) GetAddress() *wrapperspb.StringValue {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *SecondMsg) GetCell() *wrapperspb.BytesValue {
	if x != nil {
		return x.Cell
	}
	return nil
}

var File_proto_second_proto protoreflect.FileDescriptor

var file_proto_second_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x68, 0x32, 0x30, 0x35, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x09, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4d,
	0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x2f, 0x0a, 0x04, 0x63, 0x65, 0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04,
	0x63, 0x65, 0x6c, 0x6c, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x3b, 0x63, 0x68, 0x32, 0x30, 0x35,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_second_proto_rawDescOnce sync.Once
	file_proto_second_proto_rawDescData = file_proto_second_proto_rawDesc
)

func file_proto_second_proto_rawDescGZIP() []byte {
	file_proto_second_proto_rawDescOnce.Do(func() {
		file_proto_second_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_second_proto_rawDescData)
	})
	return file_proto_second_proto_rawDescData
}

var file_proto_second_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_second_proto_goTypes = []interface{}{
	(*SecondMsg)(nil),              // 0: ch205_protobuf.SecondMsg
	(*wrapperspb.Int32Value)(nil),  // 1: google.protobuf.Int32Value
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
	(*wrapperspb.BytesValue)(nil),  // 3: google.protobuf.BytesValue
}
var file_proto_second_proto_depIdxs = []int32{
	1, // 0: ch205_protobuf.SecondMsg.number:type_name -> google.protobuf.Int32Value
	2, // 1: ch205_protobuf.SecondMsg.address:type_name -> google.protobuf.StringValue
	3, // 2: ch205_protobuf.SecondMsg.cell:type_name -> google.protobuf.BytesValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_second_proto_init() }
func file_proto_second_proto_init() {
	if File_proto_second_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_second_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecondMsg); i {
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
			RawDescriptor: file_proto_second_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_second_proto_goTypes,
		DependencyIndexes: file_proto_second_proto_depIdxs,
		MessageInfos:      file_proto_second_proto_msgTypes,
	}.Build()
	File_proto_second_proto = out.File
	file_proto_second_proto_rawDesc = nil
	file_proto_second_proto_goTypes = nil
	file_proto_second_proto_depIdxs = nil
}