// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: motif/motif.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The request message containing either the image file path or image data.
type ImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Source:
	//
	//	*ImageRequest_FilePath
	//	*ImageRequest_ImageData
	Source isImageRequest_Source `protobuf_oneof:"source"`
}

func (x *ImageRequest) Reset() {
	*x = ImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motif_motif_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageRequest) ProtoMessage() {}

func (x *ImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_motif_motif_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageRequest.ProtoReflect.Descriptor instead.
func (*ImageRequest) Descriptor() ([]byte, []int) {
	return file_motif_motif_proto_rawDescGZIP(), []int{0}
}

func (m *ImageRequest) GetSource() isImageRequest_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *ImageRequest) GetFilePath() string {
	if x, ok := x.GetSource().(*ImageRequest_FilePath); ok {
		return x.FilePath
	}
	return ""
}

func (x *ImageRequest) GetImageData() []byte {
	if x, ok := x.GetSource().(*ImageRequest_ImageData); ok {
		return x.ImageData
	}
	return nil
}

type isImageRequest_Source interface {
	isImageRequest_Source()
}

type ImageRequest_FilePath struct {
	FilePath string `protobuf:"bytes,1,opt,name=filePath,proto3,oneof"`
}

type ImageRequest_ImageData struct {
	ImageData []byte `protobuf:"bytes,2,opt,name=imageData,proto3,oneof"`
}

func (*ImageRequest_FilePath) isImageRequest_Source() {}

func (*ImageRequest_ImageData) isImageRequest_Source() {}

// The response message containing bounding box coordinates.
type BoundingBoxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X      int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y      int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Width  int32 `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height int32 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *BoundingBoxResponse) Reset() {
	*x = BoundingBoxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motif_motif_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoundingBoxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoundingBoxResponse) ProtoMessage() {}

func (x *BoundingBoxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_motif_motif_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoundingBoxResponse.ProtoReflect.Descriptor instead.
func (*BoundingBoxResponse) Descriptor() ([]byte, []int) {
	return file_motif_motif_proto_rawDescGZIP(), []int{1}
}

func (x *BoundingBoxResponse) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *BoundingBoxResponse) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *BoundingBoxResponse) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *BoundingBoxResponse) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

// The response message containing the processed image data.
type ImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The image data as bytes.
	ImageData []byte `protobuf:"bytes,1,opt,name=imageData,proto3" json:"imageData,omitempty"`
}

func (x *ImageResponse) Reset() {
	*x = ImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motif_motif_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageResponse) ProtoMessage() {}

func (x *ImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_motif_motif_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageResponse.ProtoReflect.Descriptor instead.
func (*ImageResponse) Descriptor() ([]byte, []int) {
	return file_motif_motif_proto_rawDescGZIP(), []int{2}
}

func (x *ImageResponse) GetImageData() []byte {
	if x != nil {
		return x.ImageData
	}
	return nil
}

var File_motif_motif_proto protoreflect.FileDescriptor

var file_motif_motif_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x74, 0x69, 0x66, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0x5f, 0x0a, 0x13, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x22, 0x2d, 0x0a, 0x0d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x32, 0xbf, 0x01, 0x0a, 0x0c, 0x4d, 0x6f, 0x74, 0x69, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x42, 0x6f, 0x78, 0x12, 0x13, 0x2e, 0x6d, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x6f, 0x74, 0x69, 0x66,
	0x2e, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22,
	0x10, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2d, 0x62, 0x6f,
	0x78, 0x12, 0x4f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x42, 0x47, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x13, 0x2e, 0x6d, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x2d,
	0x62, 0x67, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6e, 0x69, 0x63, 0x6f, 0x2d, 0x69, 0x2f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2d, 0x6f, 0x70,
	0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x74, 0x69, 0x66,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_motif_motif_proto_rawDescOnce sync.Once
	file_motif_motif_proto_rawDescData = file_motif_motif_proto_rawDesc
)

func file_motif_motif_proto_rawDescGZIP() []byte {
	file_motif_motif_proto_rawDescOnce.Do(func() {
		file_motif_motif_proto_rawDescData = protoimpl.X.CompressGZIP(file_motif_motif_proto_rawDescData)
	})
	return file_motif_motif_proto_rawDescData
}

var file_motif_motif_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_motif_motif_proto_goTypes = []interface{}{
	(*ImageRequest)(nil),        // 0: motif.ImageRequest
	(*BoundingBoxResponse)(nil), // 1: motif.BoundingBoxResponse
	(*ImageResponse)(nil),       // 2: motif.ImageResponse
}
var file_motif_motif_proto_depIdxs = []int32{
	0, // 0: motif.MotifService.GetBoundingBox:input_type -> motif.ImageRequest
	0, // 1: motif.MotifService.GetNoBGImage:input_type -> motif.ImageRequest
	1, // 2: motif.MotifService.GetBoundingBox:output_type -> motif.BoundingBoxResponse
	2, // 3: motif.MotifService.GetNoBGImage:output_type -> motif.ImageResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_motif_motif_proto_init() }
func file_motif_motif_proto_init() {
	if File_motif_motif_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_motif_motif_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageRequest); i {
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
		file_motif_motif_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoundingBoxResponse); i {
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
		file_motif_motif_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageResponse); i {
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
	file_motif_motif_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ImageRequest_FilePath)(nil),
		(*ImageRequest_ImageData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_motif_motif_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_motif_motif_proto_goTypes,
		DependencyIndexes: file_motif_motif_proto_depIdxs,
		MessageInfos:      file_motif_motif_proto_msgTypes,
	}.Build()
	File_motif_motif_proto = out.File
	file_motif_motif_proto_rawDesc = nil
	file_motif_motif_proto_goTypes = nil
	file_motif_motif_proto_depIdxs = nil
}
