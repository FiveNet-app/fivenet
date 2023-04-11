// source: resources/rector/audit.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global =
    (typeof globalThis !== 'undefined' && globalThis) ||
    (typeof window !== 'undefined' && window) ||
    (typeof global !== 'undefined' && global) ||
    (typeof self !== 'undefined' && self) ||
    (function () { return this; }).call(null) ||
    Function('return this')();

var resources_timestamp_timestamp_pb = require('../../resources/timestamp/timestamp_pb.js');
goog.object.extend(proto, resources_timestamp_timestamp_pb);
var resources_users_users_pb = require('../../resources/users/users_pb.js');
goog.object.extend(proto, resources_users_users_pb);
goog.exportSymbol('proto.resources.rector.AuditEntry', null, global);
goog.exportSymbol('proto.resources.rector.EVENT_TYPE', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.resources.rector.AuditEntry = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.resources.rector.AuditEntry, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.resources.rector.AuditEntry.displayName = 'proto.resources.rector.AuditEntry';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.resources.rector.AuditEntry.prototype.toObject = function(opt_includeInstance) {
  return proto.resources.rector.AuditEntry.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.resources.rector.AuditEntry} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.resources.rector.AuditEntry.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    createdAt: (f = msg.getCreatedAt()) && resources_timestamp_timestamp_pb.Timestamp.toObject(includeInstance, f),
    userId: jspb.Message.getFieldWithDefault(msg, 3, 0),
    user: (f = msg.getUser()) && resources_users_users_pb.UserShort.toObject(includeInstance, f),
    service: jspb.Message.getFieldWithDefault(msg, 5, ""),
    method: jspb.Message.getFieldWithDefault(msg, 6, ""),
    state: jspb.Message.getFieldWithDefault(msg, 7, 0),
    data: jspb.Message.getFieldWithDefault(msg, 8, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.resources.rector.AuditEntry}
 */
proto.resources.rector.AuditEntry.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.resources.rector.AuditEntry;
  return proto.resources.rector.AuditEntry.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.resources.rector.AuditEntry} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.resources.rector.AuditEntry}
 */
proto.resources.rector.AuditEntry.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setId(value);
      break;
    case 2:
      var value = new resources_timestamp_timestamp_pb.Timestamp;
      reader.readMessage(value,resources_timestamp_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setCreatedAt(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setUserId(value);
      break;
    case 4:
      var value = new resources_users_users_pb.UserShort;
      reader.readMessage(value,resources_users_users_pb.UserShort.deserializeBinaryFromReader);
      msg.setUser(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setService(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setMethod(value);
      break;
    case 7:
      var value = /** @type {!proto.resources.rector.EVENT_TYPE} */ (reader.readEnum());
      msg.setState(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setData(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.resources.rector.AuditEntry.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.resources.rector.AuditEntry.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.resources.rector.AuditEntry} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.resources.rector.AuditEntry.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeUint64(
      1,
      f
    );
  }
  f = message.getCreatedAt();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      resources_timestamp_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getUserId();
  if (f !== 0) {
    writer.writeUint64(
      3,
      f
    );
  }
  f = message.getUser();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      resources_users_users_pb.UserShort.serializeBinaryToWriter
    );
  }
  f = message.getService();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getMethod();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getState();
  if (f !== 0.0) {
    writer.writeEnum(
      7,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 8));
  if (f != null) {
    writer.writeString(
      8,
      f
    );
  }
};


/**
 * optional uint64 id = 1;
 * @return {number}
 */
proto.resources.rector.AuditEntry.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.resources.rector.AuditEntry} returns this
 */
proto.resources.rector.AuditEntry.prototype.setId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional resources.timestamp.Timestamp created_at = 2;
 * @return {?proto.resources.timestamp.Timestamp}
 */
proto.resources.rector.AuditEntry.prototype.getCreatedAt = function() {
  return /** @type{?proto.resources.timestamp.Timestamp} */ (
    jspb.Message.getWrapperField(this, resources_timestamp_timestamp_pb.Timestamp, 2));
};


/**
 * @param {?proto.resources.timestamp.Timestamp|undefined} value
 * @return {!proto.resources.rector.AuditEntry} returns this
*/
proto.resources.rector.AuditEntry.prototype.setCreatedAt = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.resources.rector.AuditEntry} returns this
 */
proto.resources.rector.AuditEntry.prototype.clearCreatedAt = function() {
  return this.setCreatedAt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.resources.rector.AuditEntry.prototype.hasCreatedAt = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional uint64 user_id = 3;
 * @return {number}
 */
proto.resources.rector.AuditEntry.prototype.getUserId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.resources.rector.AuditEntry} returns this
 */
proto.resources.rector.AuditEntry.prototype.setUserId = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional resources.users.UserShort user = 4;
 * @return {?proto.resources.users.UserShort}
 */
proto.resources.rector.AuditEntry.prototype.getUser = function() {
  return /** @type{?proto.resources.users.UserShort} */ (
    jspb.Message.getWrapperField(this, resources_users_users_pb.UserShort, 4));
};


/**
 * @param {?proto.resources.users.UserShort|undefined} value
 * @return {!proto.resources.rector.AuditEntry} returns this
*/
proto.resources.rector.AuditEntry.prototype.setUser = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.resources.rector.AuditEntry} returns this
 */
proto.resources.rector.AuditEntry.prototype.clearUser = function() {
  return this.setUser(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.resources.rector.AuditEntry.prototype.hasUser = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional string service = 5;
 * @return {string}
 */
proto.resources.rector.AuditEntry.prototype.getService = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.resources.rector.AuditEntry} returns this
 */
proto.resources.rector.AuditEntry.prototype.setService = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string method = 6;
 * @return {string}
 */
proto.resources.rector.AuditEntry.prototype.getMethod = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.resources.rector.AuditEntry} returns this
 */
proto.resources.rector.AuditEntry.prototype.setMethod = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional EVENT_TYPE state = 7;
 * @return {!proto.resources.rector.EVENT_TYPE}
 */
proto.resources.rector.AuditEntry.prototype.getState = function() {
  return /** @type {!proto.resources.rector.EVENT_TYPE} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/**
 * @param {!proto.resources.rector.EVENT_TYPE} value
 * @return {!proto.resources.rector.AuditEntry} returns this
 */
proto.resources.rector.AuditEntry.prototype.setState = function(value) {
  return jspb.Message.setProto3EnumField(this, 7, value);
};


/**
 * optional string data = 8;
 * @return {string}
 */
proto.resources.rector.AuditEntry.prototype.getData = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.resources.rector.AuditEntry} returns this
 */
proto.resources.rector.AuditEntry.prototype.setData = function(value) {
  return jspb.Message.setField(this, 8, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.resources.rector.AuditEntry} returns this
 */
proto.resources.rector.AuditEntry.prototype.clearData = function() {
  return jspb.Message.setField(this, 8, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.resources.rector.AuditEntry.prototype.hasData = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * @enum {number}
 */
proto.resources.rector.EVENT_TYPE = {
  UNKNOWN: 0,
  CREATE: 1,
  UPDATE: 2,
  DELETE: 3
};

goog.object.extend(exports, proto.resources.rector);
