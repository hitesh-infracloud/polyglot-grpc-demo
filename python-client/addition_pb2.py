# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: addition.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0e\x61\x64\x64ition.proto\x12\x08\x61\x64\x64ition\"3\n\nAddRequest\x12\x0e\n\x06number\x18\x01 \x01(\x05\x12\x15\n\ranotherNumber\x18\x02 \x01(\x05\"\x1a\n\x0b\x41\x64\x64Response\x12\x0b\n\x03sum\x18\x01 \x01(\x03\x32@\n\x08\x41\x64\x64ition\x12\x34\n\x03\x41\x64\x64\x12\x14.addition.AddRequest\x1a\x15.addition.AddResponse\"\x00\x62\x06proto3')



_ADDREQUEST = DESCRIPTOR.message_types_by_name['AddRequest']
_ADDRESPONSE = DESCRIPTOR.message_types_by_name['AddResponse']
AddRequest = _reflection.GeneratedProtocolMessageType('AddRequest', (_message.Message,), {
  'DESCRIPTOR' : _ADDREQUEST,
  '__module__' : 'addition_pb2'
  # @@protoc_insertion_point(class_scope:addition.AddRequest)
  })
_sym_db.RegisterMessage(AddRequest)

AddResponse = _reflection.GeneratedProtocolMessageType('AddResponse', (_message.Message,), {
  'DESCRIPTOR' : _ADDRESPONSE,
  '__module__' : 'addition_pb2'
  # @@protoc_insertion_point(class_scope:addition.AddResponse)
  })
_sym_db.RegisterMessage(AddResponse)

_ADDITION = DESCRIPTOR.services_by_name['Addition']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _ADDREQUEST._serialized_start=28
  _ADDREQUEST._serialized_end=79
  _ADDRESPONSE._serialized_start=81
  _ADDRESPONSE._serialized_end=107
  _ADDITION._serialized_start=109
  _ADDITION._serialized_end=173
# @@protoc_insertion_point(module_scope)