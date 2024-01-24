# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from kuscia.proto.api.v1alpha1.confmanager import configuration_pb2 as kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_configuration__pb2


class ConfigurationServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateConfiguration = channel.unary_unary(
                '/kuscia.proto.api.v1alpha1.confmanager.ConfigurationService/CreateConfiguration',
                request_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_configuration__pb2.CreateConfigurationRequest.SerializeToString,
                response_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_configuration__pb2.CreateConfigurationResponse.FromString,
                )
        self.QueryConfiguration = channel.unary_unary(
                '/kuscia.proto.api.v1alpha1.confmanager.ConfigurationService/QueryConfiguration',
                request_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_configuration__pb2.QueryConfigurationRequest.SerializeToString,
                response_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_configuration__pb2.QueryConfigurationResponse.FromString,
                )


class ConfigurationServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateConfiguration(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def QueryConfiguration(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ConfigurationServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateConfiguration': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateConfiguration,
                    request_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_configuration__pb2.CreateConfigurationRequest.FromString,
                    response_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_configuration__pb2.CreateConfigurationResponse.SerializeToString,
            ),
            'QueryConfiguration': grpc.unary_unary_rpc_method_handler(
                    servicer.QueryConfiguration,
                    request_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_configuration__pb2.QueryConfigurationRequest.FromString,
                    response_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_configuration__pb2.QueryConfigurationResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'kuscia.proto.api.v1alpha1.confmanager.ConfigurationService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ConfigurationService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateConfiguration(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kuscia.proto.api.v1alpha1.confmanager.ConfigurationService/CreateConfiguration',
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_configuration__pb2.CreateConfigurationRequest.SerializeToString,
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_configuration__pb2.CreateConfigurationResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def QueryConfiguration(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kuscia.proto.api.v1alpha1.confmanager.ConfigurationService/QueryConfiguration',
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_configuration__pb2.QueryConfigurationRequest.SerializeToString,
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_configuration__pb2.QueryConfigurationResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
