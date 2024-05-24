import numpy as np
from grpc import RpcContext
from PIL import Image
from rembg import new_session, remove

from src.domain.value_objects.b_box import BBox
from src.infrastructure.__generated__.python.b_box_pb2 import GetBBoxRequest
from src.infrastructure.__generated__.python.b_box_pb2_grpc import \
    BBoxServiceServicer


class BBoxService(BBoxServiceServicer):
    """
    Service to handle operations related to bounding boxes
    """
    
    def __init__(self):
        model_name = "isnet-general-use"
        self.__session = new_session(model_name)
        
    
    def __find_bounding_box(self, alpha_channel: np.array):
        """ Find the bounding box of a given image

        Args:
            alpha_channel (np.array): The alpha channel of the image

        Returns:
            _type_: The bounding box of the image
        """
        # Find all non-zero pixels
        rows = np.any(alpha_channel, axis=1)
        cols = np.any(alpha_channel, axis=0)

        # Find bounding box coordinates
        ymin, ymax = np.nonzero(rows)[0][[0, -1]]
        xmin, xmax = np.nonzero(cols)[0][[0, -1]]
        
        return BBox(xmin, ymin, xmax - xmin, ymax - ymin)
        
    
    def get_b_box(self, request:GetBBoxRequest, context: RpcContext):
        req_img = request.image
        if req_img.HasField("path"):
            img = Image.open(req_img.path)
            no_bg_img = remove(img, session=self.__session, only_mask=True)
            no_bg_img_arr = np.array(no_bg_img)
            
            # Threshold the image to remove pixels that are not fully opaque
            threshold = 150
            no_bg_img_arr[no_bg_img_arr < threshold] = 0  # Set pixels below the threshold to 0 (black)
            
            print(no_bg_img_arr.shape)
            
            bbox = self.__find_bounding_box(no_bg_img_arr)
        else:
            # request must contain base64 encoded image
            pass