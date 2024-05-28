// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: services/metadata_service/v1/metadata_service.proto

package metadata_servicev1

import (
	v1 "github.com/nico-i/photo-ops/shared/go/__generated__/proto/messages/image_dto/v1"
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

// Request to get caption for an image
type UpdateIPTCDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Image which
	Image      *v1.ImageDto `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Caption    string       `protobuf:"bytes,2,opt,name=caption,proto3" json:"caption,omitempty"`
	HashtagCsv string       `protobuf:"bytes,3,opt,name=hashtag_csv,json=hashtagCsv,proto3" json:"hashtag_csv,omitempty"`
}

func (x *UpdateIPTCDataRequest) Reset() {
	*x = UpdateIPTCDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_metadata_service_v1_metadata_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIPTCDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIPTCDataRequest) ProtoMessage() {}

func (x *UpdateIPTCDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_metadata_service_v1_metadata_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIPTCDataRequest.ProtoReflect.Descriptor instead.
func (*UpdateIPTCDataRequest) Descriptor() ([]byte, []int) {
	return file_services_metadata_service_v1_metadata_service_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateIPTCDataRequest) GetImage() *v1.ImageDto {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *UpdateIPTCDataRequest) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

func (x *UpdateIPTCDataRequest) GetHashtagCsv() string {
	if x != nil {
		return x.HashtagCsv
	}
	return ""
}

type UpdateIPTCDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateIPTCDataResponse) Reset() {
	*x = UpdateIPTCDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_metadata_service_v1_metadata_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIPTCDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIPTCDataResponse) ProtoMessage() {}

func (x *UpdateIPTCDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_metadata_service_v1_metadata_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIPTCDataResponse.ProtoReflect.Descriptor instead.
func (*UpdateIPTCDataResponse) Descriptor() ([]byte, []int) {
	return file_services_metadata_service_v1_metadata_service_proto_rawDescGZIP(), []int{1}
}

var File_services_metadata_service_v1_metadata_service_proto protoreflect.FileDescriptor

var file_services_metadata_service_v1_metadata_service_proto_rawDesc = []byte{
	0x0a, 0x33, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x90, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x50, 0x54, 0x43, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61, 0x67, 0x5f, 0x63, 0x73, 0x76,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61, 0x67, 0x43,
	0x73, 0x76, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x50, 0x54, 0x43,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x94, 0x01, 0x0a,
	0x0f, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x80, 0x01, 0x0a, 0x10, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x70, 0x74, 0x63,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x50, 0x54, 0x43, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x50,
	0x54, 0x43, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x70, 0x74, 0x63, 0x42, 0x83, 0x02, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x14, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x69, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x63, 0x6f, 0x2d, 0x69, 0x2f, 0x70, 0x68, 0x6f, 0x74, 0x6f,
	0x2d, 0x6f, 0x70, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x5f,
	0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x5f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x12, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x1e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_services_metadata_service_v1_metadata_service_proto_rawDescOnce sync.Once
	file_services_metadata_service_v1_metadata_service_proto_rawDescData = file_services_metadata_service_v1_metadata_service_proto_rawDesc
)

func file_services_metadata_service_v1_metadata_service_proto_rawDescGZIP() []byte {
	file_services_metadata_service_v1_metadata_service_proto_rawDescOnce.Do(func() {
		file_services_metadata_service_v1_metadata_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_metadata_service_v1_metadata_service_proto_rawDescData)
	})
	return file_services_metadata_service_v1_metadata_service_proto_rawDescData
}

var file_services_metadata_service_v1_metadata_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_metadata_service_v1_metadata_service_proto_goTypes = []interface{}{
	(*UpdateIPTCDataRequest)(nil),  // 0: metadata_service.v1.UpdateIPTCDataRequest
	(*UpdateIPTCDataResponse)(nil), // 1: metadata_service.v1.UpdateIPTCDataResponse
	(*v1.ImageDto)(nil),            // 2: shared.messages.image_dto.v1.ImageDto
}
var file_services_metadata_service_v1_metadata_service_proto_depIdxs = []int32{
	2, // 0: metadata_service.v1.UpdateIPTCDataRequest.image:type_name -> shared.messages.image_dto.v1.ImageDto
	0, // 1: metadata_service.v1.MetaDataService.update_iptc_data:input_type -> metadata_service.v1.UpdateIPTCDataRequest
	1, // 2: metadata_service.v1.MetaDataService.update_iptc_data:output_type -> metadata_service.v1.UpdateIPTCDataResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_services_metadata_service_v1_metadata_service_proto_init() }
func file_services_metadata_service_v1_metadata_service_proto_init() {
	if File_services_metadata_service_v1_metadata_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_metadata_service_v1_metadata_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIPTCDataRequest); i {
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
		file_services_metadata_service_v1_metadata_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIPTCDataResponse); i {
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
			RawDescriptor: file_services_metadata_service_v1_metadata_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_metadata_service_v1_metadata_service_proto_goTypes,
		DependencyIndexes: file_services_metadata_service_v1_metadata_service_proto_depIdxs,
		MessageInfos:      file_services_metadata_service_v1_metadata_service_proto_msgTypes,
	}.Build()
	File_services_metadata_service_v1_metadata_service_proto = out.File
	file_services_metadata_service_v1_metadata_service_proto_rawDesc = nil
	file_services_metadata_service_v1_metadata_service_proto_goTypes = nil
	file_services_metadata_service_v1_metadata_service_proto_depIdxs = nil
}
