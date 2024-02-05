// @generated by protobuf-ts 2.9.3 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "services/jobs/conduct.proto" (package "services.jobs", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { JobsConductService } from "./conduct";
import type { DeleteConductEntryResponse } from "./conduct";
import type { DeleteConductEntryRequest } from "./conduct";
import type { UpdateConductEntryResponse } from "./conduct";
import type { UpdateConductEntryRequest } from "./conduct";
import type { CreateConductEntryResponse } from "./conduct";
import type { CreateConductEntryRequest } from "./conduct";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { ListConductEntriesResponse } from "./conduct";
import type { ListConductEntriesRequest } from "./conduct";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service services.jobs.JobsConductService
 */
export interface IJobsConductServiceClient {
    /**
     * @perm: Attrs=Access/StringList:[]string{"Own", "All"}
     *
     * @generated from protobuf rpc: ListConductEntries(services.jobs.ListConductEntriesRequest) returns (services.jobs.ListConductEntriesResponse);
     */
    listConductEntries(input: ListConductEntriesRequest, options?: RpcOptions): UnaryCall<ListConductEntriesRequest, ListConductEntriesResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: CreateConductEntry(services.jobs.CreateConductEntryRequest) returns (services.jobs.CreateConductEntryResponse);
     */
    createConductEntry(input: CreateConductEntryRequest, options?: RpcOptions): UnaryCall<CreateConductEntryRequest, CreateConductEntryResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: UpdateConductEntry(services.jobs.UpdateConductEntryRequest) returns (services.jobs.UpdateConductEntryResponse);
     */
    updateConductEntry(input: UpdateConductEntryRequest, options?: RpcOptions): UnaryCall<UpdateConductEntryRequest, UpdateConductEntryResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: DeleteConductEntry(services.jobs.DeleteConductEntryRequest) returns (services.jobs.DeleteConductEntryResponse);
     */
    deleteConductEntry(input: DeleteConductEntryRequest, options?: RpcOptions): UnaryCall<DeleteConductEntryRequest, DeleteConductEntryResponse>;
}
/**
 * @generated from protobuf service services.jobs.JobsConductService
 */
export class JobsConductServiceClient implements IJobsConductServiceClient, ServiceInfo {
    typeName = JobsConductService.typeName;
    methods = JobsConductService.methods;
    options = JobsConductService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @perm: Attrs=Access/StringList:[]string{"Own", "All"}
     *
     * @generated from protobuf rpc: ListConductEntries(services.jobs.ListConductEntriesRequest) returns (services.jobs.ListConductEntriesResponse);
     */
    listConductEntries(input: ListConductEntriesRequest, options?: RpcOptions): UnaryCall<ListConductEntriesRequest, ListConductEntriesResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListConductEntriesRequest, ListConductEntriesResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: CreateConductEntry(services.jobs.CreateConductEntryRequest) returns (services.jobs.CreateConductEntryResponse);
     */
    createConductEntry(input: CreateConductEntryRequest, options?: RpcOptions): UnaryCall<CreateConductEntryRequest, CreateConductEntryResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateConductEntryRequest, CreateConductEntryResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: UpdateConductEntry(services.jobs.UpdateConductEntryRequest) returns (services.jobs.UpdateConductEntryResponse);
     */
    updateConductEntry(input: UpdateConductEntryRequest, options?: RpcOptions): UnaryCall<UpdateConductEntryRequest, UpdateConductEntryResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateConductEntryRequest, UpdateConductEntryResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: DeleteConductEntry(services.jobs.DeleteConductEntryRequest) returns (services.jobs.DeleteConductEntryResponse);
     */
    deleteConductEntry(input: DeleteConductEntryRequest, options?: RpcOptions): UnaryCall<DeleteConductEntryRequest, DeleteConductEntryResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteConductEntryRequest, DeleteConductEntryResponse>("unary", this._transport, method, opt, input);
    }
}
