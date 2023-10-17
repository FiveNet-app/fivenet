// @generated by protobuf-ts 2.9.1 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "services/completor/completor.proto" (package "services.completor", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { MessageType } from "@protobuf-ts/runtime";
import { LawBook } from "../../resources/laws/laws.js";
import { Category } from "../../resources/documents/category.js";
import { Job } from "../../resources/users/jobs.js";
import { UserShort } from "../../resources/users/users.js";
/**
 * @generated from protobuf message services.completor.CompleteCitizensRequest
 */
export interface CompleteCitizensRequest {
    /**
     * @generated from protobuf field: string search = 1;
     */
    search: string;
    /**
     * @generated from protobuf field: optional bool current_job = 2;
     */
    currentJob?: boolean;
    /**
     * @generated from protobuf field: optional bool on_duty = 3;
     */
    onDuty?: boolean;
    /**
     * @generated from protobuf field: optional int32 user_id = 4;
     */
    userId?: number;
}
/**
 * @generated from protobuf message services.completor.CompleteCitizensRespoonse
 */
export interface CompleteCitizensRespoonse {
    /**
     * @generated from protobuf field: repeated resources.users.UserShort users = 1;
     */
    users: UserShort[]; // @gotags: alias:"user"
}
/**
 * @generated from protobuf message services.completor.CompleteJobsRequest
 */
export interface CompleteJobsRequest {
    /**
     * @generated from protobuf field: optional string search = 1;
     */
    search?: string;
    /**
     * @generated from protobuf field: optional bool exact_match = 2;
     */
    exactMatch?: boolean;
    /**
     * @generated from protobuf field: optional bool current_job = 3;
     */
    currentJob?: boolean;
}
/**
 * @generated from protobuf message services.completor.CompleteJobsResponse
 */
export interface CompleteJobsResponse {
    /**
     * @generated from protobuf field: repeated resources.users.Job jobs = 1;
     */
    jobs: Job[];
}
/**
 * @generated from protobuf message services.completor.CompleteDocumentCategoriesRequest
 */
export interface CompleteDocumentCategoriesRequest {
    /**
     * @generated from protobuf field: string search = 1;
     */
    search: string;
}
/**
 * @generated from protobuf message services.completor.CompleteDocumentCategoriesResponse
 */
export interface CompleteDocumentCategoriesResponse {
    /**
     * @generated from protobuf field: repeated resources.documents.Category categories = 1;
     */
    categories: Category[];
}
/**
 * @generated from protobuf message services.completor.ListLawBooksRequest
 */
export interface ListLawBooksRequest {
}
/**
 * @generated from protobuf message services.completor.ListLawBooksResponse
 */
export interface ListLawBooksResponse {
    /**
     * @generated from protobuf field: repeated resources.laws.LawBook books = 1;
     */
    books: LawBook[];
}
// @generated message type with reflection information, may provide speed optimized methods
class CompleteCitizensRequest$Type extends MessageType<CompleteCitizensRequest> {
    constructor() {
        super("services.completor.CompleteCitizensRequest", [
            { no: 1, name: "search", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 2, name: "current_job", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "on_duty", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "user_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.completor.CompleteCitizensRequest
 */
export const CompleteCitizensRequest = new CompleteCitizensRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CompleteCitizensRespoonse$Type extends MessageType<CompleteCitizensRespoonse> {
    constructor() {
        super("services.completor.CompleteCitizensRespoonse", [
            { no: 1, name: "users", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => UserShort }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.completor.CompleteCitizensRespoonse
 */
export const CompleteCitizensRespoonse = new CompleteCitizensRespoonse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CompleteJobsRequest$Type extends MessageType<CompleteJobsRequest> {
    constructor() {
        super("services.completor.CompleteJobsRequest", [
            { no: 1, name: "search", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 2, name: "exact_match", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "current_job", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.completor.CompleteJobsRequest
 */
export const CompleteJobsRequest = new CompleteJobsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CompleteJobsResponse$Type extends MessageType<CompleteJobsResponse> {
    constructor() {
        super("services.completor.CompleteJobsResponse", [
            { no: 1, name: "jobs", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Job }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.completor.CompleteJobsResponse
 */
export const CompleteJobsResponse = new CompleteJobsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CompleteDocumentCategoriesRequest$Type extends MessageType<CompleteDocumentCategoriesRequest> {
    constructor() {
        super("services.completor.CompleteDocumentCategoriesRequest", [
            { no: 1, name: "search", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "128" } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.completor.CompleteDocumentCategoriesRequest
 */
export const CompleteDocumentCategoriesRequest = new CompleteDocumentCategoriesRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CompleteDocumentCategoriesResponse$Type extends MessageType<CompleteDocumentCategoriesResponse> {
    constructor() {
        super("services.completor.CompleteDocumentCategoriesResponse", [
            { no: 1, name: "categories", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Category }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.completor.CompleteDocumentCategoriesResponse
 */
export const CompleteDocumentCategoriesResponse = new CompleteDocumentCategoriesResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListLawBooksRequest$Type extends MessageType<ListLawBooksRequest> {
    constructor() {
        super("services.completor.ListLawBooksRequest", []);
    }
}
/**
 * @generated MessageType for protobuf message services.completor.ListLawBooksRequest
 */
export const ListLawBooksRequest = new ListLawBooksRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListLawBooksResponse$Type extends MessageType<ListLawBooksResponse> {
    constructor() {
        super("services.completor.ListLawBooksResponse", [
            { no: 1, name: "books", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => LawBook }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.completor.ListLawBooksResponse
 */
export const ListLawBooksResponse = new ListLawBooksResponse$Type();
/**
 * @generated ServiceType for protobuf service services.completor.CompletorService
 */
export const CompletorService = new ServiceType("services.completor.CompletorService", [
    { name: "CompleteCitizens", options: {}, I: CompleteCitizensRequest, O: CompleteCitizensRespoonse },
    { name: "CompleteJobs", options: {}, I: CompleteJobsRequest, O: CompleteJobsResponse },
    { name: "CompleteDocumentCategories", options: {}, I: CompleteDocumentCategoriesRequest, O: CompleteDocumentCategoriesResponse },
    { name: "ListLawBooks", options: {}, I: ListLawBooksRequest, O: ListLawBooksResponse }
]);