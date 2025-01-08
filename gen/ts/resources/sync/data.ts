// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "resources/sync/data.proto" (package "resources.sync", syntax proto3)
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
import { UserLicenses } from "../users/users";
import { License } from "../users/users";
import { Vehicle } from "../vehicles/vehicles";
import { User } from "../users/users";
import { Job } from "../users/jobs";
/**
 * @generated from protobuf message resources.sync.DataStatus
 */
export interface DataStatus {
    /**
     * @generated from protobuf field: int64 count = 1;
     */
    count: number;
}
/**
 * @generated from protobuf message resources.sync.DataJobs
 */
export interface DataJobs {
    /**
     * @generated from protobuf field: repeated resources.users.Job jobs = 1;
     */
    jobs: Job[];
}
/**
 * @generated from protobuf message resources.sync.DataUsers
 */
export interface DataUsers {
    /**
     * @generated from protobuf field: repeated resources.users.User users = 1;
     */
    users: User[];
}
/**
 * @generated from protobuf message resources.sync.DataVehicles
 */
export interface DataVehicles {
    /**
     * @generated from protobuf field: repeated resources.vehicles.Vehicle vehicles = 1;
     */
    vehicles: Vehicle[];
}
/**
 * @generated from protobuf message resources.sync.DataLicenses
 */
export interface DataLicenses {
    /**
     * @generated from protobuf field: repeated resources.users.License licenses = 1;
     */
    licenses: License[];
}
/**
 * @generated from protobuf message resources.sync.DataUserLicenses
 */
export interface DataUserLicenses {
    /**
     * @generated from protobuf field: repeated resources.users.UserLicenses user_licenses = 1;
     */
    userLicenses: UserLicenses[];
}
// @generated message type with reflection information, may provide speed optimized methods
class DataStatus$Type extends MessageType<DataStatus> {
    constructor() {
        super("resources.sync.DataStatus", [
            { no: 1, name: "count", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 2 /*LongType.NUMBER*/ }
        ]);
    }
    create(value?: PartialMessage<DataStatus>): DataStatus {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.count = 0;
        if (value !== undefined)
            reflectionMergePartial<DataStatus>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DataStatus): DataStatus {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int64 count */ 1:
                    message.count = reader.int64().toNumber();
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
    internalBinaryWrite(message: DataStatus, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int64 count = 1; */
        if (message.count !== 0)
            writer.tag(1, WireType.Varint).int64(message.count);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.sync.DataStatus
 */
export const DataStatus = new DataStatus$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DataJobs$Type extends MessageType<DataJobs> {
    constructor() {
        super("resources.sync.DataJobs", [
            { no: 1, name: "jobs", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Job, options: { "validate.rules": { repeated: { maxItems: "200" } } } }
        ]);
    }
    create(value?: PartialMessage<DataJobs>): DataJobs {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.jobs = [];
        if (value !== undefined)
            reflectionMergePartial<DataJobs>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DataJobs): DataJobs {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated resources.users.Job jobs */ 1:
                    message.jobs.push(Job.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: DataJobs, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated resources.users.Job jobs = 1; */
        for (let i = 0; i < message.jobs.length; i++)
            Job.internalBinaryWrite(message.jobs[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.sync.DataJobs
 */
export const DataJobs = new DataJobs$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DataUsers$Type extends MessageType<DataUsers> {
    constructor() {
        super("resources.sync.DataUsers", [
            { no: 1, name: "users", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => User, options: { "validate.rules": { repeated: { maxItems: "500" } } } }
        ]);
    }
    create(value?: PartialMessage<DataUsers>): DataUsers {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.users = [];
        if (value !== undefined)
            reflectionMergePartial<DataUsers>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DataUsers): DataUsers {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated resources.users.User users */ 1:
                    message.users.push(User.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: DataUsers, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated resources.users.User users = 1; */
        for (let i = 0; i < message.users.length; i++)
            User.internalBinaryWrite(message.users[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.sync.DataUsers
 */
export const DataUsers = new DataUsers$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DataVehicles$Type extends MessageType<DataVehicles> {
    constructor() {
        super("resources.sync.DataVehicles", [
            { no: 1, name: "vehicles", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Vehicle, options: { "validate.rules": { repeated: { maxItems: "1000" } } } }
        ]);
    }
    create(value?: PartialMessage<DataVehicles>): DataVehicles {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.vehicles = [];
        if (value !== undefined)
            reflectionMergePartial<DataVehicles>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DataVehicles): DataVehicles {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated resources.vehicles.Vehicle vehicles */ 1:
                    message.vehicles.push(Vehicle.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: DataVehicles, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated resources.vehicles.Vehicle vehicles = 1; */
        for (let i = 0; i < message.vehicles.length; i++)
            Vehicle.internalBinaryWrite(message.vehicles[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.sync.DataVehicles
 */
export const DataVehicles = new DataVehicles$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DataLicenses$Type extends MessageType<DataLicenses> {
    constructor() {
        super("resources.sync.DataLicenses", [
            { no: 1, name: "licenses", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => License, options: { "validate.rules": { repeated: { maxItems: "200" } } } }
        ]);
    }
    create(value?: PartialMessage<DataLicenses>): DataLicenses {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.licenses = [];
        if (value !== undefined)
            reflectionMergePartial<DataLicenses>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DataLicenses): DataLicenses {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated resources.users.License licenses */ 1:
                    message.licenses.push(License.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: DataLicenses, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated resources.users.License licenses = 1; */
        for (let i = 0; i < message.licenses.length; i++)
            License.internalBinaryWrite(message.licenses[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.sync.DataLicenses
 */
export const DataLicenses = new DataLicenses$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DataUserLicenses$Type extends MessageType<DataUserLicenses> {
    constructor() {
        super("resources.sync.DataUserLicenses", [
            { no: 1, name: "user_licenses", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => UserLicenses, options: { "validate.rules": { repeated: { maxItems: "1000" } } } }
        ]);
    }
    create(value?: PartialMessage<DataUserLicenses>): DataUserLicenses {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.userLicenses = [];
        if (value !== undefined)
            reflectionMergePartial<DataUserLicenses>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DataUserLicenses): DataUserLicenses {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated resources.users.UserLicenses user_licenses */ 1:
                    message.userLicenses.push(UserLicenses.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: DataUserLicenses, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated resources.users.UserLicenses user_licenses = 1; */
        for (let i = 0; i < message.userLicenses.length; i++)
            UserLicenses.internalBinaryWrite(message.userLicenses[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.sync.DataUserLicenses
 */
export const DataUserLicenses = new DataUserLicenses$Type();