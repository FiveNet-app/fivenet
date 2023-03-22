/**
 * @fileoverview gRPC-Web generated client stub for services.citizenstore
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.21.12
// source: services/citizenstore/citizenstore.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as services_citizenstore_citizenstore_pb from '../../services/citizenstore/citizenstore_pb';


export class CitizenStoreServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorFindUsers = new grpcWeb.MethodDescriptor(
    '/services.citizenstore.CitizenStoreService/FindUsers',
    grpcWeb.MethodType.UNARY,
    services_citizenstore_citizenstore_pb.FindUsersRequest,
    services_citizenstore_citizenstore_pb.FindUsersResponse,
    (request: services_citizenstore_citizenstore_pb.FindUsersRequest) => {
      return request.serializeBinary();
    },
    services_citizenstore_citizenstore_pb.FindUsersResponse.deserializeBinary
  );

  findUsers(
    request: services_citizenstore_citizenstore_pb.FindUsersRequest,
    metadata: grpcWeb.Metadata | null): Promise<services_citizenstore_citizenstore_pb.FindUsersResponse>;

  findUsers(
    request: services_citizenstore_citizenstore_pb.FindUsersRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: services_citizenstore_citizenstore_pb.FindUsersResponse) => void): grpcWeb.ClientReadableStream<services_citizenstore_citizenstore_pb.FindUsersResponse>;

  findUsers(
    request: services_citizenstore_citizenstore_pb.FindUsersRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: services_citizenstore_citizenstore_pb.FindUsersResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/services.citizenstore.CitizenStoreService/FindUsers',
        request,
        metadata || {},
        this.methodDescriptorFindUsers,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/services.citizenstore.CitizenStoreService/FindUsers',
    request,
    metadata || {},
    this.methodDescriptorFindUsers);
  }

  methodDescriptorGetUser = new grpcWeb.MethodDescriptor(
    '/services.citizenstore.CitizenStoreService/GetUser',
    grpcWeb.MethodType.UNARY,
    services_citizenstore_citizenstore_pb.GetUserRequest,
    services_citizenstore_citizenstore_pb.GetUserResponse,
    (request: services_citizenstore_citizenstore_pb.GetUserRequest) => {
      return request.serializeBinary();
    },
    services_citizenstore_citizenstore_pb.GetUserResponse.deserializeBinary
  );

  getUser(
    request: services_citizenstore_citizenstore_pb.GetUserRequest,
    metadata: grpcWeb.Metadata | null): Promise<services_citizenstore_citizenstore_pb.GetUserResponse>;

  getUser(
    request: services_citizenstore_citizenstore_pb.GetUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: services_citizenstore_citizenstore_pb.GetUserResponse) => void): grpcWeb.ClientReadableStream<services_citizenstore_citizenstore_pb.GetUserResponse>;

  getUser(
    request: services_citizenstore_citizenstore_pb.GetUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: services_citizenstore_citizenstore_pb.GetUserResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/services.citizenstore.CitizenStoreService/GetUser',
        request,
        metadata || {},
        this.methodDescriptorGetUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/services.citizenstore.CitizenStoreService/GetUser',
    request,
    metadata || {},
    this.methodDescriptorGetUser);
  }

  methodDescriptorGetUserActivity = new grpcWeb.MethodDescriptor(
    '/services.citizenstore.CitizenStoreService/GetUserActivity',
    grpcWeb.MethodType.UNARY,
    services_citizenstore_citizenstore_pb.GetUserActivityRequest,
    services_citizenstore_citizenstore_pb.GetUserActivityResponse,
    (request: services_citizenstore_citizenstore_pb.GetUserActivityRequest) => {
      return request.serializeBinary();
    },
    services_citizenstore_citizenstore_pb.GetUserActivityResponse.deserializeBinary
  );

  getUserActivity(
    request: services_citizenstore_citizenstore_pb.GetUserActivityRequest,
    metadata: grpcWeb.Metadata | null): Promise<services_citizenstore_citizenstore_pb.GetUserActivityResponse>;

  getUserActivity(
    request: services_citizenstore_citizenstore_pb.GetUserActivityRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: services_citizenstore_citizenstore_pb.GetUserActivityResponse) => void): grpcWeb.ClientReadableStream<services_citizenstore_citizenstore_pb.GetUserActivityResponse>;

  getUserActivity(
    request: services_citizenstore_citizenstore_pb.GetUserActivityRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: services_citizenstore_citizenstore_pb.GetUserActivityResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/services.citizenstore.CitizenStoreService/GetUserActivity',
        request,
        metadata || {},
        this.methodDescriptorGetUserActivity,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/services.citizenstore.CitizenStoreService/GetUserActivity',
    request,
    metadata || {},
    this.methodDescriptorGetUserActivity);
  }

  methodDescriptorSetUserProps = new grpcWeb.MethodDescriptor(
    '/services.citizenstore.CitizenStoreService/SetUserProps',
    grpcWeb.MethodType.UNARY,
    services_citizenstore_citizenstore_pb.SetUserPropsRequest,
    services_citizenstore_citizenstore_pb.SetUserPropsResponse,
    (request: services_citizenstore_citizenstore_pb.SetUserPropsRequest) => {
      return request.serializeBinary();
    },
    services_citizenstore_citizenstore_pb.SetUserPropsResponse.deserializeBinary
  );

  setUserProps(
    request: services_citizenstore_citizenstore_pb.SetUserPropsRequest,
    metadata: grpcWeb.Metadata | null): Promise<services_citizenstore_citizenstore_pb.SetUserPropsResponse>;

  setUserProps(
    request: services_citizenstore_citizenstore_pb.SetUserPropsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: services_citizenstore_citizenstore_pb.SetUserPropsResponse) => void): grpcWeb.ClientReadableStream<services_citizenstore_citizenstore_pb.SetUserPropsResponse>;

  setUserProps(
    request: services_citizenstore_citizenstore_pb.SetUserPropsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: services_citizenstore_citizenstore_pb.SetUserPropsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/services.citizenstore.CitizenStoreService/SetUserProps',
        request,
        metadata || {},
        this.methodDescriptorSetUserProps,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/services.citizenstore.CitizenStoreService/SetUserProps',
    request,
    metadata || {},
    this.methodDescriptorSetUserProps);
  }

}

