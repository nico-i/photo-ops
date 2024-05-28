from google.api import annotations_pb2 as _annotations_pb2
from messages.image_dto.v1 import image_dto_pb2 as _image_dto_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class GetHashtagsRequest(_message.Message):
    __slots__ = ("caption",)
    CAPTION_FIELD_NUMBER: _ClassVar[int]
    caption: str
    def __init__(self, caption: _Optional[str] = ...) -> None: ...

class GetHashtagsResponse(_message.Message):
    __slots__ = ("hashtags_csv",)
    HASHTAGS_CSV_FIELD_NUMBER: _ClassVar[int]
    hashtags_csv: str
    def __init__(self, hashtags_csv: _Optional[str] = ...) -> None: ...
