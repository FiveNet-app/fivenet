// @generated by protobuf-ts 2.9.1 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "services/centrum/centrum.proto" (package "services.centrum", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { CentrumService } from "./centrum.js";
import type { StreamResponse } from "./centrum.js";
import type { StreamRequest } from "./centrum.js";
import type { ServerStreamingCall } from "@protobuf-ts/runtime-rpc";
import type { TakeDispatchResponse } from "./centrum.js";
import type { TakeDispatchRequest } from "./centrum.js";
import type { AssignDispatchResponse } from "./centrum.js";
import type { AssignDispatchRequest } from "./centrum.js";
import type { ListDispatchActivityResponse } from "./centrum.js";
import type { UpdateDispatchStatusResponse } from "./centrum.js";
import type { UpdateDispatchStatusRequest } from "./centrum.js";
import type { UpdateDispatchResponse } from "./centrum.js";
import type { UpdateDispatchRequest } from "./centrum.js";
import type { CreateDispatchResponse } from "./centrum.js";
import type { CreateDispatchRequest } from "./centrum.js";
import type { ListDispatchesResponse } from "./centrum.js";
import type { ListDispatchesRequest } from "./centrum.js";
import type { ListUnitActivityResponse } from "./centrum.js";
import type { ListActivityRequest } from "./centrum.js";
import type { AssignUnitResponse } from "./centrum.js";
import type { AssignUnitRequest } from "./centrum.js";
import type { UpdateUnitStatusResponse } from "./centrum.js";
import type { UpdateUnitStatusRequest } from "./centrum.js";
import type { DeleteUnitResponse } from "./centrum.js";
import type { DeleteUnitRequest } from "./centrum.js";
import type { CreateOrUpdateUnitResponse } from "./centrum.js";
import type { CreateOrUpdateUnitRequest } from "./centrum.js";
import type { ListUnitsResponse } from "./centrum.js";
import type { ListUnitsRequest } from "./centrum.js";
import type { TakeControlResponse } from "./centrum.js";
import type { TakeControlRequest } from "./centrum.js";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { Settings } from "../../resources/dispatch/settings.js";
import type { GetSettingsRequest } from "./centrum.js";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service services.centrum.CentrumService
 */
export interface ICentrumServiceClient {
    /**
     * @perm: Name=Stream
     *
     * @generated from protobuf rpc: GetSettings(services.centrum.GetSettingsRequest) returns (resources.dispatch.Settings);
     */
    getSettings(input: GetSettingsRequest, options?: RpcOptions): UnaryCall<GetSettingsRequest, Settings>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: UpdateSettings(resources.dispatch.Settings) returns (resources.dispatch.Settings);
     */
    updateSettings(input: Settings, options?: RpcOptions): UnaryCall<Settings, Settings>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: TakeControl(services.centrum.TakeControlRequest) returns (services.centrum.TakeControlResponse);
     */
    takeControl(input: TakeControlRequest, options?: RpcOptions): UnaryCall<TakeControlRequest, TakeControlResponse>;
    /**
     * @perm: Name=Stream
     *
     * @generated from protobuf rpc: ListUnits(services.centrum.ListUnitsRequest) returns (services.centrum.ListUnitsResponse);
     */
    listUnits(input: ListUnitsRequest, options?: RpcOptions): UnaryCall<ListUnitsRequest, ListUnitsResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: CreateOrUpdateUnit(services.centrum.CreateOrUpdateUnitRequest) returns (services.centrum.CreateOrUpdateUnitResponse);
     */
    createOrUpdateUnit(input: CreateOrUpdateUnitRequest, options?: RpcOptions): UnaryCall<CreateOrUpdateUnitRequest, CreateOrUpdateUnitResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: DeleteUnit(services.centrum.DeleteUnitRequest) returns (services.centrum.DeleteUnitResponse);
     */
    deleteUnit(input: DeleteUnitRequest, options?: RpcOptions): UnaryCall<DeleteUnitRequest, DeleteUnitResponse>;
    /**
     * @perm: Name=TakeDispatch
     *
     * @generated from protobuf rpc: UpdateUnitStatus(services.centrum.UpdateUnitStatusRequest) returns (services.centrum.UpdateUnitStatusResponse);
     */
    updateUnitStatus(input: UpdateUnitStatusRequest, options?: RpcOptions): UnaryCall<UpdateUnitStatusRequest, UpdateUnitStatusResponse>;
    /**
     * @perm: Name=TakeControl
     *
     * @generated from protobuf rpc: AssignUnit(services.centrum.AssignUnitRequest) returns (services.centrum.AssignUnitResponse);
     */
    assignUnit(input: AssignUnitRequest, options?: RpcOptions): UnaryCall<AssignUnitRequest, AssignUnitResponse>;
    /**
     * @perm: Name=Stream
     *
     * @generated from protobuf rpc: ListUnitActivity(services.centrum.ListActivityRequest) returns (services.centrum.ListUnitActivityResponse);
     */
    listUnitActivity(input: ListActivityRequest, options?: RpcOptions): UnaryCall<ListActivityRequest, ListUnitActivityResponse>;
    /**
     * @perm: Name=Stream
     *
     * @generated from protobuf rpc: ListDispatches(services.centrum.ListDispatchesRequest) returns (services.centrum.ListDispatchesResponse);
     */
    listDispatches(input: ListDispatchesRequest, options?: RpcOptions): UnaryCall<ListDispatchesRequest, ListDispatchesResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: CreateDispatch(services.centrum.CreateDispatchRequest) returns (services.centrum.CreateDispatchResponse);
     */
    createDispatch(input: CreateDispatchRequest, options?: RpcOptions): UnaryCall<CreateDispatchRequest, CreateDispatchResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: UpdateDispatch(services.centrum.UpdateDispatchRequest) returns (services.centrum.UpdateDispatchResponse);
     */
    updateDispatch(input: UpdateDispatchRequest, options?: RpcOptions): UnaryCall<UpdateDispatchRequest, UpdateDispatchResponse>;
    /**
     * @perm: Name=TakeDispatch
     *
     * @generated from protobuf rpc: UpdateDispatchStatus(services.centrum.UpdateDispatchStatusRequest) returns (services.centrum.UpdateDispatchStatusResponse);
     */
    updateDispatchStatus(input: UpdateDispatchStatusRequest, options?: RpcOptions): UnaryCall<UpdateDispatchStatusRequest, UpdateDispatchStatusResponse>;
    /**
     * @perm: Name=Stream
     *
     * @generated from protobuf rpc: ListDispatchActivity(services.centrum.ListActivityRequest) returns (services.centrum.ListDispatchActivityResponse);
     */
    listDispatchActivity(input: ListActivityRequest, options?: RpcOptions): UnaryCall<ListActivityRequest, ListDispatchActivityResponse>;
    /**
     * @perm: Name=TakeControl
     *
     * @generated from protobuf rpc: AssignDispatch(services.centrum.AssignDispatchRequest) returns (services.centrum.AssignDispatchResponse);
     */
    assignDispatch(input: AssignDispatchRequest, options?: RpcOptions): UnaryCall<AssignDispatchRequest, AssignDispatchResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: TakeDispatch(services.centrum.TakeDispatchRequest) returns (services.centrum.TakeDispatchResponse);
     */
    takeDispatch(input: TakeDispatchRequest, options?: RpcOptions): UnaryCall<TakeDispatchRequest, TakeDispatchResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: Stream(services.centrum.StreamRequest) returns (stream services.centrum.StreamResponse);
     */
    stream(input: StreamRequest, options?: RpcOptions): ServerStreamingCall<StreamRequest, StreamResponse>;
}
/**
 * @generated from protobuf service services.centrum.CentrumService
 */
export class CentrumServiceClient implements ICentrumServiceClient, ServiceInfo {
    typeName = CentrumService.typeName;
    methods = CentrumService.methods;
    options = CentrumService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @perm: Name=Stream
     *
     * @generated from protobuf rpc: GetSettings(services.centrum.GetSettingsRequest) returns (resources.dispatch.Settings);
     */
    getSettings(input: GetSettingsRequest, options?: RpcOptions): UnaryCall<GetSettingsRequest, Settings> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetSettingsRequest, Settings>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: UpdateSettings(resources.dispatch.Settings) returns (resources.dispatch.Settings);
     */
    updateSettings(input: Settings, options?: RpcOptions): UnaryCall<Settings, Settings> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<Settings, Settings>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: TakeControl(services.centrum.TakeControlRequest) returns (services.centrum.TakeControlResponse);
     */
    takeControl(input: TakeControlRequest, options?: RpcOptions): UnaryCall<TakeControlRequest, TakeControlResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<TakeControlRequest, TakeControlResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=Stream
     *
     * @generated from protobuf rpc: ListUnits(services.centrum.ListUnitsRequest) returns (services.centrum.ListUnitsResponse);
     */
    listUnits(input: ListUnitsRequest, options?: RpcOptions): UnaryCall<ListUnitsRequest, ListUnitsResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListUnitsRequest, ListUnitsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: CreateOrUpdateUnit(services.centrum.CreateOrUpdateUnitRequest) returns (services.centrum.CreateOrUpdateUnitResponse);
     */
    createOrUpdateUnit(input: CreateOrUpdateUnitRequest, options?: RpcOptions): UnaryCall<CreateOrUpdateUnitRequest, CreateOrUpdateUnitResponse> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateOrUpdateUnitRequest, CreateOrUpdateUnitResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: DeleteUnit(services.centrum.DeleteUnitRequest) returns (services.centrum.DeleteUnitResponse);
     */
    deleteUnit(input: DeleteUnitRequest, options?: RpcOptions): UnaryCall<DeleteUnitRequest, DeleteUnitResponse> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteUnitRequest, DeleteUnitResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=TakeDispatch
     *
     * @generated from protobuf rpc: UpdateUnitStatus(services.centrum.UpdateUnitStatusRequest) returns (services.centrum.UpdateUnitStatusResponse);
     */
    updateUnitStatus(input: UpdateUnitStatusRequest, options?: RpcOptions): UnaryCall<UpdateUnitStatusRequest, UpdateUnitStatusResponse> {
        const method = this.methods[6], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateUnitStatusRequest, UpdateUnitStatusResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=TakeControl
     *
     * @generated from protobuf rpc: AssignUnit(services.centrum.AssignUnitRequest) returns (services.centrum.AssignUnitResponse);
     */
    assignUnit(input: AssignUnitRequest, options?: RpcOptions): UnaryCall<AssignUnitRequest, AssignUnitResponse> {
        const method = this.methods[7], opt = this._transport.mergeOptions(options);
        return stackIntercept<AssignUnitRequest, AssignUnitResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=Stream
     *
     * @generated from protobuf rpc: ListUnitActivity(services.centrum.ListActivityRequest) returns (services.centrum.ListUnitActivityResponse);
     */
    listUnitActivity(input: ListActivityRequest, options?: RpcOptions): UnaryCall<ListActivityRequest, ListUnitActivityResponse> {
        const method = this.methods[8], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListActivityRequest, ListUnitActivityResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=Stream
     *
     * @generated from protobuf rpc: ListDispatches(services.centrum.ListDispatchesRequest) returns (services.centrum.ListDispatchesResponse);
     */
    listDispatches(input: ListDispatchesRequest, options?: RpcOptions): UnaryCall<ListDispatchesRequest, ListDispatchesResponse> {
        const method = this.methods[9], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListDispatchesRequest, ListDispatchesResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: CreateDispatch(services.centrum.CreateDispatchRequest) returns (services.centrum.CreateDispatchResponse);
     */
    createDispatch(input: CreateDispatchRequest, options?: RpcOptions): UnaryCall<CreateDispatchRequest, CreateDispatchResponse> {
        const method = this.methods[10], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateDispatchRequest, CreateDispatchResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: UpdateDispatch(services.centrum.UpdateDispatchRequest) returns (services.centrum.UpdateDispatchResponse);
     */
    updateDispatch(input: UpdateDispatchRequest, options?: RpcOptions): UnaryCall<UpdateDispatchRequest, UpdateDispatchResponse> {
        const method = this.methods[11], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateDispatchRequest, UpdateDispatchResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=TakeDispatch
     *
     * @generated from protobuf rpc: UpdateDispatchStatus(services.centrum.UpdateDispatchStatusRequest) returns (services.centrum.UpdateDispatchStatusResponse);
     */
    updateDispatchStatus(input: UpdateDispatchStatusRequest, options?: RpcOptions): UnaryCall<UpdateDispatchStatusRequest, UpdateDispatchStatusResponse> {
        const method = this.methods[12], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateDispatchStatusRequest, UpdateDispatchStatusResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=Stream
     *
     * @generated from protobuf rpc: ListDispatchActivity(services.centrum.ListActivityRequest) returns (services.centrum.ListDispatchActivityResponse);
     */
    listDispatchActivity(input: ListActivityRequest, options?: RpcOptions): UnaryCall<ListActivityRequest, ListDispatchActivityResponse> {
        const method = this.methods[13], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListActivityRequest, ListDispatchActivityResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=TakeControl
     *
     * @generated from protobuf rpc: AssignDispatch(services.centrum.AssignDispatchRequest) returns (services.centrum.AssignDispatchResponse);
     */
    assignDispatch(input: AssignDispatchRequest, options?: RpcOptions): UnaryCall<AssignDispatchRequest, AssignDispatchResponse> {
        const method = this.methods[14], opt = this._transport.mergeOptions(options);
        return stackIntercept<AssignDispatchRequest, AssignDispatchResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: TakeDispatch(services.centrum.TakeDispatchRequest) returns (services.centrum.TakeDispatchResponse);
     */
    takeDispatch(input: TakeDispatchRequest, options?: RpcOptions): UnaryCall<TakeDispatchRequest, TakeDispatchResponse> {
        const method = this.methods[15], opt = this._transport.mergeOptions(options);
        return stackIntercept<TakeDispatchRequest, TakeDispatchResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: Stream(services.centrum.StreamRequest) returns (stream services.centrum.StreamResponse);
     */
    stream(input: StreamRequest, options?: RpcOptions): ServerStreamingCall<StreamRequest, StreamResponse> {
        const method = this.methods[16], opt = this._transport.mergeOptions(options);
        return stackIntercept<StreamRequest, StreamResponse>("serverStreaming", this._transport, method, opt, input);
    }
}
