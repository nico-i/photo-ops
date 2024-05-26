
import logging
from concurrent import futures

import grpc

from services.caption_service.src.application.services.caption_service import \
    CaptionService
from services.caption_service.src.infrastructure.__generated__.python.caption_service.v1 import \
    caption_service_pb2_grpc


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    caption_service = CaptionService()
    caption_service_pb2_grpc.add_CaptionServiceServicer_to_server(caption_service, server)
    server.add_insecure_port("[::]:9090")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    serve()