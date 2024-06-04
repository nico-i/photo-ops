from google.api import annotations_pb2 as _annotations_pb2
from messages.image_dto.v1 import image_dto_pb2 as _image_dto_pb2
from messages.rect_dto.v1 import rect_dto_pb2 as _rect_dto_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CropImageRequest(_message.Message):
    __slots__ = ("input_image", "crop_spec", "output_path")
    INPUT_IMAGE_FIELD_NUMBER: _ClassVar[int]
    CROP_SPEC_FIELD_NUMBER: _ClassVar[int]
    OUTPUT_PATH_FIELD_NUMBER: _ClassVar[int]
    input_image: _image_dto_pb2.ImageDto
    crop_spec: _rect_dto_pb2.RectDto
    output_path: str
    def __init__(self, input_image: _Optional[_Union[_image_dto_pb2.ImageDto, _Mapping]] = ..., crop_spec: _Optional[_Union[_rect_dto_pb2.RectDto, _Mapping]] = ..., output_path: _Optional[str] = ...) -> None: ...

class CropImageResponse(_message.Message):
    __slots__ = ("cropped_image",)
    CROPPED_IMAGE_FIELD_NUMBER: _ClassVar[int]
    cropped_image: _image_dto_pb2.ImageDto
    def __init__(self, cropped_image: _Optional[_Union[_image_dto_pb2.ImageDto, _Mapping]] = ...) -> None: ...

class GetCropSuggestionRequest(_message.Message):
    __slots__ = ("image_motif", "image_width", "image_height", "target_aspect_ratio")
    IMAGE_MOTIF_FIELD_NUMBER: _ClassVar[int]
    IMAGE_WIDTH_FIELD_NUMBER: _ClassVar[int]
    IMAGE_HEIGHT_FIELD_NUMBER: _ClassVar[int]
    TARGET_ASPECT_RATIO_FIELD_NUMBER: _ClassVar[int]
    image_motif: _rect_dto_pb2.RectDto
    image_width: int
    image_height: int
    target_aspect_ratio: float
    def __init__(self, image_motif: _Optional[_Union[_rect_dto_pb2.RectDto, _Mapping]] = ..., image_width: _Optional[int] = ..., image_height: _Optional[int] = ..., target_aspect_ratio: _Optional[float] = ...) -> None: ...

class GetCropSuggestionResponse(_message.Message):
    __slots__ = ("crop_suggestion", "suggestion_out_of_bounds")
    CROP_SUGGESTION_FIELD_NUMBER: _ClassVar[int]
    SUGGESTION_OUT_OF_BOUNDS_FIELD_NUMBER: _ClassVar[int]
    crop_suggestion: _rect_dto_pb2.RectDto
    suggestion_out_of_bounds: bool
    def __init__(self, crop_suggestion: _Optional[_Union[_rect_dto_pb2.RectDto, _Mapping]] = ..., suggestion_out_of_bounds: bool = ...) -> None: ...
