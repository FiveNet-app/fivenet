// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size,long_type_number
// @generated from protobuf file "resources/permissions/permissions.proto" (package "resources.permissions", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { Timestamp } from "../timestamp/timestamp.js";
/**
 * @generated from protobuf message resources.permissions.Permission
 */
export interface Permission {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: number; // @gotags: alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp; // @gotags: alias:"created_at"
    /**
     * @generated from protobuf field: string category = 3;
     */
    category: string; // @gotags: alias:"category"
    /**
     * @generated from protobuf field: string name = 4;
     */
    name: string; // @gotags: alias:"name"
    /**
     * @generated from protobuf field: string guard_name = 5;
     */
    guardName: string; // @gotags: alias:"guard_name"
    /**
     * @generated from protobuf field: bool val = 6;
     */
    val: boolean; // @gotags: alias:"val"
}
/**
 * @generated from protobuf message resources.permissions.Role
 */
export interface Role {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: number; // @gotags: alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp; // @gotags: alias:"created_at"
    /**
     * @generated from protobuf field: string job = 3;
     */
    job: string; // @gotags: alias:"job"
    /**
     * @generated from protobuf field: string job_label = 4;
     */
    jobLabel: string; // @gotags: alias:"job_label"
    /**
     * @generated from protobuf field: int32 grade = 5;
     */
    grade: number; // @gotags: alias:"grade"
    /**
     * @generated from protobuf field: string job_grade_label = 6;
     */
    jobGradeLabel: string; // @gotags: alias:"job_grade_label"
    /**
     * @generated from protobuf field: repeated resources.permissions.Permission permissions = 7;
     */
    permissions: Permission[];
    /**
     * @generated from protobuf field: repeated resources.permissions.RoleAttribute attributes = 8;
     */
    attributes: RoleAttribute[];
}
/**
 * @generated from protobuf message resources.permissions.RawAttribute
 */
export interface RawAttribute {
    /**
     * @generated from protobuf field: uint64 role_id = 1;
     */
    roleId: number; // @gotags: alias:"role_id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp; // @gotags: alias:"created_at"
    /**
     * @generated from protobuf field: uint64 attr_id = 3;
     */
    attrId: number; // @gotags: alias:"attr_id"
    /**
     * @generated from protobuf field: uint64 permission_id = 4;
     */
    permissionId: number; // @gotags: alias:"permission_id"
    /**
     * @generated from protobuf field: string category = 5;
     */
    category: string; // @gotags: alias:"category"
    /**
     * @generated from protobuf field: string name = 6;
     */
    name: string; // @gotags: alias:"name"
    /**
     * @generated from protobuf field: string key = 7;
     */
    key: string; // @gotags: alias:"key"
    /**
     * @generated from protobuf field: string type = 8;
     */
    type: string; // @gotags: alias:"type"
    /**
     * @generated from protobuf field: string raw_value = 9;
     */
    rawValue: string; // @gotags: alias:"value"
    /**
     * @generated from protobuf field: string raw_valid_values = 10;
     */
    rawValidValues: string; // @gotags: alias:"valid_values"
}
/**
 * @generated from protobuf message resources.permissions.RoleAttribute
 */
export interface RoleAttribute {
    /**
     * @generated from protobuf field: uint64 role_id = 1;
     */
    roleId: number; // @gotags: alias:"role_id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp; // @gotags: alias:"created_at"
    /**
     * @generated from protobuf field: uint64 attr_id = 3;
     */
    attrId: number; // @gotags: alias:"attr_id"
    /**
     * @generated from protobuf field: uint64 permission_id = 4;
     */
    permissionId: number; // @gotags: alias:"permission_id"
    /**
     * @generated from protobuf field: string category = 5;
     */
    category: string; // @gotags: alias:"category"
    /**
     * @generated from protobuf field: string name = 6;
     */
    name: string; // @gotags: alias:"name"
    /**
     * @generated from protobuf field: string key = 7;
     */
    key: string; // @gotags: alias:"key"
    /**
     * @generated from protobuf field: string type = 8;
     */
    type: string; // @gotags: alias:"type"
    /**
     * @generated from protobuf field: resources.permissions.AttributeValues value = 11;
     */
    value?: AttributeValues;
    /**
     * @generated from protobuf field: resources.permissions.AttributeValues valid_values = 12;
     */
    validValues?: AttributeValues;
}
/**
 * @generated from protobuf message resources.permissions.AttributeValues
 */
export interface AttributeValues {
    /**
     * @generated from protobuf oneof: valid_values
     */
    validValues: {
        oneofKind: "stringList";
        /**
         * @generated from protobuf field: resources.permissions.StringList string_list = 1;
         */
        stringList: StringList;
    } | {
        oneofKind: "jobList";
        /**
         * @generated from protobuf field: resources.permissions.StringList job_list = 2;
         */
        jobList: StringList;
    } | {
        oneofKind: "jobGradeList";
        /**
         * @generated from protobuf field: resources.permissions.JobGradeList job_grade_list = 3;
         */
        jobGradeList: JobGradeList;
    } | {
        oneofKind: undefined;
    };
}
/**
 * @generated from protobuf message resources.permissions.StringList
 */
export interface StringList {
    /**
     * @generated from protobuf field: repeated string strings = 1;
     */
    strings: string[];
}
/**
 * @generated from protobuf message resources.permissions.JobGradeList
 */
export interface JobGradeList {
    /**
     * @generated from protobuf field: map<string, int32> jobs = 1;
     */
    jobs: {
        [key: string]: number;
    };
}
// @generated message type with reflection information, may provide speed optimized methods
class Permission$Type extends MessageType<Permission> {
    constructor() {
        super("resources.permissions.Permission", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 2 /*LongType.NUMBER*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "category", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "128" } } } },
            { no: 4, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 5, name: "guard_name", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 6, name: "val", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.permissions.Permission
 */
export const Permission = new Permission$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Role$Type extends MessageType<Role> {
    constructor() {
        super("resources.permissions.Role", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 2 /*LongType.NUMBER*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 4, name: "job_label", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 5, name: "grade", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 6, name: "job_grade_label", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 7, name: "permissions", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Permission },
            { no: 8, name: "attributes", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => RoleAttribute }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.permissions.Role
 */
export const Role = new Role$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RawAttribute$Type extends MessageType<RawAttribute> {
    constructor() {
        super("resources.permissions.RawAttribute", [
            { no: 1, name: "role_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 2 /*LongType.NUMBER*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "attr_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 2 /*LongType.NUMBER*/ },
            { no: 4, name: "permission_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 2 /*LongType.NUMBER*/ },
            { no: 5, name: "category", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "128" } } } },
            { no: 6, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 7, name: "key", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 8, name: "type", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 9, name: "raw_value", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 10, name: "raw_valid_values", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.permissions.RawAttribute
 */
export const RawAttribute = new RawAttribute$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RoleAttribute$Type extends MessageType<RoleAttribute> {
    constructor() {
        super("resources.permissions.RoleAttribute", [
            { no: 1, name: "role_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 2 /*LongType.NUMBER*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "attr_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 2 /*LongType.NUMBER*/ },
            { no: 4, name: "permission_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 2 /*LongType.NUMBER*/ },
            { no: 5, name: "category", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "128" } } } },
            { no: 6, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 7, name: "key", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 8, name: "type", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 11, name: "value", kind: "message", T: () => AttributeValues },
            { no: 12, name: "valid_values", kind: "message", T: () => AttributeValues }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.permissions.RoleAttribute
 */
export const RoleAttribute = new RoleAttribute$Type();
// @generated message type with reflection information, may provide speed optimized methods
class AttributeValues$Type extends MessageType<AttributeValues> {
    constructor() {
        super("resources.permissions.AttributeValues", [
            { no: 1, name: "string_list", kind: "message", oneof: "validValues", T: () => StringList },
            { no: 2, name: "job_list", kind: "message", oneof: "validValues", T: () => StringList },
            { no: 3, name: "job_grade_list", kind: "message", oneof: "validValues", T: () => JobGradeList }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.permissions.AttributeValues
 */
export const AttributeValues = new AttributeValues$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StringList$Type extends MessageType<StringList> {
    constructor() {
        super("resources.permissions.StringList", [
            { no: 1, name: "strings", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.permissions.StringList
 */
export const StringList = new StringList$Type();
// @generated message type with reflection information, may provide speed optimized methods
class JobGradeList$Type extends MessageType<JobGradeList> {
    constructor() {
        super("resources.permissions.JobGradeList", [
            { no: 1, name: "jobs", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 5 /*ScalarType.INT32*/ } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.permissions.JobGradeList
 */
export const JobGradeList = new JobGradeList$Type();
