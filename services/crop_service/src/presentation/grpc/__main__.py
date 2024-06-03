
import logging

import shared.python.__generated__.proto.services.crop_service.v1.crop_service_pb2 as caption_service_pb2
import shared.python.__generated__.proto.services.crop_service.v1.crop_service_pb2_grpc as crop_service_pb2_grpc
from services.crop_service.src.application.services.caption_service import \
    CropService
from shared.python.infrastructure.grpc.server import ServiceGrpcServer

if __name__ == "__main__":
    service = CropService()
    service_name = caption_service_pb2.DESCRIPTOR.services_by_name["CropService"].full_name
    add_func = crop_service_pb2_grpc.add_CropServiceServicer_to_server
    server = ServiceGrpcServer(service=service, service_name=service_name, port=9093, add_func=add_func, enable_reflection=True, log_level=logging.INFO)
    server.serve()
