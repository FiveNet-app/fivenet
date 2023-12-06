// @generated by protobuf-ts 2.9.2 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "resources/documents/access.proto" (package "resources.documents", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { UserShort } from "../users/users";
import { Timestamp } from "../timestamp/timestamp";
/**
 * @generated from protobuf message resources.documents.DocumentAccess
 */
export interface DocumentAccess {
    /**
     * @generated from protobuf field: repeated resources.documents.DocumentJobAccess jobs = 1;
     */
    jobs: DocumentJobAccess[]; // @gotags: alias:"job_access"
    /**
     * @generated from protobuf field: repeated resources.documents.DocumentUserAccess users = 2;
     */
    users: DocumentUserAccess[]; // @gotags: alias:"user_access"
}
/**
 * @generated from protobuf message resources.documents.DocumentJobAccess
 */
export interface DocumentJobAccess {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: uint64 document_id = 3 [jstype = JS_STRING];
     */
    documentId: string;
    /**
     * @generated from protobuf field: string job = 4;
     */
    job: string; // @gotags: alias:"job"
    /**
     * @generated from protobuf field: optional string job_label = 5;
     */
    jobLabel?: string; // @gotags: alias:"job_label"
    /**
     * @generated from protobuf field: int32 minimum_grade = 6;
     */
    minimumGrade: number; // @gotags: alias:"minimum_grade"
    /**
     * @generated from protobuf field: optional string job_grade_label = 7;
     */
    jobGradeLabel?: string; // @gotags: alias:"job_grade_label"
    /**
     * @generated from protobuf field: resources.documents.AccessLevel access = 8;
     */
    access: AccessLevel; // @gotags: alias:"access"
    /**
     * @generated from protobuf field: optional bool required = 9;
     */
    required?: boolean; // @gotags: alias:"required"
}
/**
 * @generated from protobuf message resources.documents.DocumentUserAccess
 */
export interface DocumentUserAccess {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: uint64 document_id = 3 [jstype = JS_STRING];
     */
    documentId: string;
    /**
     * @generated from protobuf field: int32 user_id = 4;
     */
    userId: number;
    /**
     * @generated from protobuf field: optional resources.users.UserShort user = 5;
     */
    user?: UserShort;
    /**
     * @generated from protobuf field: resources.documents.AccessLevel access = 6;
     */
    access: AccessLevel; // @gotags: alias:"access"
    /**
     * @generated from protobuf field: optional bool required = 7;
     */
    required?: boolean; // @gotags: alias:"required"
}
/**
 * @generated from protobuf enum resources.documents.AccessLevel
 */
export enum AccessLevel {
    /**
     * @generated from protobuf enum value: ACCESS_LEVEL_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: ACCESS_LEVEL_BLOCKED = 1;
     */
    BLOCKED = 1,
    /**
     * @generated from protobuf enum value: ACCESS_LEVEL_VIEW = 2;
     */
    VIEW = 2,
    /**
     * @generated from protobuf enum value: ACCESS_LEVEL_COMMENT = 3;
     */
    COMMENT = 3,
    /**
     * @generated from protobuf enum value: ACCESS_LEVEL_STATUS = 4;
     */
    STATUS = 4,
    /**
     * @generated from protobuf enum value: ACCESS_LEVEL_ACCESS = 5;
     */
    ACCESS = 5,
    /**
     * @generated from protobuf enum value: ACCESS_LEVEL_EDIT = 6;
     */
    EDIT = 6
}
// @generated message type with reflection information, may provide speed optimized methods
class DocumentAccess$Type extends MessageType<DocumentAccess> {
    constructor() {
        super("resources.documents.DocumentAccess", [
            { no: 1, name: "jobs", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => DocumentJobAccess },
            { no: 2, name: "users", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => DocumentUserAccess }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.DocumentAccess
 */
export const DocumentAccess = new DocumentAccess$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DocumentJobAccess$Type extends MessageType<DocumentJobAccess> {
    constructor() {
        super("resources.documents.DocumentJobAccess", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "document_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 4, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 5, name: "job_label", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 6, name: "minimum_grade", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 7, name: "job_grade_label", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 8, name: "access", kind: "enum", T: () => ["resources.documents.AccessLevel", AccessLevel, "ACCESS_LEVEL_"], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 9, name: "required", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.DocumentJobAccess
 */
export const DocumentJobAccess = new DocumentJobAccess$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DocumentUserAccess$Type extends MessageType<DocumentUserAccess> {
    constructor() {
        super("resources.documents.DocumentUserAccess", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "document_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 4, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 5, name: "user", kind: "message", T: () => UserShort },
            { no: 6, name: "access", kind: "enum", T: () => ["resources.documents.AccessLevel", AccessLevel, "ACCESS_LEVEL_"], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 7, name: "required", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.DocumentUserAccess
 */
export const DocumentUserAccess = new DocumentUserAccess$Type();
