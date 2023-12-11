// @generated by protobuf-ts 2.9.3 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "services/jobs/jobs.proto" (package "services.jobs", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { MessageType } from "@protobuf-ts/runtime";
import { RequestComment } from "../../resources/jobs/requests";
import { RequestType } from "../../resources/jobs/requests";
import { Request } from "../../resources/jobs/requests";
import { TimeclockStats } from "../../resources/jobs/timeclock";
import { TimeclockEntry } from "../../resources/jobs/timeclock";
import { Timestamp } from "../../resources/timestamp/timestamp";
import { ConductEntry } from "../../resources/jobs/conduct";
import { ConductType } from "../../resources/jobs/conduct";
import { User } from "../../resources/users/users";
import { PaginationResponse } from "../../resources/common/database/database";
import { PaginationRequest } from "../../resources/common/database/database";
// Colleagues

/**
 * @generated from protobuf message services.jobs.ColleaguesListRequest
 */
export interface ColleaguesListRequest {
    /**
     * @generated from protobuf field: resources.common.database.PaginationRequest pagination = 1;
     */
    pagination?: PaginationRequest;
    /**
     * Search params
     *
     * @generated from protobuf field: string search_name = 2;
     */
    searchName: string;
}
/**
 * @generated from protobuf message services.jobs.ColleaguesListResponse
 */
export interface ColleaguesListResponse {
    /**
     * @generated from protobuf field: resources.common.database.PaginationResponse pagination = 1;
     */
    pagination?: PaginationResponse;
    /**
     * @generated from protobuf field: repeated resources.users.User users = 2;
     */
    users: User[];
}
// Conduct Register

/**
 * @generated from protobuf message services.jobs.ConductListEntriesRequest
 */
export interface ConductListEntriesRequest {
    /**
     * @generated from protobuf field: resources.common.database.PaginationRequest pagination = 1;
     */
    pagination?: PaginationRequest;
    /**
     * Search params
     *
     * @generated from protobuf field: repeated resources.jobs.ConductType types = 2;
     */
    types: ConductType[];
    /**
     * @generated from protobuf field: optional bool show_expired = 3;
     */
    showExpired?: boolean;
    /**
     * @generated from protobuf field: repeated int32 user_ids = 4;
     */
    userIds: number[];
}
/**
 * @generated from protobuf message services.jobs.ConductListEntriesResponse
 */
export interface ConductListEntriesResponse {
    /**
     * @generated from protobuf field: resources.common.database.PaginationResponse pagination = 1;
     */
    pagination?: PaginationResponse;
    /**
     * @generated from protobuf field: repeated resources.jobs.ConductEntry entries = 2;
     */
    entries: ConductEntry[];
}
/**
 * @generated from protobuf message services.jobs.ConductCreateEntryRequest
 */
export interface ConductCreateEntryRequest {
    /**
     * @generated from protobuf field: resources.jobs.ConductEntry entry = 1;
     */
    entry?: ConductEntry;
}
/**
 * @generated from protobuf message services.jobs.ConductCreateEntryResponse
 */
export interface ConductCreateEntryResponse {
    /**
     * @generated from protobuf field: resources.jobs.ConductEntry entry = 1;
     */
    entry?: ConductEntry;
}
/**
 * @generated from protobuf message services.jobs.ConductUpdateEntryRequest
 */
export interface ConductUpdateEntryRequest {
    /**
     * @generated from protobuf field: resources.jobs.ConductEntry entry = 1;
     */
    entry?: ConductEntry;
}
/**
 * @generated from protobuf message services.jobs.ConductUpdateEntryResponse
 */
export interface ConductUpdateEntryResponse {
    /**
     * @generated from protobuf field: resources.jobs.ConductEntry entry = 1;
     */
    entry?: ConductEntry;
}
/**
 * @generated from protobuf message services.jobs.ConductDeleteEntryRequest
 */
export interface ConductDeleteEntryRequest {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
}
/**
 * @generated from protobuf message services.jobs.ConductDeleteEntryResponse
 */
export interface ConductDeleteEntryResponse {
}
// Time Clock

/**
 * @generated from protobuf message services.jobs.TimeclockListEntriesRequest
 */
export interface TimeclockListEntriesRequest {
    /**
     * @generated from protobuf field: resources.common.database.PaginationRequest pagination = 1;
     */
    pagination?: PaginationRequest;
    /**
     * Search
     *
     * @generated from protobuf field: repeated int32 user_ids = 2;
     */
    userIds: number[];
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp from = 3;
     */
    from?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp to = 4;
     */
    to?: Timestamp;
    /**
     * @generated from protobuf field: optional bool per_day = 5;
     */
    perDay?: boolean;
}
/**
 * @generated from protobuf message services.jobs.TimeclockListEntriesResponse
 */
export interface TimeclockListEntriesResponse {
    /**
     * @generated from protobuf field: resources.common.database.PaginationResponse pagination = 1;
     */
    pagination?: PaginationResponse;
    /**
     * @generated from protobuf field: repeated resources.jobs.TimeclockEntry entries = 2;
     */
    entries: TimeclockEntry[];
    /**
     * @generated from protobuf field: resources.jobs.TimeclockStats stats = 3;
     */
    stats?: TimeclockStats;
}
/**
 * @generated from protobuf message services.jobs.TimeclockStatsRequest
 */
export interface TimeclockStatsRequest {
}
/**
 * @generated from protobuf message services.jobs.TimeclockStatsResponse
 */
export interface TimeclockStatsResponse {
    /**
     * @generated from protobuf field: resources.jobs.TimeclockStats stats = 1;
     */
    stats?: TimeclockStats;
}
// Requests

/**
 * @generated from protobuf message services.jobs.RequestsListEntriesRequest
 */
export interface RequestsListEntriesRequest {
    /**
     * @generated from protobuf field: resources.common.database.PaginationRequest pagination = 1;
     */
    pagination?: PaginationRequest;
    /**
     * Search
     *
     * @generated from protobuf field: repeated int32 user_ids = 2;
     */
    userIds: number[];
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp from = 3;
     */
    from?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp to = 4;
     */
    to?: Timestamp;
    /**
     * @generated from protobuf field: optional string search = 5;
     */
    search?: string;
}
/**
 * @generated from protobuf message services.jobs.RequestsListEntriesResponse
 */
export interface RequestsListEntriesResponse {
    /**
     * @generated from protobuf field: resources.common.database.PaginationResponse pagination = 1;
     */
    pagination?: PaginationResponse;
    /**
     * @generated from protobuf field: repeated resources.jobs.Request entries = 2;
     */
    entries: Request[];
}
/**
 * @generated from protobuf message services.jobs.RequestsCreateEntryRequest
 */
export interface RequestsCreateEntryRequest {
    /**
     * @generated from protobuf field: resources.jobs.Request entry = 1;
     */
    entry?: Request;
}
/**
 * @generated from protobuf message services.jobs.RequestsCreateEntryResponse
 */
export interface RequestsCreateEntryResponse {
    /**
     * @generated from protobuf field: resources.jobs.Request entry = 1;
     */
    entry?: Request;
}
/**
 * @generated from protobuf message services.jobs.RequestsUpdateEntryRequest
 */
export interface RequestsUpdateEntryRequest {
    /**
     * @generated from protobuf field: resources.jobs.Request entry = 1;
     */
    entry?: Request;
}
/**
 * @generated from protobuf message services.jobs.RequestsUpdateEntryResponse
 */
export interface RequestsUpdateEntryResponse {
    /**
     * @generated from protobuf field: resources.jobs.Request entry = 1;
     */
    entry?: Request;
}
/**
 * @generated from protobuf message services.jobs.RequestsDeleteEntryRequest
 */
export interface RequestsDeleteEntryRequest {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
}
/**
 * @generated from protobuf message services.jobs.RequestsDeleteEntryResponse
 */
export interface RequestsDeleteEntryResponse {
}
/**
 * @generated from protobuf message services.jobs.RequestsApproveEntryRequest
 */
export interface RequestsApproveEntryRequest {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
    /**
     * @generated from protobuf field: bool approve = 2;
     */
    approve: boolean;
}
/**
 * @generated from protobuf message services.jobs.RequestsApproveEntryResponse
 */
export interface RequestsApproveEntryResponse {
}
/**
 * @generated from protobuf message services.jobs.RequestsCloseEntryRequest
 */
export interface RequestsCloseEntryRequest {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
    /**
     * @generated from protobuf field: bool close = 2;
     */
    close: boolean;
}
/**
 * @generated from protobuf message services.jobs.RequestsCloseEntryResponse
 */
export interface RequestsCloseEntryResponse {
}
/**
 * @generated from protobuf message services.jobs.RequestsListTypesRequest
 */
export interface RequestsListTypesRequest {
}
/**
 * @generated from protobuf message services.jobs.RequestsListTypesResponse
 */
export interface RequestsListTypesResponse {
    /**
     * @generated from protobuf field: repeated resources.jobs.RequestType types = 1;
     */
    types: RequestType[];
}
/**
 * @generated from protobuf message services.jobs.RequestsCreateOrUpdateTypeRequest
 */
export interface RequestsCreateOrUpdateTypeRequest {
    /**
     * @generated from protobuf field: resources.jobs.RequestType request_type = 1;
     */
    requestType?: RequestType;
}
/**
 * @generated from protobuf message services.jobs.RequestsCreateOrUpdateTypeResponse
 */
export interface RequestsCreateOrUpdateTypeResponse {
    /**
     * @generated from protobuf field: resources.jobs.RequestType request_type = 1;
     */
    requestType?: RequestType;
}
/**
 * @generated from protobuf message services.jobs.RequestsDeleteTypeRequest
 */
export interface RequestsDeleteTypeRequest {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
}
/**
 * @generated from protobuf message services.jobs.RequestsDeleteTypeResponse
 */
export interface RequestsDeleteTypeResponse {
}
/**
 * @generated from protobuf message services.jobs.RequestsListCommentsRequest
 */
export interface RequestsListCommentsRequest {
    /**
     * @generated from protobuf field: resources.common.database.PaginationRequest pagination = 1;
     */
    pagination?: PaginationRequest;
    /**
     * @generated from protobuf field: uint64 request_id = 2 [jstype = JS_STRING];
     */
    requestId: string;
}
/**
 * @generated from protobuf message services.jobs.RequestsListCommentsResponse
 */
export interface RequestsListCommentsResponse {
    /**
     * @generated from protobuf field: resources.common.database.PaginationResponse pagination = 1;
     */
    pagination?: PaginationResponse;
    /**
     * @generated from protobuf field: repeated resources.jobs.RequestComment comments = 2;
     */
    comments: RequestComment[];
}
/**
 * @generated from protobuf message services.jobs.RequestsPostCommentRequest
 */
export interface RequestsPostCommentRequest {
    /**
     * @generated from protobuf field: resources.jobs.RequestComment comment = 1;
     */
    comment?: RequestComment;
}
/**
 * @generated from protobuf message services.jobs.RequestsPostCommentResponse
 */
export interface RequestsPostCommentResponse {
    /**
     * @generated from protobuf field: resources.jobs.RequestComment comment = 1;
     */
    comment?: RequestComment;
}
/**
 * @generated from protobuf message services.jobs.RequestsDeleteCommentRequest
 */
export interface RequestsDeleteCommentRequest {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
}
/**
 * @generated from protobuf message services.jobs.RequestsDeleteCommentResponse
 */
export interface RequestsDeleteCommentResponse {
}
// @generated message type with reflection information, may provide speed optimized methods
class ColleaguesListRequest$Type extends MessageType<ColleaguesListRequest> {
    constructor() {
        super("services.jobs.ColleaguesListRequest", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationRequest, options: { "validate.rules": { message: { required: true } } } },
            { no: 2, name: "search_name", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.ColleaguesListRequest
 */
export const ColleaguesListRequest = new ColleaguesListRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ColleaguesListResponse$Type extends MessageType<ColleaguesListResponse> {
    constructor() {
        super("services.jobs.ColleaguesListResponse", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationResponse },
            { no: 2, name: "users", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => User }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.ColleaguesListResponse
 */
export const ColleaguesListResponse = new ColleaguesListResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ConductListEntriesRequest$Type extends MessageType<ConductListEntriesRequest> {
    constructor() {
        super("services.jobs.ConductListEntriesRequest", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationRequest, options: { "validate.rules": { message: { required: true } } } },
            { no: 2, name: "types", kind: "enum", repeat: 1 /*RepeatType.PACKED*/, T: () => ["resources.jobs.ConductType", ConductType, "CONDUCT_TYPE_"] },
            { no: 3, name: "show_expired", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "user_ids", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 5 /*ScalarType.INT32*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.ConductListEntriesRequest
 */
export const ConductListEntriesRequest = new ConductListEntriesRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ConductListEntriesResponse$Type extends MessageType<ConductListEntriesResponse> {
    constructor() {
        super("services.jobs.ConductListEntriesResponse", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationResponse },
            { no: 2, name: "entries", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => ConductEntry }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.ConductListEntriesResponse
 */
export const ConductListEntriesResponse = new ConductListEntriesResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ConductCreateEntryRequest$Type extends MessageType<ConductCreateEntryRequest> {
    constructor() {
        super("services.jobs.ConductCreateEntryRequest", [
            { no: 1, name: "entry", kind: "message", T: () => ConductEntry, options: { "validate.rules": { message: { required: true } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.ConductCreateEntryRequest
 */
export const ConductCreateEntryRequest = new ConductCreateEntryRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ConductCreateEntryResponse$Type extends MessageType<ConductCreateEntryResponse> {
    constructor() {
        super("services.jobs.ConductCreateEntryResponse", [
            { no: 1, name: "entry", kind: "message", T: () => ConductEntry }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.ConductCreateEntryResponse
 */
export const ConductCreateEntryResponse = new ConductCreateEntryResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ConductUpdateEntryRequest$Type extends MessageType<ConductUpdateEntryRequest> {
    constructor() {
        super("services.jobs.ConductUpdateEntryRequest", [
            { no: 1, name: "entry", kind: "message", T: () => ConductEntry, options: { "validate.rules": { message: { required: true } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.ConductUpdateEntryRequest
 */
export const ConductUpdateEntryRequest = new ConductUpdateEntryRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ConductUpdateEntryResponse$Type extends MessageType<ConductUpdateEntryResponse> {
    constructor() {
        super("services.jobs.ConductUpdateEntryResponse", [
            { no: 1, name: "entry", kind: "message", T: () => ConductEntry, options: { "validate.rules": { message: { required: true } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.ConductUpdateEntryResponse
 */
export const ConductUpdateEntryResponse = new ConductUpdateEntryResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ConductDeleteEntryRequest$Type extends MessageType<ConductDeleteEntryRequest> {
    constructor() {
        super("services.jobs.ConductDeleteEntryRequest", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.ConductDeleteEntryRequest
 */
export const ConductDeleteEntryRequest = new ConductDeleteEntryRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ConductDeleteEntryResponse$Type extends MessageType<ConductDeleteEntryResponse> {
    constructor() {
        super("services.jobs.ConductDeleteEntryResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.ConductDeleteEntryResponse
 */
export const ConductDeleteEntryResponse = new ConductDeleteEntryResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TimeclockListEntriesRequest$Type extends MessageType<TimeclockListEntriesRequest> {
    constructor() {
        super("services.jobs.TimeclockListEntriesRequest", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationRequest, options: { "validate.rules": { message: { required: true } } } },
            { no: 2, name: "user_ids", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "from", kind: "message", T: () => Timestamp },
            { no: 4, name: "to", kind: "message", T: () => Timestamp },
            { no: 5, name: "per_day", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.TimeclockListEntriesRequest
 */
export const TimeclockListEntriesRequest = new TimeclockListEntriesRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TimeclockListEntriesResponse$Type extends MessageType<TimeclockListEntriesResponse> {
    constructor() {
        super("services.jobs.TimeclockListEntriesResponse", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationResponse },
            { no: 2, name: "entries", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => TimeclockEntry },
            { no: 3, name: "stats", kind: "message", T: () => TimeclockStats }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.TimeclockListEntriesResponse
 */
export const TimeclockListEntriesResponse = new TimeclockListEntriesResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TimeclockStatsRequest$Type extends MessageType<TimeclockStatsRequest> {
    constructor() {
        super("services.jobs.TimeclockStatsRequest", []);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.TimeclockStatsRequest
 */
export const TimeclockStatsRequest = new TimeclockStatsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TimeclockStatsResponse$Type extends MessageType<TimeclockStatsResponse> {
    constructor() {
        super("services.jobs.TimeclockStatsResponse", [
            { no: 1, name: "stats", kind: "message", T: () => TimeclockStats }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.TimeclockStatsResponse
 */
export const TimeclockStatsResponse = new TimeclockStatsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsListEntriesRequest$Type extends MessageType<RequestsListEntriesRequest> {
    constructor() {
        super("services.jobs.RequestsListEntriesRequest", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationRequest, options: { "validate.rules": { message: { required: true } } } },
            { no: 2, name: "user_ids", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "from", kind: "message", T: () => Timestamp },
            { no: 4, name: "to", kind: "message", T: () => Timestamp },
            { no: 5, name: "search", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsListEntriesRequest
 */
export const RequestsListEntriesRequest = new RequestsListEntriesRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsListEntriesResponse$Type extends MessageType<RequestsListEntriesResponse> {
    constructor() {
        super("services.jobs.RequestsListEntriesResponse", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationResponse },
            { no: 2, name: "entries", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Request }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsListEntriesResponse
 */
export const RequestsListEntriesResponse = new RequestsListEntriesResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsCreateEntryRequest$Type extends MessageType<RequestsCreateEntryRequest> {
    constructor() {
        super("services.jobs.RequestsCreateEntryRequest", [
            { no: 1, name: "entry", kind: "message", T: () => Request, options: { "validate.rules": { message: { required: true } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsCreateEntryRequest
 */
export const RequestsCreateEntryRequest = new RequestsCreateEntryRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsCreateEntryResponse$Type extends MessageType<RequestsCreateEntryResponse> {
    constructor() {
        super("services.jobs.RequestsCreateEntryResponse", [
            { no: 1, name: "entry", kind: "message", T: () => Request }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsCreateEntryResponse
 */
export const RequestsCreateEntryResponse = new RequestsCreateEntryResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsUpdateEntryRequest$Type extends MessageType<RequestsUpdateEntryRequest> {
    constructor() {
        super("services.jobs.RequestsUpdateEntryRequest", [
            { no: 1, name: "entry", kind: "message", T: () => Request, options: { "validate.rules": { message: { required: true } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsUpdateEntryRequest
 */
export const RequestsUpdateEntryRequest = new RequestsUpdateEntryRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsUpdateEntryResponse$Type extends MessageType<RequestsUpdateEntryResponse> {
    constructor() {
        super("services.jobs.RequestsUpdateEntryResponse", [
            { no: 1, name: "entry", kind: "message", T: () => Request }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsUpdateEntryResponse
 */
export const RequestsUpdateEntryResponse = new RequestsUpdateEntryResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsDeleteEntryRequest$Type extends MessageType<RequestsDeleteEntryRequest> {
    constructor() {
        super("services.jobs.RequestsDeleteEntryRequest", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsDeleteEntryRequest
 */
export const RequestsDeleteEntryRequest = new RequestsDeleteEntryRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsDeleteEntryResponse$Type extends MessageType<RequestsDeleteEntryResponse> {
    constructor() {
        super("services.jobs.RequestsDeleteEntryResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsDeleteEntryResponse
 */
export const RequestsDeleteEntryResponse = new RequestsDeleteEntryResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsApproveEntryRequest$Type extends MessageType<RequestsApproveEntryRequest> {
    constructor() {
        super("services.jobs.RequestsApproveEntryRequest", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "approve", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsApproveEntryRequest
 */
export const RequestsApproveEntryRequest = new RequestsApproveEntryRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsApproveEntryResponse$Type extends MessageType<RequestsApproveEntryResponse> {
    constructor() {
        super("services.jobs.RequestsApproveEntryResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsApproveEntryResponse
 */
export const RequestsApproveEntryResponse = new RequestsApproveEntryResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsCloseEntryRequest$Type extends MessageType<RequestsCloseEntryRequest> {
    constructor() {
        super("services.jobs.RequestsCloseEntryRequest", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "close", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsCloseEntryRequest
 */
export const RequestsCloseEntryRequest = new RequestsCloseEntryRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsCloseEntryResponse$Type extends MessageType<RequestsCloseEntryResponse> {
    constructor() {
        super("services.jobs.RequestsCloseEntryResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsCloseEntryResponse
 */
export const RequestsCloseEntryResponse = new RequestsCloseEntryResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsListTypesRequest$Type extends MessageType<RequestsListTypesRequest> {
    constructor() {
        super("services.jobs.RequestsListTypesRequest", []);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsListTypesRequest
 */
export const RequestsListTypesRequest = new RequestsListTypesRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsListTypesResponse$Type extends MessageType<RequestsListTypesResponse> {
    constructor() {
        super("services.jobs.RequestsListTypesResponse", [
            { no: 1, name: "types", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => RequestType }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsListTypesResponse
 */
export const RequestsListTypesResponse = new RequestsListTypesResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsCreateOrUpdateTypeRequest$Type extends MessageType<RequestsCreateOrUpdateTypeRequest> {
    constructor() {
        super("services.jobs.RequestsCreateOrUpdateTypeRequest", [
            { no: 1, name: "request_type", kind: "message", T: () => RequestType }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsCreateOrUpdateTypeRequest
 */
export const RequestsCreateOrUpdateTypeRequest = new RequestsCreateOrUpdateTypeRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsCreateOrUpdateTypeResponse$Type extends MessageType<RequestsCreateOrUpdateTypeResponse> {
    constructor() {
        super("services.jobs.RequestsCreateOrUpdateTypeResponse", [
            { no: 1, name: "request_type", kind: "message", T: () => RequestType }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsCreateOrUpdateTypeResponse
 */
export const RequestsCreateOrUpdateTypeResponse = new RequestsCreateOrUpdateTypeResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsDeleteTypeRequest$Type extends MessageType<RequestsDeleteTypeRequest> {
    constructor() {
        super("services.jobs.RequestsDeleteTypeRequest", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsDeleteTypeRequest
 */
export const RequestsDeleteTypeRequest = new RequestsDeleteTypeRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsDeleteTypeResponse$Type extends MessageType<RequestsDeleteTypeResponse> {
    constructor() {
        super("services.jobs.RequestsDeleteTypeResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsDeleteTypeResponse
 */
export const RequestsDeleteTypeResponse = new RequestsDeleteTypeResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsListCommentsRequest$Type extends MessageType<RequestsListCommentsRequest> {
    constructor() {
        super("services.jobs.RequestsListCommentsRequest", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationRequest, options: { "validate.rules": { message: { required: true } } } },
            { no: 2, name: "request_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsListCommentsRequest
 */
export const RequestsListCommentsRequest = new RequestsListCommentsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsListCommentsResponse$Type extends MessageType<RequestsListCommentsResponse> {
    constructor() {
        super("services.jobs.RequestsListCommentsResponse", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationResponse },
            { no: 2, name: "comments", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => RequestComment }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsListCommentsResponse
 */
export const RequestsListCommentsResponse = new RequestsListCommentsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsPostCommentRequest$Type extends MessageType<RequestsPostCommentRequest> {
    constructor() {
        super("services.jobs.RequestsPostCommentRequest", [
            { no: 1, name: "comment", kind: "message", T: () => RequestComment, options: { "validate.rules": { message: { required: true } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsPostCommentRequest
 */
export const RequestsPostCommentRequest = new RequestsPostCommentRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsPostCommentResponse$Type extends MessageType<RequestsPostCommentResponse> {
    constructor() {
        super("services.jobs.RequestsPostCommentResponse", [
            { no: 1, name: "comment", kind: "message", T: () => RequestComment, options: { "validate.rules": { message: { required: true } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsPostCommentResponse
 */
export const RequestsPostCommentResponse = new RequestsPostCommentResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsDeleteCommentRequest$Type extends MessageType<RequestsDeleteCommentRequest> {
    constructor() {
        super("services.jobs.RequestsDeleteCommentRequest", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsDeleteCommentRequest
 */
export const RequestsDeleteCommentRequest = new RequestsDeleteCommentRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RequestsDeleteCommentResponse$Type extends MessageType<RequestsDeleteCommentResponse> {
    constructor() {
        super("services.jobs.RequestsDeleteCommentResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.RequestsDeleteCommentResponse
 */
export const RequestsDeleteCommentResponse = new RequestsDeleteCommentResponse$Type();
/**
 * @generated ServiceType for protobuf service services.jobs.JobsService
 */
export const JobsService = new ServiceType("services.jobs.JobsService", [
    { name: "ColleaguesList", options: {}, I: ColleaguesListRequest, O: ColleaguesListResponse },
    { name: "ConductListEntries", options: {}, I: ConductListEntriesRequest, O: ConductListEntriesResponse },
    { name: "ConductCreateEntry", options: {}, I: ConductCreateEntryRequest, O: ConductCreateEntryResponse },
    { name: "ConductUpdateEntry", options: {}, I: ConductUpdateEntryRequest, O: ConductUpdateEntryResponse },
    { name: "ConductDeleteEntry", options: {}, I: ConductDeleteEntryRequest, O: ConductDeleteEntryResponse },
    { name: "TimeclockListEntries", options: {}, I: TimeclockListEntriesRequest, O: TimeclockListEntriesResponse },
    { name: "TimeclockStats", options: {}, I: TimeclockStatsRequest, O: TimeclockStatsResponse },
    { name: "RequestsListEntries", options: {}, I: RequestsListEntriesRequest, O: RequestsListEntriesResponse },
    { name: "RequestsCreateEntry", options: {}, I: RequestsCreateEntryRequest, O: RequestsCreateEntryResponse },
    { name: "RequestsUpdateEntry", options: {}, I: RequestsUpdateEntryRequest, O: RequestsUpdateEntryResponse },
    { name: "RequestsDeleteEntry", options: {}, I: RequestsDeleteEntryRequest, O: RequestsDeleteEntryResponse },
    { name: "RequestsListTypes", options: {}, I: RequestsListTypesRequest, O: RequestsListTypesResponse },
    { name: "RequestsCreateOrUpdateType", options: {}, I: RequestsCreateOrUpdateTypeRequest, O: RequestsCreateOrUpdateTypeResponse },
    { name: "RequestsDeleteType", options: {}, I: RequestsDeleteTypeRequest, O: RequestsDeleteTypeResponse },
    { name: "RequestsListComments", options: {}, I: RequestsListCommentsRequest, O: RequestsListCommentsResponse },
    { name: "RequestsPostComment", options: {}, I: RequestsPostCommentRequest, O: RequestsPostCommentResponse },
    { name: "RequestsDeleteComment", options: {}, I: RequestsDeleteCommentRequest, O: RequestsDeleteCommentResponse }
]);
