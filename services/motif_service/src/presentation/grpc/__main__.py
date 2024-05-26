
import logging
from concurrent import futures

import grpc

import services.motif_service.src.infrastructure.__generated__.python.v1.b_box_service_pb2_grpc as b_box_pb2_grpc
from services.motif_service.src.application.services.b_box_service import \
    BBoxService
from services.motif_service.src.application.services.image_service import \
    ImageService


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    img_service = ImageService()
    b_box_service = BBoxService(img_service=img_service)
    b_box_pb2_grpc.add_BBoxServiceServicer_to_server(b_box_service, server)
    server.add_insecure_port("[::]:9090")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    logging.basicConfig()
    serve()