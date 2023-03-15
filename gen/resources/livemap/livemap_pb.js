// source: resources/livemap/livemap.proto
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
goog.exportSymbol('proto.resources.livemap.DispatchMarker', null, global);
goog.exportSymbol('proto.resources.livemap.Marker', null, global);
goog.exportSymbol('proto.resources.livemap.UserMarker', null, global);
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
proto.resources.livemap.Marker = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.resources.livemap.Marker, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.resources.livemap.Marker.displayName = 'proto.resources.livemap.Marker';
}
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
proto.resources.livemap.UserMarker = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.resources.livemap.UserMarker, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.resources.livemap.UserMarker.displayName = 'proto.resources.livemap.UserMarker';
}
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
proto.resources.livemap.DispatchMarker = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.resources.livemap.DispatchMarker, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.resources.livemap.DispatchMarker.displayName = 'proto.resources.livemap.DispatchMarker';
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
proto.resources.livemap.Marker.prototype.toObject = function(opt_includeInstance) {
  return proto.resources.livemap.Marker.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.resources.livemap.Marker} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.resources.livemap.Marker.toObject = function(includeInstance, msg) {
  var f, obj = {
    x: jspb.Message.getFloatingPointFieldWithDefault(msg, 1, 0.0),
    y: jspb.Message.getFloatingPointFieldWithDefault(msg, 2, 0.0),
    updatedAt: (f = msg.getUpdatedAt()) && resources_timestamp_timestamp_pb.Timestamp.toObject(includeInstance, f)
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
 * @return {!proto.resources.livemap.Marker}
 */
proto.resources.livemap.Marker.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.resources.livemap.Marker;
  return proto.resources.livemap.Marker.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.resources.livemap.Marker} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.resources.livemap.Marker}
 */
proto.resources.livemap.Marker.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setX(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setY(value);
      break;
    case 3:
      var value = new resources_timestamp_timestamp_pb.Timestamp;
      reader.readMessage(value,resources_timestamp_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setUpdatedAt(value);
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
proto.resources.livemap.Marker.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.resources.livemap.Marker.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.resources.livemap.Marker} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.resources.livemap.Marker.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getX();
  if (f !== 0.0) {
    writer.writeFloat(
      1,
      f
    );
  }
  f = message.getY();
  if (f !== 0.0) {
    writer.writeFloat(
      2,
      f
    );
  }
  f = message.getUpdatedAt();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      resources_timestamp_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
};


/**
 * optional float x = 1;
 * @return {number}
 */
proto.resources.livemap.Marker.prototype.getX = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 1, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.resources.livemap.Marker} returns this
 */
proto.resources.livemap.Marker.prototype.setX = function(value) {
  return jspb.Message.setProto3FloatField(this, 1, value);
};


/**
 * optional float y = 2;
 * @return {number}
 */
proto.resources.livemap.Marker.prototype.getY = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 2, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.resources.livemap.Marker} returns this
 */
proto.resources.livemap.Marker.prototype.setY = function(value) {
  return jspb.Message.setProto3FloatField(this, 2, value);
};


/**
 * optional resources.timestamp.Timestamp updated_at = 3;
 * @return {?proto.resources.timestamp.Timestamp}
 */
proto.resources.livemap.Marker.prototype.getUpdatedAt = function() {
  return /** @type{?proto.resources.timestamp.Timestamp} */ (
    jspb.Message.getWrapperField(this, resources_timestamp_timestamp_pb.Timestamp, 3));
};


/**
 * @param {?proto.resources.timestamp.Timestamp|undefined} value
 * @return {!proto.resources.livemap.Marker} returns this
*/
proto.resources.livemap.Marker.prototype.setUpdatedAt = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.resources.livemap.Marker} returns this
 */
proto.resources.livemap.Marker.prototype.clearUpdatedAt = function() {
  return this.setUpdatedAt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.resources.livemap.Marker.prototype.hasUpdatedAt = function() {
  return jspb.Message.getField(this, 3) != null;
};





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
proto.resources.livemap.UserMarker.prototype.toObject = function(opt_includeInstance) {
  return proto.resources.livemap.UserMarker.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.resources.livemap.UserMarker} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.resources.livemap.UserMarker.toObject = function(includeInstance, msg) {
  var f, obj = {
    x: jspb.Message.getFloatingPointFieldWithDefault(msg, 1, 0.0),
    y: jspb.Message.getFloatingPointFieldWithDefault(msg, 2, 0.0),
    updatedAt: (f = msg.getUpdatedAt()) && resources_timestamp_timestamp_pb.Timestamp.toObject(includeInstance, f),
    job: jspb.Message.getFieldWithDefault(msg, 4, ""),
    userId: jspb.Message.getFieldWithDefault(msg, 5, 0),
    user: (f = msg.getUser()) && resources_users_users_pb.UserShort.toObject(includeInstance, f),
    name: jspb.Message.getFieldWithDefault(msg, 7, ""),
    icon: jspb.Message.getFieldWithDefault(msg, 8, ""),
    popup: jspb.Message.getFieldWithDefault(msg, 9, ""),
    link: jspb.Message.getFieldWithDefault(msg, 10, "")
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
 * @return {!proto.resources.livemap.UserMarker}
 */
proto.resources.livemap.UserMarker.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.resources.livemap.UserMarker;
  return proto.resources.livemap.UserMarker.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.resources.livemap.UserMarker} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.resources.livemap.UserMarker}
 */
proto.resources.livemap.UserMarker.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setX(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setY(value);
      break;
    case 3:
      var value = new resources_timestamp_timestamp_pb.Timestamp;
      reader.readMessage(value,resources_timestamp_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setUpdatedAt(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setJob(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setUserId(value);
      break;
    case 6:
      var value = new resources_users_users_pb.UserShort;
      reader.readMessage(value,resources_users_users_pb.UserShort.deserializeBinaryFromReader);
      msg.setUser(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setIcon(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setPopup(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setLink(value);
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
proto.resources.livemap.UserMarker.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.resources.livemap.UserMarker.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.resources.livemap.UserMarker} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.resources.livemap.UserMarker.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getX();
  if (f !== 0.0) {
    writer.writeFloat(
      1,
      f
    );
  }
  f = message.getY();
  if (f !== 0.0) {
    writer.writeFloat(
      2,
      f
    );
  }
  f = message.getUpdatedAt();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      resources_timestamp_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getJob();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getUserId();
  if (f !== 0) {
    writer.writeInt32(
      5,
      f
    );
  }
  f = message.getUser();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      resources_users_users_pb.UserShort.serializeBinaryToWriter
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getIcon();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getPopup();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
  f = message.getLink();
  if (f.length > 0) {
    writer.writeString(
      10,
      f
    );
  }
};


/**
 * optional float x = 1;
 * @return {number}
 */
proto.resources.livemap.UserMarker.prototype.getX = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 1, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.resources.livemap.UserMarker} returns this
 */
proto.resources.livemap.UserMarker.prototype.setX = function(value) {
  return jspb.Message.setProto3FloatField(this, 1, value);
};


/**
 * optional float y = 2;
 * @return {number}
 */
proto.resources.livemap.UserMarker.prototype.getY = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 2, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.resources.livemap.UserMarker} returns this
 */
proto.resources.livemap.UserMarker.prototype.setY = function(value) {
  return jspb.Message.setProto3FloatField(this, 2, value);
};


/**
 * optional resources.timestamp.Timestamp updated_at = 3;
 * @return {?proto.resources.timestamp.Timestamp}
 */
proto.resources.livemap.UserMarker.prototype.getUpdatedAt = function() {
  return /** @type{?proto.resources.timestamp.Timestamp} */ (
    jspb.Message.getWrapperField(this, resources_timestamp_timestamp_pb.Timestamp, 3));
};


/**
 * @param {?proto.resources.timestamp.Timestamp|undefined} value
 * @return {!proto.resources.livemap.UserMarker} returns this
*/
proto.resources.livemap.UserMarker.prototype.setUpdatedAt = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.resources.livemap.UserMarker} returns this
 */
proto.resources.livemap.UserMarker.prototype.clearUpdatedAt = function() {
  return this.setUpdatedAt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.resources.livemap.UserMarker.prototype.hasUpdatedAt = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional string job = 4;
 * @return {string}
 */
proto.resources.livemap.UserMarker.prototype.getJob = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.resources.livemap.UserMarker} returns this
 */
proto.resources.livemap.UserMarker.prototype.setJob = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional int32 user_id = 5;
 * @return {number}
 */
proto.resources.livemap.UserMarker.prototype.getUserId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.resources.livemap.UserMarker} returns this
 */
proto.resources.livemap.UserMarker.prototype.setUserId = function(value) {
  return jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional resources.users.UserShort user = 6;
 * @return {?proto.resources.users.UserShort}
 */
proto.resources.livemap.UserMarker.prototype.getUser = function() {
  return /** @type{?proto.resources.users.UserShort} */ (
    jspb.Message.getWrapperField(this, resources_users_users_pb.UserShort, 6));
};


/**
 * @param {?proto.resources.users.UserShort|undefined} value
 * @return {!proto.resources.livemap.UserMarker} returns this
*/
proto.resources.livemap.UserMarker.prototype.setUser = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.resources.livemap.UserMarker} returns this
 */
proto.resources.livemap.UserMarker.prototype.clearUser = function() {
  return this.setUser(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.resources.livemap.UserMarker.prototype.hasUser = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional string name = 7;
 * @return {string}
 */
proto.resources.livemap.UserMarker.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.resources.livemap.UserMarker} returns this
 */
proto.resources.livemap.UserMarker.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string icon = 8;
 * @return {string}
 */
proto.resources.livemap.UserMarker.prototype.getIcon = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.resources.livemap.UserMarker} returns this
 */
proto.resources.livemap.UserMarker.prototype.setIcon = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional string popup = 9;
 * @return {string}
 */
proto.resources.livemap.UserMarker.prototype.getPopup = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/**
 * @param {string} value
 * @return {!proto.resources.livemap.UserMarker} returns this
 */
proto.resources.livemap.UserMarker.prototype.setPopup = function(value) {
  return jspb.Message.setProto3StringField(this, 9, value);
};


/**
 * optional string link = 10;
 * @return {string}
 */
proto.resources.livemap.UserMarker.prototype.getLink = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/**
 * @param {string} value
 * @return {!proto.resources.livemap.UserMarker} returns this
 */
proto.resources.livemap.UserMarker.prototype.setLink = function(value) {
  return jspb.Message.setProto3StringField(this, 10, value);
};





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
proto.resources.livemap.DispatchMarker.prototype.toObject = function(opt_includeInstance) {
  return proto.resources.livemap.DispatchMarker.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.resources.livemap.DispatchMarker} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.resources.livemap.DispatchMarker.toObject = function(includeInstance, msg) {
  var f, obj = {
    x: jspb.Message.getFloatingPointFieldWithDefault(msg, 1, 0.0),
    y: jspb.Message.getFloatingPointFieldWithDefault(msg, 2, 0.0),
    updatedAt: (f = msg.getUpdatedAt()) && resources_timestamp_timestamp_pb.Timestamp.toObject(includeInstance, f),
    name: jspb.Message.getFieldWithDefault(msg, 4, ""),
    icon: jspb.Message.getFieldWithDefault(msg, 5, ""),
    popup: jspb.Message.getFieldWithDefault(msg, 6, ""),
    link: jspb.Message.getFieldWithDefault(msg, 7, "")
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
 * @return {!proto.resources.livemap.DispatchMarker}
 */
proto.resources.livemap.DispatchMarker.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.resources.livemap.DispatchMarker;
  return proto.resources.livemap.DispatchMarker.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.resources.livemap.DispatchMarker} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.resources.livemap.DispatchMarker}
 */
proto.resources.livemap.DispatchMarker.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setX(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setY(value);
      break;
    case 3:
      var value = new resources_timestamp_timestamp_pb.Timestamp;
      reader.readMessage(value,resources_timestamp_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setUpdatedAt(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setIcon(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setPopup(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setLink(value);
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
proto.resources.livemap.DispatchMarker.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.resources.livemap.DispatchMarker.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.resources.livemap.DispatchMarker} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.resources.livemap.DispatchMarker.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getX();
  if (f !== 0.0) {
    writer.writeFloat(
      1,
      f
    );
  }
  f = message.getY();
  if (f !== 0.0) {
    writer.writeFloat(
      2,
      f
    );
  }
  f = message.getUpdatedAt();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      resources_timestamp_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getIcon();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getPopup();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getLink();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
};


/**
 * optional float x = 1;
 * @return {number}
 */
proto.resources.livemap.DispatchMarker.prototype.getX = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 1, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.resources.livemap.DispatchMarker} returns this
 */
proto.resources.livemap.DispatchMarker.prototype.setX = function(value) {
  return jspb.Message.setProto3FloatField(this, 1, value);
};


/**
 * optional float y = 2;
 * @return {number}
 */
proto.resources.livemap.DispatchMarker.prototype.getY = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 2, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.resources.livemap.DispatchMarker} returns this
 */
proto.resources.livemap.DispatchMarker.prototype.setY = function(value) {
  return jspb.Message.setProto3FloatField(this, 2, value);
};


/**
 * optional resources.timestamp.Timestamp updated_at = 3;
 * @return {?proto.resources.timestamp.Timestamp}
 */
proto.resources.livemap.DispatchMarker.prototype.getUpdatedAt = function() {
  return /** @type{?proto.resources.timestamp.Timestamp} */ (
    jspb.Message.getWrapperField(this, resources_timestamp_timestamp_pb.Timestamp, 3));
};


/**
 * @param {?proto.resources.timestamp.Timestamp|undefined} value
 * @return {!proto.resources.livemap.DispatchMarker} returns this
*/
proto.resources.livemap.DispatchMarker.prototype.setUpdatedAt = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.resources.livemap.DispatchMarker} returns this
 */
proto.resources.livemap.DispatchMarker.prototype.clearUpdatedAt = function() {
  return this.setUpdatedAt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.resources.livemap.DispatchMarker.prototype.hasUpdatedAt = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional string name = 4;
 * @return {string}
 */
proto.resources.livemap.DispatchMarker.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.resources.livemap.DispatchMarker} returns this
 */
proto.resources.livemap.DispatchMarker.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string icon = 5;
 * @return {string}
 */
proto.resources.livemap.DispatchMarker.prototype.getIcon = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.resources.livemap.DispatchMarker} returns this
 */
proto.resources.livemap.DispatchMarker.prototype.setIcon = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string popup = 6;
 * @return {string}
 */
proto.resources.livemap.DispatchMarker.prototype.getPopup = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.resources.livemap.DispatchMarker} returns this
 */
proto.resources.livemap.DispatchMarker.prototype.setPopup = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string link = 7;
 * @return {string}
 */
proto.resources.livemap.DispatchMarker.prototype.getLink = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.resources.livemap.DispatchMarker} returns this
 */
proto.resources.livemap.DispatchMarker.prototype.setLink = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


goog.object.extend(exports, proto.resources.livemap);
