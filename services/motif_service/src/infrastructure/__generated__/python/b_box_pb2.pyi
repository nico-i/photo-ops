from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ImageDto(_message.Message):
    __slots__ = ("path", "base64_image")
    PATH_FIELD_NUMBER: _ClassVar[int]
    BASE64_IMAGE_FIELD_NUMBER: _ClassVar[int]
    path: str
    base64_image: Base64ImageDto
    def __init__(self, path: _Optional[str] = ..., base64_image: _Optional[_Union[Base64ImageDto, _Mapping]] = ...) -> None: ...

class Base64ImageDto(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: str
    def __init__(self, data: _Optional[str] = ...) -> None: ...

class BBoxDto(_message.Message):
    __slots__ = ("x", "y", "width", "height")
    X_FIELD_NUMBER: _ClassVar[int]
    Y_FIELD_NUMBER: _ClassVar[int]
    WIDTH_FIELD_NUMBER: _ClassVar[int]
    HEIGHT_FIELD_NUMBER: _ClassVar[int]
    x: int
    y: int
    width: int
    height: int
    def __init__(self, x: _Optional[int] = ..., y: _Optional[int] = ..., width: _Optional[int] = ..., height: _Optional[int] = ...) -> None: ...

class GetBBoxRequest(_message.Message):
    __slots__ = ("image",)
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    image: ImageDto
    def __init__(self, image: _Optional[_Union[ImageDto, _Mapping]] = ...) -> None: ...

class GetBBoxDebugRequest(_message.Message):
    __slots__ = ("image",)
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    image: ImageDto
    def __init__(self, image: _Optional[_Union[ImageDto, _Mapping]] = ...) -> None: ...

class GetBBoxResponse(_message.Message):
    __slots__ = ("b_box",)
    B_BOX_FIELD_NUMBER: _ClassVar[int]
    b_box: BBoxDto
    def __init__(self, b_box: _Optional[_Union[BBoxDto, _Mapping]] = ...) -> None: ...

class GetBBoxDebugResponse(_message.Message):
    __slots__ = ("image",)
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    image: Base64ImageDto
    def __init__(self, image: _Optional[_Union[Base64ImageDto, _Mapping]] = ...) -> None: ...
