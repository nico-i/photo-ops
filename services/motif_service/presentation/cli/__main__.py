
import argparse
from bbox_service import get_motif_bbox

def main():
    parser = argparse.ArgumentParser(description='Remove image background and/or draw bounding box on the original image.')
    parser.add_argument('input_path', type=str, help='Path to the input image.', metavar='input_path')
    parser.add_argument('--debug','-d',type=str, help='Provides path to a debug directory where an image with a bounding box on top of the backgroundless image will be saved.', default=None, required=False, metavar='debug_dir_path')

    args = parser.parse_args()

    bbox = get_motif_bbox(args.input_path, args.debug)

    print(bbox)

if __name__ == "__main__":
    main()