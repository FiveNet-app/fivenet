import * as jspb from 'google-protobuf'

import * as resources_timestamp_timestamp_pb from '../../resources/timestamp/timestamp_pb';


export class UserShort extends jspb.Message {
  getUserId(): number;
  setUserId(value: number): UserShort;

  getIdentifier(): string;
  setIdentifier(value: string): UserShort;

  getJob(): string;
  setJob(value: string): UserShort;

  getJobLabel(): string;
  setJobLabel(value: string): UserShort;

  getJobGrade(): number;
  setJobGrade(value: number): UserShort;

  getJobGradeLabel(): string;
  setJobGradeLabel(value: string): UserShort;

  getFirstname(): string;
  setFirstname(value: string): UserShort;

  getLastname(): string;
  setLastname(value: string): UserShort;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserShort.AsObject;
  static toObject(includeInstance: boolean, msg: UserShort): UserShort.AsObject;
  static serializeBinaryToWriter(message: UserShort, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserShort;
  static deserializeBinaryFromReader(message: UserShort, reader: jspb.BinaryReader): UserShort;
}

export namespace UserShort {
  export type AsObject = {
    userId: number,
    identifier: string,
    job: string,
    jobLabel: string,
    jobGrade: number,
    jobGradeLabel: string,
    firstname: string,
    lastname: string,
  }
}

export class UserShortNI extends jspb.Message {
  getUserId(): number;
  setUserId(value: number): UserShortNI;

  getIdentifier(): string;
  setIdentifier(value: string): UserShortNI;

  getJob(): string;
  setJob(value: string): UserShortNI;

  getJobLabel(): string;
  setJobLabel(value: string): UserShortNI;

  getJobGrade(): number;
  setJobGrade(value: number): UserShortNI;

  getJobGradeLabel(): string;
  setJobGradeLabel(value: string): UserShortNI;

  getFirstname(): string;
  setFirstname(value: string): UserShortNI;

  getLastname(): string;
  setLastname(value: string): UserShortNI;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserShortNI.AsObject;
  static toObject(includeInstance: boolean, msg: UserShortNI): UserShortNI.AsObject;
  static serializeBinaryToWriter(message: UserShortNI, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserShortNI;
  static deserializeBinaryFromReader(message: UserShortNI, reader: jspb.BinaryReader): UserShortNI;
}

export namespace UserShortNI {
  export type AsObject = {
    userId: number,
    identifier: string,
    job: string,
    jobLabel: string,
    jobGrade: number,
    jobGradeLabel: string,
    firstname: string,
    lastname: string,
  }
}

export class User extends jspb.Message {
  getUserId(): number;
  setUserId(value: number): User;

  getIdentifier(): string;
  setIdentifier(value: string): User;

  getJob(): string;
  setJob(value: string): User;

  getJobLabel(): string;
  setJobLabel(value: string): User;

  getJobGrade(): number;
  setJobGrade(value: number): User;

  getJobGradeLabel(): string;
  setJobGradeLabel(value: string): User;

  getFirstname(): string;
  setFirstname(value: string): User;

  getLastname(): string;
  setLastname(value: string): User;

  getDateofbirth(): string;
  setDateofbirth(value: string): User;

  getSex(): string;
  setSex(value: string): User;

  getHeight(): string;
  setHeight(value: string): User;

  getPhoneNumber(): string;
  setPhoneNumber(value: string): User;

  getVisum(): number;
  setVisum(value: number): User;
  hasVisum(): boolean;
  clearVisum(): User;

  getPlaytime(): number;
  setPlaytime(value: number): User;
  hasPlaytime(): boolean;
  clearPlaytime(): User;

  getProps(): UserProps | undefined;
  setProps(value?: UserProps): User;
  hasProps(): boolean;
  clearProps(): User;

  getLicensesList(): Array<License>;
  setLicensesList(value: Array<License>): User;
  clearLicensesList(): User;
  addLicenses(value?: License, index?: number): License;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    userId: number,
    identifier: string,
    job: string,
    jobLabel: string,
    jobGrade: number,
    jobGradeLabel: string,
    firstname: string,
    lastname: string,
    dateofbirth: string,
    sex: string,
    height: string,
    phoneNumber: string,
    visum?: number,
    playtime?: number,
    props?: UserProps.AsObject,
    licensesList: Array<License.AsObject>,
  }

  export enum VisumCase { 
    _VISUM_NOT_SET = 0,
    VISUM = 13,
  }

  export enum PlaytimeCase { 
    _PLAYTIME_NOT_SET = 0,
    PLAYTIME = 14,
  }
}

export class License extends jspb.Message {
  getType(): string;
  setType(value: string): License;

  getLabel(): string;
  setLabel(value: string): License;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): License.AsObject;
  static toObject(includeInstance: boolean, msg: License): License.AsObject;
  static serializeBinaryToWriter(message: License, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): License;
  static deserializeBinaryFromReader(message: License, reader: jspb.BinaryReader): License;
}

export namespace License {
  export type AsObject = {
    type: string,
    label: string,
  }
}

export class UserProps extends jspb.Message {
  getUserId(): number;
  setUserId(value: number): UserProps;

  getWanted(): boolean;
  setWanted(value: boolean): UserProps;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserProps.AsObject;
  static toObject(includeInstance: boolean, msg: UserProps): UserProps.AsObject;
  static serializeBinaryToWriter(message: UserProps, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserProps;
  static deserializeBinaryFromReader(message: UserProps, reader: jspb.BinaryReader): UserProps;
}

export namespace UserProps {
  export type AsObject = {
    userId: number,
    wanted: boolean,
  }
}

export class UserActivity extends jspb.Message {
  getId(): number;
  setId(value: number): UserActivity;

  getType(): USER_ACTIVITY_TYPE;
  setType(value: USER_ACTIVITY_TYPE): UserActivity;

  getCreatedAt(): resources_timestamp_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: resources_timestamp_timestamp_pb.Timestamp): UserActivity;
  hasCreatedAt(): boolean;
  clearCreatedAt(): UserActivity;

  getSourceUser(): UserShort | undefined;
  setSourceUser(value?: UserShort): UserActivity;
  hasSourceUser(): boolean;
  clearSourceUser(): UserActivity;

  getTargetUser(): UserShort | undefined;
  setTargetUser(value?: UserShort): UserActivity;
  hasTargetUser(): boolean;
  clearTargetUser(): UserActivity;

  getKey(): string;
  setKey(value: string): UserActivity;

  getOldvalue(): string;
  setOldvalue(value: string): UserActivity;

  getNewvalue(): string;
  setNewvalue(value: string): UserActivity;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserActivity.AsObject;
  static toObject(includeInstance: boolean, msg: UserActivity): UserActivity.AsObject;
  static serializeBinaryToWriter(message: UserActivity, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserActivity;
  static deserializeBinaryFromReader(message: UserActivity, reader: jspb.BinaryReader): UserActivity;
}

export namespace UserActivity {
  export type AsObject = {
    id: number,
    type: USER_ACTIVITY_TYPE,
    createdAt?: resources_timestamp_timestamp_pb.Timestamp.AsObject,
    sourceUser?: UserShort.AsObject,
    targetUser?: UserShort.AsObject,
    key: string,
    oldvalue: string,
    newvalue: string,
  }
}

export enum USER_ACTIVITY_TYPE { 
  CHANGED = 0,
  MENTIONED = 1,
  CREATED = 2,
}
