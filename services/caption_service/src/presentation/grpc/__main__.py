
import logging
from concurrent import futures

import grpc

import shared.python.__generated__.proto.services.caption_service.v1.caption_service_pb2_grpc as caption_service_pb2_grpc
from services.caption_service.src.application.services.caption_service import \
    CaptionService


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    caption_service = CaptionService()
    caption_service_pb2_grpc.add_CaptionServiceServicer_to_server(caption_service, server)
    server.add_insecure_port("[::]:9091")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    serve()