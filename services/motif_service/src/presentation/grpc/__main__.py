
import logging
from concurrent import futures

import grpc
import src.infrastructure.__generated__.python.b_box_pb2_grpc as b_box_pb2_grpc
from src.application.services.b_box_service import BBoxService
from src.application.services.image_service import ImageService


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    b_box_pb2_grpc.add_BBoxServiceServicer_to_server(BBoxService(img_service=ImageService()), server)
    server.add_insecure_port("[::]:9090")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    logging.basicConfig()
    serve()