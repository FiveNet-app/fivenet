// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size,long_type_number
// @generated from protobuf file "resources/livemap/livemap.proto" (package "resources.livemap", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { UserShort } from "../users/users.js";
import { Timestamp } from "../timestamp/timestamp.js";
/**
 * @generated from protobuf message resources.livemap.DispatchMarker
 */
export interface DispatchMarker {
    /**
     * @generated from protobuf field: float x = 1;
     */
    x: number; // @gotags: alias:"x"
    /**
     * @generated from protobuf field: float y = 2;
     */
    y: number; // @gotags: alias:"y"
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp; // @gotags: alias:"updated_at"
    /**
     * @generated from protobuf field: int32 id = 4;
     */
    id: number;
    /**
     * @generated from protobuf field: string job = 5;
     */
    job: string; // @gotags: alias:"job"
    /**
     * @generated from protobuf field: string job_label = 6;
     */
    jobLabel: string; // @gotags: alias:"job_label"
    /**
     * @generated from protobuf field: string name = 7;
     */
    name: string;
    /**
     * @generated from protobuf field: string icon = 8;
     */
    icon: string;
    /**
     * @generated from protobuf field: string icon_color = 9;
     */
    iconColor: string; // @gotags: alias:"icon_color"
    /**
     * @generated from protobuf field: string popup = 10;
     */
    popup: string;
    /**
     * @generated from protobuf field: bool active = 11;
     */
    active: boolean;
}
/**
 * @generated from protobuf message resources.livemap.UserMarker
 */
export interface UserMarker {
    /**
     * @generated from protobuf field: float x = 1;
     */
    x: number; // @gotags: alias:"x"
    /**
     * @generated from protobuf field: float y = 2;
     */
    y: number; // @gotags: alias:"y"
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp; // @gotags: alias:"updated_at"
    /**
     * @generated from protobuf field: int32 id = 4;
     */
    id: number;
    /**
     * @generated from protobuf field: string name = 5;
     */
    name: string;
    /**
     * @generated from protobuf field: string icon = 6;
     */
    icon: string;
    /**
     * @generated from protobuf field: string icon_color = 7;
     */
    iconColor: string; // @gotags: alias:"icon_color"
    /**
     * @generated from protobuf field: resources.users.UserShort user = 8;
     */
    user?: UserShort; // @gotags: alias:"user"
}
// @generated message type with reflection information, may provide speed optimized methods
class DispatchMarker$Type extends MessageType<DispatchMarker> {
    constructor() {
        super("resources.livemap.DispatchMarker", [
            { no: 1, name: "x", kind: "scalar", T: 2 /*ScalarType.FLOAT*/ },
            { no: 2, name: "y", kind: "scalar", T: 2 /*ScalarType.FLOAT*/ },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 5, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 6, name: "job_label", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 7, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 8, name: "icon", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 9, name: "icon_color", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 10, name: "popup", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 11, name: "active", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.livemap.DispatchMarker
 */
export const DispatchMarker = new DispatchMarker$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UserMarker$Type extends MessageType<UserMarker> {
    constructor() {
        super("resources.livemap.UserMarker", [
            { no: 1, name: "x", kind: "scalar", T: 2 /*ScalarType.FLOAT*/ },
            { no: 2, name: "y", kind: "scalar", T: 2 /*ScalarType.FLOAT*/ },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 5, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 6, name: "icon", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 7, name: "icon_color", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 8, name: "user", kind: "message", T: () => UserShort }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.livemap.UserMarker
 */
export const UserMarker = new UserMarker$Type();
