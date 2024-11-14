// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "services/completor/completor.proto" (package "services.completor", syntax proto3)
// @ts-nocheck
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { CitizenAttribute } from "../../resources/users/users";
import { LawBook } from "../../resources/laws/laws";
import { Category } from "../../resources/documents/category";
import { Job } from "../../resources/users/jobs";
import { UserShort } from "../../resources/users/users";
/**
 * @generated from protobuf message services.completor.CompleteCitizensRequest
 */
export interface CompleteCitizensRequest {
    /**
     * @generated from protobuf field: string search = 1;
     */
    search: string;
    /**
     * @generated from protobuf field: optional bool current_job = 2;
     */
    currentJob?: boolean;
    /**
     * @generated from protobuf field: optional bool on_duty = 3;
     */
    onDuty?: boolean;
    /**
     * @generated from protobuf field: optional int32 user_id = 4;
     */
    userId?: number;
}
/**
 * @generated from protobuf message services.completor.CompleteCitizensRespoonse
 */
export interface CompleteCitizensRespoonse {
    /**
     * @generated from protobuf field: repeated resources.users.UserShort users = 1;
     */
    users: UserShort[]; // @gotags: alias:"user"
}
/**
 * @generated from protobuf message services.completor.CompleteJobsRequest
 */
export interface CompleteJobsRequest {
    /**
     * @generated from protobuf field: optional string search = 1;
     */
    search?: string;
    /**
     * @generated from protobuf field: optional bool exact_match = 2;
     */
    exactMatch?: boolean;
    /**
     * @generated from protobuf field: optional bool current_job = 3;
     */
    currentJob?: boolean;
}
/**
 * @generated from protobuf message services.completor.CompleteJobsResponse
 */
export interface CompleteJobsResponse {
    /**
     * @generated from protobuf field: repeated resources.users.Job jobs = 1;
     */
    jobs: Job[];
}
/**
 * @generated from protobuf message services.completor.CompleteDocumentCategoriesRequest
 */
export interface CompleteDocumentCategoriesRequest {
    /**
     * @generated from protobuf field: string search = 1;
     */
    search: string;
}
/**
 * @generated from protobuf message services.completor.CompleteDocumentCategoriesResponse
 */
export interface CompleteDocumentCategoriesResponse {
    /**
     * @generated from protobuf field: repeated resources.documents.Category categories = 1;
     */
    categories: Category[];
}
/**
 * @generated from protobuf message services.completor.ListLawBooksRequest
 */
export interface ListLawBooksRequest {
}
/**
 * @generated from protobuf message services.completor.ListLawBooksResponse
 */
export interface ListLawBooksResponse {
    /**
     * @generated from protobuf field: repeated resources.laws.LawBook books = 1;
     */
    books: LawBook[];
}
/**
 * @generated from protobuf message services.completor.CompleteCitizenAttributesRequest
 */
export interface CompleteCitizenAttributesRequest {
    /**
     * @generated from protobuf field: string search = 1;
     */
    search: string;
}
/**
 * @generated from protobuf message services.completor.CompleteCitizenAttributesResponse
 */
export interface CompleteCitizenAttributesResponse {
    /**
     * @generated from protobuf field: repeated resources.users.CitizenAttribute attributes = 1;
     */
    attributes: CitizenAttribute[];
}
// @generated message type with reflection information, may provide speed optimized methods
class CompleteCitizensRequest$Type extends MessageType<CompleteCitizensRequest> {
    constructor() {
        super("services.completor.CompleteCitizensRequest", [
            { no: 1, name: "search", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "64" } } } },
            { no: 2, name: "current_job", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "on_duty", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "user_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } }
        ]);
    }
    create(value?: PartialMessage<CompleteCitizensRequest>): CompleteCitizensRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.search = "";
        if (value !== undefined)
            reflectionMergePartial<CompleteCitizensRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CompleteCitizensRequest): CompleteCitizensRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string search */ 1:
                    message.search = reader.string();
                    break;
                case /* optional bool current_job */ 2:
                    message.currentJob = reader.bool();
                    break;
                case /* optional bool on_duty */ 3:
                    message.onDuty = reader.bool();
                    break;
                case /* optional int32 user_id */ 4:
                    message.userId = reader.int32();
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
    internalBinaryWrite(message: CompleteCitizensRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string search = 1; */
        if (message.search !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.search);
        /* optional bool current_job = 2; */
        if (message.currentJob !== undefined)
            writer.tag(2, WireType.Varint).bool(message.currentJob);
        /* optional bool on_duty = 3; */
        if (message.onDuty !== undefined)
            writer.tag(3, WireType.Varint).bool(message.onDuty);
        /* optional int32 user_id = 4; */
        if (message.userId !== undefined)
            writer.tag(4, WireType.Varint).int32(message.userId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.completor.CompleteCitizensRequest
 */
export const CompleteCitizensRequest = new CompleteCitizensRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CompleteCitizensRespoonse$Type extends MessageType<CompleteCitizensRespoonse> {
    constructor() {
        super("services.completor.CompleteCitizensRespoonse", [
            { no: 1, name: "users", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => UserShort }
        ]);
    }
    create(value?: PartialMessage<CompleteCitizensRespoonse>): CompleteCitizensRespoonse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.users = [];
        if (value !== undefined)
            reflectionMergePartial<CompleteCitizensRespoonse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CompleteCitizensRespoonse): CompleteCitizensRespoonse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated resources.users.UserShort users */ 1:
                    message.users.push(UserShort.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: CompleteCitizensRespoonse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated resources.users.UserShort users = 1; */
        for (let i = 0; i < message.users.length; i++)
            UserShort.internalBinaryWrite(message.users[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.completor.CompleteCitizensRespoonse
 */
export const CompleteCitizensRespoonse = new CompleteCitizensRespoonse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CompleteJobsRequest$Type extends MessageType<CompleteJobsRequest> {
    constructor() {
        super("services.completor.CompleteJobsRequest", [
            { no: 1, name: "search", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "64" } } } },
            { no: 2, name: "exact_match", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "current_job", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value?: PartialMessage<CompleteJobsRequest>): CompleteJobsRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<CompleteJobsRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CompleteJobsRequest): CompleteJobsRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* optional string search */ 1:
                    message.search = reader.string();
                    break;
                case /* optional bool exact_match */ 2:
                    message.exactMatch = reader.bool();
                    break;
                case /* optional bool current_job */ 3:
                    message.currentJob = reader.bool();
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
    internalBinaryWrite(message: CompleteJobsRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* optional string search = 1; */
        if (message.search !== undefined)
            writer.tag(1, WireType.LengthDelimited).string(message.search);
        /* optional bool exact_match = 2; */
        if (message.exactMatch !== undefined)
            writer.tag(2, WireType.Varint).bool(message.exactMatch);
        /* optional bool current_job = 3; */
        if (message.currentJob !== undefined)
            writer.tag(3, WireType.Varint).bool(message.currentJob);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.completor.CompleteJobsRequest
 */
export const CompleteJobsRequest = new CompleteJobsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CompleteJobsResponse$Type extends MessageType<CompleteJobsResponse> {
    constructor() {
        super("services.completor.CompleteJobsResponse", [
            { no: 1, name: "jobs", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Job }
        ]);
    }
    create(value?: PartialMessage<CompleteJobsResponse>): CompleteJobsResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.jobs = [];
        if (value !== undefined)
            reflectionMergePartial<CompleteJobsResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CompleteJobsResponse): CompleteJobsResponse {
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
    internalBinaryWrite(message: CompleteJobsResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
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
 * @generated MessageType for protobuf message services.completor.CompleteJobsResponse
 */
export const CompleteJobsResponse = new CompleteJobsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CompleteDocumentCategoriesRequest$Type extends MessageType<CompleteDocumentCategoriesRequest> {
    constructor() {
        super("services.completor.CompleteDocumentCategoriesRequest", [
            { no: 1, name: "search", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "64" } } } }
        ]);
    }
    create(value?: PartialMessage<CompleteDocumentCategoriesRequest>): CompleteDocumentCategoriesRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.search = "";
        if (value !== undefined)
            reflectionMergePartial<CompleteDocumentCategoriesRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CompleteDocumentCategoriesRequest): CompleteDocumentCategoriesRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string search */ 1:
                    message.search = reader.string();
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
    internalBinaryWrite(message: CompleteDocumentCategoriesRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string search = 1; */
        if (message.search !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.search);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.completor.CompleteDocumentCategoriesRequest
 */
export const CompleteDocumentCategoriesRequest = new CompleteDocumentCategoriesRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CompleteDocumentCategoriesResponse$Type extends MessageType<CompleteDocumentCategoriesResponse> {
    constructor() {
        super("services.completor.CompleteDocumentCategoriesResponse", [
            { no: 1, name: "categories", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Category }
        ]);
    }
    create(value?: PartialMessage<CompleteDocumentCategoriesResponse>): CompleteDocumentCategoriesResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.categories = [];
        if (value !== undefined)
            reflectionMergePartial<CompleteDocumentCategoriesResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CompleteDocumentCategoriesResponse): CompleteDocumentCategoriesResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated resources.documents.Category categories */ 1:
                    message.categories.push(Category.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: CompleteDocumentCategoriesResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated resources.documents.Category categories = 1; */
        for (let i = 0; i < message.categories.length; i++)
            Category.internalBinaryWrite(message.categories[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.completor.CompleteDocumentCategoriesResponse
 */
export const CompleteDocumentCategoriesResponse = new CompleteDocumentCategoriesResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListLawBooksRequest$Type extends MessageType<ListLawBooksRequest> {
    constructor() {
        super("services.completor.ListLawBooksRequest", []);
    }
    create(value?: PartialMessage<ListLawBooksRequest>): ListLawBooksRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<ListLawBooksRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListLawBooksRequest): ListLawBooksRequest {
        return target ?? this.create();
    }
    internalBinaryWrite(message: ListLawBooksRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.completor.ListLawBooksRequest
 */
export const ListLawBooksRequest = new ListLawBooksRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListLawBooksResponse$Type extends MessageType<ListLawBooksResponse> {
    constructor() {
        super("services.completor.ListLawBooksResponse", [
            { no: 1, name: "books", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => LawBook }
        ]);
    }
    create(value?: PartialMessage<ListLawBooksResponse>): ListLawBooksResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.books = [];
        if (value !== undefined)
            reflectionMergePartial<ListLawBooksResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListLawBooksResponse): ListLawBooksResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated resources.laws.LawBook books */ 1:
                    message.books.push(LawBook.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: ListLawBooksResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated resources.laws.LawBook books = 1; */
        for (let i = 0; i < message.books.length; i++)
            LawBook.internalBinaryWrite(message.books[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.completor.ListLawBooksResponse
 */
export const ListLawBooksResponse = new ListLawBooksResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CompleteCitizenAttributesRequest$Type extends MessageType<CompleteCitizenAttributesRequest> {
    constructor() {
        super("services.completor.CompleteCitizenAttributesRequest", [
            { no: 1, name: "search", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "64" } } } }
        ]);
    }
    create(value?: PartialMessage<CompleteCitizenAttributesRequest>): CompleteCitizenAttributesRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.search = "";
        if (value !== undefined)
            reflectionMergePartial<CompleteCitizenAttributesRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CompleteCitizenAttributesRequest): CompleteCitizenAttributesRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string search */ 1:
                    message.search = reader.string();
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
    internalBinaryWrite(message: CompleteCitizenAttributesRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string search = 1; */
        if (message.search !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.search);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.completor.CompleteCitizenAttributesRequest
 */
export const CompleteCitizenAttributesRequest = new CompleteCitizenAttributesRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CompleteCitizenAttributesResponse$Type extends MessageType<CompleteCitizenAttributesResponse> {
    constructor() {
        super("services.completor.CompleteCitizenAttributesResponse", [
            { no: 1, name: "attributes", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => CitizenAttribute }
        ]);
    }
    create(value?: PartialMessage<CompleteCitizenAttributesResponse>): CompleteCitizenAttributesResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.attributes = [];
        if (value !== undefined)
            reflectionMergePartial<CompleteCitizenAttributesResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CompleteCitizenAttributesResponse): CompleteCitizenAttributesResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated resources.users.CitizenAttribute attributes */ 1:
                    message.attributes.push(CitizenAttribute.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: CompleteCitizenAttributesResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated resources.users.CitizenAttribute attributes = 1; */
        for (let i = 0; i < message.attributes.length; i++)
            CitizenAttribute.internalBinaryWrite(message.attributes[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.completor.CompleteCitizenAttributesResponse
 */
export const CompleteCitizenAttributesResponse = new CompleteCitizenAttributesResponse$Type();
/**
 * @generated ServiceType for protobuf service services.completor.CompletorService
 */
export const CompletorService = new ServiceType("services.completor.CompletorService", [
    { name: "CompleteCitizens", options: {}, I: CompleteCitizensRequest, O: CompleteCitizensRespoonse },
    { name: "CompleteJobs", options: {}, I: CompleteJobsRequest, O: CompleteJobsResponse },
    { name: "CompleteDocumentCategories", options: {}, I: CompleteDocumentCategoriesRequest, O: CompleteDocumentCategoriesResponse },
    { name: "ListLawBooks", options: {}, I: ListLawBooksRequest, O: ListLawBooksResponse },
    { name: "CompleteCitizenAttributes", options: {}, I: CompleteCitizenAttributesRequest, O: CompleteCitizenAttributesResponse }
]);
