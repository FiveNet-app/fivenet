// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_bigint
// @generated from protobuf file "services/notificator/notificator.proto" (package "services.notificator", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { JobProps } from "../../resources/users/jobs";
import { User } from "../../resources/users/users";
import { Timestamp } from "../../resources/timestamp/timestamp";
import { Notification } from "../../resources/notifications/notifications";
import { PaginationResponse } from "../../resources/common/database/database";
import { PaginationRequest } from "../../resources/common/database/database";
/**
 * @generated from protobuf message services.notificator.GetNotificationsRequest
 */
export interface GetNotificationsRequest {
    /**
     * @generated from protobuf field: resources.common.database.PaginationRequest pagination = 1;
     */
    pagination?: PaginationRequest;
    /**
     * @generated from protobuf field: optional bool include_read = 2;
     */
    includeRead?: boolean;
}
/**
 * @generated from protobuf message services.notificator.GetNotificationsResponse
 */
export interface GetNotificationsResponse {
    /**
     * @generated from protobuf field: resources.common.database.PaginationResponse pagination = 1;
     */
    pagination?: PaginationResponse;
    /**
     * @generated from protobuf field: repeated resources.notifications.Notification notifications = 2;
     */
    notifications: Notification[];
}
/**
 * @generated from protobuf message services.notificator.MarkNotificationsRequest
 */
export interface MarkNotificationsRequest {
    /**
     * @generated from protobuf field: repeated uint64 ids = 1 [jstype = JS_STRING];
     */
    ids: string[];
    /**
     * @generated from protobuf field: optional bool all = 2;
     */
    all?: boolean;
}
/**
 * @generated from protobuf message services.notificator.MarkNotificationsResponse
 */
export interface MarkNotificationsResponse {
    /**
     * @generated from protobuf field: uint64 updated = 1;
     */
    updated: bigint;
}
/**
 * @generated from protobuf message services.notificator.StreamRequest
 */
export interface StreamRequest {
    /**
     * @generated from protobuf field: uint64 last_id = 1;
     */
    lastId: bigint;
}
/**
 * @generated from protobuf message services.notificator.StreamResponse
 */
export interface StreamResponse {
    /**
     * @generated from protobuf field: uint64 last_id = 1;
     */
    lastId: bigint;
    /**
     * @generated from protobuf oneof: data
     */
    data: {
        oneofKind: "notifications";
        /**
         * @generated from protobuf field: services.notificator.StreamNotifications notifications = 2;
         */
        notifications: StreamNotifications;
    } | {
        oneofKind: "token";
        /**
         * @generated from protobuf field: services.notificator.TokenUpdate token = 3;
         */
        token: TokenUpdate;
    } | {
        oneofKind: undefined;
    };
    /**
     * @generated from protobuf field: optional bool restart = 5;
     */
    restart?: boolean;
}
/**
 * @generated from protobuf message services.notificator.StreamNotifications
 */
export interface StreamNotifications {
    /**
     * @generated from protobuf field: repeated resources.notifications.Notification notifications = 1;
     */
    notifications: Notification[];
}
/**
 * @generated from protobuf message services.notificator.TokenUpdate
 */
export interface TokenUpdate {
    /**
     * @generated from protobuf field: optional string new_token = 1;
     */
    newToken?: string;
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp expires = 2;
     */
    expires?: Timestamp;
    /**
     * @generated from protobuf field: repeated string permissions = 3;
     */
    permissions: string[];
    /**
     * @generated from protobuf field: optional resources.users.User user_info = 4;
     */
    userInfo?: User;
    /**
     * @generated from protobuf field: optional resources.users.JobProps job_props = 5;
     */
    jobProps?: JobProps;
}
// @generated message type with reflection information, may provide speed optimized methods
class GetNotificationsRequest$Type extends MessageType<GetNotificationsRequest> {
    constructor() {
        super("services.notificator.GetNotificationsRequest", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationRequest, options: { "validate.rules": { message: { required: true } } } },
            { no: 2, name: "include_read", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value?: PartialMessage<GetNotificationsRequest>): GetNotificationsRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<GetNotificationsRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetNotificationsRequest): GetNotificationsRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.common.database.PaginationRequest pagination */ 1:
                    message.pagination = PaginationRequest.internalBinaryRead(reader, reader.uint32(), options, message.pagination);
                    break;
                case /* optional bool include_read */ 2:
                    message.includeRead = reader.bool();
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
    internalBinaryWrite(message: GetNotificationsRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.common.database.PaginationRequest pagination = 1; */
        if (message.pagination)
            PaginationRequest.internalBinaryWrite(message.pagination, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* optional bool include_read = 2; */
        if (message.includeRead !== undefined)
            writer.tag(2, WireType.Varint).bool(message.includeRead);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.notificator.GetNotificationsRequest
 */
export const GetNotificationsRequest = new GetNotificationsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetNotificationsResponse$Type extends MessageType<GetNotificationsResponse> {
    constructor() {
        super("services.notificator.GetNotificationsResponse", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationResponse },
            { no: 2, name: "notifications", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Notification }
        ]);
    }
    create(value?: PartialMessage<GetNotificationsResponse>): GetNotificationsResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.notifications = [];
        if (value !== undefined)
            reflectionMergePartial<GetNotificationsResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetNotificationsResponse): GetNotificationsResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.common.database.PaginationResponse pagination */ 1:
                    message.pagination = PaginationResponse.internalBinaryRead(reader, reader.uint32(), options, message.pagination);
                    break;
                case /* repeated resources.notifications.Notification notifications */ 2:
                    message.notifications.push(Notification.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: GetNotificationsResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.common.database.PaginationResponse pagination = 1; */
        if (message.pagination)
            PaginationResponse.internalBinaryWrite(message.pagination, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* repeated resources.notifications.Notification notifications = 2; */
        for (let i = 0; i < message.notifications.length; i++)
            Notification.internalBinaryWrite(message.notifications[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.notificator.GetNotificationsResponse
 */
export const GetNotificationsResponse = new GetNotificationsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MarkNotificationsRequest$Type extends MessageType<MarkNotificationsRequest> {
    constructor() {
        super("services.notificator.MarkNotificationsRequest", [
            { no: 1, name: "ids", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 4 /*ScalarType.UINT64*/, options: { "validate.rules": { repeated: { minItems: "1", maxItems: "20", ignoreEmpty: true } } } },
            { no: 2, name: "all", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value?: PartialMessage<MarkNotificationsRequest>): MarkNotificationsRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.ids = [];
        if (value !== undefined)
            reflectionMergePartial<MarkNotificationsRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: MarkNotificationsRequest): MarkNotificationsRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated uint64 ids = 1 [jstype = JS_STRING];*/ 1:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.ids.push(reader.uint64().toString());
                    else
                        message.ids.push(reader.uint64().toString());
                    break;
                case /* optional bool all */ 2:
                    message.all = reader.bool();
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
    internalBinaryWrite(message: MarkNotificationsRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated uint64 ids = 1 [jstype = JS_STRING]; */
        if (message.ids.length) {
            writer.tag(1, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.ids.length; i++)
                writer.uint64(message.ids[i]);
            writer.join();
        }
        /* optional bool all = 2; */
        if (message.all !== undefined)
            writer.tag(2, WireType.Varint).bool(message.all);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.notificator.MarkNotificationsRequest
 */
export const MarkNotificationsRequest = new MarkNotificationsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MarkNotificationsResponse$Type extends MessageType<MarkNotificationsResponse> {
    constructor() {
        super("services.notificator.MarkNotificationsResponse", [
            { no: 1, name: "updated", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<MarkNotificationsResponse>): MarkNotificationsResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.updated = 0n;
        if (value !== undefined)
            reflectionMergePartial<MarkNotificationsResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: MarkNotificationsResponse): MarkNotificationsResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 updated */ 1:
                    message.updated = reader.uint64().toBigInt();
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
    internalBinaryWrite(message: MarkNotificationsResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 updated = 1; */
        if (message.updated !== 0n)
            writer.tag(1, WireType.Varint).uint64(message.updated);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.notificator.MarkNotificationsResponse
 */
export const MarkNotificationsResponse = new MarkNotificationsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StreamRequest$Type extends MessageType<StreamRequest> {
    constructor() {
        super("services.notificator.StreamRequest", [
            { no: 1, name: "last_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<StreamRequest>): StreamRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.lastId = 0n;
        if (value !== undefined)
            reflectionMergePartial<StreamRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: StreamRequest): StreamRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 last_id */ 1:
                    message.lastId = reader.uint64().toBigInt();
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
    internalBinaryWrite(message: StreamRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 last_id = 1; */
        if (message.lastId !== 0n)
            writer.tag(1, WireType.Varint).uint64(message.lastId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.notificator.StreamRequest
 */
export const StreamRequest = new StreamRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StreamResponse$Type extends MessageType<StreamResponse> {
    constructor() {
        super("services.notificator.StreamResponse", [
            { no: 1, name: "last_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "notifications", kind: "message", oneof: "data", T: () => StreamNotifications },
            { no: 3, name: "token", kind: "message", oneof: "data", T: () => TokenUpdate },
            { no: 5, name: "restart", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value?: PartialMessage<StreamResponse>): StreamResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.lastId = 0n;
        message.data = { oneofKind: undefined };
        if (value !== undefined)
            reflectionMergePartial<StreamResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: StreamResponse): StreamResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 last_id */ 1:
                    message.lastId = reader.uint64().toBigInt();
                    break;
                case /* services.notificator.StreamNotifications notifications */ 2:
                    message.data = {
                        oneofKind: "notifications",
                        notifications: StreamNotifications.internalBinaryRead(reader, reader.uint32(), options, (message.data as any).notifications)
                    };
                    break;
                case /* services.notificator.TokenUpdate token */ 3:
                    message.data = {
                        oneofKind: "token",
                        token: TokenUpdate.internalBinaryRead(reader, reader.uint32(), options, (message.data as any).token)
                    };
                    break;
                case /* optional bool restart */ 5:
                    message.restart = reader.bool();
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
        /* uint64 last_id = 1; */
        if (message.lastId !== 0n)
            writer.tag(1, WireType.Varint).uint64(message.lastId);
        /* services.notificator.StreamNotifications notifications = 2; */
        if (message.data.oneofKind === "notifications")
            StreamNotifications.internalBinaryWrite(message.data.notifications, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* services.notificator.TokenUpdate token = 3; */
        if (message.data.oneofKind === "token")
            TokenUpdate.internalBinaryWrite(message.data.token, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* optional bool restart = 5; */
        if (message.restart !== undefined)
            writer.tag(5, WireType.Varint).bool(message.restart);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.notificator.StreamResponse
 */
export const StreamResponse = new StreamResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StreamNotifications$Type extends MessageType<StreamNotifications> {
    constructor() {
        super("services.notificator.StreamNotifications", [
            { no: 1, name: "notifications", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Notification }
        ]);
    }
    create(value?: PartialMessage<StreamNotifications>): StreamNotifications {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.notifications = [];
        if (value !== undefined)
            reflectionMergePartial<StreamNotifications>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: StreamNotifications): StreamNotifications {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated resources.notifications.Notification notifications */ 1:
                    message.notifications.push(Notification.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: StreamNotifications, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated resources.notifications.Notification notifications = 1; */
        for (let i = 0; i < message.notifications.length; i++)
            Notification.internalBinaryWrite(message.notifications[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.notificator.StreamNotifications
 */
export const StreamNotifications = new StreamNotifications$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TokenUpdate$Type extends MessageType<TokenUpdate> {
    constructor() {
        super("services.notificator.TokenUpdate", [
            { no: 1, name: "new_token", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "expires", kind: "message", T: () => Timestamp },
            { no: 3, name: "permissions", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "user_info", kind: "message", T: () => User },
            { no: 5, name: "job_props", kind: "message", T: () => JobProps }
        ]);
    }
    create(value?: PartialMessage<TokenUpdate>): TokenUpdate {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.permissions = [];
        if (value !== undefined)
            reflectionMergePartial<TokenUpdate>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: TokenUpdate): TokenUpdate {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* optional string new_token */ 1:
                    message.newToken = reader.string();
                    break;
                case /* resources.timestamp.Timestamp expires */ 2:
                    message.expires = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.expires);
                    break;
                case /* repeated string permissions */ 3:
                    message.permissions.push(reader.string());
                    break;
                case /* optional resources.users.User user_info */ 4:
                    message.userInfo = User.internalBinaryRead(reader, reader.uint32(), options, message.userInfo);
                    break;
                case /* optional resources.users.JobProps job_props */ 5:
                    message.jobProps = JobProps.internalBinaryRead(reader, reader.uint32(), options, message.jobProps);
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
    internalBinaryWrite(message: TokenUpdate, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* optional string new_token = 1; */
        if (message.newToken !== undefined)
            writer.tag(1, WireType.LengthDelimited).string(message.newToken);
        /* resources.timestamp.Timestamp expires = 2; */
        if (message.expires)
            Timestamp.internalBinaryWrite(message.expires, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* repeated string permissions = 3; */
        for (let i = 0; i < message.permissions.length; i++)
            writer.tag(3, WireType.LengthDelimited).string(message.permissions[i]);
        /* optional resources.users.User user_info = 4; */
        if (message.userInfo)
            User.internalBinaryWrite(message.userInfo, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.users.JobProps job_props = 5; */
        if (message.jobProps)
            JobProps.internalBinaryWrite(message.jobProps, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.notificator.TokenUpdate
 */
export const TokenUpdate = new TokenUpdate$Type();
/**
 * @generated ServiceType for protobuf service services.notificator.NotificatorService
 */
export const NotificatorService = new ServiceType("services.notificator.NotificatorService", [
    { name: "GetNotifications", options: {}, I: GetNotificationsRequest, O: GetNotificationsResponse },
    { name: "MarkNotifications", options: {}, I: MarkNotificationsRequest, O: MarkNotificationsResponse },
    { name: "Stream", serverStreaming: true, options: {}, I: StreamRequest, O: StreamResponse }
]);
