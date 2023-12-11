// @generated by protobuf-ts 2.9.3 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "resources/documents/category.proto" (package "resources.documents", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message resources.documents.Category
 */
export interface Category {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
    /**
     * @sanitize
     *
     * @generated from protobuf field: string name = 2;
     */
    name: string;
    /**
     * @sanitize
     *
     * @generated from protobuf field: optional string description = 3;
     */
    description?: string;
    /**
     * @generated from protobuf field: optional string job = 4;
     */
    job?: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class Category$Type extends MessageType<Category> {
    constructor() {
        super("resources.documents.Category", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3", maxLen: "128" } } } },
            { no: 3, name: "description", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 4, name: "job", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.Category
 */
export const Category = new Category$Type();
