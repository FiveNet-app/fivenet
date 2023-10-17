// @generated by protobuf-ts 2.9.1 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "resources/dispatch/dispatches.proto" (package "resources.dispatch", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { UserShort } from "../users/users.js";
import { Unit } from "./units.js";
import { User } from "../users/users.js";
import { Timestamp } from "../timestamp/timestamp.js";
/**
 * @generated from protobuf message resources.dispatch.Dispatch
 */
export interface Dispatch {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint; // @gotags: sql:"primary_key" alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp;
    /**
     * @generated from protobuf field: string job = 4;
     */
    job: string;
    /**
     * @generated from protobuf field: optional resources.dispatch.DispatchStatus status = 5;
     */
    status?: DispatchStatus;
    /**
     * @sanitize
     *
     * @generated from protobuf field: string message = 7;
     */
    message: string;
    /**
     * @sanitize
     *
     * @generated from protobuf field: optional string description = 8;
     */
    description?: string;
    /**
     * @generated from protobuf field: optional resources.dispatch.Attributes attributes = 9;
     */
    attributes?: Attributes;
    /**
     * @generated from protobuf field: double x = 10;
     */
    x: number;
    /**
     * @generated from protobuf field: double y = 11;
     */
    y: number;
    /**
     * @sanitize
     *
     * @generated from protobuf field: optional string postal = 12;
     */
    postal?: string;
    /**
     * @generated from protobuf field: bool anon = 13;
     */
    anon: boolean;
    /**
     * @generated from protobuf field: optional int32 creator_id = 14;
     */
    creatorId?: number;
    /**
     * @generated from protobuf field: optional resources.users.User creator = 15;
     */
    creator?: User;
    /**
     * @generated from protobuf field: repeated resources.dispatch.DispatchAssignment units = 16;
     */
    units: DispatchAssignment[];
}
/**
 * @generated from protobuf message resources.dispatch.Attributes
 */
export interface Attributes {
    /**
     * @generated from protobuf field: repeated string list = 1;
     */
    list: string[];
}
/**
 * @generated from protobuf message resources.dispatch.DispatchAssignments
 */
export interface DispatchAssignments {
    /**
     * @generated from protobuf field: uint64 dispatch_id = 1;
     */
    dispatchId: bigint;
    /**
     * @generated from protobuf field: string job = 2;
     */
    job: string;
    /**
     * @generated from protobuf field: repeated resources.dispatch.DispatchAssignment units = 3;
     */
    units: DispatchAssignment[];
}
/**
 * @generated from protobuf message resources.dispatch.DispatchAssignment
 */
export interface DispatchAssignment {
    /**
     * @generated from protobuf field: uint64 dispatch_id = 1;
     */
    dispatchId: bigint; // @gotags: sql:"primary_key" alias:"dispatch_id"
    /**
     * @generated from protobuf field: uint64 unit_id = 2;
     */
    unitId: bigint; // @gotags: sql:"primary_key" alias:"unit_id"
    /**
     * @generated from protobuf field: optional resources.dispatch.Unit unit = 3;
     */
    unit?: Unit;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 4;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp expires_at = 5;
     */
    expiresAt?: Timestamp;
}
/**
 * @generated from protobuf message resources.dispatch.DispatchStatus
 */
export interface DispatchStatus {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint; // @gotags: sql:"primary_key" alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: uint64 dispatch_id = 3;
     */
    dispatchId: bigint;
    /**
     * @generated from protobuf field: optional uint64 unit_id = 4;
     */
    unitId?: bigint;
    /**
     * @generated from protobuf field: optional resources.dispatch.Unit unit = 5;
     */
    unit?: Unit;
    /**
     * @generated from protobuf field: resources.dispatch.StatusDispatch status = 6;
     */
    status: StatusDispatch;
    /**
     * @sanitize
     *
     * @generated from protobuf field: optional string reason = 7;
     */
    reason?: string;
    /**
     * @sanitize
     *
     * @generated from protobuf field: optional string code = 8;
     */
    code?: string;
    /**
     * @generated from protobuf field: optional int32 user_id = 9;
     */
    userId?: number;
    /**
     * @generated from protobuf field: optional resources.users.UserShort user = 10;
     */
    user?: UserShort;
    /**
     * @generated from protobuf field: optional double x = 11;
     */
    x?: number;
    /**
     * @generated from protobuf field: optional double y = 12;
     */
    y?: number;
    /**
     * @sanitize
     *
     * @generated from protobuf field: optional string postal = 13;
     */
    postal?: string;
}
/**
 * @generated from protobuf enum resources.dispatch.StatusDispatch
 */
export enum StatusDispatch {
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_NEW = 1;
     */
    NEW = 1,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_UNASSIGNED = 2;
     */
    UNASSIGNED = 2,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_UPDATED = 3;
     */
    UPDATED = 3,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_UNIT_ASSIGNED = 4;
     */
    UNIT_ASSIGNED = 4,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_UNIT_UNASSIGNED = 5;
     */
    UNIT_UNASSIGNED = 5,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_UNIT_ACCEPTED = 6;
     */
    UNIT_ACCEPTED = 6,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_UNIT_DECLINED = 7;
     */
    UNIT_DECLINED = 7,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_EN_ROUTE = 8;
     */
    EN_ROUTE = 8,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_ON_SCENE = 9;
     */
    ON_SCENE = 9,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_NEED_ASSISTANCE = 10;
     */
    NEED_ASSISTANCE = 10,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_COMPLETED = 11;
     */
    COMPLETED = 11,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_CANCELLED = 12;
     */
    CANCELLED = 12,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_ARCHIVED = 13;
     */
    ARCHIVED = 13
}
// @generated message type with reflection information, may provide speed optimized methods
class Dispatch$Type extends MessageType<Dispatch> {
    constructor() {
        super("resources.dispatch.Dispatch", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 5, name: "status", kind: "message", T: () => DispatchStatus },
            { no: 7, name: "message", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 8, name: "description", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "1024" } } } },
            { no: 9, name: "attributes", kind: "message", T: () => Attributes },
            { no: 10, name: "x", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 11, name: "y", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 12, name: "postal", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "48" } } } },
            { no: 13, name: "anon", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 14, name: "creator_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 15, name: "creator", kind: "message", T: () => User },
            { no: 16, name: "units", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => DispatchAssignment }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.dispatch.Dispatch
 */
export const Dispatch = new Dispatch$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Attributes$Type extends MessageType<Attributes> {
    constructor() {
        super("resources.dispatch.Attributes", [
            { no: 1, name: "list", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.dispatch.Attributes
 */
export const Attributes = new Attributes$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DispatchAssignments$Type extends MessageType<DispatchAssignments> {
    constructor() {
        super("resources.dispatch.DispatchAssignments", [
            { no: 1, name: "dispatch_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 3, name: "units", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => DispatchAssignment }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.dispatch.DispatchAssignments
 */
export const DispatchAssignments = new DispatchAssignments$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DispatchAssignment$Type extends MessageType<DispatchAssignment> {
    constructor() {
        super("resources.dispatch.DispatchAssignment", [
            { no: 1, name: "dispatch_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "unit_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 3, name: "unit", kind: "message", T: () => Unit },
            { no: 4, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 5, name: "expires_at", kind: "message", T: () => Timestamp }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.dispatch.DispatchAssignment
 */
export const DispatchAssignment = new DispatchAssignment$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DispatchStatus$Type extends MessageType<DispatchStatus> {
    constructor() {
        super("resources.dispatch.DispatchStatus", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "dispatch_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 4, name: "unit_id", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 5, name: "unit", kind: "message", T: () => Unit },
            { no: 6, name: "status", kind: "enum", T: () => ["resources.dispatch.StatusDispatch", StatusDispatch, "STATUS_DISPATCH_"], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 7, name: "reason", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 8, name: "code", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 9, name: "user_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 10, name: "user", kind: "message", T: () => UserShort },
            { no: 11, name: "x", kind: "scalar", opt: true, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 12, name: "y", kind: "scalar", opt: true, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 13, name: "postal", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "48" } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.dispatch.DispatchStatus
 */
export const DispatchStatus = new DispatchStatus$Type();