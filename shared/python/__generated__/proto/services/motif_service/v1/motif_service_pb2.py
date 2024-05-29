# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: services/motif_service/v1/motif_service.proto
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
    'services/motif_service/v1/motif_service.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from messages.base64_image_dto.v1 import base64_image_dto_pb2 as messages_dot_base64__image__dto_dot_v1_dot_base64__image__dto__pb2
from messages.image_dto.v1 import image_dto_pb2 as messages_dot_image__dto_dot_v1_dot_image__dto__pb2
from messages.rect_dto.v1 import rect_dto_pb2 as messages_dot_rect__dto_dot_v1_dot_rect__dto__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n-services/motif_service/v1/motif_service.proto\x12\x10motif_service.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x33messages/base64_image_dto/v1/base64_image_dto.proto\x1a%messages/image_dto/v1/image_dto.proto\x1a#messages/rect_dto/v1/rect_dto.proto\"N\n\x0eGetBBoxRequest\x12<\n\x05image\x18\x01 \x01(\x0b\x32&.shared.messages.image_dto.v1.ImageDtoR\x05image\"S\n\x13GetBBoxDebugRequest\x12<\n\x05image\x18\x01 \x01(\x0b\x32&.shared.messages.image_dto.v1.ImageDtoR\x05image\"L\n\x0fGetBBoxResponse\x12\x39\n\x05\x62_box\x18\x01 \x01(\x0b\x32$.shared.messages.rect_dto.v1.RectDtoR\x04\x62\x42ox\"a\n\x14GetBBoxDebugResponse\x12I\n\x05image\x18\x01 \x01(\x0b\x32\x33.shared.messages.base64_image_dto.v1.Base64ImageDtoR\x05image2\xf2\x01\n\x0cMotifService\x12\x65\n\tget_b_box\x12 .motif_service.v1.GetBBoxRequest\x1a!.motif_service.v1.GetBBoxResponse\"\x13\x82\xd3\xe4\x93\x02\r\"\x08/v1/bbox:\x01*\x12{\n\x0fget_b_box_debug\x12%.motif_service.v1.GetBBoxDebugRequest\x1a&.motif_service.v1.GetBBoxDebugResponse\"\x19\x82\xd3\xe4\x93\x02\x13\"\x0e/v1/bbox/debug:\x01*B\xeb\x01\n\x14\x63om.motif_service.v1B\x11MotifServiceProtoP\x01Zcgithub.com/nico-i/photo-ops/shared/go/__generated__/proto/services/motif_service/v1;motif_servicev1\xa2\x02\x03MXX\xaa\x02\x0fMotifService.V1\xca\x02\x0fMotifService\\V1\xe2\x02\x1bMotifService\\V1\\GPBMetadata\xea\x02\x10MotifService::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'services.motif_service.v1.motif_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\024com.motif_service.v1B\021MotifServiceProtoP\001Zcgithub.com/nico-i/photo-ops/shared/go/__generated__/proto/services/motif_service/v1;motif_servicev1\242\002\003MXX\252\002\017MotifService.V1\312\002\017MotifService\\V1\342\002\033MotifService\\V1\\GPBMetadata\352\002\020MotifService::V1'
  _globals['_MOTIFSERVICE'].methods_by_name['get_b_box']._loaded_options = None
  _globals['_MOTIFSERVICE'].methods_by_name['get_b_box']._serialized_options = b'\202\323\344\223\002\r\"\010/v1/bbox:\001*'
  _globals['_MOTIFSERVICE'].methods_by_name['get_b_box_debug']._loaded_options = None
  _globals['_MOTIFSERVICE'].methods_by_name['get_b_box_debug']._serialized_options = b'\202\323\344\223\002\023\"\016/v1/bbox/debug:\001*'
  _globals['_GETBBOXREQUEST']._serialized_start=226
  _globals['_GETBBOXREQUEST']._serialized_end=304
  _globals['_GETBBOXDEBUGREQUEST']._serialized_start=306
  _globals['_GETBBOXDEBUGREQUEST']._serialized_end=389
  _globals['_GETBBOXRESPONSE']._serialized_start=391
  _globals['_GETBBOXRESPONSE']._serialized_end=467
  _globals['_GETBBOXDEBUGRESPONSE']._serialized_start=469
  _globals['_GETBBOXDEBUGRESPONSE']._serialized_end=566
  _globals['_MOTIFSERVICE']._serialized_start=569
  _globals['_MOTIFSERVICE']._serialized_end=811
# @@protoc_insertion_point(module_scope)
