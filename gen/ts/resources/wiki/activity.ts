// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "resources/wiki/activity.proto" (package "resources.wiki", syntax proto3)
// @ts-nocheck
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * TODO
 *
 * @generated from protobuf message resources.wiki.PageActivity
 */
export interface PageActivity {
}
// @generated message type with reflection information, may provide speed optimized methods
class PageActivity$Type extends MessageType<PageActivity> {
    constructor() {
        super("resources.wiki.PageActivity", []);
    }
    create(value?: PartialMessage<PageActivity>): PageActivity {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<PageActivity>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PageActivity): PageActivity {
        return target ?? this.create();
    }
    internalBinaryWrite(message: PageActivity, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.wiki.PageActivity
 */
export const PageActivity = new PageActivity$Type();
