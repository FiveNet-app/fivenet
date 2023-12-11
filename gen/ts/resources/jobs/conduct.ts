// @generated by protobuf-ts 2.9.3 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "resources/jobs/conduct.proto" (package "resources.jobs", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { UserShort } from "../users/users";
import { Timestamp } from "../timestamp/timestamp";
/**
 * @generated from protobuf message resources.jobs.ConductEntry
 */
export interface ConductEntry {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string; // @gotags: sql:"primary_key" alias:"id"
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
     * @generated from protobuf field: resources.jobs.ConductType type = 5;
     */
    type: ConductType;
    /**
     * @sanitize
     *
     * @generated from protobuf field: string message = 6;
     */
    message: string;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp expires_at = 7;
     */
    expiresAt?: Timestamp;
    /**
     * @generated from protobuf field: int32 target_user_id = 8;
     */
    targetUserId: number;
    /**
     * @generated from protobuf field: optional resources.users.UserShort target_user = 9;
     */
    targetUser?: UserShort; // @gotags: alias:"target_user"
    /**
     * @generated from protobuf field: int32 creator_id = 10;
     */
    creatorId: number;
    /**
     * @generated from protobuf field: optional resources.users.UserShort creator = 11;
     */
    creator?: UserShort; // @gotags: alias:"creator"
}
/**
 * @generated from protobuf enum resources.jobs.ConductType
 */
export enum ConductType {
    /**
     * @generated from protobuf enum value: CONDUCT_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: CONDUCT_TYPE_NEUTRAL = 1;
     */
    NEUTRAL = 1,
    /**
     * @generated from protobuf enum value: CONDUCT_TYPE_POSITIVE = 2;
     */
    POSITIVE = 2,
    /**
     * @generated from protobuf enum value: CONDUCT_TYPE_NEGATIVE = 3;
     */
    NEGATIVE = 3,
    /**
     * @generated from protobuf enum value: CONDUCT_TYPE_WARNING = 4;
     */
    WARNING = 4,
    /**
     * @generated from protobuf enum value: CONDUCT_TYPE_SUSPENSION = 5;
     */
    SUSPENSION = 5
}
// @generated message type with reflection information, may provide speed optimized methods
class ConductEntry$Type extends MessageType<ConductEntry> {
    constructor() {
        super("resources.jobs.ConductEntry", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 5, name: "type", kind: "enum", T: () => ["resources.jobs.ConductType", ConductType, "CONDUCT_TYPE_"], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 6, name: "message", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3", maxLen: "2048" } } } },
            { no: 7, name: "expires_at", kind: "message", T: () => Timestamp },
            { no: 8, name: "target_user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 9, name: "target_user", kind: "message", T: () => UserShort },
            { no: 10, name: "creator_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 11, name: "creator", kind: "message", T: () => UserShort }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.jobs.ConductEntry
 */
export const ConductEntry = new ConductEntry$Type();
