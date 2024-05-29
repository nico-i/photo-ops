// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: services/motif_service/v1/motif_service.proto

package motif_servicev1

import (
	v12 "github.com/nico-i/photo-ops/shared/go/__generated__/proto/messages/base64_image_dto/v1"
	v1 "github.com/nico-i/photo-ops/shared/go/__generated__/proto/messages/image_dto/v1"
	v11 "github.com/nico-i/photo-ops/shared/go/__generated__/proto/messages/rect_dto/v1"
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

type GetBBoxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *v1.ImageDto `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GetBBoxRequest) Reset() {
	*x = GetBBoxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_motif_service_v1_motif_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBBoxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBBoxRequest) ProtoMessage() {}

func (x *GetBBoxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_motif_service_v1_motif_service_proto_msgTypes[0]
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
	return file_services_motif_service_v1_motif_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetBBoxRequest) GetImage() *v1.ImageDto {
	if x != nil {
		return x.Image
	}
	return nil
}

type GetBBoxDebugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *v1.ImageDto `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GetBBoxDebugRequest) Reset() {
	*x = GetBBoxDebugRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_motif_service_v1_motif_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBBoxDebugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBBoxDebugRequest) ProtoMessage() {}

func (x *GetBBoxDebugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_motif_service_v1_motif_service_proto_msgTypes[1]
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
	return file_services_motif_service_v1_motif_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetBBoxDebugRequest) GetImage() *v1.ImageDto {
	if x != nil {
		return x.Image
	}
	return nil
}

type GetBBoxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BBox *v11.RectDto `protobuf:"bytes,1,opt,name=b_box,json=bBox,proto3" json:"b_box,omitempty"`
}

func (x *GetBBoxResponse) Reset() {
	*x = GetBBoxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_motif_service_v1_motif_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBBoxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBBoxResponse) ProtoMessage() {}

func (x *GetBBoxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_motif_service_v1_motif_service_proto_msgTypes[2]
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
	return file_services_motif_service_v1_motif_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetBBoxResponse) GetBBox() *v11.RectDto {
	if x != nil {
		return x.BBox
	}
	return nil
}

type GetBBoxDebugResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *v12.Base64ImageDto `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GetBBoxDebugResponse) Reset() {
	*x = GetBBoxDebugResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_motif_service_v1_motif_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBBoxDebugResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBBoxDebugResponse) ProtoMessage() {}

func (x *GetBBoxDebugResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_motif_service_v1_motif_service_proto_msgTypes[3]
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
	return file_services_motif_service_v1_motif_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetBBoxDebugResponse) GetImage() *v12.Base64ImageDto {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_services_motif_service_v1_motif_service_proto protoreflect.FileDescriptor

var file_services_motif_service_v1_motif_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x74, 0x69, 0x66,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x74, 0x69,
	0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x6d, 0x6f, 0x74, 0x69, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x33, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x36, 0x34, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x22, 0x53, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x4c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x42, 0x6f, 0x78,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x62, 0x5f, 0x62, 0x6f,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x64,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x74, 0x44, 0x74, 0x6f, 0x52, 0x04, 0x62,
	0x42, 0x6f, 0x78, 0x22, 0x61, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x44, 0x65,
	0x62, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x36, 0x34, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x52,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x32, 0xf2, 0x01, 0x0a, 0x0c, 0x4d, 0x6f, 0x74, 0x69, 0x66,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x5f, 0x62,
	0x5f, 0x62, 0x6f, 0x78, 0x12, 0x20, 0x2e, 0x6d, 0x6f, 0x74, 0x69, 0x66, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x6f, 0x74, 0x69, 0x66, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x42, 0x6f,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x62, 0x6f, 0x78, 0x12, 0x7b,
	0x0a, 0x0f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x64, 0x65, 0x62, 0x75,
	0x67, 0x12, 0x25, 0x2e, 0x6d, 0x6f, 0x74, 0x69, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x44, 0x65, 0x62, 0x75,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x6f, 0x74, 0x69, 0x66,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x42, 0x6f, 0x78, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x62, 0x6f, 0x78, 0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x42, 0xeb, 0x01, 0x0a, 0x14,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x74, 0x69, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x4d, 0x6f, 0x74, 0x69, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x63, 0x6f, 0x2d, 0x69, 0x2f, 0x70, 0x68, 0x6f,
	0x74, 0x6f, 0x2d, 0x6f, 0x70, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x67, 0x6f,
	0x2f, 0x5f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x5f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x6f,
	0x74, 0x69, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d,
	0x6f, 0x74, 0x69, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x4d, 0x6f, 0x74, 0x69, 0x66, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x4d, 0x6f, 0x74, 0x69, 0x66, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x4d, 0x6f, 0x74, 0x69, 0x66,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x4d, 0x6f, 0x74, 0x69, 0x66, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_services_motif_service_v1_motif_service_proto_rawDescOnce sync.Once
	file_services_motif_service_v1_motif_service_proto_rawDescData = file_services_motif_service_v1_motif_service_proto_rawDesc
)

func file_services_motif_service_v1_motif_service_proto_rawDescGZIP() []byte {
	file_services_motif_service_v1_motif_service_proto_rawDescOnce.Do(func() {
		file_services_motif_service_v1_motif_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_motif_service_v1_motif_service_proto_rawDescData)
	})
	return file_services_motif_service_v1_motif_service_proto_rawDescData
}

var file_services_motif_service_v1_motif_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_services_motif_service_v1_motif_service_proto_goTypes = []interface{}{
	(*GetBBoxRequest)(nil),       // 0: motif_service.v1.GetBBoxRequest
	(*GetBBoxDebugRequest)(nil),  // 1: motif_service.v1.GetBBoxDebugRequest
	(*GetBBoxResponse)(nil),      // 2: motif_service.v1.GetBBoxResponse
	(*GetBBoxDebugResponse)(nil), // 3: motif_service.v1.GetBBoxDebugResponse
	(*v1.ImageDto)(nil),          // 4: shared.messages.image_dto.v1.ImageDto
	(*v11.RectDto)(nil),          // 5: shared.messages.rect_dto.v1.RectDto
	(*v12.Base64ImageDto)(nil),   // 6: shared.messages.base64_image_dto.v1.Base64ImageDto
}
var file_services_motif_service_v1_motif_service_proto_depIdxs = []int32{
	4, // 0: motif_service.v1.GetBBoxRequest.image:type_name -> shared.messages.image_dto.v1.ImageDto
	4, // 1: motif_service.v1.GetBBoxDebugRequest.image:type_name -> shared.messages.image_dto.v1.ImageDto
	5, // 2: motif_service.v1.GetBBoxResponse.b_box:type_name -> shared.messages.rect_dto.v1.RectDto
	6, // 3: motif_service.v1.GetBBoxDebugResponse.image:type_name -> shared.messages.base64_image_dto.v1.Base64ImageDto
	0, // 4: motif_service.v1.MotifService.get_b_box:input_type -> motif_service.v1.GetBBoxRequest
	1, // 5: motif_service.v1.MotifService.get_b_box_debug:input_type -> motif_service.v1.GetBBoxDebugRequest
	2, // 6: motif_service.v1.MotifService.get_b_box:output_type -> motif_service.v1.GetBBoxResponse
	3, // 7: motif_service.v1.MotifService.get_b_box_debug:output_type -> motif_service.v1.GetBBoxDebugResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_services_motif_service_v1_motif_service_proto_init() }
func file_services_motif_service_v1_motif_service_proto_init() {
	if File_services_motif_service_v1_motif_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_motif_service_v1_motif_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_motif_service_v1_motif_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_motif_service_v1_motif_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_motif_service_v1_motif_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_motif_service_v1_motif_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_motif_service_v1_motif_service_proto_goTypes,
		DependencyIndexes: file_services_motif_service_v1_motif_service_proto_depIdxs,
		MessageInfos:      file_services_motif_service_v1_motif_service_proto_msgTypes,
	}.Build()
	File_services_motif_service_v1_motif_service_proto = out.File
	file_services_motif_service_v1_motif_service_proto_rawDesc = nil
	file_services_motif_service_v1_motif_service_proto_goTypes = nil
	file_services_motif_service_v1_motif_service_proto_depIdxs = nil
}