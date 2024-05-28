
import logging
import os
from concurrent import futures

import grpc
from grpc_reflection.v1alpha import reflection

import shared.python.__generated__.proto.services.hashtag_service.v1.hashtag_service_pb2_grpc as hashtag_service_pb2_grpc
from services.hashtag_service.src.application.services.hashtag_service import \
    HashtagService
from shared.python.__generated__.proto.services.hashtag_service.v1 import \
    hashtag_service_pb2


def serve():
    port = 9090
    
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    logging.info("Initializing service")
    hashtag_service = HashtagService()
    logging.info("Service intialized")
    hashtag_service_pb2_grpc.add_HashtagServiceServicer_to_server(hashtag_service, server)
    
    SERVICE_NAMES = (
        hashtag_service_pb2.DESCRIPTOR.services_by_name["HashtagService"].full_name,
        reflection.SERVICE_NAME,
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)
    
    logging.info(f"Server started at port {port}")
    server.add_insecure_port(f"[::]:{port}")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    serve()