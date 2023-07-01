// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "services/centrum/centrum.proto" (package "services.centrum", syntax proto3)
// tslint:disable
import { CentrumService } from "./centrum.js";
import type { CentrumStreamResponse } from "./centrum.js";
import type { CentrumStreamRequest } from "./centrum.js";
import type { CreateDispatchResponse } from "./centrum.js";
import type { CreateDispatchRequest } from "./centrum.js";
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { UnitService } from "./centrum.js";
import type { UnitStreamResponse } from "./centrum.js";
import type { UnitStreamRequest } from "./centrum.js";
import type { ServerStreamingCall } from "@protobuf-ts/runtime-rpc";
import type { UpdateUnitStatusResponse } from "./centrum.js";
import type { UpdateUnitStatusRequest } from "./centrum.js";
import type { AssignUnitResponse } from "./centrum.js";
import type { AssignUnitRequest } from "./centrum.js";
import type { DeleteUnitResponse } from "./centrum.js";
import type { DeleteUnitRequest } from "./centrum.js";
import type { UpdateUnitResponse } from "./centrum.js";
import type { UpdateUnitRequest } from "./centrum.js";
import type { CreateUnitResponse } from "./centrum.js";
import type { CreateUnitRequest } from "./centrum.js";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { ListUnitsResponse } from "./centrum.js";
import type { ListUnitsRequest } from "./centrum.js";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service services.centrum.UnitService
 */
export interface IUnitServiceClient {
    /**
     * @perm
     *
     * @generated from protobuf rpc: ListUnits(services.centrum.ListUnitsRequest) returns (services.centrum.ListUnitsResponse);
     */
    listUnits(input: ListUnitsRequest, options?: RpcOptions): UnaryCall<ListUnitsRequest, ListUnitsResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: CreateUnit(services.centrum.CreateUnitRequest) returns (services.centrum.CreateUnitResponse);
     */
    createUnit(input: CreateUnitRequest, options?: RpcOptions): UnaryCall<CreateUnitRequest, CreateUnitResponse>;
    /**
     * @perm: Name=CreateUnit
     *
     * @generated from protobuf rpc: UpdateUnit(services.centrum.UpdateUnitRequest) returns (services.centrum.UpdateUnitResponse);
     */
    updateUnit(input: UpdateUnitRequest, options?: RpcOptions): UnaryCall<UpdateUnitRequest, UpdateUnitResponse>;
    /**
     * @perm: Name=DeleteUnit
     *
     * @generated from protobuf rpc: DeleteUnit(services.centrum.DeleteUnitRequest) returns (services.centrum.DeleteUnitResponse);
     */
    deleteUnit(input: DeleteUnitRequest, options?: RpcOptions): UnaryCall<DeleteUnitRequest, DeleteUnitResponse>;
    /**
     * @perm: Attrs=Access/StringList:[]string{"Own", "Lower_Rank", "Same_Rank"}§[]string{"Own"}
     *
     * @generated from protobuf rpc: AssignUnit(services.centrum.AssignUnitRequest) returns (services.centrum.AssignUnitResponse);
     */
    assignUnit(input: AssignUnitRequest, options?: RpcOptions): UnaryCall<AssignUnitRequest, AssignUnitResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: UpdateUnitStatus(services.centrum.UpdateUnitStatusRequest) returns (services.centrum.UpdateUnitStatusResponse);
     */
    updateUnitStatus(input: UpdateUnitStatusRequest, options?: RpcOptions): UnaryCall<UpdateUnitStatusRequest, UpdateUnitStatusResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: StreamUnits(services.centrum.UnitStreamRequest) returns (stream services.centrum.UnitStreamResponse);
     */
    streamUnits(input: UnitStreamRequest, options?: RpcOptions): ServerStreamingCall<UnitStreamRequest, UnitStreamResponse>;
}
/**
 * @generated from protobuf service services.centrum.UnitService
 */
export class UnitServiceClient implements IUnitServiceClient, ServiceInfo {
    typeName = UnitService.typeName;
    methods = UnitService.methods;
    options = UnitService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: ListUnits(services.centrum.ListUnitsRequest) returns (services.centrum.ListUnitsResponse);
     */
    listUnits(input: ListUnitsRequest, options?: RpcOptions): UnaryCall<ListUnitsRequest, ListUnitsResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListUnitsRequest, ListUnitsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: CreateUnit(services.centrum.CreateUnitRequest) returns (services.centrum.CreateUnitResponse);
     */
    createUnit(input: CreateUnitRequest, options?: RpcOptions): UnaryCall<CreateUnitRequest, CreateUnitResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateUnitRequest, CreateUnitResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=CreateUnit
     *
     * @generated from protobuf rpc: UpdateUnit(services.centrum.UpdateUnitRequest) returns (services.centrum.UpdateUnitResponse);
     */
    updateUnit(input: UpdateUnitRequest, options?: RpcOptions): UnaryCall<UpdateUnitRequest, UpdateUnitResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateUnitRequest, UpdateUnitResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=DeleteUnit
     *
     * @generated from protobuf rpc: DeleteUnit(services.centrum.DeleteUnitRequest) returns (services.centrum.DeleteUnitResponse);
     */
    deleteUnit(input: DeleteUnitRequest, options?: RpcOptions): UnaryCall<DeleteUnitRequest, DeleteUnitResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteUnitRequest, DeleteUnitResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Attrs=Access/StringList:[]string{"Own", "Lower_Rank", "Same_Rank"}§[]string{"Own"}
     *
     * @generated from protobuf rpc: AssignUnit(services.centrum.AssignUnitRequest) returns (services.centrum.AssignUnitResponse);
     */
    assignUnit(input: AssignUnitRequest, options?: RpcOptions): UnaryCall<AssignUnitRequest, AssignUnitResponse> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<AssignUnitRequest, AssignUnitResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: UpdateUnitStatus(services.centrum.UpdateUnitStatusRequest) returns (services.centrum.UpdateUnitStatusResponse);
     */
    updateUnitStatus(input: UpdateUnitStatusRequest, options?: RpcOptions): UnaryCall<UpdateUnitStatusRequest, UpdateUnitStatusResponse> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateUnitStatusRequest, UpdateUnitStatusResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: StreamUnits(services.centrum.UnitStreamRequest) returns (stream services.centrum.UnitStreamResponse);
     */
    streamUnits(input: UnitStreamRequest, options?: RpcOptions): ServerStreamingCall<UnitStreamRequest, UnitStreamResponse> {
        const method = this.methods[6], opt = this._transport.mergeOptions(options);
        return stackIntercept<UnitStreamRequest, UnitStreamResponse>("serverStreaming", this._transport, method, opt, input);
    }
}
/**
 * @generated from protobuf service services.centrum.CentrumService
 */
export interface ICentrumServiceClient {
    /**
     * @perm
     *
     * TODO
     *
     * @generated from protobuf rpc: CreateDispatch(services.centrum.CreateDispatchRequest) returns (services.centrum.CreateDispatchResponse);
     */
    createDispatch(input: CreateDispatchRequest, options?: RpcOptions): UnaryCall<CreateDispatchRequest, CreateDispatchResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: Stream(services.centrum.CentrumStreamRequest) returns (stream services.centrum.CentrumStreamResponse);
     */
    stream(input: CentrumStreamRequest, options?: RpcOptions): ServerStreamingCall<CentrumStreamRequest, CentrumStreamResponse>;
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
     * @perm
     *
     * TODO
     *
     * @generated from protobuf rpc: CreateDispatch(services.centrum.CreateDispatchRequest) returns (services.centrum.CreateDispatchResponse);
     */
    createDispatch(input: CreateDispatchRequest, options?: RpcOptions): UnaryCall<CreateDispatchRequest, CreateDispatchResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateDispatchRequest, CreateDispatchResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: Stream(services.centrum.CentrumStreamRequest) returns (stream services.centrum.CentrumStreamResponse);
     */
    stream(input: CentrumStreamRequest, options?: RpcOptions): ServerStreamingCall<CentrumStreamRequest, CentrumStreamResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<CentrumStreamRequest, CentrumStreamResponse>("serverStreaming", this._transport, method, opt, input);
    }
}
