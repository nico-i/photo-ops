import numpy as np
from rembg import new_session, remove

from shared.python.domain.value_objects.image import Image


class ImageService:
    """
    Service to handle operations related to images
    """
    
    def __init__(self, model_name: str = "u2net"):
        self.__session = new_session(model_name)
        
    def remove_bg(self, image: Image) -> np.ndarray:
        """
        Remove background from the image
        """
        img_arr = image.get_cv2_img()
        no_bg_img_arr = remove(img_arr, session=self.__session, only_mask=True)
        
        return no_bg_img_arr