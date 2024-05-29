# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from services.motif_service.v1 import motif_service_pb2 as services_dot_motif__service_dot_v1_dot_motif__service__pb2


class MotifServiceStub(object):
    """Service responsible for handling all motif related operations.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.get_b_box = channel.unary_unary(
                '/motif_service.v1.MotifService/get_b_box',
                request_serializer=services_dot_motif__service_dot_v1_dot_motif__service__pb2.GetBBoxRequest.SerializeToString,
                response_deserializer=services_dot_motif__service_dot_v1_dot_motif__service__pb2.GetBBoxResponse.FromString,
                _registered_method=True)
        self.get_b_box_debug = channel.unary_unary(
                '/motif_service.v1.MotifService/get_b_box_debug',
                request_serializer=services_dot_motif__service_dot_v1_dot_motif__service__pb2.GetBBoxDebugRequest.SerializeToString,
                response_deserializer=services_dot_motif__service_dot_v1_dot_motif__service__pb2.GetBBoxDebugResponse.FromString,
                _registered_method=True)


class MotifServiceServicer(object):
    """Service responsible for handling all motif related operations.
    """

    def get_b_box(self, request, context):
        """Returns bounding box information for the subject in the image.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get_b_box_debug(self, request, context):
        """Returns the image used to generate the bounding box with the bounding box drawn on it.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MotifServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'get_b_box': grpc.unary_unary_rpc_method_handler(
                    servicer.get_b_box,
                    request_deserializer=services_dot_motif__service_dot_v1_dot_motif__service__pb2.GetBBoxRequest.FromString,
                    response_serializer=services_dot_motif__service_dot_v1_dot_motif__service__pb2.GetBBoxResponse.SerializeToString,
            ),
            'get_b_box_debug': grpc.unary_unary_rpc_method_handler(
                    servicer.get_b_box_debug,
                    request_deserializer=services_dot_motif__service_dot_v1_dot_motif__service__pb2.GetBBoxDebugRequest.FromString,
                    response_serializer=services_dot_motif__service_dot_v1_dot_motif__service__pb2.GetBBoxDebugResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'motif_service.v1.MotifService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('motif_service.v1.MotifService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class MotifService(object):
    """Service responsible for handling all motif related operations.
    """

    @staticmethod
    def get_b_box(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/motif_service.v1.MotifService/get_b_box',
            services_dot_motif__service_dot_v1_dot_motif__service__pb2.GetBBoxRequest.SerializeToString,
            services_dot_motif__service_dot_v1_dot_motif__service__pb2.GetBBoxResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def get_b_box_debug(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/motif_service.v1.MotifService/get_b_box_debug',
            services_dot_motif__service_dot_v1_dot_motif__service__pb2.GetBBoxDebugRequest.SerializeToString,
            services_dot_motif__service_dot_v1_dot_motif__service__pb2.GetBBoxDebugResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
