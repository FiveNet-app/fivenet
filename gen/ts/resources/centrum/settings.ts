// @generated by protobuf-ts 2.9.3 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "resources/centrum/settings.proto" (package "resources.centrum", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message resources.centrum.Settings
 */
export interface Settings {
    /**
     * @generated from protobuf field: string job = 1;
     */
    job: string;
    /**
     * @generated from protobuf field: bool enabled = 2;
     */
    enabled: boolean;
    /**
     * @generated from protobuf field: resources.centrum.CentrumMode mode = 3;
     */
    mode: CentrumMode;
    /**
     * @generated from protobuf field: resources.centrum.CentrumMode fallback_mode = 4;
     */
    fallbackMode: CentrumMode;
}
/**
 * @generated from protobuf enum resources.centrum.CentrumMode
 */
export enum CentrumMode {
    /**
     * @generated from protobuf enum value: CENTRUM_MODE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: CENTRUM_MODE_MANUAL = 1;
     */
    MANUAL = 1,
    /**
     * @generated from protobuf enum value: CENTRUM_MODE_CENTRAL_COMMAND = 2;
     */
    CENTRAL_COMMAND = 2,
    /**
     * @generated from protobuf enum value: CENTRUM_MODE_AUTO_ROUND_ROBIN = 3;
     */
    AUTO_ROUND_ROBIN = 3,
    /**
     * @generated from protobuf enum value: CENTRUM_MODE_SIMPLIFIED = 4;
     */
    SIMPLIFIED = 4
}
// @generated message type with reflection information, may provide speed optimized methods
class Settings$Type extends MessageType<Settings> {
    constructor() {
        super("resources.centrum.Settings", [
            { no: 1, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 2, name: "enabled", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "mode", kind: "enum", T: () => ["resources.centrum.CentrumMode", CentrumMode, "CENTRUM_MODE_"], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 4, name: "fallback_mode", kind: "enum", T: () => ["resources.centrum.CentrumMode", CentrumMode, "CENTRUM_MODE_"], options: { "validate.rules": { enum: { definedOnly: true } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.centrum.Settings
 */
export const Settings = new Settings$Type();
