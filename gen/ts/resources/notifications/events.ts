// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "resources/notifications/events.proto" (package "resources.notifications", syntax proto3)
// @ts-nocheck
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { JobProps } from "../users/job_props";
import { Notification } from "./notifications";
// User Events

/**
 * @generated from protobuf message resources.notifications.UserEvent
 */
export interface UserEvent {
    /**
     * @generated from protobuf oneof: data
     */
    data: {
        oneofKind: "refreshToken";
        /**
         * @generated from protobuf field: bool refresh_token = 1;
         */
        refreshToken: boolean;
    } | {
        oneofKind: "notification";
        /**
         * Notifications
         *
         * @generated from protobuf field: resources.notifications.Notification notification = 2;
         */
        notification: Notification;
    } | {
        oneofKind: "notificationsReadCount";
        /**
         * @generated from protobuf field: int32 notifications_read_count = 3;
         */
        notificationsReadCount: number;
    } | {
        oneofKind: undefined;
    };
}
// Job Events

/**
 * @generated from protobuf message resources.notifications.JobEvent
 */
export interface JobEvent {
    /**
     * @generated from protobuf oneof: data
     */
    data: {
        oneofKind: "jobProps";
        /**
         * @generated from protobuf field: resources.users.JobProps job_props = 1;
         */
        jobProps: JobProps;
    } | {
        oneofKind: undefined;
    };
}
// Job Grade Events

/**
 * @generated from protobuf message resources.notifications.JobGradeEvent
 */
export interface JobGradeEvent {
    /**
     * @generated from protobuf oneof: data
     */
    data: {
        oneofKind: "refreshToken";
        /**
         * @generated from protobuf field: bool refresh_token = 1;
         */
        refreshToken: boolean;
    } | {
        oneofKind: undefined;
    };
}
// System Events

/**
 * @generated from protobuf message resources.notifications.SystemEvent
 */
export interface SystemEvent {
    /**
     * @generated from protobuf oneof: data
     */
    data: {
        oneofKind: "ping";
        /**
         * @generated from protobuf field: bool ping = 1;
         */
        ping: boolean;
    } | {
        oneofKind: undefined;
    };
}
// @generated message type with reflection information, may provide speed optimized methods
class UserEvent$Type extends MessageType<UserEvent> {
    constructor() {
        super("resources.notifications.UserEvent", [
            { no: 1, name: "refresh_token", kind: "scalar", oneof: "data", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "notification", kind: "message", oneof: "data", T: () => Notification },
            { no: 3, name: "notifications_read_count", kind: "scalar", oneof: "data", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value?: PartialMessage<UserEvent>): UserEvent {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.data = { oneofKind: undefined };
        if (value !== undefined)
            reflectionMergePartial<UserEvent>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UserEvent): UserEvent {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool refresh_token */ 1:
                    message.data = {
                        oneofKind: "refreshToken",
                        refreshToken: reader.bool()
                    };
                    break;
                case /* resources.notifications.Notification notification */ 2:
                    message.data = {
                        oneofKind: "notification",
                        notification: Notification.internalBinaryRead(reader, reader.uint32(), options, (message.data as any).notification)
                    };
                    break;
                case /* int32 notifications_read_count */ 3:
                    message.data = {
                        oneofKind: "notificationsReadCount",
                        notificationsReadCount: reader.int32()
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
    internalBinaryWrite(message: UserEvent, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bool refresh_token = 1; */
        if (message.data.oneofKind === "refreshToken")
            writer.tag(1, WireType.Varint).bool(message.data.refreshToken);
        /* resources.notifications.Notification notification = 2; */
        if (message.data.oneofKind === "notification")
            Notification.internalBinaryWrite(message.data.notification, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* int32 notifications_read_count = 3; */
        if (message.data.oneofKind === "notificationsReadCount")
            writer.tag(3, WireType.Varint).int32(message.data.notificationsReadCount);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.notifications.UserEvent
 */
export const UserEvent = new UserEvent$Type();
// @generated message type with reflection information, may provide speed optimized methods
class JobEvent$Type extends MessageType<JobEvent> {
    constructor() {
        super("resources.notifications.JobEvent", [
            { no: 1, name: "job_props", kind: "message", oneof: "data", T: () => JobProps }
        ]);
    }
    create(value?: PartialMessage<JobEvent>): JobEvent {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.data = { oneofKind: undefined };
        if (value !== undefined)
            reflectionMergePartial<JobEvent>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: JobEvent): JobEvent {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.users.JobProps job_props */ 1:
                    message.data = {
                        oneofKind: "jobProps",
                        jobProps: JobProps.internalBinaryRead(reader, reader.uint32(), options, (message.data as any).jobProps)
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
    internalBinaryWrite(message: JobEvent, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.users.JobProps job_props = 1; */
        if (message.data.oneofKind === "jobProps")
            JobProps.internalBinaryWrite(message.data.jobProps, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.notifications.JobEvent
 */
export const JobEvent = new JobEvent$Type();
// @generated message type with reflection information, may provide speed optimized methods
class JobGradeEvent$Type extends MessageType<JobGradeEvent> {
    constructor() {
        super("resources.notifications.JobGradeEvent", [
            { no: 1, name: "refresh_token", kind: "scalar", oneof: "data", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value?: PartialMessage<JobGradeEvent>): JobGradeEvent {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.data = { oneofKind: undefined };
        if (value !== undefined)
            reflectionMergePartial<JobGradeEvent>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: JobGradeEvent): JobGradeEvent {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool refresh_token */ 1:
                    message.data = {
                        oneofKind: "refreshToken",
                        refreshToken: reader.bool()
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
    internalBinaryWrite(message: JobGradeEvent, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bool refresh_token = 1; */
        if (message.data.oneofKind === "refreshToken")
            writer.tag(1, WireType.Varint).bool(message.data.refreshToken);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.notifications.JobGradeEvent
 */
export const JobGradeEvent = new JobGradeEvent$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SystemEvent$Type extends MessageType<SystemEvent> {
    constructor() {
        super("resources.notifications.SystemEvent", [
            { no: 1, name: "ping", kind: "scalar", oneof: "data", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value?: PartialMessage<SystemEvent>): SystemEvent {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.data = { oneofKind: undefined };
        if (value !== undefined)
            reflectionMergePartial<SystemEvent>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SystemEvent): SystemEvent {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool ping */ 1:
                    message.data = {
                        oneofKind: "ping",
                        ping: reader.bool()
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
    internalBinaryWrite(message: SystemEvent, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bool ping = 1; */
        if (message.data.oneofKind === "ping")
            writer.tag(1, WireType.Varint).bool(message.data.ping);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.notifications.SystemEvent
 */
export const SystemEvent = new SystemEvent$Type();
