// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "services/sync/sync.proto" (package "services.sync", syntax proto3)
// @ts-nocheck
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { DataVehicles } from "../../resources/sync/data";
import { DataUsers } from "../../resources/sync/data";
import { DataLicenses } from "../../resources/sync/data";
import { DataJobs } from "../../resources/sync/data";
import { TimeclockEntry } from "../../resources/jobs/timeclock";
import { JobsUserProps } from "../../resources/jobs/colleagues";
import { JobsUserActivity } from "../../resources/jobs/activity";
import { UserProps } from "../../resources/users/users";
import { UserActivity } from "../../resources/users/activity";
import { UserOAuth2Conn } from "../../resources/sync/activity";
import { DataStatus } from "../../resources/sync/data";
/**
 * @generated from protobuf message services.sync.GetStatusRequest
 */
export interface GetStatusRequest {
}
/**
 * @generated from protobuf message services.sync.GetStatusResponse
 */
export interface GetStatusResponse {
    /**
     * @generated from protobuf field: resources.sync.DataStatus jobs = 1;
     */
    jobs?: DataStatus;
    /**
     * @generated from protobuf field: resources.sync.DataStatus licenses = 2;
     */
    licenses?: DataStatus;
    /**
     * @generated from protobuf field: resources.sync.DataStatus users = 3;
     */
    users?: DataStatus;
    /**
     * @generated from protobuf field: resources.sync.DataStatus vehicles = 4;
     */
    vehicles?: DataStatus;
}
/**
 * @generated from protobuf message services.sync.AddActivityRequest
 */
export interface AddActivityRequest {
    /**
     * @generated from protobuf oneof: activity
     */
    activity: {
        oneofKind: "userOauth2";
        /**
         * @generated from protobuf field: resources.sync.UserOAuth2Conn user_oauth2 = 1;
         */
        userOauth2: UserOAuth2Conn;
    } | {
        oneofKind: "userActivity";
        /**
         * User activity
         *
         * @generated from protobuf field: resources.users.UserActivity user_activity = 2;
         */
        userActivity: UserActivity;
    } | {
        oneofKind: "userProps";
        /**
         * Setting props will cause activity to be created automtically
         *
         * @generated from protobuf field: resources.users.UserProps user_props = 3;
         */
        userProps: UserProps;
    } | {
        oneofKind: "jobsUserActivity";
        /**
         * Jobs user activity
         *
         * @generated from protobuf field: resources.jobs.JobsUserActivity jobs_user_activity = 4;
         */
        jobsUserActivity: JobsUserActivity;
    } | {
        oneofKind: "jobsUserProps";
        /**
         * Setting props will cause activity to be created automtically
         *
         * @generated from protobuf field: resources.jobs.JobsUserProps jobs_user_props = 5;
         */
        jobsUserProps: JobsUserProps;
    } | {
        oneofKind: "jobsTimeclock";
        /**
         * Timeclock
         *
         * @generated from protobuf field: resources.jobs.TimeclockEntry jobs_timeclock = 6;
         */
        jobsTimeclock: TimeclockEntry;
    } | {
        oneofKind: undefined;
    };
}
/**
 * @generated from protobuf message services.sync.AddActivityResponse
 */
export interface AddActivityResponse {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: number;
}
/**
 * @generated from protobuf message services.sync.SendDataRequest
 */
export interface SendDataRequest {
    /**
     * @generated from protobuf oneof: data
     */
    data: {
        oneofKind: "jobs";
        /**
         * @generated from protobuf field: resources.sync.DataJobs jobs = 1;
         */
        jobs: DataJobs;
    } | {
        oneofKind: "licenses";
        /**
         * @generated from protobuf field: resources.sync.DataLicenses licenses = 2;
         */
        licenses: DataLicenses;
    } | {
        oneofKind: "users";
        /**
         * @generated from protobuf field: resources.sync.DataUsers users = 3;
         */
        users: DataUsers;
    } | {
        oneofKind: "vehicles";
        /**
         * @generated from protobuf field: resources.sync.DataVehicles vehicles = 4;
         */
        vehicles: DataVehicles;
    } | {
        oneofKind: undefined;
    };
}
/**
 * @generated from protobuf message services.sync.SendDataResponse
 */
export interface SendDataResponse {
    /**
     * @generated from protobuf field: int64 affected_rows = 1;
     */
    affectedRows: number;
}
/**
 * @generated from protobuf message services.sync.StreamRequest
 */
export interface StreamRequest {
}
/**
 * @generated from protobuf message services.sync.StreamResponse
 */
export interface StreamResponse {
    /**
     * @generated from protobuf field: int32 user_id = 1;
     */
    userId: number;
}
// @generated message type with reflection information, may provide speed optimized methods
class GetStatusRequest$Type extends MessageType<GetStatusRequest> {
    constructor() {
        super("services.sync.GetStatusRequest", []);
    }
    create(value?: PartialMessage<GetStatusRequest>): GetStatusRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<GetStatusRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetStatusRequest): GetStatusRequest {
        return target ?? this.create();
    }
    internalBinaryWrite(message: GetStatusRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.sync.GetStatusRequest
 */
export const GetStatusRequest = new GetStatusRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetStatusResponse$Type extends MessageType<GetStatusResponse> {
    constructor() {
        super("services.sync.GetStatusResponse", [
            { no: 1, name: "jobs", kind: "message", T: () => DataStatus },
            { no: 2, name: "licenses", kind: "message", T: () => DataStatus },
            { no: 3, name: "users", kind: "message", T: () => DataStatus },
            { no: 4, name: "vehicles", kind: "message", T: () => DataStatus }
        ]);
    }
    create(value?: PartialMessage<GetStatusResponse>): GetStatusResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<GetStatusResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetStatusResponse): GetStatusResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.sync.DataStatus jobs */ 1:
                    message.jobs = DataStatus.internalBinaryRead(reader, reader.uint32(), options, message.jobs);
                    break;
                case /* resources.sync.DataStatus licenses */ 2:
                    message.licenses = DataStatus.internalBinaryRead(reader, reader.uint32(), options, message.licenses);
                    break;
                case /* resources.sync.DataStatus users */ 3:
                    message.users = DataStatus.internalBinaryRead(reader, reader.uint32(), options, message.users);
                    break;
                case /* resources.sync.DataStatus vehicles */ 4:
                    message.vehicles = DataStatus.internalBinaryRead(reader, reader.uint32(), options, message.vehicles);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: GetStatusResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.sync.DataStatus jobs = 1; */
        if (message.jobs)
            DataStatus.internalBinaryWrite(message.jobs, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* resources.sync.DataStatus licenses = 2; */
        if (message.licenses)
            DataStatus.internalBinaryWrite(message.licenses, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* resources.sync.DataStatus users = 3; */
        if (message.users)
            DataStatus.internalBinaryWrite(message.users, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* resources.sync.DataStatus vehicles = 4; */
        if (message.vehicles)
            DataStatus.internalBinaryWrite(message.vehicles, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.sync.GetStatusResponse
 */
export const GetStatusResponse = new GetStatusResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class AddActivityRequest$Type extends MessageType<AddActivityRequest> {
    constructor() {
        super("services.sync.AddActivityRequest", [
            { no: 1, name: "user_oauth2", kind: "message", oneof: "activity", T: () => UserOAuth2Conn },
            { no: 2, name: "user_activity", kind: "message", oneof: "activity", T: () => UserActivity },
            { no: 3, name: "user_props", kind: "message", oneof: "activity", T: () => UserProps },
            { no: 4, name: "jobs_user_activity", kind: "message", oneof: "activity", T: () => JobsUserActivity },
            { no: 5, name: "jobs_user_props", kind: "message", oneof: "activity", T: () => JobsUserProps },
            { no: 6, name: "jobs_timeclock", kind: "message", oneof: "activity", T: () => TimeclockEntry }
        ]);
    }
    create(value?: PartialMessage<AddActivityRequest>): AddActivityRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.activity = { oneofKind: undefined };
        if (value !== undefined)
            reflectionMergePartial<AddActivityRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: AddActivityRequest): AddActivityRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.sync.UserOAuth2Conn user_oauth2 */ 1:
                    message.activity = {
                        oneofKind: "userOauth2",
                        userOauth2: UserOAuth2Conn.internalBinaryRead(reader, reader.uint32(), options, (message.activity as any).userOauth2)
                    };
                    break;
                case /* resources.users.UserActivity user_activity */ 2:
                    message.activity = {
                        oneofKind: "userActivity",
                        userActivity: UserActivity.internalBinaryRead(reader, reader.uint32(), options, (message.activity as any).userActivity)
                    };
                    break;
                case /* resources.users.UserProps user_props */ 3:
                    message.activity = {
                        oneofKind: "userProps",
                        userProps: UserProps.internalBinaryRead(reader, reader.uint32(), options, (message.activity as any).userProps)
                    };
                    break;
                case /* resources.jobs.JobsUserActivity jobs_user_activity */ 4:
                    message.activity = {
                        oneofKind: "jobsUserActivity",
                        jobsUserActivity: JobsUserActivity.internalBinaryRead(reader, reader.uint32(), options, (message.activity as any).jobsUserActivity)
                    };
                    break;
                case /* resources.jobs.JobsUserProps jobs_user_props */ 5:
                    message.activity = {
                        oneofKind: "jobsUserProps",
                        jobsUserProps: JobsUserProps.internalBinaryRead(reader, reader.uint32(), options, (message.activity as any).jobsUserProps)
                    };
                    break;
                case /* resources.jobs.TimeclockEntry jobs_timeclock */ 6:
                    message.activity = {
                        oneofKind: "jobsTimeclock",
                        jobsTimeclock: TimeclockEntry.internalBinaryRead(reader, reader.uint32(), options, (message.activity as any).jobsTimeclock)
                    };
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: AddActivityRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.sync.UserOAuth2Conn user_oauth2 = 1; */
        if (message.activity.oneofKind === "userOauth2")
            UserOAuth2Conn.internalBinaryWrite(message.activity.userOauth2, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* resources.users.UserActivity user_activity = 2; */
        if (message.activity.oneofKind === "userActivity")
            UserActivity.internalBinaryWrite(message.activity.userActivity, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* resources.users.UserProps user_props = 3; */
        if (message.activity.oneofKind === "userProps")
            UserProps.internalBinaryWrite(message.activity.userProps, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* resources.jobs.JobsUserActivity jobs_user_activity = 4; */
        if (message.activity.oneofKind === "jobsUserActivity")
            JobsUserActivity.internalBinaryWrite(message.activity.jobsUserActivity, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* resources.jobs.JobsUserProps jobs_user_props = 5; */
        if (message.activity.oneofKind === "jobsUserProps")
            JobsUserProps.internalBinaryWrite(message.activity.jobsUserProps, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* resources.jobs.TimeclockEntry jobs_timeclock = 6; */
        if (message.activity.oneofKind === "jobsTimeclock")
            TimeclockEntry.internalBinaryWrite(message.activity.jobsTimeclock, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.sync.AddActivityRequest
 */
export const AddActivityRequest = new AddActivityRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class AddActivityResponse$Type extends MessageType<AddActivityResponse> {
    constructor() {
        super("services.sync.AddActivityResponse", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 2 /*LongType.NUMBER*/ }
        ]);
    }
    create(value?: PartialMessage<AddActivityResponse>): AddActivityResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.id = 0;
        if (value !== undefined)
            reflectionMergePartial<AddActivityResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: AddActivityResponse): AddActivityResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 id */ 1:
                    message.id = reader.uint64().toNumber();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: AddActivityResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 id = 1; */
        if (message.id !== 0)
            writer.tag(1, WireType.Varint).uint64(message.id);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.sync.AddActivityResponse
 */
export const AddActivityResponse = new AddActivityResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SendDataRequest$Type extends MessageType<SendDataRequest> {
    constructor() {
        super("services.sync.SendDataRequest", [
            { no: 1, name: "jobs", kind: "message", oneof: "data", T: () => DataJobs },
            { no: 2, name: "licenses", kind: "message", oneof: "data", T: () => DataLicenses },
            { no: 3, name: "users", kind: "message", oneof: "data", T: () => DataUsers },
            { no: 4, name: "vehicles", kind: "message", oneof: "data", T: () => DataVehicles }
        ]);
    }
    create(value?: PartialMessage<SendDataRequest>): SendDataRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.data = { oneofKind: undefined };
        if (value !== undefined)
            reflectionMergePartial<SendDataRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SendDataRequest): SendDataRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.sync.DataJobs jobs */ 1:
                    message.data = {
                        oneofKind: "jobs",
                        jobs: DataJobs.internalBinaryRead(reader, reader.uint32(), options, (message.data as any).jobs)
                    };
                    break;
                case /* resources.sync.DataLicenses licenses */ 2:
                    message.data = {
                        oneofKind: "licenses",
                        licenses: DataLicenses.internalBinaryRead(reader, reader.uint32(), options, (message.data as any).licenses)
                    };
                    break;
                case /* resources.sync.DataUsers users */ 3:
                    message.data = {
                        oneofKind: "users",
                        users: DataUsers.internalBinaryRead(reader, reader.uint32(), options, (message.data as any).users)
                    };
                    break;
                case /* resources.sync.DataVehicles vehicles */ 4:
                    message.data = {
                        oneofKind: "vehicles",
                        vehicles: DataVehicles.internalBinaryRead(reader, reader.uint32(), options, (message.data as any).vehicles)
                    };
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: SendDataRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.sync.DataJobs jobs = 1; */
        if (message.data.oneofKind === "jobs")
            DataJobs.internalBinaryWrite(message.data.jobs, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* resources.sync.DataLicenses licenses = 2; */
        if (message.data.oneofKind === "licenses")
            DataLicenses.internalBinaryWrite(message.data.licenses, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* resources.sync.DataUsers users = 3; */
        if (message.data.oneofKind === "users")
            DataUsers.internalBinaryWrite(message.data.users, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* resources.sync.DataVehicles vehicles = 4; */
        if (message.data.oneofKind === "vehicles")
            DataVehicles.internalBinaryWrite(message.data.vehicles, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.sync.SendDataRequest
 */
export const SendDataRequest = new SendDataRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SendDataResponse$Type extends MessageType<SendDataResponse> {
    constructor() {
        super("services.sync.SendDataResponse", [
            { no: 1, name: "affected_rows", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 2 /*LongType.NUMBER*/ }
        ]);
    }
    create(value?: PartialMessage<SendDataResponse>): SendDataResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.affectedRows = 0;
        if (value !== undefined)
            reflectionMergePartial<SendDataResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SendDataResponse): SendDataResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int64 affected_rows */ 1:
                    message.affectedRows = reader.int64().toNumber();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: SendDataResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int64 affected_rows = 1; */
        if (message.affectedRows !== 0)
            writer.tag(1, WireType.Varint).int64(message.affectedRows);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.sync.SendDataResponse
 */
export const SendDataResponse = new SendDataResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StreamRequest$Type extends MessageType<StreamRequest> {
    constructor() {
        super("services.sync.StreamRequest", []);
    }
    create(value?: PartialMessage<StreamRequest>): StreamRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<StreamRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: StreamRequest): StreamRequest {
        return target ?? this.create();
    }
    internalBinaryWrite(message: StreamRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.sync.StreamRequest
 */
export const StreamRequest = new StreamRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StreamResponse$Type extends MessageType<StreamResponse> {
    constructor() {
        super("services.sync.StreamResponse", [
            { no: 1, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value?: PartialMessage<StreamResponse>): StreamResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.userId = 0;
        if (value !== undefined)
            reflectionMergePartial<StreamResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: StreamResponse): StreamResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 user_id */ 1:
                    message.userId = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: StreamResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 user_id = 1; */
        if (message.userId !== 0)
            writer.tag(1, WireType.Varint).int32(message.userId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.sync.StreamResponse
 */
export const StreamResponse = new StreamResponse$Type();
/**
 * @generated ServiceType for protobuf service services.sync.SyncService
 */
export const SyncService = new ServiceType("services.sync.SyncService", [
    { name: "GetStatus", options: {}, I: GetStatusRequest, O: GetStatusResponse },
    { name: "AddActivity", options: {}, I: AddActivityRequest, O: AddActivityResponse },
    { name: "SendData", options: {}, I: SendDataRequest, O: SendDataResponse },
    { name: "Stream", serverStreaming: true, options: {}, I: StreamRequest, O: StreamResponse }
]);
