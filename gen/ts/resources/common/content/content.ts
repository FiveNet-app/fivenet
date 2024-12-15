// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "resources/common/content/content.proto" (package "resources.common.content", syntax proto3)
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
 * @generated from protobuf message resources.common.content.Content
 */
export interface Content {
    /**
     * @generated from protobuf field: optional string version = 1;
     */
    version?: string;
    /**
     * @generated from protobuf field: optional resources.common.content.JSONNode content = 2;
     */
    content?: JSONNode;
    /**
     * @sanitize
     *
     * @generated from protobuf field: optional string raw_content = 3;
     */
    rawContent?: string;
}
/**
 * @generated from protobuf message resources.common.content.JSONNode
 */
export interface JSONNode {
    /**
     * @sanitize: method=StripTags
     *
     * @generated from protobuf field: string type = 1;
     */
    type: string;
    /**
     * @sanitize: method=StripTags
     *
     * @generated from protobuf field: string id = 2;
     */
    id: string;
    /**
     * @sanitize: method=StripTags
     *
     * @generated from protobuf field: string tag = 3;
     */
    tag: string;
    /**
     * @sanitize: method=StripTags
     *
     * @generated from protobuf field: map<string, string> attributes = 4;
     */
    attributes: {
        [key: string]: string;
    };
    /**
     * @sanitize: method=StripTags
     *
     * @generated from protobuf field: string class = 5;
     */
    class: string;
    /**
     * @sanitize: method=StripTags
     *
     * @generated from protobuf field: string text = 6;
     */
    text: string;
    /**
     * @generated from protobuf field: repeated resources.common.content.JSONNode children = 7;
     */
    children: JSONNode[];
}
/**
 * @generated from protobuf enum resources.common.content.ContentType
 */
export enum ContentType {
    /**
     * @generated from protobuf enum value: CONTENT_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: CONTENT_TYPE_HTML = 1;
     */
    HTML = 1,
    /**
     * @generated from protobuf enum value: CONTENT_TYPE_PLAIN = 2;
     */
    PLAIN = 2
}
// @generated message type with reflection information, may provide speed optimized methods
class Content$Type extends MessageType<Content> {
    constructor() {
        super("resources.common.content.Content", [
            { no: 1, name: "version", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "24" } } } },
            { no: 2, name: "content", kind: "message", T: () => JSONNode },
            { no: 3, name: "raw_content", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxBytes: "2000000" } } } }
        ]);
    }
    create(value?: PartialMessage<Content>): Content {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<Content>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Content): Content {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* optional string version */ 1:
                    message.version = reader.string();
                    break;
                case /* optional resources.common.content.JSONNode content */ 2:
                    message.content = JSONNode.internalBinaryRead(reader, reader.uint32(), options, message.content);
                    break;
                case /* optional string raw_content */ 3:
                    message.rawContent = reader.string();
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
    internalBinaryWrite(message: Content, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* optional string version = 1; */
        if (message.version !== undefined)
            writer.tag(1, WireType.LengthDelimited).string(message.version);
        /* optional resources.common.content.JSONNode content = 2; */
        if (message.content)
            JSONNode.internalBinaryWrite(message.content, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* optional string raw_content = 3; */
        if (message.rawContent !== undefined)
            writer.tag(3, WireType.LengthDelimited).string(message.rawContent);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.common.content.Content
 */
export const Content = new Content$Type();
// @generated message type with reflection information, may provide speed optimized methods
class JSONNode$Type extends MessageType<JSONNode> {
    constructor() {
        super("resources.common.content.JSONNode", [
            { no: 1, name: "type", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "tag", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "attributes", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } },
            { no: 5, name: "class", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 6, name: "text", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 7, name: "children", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => JSONNode }
        ]);
    }
    create(value?: PartialMessage<JSONNode>): JSONNode {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.type = "";
        message.id = "";
        message.tag = "";
        message.attributes = {};
        message.class = "";
        message.text = "";
        message.children = [];
        if (value !== undefined)
            reflectionMergePartial<JSONNode>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: JSONNode): JSONNode {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string type */ 1:
                    message.type = reader.string();
                    break;
                case /* string id */ 2:
                    message.id = reader.string();
                    break;
                case /* string tag */ 3:
                    message.tag = reader.string();
                    break;
                case /* map<string, string> attributes */ 4:
                    this.binaryReadMap4(message.attributes, reader, options);
                    break;
                case /* string class */ 5:
                    message.class = reader.string();
                    break;
                case /* string text */ 6:
                    message.text = reader.string();
                    break;
                case /* repeated resources.common.content.JSONNode children */ 7:
                    message.children.push(JSONNode.internalBinaryRead(reader, reader.uint32(), options));
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
    private binaryReadMap4(map: JSONNode["attributes"], reader: IBinaryReader, options: BinaryReadOptions): void {
        let len = reader.uint32(), end = reader.pos + len, key: keyof JSONNode["attributes"] | undefined, val: JSONNode["attributes"][any] | undefined;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case 1:
                    key = reader.string();
                    break;
                case 2:
                    val = reader.string();
                    break;
                default: throw new globalThis.Error("unknown map entry field for field resources.common.content.JSONNode.attributes");
            }
        }
        map[key ?? ""] = val ?? "";
    }
    internalBinaryWrite(message: JSONNode, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string type = 1; */
        if (message.type !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.type);
        /* string id = 2; */
        if (message.id !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.id);
        /* string tag = 3; */
        if (message.tag !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.tag);
        /* map<string, string> attributes = 4; */
        for (let k of globalThis.Object.keys(message.attributes))
            writer.tag(4, WireType.LengthDelimited).fork().tag(1, WireType.LengthDelimited).string(k).tag(2, WireType.LengthDelimited).string(message.attributes[k]).join();
        /* string class = 5; */
        if (message.class !== "")
            writer.tag(5, WireType.LengthDelimited).string(message.class);
        /* string text = 6; */
        if (message.text !== "")
            writer.tag(6, WireType.LengthDelimited).string(message.text);
        /* repeated resources.common.content.JSONNode children = 7; */
        for (let i = 0; i < message.children.length; i++)
            JSONNode.internalBinaryWrite(message.children[i], writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.common.content.JSONNode
 */
export const JSONNode = new JSONNode$Type();
