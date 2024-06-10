from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class ImageDto(_message.Message):
    __slots__ = ("path", "base64")
    PATH_FIELD_NUMBER: _ClassVar[int]
    BASE64_FIELD_NUMBER: _ClassVar[int]
    path: str
    base64: bytes
    def __init__(self, path: _Optional[str] = ..., base64: _Optional[bytes] = ...) -> None: ...
