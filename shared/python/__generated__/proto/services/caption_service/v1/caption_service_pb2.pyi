from google.api import annotations_pb2 as _annotations_pb2
from messages.image_dto.v1 import image_dto_pb2 as _image_dto_pb2
from messages.rect_dto.v1 import rect_dto_pb2 as _rect_dto_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetCaptionRequest(_message.Message):
    __slots__ = ("image", "crop")
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    CROP_FIELD_NUMBER: _ClassVar[int]
    image: _image_dto_pb2.ImageDto
    crop: _rect_dto_pb2.RectDto
    def __init__(self, image: _Optional[_Union[_image_dto_pb2.ImageDto, _Mapping]] = ..., crop: _Optional[_Union[_rect_dto_pb2.RectDto, _Mapping]] = ...) -> None: ...

class GetCaptionResponse(_message.Message):
    __slots__ = ("caption",)
    CAPTION_FIELD_NUMBER: _ClassVar[int]
    caption: str
    def __init__(self, caption: _Optional[str] = ...) -> None: ...
