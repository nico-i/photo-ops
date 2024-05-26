# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: services/caption_service/v1/caption_service.proto
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
    'services/caption_service/v1/caption_service.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from messages.image_dto.v1 import image_dto_pb2 as messages_dot_image__dto_dot_v1_dot_image__dto__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n1services/caption_service/v1/caption_service.proto\x12\x12\x63\x61ption_service.v1\x1a\x1cgoogle/api/annotations.proto\x1a%messages/image_dto/v1/image_dto.proto\"Q\n\x11GetCaptionRequest\x12<\n\x05image\x18\x01 \x01(\x0b\x32&.shared.messages.image_dto.v1.ImageDtoR\x05image\".\n\x12GetCaptionResponse\x12\x18\n\x07\x63\x61ption\x18\x01 \x01(\tR\x07\x63\x61ption2\x86\x01\n\x0e\x43\x61ptionService\x12t\n\x0bget_caption\x12%.caption_service.v1.GetCaptionRequest\x1a&.caption_service.v1.GetCaptionResponse\"\x16\x82\xd3\xe4\x93\x02\x10\"\x0b/v1/caption:\x01*B\xfb\x01\n\x16\x63om.caption_service.v1B\x13\x43\x61ptionServiceProtoP\x01Zggithub.com/nico-i/photo-ops/shared/go/__generated__/proto/services/caption_service/v1;caption_servicev1\xa2\x02\x03\x43XX\xaa\x02\x11\x43\x61ptionService.V1\xca\x02\x11\x43\x61ptionService\\V1\xe2\x02\x1d\x43\x61ptionService\\V1\\GPBMetadata\xea\x02\x12\x43\x61ptionService::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'services.caption_service.v1.caption_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\026com.caption_service.v1B\023CaptionServiceProtoP\001Zggithub.com/nico-i/photo-ops/shared/go/__generated__/proto/services/caption_service/v1;caption_servicev1\242\002\003CXX\252\002\021CaptionService.V1\312\002\021CaptionService\\V1\342\002\035CaptionService\\V1\\GPBMetadata\352\002\022CaptionService::V1'
  _globals['_CAPTIONSERVICE'].methods_by_name['get_caption']._loaded_options = None
  _globals['_CAPTIONSERVICE'].methods_by_name['get_caption']._serialized_options = b'\202\323\344\223\002\020\"\013/v1/caption:\001*'
  _globals['_GETCAPTIONREQUEST']._serialized_start=142
  _globals['_GETCAPTIONREQUEST']._serialized_end=223
  _globals['_GETCAPTIONRESPONSE']._serialized_start=225
  _globals['_GETCAPTIONRESPONSE']._serialized_end=271
  _globals['_CAPTIONSERVICE']._serialized_start=274
  _globals['_CAPTIONSERVICE']._serialized_end=408
# @@protoc_insertion_point(module_scope)
