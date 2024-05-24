import json

import numpy as np


class BoundingBox:
    def __init__(self, x, y, width, height):
        self.x = int(x)
        self.y = int(y)
        self.width = int(width)
        self.height =int(height)
    
    def __repr__(self):
        return json.dumps(self.__dict__)
    
def find_bounding_box(alpha_channel):
    # Find all non-zero pixels
    rows = np.any(alpha_channel, axis=1)
    cols = np.any(alpha_channel, axis=0)

    # Find bounding box coordinates
    ymin, ymax = np.nonzero(rows)[0][[0, -1]]
    xmin, xmax = np.nonzero(cols)[0][[0, -1]]

    return BoundingBox(xmin, ymin, xmax - xmin, ymax - ymin)