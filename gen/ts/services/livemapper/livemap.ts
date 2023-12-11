// @generated by protobuf-ts 2.9.3 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "services/livemapper/livemap.proto" (package "services.livemapper", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { MessageType } from "@protobuf-ts/runtime";
import { Marker } from "../../resources/livemap/livemap";
import { UserMarker } from "../../resources/livemap/livemap";
import { Job } from "../../resources/users/jobs";
/**
 * @generated from protobuf message services.livemapper.StreamRequest
 */
export interface StreamRequest {
}
/**
 * @generated from protobuf message services.livemapper.StreamResponse
 */
export interface StreamResponse {
    /**
     * @generated from protobuf field: repeated resources.users.Job jobs_users = 1;
     */
    jobsUsers: Job[];
    /**
     * @generated from protobuf field: repeated resources.livemap.UserMarker users = 2;
     */
    users: UserMarker[];
    /**
     * @generated from protobuf field: repeated resources.users.Job jobs_markers = 3;
     */
    jobsMarkers: Job[];
    /**
     * @generated from protobuf field: repeated resources.livemap.Marker markers = 4;
     */
    markers: Marker[];
}
/**
 * @generated from protobuf message services.livemapper.CreateOrUpdateMarkerRequest
 */
export interface CreateOrUpdateMarkerRequest {
    /**
     * @generated from protobuf field: resources.livemap.Marker marker = 1;
     */
    marker?: Marker;
}
/**
 * @generated from protobuf message services.livemapper.CreateOrUpdateMarkerResponse
 */
export interface CreateOrUpdateMarkerResponse {
    /**
     * @generated from protobuf field: resources.livemap.Marker marker = 1;
     */
    marker?: Marker;
}
/**
 * @generated from protobuf message services.livemapper.DeleteMarkerRequest
 */
export interface DeleteMarkerRequest {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
}
/**
 * @generated from protobuf message services.livemapper.DeleteMarkerResponse
 */
export interface DeleteMarkerResponse {
}
// @generated message type with reflection information, may provide speed optimized methods
class StreamRequest$Type extends MessageType<StreamRequest> {
    constructor() {
        super("services.livemapper.StreamRequest", []);
    }
}
/**
 * @generated MessageType for protobuf message services.livemapper.StreamRequest
 */
export const StreamRequest = new StreamRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StreamResponse$Type extends MessageType<StreamResponse> {
    constructor() {
        super("services.livemapper.StreamResponse", [
            { no: 1, name: "jobs_users", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Job },
            { no: 2, name: "users", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => UserMarker },
            { no: 3, name: "jobs_markers", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Job },
            { no: 4, name: "markers", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Marker }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.livemapper.StreamResponse
 */
export const StreamResponse = new StreamResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateOrUpdateMarkerRequest$Type extends MessageType<CreateOrUpdateMarkerRequest> {
    constructor() {
        super("services.livemapper.CreateOrUpdateMarkerRequest", [
            { no: 1, name: "marker", kind: "message", T: () => Marker, options: { "validate.rules": { message: { required: true } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.livemapper.CreateOrUpdateMarkerRequest
 */
export const CreateOrUpdateMarkerRequest = new CreateOrUpdateMarkerRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateOrUpdateMarkerResponse$Type extends MessageType<CreateOrUpdateMarkerResponse> {
    constructor() {
        super("services.livemapper.CreateOrUpdateMarkerResponse", [
            { no: 1, name: "marker", kind: "message", T: () => Marker }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.livemapper.CreateOrUpdateMarkerResponse
 */
export const CreateOrUpdateMarkerResponse = new CreateOrUpdateMarkerResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteMarkerRequest$Type extends MessageType<DeleteMarkerRequest> {
    constructor() {
        super("services.livemapper.DeleteMarkerRequest", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.livemapper.DeleteMarkerRequest
 */
export const DeleteMarkerRequest = new DeleteMarkerRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteMarkerResponse$Type extends MessageType<DeleteMarkerResponse> {
    constructor() {
        super("services.livemapper.DeleteMarkerResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.livemapper.DeleteMarkerResponse
 */
export const DeleteMarkerResponse = new DeleteMarkerResponse$Type();
/**
 * @generated ServiceType for protobuf service services.livemapper.LivemapperService
 */
export const LivemapperService = new ServiceType("services.livemapper.LivemapperService", [
    { name: "Stream", serverStreaming: true, options: {}, I: StreamRequest, O: StreamResponse },
    { name: "CreateOrUpdateMarker", options: {}, I: CreateOrUpdateMarkerRequest, O: CreateOrUpdateMarkerResponse },
    { name: "DeleteMarker", options: {}, I: DeleteMarkerRequest, O: DeleteMarkerResponse }
]);
