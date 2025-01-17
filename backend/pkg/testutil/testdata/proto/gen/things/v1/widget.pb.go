// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: things/v1/widget.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Widget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Widget) Reset() {
	*x = Widget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_things_v1_widget_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Widget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Widget) ProtoMessage() {}

func (x *Widget) ProtoReflect() protoreflect.Message {
	mi := &file_things_v1_widget_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Widget.ProtoReflect.Descriptor instead.
func (*Widget) Descriptor() ([]byte, []int) {
	return file_things_v1_widget_proto_rawDescGZIP(), []int{0}
}

func (x *Widget) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Widget) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_things_v1_widget_proto protoreflect.FileDescriptor

var file_things_v1_widget_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x73, 0x74, 0x75, 0x74,
	0x69, 0x6c, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a,
	0x06, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_things_v1_widget_proto_rawDescOnce sync.Once
	file_things_v1_widget_proto_rawDescData = file_things_v1_widget_proto_rawDesc
)

func file_things_v1_widget_proto_rawDescGZIP() []byte {
	file_things_v1_widget_proto_rawDescOnce.Do(func() {
		file_things_v1_widget_proto_rawDescData = protoimpl.X.CompressGZIP(file_things_v1_widget_proto_rawDescData)
	})
	return file_things_v1_widget_proto_rawDescData
}

var file_things_v1_widget_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_things_v1_widget_proto_goTypes = []interface{}{
	(*Widget)(nil),                // 0: testutil.things.v1.Widget
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_things_v1_widget_proto_depIdxs = []int32{
	1, // 0: testutil.things.v1.Widget.created_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_things_v1_widget_proto_init() }
func file_things_v1_widget_proto_init() {
	if File_things_v1_widget_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_things_v1_widget_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Widget); i {
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
			RawDescriptor: file_things_v1_widget_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_things_v1_widget_proto_goTypes,
		DependencyIndexes: file_things_v1_widget_proto_depIdxs,
		MessageInfos:      file_things_v1_widget_proto_msgTypes,
	}.Build()
	File_things_v1_widget_proto = out.File
	file_things_v1_widget_proto_rawDesc = nil
	file_things_v1_widget_proto_goTypes = nil
	file_things_v1_widget_proto_depIdxs = nil
}
