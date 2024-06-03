# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: services/crop_service/v1/crop_service.proto
# Protobuf Python Version: 5.27.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    0,
    '',
    'services/crop_service/v1/crop_service.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from messages.image_dto.v1 import image_dto_pb2 as messages_dot_image__dto_dot_v1_dot_image__dto__pb2
from messages.rect_dto.v1 import rect_dto_pb2 as messages_dot_rect__dto_dot_v1_dot_rect__dto__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n+services/crop_service/v1/crop_service.proto\x12\x0f\x63rop_service.v1\x1a\x1cgoogle/api/annotations.proto\x1a%messages/image_dto/v1/image_dto.proto\x1a#messages/rect_dto/v1/rect_dto.proto\"\x9e\x01\n\x10\x43ropImageRequest\x12G\n\x0binput_image\x18\x01 \x01(\x0b\x32&.shared.messages.image_dto.v1.ImageDtoR\ninputImage\x12\x41\n\tcrop_spec\x18\x02 \x01(\x0b\x32$.shared.messages.rect_dto.v1.RectDtoR\x08\x63ropSpec\"`\n\x11\x43ropImageResponse\x12K\n\rcropped_image\x18\x01 \x01(\x0b\x32&.shared.messages.image_dto.v1.ImageDtoR\x0c\x63roppedImage\"\xff\x01\n\x18GetCropSuggestionRequest\x12\x45\n\x0bimage_motif\x18\x01 \x01(\x0b\x32$.shared.messages.rect_dto.v1.RectDtoR\nimageMotif\x12\x1f\n\x0bimage_width\x18\x02 \x01(\rR\nimageWidth\x12!\n\x0cimage_height\x18\x03 \x01(\rR\x0bimageHeight\x12*\n\x11target_crop_width\x18\x04 \x01(\rR\x0ftargetCropWidth\x12,\n\x12target_crop_height\x18\x05 \x01(\rR\x10targetCropHeight\"j\n\x19GetCropSuggestionResponse\x12M\n\x0f\x63rop_suggestion\x18\x01 \x01(\x0b\x32$.shared.messages.rect_dto.v1.RectDtoR\x0e\x63ropSuggestion2\x86\x02\n\x0b\x43ropService\x12h\n\ncrop_image\x12!.crop_service.v1.CropImageRequest\x1a\".crop_service.v1.CropImageResponse\"\x13\x82\xd3\xe4\x93\x02\r\"\x08/v1/crop:\x01*\x12\x8c\x01\n\x13get_crop_suggestion\x12).crop_service.v1.GetCropSuggestionRequest\x1a*.crop_service.v1.GetCropSuggestionResponse\"\x1e\x82\xd3\xe4\x93\x02\x18\"\x13/v1/crop/suggestion:\x01*B\xe3\x01\n\x13\x63om.crop_service.v1B\x10\x43ropServiceProtoP\x01Zagithub.com/nico-i/photo-ops/shared/go/__generated__/proto/services/crop_service/v1;crop_servicev1\xa2\x02\x03\x43XX\xaa\x02\x0e\x43ropService.V1\xca\x02\x0e\x43ropService\\V1\xe2\x02\x1a\x43ropService\\V1\\GPBMetadata\xea\x02\x0f\x43ropService::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'services.crop_service.v1.crop_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\023com.crop_service.v1B\020CropServiceProtoP\001Zagithub.com/nico-i/photo-ops/shared/go/__generated__/proto/services/crop_service/v1;crop_servicev1\242\002\003CXX\252\002\016CropService.V1\312\002\016CropService\\V1\342\002\032CropService\\V1\\GPBMetadata\352\002\017CropService::V1'
  _globals['_CROPSERVICE'].methods_by_name['crop_image']._loaded_options = None
  _globals['_CROPSERVICE'].methods_by_name['crop_image']._serialized_options = b'\202\323\344\223\002\r\"\010/v1/crop:\001*'
  _globals['_CROPSERVICE'].methods_by_name['get_crop_suggestion']._loaded_options = None
  _globals['_CROPSERVICE'].methods_by_name['get_crop_suggestion']._serialized_options = b'\202\323\344\223\002\030\"\023/v1/crop/suggestion:\001*'
  _globals['_CROPIMAGEREQUEST']._serialized_start=171
  _globals['_CROPIMAGEREQUEST']._serialized_end=329
  _globals['_CROPIMAGERESPONSE']._serialized_start=331
  _globals['_CROPIMAGERESPONSE']._serialized_end=427
  _globals['_GETCROPSUGGESTIONREQUEST']._serialized_start=430
  _globals['_GETCROPSUGGESTIONREQUEST']._serialized_end=685
  _globals['_GETCROPSUGGESTIONRESPONSE']._serialized_start=687
  _globals['_GETCROPSUGGESTIONRESPONSE']._serialized_end=793
  _globals['_CROPSERVICE']._serialized_start=796
  _globals['_CROPSERVICE']._serialized_end=1058
# @@protoc_insertion_point(module_scope)
