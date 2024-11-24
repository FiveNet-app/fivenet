// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "resources/common/access/dummy.proto" (package "resources.common.access", syntax proto3)
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
/**
 * @generated from protobuf message resources.common.access.DummyJobAccess
 */
export interface DummyJobAccess {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
    /**
     * @generated from protobuf field: uint64 target_id = 2 [jstype = JS_STRING];
     */
    targetId: string;
    /**
     * @generated from protobuf field: string job = 3;
     */
    job: string;
    /**
     * @generated from protobuf field: int32 minimum_grade = 4;
     */
    minimumGrade: number;
    /**
     * @generated from protobuf field: resources.common.access.AccessLevel access = 5;
     */
    access: AccessLevel;
}
/**
 * @generated from protobuf message resources.common.access.DummyUserAccess
 */
export interface DummyUserAccess {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
    /**
     * @generated from protobuf field: uint64 target_id = 2 [jstype = JS_STRING];
     */
    targetId: string;
    /**
     * @generated from protobuf field: int32 user_id = 3;
     */
    userId: number;
    /**
     * @generated from protobuf field: resources.common.access.AccessLevel access = 4;
     */
    access: AccessLevel;
}
/**
 * @generated from protobuf message resources.common.access.DummyQualificationAccess
 */
export interface DummyQualificationAccess {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
    /**
     * @generated from protobuf field: uint64 target_id = 2 [jstype = JS_STRING];
     */
    targetId: string;
    /**
     * @generated from protobuf field: uint64 qualification_id = 3 [jstype = JS_STRING];
     */
    qualificationId: string;
    /**
     * @generated from protobuf field: resources.common.access.AccessLevel access = 4;
     */
    access: AccessLevel;
}
/**
 * @generated from protobuf enum resources.common.access.AccessLevel
 */
export enum AccessLevel {
    /**
     * @generated from protobuf enum value: ACCESS_LEVEL_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0
}
// @generated message type with reflection information, may provide speed optimized methods
class DummyJobAccess$Type extends MessageType<DummyJobAccess> {
    constructor() {
        super("resources.common.access.DummyJobAccess", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "target_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 3, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 4, name: "minimum_grade", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gte: 0 } } } },
            { no: 5, name: "access", kind: "enum", T: () => ["resources.common.access.AccessLevel", AccessLevel, "ACCESS_LEVEL_"], options: { "validate.rules": { enum: { definedOnly: true } } } }
        ]);
    }
    create(value?: PartialMessage<DummyJobAccess>): DummyJobAccess {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.id = "0";
        message.targetId = "0";
        message.job = "";
        message.minimumGrade = 0;
        message.access = 0;
        if (value !== undefined)
            reflectionMergePartial<DummyJobAccess>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DummyJobAccess): DummyJobAccess {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 id = 1 [jstype = JS_STRING];*/ 1:
                    message.id = reader.uint64().toString();
                    break;
                case /* uint64 target_id = 2 [jstype = JS_STRING];*/ 2:
                    message.targetId = reader.uint64().toString();
                    break;
                case /* string job */ 3:
                    message.job = reader.string();
                    break;
                case /* int32 minimum_grade */ 4:
                    message.minimumGrade = reader.int32();
                    break;
                case /* resources.common.access.AccessLevel access */ 5:
                    message.access = reader.int32();
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
    internalBinaryWrite(message: DummyJobAccess, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 id = 1 [jstype = JS_STRING]; */
        if (message.id !== "0")
            writer.tag(1, WireType.Varint).uint64(message.id);
        /* uint64 target_id = 2 [jstype = JS_STRING]; */
        if (message.targetId !== "0")
            writer.tag(2, WireType.Varint).uint64(message.targetId);
        /* string job = 3; */
        if (message.job !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.job);
        /* int32 minimum_grade = 4; */
        if (message.minimumGrade !== 0)
            writer.tag(4, WireType.Varint).int32(message.minimumGrade);
        /* resources.common.access.AccessLevel access = 5; */
        if (message.access !== 0)
            writer.tag(5, WireType.Varint).int32(message.access);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.common.access.DummyJobAccess
 */
export const DummyJobAccess = new DummyJobAccess$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DummyUserAccess$Type extends MessageType<DummyUserAccess> {
    constructor() {
        super("resources.common.access.DummyUserAccess", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "target_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 3, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 4, name: "access", kind: "enum", T: () => ["resources.common.access.AccessLevel", AccessLevel, "ACCESS_LEVEL_"], options: { "validate.rules": { enum: { definedOnly: true } } } }
        ]);
    }
    create(value?: PartialMessage<DummyUserAccess>): DummyUserAccess {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.id = "0";
        message.targetId = "0";
        message.userId = 0;
        message.access = 0;
        if (value !== undefined)
            reflectionMergePartial<DummyUserAccess>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DummyUserAccess): DummyUserAccess {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 id = 1 [jstype = JS_STRING];*/ 1:
                    message.id = reader.uint64().toString();
                    break;
                case /* uint64 target_id = 2 [jstype = JS_STRING];*/ 2:
                    message.targetId = reader.uint64().toString();
                    break;
                case /* int32 user_id */ 3:
                    message.userId = reader.int32();
                    break;
                case /* resources.common.access.AccessLevel access */ 4:
                    message.access = reader.int32();
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
    internalBinaryWrite(message: DummyUserAccess, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 id = 1 [jstype = JS_STRING]; */
        if (message.id !== "0")
            writer.tag(1, WireType.Varint).uint64(message.id);
        /* uint64 target_id = 2 [jstype = JS_STRING]; */
        if (message.targetId !== "0")
            writer.tag(2, WireType.Varint).uint64(message.targetId);
        /* int32 user_id = 3; */
        if (message.userId !== 0)
            writer.tag(3, WireType.Varint).int32(message.userId);
        /* resources.common.access.AccessLevel access = 4; */
        if (message.access !== 0)
            writer.tag(4, WireType.Varint).int32(message.access);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.common.access.DummyUserAccess
 */
export const DummyUserAccess = new DummyUserAccess$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DummyQualificationAccess$Type extends MessageType<DummyQualificationAccess> {
    constructor() {
        super("resources.common.access.DummyQualificationAccess", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "target_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 3, name: "qualification_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 4, name: "access", kind: "enum", T: () => ["resources.common.access.AccessLevel", AccessLevel, "ACCESS_LEVEL_"], options: { "validate.rules": { enum: { definedOnly: true } } } }
        ]);
    }
    create(value?: PartialMessage<DummyQualificationAccess>): DummyQualificationAccess {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.id = "0";
        message.targetId = "0";
        message.qualificationId = "0";
        message.access = 0;
        if (value !== undefined)
            reflectionMergePartial<DummyQualificationAccess>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DummyQualificationAccess): DummyQualificationAccess {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 id = 1 [jstype = JS_STRING];*/ 1:
                    message.id = reader.uint64().toString();
                    break;
                case /* uint64 target_id = 2 [jstype = JS_STRING];*/ 2:
                    message.targetId = reader.uint64().toString();
                    break;
                case /* uint64 qualification_id = 3 [jstype = JS_STRING];*/ 3:
                    message.qualificationId = reader.uint64().toString();
                    break;
                case /* resources.common.access.AccessLevel access */ 4:
                    message.access = reader.int32();
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
    internalBinaryWrite(message: DummyQualificationAccess, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 id = 1 [jstype = JS_STRING]; */
        if (message.id !== "0")
            writer.tag(1, WireType.Varint).uint64(message.id);
        /* uint64 target_id = 2 [jstype = JS_STRING]; */
        if (message.targetId !== "0")
            writer.tag(2, WireType.Varint).uint64(message.targetId);
        /* uint64 qualification_id = 3 [jstype = JS_STRING]; */
        if (message.qualificationId !== "0")
            writer.tag(3, WireType.Varint).uint64(message.qualificationId);
        /* resources.common.access.AccessLevel access = 4; */
        if (message.access !== 0)
            writer.tag(4, WireType.Varint).int32(message.access);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.common.access.DummyQualificationAccess
 */
export const DummyQualificationAccess = new DummyQualificationAccess$Type();