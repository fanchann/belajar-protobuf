// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.14.0
// source: application.proto

package gen

import (
	first "gen/first/"
	second "gen/second/"
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

type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App         *first.FirstApplication `protobuf:"bytes,1,opt,name=App,proto3" json:"App,omitempty"`
	Description string                  `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_application_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetApp() *first.FirstApplication {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *Application) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type JobSoftware struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App         *second.SecondApplication `protobuf:"bytes,1,opt,name=App,proto3" json:"App,omitempty"`
	Description string                    `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *JobSoftware) Reset() {
	*x = JobSoftware{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobSoftware) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobSoftware) ProtoMessage() {}

func (x *JobSoftware) ProtoReflect() protoreflect.Message {
	mi := &file_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobSoftware.ProtoReflect.Descriptor instead.
func (*JobSoftware) Descriptor() ([]byte, []int) {
	return file_application_proto_rawDescGZIP(), []int{1}
}

func (x *JobSoftware) GetApp() *second.SecondApplication {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *JobSoftware) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_application_proto protoreflect.FileDescriptor

var file_application_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x0f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x10, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x41, 0x70,
	0x70, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x60, 0x0a, 0x0b, 0x4a, 0x6f, 0x62, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61,
	0x72, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03,
	0x41, 0x70, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_application_proto_rawDescOnce sync.Once
	file_application_proto_rawDescData = file_application_proto_rawDesc
)

func file_application_proto_rawDescGZIP() []byte {
	file_application_proto_rawDescOnce.Do(func() {
		file_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_application_proto_rawDescData)
	})
	return file_application_proto_rawDescData
}

var file_application_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_application_proto_goTypes = []interface{}{
	(*Application)(nil),              // 0: application.Application
	(*JobSoftware)(nil),              // 1: application.JobSoftware
	(*first.FirstApplication)(nil),   // 2: first.app.FirstApplication
	(*second.SecondApplication)(nil), // 3: second.app.SecondApplication
}
var file_application_proto_depIdxs = []int32{
	2, // 0: application.Application.App:type_name -> first.app.FirstApplication
	3, // 1: application.JobSoftware.App:type_name -> second.app.SecondApplication
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_application_proto_init() }
func file_application_proto_init() {
	if File_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobSoftware); i {
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
			RawDescriptor: file_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_application_proto_goTypes,
		DependencyIndexes: file_application_proto_depIdxs,
		MessageInfos:      file_application_proto_msgTypes,
	}.Build()
	File_application_proto = out.File
	file_application_proto_rawDesc = nil
	file_application_proto_goTypes = nil
	file_application_proto_depIdxs = nil
}
