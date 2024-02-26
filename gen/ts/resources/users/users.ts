// @generated by protobuf-ts 2.9.3 with parameter optimize_speed,long_type_bigint
// @generated from protobuf file "resources/users/users.proto" (package "resources.users", syntax proto3)
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
import { Timestamp } from "../timestamp/timestamp";
import { JobGrade } from "./jobs";
import { Job } from "./jobs";
import { File } from "../filestore/file";
/**
 * @generated from protobuf message resources.users.UserShort
 */
export interface UserShort {
    /**
     * @generated from protobuf field: int32 user_id = 1;
     */
    userId: number; // @gotags: alias:"id"
    /**
     * @generated from protobuf field: string identifier = 2;
     */
    identifier: string;
    /**
     * @generated from protobuf field: string job = 3;
     */
    job: string;
    /**
     * @generated from protobuf field: optional string job_label = 4;
     */
    jobLabel?: string;
    /**
     * @generated from protobuf field: int32 job_grade = 5;
     */
    jobGrade: number;
    /**
     * @generated from protobuf field: optional string job_grade_label = 6;
     */
    jobGradeLabel?: string;
    /**
     * @generated from protobuf field: string firstname = 7;
     */
    firstname: string;
    /**
     * @generated from protobuf field: string lastname = 8;
     */
    lastname: string;
    /**
     * @generated from protobuf field: string dateofbirth = 9;
     */
    dateofbirth: string;
    /**
     * @generated from protobuf field: optional string phone_number = 12;
     */
    phoneNumber?: string;
    /**
     * @generated from protobuf field: optional resources.filestore.File avatar_url = 17;
     */
    avatarUrl?: File;
}
/**
 * @generated from protobuf message resources.users.User
 */
export interface User {
    /**
     * @generated from protobuf field: int32 user_id = 1;
     */
    userId: number; // @gotags: alias:"id"
    /**
     * @generated from protobuf field: string identifier = 2;
     */
    identifier: string;
    /**
     * @generated from protobuf field: string job = 3;
     */
    job: string;
    /**
     * @generated from protobuf field: optional string job_label = 4;
     */
    jobLabel?: string;
    /**
     * @generated from protobuf field: int32 job_grade = 5;
     */
    jobGrade: number;
    /**
     * @generated from protobuf field: optional string job_grade_label = 6;
     */
    jobGradeLabel?: string;
    /**
     * @generated from protobuf field: string firstname = 7;
     */
    firstname: string;
    /**
     * @generated from protobuf field: string lastname = 8;
     */
    lastname: string;
    /**
     * @generated from protobuf field: string dateofbirth = 9;
     */
    dateofbirth: string;
    /**
     * @generated from protobuf field: optional string sex = 10;
     */
    sex?: string;
    /**
     * @generated from protobuf field: optional string height = 11;
     */
    height?: string;
    /**
     * @generated from protobuf field: optional string phone_number = 12;
     */
    phoneNumber?: string;
    /**
     * @generated from protobuf field: optional int32 visum = 13;
     */
    visum?: number;
    /**
     * @generated from protobuf field: optional int32 playtime = 14;
     */
    playtime?: number;
    /**
     * @generated from protobuf field: resources.users.UserProps props = 15;
     */
    props?: UserProps; // @gotags: alias:"fivenet_user_props"
    /**
     * @generated from protobuf field: repeated resources.users.License licenses = 16;
     */
    licenses: License[]; // @gotags: alias:"user_licenses"
    /**
     * @generated from protobuf field: optional resources.filestore.File avatar_url = 17;
     */
    avatarUrl?: File;
}
/**
 * @generated from protobuf message resources.users.License
 */
export interface License {
    /**
     * @generated from protobuf field: string type = 1;
     */
    type: string;
    /**
     * @generated from protobuf field: string label = 2;
     */
    label: string;
}
/**
 * @generated from protobuf message resources.users.UserProps
 */
export interface UserProps {
    /**
     * @generated from protobuf field: int32 user_id = 1;
     */
    userId: number;
    /**
     * @generated from protobuf field: optional bool wanted = 2;
     */
    wanted?: boolean;
    /**
     * @generated from protobuf field: optional string job_name = 3;
     */
    jobName?: string; // @gotags: alias:"job"
    /**
     * @generated from protobuf field: optional resources.users.Job job = 4;
     */
    job?: Job;
    /**
     * @generated from protobuf field: optional int32 job_grade_number = 5;
     */
    jobGradeNumber?: number; // @gotags: alias:"job_grade"
    /**
     * @generated from protobuf field: optional resources.users.JobGrade job_grade = 6;
     */
    jobGrade?: JobGrade;
    /**
     * @generated from protobuf field: optional uint32 traffic_infraction_points = 7;
     */
    trafficInfractionPoints?: number;
    /**
     * @generated from protobuf field: optional int64 open_fines = 8;
     */
    openFines?: bigint;
    /**
     * @generated from protobuf field: optional string blood_type = 9;
     */
    bloodType?: string;
}
/**
 * @generated from protobuf message resources.users.UserActivity
 */
export interface UserActivity {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string; // @gotags: alias:"fivenet_user_activity.id"
    /**
     * @generated from protobuf field: resources.users.UserActivityType type = 2;
     */
    type: UserActivityType; // @gotags: alias:"fivenet_user_activity.type"
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp created_at = 3;
     */
    createdAt?: Timestamp; // @gotags: alias:"fivenet_user_activity.created_at"
    /**
     * @generated from protobuf field: resources.users.UserShort source_user = 4;
     */
    sourceUser?: UserShort; // @gotags: alias:"source_user"
    /**
     * @generated from protobuf field: resources.users.UserShort target_user = 5;
     */
    targetUser?: UserShort; // @gotags: alias:"target_user"
    /**
     * @sanitize
     *
     * @generated from protobuf field: string key = 6;
     */
    key: string; // @gotags: alias:"fivenet_user_activity.key"
    /**
     * @generated from protobuf field: string old_value = 7;
     */
    oldValue: string; // @gotags: alias:"fivenet_user_activity.old_value"
    /**
     * @generated from protobuf field: string new_value = 8;
     */
    newValue: string; // @gotags: alias:"fivenet_user_activity.new_value"
    /**
     * @sanitize
     *
     * @generated from protobuf field: string reason = 9;
     */
    reason: string; // @gotags: alias:"fivenet_user_activity.reason"
}
/**
 * @generated from protobuf enum resources.users.UserActivityType
 */
export enum UserActivityType {
    /**
     * @generated from protobuf enum value: USER_ACTIVITY_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: USER_ACTIVITY_TYPE_CHANGED = 1;
     */
    CHANGED = 1,
    /**
     * @generated from protobuf enum value: USER_ACTIVITY_TYPE_MENTIONED = 2;
     */
    MENTIONED = 2,
    /**
     * @generated from protobuf enum value: USER_ACTIVITY_TYPE_CREATED = 3;
     */
    CREATED = 3
}
// @generated message type with reflection information, may provide speed optimized methods
class UserShort$Type extends MessageType<UserShort> {
    constructor() {
        super("resources.users.UserShort", [
            { no: 1, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 2, name: "identifier", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { len: "46" } } } },
            { no: 3, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 4, name: "job_label", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 5, name: "job_grade", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: -1 } } } },
            { no: 6, name: "job_grade_label", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 7, name: "firstname", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "1", maxLen: "50" } } } },
            { no: 8, name: "lastname", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "1", maxLen: "50" } } } },
            { no: 9, name: "dateofbirth", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { len: "10" } } } },
            { no: 12, name: "phone_number", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 17, name: "avatar_url", kind: "message", T: () => File }
        ]);
    }
    create(value?: PartialMessage<UserShort>): UserShort {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.userId = 0;
        message.identifier = "";
        message.job = "";
        message.jobGrade = 0;
        message.firstname = "";
        message.lastname = "";
        message.dateofbirth = "";
        if (value !== undefined)
            reflectionMergePartial<UserShort>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UserShort): UserShort {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 user_id */ 1:
                    message.userId = reader.int32();
                    break;
                case /* string identifier */ 2:
                    message.identifier = reader.string();
                    break;
                case /* string job */ 3:
                    message.job = reader.string();
                    break;
                case /* optional string job_label */ 4:
                    message.jobLabel = reader.string();
                    break;
                case /* int32 job_grade */ 5:
                    message.jobGrade = reader.int32();
                    break;
                case /* optional string job_grade_label */ 6:
                    message.jobGradeLabel = reader.string();
                    break;
                case /* string firstname */ 7:
                    message.firstname = reader.string();
                    break;
                case /* string lastname */ 8:
                    message.lastname = reader.string();
                    break;
                case /* string dateofbirth */ 9:
                    message.dateofbirth = reader.string();
                    break;
                case /* optional string phone_number */ 12:
                    message.phoneNumber = reader.string();
                    break;
                case /* optional resources.filestore.File avatar_url */ 17:
                    message.avatarUrl = File.internalBinaryRead(reader, reader.uint32(), options, message.avatarUrl);
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
    internalBinaryWrite(message: UserShort, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 user_id = 1; */
        if (message.userId !== 0)
            writer.tag(1, WireType.Varint).int32(message.userId);
        /* string identifier = 2; */
        if (message.identifier !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.identifier);
        /* string job = 3; */
        if (message.job !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.job);
        /* optional string job_label = 4; */
        if (message.jobLabel !== undefined)
            writer.tag(4, WireType.LengthDelimited).string(message.jobLabel);
        /* int32 job_grade = 5; */
        if (message.jobGrade !== 0)
            writer.tag(5, WireType.Varint).int32(message.jobGrade);
        /* optional string job_grade_label = 6; */
        if (message.jobGradeLabel !== undefined)
            writer.tag(6, WireType.LengthDelimited).string(message.jobGradeLabel);
        /* string firstname = 7; */
        if (message.firstname !== "")
            writer.tag(7, WireType.LengthDelimited).string(message.firstname);
        /* string lastname = 8; */
        if (message.lastname !== "")
            writer.tag(8, WireType.LengthDelimited).string(message.lastname);
        /* string dateofbirth = 9; */
        if (message.dateofbirth !== "")
            writer.tag(9, WireType.LengthDelimited).string(message.dateofbirth);
        /* optional string phone_number = 12; */
        if (message.phoneNumber !== undefined)
            writer.tag(12, WireType.LengthDelimited).string(message.phoneNumber);
        /* optional resources.filestore.File avatar_url = 17; */
        if (message.avatarUrl)
            File.internalBinaryWrite(message.avatarUrl, writer.tag(17, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.users.UserShort
 */
export const UserShort = new UserShort$Type();
// @generated message type with reflection information, may provide speed optimized methods
class User$Type extends MessageType<User> {
    constructor() {
        super("resources.users.User", [
            { no: 1, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 2, name: "identifier", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { len: "46" } } } },
            { no: 3, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 4, name: "job_label", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 5, name: "job_grade", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: -1 } } } },
            { no: 6, name: "job_grade_label", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 7, name: "firstname", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "1", maxLen: "50" } } } },
            { no: 8, name: "lastname", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "1", maxLen: "50" } } } },
            { no: 9, name: "dateofbirth", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { len: "10" } } } },
            { no: 10, name: "sex", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "1", maxLen: "2" } } } },
            { no: 11, name: "height", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 12, name: "phone_number", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 13, name: "visum", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gte: 0 } } } },
            { no: 14, name: "playtime", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gte: 0 } } } },
            { no: 15, name: "props", kind: "message", T: () => UserProps },
            { no: 16, name: "licenses", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => License },
            { no: 17, name: "avatar_url", kind: "message", T: () => File }
        ]);
    }
    create(value?: PartialMessage<User>): User {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.userId = 0;
        message.identifier = "";
        message.job = "";
        message.jobGrade = 0;
        message.firstname = "";
        message.lastname = "";
        message.dateofbirth = "";
        message.licenses = [];
        if (value !== undefined)
            reflectionMergePartial<User>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: User): User {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 user_id */ 1:
                    message.userId = reader.int32();
                    break;
                case /* string identifier */ 2:
                    message.identifier = reader.string();
                    break;
                case /* string job */ 3:
                    message.job = reader.string();
                    break;
                case /* optional string job_label */ 4:
                    message.jobLabel = reader.string();
                    break;
                case /* int32 job_grade */ 5:
                    message.jobGrade = reader.int32();
                    break;
                case /* optional string job_grade_label */ 6:
                    message.jobGradeLabel = reader.string();
                    break;
                case /* string firstname */ 7:
                    message.firstname = reader.string();
                    break;
                case /* string lastname */ 8:
                    message.lastname = reader.string();
                    break;
                case /* string dateofbirth */ 9:
                    message.dateofbirth = reader.string();
                    break;
                case /* optional string sex */ 10:
                    message.sex = reader.string();
                    break;
                case /* optional string height */ 11:
                    message.height = reader.string();
                    break;
                case /* optional string phone_number */ 12:
                    message.phoneNumber = reader.string();
                    break;
                case /* optional int32 visum */ 13:
                    message.visum = reader.int32();
                    break;
                case /* optional int32 playtime */ 14:
                    message.playtime = reader.int32();
                    break;
                case /* resources.users.UserProps props */ 15:
                    message.props = UserProps.internalBinaryRead(reader, reader.uint32(), options, message.props);
                    break;
                case /* repeated resources.users.License licenses */ 16:
                    message.licenses.push(License.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* optional resources.filestore.File avatar_url */ 17:
                    message.avatarUrl = File.internalBinaryRead(reader, reader.uint32(), options, message.avatarUrl);
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
    internalBinaryWrite(message: User, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 user_id = 1; */
        if (message.userId !== 0)
            writer.tag(1, WireType.Varint).int32(message.userId);
        /* string identifier = 2; */
        if (message.identifier !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.identifier);
        /* string job = 3; */
        if (message.job !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.job);
        /* optional string job_label = 4; */
        if (message.jobLabel !== undefined)
            writer.tag(4, WireType.LengthDelimited).string(message.jobLabel);
        /* int32 job_grade = 5; */
        if (message.jobGrade !== 0)
            writer.tag(5, WireType.Varint).int32(message.jobGrade);
        /* optional string job_grade_label = 6; */
        if (message.jobGradeLabel !== undefined)
            writer.tag(6, WireType.LengthDelimited).string(message.jobGradeLabel);
        /* string firstname = 7; */
        if (message.firstname !== "")
            writer.tag(7, WireType.LengthDelimited).string(message.firstname);
        /* string lastname = 8; */
        if (message.lastname !== "")
            writer.tag(8, WireType.LengthDelimited).string(message.lastname);
        /* string dateofbirth = 9; */
        if (message.dateofbirth !== "")
            writer.tag(9, WireType.LengthDelimited).string(message.dateofbirth);
        /* optional string sex = 10; */
        if (message.sex !== undefined)
            writer.tag(10, WireType.LengthDelimited).string(message.sex);
        /* optional string height = 11; */
        if (message.height !== undefined)
            writer.tag(11, WireType.LengthDelimited).string(message.height);
        /* optional string phone_number = 12; */
        if (message.phoneNumber !== undefined)
            writer.tag(12, WireType.LengthDelimited).string(message.phoneNumber);
        /* optional int32 visum = 13; */
        if (message.visum !== undefined)
            writer.tag(13, WireType.Varint).int32(message.visum);
        /* optional int32 playtime = 14; */
        if (message.playtime !== undefined)
            writer.tag(14, WireType.Varint).int32(message.playtime);
        /* resources.users.UserProps props = 15; */
        if (message.props)
            UserProps.internalBinaryWrite(message.props, writer.tag(15, WireType.LengthDelimited).fork(), options).join();
        /* repeated resources.users.License licenses = 16; */
        for (let i = 0; i < message.licenses.length; i++)
            License.internalBinaryWrite(message.licenses[i], writer.tag(16, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.filestore.File avatar_url = 17; */
        if (message.avatarUrl)
            File.internalBinaryWrite(message.avatarUrl, writer.tag(17, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.users.User
 */
export const User = new User$Type();
// @generated message type with reflection information, may provide speed optimized methods
class License$Type extends MessageType<License> {
    constructor() {
        super("resources.users.License", [
            { no: 1, name: "type", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3", maxLen: "60" } } } },
            { no: 2, name: "label", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<License>): License {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.type = "";
        message.label = "";
        if (value !== undefined)
            reflectionMergePartial<License>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: License): License {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string type */ 1:
                    message.type = reader.string();
                    break;
                case /* string label */ 2:
                    message.label = reader.string();
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
    internalBinaryWrite(message: License, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string type = 1; */
        if (message.type !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.type);
        /* string label = 2; */
        if (message.label !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.label);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.users.License
 */
export const License = new License$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UserProps$Type extends MessageType<UserProps> {
    constructor() {
        super("resources.users.UserProps", [
            { no: 1, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 2, name: "wanted", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "job_name", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "job", kind: "message", T: () => Job },
            { no: 5, name: "job_grade_number", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "job_grade", kind: "message", T: () => JobGrade },
            { no: 7, name: "traffic_infraction_points", kind: "scalar", opt: true, T: 13 /*ScalarType.UINT32*/ },
            { no: 8, name: "open_fines", kind: "scalar", opt: true, T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 9, name: "blood_type", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<UserProps>): UserProps {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.userId = 0;
        if (value !== undefined)
            reflectionMergePartial<UserProps>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UserProps): UserProps {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 user_id */ 1:
                    message.userId = reader.int32();
                    break;
                case /* optional bool wanted */ 2:
                    message.wanted = reader.bool();
                    break;
                case /* optional string job_name */ 3:
                    message.jobName = reader.string();
                    break;
                case /* optional resources.users.Job job */ 4:
                    message.job = Job.internalBinaryRead(reader, reader.uint32(), options, message.job);
                    break;
                case /* optional int32 job_grade_number */ 5:
                    message.jobGradeNumber = reader.int32();
                    break;
                case /* optional resources.users.JobGrade job_grade */ 6:
                    message.jobGrade = JobGrade.internalBinaryRead(reader, reader.uint32(), options, message.jobGrade);
                    break;
                case /* optional uint32 traffic_infraction_points */ 7:
                    message.trafficInfractionPoints = reader.uint32();
                    break;
                case /* optional int64 open_fines */ 8:
                    message.openFines = reader.int64().toBigInt();
                    break;
                case /* optional string blood_type */ 9:
                    message.bloodType = reader.string();
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
    internalBinaryWrite(message: UserProps, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 user_id = 1; */
        if (message.userId !== 0)
            writer.tag(1, WireType.Varint).int32(message.userId);
        /* optional bool wanted = 2; */
        if (message.wanted !== undefined)
            writer.tag(2, WireType.Varint).bool(message.wanted);
        /* optional string job_name = 3; */
        if (message.jobName !== undefined)
            writer.tag(3, WireType.LengthDelimited).string(message.jobName);
        /* optional resources.users.Job job = 4; */
        if (message.job)
            Job.internalBinaryWrite(message.job, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* optional int32 job_grade_number = 5; */
        if (message.jobGradeNumber !== undefined)
            writer.tag(5, WireType.Varint).int32(message.jobGradeNumber);
        /* optional resources.users.JobGrade job_grade = 6; */
        if (message.jobGrade)
            JobGrade.internalBinaryWrite(message.jobGrade, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        /* optional uint32 traffic_infraction_points = 7; */
        if (message.trafficInfractionPoints !== undefined)
            writer.tag(7, WireType.Varint).uint32(message.trafficInfractionPoints);
        /* optional int64 open_fines = 8; */
        if (message.openFines !== undefined)
            writer.tag(8, WireType.Varint).int64(message.openFines);
        /* optional string blood_type = 9; */
        if (message.bloodType !== undefined)
            writer.tag(9, WireType.LengthDelimited).string(message.bloodType);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.users.UserProps
 */
export const UserProps = new UserProps$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UserActivity$Type extends MessageType<UserActivity> {
    constructor() {
        super("resources.users.UserActivity", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "type", kind: "enum", T: () => ["resources.users.UserActivityType", UserActivityType, "USER_ACTIVITY_TYPE_"], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 3, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "source_user", kind: "message", T: () => UserShort },
            { no: 5, name: "target_user", kind: "message", T: () => UserShort },
            { no: 6, name: "key", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "1", maxLen: "64" } } } },
            { no: 7, name: "old_value", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "256" } } } },
            { no: 8, name: "new_value", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "256" } } } },
            { no: 9, name: "reason", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<UserActivity>): UserActivity {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.id = "0";
        message.type = 0;
        message.key = "";
        message.oldValue = "";
        message.newValue = "";
        message.reason = "";
        if (value !== undefined)
            reflectionMergePartial<UserActivity>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UserActivity): UserActivity {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 id = 1 [jstype = JS_STRING];*/ 1:
                    message.id = reader.uint64().toString();
                    break;
                case /* resources.users.UserActivityType type */ 2:
                    message.type = reader.int32();
                    break;
                case /* resources.timestamp.Timestamp created_at */ 3:
                    message.createdAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.createdAt);
                    break;
                case /* resources.users.UserShort source_user */ 4:
                    message.sourceUser = UserShort.internalBinaryRead(reader, reader.uint32(), options, message.sourceUser);
                    break;
                case /* resources.users.UserShort target_user */ 5:
                    message.targetUser = UserShort.internalBinaryRead(reader, reader.uint32(), options, message.targetUser);
                    break;
                case /* string key */ 6:
                    message.key = reader.string();
                    break;
                case /* string old_value */ 7:
                    message.oldValue = reader.string();
                    break;
                case /* string new_value */ 8:
                    message.newValue = reader.string();
                    break;
                case /* string reason */ 9:
                    message.reason = reader.string();
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
    internalBinaryWrite(message: UserActivity, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 id = 1 [jstype = JS_STRING]; */
        if (message.id !== "0")
            writer.tag(1, WireType.Varint).uint64(message.id);
        /* resources.users.UserActivityType type = 2; */
        if (message.type !== 0)
            writer.tag(2, WireType.Varint).int32(message.type);
        /* resources.timestamp.Timestamp created_at = 3; */
        if (message.createdAt)
            Timestamp.internalBinaryWrite(message.createdAt, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* resources.users.UserShort source_user = 4; */
        if (message.sourceUser)
            UserShort.internalBinaryWrite(message.sourceUser, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* resources.users.UserShort target_user = 5; */
        if (message.targetUser)
            UserShort.internalBinaryWrite(message.targetUser, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* string key = 6; */
        if (message.key !== "")
            writer.tag(6, WireType.LengthDelimited).string(message.key);
        /* string old_value = 7; */
        if (message.oldValue !== "")
            writer.tag(7, WireType.LengthDelimited).string(message.oldValue);
        /* string new_value = 8; */
        if (message.newValue !== "")
            writer.tag(8, WireType.LengthDelimited).string(message.newValue);
        /* string reason = 9; */
        if (message.reason !== "")
            writer.tag(9, WireType.LengthDelimited).string(message.reason);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.users.UserActivity
 */
export const UserActivity = new UserActivity$Type();
