class Caption:
    """
    Image caption value object
    """
    
    def __init__(self, caption: str):
        self.__caption = caption
        
    def get_caption(self):
        return self.__caption