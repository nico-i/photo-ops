import logging
import os
from operator import is_

import cv2
from grpc import RpcContext

from shared.python.__generated__.proto.services.crop_service.v1.crop_service_pb2 import (
    CropImageRequest, CropImageResponse, GetCropSuggestionRequest,
    GetCropSuggestionResponse)
from shared.python.__generated__.proto.services.crop_service.v1.crop_service_pb2_grpc import \
    CropServiceServicer
from shared.python.domain.value_objects.image import Image
from shared.python.domain.value_objects.rectangle import Rectangle


class CropService(CropServiceServicer):
    
    def crop_image(self, request: CropImageRequest, context: RpcContext) -> CropImageResponse:
        input_image = Image.from_dto(request.input_image)
        logging.info(f"Received image with shape {input_image.get_cv2_img().shape}")
        
        crop_spec = Rectangle.from_dto(request.crop_spec)
        logging.info(f"Received crop spec {crop_spec.get_values()}")
        
        input_image_cv2 = input_image.get_cv2_img()
        x, y, w, h = crop_spec.get_values()
        
        cropped_image_cv2 = input_image_cv2[y:y+h, x:x+w]
        
        cropped_image = Image(cropped_image_cv2)
        
        output_path = request.output_path
        
        if len(output_path) > 0:
            if not os.path.exists(os.path.dirname(output_path)):
                logging.warning(f"Output directory {os.path.dirname(output_path)} does not exist. Creating it.")
                os.makedirs(os.path.dirname(output_path))
            cv2.imwrite(output_path, cropped_image_cv2)
            logging.info(f"Image cropped and saved to {output_path}")
            return CropImageResponse(cropped_image=cropped_image.to_dto(output_path))
        
        return CropImageResponse(cropped_image=cropped_image.to_dto())
        
        

    def get_crop_suggestion(self, request: GetCropSuggestionRequest, context: RpcContext) -> CropImageResponse:
        motif = Rectangle.from_dto(request.image_motif)
        image_width, image_height = request.image_width, request.image_height
        target_ratio = request.target_aspect_ratio
        
        motif_x, motif_y, motif_width, motif_height = motif.get_values()
        
        crop_width = max(motif_width, target_ratio * motif_height) # either the width of the motif or the target aspect ratio times the height of the motif
        crop_height = max(motif_height, motif_width / target_ratio) # either the height of the motif or the width of the motif divided by the target aspect ratio

        crop_width = max(crop_width, crop_height * target_ratio) # make sure W is at least as large as H * target_ratio
        
        # center the motif in the crop rectangle by scaling the crop rectangle to the image bounds
        
        crop_x = motif_x + motif_width / 2 - crop_width / 2
        crop_y = motif_y + motif_height / 2 - crop_height / 2
        
        crop_center_x = crop_x + crop_width / 2
        crop_center_y = crop_y + crop_height / 2
        
        scale_width = min((image_width - crop_center_x) / (crop_width / 2), crop_center_x / (crop_width / 2))
        scale_height = min((image_height - crop_center_y) / (crop_height / 2), crop_center_y / (crop_height / 2))
        
        scale_factor = min(scale_width, scale_height)
        
        new_crop_width = crop_width * scale_factor
        new_crop_height = crop_height * scale_factor
        
        new_crop_x = crop_center_x - new_crop_width / 2
        new_crop_y = crop_center_y - new_crop_height / 2
        
        suggestion = Rectangle(int(new_crop_x), int(new_crop_y), int(new_crop_width), int(new_crop_height))
        
        crop_x, crop_y, crop_w, crop_h = suggestion.get_values()
        
        suggestion_dto = suggestion.to_dto()
        
        # if the crop rectangle is smaller than the motif, return a suggestion with the out_of_bounds flag set to true
        if crop_w < motif_width or crop_h < motif_height:
            return GetCropSuggestionResponse(crop_suggestion=suggestion_dto, invalid_bounds=True)
            
        return GetCropSuggestionResponse(crop_suggestion=suggestion_dto)