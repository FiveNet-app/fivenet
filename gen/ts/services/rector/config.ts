// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_bigint
// @generated from protobuf file "services/rector/config.proto" (package "services.rector", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { AppConfig } from "../../resources/rector/config";
/**
 * @generated from protobuf message services.rector.GetAppConfigRequest
 */
export interface GetAppConfigRequest {
}
/**
 * @generated from protobuf message services.rector.GetAppConfigResponse
 */
export interface GetAppConfigResponse {
    /**
     * @generated from protobuf field: resources.rector.AppConfig config = 1;
     */
    config?: AppConfig;
}
/**
 * @generated from protobuf message services.rector.UpdateAppConfigRequest
 */
export interface UpdateAppConfigRequest {
    /**
     * @generated from protobuf field: resources.rector.AppConfig config = 1;
     */
    config?: AppConfig;
}
/**
 * @generated from protobuf message services.rector.UpdateAppConfigResponse
 */
export interface UpdateAppConfigResponse {
    /**
     * @generated from protobuf field: resources.rector.AppConfig config = 1;
     */
    config?: AppConfig;
}
// @generated message type with reflection information, may provide speed optimized methods
class GetAppConfigRequest$Type extends MessageType<GetAppConfigRequest> {
    constructor() {
        super("services.rector.GetAppConfigRequest", []);
    }
    create(value?: PartialMessage<GetAppConfigRequest>): GetAppConfigRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<GetAppConfigRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetAppConfigRequest): GetAppConfigRequest {
        return target ?? this.create();
    }
    internalBinaryWrite(message: GetAppConfigRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.rector.GetAppConfigRequest
 */
export const GetAppConfigRequest = new GetAppConfigRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetAppConfigResponse$Type extends MessageType<GetAppConfigResponse> {
    constructor() {
        super("services.rector.GetAppConfigResponse", [
            { no: 1, name: "config", kind: "message", T: () => AppConfig }
        ]);
    }
    create(value?: PartialMessage<GetAppConfigResponse>): GetAppConfigResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<GetAppConfigResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetAppConfigResponse): GetAppConfigResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.rector.AppConfig config */ 1:
                    message.config = AppConfig.internalBinaryRead(reader, reader.uint32(), options, message.config);
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
    internalBinaryWrite(message: GetAppConfigResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.rector.AppConfig config = 1; */
        if (message.config)
            AppConfig.internalBinaryWrite(message.config, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.rector.GetAppConfigResponse
 */
export const GetAppConfigResponse = new GetAppConfigResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateAppConfigRequest$Type extends MessageType<UpdateAppConfigRequest> {
    constructor() {
        super("services.rector.UpdateAppConfigRequest", [
            { no: 1, name: "config", kind: "message", T: () => AppConfig, options: { "validate.rules": { message: { required: true } } } }
        ]);
    }
    create(value?: PartialMessage<UpdateAppConfigRequest>): UpdateAppConfigRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<UpdateAppConfigRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UpdateAppConfigRequest): UpdateAppConfigRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.rector.AppConfig config */ 1:
                    message.config = AppConfig.internalBinaryRead(reader, reader.uint32(), options, message.config);
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
    internalBinaryWrite(message: UpdateAppConfigRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.rector.AppConfig config = 1; */
        if (message.config)
            AppConfig.internalBinaryWrite(message.config, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.rector.UpdateAppConfigRequest
 */
export const UpdateAppConfigRequest = new UpdateAppConfigRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateAppConfigResponse$Type extends MessageType<UpdateAppConfigResponse> {
    constructor() {
        super("services.rector.UpdateAppConfigResponse", [
            { no: 1, name: "config", kind: "message", T: () => AppConfig }
        ]);
    }
    create(value?: PartialMessage<UpdateAppConfigResponse>): UpdateAppConfigResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<UpdateAppConfigResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UpdateAppConfigResponse): UpdateAppConfigResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.rector.AppConfig config */ 1:
                    message.config = AppConfig.internalBinaryRead(reader, reader.uint32(), options, message.config);
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
    internalBinaryWrite(message: UpdateAppConfigResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.rector.AppConfig config = 1; */
        if (message.config)
            AppConfig.internalBinaryWrite(message.config, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.rector.UpdateAppConfigResponse
 */
export const UpdateAppConfigResponse = new UpdateAppConfigResponse$Type();
/**
 * @generated ServiceType for protobuf service services.rector.RectorConfigService
 */
export const RectorConfigService = new ServiceType("services.rector.RectorConfigService", [
    { name: "GetAppConfig", options: {}, I: GetAppConfigRequest, O: GetAppConfigResponse },
    { name: "UpdateAppConfig", options: {}, I: UpdateAppConfigRequest, O: UpdateAppConfigResponse }
]);
