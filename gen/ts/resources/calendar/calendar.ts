// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "resources/calendar/calendar.proto" (package "resources.calendar", syntax proto3)
// tslint:disable
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { UserShort } from "../users/users";
import { Timestamp } from "../timestamp/timestamp";
/**
 * @generated from protobuf message resources.calendar.Calendar
 */
export interface Calendar {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string; // @gotags: sql:"primary_key" alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
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
     * @generated from protobuf field: optional string job = 5;
     */
    job?: string;
    /**
     * @sanitize
     *
     * @generated from protobuf field: string name = 6;
     */
    name: string;
    /**
     * @sanitize: method=StripTags
     *
     * @generated from protobuf field: optional string description = 7;
     */
    description?: string;
    /**
     * @generated from protobuf field: bool public = 8;
     */
    public: boolean;
    /**
     * @generated from protobuf field: bool closed = 9;
     */
    closed: boolean;
    /**
     * @generated from protobuf field: optional int32 creator_id = 10;
     */
    creatorId?: number;
    /**
     * @generated from protobuf field: optional resources.users.UserShort creator = 11;
     */
    creator?: UserShort; // @gotags: alias:"creator"
    /**
     * @generated from protobuf field: string creator_job = 12;
     */
    creatorJob: string;
}
/**
 * @generated from protobuf message resources.calendar.CalendarEntry
 */
export interface CalendarEntry {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string; // @gotags: sql:"primary_key" alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
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
     * @generated from protobuf field: uint64 calendar_id = 5;
     */
    calendarId: number;
    /**
     * @generated from protobuf field: optional string job = 6;
     */
    job?: string;
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp start_time = 7;
     */
    startTime?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp end_time = 8;
     */
    endTime?: Timestamp;
    /**
     * @sanitize
     *
     * @generated from protobuf field: string title = 9;
     */
    title: string;
    /**
     * @sanitize: method=StripTags
     *
     * @generated from protobuf field: optional string color = 10;
     */
    color?: string;
    /**
     * @sanitize
     *
     * @generated from protobuf field: string content = 11;
     */
    content: string;
    /**
     * @generated from protobuf field: bool public = 12;
     */
    public: boolean;
    /**
     * @generated from protobuf field: optional int32 creator_id = 13;
     */
    creatorId?: number;
    /**
     * @generated from protobuf field: optional resources.users.UserShort creator = 14;
     */
    creator?: UserShort; // @gotags: alias:"creator"
    /**
     * @generated from protobuf field: string creator_job = 15;
     */
    creatorJob: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class Calendar$Type extends MessageType<Calendar> {
    constructor() {
        super("resources.calendar.Calendar", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "deleted_at", kind: "message", T: () => Timestamp },
            { no: 5, name: "job", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 6, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3", maxLen: "255" } } } },
            { no: 7, name: "description", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "512" } } } },
            { no: 8, name: "public", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 9, name: "closed", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 10, name: "creator_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ },
            { no: 11, name: "creator", kind: "message", T: () => UserShort },
            { no: 12, name: "creator_job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } }
        ]);
    }
    create(value?: PartialMessage<Calendar>): Calendar {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.id = "0";
        message.name = "";
        message.public = false;
        message.closed = false;
        message.creatorJob = "";
        if (value !== undefined)
            reflectionMergePartial<Calendar>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Calendar): Calendar {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 id = 1 [jstype = JS_STRING];*/ 1:
                    message.id = reader.uint64().toString();
                    break;
                case /* optional resources.timestamp.Timestamp created_at */ 2:
                    message.createdAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.createdAt);
                    break;
                case /* optional resources.timestamp.Timestamp updated_at */ 3:
                    message.updatedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.updatedAt);
                    break;
                case /* optional resources.timestamp.Timestamp deleted_at */ 4:
                    message.deletedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.deletedAt);
                    break;
                case /* optional string job */ 5:
                    message.job = reader.string();
                    break;
                case /* string name */ 6:
                    message.name = reader.string();
                    break;
                case /* optional string description */ 7:
                    message.description = reader.string();
                    break;
                case /* bool public */ 8:
                    message.public = reader.bool();
                    break;
                case /* bool closed */ 9:
                    message.closed = reader.bool();
                    break;
                case /* optional int32 creator_id */ 10:
                    message.creatorId = reader.int32();
                    break;
                case /* optional resources.users.UserShort creator */ 11:
                    message.creator = UserShort.internalBinaryRead(reader, reader.uint32(), options, message.creator);
                    break;
                case /* string creator_job */ 12:
                    message.creatorJob = reader.string();
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
    internalBinaryWrite(message: Calendar, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 id = 1 [jstype = JS_STRING]; */
        if (message.id !== "0")
            writer.tag(1, WireType.Varint).uint64(message.id);
        /* optional resources.timestamp.Timestamp created_at = 2; */
        if (message.createdAt)
            Timestamp.internalBinaryWrite(message.createdAt, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.timestamp.Timestamp updated_at = 3; */
        if (message.updatedAt)
            Timestamp.internalBinaryWrite(message.updatedAt, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.timestamp.Timestamp deleted_at = 4; */
        if (message.deletedAt)
            Timestamp.internalBinaryWrite(message.deletedAt, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* optional string job = 5; */
        if (message.job !== undefined)
            writer.tag(5, WireType.LengthDelimited).string(message.job);
        /* string name = 6; */
        if (message.name !== "")
            writer.tag(6, WireType.LengthDelimited).string(message.name);
        /* optional string description = 7; */
        if (message.description !== undefined)
            writer.tag(7, WireType.LengthDelimited).string(message.description);
        /* bool public = 8; */
        if (message.public !== false)
            writer.tag(8, WireType.Varint).bool(message.public);
        /* bool closed = 9; */
        if (message.closed !== false)
            writer.tag(9, WireType.Varint).bool(message.closed);
        /* optional int32 creator_id = 10; */
        if (message.creatorId !== undefined)
            writer.tag(10, WireType.Varint).int32(message.creatorId);
        /* optional resources.users.UserShort creator = 11; */
        if (message.creator)
            UserShort.internalBinaryWrite(message.creator, writer.tag(11, WireType.LengthDelimited).fork(), options).join();
        /* string creator_job = 12; */
        if (message.creatorJob !== "")
            writer.tag(12, WireType.LengthDelimited).string(message.creatorJob);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.calendar.Calendar
 */
export const Calendar = new Calendar$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CalendarEntry$Type extends MessageType<CalendarEntry> {
    constructor() {
        super("resources.calendar.CalendarEntry", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "deleted_at", kind: "message", T: () => Timestamp },
            { no: 5, name: "calendar_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 2 /*LongType.NUMBER*/ },
            { no: 6, name: "job", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 7, name: "start_time", kind: "message", T: () => Timestamp },
            { no: 8, name: "end_time", kind: "message", T: () => Timestamp },
            { no: 9, name: "title", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3", maxLen: "1024" } } } },
            { no: 10, name: "color", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "12" } } } },
            { no: 11, name: "content", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "20", maxBytes: "1000000" } } } },
            { no: 12, name: "public", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 13, name: "creator_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ },
            { no: 14, name: "creator", kind: "message", T: () => UserShort },
            { no: 15, name: "creator_job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } }
        ]);
    }
    create(value?: PartialMessage<CalendarEntry>): CalendarEntry {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.id = "0";
        message.calendarId = 0;
        message.title = "";
        message.content = "";
        message.public = false;
        message.creatorJob = "";
        if (value !== undefined)
            reflectionMergePartial<CalendarEntry>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CalendarEntry): CalendarEntry {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 id = 1 [jstype = JS_STRING];*/ 1:
                    message.id = reader.uint64().toString();
                    break;
                case /* optional resources.timestamp.Timestamp created_at */ 2:
                    message.createdAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.createdAt);
                    break;
                case /* optional resources.timestamp.Timestamp updated_at */ 3:
                    message.updatedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.updatedAt);
                    break;
                case /* optional resources.timestamp.Timestamp deleted_at */ 4:
                    message.deletedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.deletedAt);
                    break;
                case /* uint64 calendar_id */ 5:
                    message.calendarId = reader.uint64().toNumber();
                    break;
                case /* optional string job */ 6:
                    message.job = reader.string();
                    break;
                case /* resources.timestamp.Timestamp start_time */ 7:
                    message.startTime = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.startTime);
                    break;
                case /* optional resources.timestamp.Timestamp end_time */ 8:
                    message.endTime = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.endTime);
                    break;
                case /* string title */ 9:
                    message.title = reader.string();
                    break;
                case /* optional string color */ 10:
                    message.color = reader.string();
                    break;
                case /* string content */ 11:
                    message.content = reader.string();
                    break;
                case /* bool public */ 12:
                    message.public = reader.bool();
                    break;
                case /* optional int32 creator_id */ 13:
                    message.creatorId = reader.int32();
                    break;
                case /* optional resources.users.UserShort creator */ 14:
                    message.creator = UserShort.internalBinaryRead(reader, reader.uint32(), options, message.creator);
                    break;
                case /* string creator_job */ 15:
                    message.creatorJob = reader.string();
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
    internalBinaryWrite(message: CalendarEntry, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 id = 1 [jstype = JS_STRING]; */
        if (message.id !== "0")
            writer.tag(1, WireType.Varint).uint64(message.id);
        /* optional resources.timestamp.Timestamp created_at = 2; */
        if (message.createdAt)
            Timestamp.internalBinaryWrite(message.createdAt, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.timestamp.Timestamp updated_at = 3; */
        if (message.updatedAt)
            Timestamp.internalBinaryWrite(message.updatedAt, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.timestamp.Timestamp deleted_at = 4; */
        if (message.deletedAt)
            Timestamp.internalBinaryWrite(message.deletedAt, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* uint64 calendar_id = 5; */
        if (message.calendarId !== 0)
            writer.tag(5, WireType.Varint).uint64(message.calendarId);
        /* optional string job = 6; */
        if (message.job !== undefined)
            writer.tag(6, WireType.LengthDelimited).string(message.job);
        /* resources.timestamp.Timestamp start_time = 7; */
        if (message.startTime)
            Timestamp.internalBinaryWrite(message.startTime, writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.timestamp.Timestamp end_time = 8; */
        if (message.endTime)
            Timestamp.internalBinaryWrite(message.endTime, writer.tag(8, WireType.LengthDelimited).fork(), options).join();
        /* string title = 9; */
        if (message.title !== "")
            writer.tag(9, WireType.LengthDelimited).string(message.title);
        /* optional string color = 10; */
        if (message.color !== undefined)
            writer.tag(10, WireType.LengthDelimited).string(message.color);
        /* string content = 11; */
        if (message.content !== "")
            writer.tag(11, WireType.LengthDelimited).string(message.content);
        /* bool public = 12; */
        if (message.public !== false)
            writer.tag(12, WireType.Varint).bool(message.public);
        /* optional int32 creator_id = 13; */
        if (message.creatorId !== undefined)
            writer.tag(13, WireType.Varint).int32(message.creatorId);
        /* optional resources.users.UserShort creator = 14; */
        if (message.creator)
            UserShort.internalBinaryWrite(message.creator, writer.tag(14, WireType.LengthDelimited).fork(), options).join();
        /* string creator_job = 15; */
        if (message.creatorJob !== "")
            writer.tag(15, WireType.LengthDelimited).string(message.creatorJob);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.calendar.CalendarEntry
 */
export const CalendarEntry = new CalendarEntry$Type();
