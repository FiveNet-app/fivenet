// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size,long_type_number
// @generated from protobuf file "resources/users/users.proto" (package "resources.users", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { Timestamp } from "../timestamp/timestamp.js";
import { Job } from "../jobs/jobs.js";
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
    identifier: string; // @gotags: alias:"identifier"
    /**
     * @generated from protobuf field: string job = 3;
     */
    job: string; // @gotags: alias:"job"
    /**
     * @generated from protobuf field: string job_label = 4;
     */
    jobLabel: string; // @gotags: alias:"job_label"
    /**
     * @generated from protobuf field: int32 job_grade = 5;
     */
    jobGrade: number; // @gotags: alias:"job_grade"
    /**
     * @generated from protobuf field: string job_grade_label = 6;
     */
    jobGradeLabel: string; // @gotags: alias:"job_grade_label"
    /**
     * @generated from protobuf field: string firstname = 7;
     */
    firstname: string; // @gotags: alias:"firstname"
    /**
     * @generated from protobuf field: string lastname = 8;
     */
    lastname: string; // @gotags: alias:"lastname"
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
    identifier: string; // @gotags: alias:"identifier"
    /**
     * @generated from protobuf field: string job = 3;
     */
    job: string; // @gotags: alias:"job"
    /**
     * @generated from protobuf field: string job_label = 4;
     */
    jobLabel: string; // @gotags: alias:"job_label"
    /**
     * @generated from protobuf field: int32 job_grade = 5;
     */
    jobGrade: number; // @gotags: alias:"job_grade"
    /**
     * @generated from protobuf field: string job_grade_label = 6;
     */
    jobGradeLabel: string; // @gotags: alias:"job_grade_label"
    /**
     * @generated from protobuf field: string firstname = 7;
     */
    firstname: string; // @gotags: alias:"firstname"
    /**
     * @generated from protobuf field: string lastname = 8;
     */
    lastname: string; // @gotags: alias:"lastname"
    /**
     * @generated from protobuf field: string dateofbirth = 9;
     */
    dateofbirth: string; // @gotags: alias:"dateofbirth"
    /**
     * @generated from protobuf field: optional string sex = 10;
     */
    sex?: string; // @gotags: alias:"sex"
    /**
     * @generated from protobuf field: optional string height = 11;
     */
    height?: string; // @gotags: alias:"height"
    /**
     * @generated from protobuf field: optional string phone_number = 12;
     */
    phoneNumber?: string; // @gotags: alias:"phone_number"
    /**
     * @generated from protobuf field: optional int32 visum = 13;
     */
    visum?: number; // @gotags: alias:"visum"
    /**
     * @generated from protobuf field: optional int32 playtime = 14;
     */
    playtime?: number; // @gotags: alias:"playtime"
    /**
     * @generated from protobuf field: resources.users.UserProps props = 15;
     */
    props?: UserProps; // @gotags: alias:"fivenet_user_props"
    /**
     * @generated from protobuf field: repeated resources.users.License licenses = 16;
     */
    licenses: License[]; // @gotags: alias:"user_licenses"
}
/**
 * @generated from protobuf message resources.users.License
 */
export interface License {
    /**
     * @generated from protobuf field: string type = 1;
     */
    type: string; // @gotags: alias:"type"
    /**
     * @generated from protobuf field: string label = 2;
     */
    label: string; // @gotags: alias:"label"
}
/**
 * @generated from protobuf message resources.users.UserProps
 */
export interface UserProps {
    /**
     * @generated from protobuf field: int32 user_id = 1;
     */
    userId: number; // @gotags: alias:"user_id"
    /**
     * @generated from protobuf field: optional bool wanted = 2;
     */
    wanted?: boolean; // @gotags: alias:"wanted"
    /**
     * @generated from protobuf field: optional string job_name = 3;
     */
    jobName?: string; // @gotags: alias:"job"
    /**
     * @generated from protobuf field: optional resources.jobs.Job job = 4;
     */
    job?: Job;
}
/**
 * @generated from protobuf message resources.users.UserActivity
 */
export interface UserActivity {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: number; // @gotags: alias:"fivenet_user_activity.id"
    /**
     * @generated from protobuf field: resources.users.USER_ACTIVITY_TYPE type = 2;
     */
    type: USER_ACTIVITY_TYPE; // @gotags: alias:"fivenet_user_activity.type"
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
     * @generated from protobuf field: string reason = 9;
     */
    reason: string; // @gotags: alias:"fivenet_user_activity.reason"
}
/**
 * @generated from protobuf enum resources.users.USER_ACTIVITY_TYPE
 */
export enum USER_ACTIVITY_TYPE {
    /**
     * @generated from protobuf enum value: CHANGED = 0;
     */
    CHANGED = 0,
    /**
     * @generated from protobuf enum value: MENTIONED = 1;
     */
    MENTIONED = 1,
    /**
     * @generated from protobuf enum value: CREATED = 2;
     */
    CREATED = 2
}
// @generated message type with reflection information, may provide speed optimized methods
class UserShort$Type extends MessageType<UserShort> {
    constructor() {
        super("resources.users.UserShort", [
            { no: 1, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 2, name: "identifier", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { len: "46" } } } },
            { no: 3, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 4, name: "job_label", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 5, name: "job_grade", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: -1 } } } },
            { no: 6, name: "job_grade_label", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 7, name: "firstname", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "1", maxLen: "50" } } } },
            { no: 8, name: "lastname", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "1", maxLen: "50" } } } }
        ]);
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
            { no: 3, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 4, name: "job_label", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 5, name: "job_grade", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: -1 } } } },
            { no: 6, name: "job_grade_label", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 7, name: "firstname", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "1", maxLen: "50" } } } },
            { no: 8, name: "lastname", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "1", maxLen: "50" } } } },
            { no: 9, name: "dateofbirth", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { len: "10" } } } },
            { no: 10, name: "sex", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "1", maxLen: "2" } } } },
            { no: 11, name: "height", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 12, name: "phone_number", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 13, name: "visum", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gte: 0 } } } },
            { no: 14, name: "playtime", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gte: 0 } } } },
            { no: 15, name: "props", kind: "message", T: () => UserProps },
            { no: 16, name: "licenses", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => License }
        ]);
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
            { no: 4, name: "job", kind: "message", T: () => Job }
        ]);
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
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 2 /*LongType.NUMBER*/ },
            { no: 2, name: "type", kind: "enum", T: () => ["resources.users.USER_ACTIVITY_TYPE", USER_ACTIVITY_TYPE], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 3, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "source_user", kind: "message", T: () => UserShort },
            { no: 5, name: "target_user", kind: "message", T: () => UserShort },
            { no: 6, name: "key", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "1", maxLen: "64" } } } },
            { no: 7, name: "old_value", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "256" } } } },
            { no: 8, name: "new_value", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "256" } } } },
            { no: 9, name: "reason", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.users.UserActivity
 */
export const UserActivity = new UserActivity$Type();
