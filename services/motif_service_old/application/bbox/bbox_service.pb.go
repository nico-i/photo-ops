// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: application/bbox/bbox_service.proto

package bbox

import (
	value_objects "github.com/nico-i/photo-ops/services/motif_service/domain/value_objects"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_application_bbox_bbox_service_proto protoreflect.FileDescriptor

var file_application_bbox_bbox_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x62,
	0x6f, 0x78, 0x2f, 0x62, 0x62, 0x6f, 0x78, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x62, 0x62, 0x6f, 0x78, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x62, 0x62,
	0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd6, 0x01, 0x0a, 0x0b, 0x42, 0x42, 0x6f,
	0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42,
	0x42, 0x6f, 0x78, 0x12, 0x1b, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x1a, 0x1a, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x42, 0x42, 0x6f, 0x78, 0x22, 0x19, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x62, 0x6f,
	0x78, 0x2d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x68, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x42,
	0x6f, 0x78, 0x44, 0x65, 0x62, 0x75, 0x67, 0x12, 0x1b, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x1a, 0x1a, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x42, 0x42, 0x6f, 0x78,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x62, 0x6f, 0x78, 0x2d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x62, 0x75,
	0x67, 0x42, 0xce, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x62, 0x6f, 0x78, 0x42, 0x10, 0x42, 0x62, 0x6f, 0x78,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x63, 0x6f, 0x2d,
	0x69, 0x2f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2d, 0x6f, 0x70, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x74, 0x69, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62,
	0x62, 0x6f, 0x78, 0xa2, 0x02, 0x03, 0x41, 0x42, 0x58, 0xaa, 0x02, 0x10, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x62, 0x6f, 0x78, 0xca, 0x02, 0x10, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x42, 0x62, 0x6f, 0x78, 0xe2,
	0x02, 0x1c, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x42, 0x62,
	0x6f, 0x78, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x11, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x42, 0x62,
	0x6f, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_application_bbox_bbox_service_proto_goTypes = []interface{}{
	(*value_objects.Image)(nil), // 0: domain.value_objects.Image
	(*value_objects.BBox)(nil),  // 1: domain.value_objects.BBox
}
var file_application_bbox_bbox_service_proto_depIdxs = []int32{
	0, // 0: application.bbox.BBoxService.GetBBox:input_type -> domain.value_objects.Image
	0, // 1: application.bbox.BBoxService.GetBBoxDebug:input_type -> domain.value_objects.Image
	1, // 2: application.bbox.BBoxService.GetBBox:output_type -> domain.value_objects.BBox
	1, // 3: application.bbox.BBoxService.GetBBoxDebug:output_type -> domain.value_objects.BBox
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_application_bbox_bbox_service_proto_init() }
func file_application_bbox_bbox_service_proto_init() {
	if File_application_bbox_bbox_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_application_bbox_bbox_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_application_bbox_bbox_service_proto_goTypes,
		DependencyIndexes: file_application_bbox_bbox_service_proto_depIdxs,
	}.Build()
	File_application_bbox_bbox_service_proto = out.File
	file_application_bbox_bbox_service_proto_rawDesc = nil
	file_application_bbox_bbox_service_proto_goTypes = nil
	file_application_bbox_bbox_service_proto_depIdxs = nil
}