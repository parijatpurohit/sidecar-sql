// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var User_FindByRollAndName_pb = require('./User_FindByRollAndName_pb.js');

function serialize_storage_User_FindByRollAndNameRequest(arg) {
  if (!(arg instanceof User_FindByRollAndName_pb.User_FindByRollAndNameRequest)) {
    throw new Error('Expected argument of type storage.User_FindByRollAndNameRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_storage_User_FindByRollAndNameRequest(buffer_arg) {
  return User_FindByRollAndName_pb.User_FindByRollAndNameRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_storage_User_FindByRollAndNameResponse(arg) {
  if (!(arg instanceof User_FindByRollAndName_pb.User_FindByRollAndNameResponse)) {
    throw new Error('Expected argument of type storage.User_FindByRollAndNameResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_storage_User_FindByRollAndNameResponse(buffer_arg) {
  return User_FindByRollAndName_pb.User_FindByRollAndNameResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Server definition
var StorageServiceService = exports.StorageServiceService = {
  // Sends a greeting
user_FindByRollAndName: {
    path: '/storage.StorageService/User_FindByRollAndName',
    requestStream: false,
    responseStream: false,
    requestType: User_FindByRollAndName_pb.User_FindByRollAndNameRequest,
    responseType: User_FindByRollAndName_pb.User_FindByRollAndNameResponse,
    requestSerialize: serialize_storage_User_FindByRollAndNameRequest,
    requestDeserialize: deserialize_storage_User_FindByRollAndNameRequest,
    responseSerialize: serialize_storage_User_FindByRollAndNameResponse,
    responseDeserialize: deserialize_storage_User_FindByRollAndNameResponse,
  },
};

exports.StorageServiceClient = grpc.makeGenericClientConstructor(StorageServiceService);
