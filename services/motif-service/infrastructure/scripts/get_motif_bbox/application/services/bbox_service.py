
import os
from PIL import Image, ImageDraw
from domain.aggregates.processed_image import ProcessedImage

def get_motif_bbox(input_path, debug_dir_path = None):
    processed_img = ProcessedImage(input_path)
        
    if debug_dir_path is not None:
        # check that last char is a slash
        if debug_dir_path[-1] != '/':
            raise ValueError("Debug directory path must end with a slash.")
        
        if not os.path.exists(debug_dir_path):
            print("Debug directory does not exist yet. Creating...")
            os.makedirs(debug_dir_path)    
        
        bg_less_img = Image.fromarray(processed_img.bg_less_np_img)
        draw = ImageDraw.Draw(bg_less_img)
        
        bbox_top_left = (processed_img.bbox.x, processed_img.bbox.y)
        bbox_bottom_right = (processed_img.bbox.x + processed_img.bbox.width, processed_img.bbox.y + processed_img.bbox.height)
        
        draw.rectangle([bbox_top_left, bbox_bottom_right], outline="red", width=3)
        
        filename = processed_img.file_name.split('.')[0]
        bg_less_img.save(os.path.join(debug_dir_path, f"{filename}_debug.png"))

    return str(processed_img.bbox)