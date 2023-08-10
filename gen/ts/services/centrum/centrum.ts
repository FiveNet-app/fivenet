// @generated by protobuf-ts 2.9.1 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "services/centrum/centrum.proto" (package "services.centrum", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { MessageType } from "@protobuf-ts/runtime";
import { Settings } from "../../resources/dispatch/settings.js";
import { DispatchStatus } from "../../resources/dispatch/dispatches.js";
import { Dispatch } from "../../resources/dispatch/dispatches.js";
import { DISPATCH_STATUS } from "../../resources/dispatch/dispatches.js";
import { UserShort } from "../../resources/users/users.js";
import { UnitStatus } from "../../resources/dispatch/units.js";
import { PaginationResponse } from "../../resources/common/database/database.js";
import { Unit } from "../../resources/dispatch/units.js";
import { UNIT_STATUS } from "../../resources/dispatch/units.js";
import { PaginationRequest } from "../../resources/common/database/database.js";
// Common

/**
 * @generated from protobuf message services.centrum.ListDispatchActivityRequest
 */
export interface ListDispatchActivityRequest {
    /**
     * @generated from protobuf field: resources.common.database.PaginationRequest pagination = 1;
     */
    pagination?: PaginationRequest;
    /**
     * @generated from protobuf field: uint64 id = 2;
     */
    id: bigint;
}
/**
 * @generated from protobuf message services.centrum.ListUnitActivityRequest
 */
export interface ListUnitActivityRequest {
    /**
     * @generated from protobuf field: resources.common.database.PaginationRequest pagination = 1;
     */
    pagination?: PaginationRequest;
    /**
     * @generated from protobuf field: uint64 id = 2;
     */
    id: bigint;
}
/**
 * @generated from protobuf message services.centrum.GetSettingsRequest
 */
export interface GetSettingsRequest {
}
// Unit Management

/**
 * @generated from protobuf message services.centrum.ListUnitsRequest
 */
export interface ListUnitsRequest {
    /**
     * @generated from protobuf field: repeated resources.dispatch.UNIT_STATUS status = 1;
     */
    status: UNIT_STATUS[];
}
/**
 * @generated from protobuf message services.centrum.ListUnitsResponse
 */
export interface ListUnitsResponse {
    /**
     * @generated from protobuf field: repeated resources.dispatch.Unit units = 1;
     */
    units: Unit[];
}
/**
 * @generated from protobuf message services.centrum.CreateOrUpdateUnitRequest
 */
export interface CreateOrUpdateUnitRequest {
    /**
     * @generated from protobuf field: resources.dispatch.Unit unit = 1;
     */
    unit?: Unit;
}
/**
 * @generated from protobuf message services.centrum.CreateOrUpdateUnitResponse
 */
export interface CreateOrUpdateUnitResponse {
    /**
     * @generated from protobuf field: resources.dispatch.Unit unit = 1;
     */
    unit?: Unit;
}
/**
 * @generated from protobuf message services.centrum.DeleteUnitRequest
 */
export interface DeleteUnitRequest {
    /**
     * @generated from protobuf field: uint64 unit_id = 1;
     */
    unitId: bigint;
}
/**
 * @generated from protobuf message services.centrum.DeleteUnitResponse
 */
export interface DeleteUnitResponse {
}
/**
 * @generated from protobuf message services.centrum.UpdateUnitStatusRequest
 */
export interface UpdateUnitStatusRequest {
    /**
     * @generated from protobuf field: uint64 unit_id = 1;
     */
    unitId: bigint;
    /**
     * @generated from protobuf field: resources.dispatch.UNIT_STATUS status = 2;
     */
    status: UNIT_STATUS;
    /**
     * @generated from protobuf field: optional string reason = 3;
     */
    reason?: string;
    /**
     * @generated from protobuf field: optional string code = 4;
     */
    code?: string;
}
/**
 * @generated from protobuf message services.centrum.UpdateUnitStatusResponse
 */
export interface UpdateUnitStatusResponse {
}
/**
 * @generated from protobuf message services.centrum.AssignUnitRequest
 */
export interface AssignUnitRequest {
    /**
     * @generated from protobuf field: uint64 unit_id = 1;
     */
    unitId: bigint;
    /**
     * @generated from protobuf field: repeated int32 to_add = 2;
     */
    toAdd: number[];
    /**
     * @generated from protobuf field: repeated int32 to_remove = 3;
     */
    toRemove: number[];
}
/**
 * @generated from protobuf message services.centrum.AssignUnitResponse
 */
export interface AssignUnitResponse {
}
/**
 * @generated from protobuf message services.centrum.ListUnitActivityResponse
 */
export interface ListUnitActivityResponse {
    /**
     * @generated from protobuf field: resources.common.database.PaginationResponse pagination = 1;
     */
    pagination?: PaginationResponse;
    /**
     * @generated from protobuf field: repeated resources.dispatch.UnitStatus activity = 2;
     */
    activity: UnitStatus[];
}
// Dispatch Management

/**
 * @generated from protobuf message services.centrum.TakeControlRequest
 */
export interface TakeControlRequest {
    /**
     * @generated from protobuf field: bool signon = 1;
     */
    signon: boolean;
}
/**
 * @generated from protobuf message services.centrum.TakeControlResponse
 */
export interface TakeControlResponse {
}
/**
 * @generated from protobuf message services.centrum.DisponentsChange
 */
export interface DisponentsChange {
    /**
     * @generated from protobuf field: repeated resources.users.UserShort disponents = 1;
     */
    disponents: UserShort[];
    /**
     * @generated from protobuf field: bool active = 2;
     */
    active: boolean;
}
/**
 * @generated from protobuf message services.centrum.ListDispatchesRequest
 */
export interface ListDispatchesRequest {
    /**
     * @generated from protobuf field: repeated resources.dispatch.DISPATCH_STATUS status = 1;
     */
    status: DISPATCH_STATUS[];
}
/**
 * @generated from protobuf message services.centrum.ListDispatchesResponse
 */
export interface ListDispatchesResponse {
    /**
     * @generated from protobuf field: repeated resources.dispatch.Dispatch dispatches = 1;
     */
    dispatches: Dispatch[];
}
/**
 * @generated from protobuf message services.centrum.CreateDispatchRequest
 */
export interface CreateDispatchRequest {
    /**
     * @generated from protobuf field: resources.dispatch.Dispatch dispatch = 1;
     */
    dispatch?: Dispatch;
}
/**
 * @generated from protobuf message services.centrum.CreateDispatchResponse
 */
export interface CreateDispatchResponse {
    /**
     * @generated from protobuf field: resources.dispatch.Dispatch dispatch = 1;
     */
    dispatch?: Dispatch;
}
/**
 * @generated from protobuf message services.centrum.UpdateDispatchRequest
 */
export interface UpdateDispatchRequest {
    /**
     * @generated from protobuf field: resources.dispatch.Dispatch dispatch = 1;
     */
    dispatch?: Dispatch;
}
/**
 * @generated from protobuf message services.centrum.UpdateDispatchResponse
 */
export interface UpdateDispatchResponse {
}
/**
 * @generated from protobuf message services.centrum.UpdateDispatchStatusRequest
 */
export interface UpdateDispatchStatusRequest {
    /**
     * @generated from protobuf field: uint64 dispatch_id = 1;
     */
    dispatchId: bigint;
    /**
     * @generated from protobuf field: resources.dispatch.DISPATCH_STATUS status = 2;
     */
    status: DISPATCH_STATUS;
    /**
     * @generated from protobuf field: optional string reason = 3;
     */
    reason?: string;
    /**
     * @generated from protobuf field: optional string code = 4;
     */
    code?: string;
}
/**
 * @generated from protobuf message services.centrum.UpdateDispatchStatusResponse
 */
export interface UpdateDispatchStatusResponse {
}
/**
 * @generated from protobuf message services.centrum.AssignDispatchRequest
 */
export interface AssignDispatchRequest {
    /**
     * @generated from protobuf field: uint64 dispatch_id = 1;
     */
    dispatchId: bigint;
    /**
     * @generated from protobuf field: repeated uint64 to_add = 2;
     */
    toAdd: bigint[];
    /**
     * @generated from protobuf field: repeated uint64 to_remove = 3;
     */
    toRemove: bigint[];
}
/**
 * @generated from protobuf message services.centrum.AssignDispatchResponse
 */
export interface AssignDispatchResponse {
}
/**
 * @generated from protobuf message services.centrum.ListDispatchActivityResponse
 */
export interface ListDispatchActivityResponse {
    /**
     * @generated from protobuf field: resources.common.database.PaginationResponse pagination = 1;
     */
    pagination?: PaginationResponse;
    /**
     * @generated from protobuf field: repeated resources.dispatch.DispatchStatus activity = 2;
     */
    activity: DispatchStatus[];
}
/**
 * @generated from protobuf message services.centrum.TakeDispatchRequest
 */
export interface TakeDispatchRequest {
    /**
     * @generated from protobuf field: uint64 dispatch_id = 1;
     */
    dispatchId: bigint;
    /**
     * @generated from protobuf field: services.centrum.TAKE_DISPATCH_RESP resp = 2;
     */
    resp: TAKE_DISPATCH_RESP;
    /**
     * @generated from protobuf field: optional string reason = 3;
     */
    reason?: string;
}
/**
 * @generated from protobuf message services.centrum.TakeDispatchResponse
 */
export interface TakeDispatchResponse {
}
/**
 * @generated from protobuf message services.centrum.JoinUnitRequest
 */
export interface JoinUnitRequest {
    /**
     * @generated from protobuf field: uint64 unit_id = 1;
     */
    unitId: bigint;
    /**
     * @generated from protobuf field: optional bool leave = 2;
     */
    leave?: boolean;
}
/**
 * @generated from protobuf message services.centrum.JoinUnitResponse
 */
export interface JoinUnitResponse {
    /**
     * @generated from protobuf field: resources.dispatch.Unit unit = 1;
     */
    unit?: Unit;
}
/**
 * @generated from protobuf message services.centrum.Initial
 */
export interface Initial {
    /**
     * @generated from protobuf field: bool is_disponent = 1;
     */
    isDisponent: boolean;
    /**
     * @generated from protobuf field: resources.dispatch.Settings settings = 2;
     */
    settings?: Settings;
    /**
     * @generated from protobuf field: resources.dispatch.Unit unit = 3;
     */
    unit?: Unit;
    /**
     * Send the current units and dispatches
     *
     * @generated from protobuf field: repeated resources.dispatch.Unit units = 4;
     */
    units: Unit[];
    /**
     * @generated from protobuf field: repeated resources.dispatch.Dispatch dispatches = 5;
     */
    dispatches: Dispatch[];
}
/**
 * @generated from protobuf message services.centrum.StreamRequest
 */
export interface StreamRequest {
}
/**
 * @generated from protobuf message services.centrum.StreamResponse
 */
export interface StreamResponse {
    /**
     * @generated from protobuf oneof: change
     */
    change: {
        oneofKind: "initial";
        /**
         * @generated from protobuf field: services.centrum.Initial initial = 1;
         */
        initial: Initial;
    } | {
        oneofKind: "settings";
        /**
         * @generated from protobuf field: resources.dispatch.Settings settings = 2;
         */
        settings: Settings;
    } | {
        oneofKind: "disponents";
        /**
         * @generated from protobuf field: services.centrum.DisponentsChange disponents = 3;
         */
        disponents: DisponentsChange;
    } | {
        oneofKind: "dispatchStatus";
        /**
         * @generated from protobuf field: resources.dispatch.DispatchStatus dispatch_status = 4;
         */
        dispatchStatus: DispatchStatus;
    } | {
        oneofKind: "dispatchUpdate";
        /**
         * @generated from protobuf field: resources.dispatch.Dispatch dispatch_update = 5;
         */
        dispatchUpdate: Dispatch;
    } | {
        oneofKind: "dispatchAssigned";
        /**
         * @generated from protobuf field: resources.dispatch.Dispatch dispatch_assigned = 6;
         */
        dispatchAssigned: Dispatch;
    } | {
        oneofKind: "dispatchUnassigned";
        /**
         * @generated from protobuf field: resources.dispatch.Dispatch dispatch_unassigned = 7;
         */
        dispatchUnassigned: Dispatch;
    } | {
        oneofKind: "unitStatus";
        /**
         * @generated from protobuf field: resources.dispatch.UnitStatus unit_status = 8;
         */
        unitStatus: UnitStatus;
    } | {
        oneofKind: "unitUpdate";
        /**
         * @generated from protobuf field: resources.dispatch.Unit unit_update = 9;
         */
        unitUpdate: Unit;
    } | {
        oneofKind: "unitAssigned";
        /**
         * @generated from protobuf field: resources.dispatch.Unit unit_assigned = 10;
         */
        unitAssigned: Unit;
    } | {
        oneofKind: "unitDeleted";
        /**
         * @generated from protobuf field: uint64 unit_deleted = 11;
         */
        unitDeleted: bigint;
    } | {
        oneofKind: undefined;
    };
}
/**
 * @generated from protobuf enum services.centrum.TAKE_DISPATCH_RESP
 */
export enum TAKE_DISPATCH_RESP {
    /**
     * @generated from protobuf enum value: TIMEOUT = 0;
     */
    TIMEOUT = 0,
    /**
     * @generated from protobuf enum value: ACCEPTED = 1;
     */
    ACCEPTED = 1,
    /**
     * @generated from protobuf enum value: DECLINED = 2;
     */
    DECLINED = 2
}
// @generated message type with reflection information, may provide speed optimized methods
class ListDispatchActivityRequest$Type extends MessageType<ListDispatchActivityRequest> {
    constructor() {
        super("services.centrum.ListDispatchActivityRequest", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationRequest, options: { "validate.rules": { message: { required: true } } } },
            { no: 2, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.ListDispatchActivityRequest
 */
export const ListDispatchActivityRequest = new ListDispatchActivityRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListUnitActivityRequest$Type extends MessageType<ListUnitActivityRequest> {
    constructor() {
        super("services.centrum.ListUnitActivityRequest", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationRequest, options: { "validate.rules": { message: { required: true } } } },
            { no: 2, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.ListUnitActivityRequest
 */
export const ListUnitActivityRequest = new ListUnitActivityRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetSettingsRequest$Type extends MessageType<GetSettingsRequest> {
    constructor() {
        super("services.centrum.GetSettingsRequest", []);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.GetSettingsRequest
 */
export const GetSettingsRequest = new GetSettingsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListUnitsRequest$Type extends MessageType<ListUnitsRequest> {
    constructor() {
        super("services.centrum.ListUnitsRequest", [
            { no: 1, name: "status", kind: "enum", repeat: 1 /*RepeatType.PACKED*/, T: () => ["resources.dispatch.UNIT_STATUS", UNIT_STATUS], options: { "validate.rules": { repeated: { items: { enum: { definedOnly: true } } } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.ListUnitsRequest
 */
export const ListUnitsRequest = new ListUnitsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListUnitsResponse$Type extends MessageType<ListUnitsResponse> {
    constructor() {
        super("services.centrum.ListUnitsResponse", [
            { no: 1, name: "units", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Unit }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.ListUnitsResponse
 */
export const ListUnitsResponse = new ListUnitsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateOrUpdateUnitRequest$Type extends MessageType<CreateOrUpdateUnitRequest> {
    constructor() {
        super("services.centrum.CreateOrUpdateUnitRequest", [
            { no: 1, name: "unit", kind: "message", T: () => Unit, options: { "validate.rules": { message: { required: true } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.CreateOrUpdateUnitRequest
 */
export const CreateOrUpdateUnitRequest = new CreateOrUpdateUnitRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateOrUpdateUnitResponse$Type extends MessageType<CreateOrUpdateUnitResponse> {
    constructor() {
        super("services.centrum.CreateOrUpdateUnitResponse", [
            { no: 1, name: "unit", kind: "message", T: () => Unit }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.CreateOrUpdateUnitResponse
 */
export const CreateOrUpdateUnitResponse = new CreateOrUpdateUnitResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteUnitRequest$Type extends MessageType<DeleteUnitRequest> {
    constructor() {
        super("services.centrum.DeleteUnitRequest", [
            { no: 1, name: "unit_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.DeleteUnitRequest
 */
export const DeleteUnitRequest = new DeleteUnitRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteUnitResponse$Type extends MessageType<DeleteUnitResponse> {
    constructor() {
        super("services.centrum.DeleteUnitResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.DeleteUnitResponse
 */
export const DeleteUnitResponse = new DeleteUnitResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateUnitStatusRequest$Type extends MessageType<UpdateUnitStatusRequest> {
    constructor() {
        super("services.centrum.UpdateUnitStatusRequest", [
            { no: 1, name: "unit_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "status", kind: "enum", T: () => ["resources.dispatch.UNIT_STATUS", UNIT_STATUS], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 3, name: "reason", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "code", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.UpdateUnitStatusRequest
 */
export const UpdateUnitStatusRequest = new UpdateUnitStatusRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateUnitStatusResponse$Type extends MessageType<UpdateUnitStatusResponse> {
    constructor() {
        super("services.centrum.UpdateUnitStatusResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.UpdateUnitStatusResponse
 */
export const UpdateUnitStatusResponse = new UpdateUnitStatusResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class AssignUnitRequest$Type extends MessageType<AssignUnitRequest> {
    constructor() {
        super("services.centrum.AssignUnitRequest", [
            { no: 1, name: "unit_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "to_add", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "to_remove", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 5 /*ScalarType.INT32*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.AssignUnitRequest
 */
export const AssignUnitRequest = new AssignUnitRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class AssignUnitResponse$Type extends MessageType<AssignUnitResponse> {
    constructor() {
        super("services.centrum.AssignUnitResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.AssignUnitResponse
 */
export const AssignUnitResponse = new AssignUnitResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListUnitActivityResponse$Type extends MessageType<ListUnitActivityResponse> {
    constructor() {
        super("services.centrum.ListUnitActivityResponse", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationResponse },
            { no: 2, name: "activity", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => UnitStatus }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.ListUnitActivityResponse
 */
export const ListUnitActivityResponse = new ListUnitActivityResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TakeControlRequest$Type extends MessageType<TakeControlRequest> {
    constructor() {
        super("services.centrum.TakeControlRequest", [
            { no: 1, name: "signon", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.TakeControlRequest
 */
export const TakeControlRequest = new TakeControlRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TakeControlResponse$Type extends MessageType<TakeControlResponse> {
    constructor() {
        super("services.centrum.TakeControlResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.TakeControlResponse
 */
export const TakeControlResponse = new TakeControlResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DisponentsChange$Type extends MessageType<DisponentsChange> {
    constructor() {
        super("services.centrum.DisponentsChange", [
            { no: 1, name: "disponents", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => UserShort },
            { no: 2, name: "active", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.DisponentsChange
 */
export const DisponentsChange = new DisponentsChange$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListDispatchesRequest$Type extends MessageType<ListDispatchesRequest> {
    constructor() {
        super("services.centrum.ListDispatchesRequest", [
            { no: 1, name: "status", kind: "enum", repeat: 1 /*RepeatType.PACKED*/, T: () => ["resources.dispatch.DISPATCH_STATUS", DISPATCH_STATUS], options: { "validate.rules": { repeated: { items: { enum: { definedOnly: true } } } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.ListDispatchesRequest
 */
export const ListDispatchesRequest = new ListDispatchesRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListDispatchesResponse$Type extends MessageType<ListDispatchesResponse> {
    constructor() {
        super("services.centrum.ListDispatchesResponse", [
            { no: 1, name: "dispatches", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Dispatch }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.ListDispatchesResponse
 */
export const ListDispatchesResponse = new ListDispatchesResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateDispatchRequest$Type extends MessageType<CreateDispatchRequest> {
    constructor() {
        super("services.centrum.CreateDispatchRequest", [
            { no: 1, name: "dispatch", kind: "message", T: () => Dispatch, options: { "validate.rules": { message: { required: true } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.CreateDispatchRequest
 */
export const CreateDispatchRequest = new CreateDispatchRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateDispatchResponse$Type extends MessageType<CreateDispatchResponse> {
    constructor() {
        super("services.centrum.CreateDispatchResponse", [
            { no: 1, name: "dispatch", kind: "message", T: () => Dispatch }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.CreateDispatchResponse
 */
export const CreateDispatchResponse = new CreateDispatchResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateDispatchRequest$Type extends MessageType<UpdateDispatchRequest> {
    constructor() {
        super("services.centrum.UpdateDispatchRequest", [
            { no: 1, name: "dispatch", kind: "message", T: () => Dispatch, options: { "validate.rules": { message: { required: true } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.UpdateDispatchRequest
 */
export const UpdateDispatchRequest = new UpdateDispatchRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateDispatchResponse$Type extends MessageType<UpdateDispatchResponse> {
    constructor() {
        super("services.centrum.UpdateDispatchResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.UpdateDispatchResponse
 */
export const UpdateDispatchResponse = new UpdateDispatchResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateDispatchStatusRequest$Type extends MessageType<UpdateDispatchStatusRequest> {
    constructor() {
        super("services.centrum.UpdateDispatchStatusRequest", [
            { no: 1, name: "dispatch_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "status", kind: "enum", T: () => ["resources.dispatch.DISPATCH_STATUS", DISPATCH_STATUS], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 3, name: "reason", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "code", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.UpdateDispatchStatusRequest
 */
export const UpdateDispatchStatusRequest = new UpdateDispatchStatusRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateDispatchStatusResponse$Type extends MessageType<UpdateDispatchStatusResponse> {
    constructor() {
        super("services.centrum.UpdateDispatchStatusResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.UpdateDispatchStatusResponse
 */
export const UpdateDispatchStatusResponse = new UpdateDispatchStatusResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class AssignDispatchRequest$Type extends MessageType<AssignDispatchRequest> {
    constructor() {
        super("services.centrum.AssignDispatchRequest", [
            { no: 1, name: "dispatch_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "to_add", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 3, name: "to_remove", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.AssignDispatchRequest
 */
export const AssignDispatchRequest = new AssignDispatchRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class AssignDispatchResponse$Type extends MessageType<AssignDispatchResponse> {
    constructor() {
        super("services.centrum.AssignDispatchResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.AssignDispatchResponse
 */
export const AssignDispatchResponse = new AssignDispatchResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListDispatchActivityResponse$Type extends MessageType<ListDispatchActivityResponse> {
    constructor() {
        super("services.centrum.ListDispatchActivityResponse", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationResponse },
            { no: 2, name: "activity", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => DispatchStatus }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.ListDispatchActivityResponse
 */
export const ListDispatchActivityResponse = new ListDispatchActivityResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TakeDispatchRequest$Type extends MessageType<TakeDispatchRequest> {
    constructor() {
        super("services.centrum.TakeDispatchRequest", [
            { no: 1, name: "dispatch_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "resp", kind: "enum", T: () => ["services.centrum.TAKE_DISPATCH_RESP", TAKE_DISPATCH_RESP], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 3, name: "reason", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.TakeDispatchRequest
 */
export const TakeDispatchRequest = new TakeDispatchRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TakeDispatchResponse$Type extends MessageType<TakeDispatchResponse> {
    constructor() {
        super("services.centrum.TakeDispatchResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.TakeDispatchResponse
 */
export const TakeDispatchResponse = new TakeDispatchResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class JoinUnitRequest$Type extends MessageType<JoinUnitRequest> {
    constructor() {
        super("services.centrum.JoinUnitRequest", [
            { no: 1, name: "unit_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "leave", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.JoinUnitRequest
 */
export const JoinUnitRequest = new JoinUnitRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class JoinUnitResponse$Type extends MessageType<JoinUnitResponse> {
    constructor() {
        super("services.centrum.JoinUnitResponse", [
            { no: 1, name: "unit", kind: "message", T: () => Unit }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.JoinUnitResponse
 */
export const JoinUnitResponse = new JoinUnitResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Initial$Type extends MessageType<Initial> {
    constructor() {
        super("services.centrum.Initial", [
            { no: 1, name: "is_disponent", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "settings", kind: "message", T: () => Settings },
            { no: 3, name: "unit", kind: "message", T: () => Unit },
            { no: 4, name: "units", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Unit },
            { no: 5, name: "dispatches", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Dispatch }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.Initial
 */
export const Initial = new Initial$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StreamRequest$Type extends MessageType<StreamRequest> {
    constructor() {
        super("services.centrum.StreamRequest", []);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.StreamRequest
 */
export const StreamRequest = new StreamRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StreamResponse$Type extends MessageType<StreamResponse> {
    constructor() {
        super("services.centrum.StreamResponse", [
            { no: 1, name: "initial", kind: "message", oneof: "change", T: () => Initial },
            { no: 2, name: "settings", kind: "message", oneof: "change", T: () => Settings },
            { no: 3, name: "disponents", kind: "message", oneof: "change", T: () => DisponentsChange },
            { no: 4, name: "dispatch_status", kind: "message", oneof: "change", T: () => DispatchStatus },
            { no: 5, name: "dispatch_update", kind: "message", oneof: "change", T: () => Dispatch },
            { no: 6, name: "dispatch_assigned", kind: "message", oneof: "change", T: () => Dispatch },
            { no: 7, name: "dispatch_unassigned", kind: "message", oneof: "change", T: () => Dispatch },
            { no: 8, name: "unit_status", kind: "message", oneof: "change", T: () => UnitStatus },
            { no: 9, name: "unit_update", kind: "message", oneof: "change", T: () => Unit },
            { no: 10, name: "unit_assigned", kind: "message", oneof: "change", T: () => Unit },
            { no: 11, name: "unit_deleted", kind: "scalar", oneof: "change", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.StreamResponse
 */
export const StreamResponse = new StreamResponse$Type();
/**
 * @generated ServiceType for protobuf service services.centrum.CentrumService
 */
export const CentrumService = new ServiceType("services.centrum.CentrumService", [
    { name: "GetSettings", options: {}, I: GetSettingsRequest, O: Settings },
    { name: "UpdateSettings", options: {}, I: Settings, O: Settings },
    { name: "TakeControl", options: {}, I: TakeControlRequest, O: TakeControlResponse },
    { name: "ListUnits", options: {}, I: ListUnitsRequest, O: ListUnitsResponse },
    { name: "CreateOrUpdateUnit", options: {}, I: CreateOrUpdateUnitRequest, O: CreateOrUpdateUnitResponse },
    { name: "DeleteUnit", options: {}, I: DeleteUnitRequest, O: DeleteUnitResponse },
    { name: "UpdateUnitStatus", options: {}, I: UpdateUnitStatusRequest, O: UpdateUnitStatusResponse },
    { name: "AssignUnit", options: {}, I: AssignUnitRequest, O: AssignUnitResponse },
    { name: "ListUnitActivity", options: {}, I: ListUnitActivityRequest, O: ListUnitActivityResponse },
    { name: "ListDispatches", options: {}, I: ListDispatchesRequest, O: ListDispatchesResponse },
    { name: "CreateDispatch", options: {}, I: CreateDispatchRequest, O: CreateDispatchResponse },
    { name: "UpdateDispatch", options: {}, I: UpdateDispatchRequest, O: UpdateDispatchResponse },
    { name: "UpdateDispatchStatus", options: {}, I: UpdateDispatchStatusRequest, O: UpdateDispatchStatusResponse },
    { name: "ListDispatchActivity", options: {}, I: ListDispatchActivityRequest, O: ListDispatchActivityResponse },
    { name: "AssignDispatch", options: {}, I: AssignDispatchRequest, O: AssignDispatchResponse },
    { name: "JoinUnit", options: {}, I: JoinUnitRequest, O: JoinUnitResponse },
    { name: "TakeDispatch", options: {}, I: TakeDispatchRequest, O: TakeDispatchResponse },
    { name: "Stream", serverStreaming: true, options: {}, I: StreamRequest, O: StreamResponse }
]);
