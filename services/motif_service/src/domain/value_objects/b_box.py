from shared.python.__generated__.proto.messages.rect_dto.v1.rect_dto_pb2 import \
    RectDto


class BBox:
    """
    Bounding box value object
    """
    def __init__(self, x: int, y: int, width: int, height: int):
        self.x = x
        self.y = y
        self.width = width
        self.height = height

    def __eq__(self, other: 'BBox') -> bool:
        return (
            self.x == other.x
            and self.y == other.y
            and self.width == other.width
            and self.height == other.height
        )

    def __str__(self) -> str:
        return f"x: {self.x}, y: {self.y}, width: {self.width}, height: {self.height}"
    
    def to_json(self):
        return {
            "x": self.x,
            "y": self.y,
            "width": self.width,
            "height": self.height
        }
        
    def to_dto(self) -> RectDto:
        """
        Convert the BBox object to a RectDto object

        Returns:
            BBoxDto: BBoxDto object
        """
        return RectDto(x=self.x, y=self.y, width=self.width, height=self.height)