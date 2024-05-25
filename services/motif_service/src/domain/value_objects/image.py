import base64

import cv2
import numpy as np
from src.infrastructure.__generated__.python.b_box_pb2 import ImageDto


class Image:
    """
    Image value object
    """
     
    def __init__(self, cv2_image):
        self.__cv2_image = cv2_image
        
    def get_arr(self) -> np.ndarray:
        """
        Get the ndarray of this image
        """
        return self.__cv2_image
    
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
        Create an Image object from ImageDto object

        Args:
            dto (ImageDto): ImageDto object

        Raises:
            ValueError: If ImageDto object does not have path or base64_image

        Returns:
            Image: Image object
        """
        if dto.HasField('path'):
            return Image.from_local_path(dto.path)
        elif dto.HasField('base64_image'):
            return Image.from_base64_string(dto.base64_image.data)
        else:
            raise ValueError("ImageDto must have either path or base64_image")
    
    @staticmethod
    def from_local_path( path: str) -> 'Image':
        img = cv2.imread(path)
        return Image(img)
    
    @staticmethod
    def from_base64_string(base64_string: str) -> 'Image':
         # Decode base64 string to bytes
        img_data = base64.b64decode(base64_string)
        
        # Convert bytes data to numpy array
        np_array = np.frombuffer(img_data, dtype=np.uint8)
        
        # Decode numpy array to an image
        img = cv2.imdecode(np_array, cv2.IMREAD_COLOR)
        
        return Image(img)
    