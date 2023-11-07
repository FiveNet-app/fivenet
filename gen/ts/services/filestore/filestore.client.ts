// @generated by protobuf-ts 2.9.1 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "services/filestore/filestore.proto" (package "services.filestore", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { DocStoreService } from "./filestore.js";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { UploadResponse } from "./filestore.js";
import type { UploadRequest } from "./filestore.js";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service services.filestore.DocStoreService
 */
export interface IDocStoreServiceClient {
    /**
     * @generated from protobuf rpc: Upload(services.filestore.UploadRequest) returns (services.filestore.UploadResponse);
     */
    upload(input: UploadRequest, options?: RpcOptions): UnaryCall<UploadRequest, UploadResponse>;
}
/**
 * @generated from protobuf service services.filestore.DocStoreService
 */
export class DocStoreServiceClient implements IDocStoreServiceClient, ServiceInfo {
    typeName = DocStoreService.typeName;
    methods = DocStoreService.methods;
    options = DocStoreService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: Upload(services.filestore.UploadRequest) returns (services.filestore.UploadResponse);
     */
    upload(input: UploadRequest, options?: RpcOptions): UnaryCall<UploadRequest, UploadResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<UploadRequest, UploadResponse>("unary", this._transport, method, opt, input);
    }
}
