import logging

from grpc import RpcContext

from shared.python.__generated__.proto.services.crop_service.v1.crop_service_pb2 import (
    CropImageRequest, CropImageResponse, GetCropSuggestionRequest)
from shared.python.__generated__.proto.services.crop_service.v1.crop_service_pb2_grpc import \
    CropServiceServicer
from shared.python.domain.value_objects.image import Image


class CropService(CropServiceServicer):
    
    def crop_image(self, request: CropImageRequest, context: RpcContext) -> CropImageResponse:
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get_crop_suggestion(self, request: GetCropSuggestionRequest, context: RpcContext) -> CropImageResponse:
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')