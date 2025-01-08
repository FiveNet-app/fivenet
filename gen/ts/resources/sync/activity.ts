// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "resources/sync/activity.proto" (package "resources.sync", syntax proto3)
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
 * Connect an identifier/license to the provider with the specified external id
 * (e.g., auto discord social connect on server join)
 *
 * @generated from protobuf message resources.sync.UserOAuth2Conn
 */
export interface UserOAuth2Conn {
    /**
     * @generated from protobuf field: string provider_name = 1;
     */
    providerName: string;
    /**
     * @generated from protobuf field: string identifier = 2;
     */
    identifier: string;
    /**
     * @generated from protobuf field: string external_id = 3;
     */
    externalId: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class UserOAuth2Conn$Type extends MessageType<UserOAuth2Conn> {
    constructor() {
        super("resources.sync.UserOAuth2Conn", [
            { no: 1, name: "provider_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "identifier", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "external_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<UserOAuth2Conn>): UserOAuth2Conn {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.providerName = "";
        message.identifier = "";
        message.externalId = "";
        if (value !== undefined)
            reflectionMergePartial<UserOAuth2Conn>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UserOAuth2Conn): UserOAuth2Conn {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string provider_name */ 1:
                    message.providerName = reader.string();
                    break;
                case /* string identifier */ 2:
                    message.identifier = reader.string();
                    break;
                case /* string external_id */ 3:
                    message.externalId = reader.string();
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
    internalBinaryWrite(message: UserOAuth2Conn, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string provider_name = 1; */
        if (message.providerName !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.providerName);
        /* string identifier = 2; */
        if (message.identifier !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.identifier);
        /* string external_id = 3; */
        if (message.externalId !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.externalId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.sync.UserOAuth2Conn
 */
export const UserOAuth2Conn = new UserOAuth2Conn$Type();