// @generated by protobuf-ts 2.9.1 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "resources/dispatch/units.proto" (package "resources.dispatch", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { UserShort } from "../users/users.js";
import { Timestamp } from "../timestamp/timestamp.js";
/**
 * @generated from protobuf message resources.dispatch.Unit
 */
export interface Unit {
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
     * @generated from protobuf field: string name = 5;
     */
    name: string;
    /**
     * @generated from protobuf field: string initials = 6;
     */
    initials: string;
    /**
     * @generated from protobuf field: optional string color = 7;
     */
    color?: string;
    /**
     * @generated from protobuf field: optional string description = 8;
     */
    description?: string;
    /**
     * @generated from protobuf field: optional resources.dispatch.UnitStatus status = 9;
     */
    status?: UnitStatus;
    /**
     * repeated UnitStatus statuses = 10;
     *
     * @generated from protobuf field: repeated resources.dispatch.UnitAssignment users = 11;
     */
    users: UnitAssignment[];
}
/**
 * @generated from protobuf message resources.dispatch.UnitAssignments
 */
export interface UnitAssignments {
    /**
     * @generated from protobuf field: uint64 unit_id = 1;
     */
    unitId: bigint;
    /**
     * @generated from protobuf field: string job = 2;
     */
    job: string;
    /**
     * @generated from protobuf field: repeated resources.dispatch.UnitAssignment users = 3;
     */
    users: UnitAssignment[];
}
/**
 * @generated from protobuf message resources.dispatch.UnitAssignment
 */
export interface UnitAssignment {
    /**
     * @generated from protobuf field: uint64 unit_id = 1;
     */
    unitId: bigint; // @gotags: sql:"primary_key" alias:"unit_id"
    /**
     * @generated from protobuf field: int32 user_id = 2;
     */
    userId: number; // @gotags: sql:"primary_key" alias:"user_id"
    /**
     * @generated from protobuf field: optional resources.users.UserShort user = 3;
     */
    user?: UserShort;
}
/**
 * @generated from protobuf message resources.dispatch.UnitStatus
 */
export interface UnitStatus {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint; // @gotags: sql:"primary_key" alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: uint64 unit_id = 3;
     */
    unitId: bigint;
    /**
     * @generated from protobuf field: resources.dispatch.UNIT_STATUS status = 4;
     */
    status: UNIT_STATUS;
    /**
     * @generated from protobuf field: optional string reason = 5;
     */
    reason?: string;
    /**
     * @generated from protobuf field: optional string code = 6;
     */
    code?: string;
    /**
     * @generated from protobuf field: optional int32 user_id = 7;
     */
    userId?: number;
    /**
     * @generated from protobuf field: optional resources.users.UserShort user = 8;
     */
    user?: UserShort;
    /**
     * @generated from protobuf field: optional double x = 9;
     */
    x?: number;
    /**
     * @generated from protobuf field: optional double y = 10;
     */
    y?: number;
    /**
     * @generated from protobuf field: optional string postal = 11;
     */
    postal?: string;
    /**
     * @generated from protobuf field: optional int32 creator_id = 12;
     */
    creatorId?: number;
    /**
     * @generated from protobuf field: optional resources.users.UserShort creator = 13;
     */
    creator?: UserShort;
}
/**
 * @generated from protobuf enum resources.dispatch.UNIT_STATUS
 */
export enum UNIT_STATUS {
    /**
     * @generated from protobuf enum value: UNKNOWN = 0;
     */
    UNKNOWN = 0,
    /**
     * @generated from protobuf enum value: USER_ADDED = 1;
     */
    USER_ADDED = 1,
    /**
     * @generated from protobuf enum value: USER_REMOVED = 2;
     */
    USER_REMOVED = 2,
    /**
     * @generated from protobuf enum value: UNAVAILABLE = 3;
     */
    UNAVAILABLE = 3,
    /**
     * @generated from protobuf enum value: AVAILABLE = 4;
     */
    AVAILABLE = 4,
    /**
     * @generated from protobuf enum value: ON_BREAK = 5;
     */
    ON_BREAK = 5,
    /**
     * @generated from protobuf enum value: BUSY = 6;
     */
    BUSY = 6
}
// @generated message type with reflection information, may provide speed optimized methods
class Unit$Type extends MessageType<Unit> {
    constructor() {
        super("resources.dispatch.Unit", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 5, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3", maxLen: "24" } } } },
            { no: 6, name: "initials", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "2", maxLen: "4" } } } },
            { no: 7, name: "color", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "6" } } } },
            { no: 8, name: "description", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 9, name: "status", kind: "message", T: () => UnitStatus },
            { no: 11, name: "users", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => UnitAssignment }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.dispatch.Unit
 */
export const Unit = new Unit$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UnitAssignments$Type extends MessageType<UnitAssignments> {
    constructor() {
        super("resources.dispatch.UnitAssignments", [
            { no: 1, name: "unit_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 3, name: "users", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => UnitAssignment }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.dispatch.UnitAssignments
 */
export const UnitAssignments = new UnitAssignments$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UnitAssignment$Type extends MessageType<UnitAssignment> {
    constructor() {
        super("resources.dispatch.UnitAssignment", [
            { no: 1, name: "unit_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "user", kind: "message", T: () => UserShort }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.dispatch.UnitAssignment
 */
export const UnitAssignment = new UnitAssignment$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UnitStatus$Type extends MessageType<UnitStatus> {
    constructor() {
        super("resources.dispatch.UnitStatus", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "unit_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 4, name: "status", kind: "enum", T: () => ["resources.dispatch.UNIT_STATUS", UNIT_STATUS], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 5, name: "reason", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 6, name: "code", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 7, name: "user_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 8, name: "user", kind: "message", T: () => UserShort },
            { no: 9, name: "x", kind: "scalar", opt: true, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 10, name: "y", kind: "scalar", opt: true, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 11, name: "postal", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "48" } } } },
            { no: 12, name: "creator_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 13, name: "creator", kind: "message", T: () => UserShort }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.dispatch.UnitStatus
 */
export const UnitStatus = new UnitStatus$Type();
