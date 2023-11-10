import argparse
import io
import json
import cv2
import numpy as np
from PIL import Image, ImageDraw
from rembg import remove

def main():
    parser = argparse.ArgumentParser(description='Remove image background and/or draw bounding box on the original image.')
    parser.add_argument('--input-path','-i', type=str, help='Path to the input image')
    parser.add_argument('--debug--path', '-d', type=str, help='Path to save the debug image without a background and a bounding box drawn around it', default=None)

    args = parser.parse_args()

   

if __name__ == "__main__":
    main()