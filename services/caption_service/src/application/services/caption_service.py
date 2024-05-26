import cv2
from grpc import RpcContext
from transformers import AutoModelForCausalLM, AutoProcessor

from services.caption_service.src.domain.value_objects.caption import Caption
from services.caption_service.src.infrastructure.__generated__.python.caption_service.v1.caption_service_pb2 import (
    GetCaptionRequest, GetCaptionResponse)
from services.caption_service.src.infrastructure.__generated__.python.caption_service.v1.caption_service_pb2_grpc import \
    CaptionServiceServicer
from shared.python.domain.value_objects.image import Image


class CaptionService(CaptionServiceServicer):
    def __init__(self):
        self.__processor = AutoProcessor.from_pretrained("microsoft/git-base-coco")
        self.__model = AutoModelForCausalLM.from_pretrained("microsoft/git-base-coco")
    
    def get_caption(self, request: GetCaptionRequest, context: RpcContext) -> GetCaptionResponse:
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
        
        caption = Caption(predictions[0])
        
        return GetCaptionResponse(caption=caption.get_caption())