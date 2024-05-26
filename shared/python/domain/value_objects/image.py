import base64

import cv2
import numpy as np
from PIL import Image as PILImage


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
    