
import io
import os

import cv2
import numpy as np
from bounding_box import find_bounding_box
from PIL import Image, ImageDraw
from rembg import remove


def get_motif_bbox(input_path, debug_dir_path = None):
    with open(input_path, 'rb') as file:
        input_image = file.read()
        
    # Remove background
    debug_img = remove(input_image)

    processed_img = Image.open(io.BytesIO(debug_img))
    processed_np = np.array(processed_img)
    
    # Threshold the image to remove pixels that are not fully opaque
    threshold = 150
    processed_np[processed_np < threshold] = 0  # Set pixels below the threshold to 0 (black)

    # Assume the alpha channel is the 4th channel
    alpha_channel = processed_np[:, :, 3]

    # denoise alpha channel
    cleaned_alpha_channel = remove_noise(alpha_channel, (10,10), 1)

    # Replace the alpha channel in the noisy image
    processed_np[:, :, 3] = cleaned_alpha_channel
    cleaned_img = Image.fromarray(processed_np)

    bbox = find_bounding_box(cleaned_img)
    base_name = os.path.basename(input_path)
    debug_img_arr = processed_np
        
    if debug_dir_path is not None:        
        if not os.path.exists(debug_dir_path):
            print("Debug directory does not exist yet. Creating...")
            os.makedirs(debug_dir_path)
            print(f"Debug directory created at {debug_dir_path}")
        
        debug_img = Image.fromarray(debug_img_arr)
        draw = ImageDraw.Draw(debug_img)
        
        bbox_top_left = (bbox.x, bbox.y)
        bbox_bottom_right = (bbox.x + bbox.width, bbox.y + bbox.height)
        
        draw.rectangle([bbox_top_left, bbox_bottom_right], outline="red", width=3)
        
        filename = base_name.split('.')[0]
        output_path = os.path.join(debug_dir_path, f"{filename}_debug.png")
        
        print(f"Saving debug image to {output_path}")
        debug_img.save(output_path)

    return str(bbox)

def remove_noise(self,image_array, kernel_size=(5, 5), iterations=1):
    """
    Apply morphological operations to remove isolated noise pixels from the alpha channel.
    """
    kernel = cv2.getStructuringElement(cv2.MORPH_RECT, kernel_size)
    # Use 'closing' morphological operation (dilation followed by erosion) to remove noise.
    opening = cv2.morphologyEx(image_array, cv2.MORPH_OPEN, kernel, iterations)
    return opening