import cv2
import numpy as np
from grpc import RpcContext
from src.application.services.image_service import ImageService
from src.domain.value_objects.b_box import BBox
from src.domain.value_objects.image import Image
from src.infrastructure.__generated__.python.messages.v1.base64_image_dto_pb2 import \
    Base64ImageDto
from src.infrastructure.__generated__.python.v1.b_box_service_pb2 import (
    GetBBoxDebugRequest, GetBBoxDebugResponse, GetBBoxRequest, GetBBoxResponse)
from src.infrastructure.__generated__.python.v1.b_box_service_pb2_grpc import \
    BBoxServiceServicer


class BBoxService(BBoxServiceServicer):
	"""
	Service to handle operations related to bounding boxes
	"""

	def __init__(self, img_service: ImageService, threshold: int = 240):
		self.__img_service = img_service
        # 240 has been tested on the test images and has provided good results
		self.__threshold = threshold


	def __threshold_ndarray(self, img_arr: np.array) -> np.array:
		"""
		Threshold a ndarray
		
		Args:
		    img_arr: Image array
		    
		Returns:
		    binary_arr: Thresholded image array
		"""
		_, binary_arr = cv2.threshold(img_arr, self.__threshold
			, 255, cv2.THRESH_BINARY)

		return binary_arr

	def __find_b_box(self, img_arr: np.array) -> BBox:
		"""
		Find the largest contour in a image mask contained in a ndarray
		
		Args:
		    img: Image object
		    
		Returns:
		    largest_contour: Largest contour in the image
		""" 
		binary_arr = self.__threshold_ndarray(img_arr)

		contours, _ = cv2.findContours(binary_arr, cv2.RETR_EXTERNAL,cv2.CHAIN_APPROX_NONE)

		largest_contour = max(contours, key=cv2.contourArea)

		x, y, w, h = cv2.boundingRect(largest_contour)

		b_box = BBox(x, y, w, h)

		return b_box

	def get_b_box(self, request:GetBBoxRequest, context: RpcContext) -> GetBBoxResponse:
		try:
			img: Image = Image.from_dto(request.image)
		except Exception as e:
			context.set_code(3)
			context.set_details(str(e))
			return

		no_bg_img_arr = self.__img_service.remove_bg(img)

		b_box = self.__find_b_box(no_bg_img_arr)

		b_box_dto = b_box.to_dto()

		res = GetBBoxResponse(b_box=b_box_dto)

		return res

	def get_b_box_debug(self, request:GetBBoxDebugRequest, context: RpcContext) -> GetBBoxDebugResponse:
		try:
			img: Image = Image.from_dto(request.image)
		except Exception as e:
			context.set_code(3)
			context.set_details(str(e))
			return

		no_bg_img_arr = self.__img_service.remove_bg(img)

		b_box = self.__find_b_box(no_bg_img_arr)

		binary_arr = self.__threshold_ndarray(no_bg_img_arr)

			# Convert the binary image to BGR format for overlay
		mask_img = cv2.cvtColor(binary_arr, cv2.COLOR_GRAY2BGR)

		# Overlay the original image onto the binary image with opacity
		original_img = img.get_arr()

		opacity = 0.4
		overlay_image = cv2.addWeighted(original_img, opacity, mask_img, 1 - opacity, 0)

		# Draw the bounding box on the overlay image
		cv2.rectangle(overlay_image, (b_box.x, b_box.y), (b_box.x + b_box.width, b_box.y + b_box.height), (0, 255, 0), 2)

		composed_img = Image(overlay_image)

		base64_str = composed_img.to_base64_string()
		base64_img_dto = Base64ImageDto(data=base64_str)

		res = GetBBoxDebugResponse(image=base64_img_dto)
		return res
