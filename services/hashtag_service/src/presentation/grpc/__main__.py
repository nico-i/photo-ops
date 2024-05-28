
import logging
import os
from concurrent import futures

import grpc
from dotenv import load_dotenv

import shared.python.__generated__.proto.services.hashtag_service.v1.hashtag_service_pb2_grpc as hashtag_service_pb2_grpc
from services.hashtag_service.src.application.services.hashtag_service import \
    HashtagService

load_dotenv()

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    access_token = os.getenv("HUGGING_UAT")
    
    if(access_token is None):
        raise ValueError("\"HUGGING_UAT\" env var is not set")
    
    hashtag_service = HashtagService(access_token=access_token)
    hashtag_service_pb2_grpc.add_HashtagServiceServicer_to_server(hashtag_service, server)
    server.add_insecure_port("[::]:9090")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    serve()