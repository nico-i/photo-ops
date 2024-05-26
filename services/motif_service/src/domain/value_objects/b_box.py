from services.motif_service.src.infrastructure.__generated__.python.motif_service.v1.b_box_service_pb2 import \
    BBoxDto


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
        
    def to_dto(self) -> BBoxDto:
        """
        Convert the BBox object to BBoxDto object

        Returns:
            BBoxDto: BBoxDto object
        """
        return BBoxDto(x=self.x, y=self.y, width=self.width, height=self.height)