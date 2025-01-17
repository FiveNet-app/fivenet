// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "resources/internet/page.proto" (package "resources.internet", syntax proto3)
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
import { Timestamp } from "../timestamp/timestamp";
/**
 * @generated from protobuf message resources.internet.Page
 */
export interface Page {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp created_at = 2;
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
     * @generated from protobuf field: uint64 domain_id = 5 [jstype = JS_STRING];
     */
    domainId: string;
    /**
     * @sanitize: method=StripTags
     *
     * @generated from protobuf field: string path = 6;
     */
    path: string;
    /**
     * @sanitize: method=StripTags
     *
     * @generated from protobuf field: string title = 7;
     */
    title: string;
    /**
     * @sanitize: method=StripTags
     *
     * @generated from protobuf field: string description = 8;
     */
    description: string;
    /**
     * @generated from protobuf field: resources.internet.PageData data = 9;
     */
    data?: PageData;
    /**
     * @generated from protobuf field: optional string creator_job = 10;
     */
    creatorJob?: string;
    /**
     * @generated from protobuf field: optional int32 creator_id = 11;
     */
    creatorId?: number;
}
/**
 * TODO
 *
 * @generated from protobuf message resources.internet.PageData
 */
export interface PageData {
}
/**
 * @generated from protobuf enum resources.internet.PageLayoutType
 */
export enum PageLayoutType {
    /**
     * @generated from protobuf enum value: PAGE_LAYOUT_TYPE_UNSPECIFIED = 0;
     */
    PAGE_LAYOUT_TYPE_UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: PAGE_LAYOUT_TYPE_ = 1;
     */
    PAGE_LAYOUT_TYPE_ = 1
}
// @generated message type with reflection information, may provide speed optimized methods
class Page$Type extends MessageType<Page> {
    constructor() {
        super("resources.internet.Page", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "deleted_at", kind: "message", T: () => Timestamp },
            { no: 5, name: "domain_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 6, name: "path", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "128" } } } },
            { no: 7, name: "title", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "1", maxLen: "255" } } } },
            { no: 8, name: "description", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3", maxLen: "512" } } } },
            { no: 9, name: "data", kind: "message", T: () => PageData, options: { "validate.rules": { message: { required: true } } } },
            { no: 10, name: "creator_job", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 11, name: "creator_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value?: PartialMessage<Page>): Page {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.id = "0";
        message.domainId = "0";
        message.path = "";
        message.title = "";
        message.description = "";
        if (value !== undefined)
            reflectionMergePartial<Page>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Page): Page {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 id = 1 [jstype = JS_STRING];*/ 1:
                    message.id = reader.uint64().toString();
                    break;
                case /* resources.timestamp.Timestamp created_at */ 2:
                    message.createdAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.createdAt);
                    break;
                case /* optional resources.timestamp.Timestamp updated_at */ 3:
                    message.updatedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.updatedAt);
                    break;
                case /* optional resources.timestamp.Timestamp deleted_at */ 4:
                    message.deletedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.deletedAt);
                    break;
                case /* uint64 domain_id = 5 [jstype = JS_STRING];*/ 5:
                    message.domainId = reader.uint64().toString();
                    break;
                case /* string path */ 6:
                    message.path = reader.string();
                    break;
                case /* string title */ 7:
                    message.title = reader.string();
                    break;
                case /* string description */ 8:
                    message.description = reader.string();
                    break;
                case /* resources.internet.PageData data */ 9:
                    message.data = PageData.internalBinaryRead(reader, reader.uint32(), options, message.data);
                    break;
                case /* optional string creator_job */ 10:
                    message.creatorJob = reader.string();
                    break;
                case /* optional int32 creator_id */ 11:
                    message.creatorId = reader.int32();
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
    internalBinaryWrite(message: Page, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 id = 1 [jstype = JS_STRING]; */
        if (message.id !== "0")
            writer.tag(1, WireType.Varint).uint64(message.id);
        /* resources.timestamp.Timestamp created_at = 2; */
        if (message.createdAt)
            Timestamp.internalBinaryWrite(message.createdAt, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.timestamp.Timestamp updated_at = 3; */
        if (message.updatedAt)
            Timestamp.internalBinaryWrite(message.updatedAt, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.timestamp.Timestamp deleted_at = 4; */
        if (message.deletedAt)
            Timestamp.internalBinaryWrite(message.deletedAt, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* uint64 domain_id = 5 [jstype = JS_STRING]; */
        if (message.domainId !== "0")
            writer.tag(5, WireType.Varint).uint64(message.domainId);
        /* string path = 6; */
        if (message.path !== "")
            writer.tag(6, WireType.LengthDelimited).string(message.path);
        /* string title = 7; */
        if (message.title !== "")
            writer.tag(7, WireType.LengthDelimited).string(message.title);
        /* string description = 8; */
        if (message.description !== "")
            writer.tag(8, WireType.LengthDelimited).string(message.description);
        /* resources.internet.PageData data = 9; */
        if (message.data)
            PageData.internalBinaryWrite(message.data, writer.tag(9, WireType.LengthDelimited).fork(), options).join();
        /* optional string creator_job = 10; */
        if (message.creatorJob !== undefined)
            writer.tag(10, WireType.LengthDelimited).string(message.creatorJob);
        /* optional int32 creator_id = 11; */
        if (message.creatorId !== undefined)
            writer.tag(11, WireType.Varint).int32(message.creatorId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.internet.Page
 */
export const Page = new Page$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PageData$Type extends MessageType<PageData> {
    constructor() {
        super("resources.internet.PageData", []);
    }
    create(value?: PartialMessage<PageData>): PageData {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<PageData>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PageData): PageData {
        return target ?? this.create();
    }
    internalBinaryWrite(message: PageData, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.internet.PageData
 */
export const PageData = new PageData$Type();
