
import argparse

from domain.aggregates.processed_image import ProcessedImage
from application.services.bbox_service import get_bbox
def main():
    parser = argparse.ArgumentParser(description='Remove image background and/or draw bounding box on the original image.')
    parser.add_argument('--input-path','-i', type=str, help='Path to the input image', required=True)
    parser.add_argument('--is-debug','-d', action='store_true', help='Draw bounding box on the original image', default=False, required=False)

    args = parser.parse_args()
    bbox = get_bbox(args.input_path, args.is_debug)

    print(bbox)

if __name__ == "__main__":
    main()