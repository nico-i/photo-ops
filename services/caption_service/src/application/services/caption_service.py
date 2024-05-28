import logging

from grpc import RpcContext
from transformers import AutoModelForCausalLM, AutoProcessor

from shared.python.__generated__.proto.services.caption_service.v1.caption_service_pb2 import (
    GetCaptionRequest, GetCaptionResponse)
from shared.python.__generated__.proto.services.caption_service.v1.caption_service_pb2_grpc import \
    CaptionServiceServicer
from shared.python.domain.value_objects.image import Image


class CaptionService(CaptionServiceServicer):
    def __init__(self):
        self.__processor = AutoProcessor.from_pretrained("microsoft/git-base-coco")
        self.__model = AutoModelForCausalLM.from_pretrained("microsoft/git-base-coco")
    
    def get_caption(self, request: GetCaptionRequest, context: RpcContext) -> GetCaptionResponse:
        logging.info(f"request: {request}")
        try:
            img = Image.from_local_path(request.image.path) if request.image.HasField('path') else Image.from_base64_string(request.image.base64_image.data)
        except Exception as e:
            context.set_code(3)
            context.set_details(str(e))
            return
        
        pil_img = img.get_pil_img()

        pixel_values = self.__processor(images=pil_img, return_tensors="pt").pixel_values

        generated_ids = self.__model.generate(pixel_values=pixel_values, max_length=50)

        predictions = self.__processor.batch_decode(generated_ids, skip_special_tokens=True)
    
        res = GetCaptionResponse(caption=predictions[0])
        
        logging.info(f"response: {res}")
        
        return res