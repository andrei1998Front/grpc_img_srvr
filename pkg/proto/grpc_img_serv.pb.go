// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0--rc2
// source: grpc_img_serv.proto

package grpc_img_server

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ImgUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Chunk    []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *ImgUploadRequest) Reset() {
	*x = ImgUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_img_serv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImgUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImgUploadRequest) ProtoMessage() {}

func (x *ImgUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_img_serv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImgUploadRequest.ProtoReflect.Descriptor instead.
func (*ImgUploadRequest) Descriptor() ([]byte, []int) {
	return file_grpc_img_serv_proto_rawDescGZIP(), []int{0}
}

func (x *ImgUploadRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *ImgUploadRequest) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type ImgInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string                 `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Size     uint32                 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	CreateDt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_dt,json=createDt,proto3" json:"create_dt,omitempty"`
	UpdateDt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_dt,json=updateDt,proto3" json:"update_dt,omitempty"`
}

func (x *ImgInfo) Reset() {
	*x = ImgInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_img_serv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImgInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImgInfo) ProtoMessage() {}

func (x *ImgInfo) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_img_serv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImgInfo.ProtoReflect.Descriptor instead.
func (*ImgInfo) Descriptor() ([]byte, []int) {
	return file_grpc_img_serv_proto_rawDescGZIP(), []int{1}
}

func (x *ImgInfo) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ImgInfo) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ImgInfo) GetCreateDt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateDt
	}
	return nil
}

func (x *ImgInfo) GetUpdateDt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateDt
	}
	return nil
}

var File_grpc_img_serv_proto protoreflect.FileDescriptor

var file_grpc_img_serv_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x69, 0x6d, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x69, 0x6d, 0x67, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x10, 0x49, 0x6d, 0x67, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0xac, 0x01,
	0x0a, 0x07, 0x49, 0x6d, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x74, 0x32, 0x55, 0x0a, 0x0a,
	0x49, 0x6d, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x69, 0x6d, 0x67, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x49, 0x6d, 0x67, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x69,
	0x6d, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x49, 0x6d, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x28, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x69, 0x6d,
	0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_img_serv_proto_rawDescOnce sync.Once
	file_grpc_img_serv_proto_rawDescData = file_grpc_img_serv_proto_rawDesc
)

func file_grpc_img_serv_proto_rawDescGZIP() []byte {
	file_grpc_img_serv_proto_rawDescOnce.Do(func() {
		file_grpc_img_serv_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_img_serv_proto_rawDescData)
	})
	return file_grpc_img_serv_proto_rawDescData
}

var file_grpc_img_serv_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_img_serv_proto_goTypes = []interface{}{
	(*ImgUploadRequest)(nil),      // 0: grpc_img_server.ImgUploadRequest
	(*ImgInfo)(nil),               // 1: grpc_img_server.ImgInfo
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_grpc_img_serv_proto_depIdxs = []int32{
	2, // 0: grpc_img_server.ImgInfo.create_dt:type_name -> google.protobuf.Timestamp
	2, // 1: grpc_img_server.ImgInfo.update_dt:type_name -> google.protobuf.Timestamp
	0, // 2: grpc_img_server.ImgService.Upload:input_type -> grpc_img_server.ImgUploadRequest
	1, // 3: grpc_img_server.ImgService.Upload:output_type -> grpc_img_server.ImgInfo
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpc_img_serv_proto_init() }
func file_grpc_img_serv_proto_init() {
	if File_grpc_img_serv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_img_serv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImgUploadRequest); i {
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
		file_grpc_img_serv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImgInfo); i {
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
			RawDescriptor: file_grpc_img_serv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_img_serv_proto_goTypes,
		DependencyIndexes: file_grpc_img_serv_proto_depIdxs,
		MessageInfos:      file_grpc_img_serv_proto_msgTypes,
	}.Build()
	File_grpc_img_serv_proto = out.File
	file_grpc_img_serv_proto_rawDesc = nil
	file_grpc_img_serv_proto_goTypes = nil
	file_grpc_img_serv_proto_depIdxs = nil
}
