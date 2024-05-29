from google.api import annotations_pb2 as _annotations_pb2
from messages.base64_image_dto.v1 import base64_image_dto_pb2 as _base64_image_dto_pb2
from messages.image_dto.v1 import image_dto_pb2 as _image_dto_pb2
from messages.rect_dto.v1 import rect_dto_pb2 as _rect_dto_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetBBoxRequest(_message.Message):
    __slots__ = ("image",)
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    image: _image_dto_pb2.ImageDto
    def __init__(self, image: _Optional[_Union[_image_dto_pb2.ImageDto, _Mapping]] = ...) -> None: ...

class GetBBoxDebugRequest(_message.Message):
    __slots__ = ("image",)
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    image: _image_dto_pb2.ImageDto
    def __init__(self, image: _Optional[_Union[_image_dto_pb2.ImageDto, _Mapping]] = ...) -> None: ...

class GetBBoxResponse(_message.Message):
    __slots__ = ("b_box",)
    B_BOX_FIELD_NUMBER: _ClassVar[int]
    b_box: _rect_dto_pb2.RectDto
    def __init__(self, b_box: _Optional[_Union[_rect_dto_pb2.RectDto, _Mapping]] = ...) -> None: ...

class GetBBoxDebugResponse(_message.Message):
    __slots__ = ("image",)
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    image: _base64_image_dto_pb2.Base64ImageDto
    def __init__(self, image: _Optional[_Union[_base64_image_dto_pb2.Base64ImageDto, _Mapping]] = ...) -> None: ...
