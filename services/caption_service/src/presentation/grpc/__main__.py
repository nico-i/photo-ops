
import logging

import shared.python.__generated__.proto.services.caption_service.v1.caption_service_pb2_grpc as caption_service_pb2_grpc
from services.caption_service.src.application.services.caption_service import \
    CaptionService
from shared.python.__generated__.proto.services.caption_service.v1 import \
    caption_service_pb2
from shared.python.infrastructure.grpc.server import ServiceGrpcServer

if __name__ == "__main__":
    service = CaptionService()
    service_name = caption_service_pb2.DESCRIPTOR.services_by_name["CaptionService"].full_name
    add_func = caption_service_pb2_grpc.add_CaptionServiceServicer_to_server
    server = ServiceGrpcServer(service=service, service_name=service_name, port=9091, add_func=add_func, enable_reflection=True, log_level=logging.INFO)
    server.serve()
