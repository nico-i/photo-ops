
import logging

import shared.python.__generated__.proto.services.hashtag_service.v1.hashtag_service_pb2_grpc as hashtag_service_pb2_grpc
from services.hashtag_service.src.application.services.hashtag_service import \
    HashtagService
from shared.python.__generated__.proto.services.hashtag_service.v1 import \
    hashtag_service_pb2
from shared.python.infrastructure.grpc.server import ServiceGrpcServer

if __name__ == "__main__":
    service = HashtagService()
    service_name = hashtag_service_pb2.DESCRIPTOR.services_by_name["HashtagService"].full_name
    add_func = hashtag_service_pb2_grpc.add_HashtagServiceServicer_to_server
    server = ServiceGrpcServer(service=service, service_name=service_name, port=9092, add_func=add_func, enable_reflection=True, log_level=logging.INFO)
    server.serve()