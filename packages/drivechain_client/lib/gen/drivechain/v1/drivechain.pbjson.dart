//
//  Generated code. Do not modify.
//  source: drivechain/v1/drivechain.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use listSidechainsRequestDescriptor instead')
const ListSidechainsRequest$json = {
  '1': 'ListSidechainsRequest',
};

/// Descriptor for `ListSidechainsRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List listSidechainsRequestDescriptor = $convert.base64Decode(
    'ChVMaXN0U2lkZWNoYWluc1JlcXVlc3Q=');

@$core.Deprecated('Use listSidechainsResponseDescriptor instead')
const ListSidechainsResponse$json = {
  '1': 'ListSidechainsResponse',
  '2': [
    {'1': 'sidechains', '3': 1, '4': 3, '5': 11, '6': '.drivechain.v1.ListSidechainsResponse.Sidechain', '10': 'sidechains'},
  ],
  '3': [ListSidechainsResponse_Sidechain$json],
};

@$core.Deprecated('Use listSidechainsResponseDescriptor instead')
const ListSidechainsResponse_Sidechain$json = {
  '1': 'Sidechain',
  '2': [
    {'1': 'title', '3': 1, '4': 1, '5': 9, '10': 'title'},
    {'1': 'description', '3': 2, '4': 1, '5': 9, '10': 'description'},
    {'1': 'nversion', '3': 3, '4': 1, '5': 13, '10': 'nversion'},
    {'1': 'hashid1', '3': 4, '4': 1, '5': 9, '10': 'hashid1'},
    {'1': 'hashid2', '3': 5, '4': 1, '5': 9, '10': 'hashid2'},
    {'1': 'slot', '3': 6, '4': 1, '5': 5, '10': 'slot'},
    {'1': 'amount_satoshi', '3': 7, '4': 1, '5': 3, '10': 'amountSatoshi'},
    {'1': 'chaintip_txid', '3': 8, '4': 1, '5': 9, '10': 'chaintipTxid'},
  ],
};

/// Descriptor for `ListSidechainsResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List listSidechainsResponseDescriptor = $convert.base64Decode(
    'ChZMaXN0U2lkZWNoYWluc1Jlc3BvbnNlEk8KCnNpZGVjaGFpbnMYASADKAsyLy5kcml2ZWNoYW'
    'luLnYxLkxpc3RTaWRlY2hhaW5zUmVzcG9uc2UuU2lkZWNoYWluUgpzaWRlY2hhaW5zGvMBCglT'
    'aWRlY2hhaW4SFAoFdGl0bGUYASABKAlSBXRpdGxlEiAKC2Rlc2NyaXB0aW9uGAIgASgJUgtkZX'
    'NjcmlwdGlvbhIaCghudmVyc2lvbhgDIAEoDVIIbnZlcnNpb24SGAoHaGFzaGlkMRgEIAEoCVIH'
    'aGFzaGlkMRIYCgdoYXNoaWQyGAUgASgJUgdoYXNoaWQyEhIKBHNsb3QYBiABKAVSBHNsb3QSJQ'
    'oOYW1vdW50X3NhdG9zaGkYByABKANSDWFtb3VudFNhdG9zaGkSIwoNY2hhaW50aXBfdHhpZBgI'
    'IAEoCVIMY2hhaW50aXBUeGlk');

