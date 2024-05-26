from messages.base64_image_dto.v1 import base64_image_dto_pb2 as _base64_image_dto_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ImageDto(_message.Message):
    __slots__ = ("path", "base64_image")
    PATH_FIELD_NUMBER: _ClassVar[int]
    BASE64_IMAGE_FIELD_NUMBER: _ClassVar[int]
    path: str
    base64_image: _base64_image_dto_pb2.Base64ImageDto
    def __init__(self, path: _Optional[str] = ..., base64_image: _Optional[_Union[_base64_image_dto_pb2.Base64ImageDto, _Mapping]] = ...) -> None: ...
