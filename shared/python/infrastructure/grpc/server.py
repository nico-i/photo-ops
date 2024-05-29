
import logging
from concurrent import futures

import grpc
from grpc_reflection.v1alpha import reflection

from shared.python.infrastructure.logging import logger


class ServiceGrpcServer:
    
    def __init__(self, service: object, service_name: str,  port: int, add_func: object, max_workers: int = 10, enable_reflection: bool = False, log_level: int = logging.ERROR):
        logger.get_logger(log_level)
        
        self.__service = service
        self.__service_name = service_name
        self.__port = port
        self.__max_workers = max_workers
        self.__add_func = add_func
        self.__enable_reflection = enable_reflection
        
        logging.info("Starting server")

    def serve(self):
        port = self.__port
        
        server = grpc.server(futures.ThreadPoolExecutor(max_workers=self.__max_workers))
        
        self.__add_func(self.__service, server)
        
        if self.__enable_reflection:
            SERVICE_NAMES = (
                self.__service_name,
                reflection.SERVICE_NAME,
            )
            reflection.enable_server_reflection(SERVICE_NAMES, server)

        server.add_insecure_port(f"[::]:{port}")
        server.start()
        
        logging.info(f"Server started at port {port}")
        
        server.wait_for_termination()