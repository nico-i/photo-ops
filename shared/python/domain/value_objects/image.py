import base64

import cv2
import numpy as np
from PIL import Image as PILImage

from shared.python.__generated__.proto.messages.image_dto.v1.image_dto_pb2 import \
    ImageDto


class Image:
    """
    Image value object
    """
     
    def __init__(self, cv2_image):
        self.__cv2_image = cv2_image
        
    def get_cv2_img(self) -> np.ndarray:
        """
        Get the cv2 image (numpy array)
        """
        return self.__cv2_image
    
    def get_pil_img(self) -> PILImage:
        """
        Get the PIL image
        """
        cv_image_rgb = cv2.cvtColor(self.__cv2_image, cv2.COLOR_BGR2RGB)
        # Convert the OpenCV image (NumPy array) to a PIL Image
        pil_image = PILImage.fromarray(cv_image_rgb)
        return pil_image
            
    def to_base64_string(self) -> str:
        """
        Convert the image to base64 string

        Returns:
            base64_string: Base64 string of the image
        """
        # Encode the image to bytes
        _, img_data = cv2.imencode(".jpg", self.__cv2_image)
        
        # Convert bytes data to base64 string
        base64_string = base64.b64encode(img_data).decode()
        
        return base64_string
    
    @staticmethod
    def from_dto(dto: ImageDto) -> 'Image':
        """
        Create an Image object from a DTO

        Args:
            dto: Image DTO

        Returns:
            Image: Image object
        """
        if dto.HasField('path'):
            return Image.from_local_path(dto.path)
        else:
            raise ValueError("Image DTO must have a \"path\" field")
    
    @staticmethod
    def from_local_path( path: str) -> 'Image':
        img = cv2.imread(path)
        return Image(img)
    