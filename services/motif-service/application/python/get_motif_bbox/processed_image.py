import io
import cv2
from PIL import Image
from rembg import remove
import numpy as np
from bounding_box import find_bounding_box
import os

class ProcessedImage:
    def __init__(self, input_path):
        # Open the image file
        with open(input_path, 'rb') as file:
            input_image = file.read()
            
        self.base_name = os.path.basename(input_path)

        # Remove background
        bgless_image = remove(input_image)

        # Convert raw bytes to Image for processing
        processed_img = Image.open(io.BytesIO(bgless_image))
        
        # Convert to numpy array for morphological operations
        processed_np = np.array(processed_img)

        # Assume the alpha channel is the 4th channel
        alpha_channel = processed_np[:, :, 3]

        # Perform noise reduction on the alpha channel
        cleaned_alpha_channel = self.__remove_noise(alpha_channel)

        # Replace the alpha channel in the processed image
        processed_np[:, :, 3] = cleaned_alpha_channel
        cleaned_img = Image.fromarray(processed_np)

        # Now, find the bounding box on the cleaned image
        self.bbox = find_bounding_box(cleaned_img)
        self.input_image = input_image
        self.bg_less_np_img = processed_np
        
    def __remove_noise(self,image_array, kernel_size=(5, 5), iterations=1):
        """
        Apply morphological operations to remove isolated noise pixels from the alpha channel.
        """
        kernel = cv2.getStructuringElement(cv2.MORPH_RECT, kernel_size)
        # Use 'closing' morphological operation (dilation followed by erosion) to remove noise.
        closing = cv2.morphologyEx(image_array, cv2.MORPH_CLOSE, kernel, iterations=iterations)
        return closing
    
    