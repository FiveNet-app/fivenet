// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "services/citizenstore/citizenstore.proto" (package "services.citizenstore", syntax proto3)
// @ts-nocheck
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { CitizenStoreService } from "./citizenstore";
import type { ManageCitizenLabelsResponse } from "./citizenstore";
import type { ManageCitizenLabelsRequest } from "./citizenstore";
import type { SetProfilePictureResponse } from "./citizenstore";
import type { SetProfilePictureRequest } from "./citizenstore";
import type { SetUserPropsResponse } from "./citizenstore";
import type { SetUserPropsRequest } from "./citizenstore";
import type { ListUserActivityResponse } from "./citizenstore";
import type { ListUserActivityRequest } from "./citizenstore";
import type { GetUserResponse } from "./citizenstore";
import type { GetUserRequest } from "./citizenstore";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { ListCitizensResponse } from "./citizenstore";
import type { ListCitizensRequest } from "./citizenstore";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service services.citizenstore.CitizenStoreService
 */
export interface ICitizenStoreServiceClient {
    /**
     * @perm: Attrs=Fields/StringList:[]string{"PhoneNumber", "Licenses", "UserProps.Wanted", "UserProps.Job", "UserProps.TrafficInfractionPoints", "UserProps.OpenFines", "UserProps.BloodType", "UserProps.MugShot", "UserProps.Labels", "UserProps.Email"}
     *
     * @generated from protobuf rpc: ListCitizens(services.citizenstore.ListCitizensRequest) returns (services.citizenstore.ListCitizensResponse);
     */
    listCitizens(input: ListCitizensRequest, options?: RpcOptions): UnaryCall<ListCitizensRequest, ListCitizensResponse>;
    /**
     * @perm: Attrs=Jobs/JobGradeList
     *
     * @generated from protobuf rpc: GetUser(services.citizenstore.GetUserRequest) returns (services.citizenstore.GetUserResponse);
     */
    getUser(input: GetUserRequest, options?: RpcOptions): UnaryCall<GetUserRequest, GetUserResponse>;
    /**
     * @perm: Attrs=Fields/StringList:[]string{"SourceUser", "Own"}
     *
     * @generated from protobuf rpc: ListUserActivity(services.citizenstore.ListUserActivityRequest) returns (services.citizenstore.ListUserActivityResponse);
     */
    listUserActivity(input: ListUserActivityRequest, options?: RpcOptions): UnaryCall<ListUserActivityRequest, ListUserActivityResponse>;
    /**
     * @perm: Attrs=Fields/StringList:[]string{"Wanted", "Job", "TrafficInfractionPoints", "MugShot", "Labels"}
     *
     * @generated from protobuf rpc: SetUserProps(services.citizenstore.SetUserPropsRequest) returns (services.citizenstore.SetUserPropsResponse);
     */
    setUserProps(input: SetUserPropsRequest, options?: RpcOptions): UnaryCall<SetUserPropsRequest, SetUserPropsResponse>;
    /**
     * @perm: Name=Any
     *
     * @generated from protobuf rpc: SetProfilePicture(services.citizenstore.SetProfilePictureRequest) returns (services.citizenstore.SetProfilePictureResponse);
     */
    setProfilePicture(input: SetProfilePictureRequest, options?: RpcOptions): UnaryCall<SetProfilePictureRequest, SetProfilePictureResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: ManageCitizenLabels(services.citizenstore.ManageCitizenLabelsRequest) returns (services.citizenstore.ManageCitizenLabelsResponse);
     */
    manageCitizenLabels(input: ManageCitizenLabelsRequest, options?: RpcOptions): UnaryCall<ManageCitizenLabelsRequest, ManageCitizenLabelsResponse>;
}
/**
 * @generated from protobuf service services.citizenstore.CitizenStoreService
 */
export class CitizenStoreServiceClient implements ICitizenStoreServiceClient, ServiceInfo {
    typeName = CitizenStoreService.typeName;
    methods = CitizenStoreService.methods;
    options = CitizenStoreService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @perm: Attrs=Fields/StringList:[]string{"PhoneNumber", "Licenses", "UserProps.Wanted", "UserProps.Job", "UserProps.TrafficInfractionPoints", "UserProps.OpenFines", "UserProps.BloodType", "UserProps.MugShot", "UserProps.Labels", "UserProps.Email"}
     *
     * @generated from protobuf rpc: ListCitizens(services.citizenstore.ListCitizensRequest) returns (services.citizenstore.ListCitizensResponse);
     */
    listCitizens(input: ListCitizensRequest, options?: RpcOptions): UnaryCall<ListCitizensRequest, ListCitizensResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListCitizensRequest, ListCitizensResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Attrs=Jobs/JobGradeList
     *
     * @generated from protobuf rpc: GetUser(services.citizenstore.GetUserRequest) returns (services.citizenstore.GetUserResponse);
     */
    getUser(input: GetUserRequest, options?: RpcOptions): UnaryCall<GetUserRequest, GetUserResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetUserRequest, GetUserResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Attrs=Fields/StringList:[]string{"SourceUser", "Own"}
     *
     * @generated from protobuf rpc: ListUserActivity(services.citizenstore.ListUserActivityRequest) returns (services.citizenstore.ListUserActivityResponse);
     */
    listUserActivity(input: ListUserActivityRequest, options?: RpcOptions): UnaryCall<ListUserActivityRequest, ListUserActivityResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListUserActivityRequest, ListUserActivityResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Attrs=Fields/StringList:[]string{"Wanted", "Job", "TrafficInfractionPoints", "MugShot", "Labels"}
     *
     * @generated from protobuf rpc: SetUserProps(services.citizenstore.SetUserPropsRequest) returns (services.citizenstore.SetUserPropsResponse);
     */
    setUserProps(input: SetUserPropsRequest, options?: RpcOptions): UnaryCall<SetUserPropsRequest, SetUserPropsResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<SetUserPropsRequest, SetUserPropsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=Any
     *
     * @generated from protobuf rpc: SetProfilePicture(services.citizenstore.SetProfilePictureRequest) returns (services.citizenstore.SetProfilePictureResponse);
     */
    setProfilePicture(input: SetProfilePictureRequest, options?: RpcOptions): UnaryCall<SetProfilePictureRequest, SetProfilePictureResponse> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<SetProfilePictureRequest, SetProfilePictureResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: ManageCitizenLabels(services.citizenstore.ManageCitizenLabelsRequest) returns (services.citizenstore.ManageCitizenLabelsResponse);
     */
    manageCitizenLabels(input: ManageCitizenLabelsRequest, options?: RpcOptions): UnaryCall<ManageCitizenLabelsRequest, ManageCitizenLabelsResponse> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<ManageCitizenLabelsRequest, ManageCitizenLabelsResponse>("unary", this._transport, method, opt, input);
    }
}
