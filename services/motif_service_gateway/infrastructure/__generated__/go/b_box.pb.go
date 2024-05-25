// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: b_box.proto

package motif_service

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

// Image to be processed by the BBoxService.
type ImageDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Source:
	//
	//	*ImageDto_Path
	//	*ImageDto_Base64Image
	Source isImageDto_Source `protobuf_oneof:"source"`
}

func (x *ImageDto) Reset() {
	*x = ImageDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_b_box_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageDto) ProtoMessage() {}

func (x *ImageDto) ProtoReflect() protoreflect.Message {
	mi := &file_b_box_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageDto.ProtoReflect.Descriptor instead.
func (*ImageDto) Descriptor() ([]byte, []int) {
	return file_b_box_proto_rawDescGZIP(), []int{0}
}

func (m *ImageDto) GetSource() isImageDto_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *ImageDto) GetPath() string {
	if x, ok := x.GetSource().(*ImageDto_Path); ok {
		return x.Path
	}
	return ""
}

func (x *ImageDto) GetBase64Image() *Base64ImageDto {
	if x, ok := x.GetSource().(*ImageDto_Base64Image); ok {
		return x.Base64Image
	}
	return nil
}

type isImageDto_Source interface {
	isImageDto_Source()
}

type ImageDto_Path struct {
	// The path to the image file on the server.
	Path string `protobuf:"bytes,1,opt,name=path,proto3,oneof"`
}

type ImageDto_Base64Image struct {
	// The image encoded in base64.
	Base64Image *Base64ImageDto `protobuf:"bytes,2,opt,name=base64_image,json=base64Image,proto3,oneof"`
}

func (*ImageDto_Path) isImageDto_Source() {}

func (*ImageDto_Base64Image) isImageDto_Source() {}

type Base64ImageDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The image data encoded in base64.
	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Base64ImageDto) Reset() {
	*x = Base64ImageDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_b_box_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Base64ImageDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Base64ImageDto) ProtoMessage() {}

func (x *Base64ImageDto) ProtoReflect() protoreflect.Message {
	mi := &file_b_box_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Base64ImageDto.ProtoReflect.Descriptor instead.
func (*Base64ImageDto) Descriptor() ([]byte, []int) {
	return file_b_box_proto_rawDescGZIP(), []int{1}
}

func (x *Base64ImageDto) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// Bounding box to encapsulate the an object in an image.
type BBoxDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The x pixel coordinate of the top left corner of the bounding box.
	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	// The y pixel coordinate of the top left corner of the bounding box.
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	// The width of the bounding box.
	Width int32 `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	// The height of the bounding box.
	Height int32 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *BBoxDto) Reset() {
	*x = BBoxDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_b_box_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BBoxDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BBoxDto) ProtoMessage() {}

func (x *BBoxDto) ProtoReflect() protoreflect.Message {
	mi := &file_b_box_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BBoxDto.ProtoReflect.Descriptor instead.
func (*BBoxDto) Descriptor() ([]byte, []int) {
	return file_b_box_proto_rawDescGZIP(), []int{2}
}

func (x *BBoxDto) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *BBoxDto) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *BBoxDto) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *BBoxDto) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type GetBBoxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *ImageDto `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GetBBoxRequest) Reset() {
	*x = GetBBoxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_b_box_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBBoxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBBoxRequest) ProtoMessage() {}

func (x *GetBBoxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_b_box_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBBoxRequest.ProtoReflect.Descriptor instead.
func (*GetBBoxRequest) Descriptor() ([]byte, []int) {
	return file_b_box_proto_rawDescGZIP(), []int{3}
}

func (x *GetBBoxRequest) GetImage() *ImageDto {
	if x != nil {
		return x.Image
	}
	return nil
}

type GetBBoxDebugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *ImageDto `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GetBBoxDebugRequest) Reset() {
	*x = GetBBoxDebugRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_b_box_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBBoxDebugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBBoxDebugRequest) ProtoMessage() {}

func (x *GetBBoxDebugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_b_box_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBBoxDebugRequest.ProtoReflect.Descriptor instead.
func (*GetBBoxDebugRequest) Descriptor() ([]byte, []int) {
	return file_b_box_proto_rawDescGZIP(), []int{4}
}

func (x *GetBBoxDebugRequest) GetImage() *ImageDto {
	if x != nil {
		return x.Image
	}
	return nil
}

type GetBBoxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BBox *BBoxDto `protobuf:"bytes,1,opt,name=b_box,json=bBox,proto3" json:"b_box,omitempty"`
}

func (x *GetBBoxResponse) Reset() {
	*x = GetBBoxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_b_box_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBBoxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBBoxResponse) ProtoMessage() {}

func (x *GetBBoxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_b_box_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBBoxResponse.ProtoReflect.Descriptor instead.
func (*GetBBoxResponse) Descriptor() ([]byte, []int) {
	return file_b_box_proto_rawDescGZIP(), []int{5}
}

func (x *GetBBoxResponse) GetBBox() *BBoxDto {
	if x != nil {
		return x.BBox
	}
	return nil
}

type GetBBoxDebugResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *Base64ImageDto `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GetBBoxDebugResponse) Reset() {
	*x = GetBBoxDebugResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_b_box_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBBoxDebugResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBBoxDebugResponse) ProtoMessage() {}

func (x *GetBBoxDebugResponse) ProtoReflect() protoreflect.Message {
	mi := &file_b_box_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBBoxDebugResponse.ProtoReflect.Descriptor instead.
func (*GetBBoxDebugResponse) Descriptor() ([]byte, []int) {
	return file_b_box_proto_rawDescGZIP(), []int{6}
}

func (x *GetBBoxDebugResponse) GetImage() *Base64ImageDto {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_b_box_proto protoreflect.FileDescriptor

var file_b_box_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x5f, 0x62, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x63, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x37, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x36, 0x34, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x48, 0x00, 0x52, 0x0b, 0x62,
	0x61, 0x73, 0x65, 0x36, 0x34, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x53, 0x0a, 0x07, 0x42, 0x42,
	0x6f, 0x78, 0x44, 0x74, 0x6f, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22,
	0x34, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x39, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x42, 0x6f, 0x78,
	0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x22, 0x33, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x62, 0x5f, 0x62, 0x6f, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x42, 0x6f, 0x78, 0x44, 0x74, 0x6f, 0x52,
	0x04, 0x62, 0x42, 0x6f, 0x78, 0x22, 0x40, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x42, 0x6f, 0x78,
	0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x32, 0xb9, 0x01, 0x0a, 0x0b, 0x42, 0x42, 0x6f, 0x78,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x5f, 0x62,
	0x5f, 0x62, 0x6f, 0x78, 0x12, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x42, 0x6f,
	0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x62,
	0x6f, 0x78, 0x12, 0x5f, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x5f, 0x62, 0x6f, 0x78, 0x5f,
	0x64, 0x65, 0x62, 0x75, 0x67, 0x12, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x42,
	0x6f, 0x78, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x62, 0x6f, 0x78, 0x2f, 0x64, 0x65,
	0x62, 0x75, 0x67, 0x42, 0x6f, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x42,
	0x42, 0x6f, 0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x63, 0x6f, 0x2d, 0x69, 0x2f, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x2d, 0x6f, 0x70, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x6d, 0x6f, 0x74, 0x69, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xa2, 0x02,
	0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02,
	0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_b_box_proto_rawDescOnce sync.Once
	file_b_box_proto_rawDescData = file_b_box_proto_rawDesc
)

func file_b_box_proto_rawDescGZIP() []byte {
	file_b_box_proto_rawDescOnce.Do(func() {
		file_b_box_proto_rawDescData = protoimpl.X.CompressGZIP(file_b_box_proto_rawDescData)
	})
	return file_b_box_proto_rawDescData
}

var file_b_box_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_b_box_proto_goTypes = []interface{}{
	(*ImageDto)(nil),             // 0: v1.ImageDto
	(*Base64ImageDto)(nil),       // 1: v1.Base64ImageDto
	(*BBoxDto)(nil),              // 2: v1.BBoxDto
	(*GetBBoxRequest)(nil),       // 3: v1.GetBBoxRequest
	(*GetBBoxDebugRequest)(nil),  // 4: v1.GetBBoxDebugRequest
	(*GetBBoxResponse)(nil),      // 5: v1.GetBBoxResponse
	(*GetBBoxDebugResponse)(nil), // 6: v1.GetBBoxDebugResponse
}
var file_b_box_proto_depIdxs = []int32{
	1, // 0: v1.ImageDto.base64_image:type_name -> v1.Base64ImageDto
	0, // 1: v1.GetBBoxRequest.image:type_name -> v1.ImageDto
	0, // 2: v1.GetBBoxDebugRequest.image:type_name -> v1.ImageDto
	2, // 3: v1.GetBBoxResponse.b_box:type_name -> v1.BBoxDto
	1, // 4: v1.GetBBoxDebugResponse.image:type_name -> v1.Base64ImageDto
	3, // 5: v1.BBoxService.get_b_box:input_type -> v1.GetBBoxRequest
	4, // 6: v1.BBoxService.get_b_box_debug:input_type -> v1.GetBBoxDebugRequest
	5, // 7: v1.BBoxService.get_b_box:output_type -> v1.GetBBoxResponse
	6, // 8: v1.BBoxService.get_b_box_debug:output_type -> v1.GetBBoxDebugResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_b_box_proto_init() }
func file_b_box_proto_init() {
	if File_b_box_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_b_box_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageDto); i {
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
		file_b_box_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Base64ImageDto); i {
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
		file_b_box_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BBoxDto); i {
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
		file_b_box_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBBoxRequest); i {
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
		file_b_box_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBBoxDebugRequest); i {
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
		file_b_box_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBBoxResponse); i {
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
		file_b_box_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBBoxDebugResponse); i {
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
	file_b_box_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ImageDto_Path)(nil),
		(*ImageDto_Base64Image)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_b_box_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_b_box_proto_goTypes,
		DependencyIndexes: file_b_box_proto_depIdxs,
		MessageInfos:      file_b_box_proto_msgTypes,
	}.Build()
	File_b_box_proto = out.File
	file_b_box_proto_rawDesc = nil
	file_b_box_proto_goTypes = nil
	file_b_box_proto_depIdxs = nil
}
