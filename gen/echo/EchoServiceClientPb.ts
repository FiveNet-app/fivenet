/**
 * @fileoverview gRPC-Web generated client stub for grpc.gateway.testing
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.21.12
// source: echo/echo.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as echo_echo_pb from '../echo/echo_pb';


export class EchoServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorEcho = new grpcWeb.MethodDescriptor(
    '/grpc.gateway.testing.EchoService/Echo',
    grpcWeb.MethodType.UNARY,
    echo_echo_pb.EchoRequest,
    echo_echo_pb.EchoResponse,
    (request: echo_echo_pb.EchoRequest) => {
      return request.serializeBinary();
    },
    echo_echo_pb.EchoResponse.deserializeBinary
  );

  echo(
    request: echo_echo_pb.EchoRequest,
    metadata: grpcWeb.Metadata | null): Promise<echo_echo_pb.EchoResponse>;

  echo(
    request: echo_echo_pb.EchoRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: echo_echo_pb.EchoResponse) => void): grpcWeb.ClientReadableStream<echo_echo_pb.EchoResponse>;

  echo(
    request: echo_echo_pb.EchoRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: echo_echo_pb.EchoResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.gateway.testing.EchoService/Echo',
        request,
        metadata || {},
        this.methodDescriptorEcho,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.gateway.testing.EchoService/Echo',
    request,
    metadata || {},
    this.methodDescriptorEcho);
  }

  methodDescriptorServerStreamingEcho = new grpcWeb.MethodDescriptor(
    '/grpc.gateway.testing.EchoService/ServerStreamingEcho',
    grpcWeb.MethodType.SERVER_STREAMING,
    echo_echo_pb.ServerStreamingEchoRequest,
    echo_echo_pb.ServerStreamingEchoResponse,
    (request: echo_echo_pb.ServerStreamingEchoRequest) => {
      return request.serializeBinary();
    },
    echo_echo_pb.ServerStreamingEchoResponse.deserializeBinary
  );

  serverStreamingEcho(
    request: echo_echo_pb.ServerStreamingEchoRequest,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<echo_echo_pb.ServerStreamingEchoResponse> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/grpc.gateway.testing.EchoService/ServerStreamingEcho',
      request,
      metadata || {},
      this.methodDescriptorServerStreamingEcho);
  }

  methodDescriptorEchoStatus = new grpcWeb.MethodDescriptor(
    '/grpc.gateway.testing.EchoService/EchoStatus',
    grpcWeb.MethodType.UNARY,
    echo_echo_pb.EchoStatusRequest,
    echo_echo_pb.EchoStatusResponse,
    (request: echo_echo_pb.EchoStatusRequest) => {
      return request.serializeBinary();
    },
    echo_echo_pb.EchoStatusResponse.deserializeBinary
  );

  echoStatus(
    request: echo_echo_pb.EchoStatusRequest,
    metadata: grpcWeb.Metadata | null): Promise<echo_echo_pb.EchoStatusResponse>;

  echoStatus(
    request: echo_echo_pb.EchoStatusRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: echo_echo_pb.EchoStatusResponse) => void): grpcWeb.ClientReadableStream<echo_echo_pb.EchoStatusResponse>;

  echoStatus(
    request: echo_echo_pb.EchoStatusRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: echo_echo_pb.EchoStatusResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.gateway.testing.EchoService/EchoStatus',
        request,
        metadata || {},
        this.methodDescriptorEchoStatus,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.gateway.testing.EchoService/EchoStatus',
    request,
    metadata || {},
    this.methodDescriptorEchoStatus);
  }

}

