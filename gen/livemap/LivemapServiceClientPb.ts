/**
 * @fileoverview gRPC-Web generated client stub for gen.livemap
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.21.12
// source: livemap/livemap.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as livemap_livemap_pb from '../livemap/livemap_pb';


export class LivemapServiceClient {
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

  methodDescriptorStream = new grpcWeb.MethodDescriptor(
    '/gen.livemap.LivemapService/Stream',
    grpcWeb.MethodType.SERVER_STREAMING,
    livemap_livemap_pb.StreamRequest,
    livemap_livemap_pb.LivemapMarker,
    (request: livemap_livemap_pb.StreamRequest) => {
      return request.serializeBinary();
    },
    livemap_livemap_pb.LivemapMarker.deserializeBinary
  );

  stream(
    request: livemap_livemap_pb.StreamRequest,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<livemap_livemap_pb.LivemapMarker> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/gen.livemap.LivemapService/Stream',
      request,
      metadata || {},
      this.methodDescriptorStream);
  }

}

