// @generated by protobuf-ts 2.9.1 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "resources/laws/laws.proto" (package "resources.laws", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { Timestamp } from "../timestamp/timestamp.js";
/**
 * @generated from protobuf message resources.laws.LawBook
 */
export interface LawBook {
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
     * @sanitize
     *
     * @generated from protobuf field: string name = 4;
     */
    name: string;
    /**
     * @sanitize
     *
     * @generated from protobuf field: optional string description = 5;
     */
    description?: string;
    /**
     * @generated from protobuf field: repeated resources.laws.Law laws = 6;
     */
    laws: Law[];
}
/**
 * @generated from protobuf message resources.laws.Law
 */
export interface Law {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint; // @gotags: sql:"primary_key" alias:"law.id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp;
    /**
     * @generated from protobuf field: uint64 lawbook_id = 4;
     */
    lawbookId: bigint;
    /**
     * @sanitize
     *
     * @generated from protobuf field: string name = 5;
     */
    name: string;
    /**
     * @sanitize
     *
     * @generated from protobuf field: optional string description = 6;
     */
    description?: string;
    /**
     * @generated from protobuf field: optional uint64 fine = 7;
     */
    fine?: bigint;
    /**
     * @generated from protobuf field: optional uint64 detention_time = 8;
     */
    detentionTime?: bigint;
    /**
     * @generated from protobuf field: optional uint64 stvo_points = 9;
     */
    stvoPoints?: bigint;
}
// @generated message type with reflection information, may provide speed optimized methods
class LawBook$Type extends MessageType<LawBook> {
    constructor() {
        super("resources.laws.LawBook", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3", maxLen: "128" } } } },
            { no: 5, name: "description", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 6, name: "laws", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Law }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.laws.LawBook
 */
export const LawBook = new LawBook$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Law$Type extends MessageType<Law> {
    constructor() {
        super("resources.laws.Law", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "lawbook_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 5, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3", maxLen: "128" } } } },
            { no: 6, name: "description", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "511" } } } },
            { no: 7, name: "fine", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 8, name: "detention_time", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 9, name: "stvo_points", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.laws.Law
 */
export const Law = new Law$Type();