import cv2
import numpy as np
from grpc import RpcContext
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
        model_name = "u2net"
        self.__session = new_session(model_name)
    
    def get_b_box(self, request:GetBBoxRequest, context: RpcContext):
        print(request)
        req_img = request.image
        if req_img.HasField("path"):
            original_img = cv2.imread(req_img.path)
            no_bg_img = remove(original_img, session=self.__session, only_mask=True)
            
            # Threshold the image to get a binary image
            _, binary = cv2.threshold(no_bg_img, 240, 255, cv2.THRESH_BINARY)

            # Find contours in the binary image
            contours, _ = cv2.findContours(binary, cv2.RETR_EXTERNAL,cv2.CHAIN_APPROX_NONE)

            # Find the contour with the largest area
            largest_contour = max(contours, key=cv2.contourArea)

            # Get the bounding box coordinates for the largest contour
            x, y, w, h = cv2.boundingRect(largest_contour)

            # Convert the binary image to BGR format for overlay
            mask_img = cv2.cvtColor(binary, cv2.COLOR_GRAY2BGR)

            # Overlay the original image onto the binary image with opacity
            opacity = 0.4
            overlay_image = cv2.addWeighted(original_img, opacity, mask_img, 1 - opacity, 0)

            # Draw the bounding box on the overlay image
            cv2.rectangle(overlay_image, (x, y), (x + w, y + h), (0, 255, 0), 2)

            # Save or display the result
            output_path = "debug.png"
            cv2.imwrite(output_path, overlay_image)
        else:
            # request must contain base64 encoded image
            pass