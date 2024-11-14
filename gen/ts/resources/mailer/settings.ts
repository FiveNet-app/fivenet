// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "resources/mailer/settings.proto" (package "resources.mailer", syntax proto3)
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
import { UserShort } from "../users/users";
/**
 * @generated from protobuf message resources.mailer.UserSettings
 */
export interface UserSettings {
    /**
     * @generated from protobuf field: int32 user_id = 1;
     */
    userId: number;
    /**
     * @generated from protobuf field: repeated resources.mailer.BlockedUser blocked_users = 2;
     */
    blockedUsers: BlockedUser[];
}
/**
 * @generated from protobuf message resources.mailer.BlockedUser
 */
export interface BlockedUser {
    /**
     * @generated from protobuf field: int32 user_id = 1;
     */
    userId: number;
    /**
     * @generated from protobuf field: optional resources.users.UserShort user = 2;
     */
    user?: UserShort;
}
// @generated message type with reflection information, may provide speed optimized methods
class UserSettings$Type extends MessageType<UserSettings> {
    constructor() {
        super("resources.mailer.UserSettings", [
            { no: 1, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gte: 0 } } } },
            { no: 2, name: "blocked_users", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => BlockedUser, options: { "validate.rules": { repeated: { maxItems: "25" } } } }
        ]);
    }
    create(value?: PartialMessage<UserSettings>): UserSettings {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.userId = 0;
        message.blockedUsers = [];
        if (value !== undefined)
            reflectionMergePartial<UserSettings>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UserSettings): UserSettings {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 user_id */ 1:
                    message.userId = reader.int32();
                    break;
                case /* repeated resources.mailer.BlockedUser blocked_users */ 2:
                    message.blockedUsers.push(BlockedUser.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: UserSettings, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 user_id = 1; */
        if (message.userId !== 0)
            writer.tag(1, WireType.Varint).int32(message.userId);
        /* repeated resources.mailer.BlockedUser blocked_users = 2; */
        for (let i = 0; i < message.blockedUsers.length; i++)
            BlockedUser.internalBinaryWrite(message.blockedUsers[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.mailer.UserSettings
 */
export const UserSettings = new UserSettings$Type();
// @generated message type with reflection information, may provide speed optimized methods
class BlockedUser$Type extends MessageType<BlockedUser> {
    constructor() {
        super("resources.mailer.BlockedUser", [
            { no: 1, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "user", kind: "message", T: () => UserShort }
        ]);
    }
    create(value?: PartialMessage<BlockedUser>): BlockedUser {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.userId = 0;
        if (value !== undefined)
            reflectionMergePartial<BlockedUser>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: BlockedUser): BlockedUser {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 user_id */ 1:
                    message.userId = reader.int32();
                    break;
                case /* optional resources.users.UserShort user */ 2:
                    message.user = UserShort.internalBinaryRead(reader, reader.uint32(), options, message.user);
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
    internalBinaryWrite(message: BlockedUser, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 user_id = 1; */
        if (message.userId !== 0)
            writer.tag(1, WireType.Varint).int32(message.userId);
        /* optional resources.users.UserShort user = 2; */
        if (message.user)
            UserShort.internalBinaryWrite(message.user, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.mailer.BlockedUser
 */
export const BlockedUser = new BlockedUser$Type();
