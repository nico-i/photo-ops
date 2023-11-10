
from PIL import Image, ImageDraw
from domain.aggregates.processed_image import ProcessedImage

def get_bbox(input_path, is_debug = False):
    processed_img = ProcessedImage(input_path)
    
    debug_path = None
    
    if is_debug:
        # replace the initial file extension with _debug.png
        debug_path = input_path[:input_path.rfind('.')] + '_debug.png'        
        
        bg_less_img = Image.fromarray(processed_img.bg_less_np_img)
        draw = ImageDraw.Draw(bg_less_img)
        
        bbox_top_left = (processed_img.bbox.x, processed_img.bbox.y)
        bbox_bottom_right = (processed_img.bbox.x + processed_img.bbox.width, processed_img.bbox.y + processed_img.bbox.height)
        
        draw.rectangle([bbox_top_left, bbox_bottom_right], outline="red", width=3)
        bg_less_img.save(debug_path)

    return str(processed_img.bbox)