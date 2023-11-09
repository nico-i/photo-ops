import argparse
import io
import json

import cv2
import numpy as np
from PIL import Image, ImageDraw
from rembg import remove


def find_bounding_box(image):
    # Convert to numpy array
    np_image = np.array(image)

    # Find all non-transparent pixels
    rows = np.any(np_image[:, :, 3], axis=1)
    cols = np.any(np_image[:, :, 3], axis=0)

    # Find bounding box coordinates
    ymin, ymax = np.nonzero(rows)[0][[0, -1]]
    xmin, xmax = np.nonzero(cols)[0][[0, -1]]

    return xmin, ymin, xmax-xmin, ymax-ymin


def remove_noise(image_array, kernel_size=(25,25), iterations=5):
    """
    Remove noise from an image array using morphological operations.
    :param image_array: NumPy array of the image
    :param kernel_size: Size of the morphological kernel
    :param iterations: Number of times to perform the operation
    :return: NumPy array of the processed image
    """
    kernel = cv2.getStructuringElement(cv2.MORPH_RECT, kernel_size)
    erosion = cv2.erode(image_array, kernel, iterations=iterations)
    dilation = cv2.dilate(erosion, kernel, iterations=iterations)
    return dilation

def process_image(input_path, no_bg_output_path=None, bbox_output_path=None):
    # Open the image file
    with open(input_path, 'rb') as file:
        input_image = file.read()

    # Remove background
    output_image = remove(input_image)

    # Convert raw bytes to Image for processing
    processed_img = Image.open(io.BytesIO(output_image))
    
    # Convert to numpy array for morphological operations
    processed_np = np.array(processed_img)

    # Assume the alpha channel is the 4th channel
    alpha_channel = processed_np[:, :, 3]

    # Perform noise reduction on the alpha channel
    cleaned_alpha_channel = remove_noise(alpha_channel)

    # Replace the alpha channel in the processed image
    processed_np[:, :, 3] = cleaned_alpha_channel
    cleaned_img = Image.fromarray(processed_np)

    # Now, find the bounding box on the cleaned image
    bbox = find_bounding_box(cleaned_img)

    # Save the no background image
    if no_bg_output_path:
        cleaned_img.save(no_bg_output_path)

    # Draw bounding box on the original image and save
    if bbox_output_path:
        original_img = Image.open(input_path)
        draw = ImageDraw.Draw(original_img)
        draw.rectangle([(bbox[0], bbox[1]), (bbox[0] + bbox[2], bbox[1] + bbox[3])], outline="red")
        original_img.save(bbox_output_path)

    return bbox

def main():
    parser = argparse.ArgumentParser(description='Remove image background and/or draw bounding box on the original image.')
    parser.add_argument('--input-path','-i', type=str, help='Path to the input image')
    parser.add_argument('--no-bg-out', '-nbgo', type=str, help='Path to save the image without background', default=None)
    parser.add_argument('--bbox-out', '-bbxo', type=str, help='Path to save the image with bounding box', default=None)

    args = parser.parse_args()

    bbox = process_image(args.input_path, args.no_bg_out, args.bbox_out)

    # Convert numpy.int64 elements to Python native int for JSON serialization
    bbox_dict = {
        "x": int(bbox[0]), 
        "y": int(bbox[1]), 
        "width": int(bbox[2]), 
        "height": int(bbox[3])
    }
    
    print(json.dumps(bbox_dict))

if __name__ == "__main__":
    main()