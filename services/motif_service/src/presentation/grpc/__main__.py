
import logging
from concurrent import futures

import grpc
from grpc_reflection.v1alpha import reflection

import shared.python.__generated__.proto.services.motif_service.v1.motif_service_pb2_grpc as motif_service_pb2_grpc
from services.motif_service.src.application.services.b_box_service import \
    MotifService
from services.motif_service.src.application.services.image_service import \
    ImageService
from shared.python.__generated__.proto.services.motif_service.v1 import \
    motif_service_pb2


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    img_service = ImageService()
    motif_service = MotifService(img_service=img_service)
    motif_service_pb2_grpc.add_MotifServiceServicer_to_server(motif_service, server)
    
    SERVICE_NAMES = (
    motif_service_pb2.DESCRIPTOR.services_by_name["MotifService"].full_name,
    reflection.SERVICE_NAME,
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)
    
    server.add_insecure_port("[::]:9090")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    logging.basicConfig()
    serve()