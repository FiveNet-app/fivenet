// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "services/mailer/mailer.proto" (package "services.mailer", syntax proto3)
// @ts-nocheck
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { MailerService } from "./mailer";
import type { SetUserSettingsResponse } from "./mailer";
import type { SetUserSettingsRequest } from "./mailer";
import type { GetUserSettingsResponse } from "./mailer";
import type { GetUserSettingsRequest } from "./mailer";
import type { DeleteMessageResponse } from "./mailer";
import type { DeleteMessageRequest } from "./mailer";
import type { PostMessageResponse } from "./mailer";
import type { PostMessageRequest } from "./mailer";
import type { ListThreadMessagesResponse } from "./mailer";
import type { ListThreadMessagesRequest } from "./mailer";
import type { LeaveThreadResponse } from "./mailer";
import type { LeaveThreadRequest } from "./mailer";
import type { SetThreadStateResponse } from "./mailer";
import type { SetThreadStateRequest } from "./mailer";
import type { DeleteThreadResponse } from "./mailer";
import type { DeleteThreadRequest } from "./mailer";
import type { CreateThreadResponse } from "./mailer";
import type { CreateThreadRequest } from "./mailer";
import type { GetThreadResponse } from "./mailer";
import type { GetThreadRequest } from "./mailer";
import type { ListThreadsResponse } from "./mailer";
import type { ListThreadsRequest } from "./mailer";
import type { DeleteTemplateResponse } from "./mailer";
import type { DeleteTemplateRequest } from "./mailer";
import type { CreateOrUpdateTemplateResponse } from "./mailer";
import type { CreateOrUpdateTemplateRequest } from "./mailer";
import type { GetTemplateResponse } from "./mailer";
import type { GetTemplateRequest } from "./mailer";
import type { ListTemplatesResponse } from "./mailer";
import type { ListTemplatesRequest } from "./mailer";
import type { DeleteEmailResponse } from "./mailer";
import type { DeleteEmailRequest } from "./mailer";
import type { CreateOrUpdateEmailResponse } from "./mailer";
import type { CreateOrUpdateEmailRequest } from "./mailer";
import type { GetEmailResponse } from "./mailer";
import type { GetEmailRequest } from "./mailer";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { ListEmailsResponse } from "./mailer";
import type { ListEmailsRequest } from "./mailer";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service services.mailer.MailerService
 */
export interface IMailerServiceClient {
    /**
     * @perm
     *
     * @generated from protobuf rpc: ListEmails(services.mailer.ListEmailsRequest) returns (services.mailer.ListEmailsResponse);
     */
    listEmails(input: ListEmailsRequest, options?: RpcOptions): UnaryCall<ListEmailsRequest, ListEmailsResponse>;
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: GetEmail(services.mailer.GetEmailRequest) returns (services.mailer.GetEmailResponse);
     */
    getEmail(input: GetEmailRequest, options?: RpcOptions): UnaryCall<GetEmailRequest, GetEmailResponse>;
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: CreateOrUpdateEmail(services.mailer.CreateOrUpdateEmailRequest) returns (services.mailer.CreateOrUpdateEmailResponse);
     */
    createOrUpdateEmail(input: CreateOrUpdateEmailRequest, options?: RpcOptions): UnaryCall<CreateOrUpdateEmailRequest, CreateOrUpdateEmailResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: DeleteEmail(services.mailer.DeleteEmailRequest) returns (services.mailer.DeleteEmailResponse);
     */
    deleteEmail(input: DeleteEmailRequest, options?: RpcOptions): UnaryCall<DeleteEmailRequest, DeleteEmailResponse>;
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: ListTemplates(services.mailer.ListTemplatesRequest) returns (services.mailer.ListTemplatesResponse);
     */
    listTemplates(input: ListTemplatesRequest, options?: RpcOptions): UnaryCall<ListTemplatesRequest, ListTemplatesResponse>;
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: GetTemplate(services.mailer.GetTemplateRequest) returns (services.mailer.GetTemplateResponse);
     */
    getTemplate(input: GetTemplateRequest, options?: RpcOptions): UnaryCall<GetTemplateRequest, GetTemplateResponse>;
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: CreateOrUpdateTemplate(services.mailer.CreateOrUpdateTemplateRequest) returns (services.mailer.CreateOrUpdateTemplateResponse);
     */
    createOrUpdateTemplate(input: CreateOrUpdateTemplateRequest, options?: RpcOptions): UnaryCall<CreateOrUpdateTemplateRequest, CreateOrUpdateTemplateResponse>;
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: DeleteTemplate(services.mailer.DeleteTemplateRequest) returns (services.mailer.DeleteTemplateResponse);
     */
    deleteTemplate(input: DeleteTemplateRequest, options?: RpcOptions): UnaryCall<DeleteTemplateRequest, DeleteTemplateResponse>;
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: ListThreads(services.mailer.ListThreadsRequest) returns (services.mailer.ListThreadsResponse);
     */
    listThreads(input: ListThreadsRequest, options?: RpcOptions): UnaryCall<ListThreadsRequest, ListThreadsResponse>;
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: GetThread(services.mailer.GetThreadRequest) returns (services.mailer.GetThreadResponse);
     */
    getThread(input: GetThreadRequest, options?: RpcOptions): UnaryCall<GetThreadRequest, GetThreadResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: CreateThread(services.mailer.CreateThreadRequest) returns (services.mailer.CreateThreadResponse);
     */
    createThread(input: CreateThreadRequest, options?: RpcOptions): UnaryCall<CreateThreadRequest, CreateThreadResponse>;
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: DeleteThread(services.mailer.DeleteThreadRequest) returns (services.mailer.DeleteThreadResponse);
     */
    deleteThread(input: DeleteThreadRequest, options?: RpcOptions): UnaryCall<DeleteThreadRequest, DeleteThreadResponse>;
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: SetThreadState(services.mailer.SetThreadStateRequest) returns (services.mailer.SetThreadStateResponse);
     */
    setThreadState(input: SetThreadStateRequest, options?: RpcOptions): UnaryCall<SetThreadStateRequest, SetThreadStateResponse>;
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: LeaveThread(services.mailer.LeaveThreadRequest) returns (services.mailer.LeaveThreadResponse);
     */
    leaveThread(input: LeaveThreadRequest, options?: RpcOptions): UnaryCall<LeaveThreadRequest, LeaveThreadResponse>;
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: ListThreadMessages(services.mailer.ListThreadMessagesRequest) returns (services.mailer.ListThreadMessagesResponse);
     */
    listThreadMessages(input: ListThreadMessagesRequest, options?: RpcOptions): UnaryCall<ListThreadMessagesRequest, ListThreadMessagesResponse>;
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: PostMessage(services.mailer.PostMessageRequest) returns (services.mailer.PostMessageResponse);
     */
    postMessage(input: PostMessageRequest, options?: RpcOptions): UnaryCall<PostMessageRequest, PostMessageResponse>;
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: DeleteMessage(services.mailer.DeleteMessageRequest) returns (services.mailer.DeleteMessageResponse);
     */
    deleteMessage(input: DeleteMessageRequest, options?: RpcOptions): UnaryCall<DeleteMessageRequest, DeleteMessageResponse>;
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: GetUserSettings(services.mailer.GetUserSettingsRequest) returns (services.mailer.GetUserSettingsResponse);
     */
    getUserSettings(input: GetUserSettingsRequest, options?: RpcOptions): UnaryCall<GetUserSettingsRequest, GetUserSettingsResponse>;
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: SetUserSettings(services.mailer.SetUserSettingsRequest) returns (services.mailer.SetUserSettingsResponse);
     */
    setUserSettings(input: SetUserSettingsRequest, options?: RpcOptions): UnaryCall<SetUserSettingsRequest, SetUserSettingsResponse>;
}
/**
 * @generated from protobuf service services.mailer.MailerService
 */
export class MailerServiceClient implements IMailerServiceClient, ServiceInfo {
    typeName = MailerService.typeName;
    methods = MailerService.methods;
    options = MailerService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: ListEmails(services.mailer.ListEmailsRequest) returns (services.mailer.ListEmailsResponse);
     */
    listEmails(input: ListEmailsRequest, options?: RpcOptions): UnaryCall<ListEmailsRequest, ListEmailsResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListEmailsRequest, ListEmailsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: GetEmail(services.mailer.GetEmailRequest) returns (services.mailer.GetEmailResponse);
     */
    getEmail(input: GetEmailRequest, options?: RpcOptions): UnaryCall<GetEmailRequest, GetEmailResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetEmailRequest, GetEmailResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: CreateOrUpdateEmail(services.mailer.CreateOrUpdateEmailRequest) returns (services.mailer.CreateOrUpdateEmailResponse);
     */
    createOrUpdateEmail(input: CreateOrUpdateEmailRequest, options?: RpcOptions): UnaryCall<CreateOrUpdateEmailRequest, CreateOrUpdateEmailResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateOrUpdateEmailRequest, CreateOrUpdateEmailResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: DeleteEmail(services.mailer.DeleteEmailRequest) returns (services.mailer.DeleteEmailResponse);
     */
    deleteEmail(input: DeleteEmailRequest, options?: RpcOptions): UnaryCall<DeleteEmailRequest, DeleteEmailResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteEmailRequest, DeleteEmailResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: ListTemplates(services.mailer.ListTemplatesRequest) returns (services.mailer.ListTemplatesResponse);
     */
    listTemplates(input: ListTemplatesRequest, options?: RpcOptions): UnaryCall<ListTemplatesRequest, ListTemplatesResponse> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListTemplatesRequest, ListTemplatesResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: GetTemplate(services.mailer.GetTemplateRequest) returns (services.mailer.GetTemplateResponse);
     */
    getTemplate(input: GetTemplateRequest, options?: RpcOptions): UnaryCall<GetTemplateRequest, GetTemplateResponse> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetTemplateRequest, GetTemplateResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: CreateOrUpdateTemplate(services.mailer.CreateOrUpdateTemplateRequest) returns (services.mailer.CreateOrUpdateTemplateResponse);
     */
    createOrUpdateTemplate(input: CreateOrUpdateTemplateRequest, options?: RpcOptions): UnaryCall<CreateOrUpdateTemplateRequest, CreateOrUpdateTemplateResponse> {
        const method = this.methods[6], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateOrUpdateTemplateRequest, CreateOrUpdateTemplateResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: DeleteTemplate(services.mailer.DeleteTemplateRequest) returns (services.mailer.DeleteTemplateResponse);
     */
    deleteTemplate(input: DeleteTemplateRequest, options?: RpcOptions): UnaryCall<DeleteTemplateRequest, DeleteTemplateResponse> {
        const method = this.methods[7], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteTemplateRequest, DeleteTemplateResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: ListThreads(services.mailer.ListThreadsRequest) returns (services.mailer.ListThreadsResponse);
     */
    listThreads(input: ListThreadsRequest, options?: RpcOptions): UnaryCall<ListThreadsRequest, ListThreadsResponse> {
        const method = this.methods[8], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListThreadsRequest, ListThreadsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: GetThread(services.mailer.GetThreadRequest) returns (services.mailer.GetThreadResponse);
     */
    getThread(input: GetThreadRequest, options?: RpcOptions): UnaryCall<GetThreadRequest, GetThreadResponse> {
        const method = this.methods[9], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetThreadRequest, GetThreadResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: CreateThread(services.mailer.CreateThreadRequest) returns (services.mailer.CreateThreadResponse);
     */
    createThread(input: CreateThreadRequest, options?: RpcOptions): UnaryCall<CreateThreadRequest, CreateThreadResponse> {
        const method = this.methods[10], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateThreadRequest, CreateThreadResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: DeleteThread(services.mailer.DeleteThreadRequest) returns (services.mailer.DeleteThreadResponse);
     */
    deleteThread(input: DeleteThreadRequest, options?: RpcOptions): UnaryCall<DeleteThreadRequest, DeleteThreadResponse> {
        const method = this.methods[11], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteThreadRequest, DeleteThreadResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: SetThreadState(services.mailer.SetThreadStateRequest) returns (services.mailer.SetThreadStateResponse);
     */
    setThreadState(input: SetThreadStateRequest, options?: RpcOptions): UnaryCall<SetThreadStateRequest, SetThreadStateResponse> {
        const method = this.methods[12], opt = this._transport.mergeOptions(options);
        return stackIntercept<SetThreadStateRequest, SetThreadStateResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: LeaveThread(services.mailer.LeaveThreadRequest) returns (services.mailer.LeaveThreadResponse);
     */
    leaveThread(input: LeaveThreadRequest, options?: RpcOptions): UnaryCall<LeaveThreadRequest, LeaveThreadResponse> {
        const method = this.methods[13], opt = this._transport.mergeOptions(options);
        return stackIntercept<LeaveThreadRequest, LeaveThreadResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: ListThreadMessages(services.mailer.ListThreadMessagesRequest) returns (services.mailer.ListThreadMessagesResponse);
     */
    listThreadMessages(input: ListThreadMessagesRequest, options?: RpcOptions): UnaryCall<ListThreadMessagesRequest, ListThreadMessagesResponse> {
        const method = this.methods[14], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListThreadMessagesRequest, ListThreadMessagesResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: PostMessage(services.mailer.PostMessageRequest) returns (services.mailer.PostMessageResponse);
     */
    postMessage(input: PostMessageRequest, options?: RpcOptions): UnaryCall<PostMessageRequest, PostMessageResponse> {
        const method = this.methods[15], opt = this._transport.mergeOptions(options);
        return stackIntercept<PostMessageRequest, PostMessageResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: DeleteMessage(services.mailer.DeleteMessageRequest) returns (services.mailer.DeleteMessageResponse);
     */
    deleteMessage(input: DeleteMessageRequest, options?: RpcOptions): UnaryCall<DeleteMessageRequest, DeleteMessageResponse> {
        const method = this.methods[16], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteMessageRequest, DeleteMessageResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: GetUserSettings(services.mailer.GetUserSettingsRequest) returns (services.mailer.GetUserSettingsResponse);
     */
    getUserSettings(input: GetUserSettingsRequest, options?: RpcOptions): UnaryCall<GetUserSettingsRequest, GetUserSettingsResponse> {
        const method = this.methods[17], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetUserSettingsRequest, GetUserSettingsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=ListEmails
     *
     * @generated from protobuf rpc: SetUserSettings(services.mailer.SetUserSettingsRequest) returns (services.mailer.SetUserSettingsResponse);
     */
    setUserSettings(input: SetUserSettingsRequest, options?: RpcOptions): UnaryCall<SetUserSettingsRequest, SetUserSettingsResponse> {
        const method = this.methods[18], opt = this._transport.mergeOptions(options);
        return stackIntercept<SetUserSettingsRequest, SetUserSettingsResponse>("unary", this._transport, method, opt, input);
    }
}
