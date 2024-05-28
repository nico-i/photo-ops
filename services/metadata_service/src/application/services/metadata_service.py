import logging

from grpc import RpcContext

from shared.python.__generated__.proto.services.metadata_service.v1.metadata_service_pb2 import \
    UpdateIPTCDataRequest
from shared.python.__generated__.proto.services.metadata_service.v1.metadata_service_pb2_grpc import \
    MetaDataServiceServicer
from shared.python.domain.value_objects.image import Image
from iptcinfo3 import IPTCInfo

class MetaDataService(MetaDataServiceServicer):
    
    def update_iptc_data(self, request: UpdateIPTCDataRequest, context: RpcContext) -> None:
        logging.info(f"request: {request}")
        
        img = Image.from_dto(request.image)
        caption = request.caption
        hashtags = request.hashtags
        
        