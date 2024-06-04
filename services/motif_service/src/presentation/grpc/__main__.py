import argparse
import logging

import shared.python.__generated__.proto.services.motif_service.v1.motif_service_pb2_grpc as motif_service_pb2_grpc
from services.motif_service.src.application.services.image_service import \
    ImageService
from services.motif_service.src.application.services.motif_service import \
    MotifService
from shared.python.__generated__.proto.services.motif_service.v1 import \
    motif_service_pb2
from shared.python.infrastructure.grpc.server import ServiceGrpcServer

arg_parser  = argparse.ArgumentParser(
                    prog='Motif Service',
                    description='Detects motifs in images and returns the results in the form of bounding boxes')
arg_parser.add_argument('--debug', action='store_true', help='Enables the debug endpoint')

if __name__ == "__main__":
	args = arg_parser.parse_args()
	enable_debug = args.debug
    
	img_service = ImageService()
	service = MotifService(img_service=img_service, enable_debug=enable_debug)
 
	service_name = motif_service_pb2.DESCRIPTOR.services_by_name["MotifService"].full_name
	add_func = motif_service_pb2_grpc.add_MotifServiceServicer_to_server
 
	server = ServiceGrpcServer(service=service, service_name=service_name, port=9090, add_func=add_func, enable_reflection=True, log_level=logging.ERROR if not enable_debug else logging.DEBUG)
	server.serve()
