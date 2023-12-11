// @generated by protobuf-ts 2.9.3 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "resources/documents/documents.proto" (package "resources.documents", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { UserShort } from "../users/users";
import { Category } from "./category";
import { Timestamp } from "../timestamp/timestamp";
/**
 * @generated from protobuf message resources.documents.Document
 */
export interface Document {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp deleted_at = 4;
     */
    deletedAt?: Timestamp;
    /**
     * @generated from protobuf field: optional uint64 category_id = 5 [jstype = JS_STRING];
     */
    categoryId?: string; // @gotags: alias:"category_id"
    /**
     * @generated from protobuf field: optional resources.documents.Category category = 6;
     */
    category?: Category; // @gotags: alias:"category"
    /**
     * @sanitize
     *
     * @generated from protobuf field: string title = 7;
     */
    title: string; // @gotags: alias:"title"
    /**
     * @generated from protobuf field: resources.documents.DocContentType content_type = 8;
     */
    contentType: DocContentType; // @gotags: alias:"content_type"
    /**
     * @sanitize
     *
     * @generated from protobuf field: string content = 9;
     */
    content: string; // @gotags: alias:"content"
    /**
     * @sanitize
     *
     * @generated from protobuf field: optional string data = 10;
     */
    data?: string; // @gotags: alias:"data"
    /**
     * @generated from protobuf field: optional int32 creator_id = 11;
     */
    creatorId?: number; // @gotags: alias:"creator_id"
    /**
     * @generated from protobuf field: optional resources.users.UserShort creator = 12;
     */
    creator?: UserShort; // @gotags: alias:"creator"
    /**
     * @generated from protobuf field: string creator_job = 13;
     */
    creatorJob: string; // @gotags: alias:"creator_job"
    /**
     * @sanitize
     *
     * @generated from protobuf field: string state = 14;
     */
    state: string; // @gotags: alias:"state"
    /**
     * @generated from protobuf field: bool closed = 15;
     */
    closed: boolean;
    /**
     * @generated from protobuf field: bool public = 16;
     */
    public: boolean;
}
/**
 * @generated from protobuf message resources.documents.DocumentShort
 */
export interface DocumentShort {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp deleted_at = 4;
     */
    deletedAt?: Timestamp;
    /**
     * @generated from protobuf field: optional uint64 category_id = 5 [jstype = JS_STRING];
     */
    categoryId?: string; // @gotags: alias:"category_id"
    /**
     * @generated from protobuf field: optional resources.documents.Category category = 6;
     */
    category?: Category; // @gotags: alias:"category"
    /**
     * @sanitize
     *
     * @generated from protobuf field: string title = 7;
     */
    title: string; // @gotags: alias:"title"
    /**
     * @generated from protobuf field: resources.documents.DocContentType content_type = 8;
     */
    contentType: DocContentType; // @gotags: alias:"content_type"
    /**
     * @sanitize
     *
     * @generated from protobuf field: string content = 9;
     */
    content: string; // @gotags: alias:"content"
    /**
     * @generated from protobuf field: optional int32 creator_id = 10;
     */
    creatorId?: number; // @gotags: alias:"creator_id"
    /**
     * @generated from protobuf field: optional resources.users.UserShort creator = 11;
     */
    creator?: UserShort; // @gotags: alias:"creator"
    /**
     * @generated from protobuf field: string creator_job = 12;
     */
    creatorJob: string; // @gotags: alias:"creator_job"
    /**
     * @sanitize
     *
     * @generated from protobuf field: string state = 13;
     */
    state: string; // @gotags: alias:"state"
    /**
     * @generated from protobuf field: bool closed = 14;
     */
    closed: boolean;
    /**
     * @generated from protobuf field: bool public = 15;
     */
    public: boolean;
}
/**
 * @generated from protobuf message resources.documents.DocumentReference
 */
export interface DocumentReference {
    /**
     * @generated from protobuf field: optional uint64 id = 1 [jstype = JS_STRING];
     */
    id?: string;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: uint64 source_document_id = 3 [jstype = JS_STRING];
     */
    sourceDocumentId: string; // @gotags: alias:"source_document_id"
    /**
     * @generated from protobuf field: optional resources.documents.DocumentShort source_document = 4;
     */
    sourceDocument?: DocumentShort; // @gotags: alias:"source_document"
    /**
     * @generated from protobuf field: resources.documents.DocReference reference = 5;
     */
    reference: DocReference; // @gotags: alias:"reference"
    /**
     * @generated from protobuf field: uint64 target_document_id = 6 [jstype = JS_STRING];
     */
    targetDocumentId: string; // @gotags: alias:"target_document_id"
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
     * @generated from protobuf field: optional uint64 id = 1 [jstype = JS_STRING];
     */
    id?: string;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: uint64 document_id = 3 [jstype = JS_STRING];
     */
    documentId: string;
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
     * @generated from protobuf field: resources.documents.DocRelation relation = 7;
     */
    relation: DocRelation; // @gotags: alias:"relation"
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
 * @generated from protobuf enum resources.documents.DocContentType
 */
export enum DocContentType {
    /**
     * @generated from protobuf enum value: DOC_CONTENT_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: DOC_CONTENT_TYPE_HTML = 1;
     */
    HTML = 1,
    /**
     * @generated from protobuf enum value: DOC_CONTENT_TYPE_PLAIN = 2;
     */
    PLAIN = 2
}
/**
 * @generated from protobuf enum resources.documents.DocReference
 */
export enum DocReference {
    /**
     * @generated from protobuf enum value: DOC_REFERENCE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: DOC_REFERENCE_LINKED = 1;
     */
    LINKED = 1,
    /**
     * @generated from protobuf enum value: DOC_REFERENCE_SOLVES = 2;
     */
    SOLVES = 2,
    /**
     * @generated from protobuf enum value: DOC_REFERENCE_CLOSES = 3;
     */
    CLOSES = 3,
    /**
     * @generated from protobuf enum value: DOC_REFERENCE_DEPRECATES = 4;
     */
    DEPRECATES = 4
}
/**
 * @generated from protobuf enum resources.documents.DocRelation
 */
export enum DocRelation {
    /**
     * @generated from protobuf enum value: DOC_RELATION_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: DOC_RELATION_MENTIONED = 1;
     */
    MENTIONED = 1,
    /**
     * @generated from protobuf enum value: DOC_RELATION_TARGETS = 2;
     */
    TARGETS = 2,
    /**
     * @generated from protobuf enum value: DOC_RELATION_CAUSED = 3;
     */
    CAUSED = 3
}
// @generated message type with reflection information, may provide speed optimized methods
class Document$Type extends MessageType<Document> {
    constructor() {
        super("resources.documents.Document", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "deleted_at", kind: "message", T: () => Timestamp },
            { no: 5, name: "category_id", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/ },
            { no: 6, name: "category", kind: "message", T: () => Category },
            { no: 7, name: "title", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3", maxBytes: "21845" } } } },
            { no: 8, name: "content_type", kind: "enum", T: () => ["resources.documents.DocContentType", DocContentType, "DOC_CONTENT_TYPE_"], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 9, name: "content", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "20", maxBytes: "1750000" } } } },
            { no: 10, name: "data", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxBytes: "1000000" } } } },
            { no: 11, name: "creator_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ },
            { no: 12, name: "creator", kind: "message", T: () => UserShort },
            { no: 13, name: "creator_job", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 14, name: "state", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "24" } } } },
            { no: 15, name: "closed", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 16, name: "public", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
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
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "deleted_at", kind: "message", T: () => Timestamp },
            { no: 5, name: "category_id", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/ },
            { no: 6, name: "category", kind: "message", T: () => Category },
            { no: 7, name: "title", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3" } } } },
            { no: 8, name: "content_type", kind: "enum", T: () => ["resources.documents.DocContentType", DocContentType, "DOC_CONTENT_TYPE_"], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 9, name: "content", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxBytes: "1024" } } } },
            { no: 10, name: "creator_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ },
            { no: 11, name: "creator", kind: "message", T: () => UserShort },
            { no: 12, name: "creator_job", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 13, name: "state", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "24" } } } },
            { no: 14, name: "closed", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 15, name: "public", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.DocumentShort
 */
export const DocumentShort = new DocumentShort$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DocumentReference$Type extends MessageType<DocumentReference> {
    constructor() {
        super("resources.documents.DocumentReference", [
            { no: 1, name: "id", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "source_document_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 4, name: "source_document", kind: "message", T: () => DocumentShort },
            { no: 5, name: "reference", kind: "enum", T: () => ["resources.documents.DocReference", DocReference, "DOC_REFERENCE_"], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 6, name: "target_document_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
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
            { no: 1, name: "id", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "document_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 4, name: "document", kind: "message", T: () => DocumentShort },
            { no: 5, name: "source_user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 6, name: "source_user", kind: "message", T: () => UserShort },
            { no: 7, name: "relation", kind: "enum", T: () => ["resources.documents.DocRelation", DocRelation, "DOC_RELATION_"], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 8, name: "target_user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 9, name: "target_user", kind: "message", T: () => UserShort }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.DocumentRelation
 */
export const DocumentRelation = new DocumentRelation$Type();
