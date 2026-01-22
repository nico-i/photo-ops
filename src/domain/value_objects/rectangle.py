from shared.python.__generated__.proto.messages.rect_dto.v1.rect_dto_pb2 import \
    RectDto


class Rectangle:
    """
    Rectangle class to represent a rectangle with x, y, width, and height
    """
    def __init__(self, x: int, y: int, width: int, height: int):
        self.x = x
        self.y = y
        self.width = width
        self.height = height

    def __eq__(self, other: "Rectangle") -> bool:
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
        Convert the Rectangle object to a RectDto object

        Returns:
            RectDto: A RectDto object
        """
        return RectDto(x=self.x, y=self.y, width=self.width, height=self.height)
    
    @staticmethod
    def from_dto(dto: RectDto) -> "Rectangle":
        """
        Create a Rectangle object from a DTO

        Args:
            dto: Rectangle DTO

        Returns:
            Rectangle: Rectangle object
        """
        return Rectangle(x=dto.x, y=dto.y, width=dto.width, height=dto.height)
    
    def get_values(self) -> tuple:
        """
        Get the values of the rectangle

        Returns:
            tuple: x, y, width, height
        """
        return self.x, self.y, self.width, self.height