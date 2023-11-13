
import os
from PIL import Image, ImageDraw
from processed_image import ProcessedImage

def get_motif_bbox(input_path, debug_dir_path = None):
    processed_img = ProcessedImage(input_path)
        
    if debug_dir_path is not None:        
        if not os.path.exists(debug_dir_path):
            print("Debug directory does not exist yet. Creating...")
            os.makedirs(debug_dir_path)
            print(f"Debug directory created at {debug_dir_path}")
        
        debug_img = Image.fromarray(processed_img.debug_img_arr)
        draw = ImageDraw.Draw(debug_img)
        
        bbox_top_left = (processed_img.bbox.x, processed_img.bbox.y)
        bbox_bottom_right = (processed_img.bbox.x + processed_img.bbox.width, processed_img.bbox.y + processed_img.bbox.height)
        
        draw.rectangle([bbox_top_left, bbox_bottom_right], outline="red", width=3)
        
        filename = processed_img.base_name.split('.')[0]
        output_path = os.path.join(debug_dir_path, f"{filename}_debug.png")
        
        print(f"Saving debug image to {output_path}")
        debug_img.save(output_path)

    return str(processed_img.bbox)