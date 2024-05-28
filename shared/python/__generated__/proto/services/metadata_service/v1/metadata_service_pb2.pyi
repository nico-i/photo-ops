from google.api import annotations_pb2 as _annotations_pb2
from messages.image_dto.v1 import image_dto_pb2 as _image_dto_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class UpdateIPTCDataRequest(_message.Message):
    __slots__ = ("image", "caption", "hashtag_csv")
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    CAPTION_FIELD_NUMBER: _ClassVar[int]
    HASHTAG_CSV_FIELD_NUMBER: _ClassVar[int]
    image: _image_dto_pb2.ImageDto
    caption: str
    hashtag_csv: str
    def __init__(self, image: _Optional[_Union[_image_dto_pb2.ImageDto, _Mapping]] = ..., caption: _Optional[str] = ..., hashtag_csv: _Optional[str] = ...) -> None: ...

class UpdateIPTCDataResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
