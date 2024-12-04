# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: set.proto
# Protobuf Python Version: 5.28.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    28,
    1,
    '',
    'set.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from config import protocol_pb2 as config_dot_protocol__pb2
from config import technology_pb2 as config_dot_technology__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\tset.proto\x12\x02pb\x1a\x15\x63onfig/protocol.proto\x1a\x17\x63onfig/technology.proto\"<\n\x15SetAutoconnectRequest\x12\x0f\n\x07\x65nabled\x18\x01 \x01(\x08\x12\x12\n\nserver_tag\x18\x02 \x01(\t\"$\n\x11SetGenericRequest\x12\x0f\n\x07\x65nabled\x18\x01 \x01(\x08\"!\n\x10SetUint32Request\x12\r\n\x05value\x18\x01 \x01(\r\"@\n\x1eSetThreatProtectionLiteRequest\x12\x1e\n\x16threat_protection_lite\x18\x01 \x01(\x08\"\xa5\x01\n\x1fSetThreatProtectionLiteResponse\x12&\n\nerror_code\x18\x01 \x01(\x0e\x32\x10.pb.SetErrorCodeH\x00\x12N\n!set_threat_protection_lite_status\x18\x02 \x01(\x0e\x32!.pb.SetThreatProtectionLiteStatusH\x00\x42\n\n\x08response\"<\n\rSetDNSRequest\x12\x0b\n\x03\x64ns\x18\x02 \x03(\t\x12\x1e\n\x16threat_protection_lite\x18\x03 \x01(\x08\"p\n\x0eSetDNSResponse\x12&\n\nerror_code\x18\x02 \x01(\x0e\x32\x10.pb.SetErrorCodeH\x00\x12*\n\x0eset_dns_status\x18\x03 \x01(\x0e\x32\x10.pb.SetDNSStatusH\x00\x42\n\n\x08response\"+\n\x14SetKillSwitchRequest\x12\x13\n\x0bkill_switch\x18\x02 \x01(\x08\"/\n\x10SetNotifyRequest\x12\x0b\n\x03uid\x18\x02 \x01(\x03\x12\x0e\n\x06notify\x18\x03 \x01(\x08\"+\n\x0eSetTrayRequest\x12\x0b\n\x03uid\x18\x02 \x01(\x03\x12\x0c\n\x04tray\x18\x03 \x01(\x08\"8\n\x12SetProtocolRequest\x12\"\n\x08protocol\x18\x02 \x01(\x0e\x32\x10.config.Protocol\"\x7f\n\x13SetProtocolResponse\x12&\n\nerror_code\x18\x01 \x01(\x0e\x32\x10.pb.SetErrorCodeH\x00\x12\x34\n\x13set_protocol_status\x18\x02 \x01(\x0e\x32\x15.pb.SetProtocolStatusH\x00\x42\n\n\x08response\">\n\x14SetTechnologyRequest\x12&\n\ntechnology\x18\x02 \x01(\x0e\x32\x12.config.Technology\"1\n\tPortRange\x12\x12\n\nstart_port\x18\x01 \x01(\x03\x12\x10\n\x08\x65nd_port\x18\x02 \x01(\x03\"+\n\x19SetAllowlistSubnetRequest\x12\x0e\n\x06subnet\x18\x01 \x01(\t\"]\n\x18SetAllowlistPortsRequest\x12\x0e\n\x06is_udp\x18\x01 \x01(\x08\x12\x0e\n\x06is_tcp\x18\x02 \x01(\x08\x12!\n\nport_range\x18\x03 \x01(\x0b\x32\r.pb.PortRange\"\xac\x01\n\x13SetAllowlistRequest\x12\x45\n\x1cset_allowlist_subnet_request\x18\x01 \x01(\x0b\x32\x1d.pb.SetAllowlistSubnetRequestH\x00\x12\x43\n\x1bset_allowlist_ports_request\x18\x02 \x01(\x0b\x32\x1c.pb.SetAllowlistPortsRequestH\x00\x42\t\n\x07request\")\n\x16SetLANDiscoveryRequest\x12\x0f\n\x07\x65nabled\x18\x01 \x01(\x08\"\x8c\x01\n\x17SetLANDiscoveryResponse\x12&\n\nerror_code\x18\x01 \x01(\x0e\x32\x10.pb.SetErrorCodeH\x00\x12=\n\x18set_lan_discovery_status\x18\x02 \x01(\x0e\x32\x19.pb.SetLANDiscoveryStatusH\x00\x42\n\n\x08response*>\n\x0cSetErrorCode\x12\x0b\n\x07\x46\x41ILURE\x10\x00\x12\x10\n\x0c\x43ONFIG_ERROR\x10\x01\x12\x0f\n\x0b\x41LREADY_SET\x10\x02*Q\n\x1dSetThreatProtectionLiteStatus\x12\x12\n\x0eTPL_CONFIGURED\x10\x00\x12\x1c\n\x18TPL_CONFIGURED_DNS_RESET\x10\x01*n\n\x0cSetDNSStatus\x12\x12\n\x0e\x44NS_CONFIGURED\x10\x00\x12\x1c\n\x18\x44NS_CONFIGURED_TPL_RESET\x10\x01\x12\x17\n\x13INVALID_DNS_ADDRESS\x10\x02\x12\x13\n\x0fTOO_MANY_VALUES\x10\x03*d\n\x11SetProtocolStatus\x12\x17\n\x13PROTOCOL_CONFIGURED\x10\x00\x12\x1e\n\x1aPROTOCOL_CONFIGURED_VPN_ON\x10\x01\x12\x16\n\x12INVALID_TECHNOLOGY\x10\x02*[\n\x15SetLANDiscoveryStatus\x12\x18\n\x14\x44ISCOVERY_CONFIGURED\x10\x00\x12(\n$DISCOVERY_CONFIGURED_ALLOWLIST_RESET\x10\x01\x42\x31Z/github.com/NordSecurity/nordvpn-linux/daemon/pbb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'set_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z/github.com/NordSecurity/nordvpn-linux/daemon/pb'
  _globals['_SETERRORCODE']._serialized_start=1552
  _globals['_SETERRORCODE']._serialized_end=1614
  _globals['_SETTHREATPROTECTIONLITESTATUS']._serialized_start=1616
  _globals['_SETTHREATPROTECTIONLITESTATUS']._serialized_end=1697
  _globals['_SETDNSSTATUS']._serialized_start=1699
  _globals['_SETDNSSTATUS']._serialized_end=1809
  _globals['_SETPROTOCOLSTATUS']._serialized_start=1811
  _globals['_SETPROTOCOLSTATUS']._serialized_end=1911
  _globals['_SETLANDISCOVERYSTATUS']._serialized_start=1913
  _globals['_SETLANDISCOVERYSTATUS']._serialized_end=2004
  _globals['_SETAUTOCONNECTREQUEST']._serialized_start=65
  _globals['_SETAUTOCONNECTREQUEST']._serialized_end=125
  _globals['_SETGENERICREQUEST']._serialized_start=127
  _globals['_SETGENERICREQUEST']._serialized_end=163
  _globals['_SETUINT32REQUEST']._serialized_start=165
  _globals['_SETUINT32REQUEST']._serialized_end=198
  _globals['_SETTHREATPROTECTIONLITEREQUEST']._serialized_start=200
  _globals['_SETTHREATPROTECTIONLITEREQUEST']._serialized_end=264
  _globals['_SETTHREATPROTECTIONLITERESPONSE']._serialized_start=267
  _globals['_SETTHREATPROTECTIONLITERESPONSE']._serialized_end=432
  _globals['_SETDNSREQUEST']._serialized_start=434
  _globals['_SETDNSREQUEST']._serialized_end=494
  _globals['_SETDNSRESPONSE']._serialized_start=496
  _globals['_SETDNSRESPONSE']._serialized_end=608
  _globals['_SETKILLSWITCHREQUEST']._serialized_start=610
  _globals['_SETKILLSWITCHREQUEST']._serialized_end=653
  _globals['_SETNOTIFYREQUEST']._serialized_start=655
  _globals['_SETNOTIFYREQUEST']._serialized_end=702
  _globals['_SETTRAYREQUEST']._serialized_start=704
  _globals['_SETTRAYREQUEST']._serialized_end=747
  _globals['_SETPROTOCOLREQUEST']._serialized_start=749
  _globals['_SETPROTOCOLREQUEST']._serialized_end=805
  _globals['_SETPROTOCOLRESPONSE']._serialized_start=807
  _globals['_SETPROTOCOLRESPONSE']._serialized_end=934
  _globals['_SETTECHNOLOGYREQUEST']._serialized_start=936
  _globals['_SETTECHNOLOGYREQUEST']._serialized_end=998
  _globals['_PORTRANGE']._serialized_start=1000
  _globals['_PORTRANGE']._serialized_end=1049
  _globals['_SETALLOWLISTSUBNETREQUEST']._serialized_start=1051
  _globals['_SETALLOWLISTSUBNETREQUEST']._serialized_end=1094
  _globals['_SETALLOWLISTPORTSREQUEST']._serialized_start=1096
  _globals['_SETALLOWLISTPORTSREQUEST']._serialized_end=1189
  _globals['_SETALLOWLISTREQUEST']._serialized_start=1192
  _globals['_SETALLOWLISTREQUEST']._serialized_end=1364
  _globals['_SETLANDISCOVERYREQUEST']._serialized_start=1366
  _globals['_SETLANDISCOVERYREQUEST']._serialized_end=1407
  _globals['_SETLANDISCOVERYRESPONSE']._serialized_start=1410
  _globals['_SETLANDISCOVERYRESPONSE']._serialized_end=1550
# @@protoc_insertion_point(module_scope)
