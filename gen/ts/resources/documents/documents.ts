// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "resources/documents/documents.proto" (package "resources.documents", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { ACCESS_LEVEL } from "./access.js";
import { DocumentCategory } from "./category.js";
import { UserShort } from "../users/users.js";
import { Timestamp } from "../timestamp/timestamp.js";
/**
 * @generated from protobuf message resources.documents.DocumentComment
 */
export interface DocumentComment {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint; // @gotags: alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp; // @gotags: alias:"created_at"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp; // @gotags: alias:"updated_at"
    /**
     * @generated from protobuf field: uint64 document_id = 4;
     */
    documentId: bigint; // @gotags: alias:"document_id"
    /**
     * @generated from protobuf field: string comment = 5;
     */
    comment: string; // @gotags: alias:"comment"
    /**
     * @generated from protobuf field: optional int32 creator_id = 6;
     */
    creatorId?: number; // @gotags: alias:"creator_id"
    /**
     * @generated from protobuf field: optional resources.users.UserShort creator = 7;
     */
    creator?: UserShort; // @gotags: alias:"creator"
}
/**
 * @generated from protobuf message resources.documents.Document
 */
export interface Document {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint; // @gotags: alias:"id"
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp; // @gotags: alias:"created_at"
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp; // @gotags: alias:"updated_at"
    /**
     * @generated from protobuf field: optional uint64 category_id = 4;
     */
    categoryId?: bigint; // @gotags: alias:"category_id"
    /**
     * @generated from protobuf field: optional resources.documents.DocumentCategory category = 5;
     */
    category?: DocumentCategory; // @gotags: alias:"category"
    /**
     * @generated from protobuf field: string title = 6;
     */
    title: string; // @gotags: alias:"title"
    /**
     * @generated from protobuf field: resources.documents.DOC_CONTENT_TYPE content_type = 7;
     */
    contentType: DOC_CONTENT_TYPE; // @gotags: alias:"content_type"
    /**
     * @generated from protobuf field: string content = 8;
     */
    content: string; // @gotags: alias:"content"
    /**
     * @generated from protobuf field: string data = 9;
     */
    data: string; // @gotags: alias:"data"
    /**
     * @generated from protobuf field: optional int32 creator_id = 10;
     */
    creatorId?: number; // @gotags: alias:"creator_id"
    /**
     * @generated from protobuf field: optional resources.users.UserShort creator = 11;
     */
    creator?: UserShort; // @gotags: alias:"creator"
    /**
     * @generated from protobuf field: string state = 12;
     */
    state: string; // @gotags: alias:"state"
    /**
     * @generated from protobuf field: bool closed = 13;
     */
    closed: boolean; // @gotags: alias:"closed"
    /**
     * @generated from protobuf field: bool public = 14;
     */
    public: boolean; // @gotags: alias:"public"
}
/**
 * @generated from protobuf message resources.documents.DocumentShort
 */
export interface DocumentShort {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint; // @gotags: alias:"id"
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp; // @gotags: alias:"created_at"
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp; // @gotags: alias:"updated_at"
    /**
     * @generated from protobuf field: optional uint64 category_id = 4;
     */
    categoryId?: bigint; // @gotags: alias:"category_id"
    /**
     * @generated from protobuf field: optional resources.documents.DocumentCategory category = 5;
     */
    category?: DocumentCategory; // @gotags: alias:"category"
    /**
     * @generated from protobuf field: string title = 6;
     */
    title: string; // @gotags: alias:"title"
    /**
     * @generated from protobuf field: resources.documents.DOC_CONTENT_TYPE content_type = 7;
     */
    contentType: DOC_CONTENT_TYPE; // @gotags: alias:"content_type"
    /**
     * @generated from protobuf field: string content = 8;
     */
    content: string; // @gotags: alias:"content"
    /**
     * @generated from protobuf field: optional int32 creator_id = 9;
     */
    creatorId?: number; // @gotags: alias:"creator_id"
    /**
     * @generated from protobuf field: optional resources.users.UserShort creator = 10;
     */
    creator?: UserShort; // @gotags: alias:"creator"
    /**
     * @generated from protobuf field: string state = 11;
     */
    state: string; // @gotags: alias:"state"
    /**
     * @generated from protobuf field: bool closed = 12;
     */
    closed: boolean; // @gotags: alias:"closed"
}
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
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint; // @gotags: alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp; // @gotags: alias:"created_at"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp; // @gotags: alias:"updated_at"
    /**
     * @generated from protobuf field: uint64 document_id = 4;
     */
    documentId: bigint; // @gotags: alias:"document_id"
    /**
     * @generated from protobuf field: string job = 5;
     */
    job: string; // @gotags: alias:"job"
    /**
     * @generated from protobuf field: optional string job_label = 6;
     */
    jobLabel?: string; // @gotags: alias:"job_label"
    /**
     * @generated from protobuf field: int32 minimumGrade = 7;
     */
    minimumGrade: number; // @gotags: alias:"minimum_grade"
    /**
     * @generated from protobuf field: optional string job_grade_label = 8;
     */
    jobGradeLabel?: string; // @gotags: alias:"job_grade_label"
    /**
     * @generated from protobuf field: resources.documents.ACCESS_LEVEL access = 9;
     */
    access: ACCESS_LEVEL; // @gotags: alias:"access"
}
/**
 * @generated from protobuf message resources.documents.DocumentUserAccess
 */
export interface DocumentUserAccess {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint; // @gotags: alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp; // @gotags: alias:"created_at"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp; // @gotags: alias:"updated_at"
    /**
     * @generated from protobuf field: uint64 document_id = 4;
     */
    documentId: bigint; // @gotags: alias:"document_id"
    /**
     * @generated from protobuf field: int32 user_id = 5;
     */
    userId: number; // @gotags: alias:"user_id"
    /**
     * @generated from protobuf field: optional resources.users.UserShort user = 6;
     */
    user?: UserShort;
    /**
     * @generated from protobuf field: resources.documents.ACCESS_LEVEL access = 7;
     */
    access: ACCESS_LEVEL; // @gotags: alias:"access"
}
/**
 * @generated from protobuf message resources.documents.DocumentReference
 */
export interface DocumentReference {
    /**
     * @generated from protobuf field: optional uint64 id = 1;
     */
    id?: bigint; // @gotags: alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp; // @gotags: alias:"created_at"
    /**
     * @generated from protobuf field: uint64 source_document_id = 3;
     */
    sourceDocumentId: bigint; // @gotags: alias:"source_document_id"
    /**
     * @generated from protobuf field: optional resources.documents.DocumentShort source_document = 4;
     */
    sourceDocument?: DocumentShort; // @gotags: alias:"source_document"
    /**
     * @generated from protobuf field: resources.documents.DOC_REFERENCE reference = 5;
     */
    reference: DOC_REFERENCE; // @gotags: alias:"reference"
    /**
     * @generated from protobuf field: uint64 target_document_id = 6;
     */
    targetDocumentId: bigint; // @gotags: alias:"target_document_id"
    /**
     * @generated from protobuf field: optional resources.documents.DocumentShort target_document = 7;
     */
    targetDocument?: DocumentShort; // @gotags: alias:"target_document"
    /**
     * @generated from protobuf field: optional int32 creator_id = 8;
     */
    creatorId?: number; // @gotags: alias:"creator_id"
    /**
     * @generated from protobuf field: optional resources.users.UserShort creator = 9;
     */
    creator?: UserShort; // @gotags: alias:"ref_creator"
}
/**
 * @generated from protobuf message resources.documents.DocumentRelation
 */
export interface DocumentRelation {
    /**
     * @generated from protobuf field: optional uint64 id = 1;
     */
    id?: bigint; // @gotags: alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp; // @gotags: alias:"created_at"
    /**
     * @generated from protobuf field: uint64 document_id = 3;
     */
    documentId: bigint; // @gotags: alias:"document_id"
    /**
     * @generated from protobuf field: optional resources.documents.DocumentShort document = 4;
     */
    document?: DocumentShort; // @gotags: alias:"document"
    /**
     * @generated from protobuf field: int32 source_user_id = 5;
     */
    sourceUserId: number; // @gotags: alias:"source_user_id"
    /**
     * @generated from protobuf field: optional resources.users.UserShort source_user = 6;
     */
    sourceUser?: UserShort; // @gotags: alias:"source_user"
    /**
     * @generated from protobuf field: resources.documents.DOC_RELATION relation = 7;
     */
    relation: DOC_RELATION; // @gotags: alias:"relation"
    /**
     * @generated from protobuf field: int32 target_user_id = 8;
     */
    targetUserId: number; // @gotags: alias:"target_user_id"
    /**
     * @generated from protobuf field: optional resources.users.UserShort target_user = 9;
     */
    targetUser?: UserShort; // @gotags: alias:"target_user"
}
/**
 * @generated from protobuf enum resources.documents.DOC_CONTENT_TYPE
 */
export enum DOC_CONTENT_TYPE {
    /**
     * @generated from protobuf enum value: HTML = 0;
     */
    HTML = 0,
    /**
     * @generated from protobuf enum value: PLAIN = 1;
     */
    PLAIN = 1
}
/**
 * @generated from protobuf enum resources.documents.DOC_REFERENCE
 */
export enum DOC_REFERENCE {
    /**
     * @generated from protobuf enum value: LINKED = 0;
     */
    LINKED = 0,
    /**
     * @generated from protobuf enum value: SOLVES = 1;
     */
    SOLVES = 1,
    /**
     * @generated from protobuf enum value: CLOSES = 2;
     */
    CLOSES = 2,
    /**
     * @generated from protobuf enum value: DEPRECATES = 3;
     */
    DEPRECATES = 3
}
/**
 * @generated from protobuf enum resources.documents.DOC_RELATION
 */
export enum DOC_RELATION {
    /**
     * @generated from protobuf enum value: MENTIONED = 0;
     */
    MENTIONED = 0,
    /**
     * @generated from protobuf enum value: TARGETS = 1;
     */
    TARGETS = 1,
    /**
     * @generated from protobuf enum value: CAUSED = 2;
     */
    CAUSED = 2
}
// @generated message type with reflection information, may provide speed optimized methods
class DocumentComment$Type extends MessageType<DocumentComment> {
    constructor() {
        super("resources.documents.DocumentComment", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "document_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 5, name: "comment", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3", maxBytes: "2048" } } } },
            { no: 6, name: "creator_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ },
            { no: 7, name: "creator", kind: "message", T: () => UserShort }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.DocumentComment
 */
export const DocumentComment = new DocumentComment$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Document$Type extends MessageType<Document> {
    constructor() {
        super("resources.documents.Document", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "category_id", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 5, name: "category", kind: "message", T: () => DocumentCategory },
            { no: 6, name: "title", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3", maxBytes: "21845" } } } },
            { no: 7, name: "content_type", kind: "enum", T: () => ["resources.documents.DOC_CONTENT_TYPE", DOC_CONTENT_TYPE], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 8, name: "content", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "20", maxBytes: "1000000" } } } },
            { no: 9, name: "data", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 10, name: "creator_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ },
            { no: 11, name: "creator", kind: "message", T: () => UserShort },
            { no: 12, name: "state", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "24" } } } },
            { no: 13, name: "closed", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 14, name: "public", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.Document
 */
export const Document = new Document$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DocumentShort$Type extends MessageType<DocumentShort> {
    constructor() {
        super("resources.documents.DocumentShort", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "category_id", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 5, name: "category", kind: "message", T: () => DocumentCategory },
            { no: 6, name: "title", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3" } } } },
            { no: 7, name: "content_type", kind: "enum", T: () => ["resources.documents.DOC_CONTENT_TYPE", DOC_CONTENT_TYPE], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 8, name: "content", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxBytes: "1024" } } } },
            { no: 9, name: "creator_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ },
            { no: 10, name: "creator", kind: "message", T: () => UserShort },
            { no: 11, name: "state", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "24" } } } },
            { no: 12, name: "closed", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.DocumentShort
 */
export const DocumentShort = new DocumentShort$Type();
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
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "document_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 5, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 6, name: "job_label", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 7, name: "minimumGrade", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 8, name: "job_grade_label", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 9, name: "access", kind: "enum", T: () => ["resources.documents.ACCESS_LEVEL", ACCESS_LEVEL], options: { "validate.rules": { enum: { definedOnly: true } } } }
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
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "document_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 5, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 6, name: "user", kind: "message", T: () => UserShort },
            { no: 7, name: "access", kind: "enum", T: () => ["resources.documents.ACCESS_LEVEL", ACCESS_LEVEL], options: { "validate.rules": { enum: { definedOnly: true } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.DocumentUserAccess
 */
export const DocumentUserAccess = new DocumentUserAccess$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DocumentReference$Type extends MessageType<DocumentReference> {
    constructor() {
        super("resources.documents.DocumentReference", [
            { no: 1, name: "id", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "source_document_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 4, name: "source_document", kind: "message", T: () => DocumentShort },
            { no: 5, name: "reference", kind: "enum", T: () => ["resources.documents.DOC_REFERENCE", DOC_REFERENCE], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 6, name: "target_document_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 7, name: "target_document", kind: "message", T: () => DocumentShort },
            { no: 8, name: "creator_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ },
            { no: 9, name: "creator", kind: "message", T: () => UserShort }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.DocumentReference
 */
export const DocumentReference = new DocumentReference$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DocumentRelation$Type extends MessageType<DocumentRelation> {
    constructor() {
        super("resources.documents.DocumentRelation", [
            { no: 1, name: "id", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "document_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 4, name: "document", kind: "message", T: () => DocumentShort },
            { no: 5, name: "source_user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 6, name: "source_user", kind: "message", T: () => UserShort },
            { no: 7, name: "relation", kind: "enum", T: () => ["resources.documents.DOC_RELATION", DOC_RELATION], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 8, name: "target_user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 9, name: "target_user", kind: "message", T: () => UserShort }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.DocumentRelation
 */
export const DocumentRelation = new DocumentRelation$Type();
