# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from kuscia.proto.api.v1alpha1.datamesh import domaindatasource_pb2 as kuscia_dot_proto_dot_api_dot_v1alpha1_dot_datamesh_dot_domaindatasource__pb2


class DomainDataSourceServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.QueryDomainDataSource = channel.unary_unary(
                '/kuscia.proto.api.v1alpha1.datamesh.DomainDataSourceService/QueryDomainDataSource',
                request_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_datamesh_dot_domaindatasource__pb2.QueryDomainDataSourceRequest.SerializeToString,
                response_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_datamesh_dot_domaindatasource__pb2.QueryDomainDataSourceResponse.FromString,
                )


class DomainDataSourceServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def QueryDomainDataSource(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DomainDataSourceServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'QueryDomainDataSource': grpc.unary_unary_rpc_method_handler(
                    servicer.QueryDomainDataSource,
                    request_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_datamesh_dot_domaindatasource__pb2.QueryDomainDataSourceRequest.FromString,
                    response_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_datamesh_dot_domaindatasource__pb2.QueryDomainDataSourceResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'kuscia.proto.api.v1alpha1.datamesh.DomainDataSourceService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class DomainDataSourceService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def QueryDomainDataSource(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kuscia.proto.api.v1alpha1.datamesh.DomainDataSourceService/QueryDomainDataSource',
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_datamesh_dot_domaindatasource__pb2.QueryDomainDataSourceRequest.SerializeToString,
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_datamesh_dot_domaindatasource__pb2.QueryDomainDataSourceResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
