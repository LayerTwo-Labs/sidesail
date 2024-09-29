// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: drivechain/v1/drivechain.proto

package drivechainv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListSidechainsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSidechainsRequest) Reset() {
	*x = ListSidechainsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drivechain_v1_drivechain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSidechainsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSidechainsRequest) ProtoMessage() {}

func (x *ListSidechainsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drivechain_v1_drivechain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSidechainsRequest.ProtoReflect.Descriptor instead.
func (*ListSidechainsRequest) Descriptor() ([]byte, []int) {
	return file_drivechain_v1_drivechain_proto_rawDescGZIP(), []int{0}
}

type ListSidechainsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sidechains []*ListSidechainsResponse_Sidechain `protobuf:"bytes,1,rep,name=sidechains,proto3" json:"sidechains,omitempty"`
}

func (x *ListSidechainsResponse) Reset() {
	*x = ListSidechainsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drivechain_v1_drivechain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSidechainsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSidechainsResponse) ProtoMessage() {}

func (x *ListSidechainsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_drivechain_v1_drivechain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSidechainsResponse.ProtoReflect.Descriptor instead.
func (*ListSidechainsResponse) Descriptor() ([]byte, []int) {
	return file_drivechain_v1_drivechain_proto_rawDescGZIP(), []int{1}
}

func (x *ListSidechainsResponse) GetSidechains() []*ListSidechainsResponse_Sidechain {
	if x != nil {
		return x.Sidechains
	}
	return nil
}

type ListSidechainProposalsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSidechainProposalsRequest) Reset() {
	*x = ListSidechainProposalsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drivechain_v1_drivechain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSidechainProposalsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSidechainProposalsRequest) ProtoMessage() {}

func (x *ListSidechainProposalsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drivechain_v1_drivechain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSidechainProposalsRequest.ProtoReflect.Descriptor instead.
func (*ListSidechainProposalsRequest) Descriptor() ([]byte, []int) {
	return file_drivechain_v1_drivechain_proto_rawDescGZIP(), []int{2}
}

type SidechainProposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot           uint32 `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	Data           []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	DataHash       string `protobuf:"bytes,3,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
	VoteCount      uint32 `protobuf:"varint,4,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`
	ProposalHeight uint32 `protobuf:"varint,5,opt,name=proposal_height,json=proposalHeight,proto3" json:"proposal_height,omitempty"`
	ProposalAge    uint32 `protobuf:"varint,6,opt,name=proposal_age,json=proposalAge,proto3" json:"proposal_age,omitempty"`
}

func (x *SidechainProposal) Reset() {
	*x = SidechainProposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drivechain_v1_drivechain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SidechainProposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SidechainProposal) ProtoMessage() {}

func (x *SidechainProposal) ProtoReflect() protoreflect.Message {
	mi := &file_drivechain_v1_drivechain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SidechainProposal.ProtoReflect.Descriptor instead.
func (*SidechainProposal) Descriptor() ([]byte, []int) {
	return file_drivechain_v1_drivechain_proto_rawDescGZIP(), []int{3}
}

func (x *SidechainProposal) GetSlot() uint32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *SidechainProposal) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SidechainProposal) GetDataHash() string {
	if x != nil {
		return x.DataHash
	}
	return ""
}

func (x *SidechainProposal) GetVoteCount() uint32 {
	if x != nil {
		return x.VoteCount
	}
	return 0
}

func (x *SidechainProposal) GetProposalHeight() uint32 {
	if x != nil {
		return x.ProposalHeight
	}
	return 0
}

func (x *SidechainProposal) GetProposalAge() uint32 {
	if x != nil {
		return x.ProposalAge
	}
	return 0
}

type ListSidechainProposalsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proposals []*SidechainProposal `protobuf:"bytes,1,rep,name=proposals,proto3" json:"proposals,omitempty"`
}

func (x *ListSidechainProposalsResponse) Reset() {
	*x = ListSidechainProposalsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drivechain_v1_drivechain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSidechainProposalsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSidechainProposalsResponse) ProtoMessage() {}

func (x *ListSidechainProposalsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_drivechain_v1_drivechain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSidechainProposalsResponse.ProtoReflect.Descriptor instead.
func (*ListSidechainProposalsResponse) Descriptor() ([]byte, []int) {
	return file_drivechain_v1_drivechain_proto_rawDescGZIP(), []int{4}
}

func (x *ListSidechainProposalsResponse) GetProposals() []*SidechainProposal {
	if x != nil {
		return x.Proposals
	}
	return nil
}

type GetSidechainCtipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SidechainNumber uint32 `protobuf:"varint,1,opt,name=sidechain_number,json=sidechainNumber,proto3" json:"sidechain_number,omitempty"`
}

func (x *GetSidechainCtipRequest) Reset() {
	*x = GetSidechainCtipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drivechain_v1_drivechain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSidechainCtipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSidechainCtipRequest) ProtoMessage() {}

func (x *GetSidechainCtipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drivechain_v1_drivechain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSidechainCtipRequest.ProtoReflect.Descriptor instead.
func (*GetSidechainCtipRequest) Descriptor() ([]byte, []int) {
	return file_drivechain_v1_drivechain_proto_rawDescGZIP(), []int{5}
}

func (x *GetSidechainCtipRequest) GetSidechainNumber() uint32 {
	if x != nil {
		return x.SidechainNumber
	}
	return 0
}

type GetSidechainCtipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Outpoint       *string `protobuf:"bytes,1,opt,name=outpoint,proto3,oneof" json:"outpoint,omitempty"`
	Value          *uint64 `protobuf:"varint,2,opt,name=value,proto3,oneof" json:"value,omitempty"`
	SequenceNumber *uint32 `protobuf:"varint,3,opt,name=sequence_number,json=sequenceNumber,proto3,oneof" json:"sequence_number,omitempty"`
}

func (x *GetSidechainCtipResponse) Reset() {
	*x = GetSidechainCtipResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drivechain_v1_drivechain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSidechainCtipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSidechainCtipResponse) ProtoMessage() {}

func (x *GetSidechainCtipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_drivechain_v1_drivechain_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSidechainCtipResponse.ProtoReflect.Descriptor instead.
func (*GetSidechainCtipResponse) Descriptor() ([]byte, []int) {
	return file_drivechain_v1_drivechain_proto_rawDescGZIP(), []int{6}
}

func (x *GetSidechainCtipResponse) GetOutpoint() string {
	if x != nil && x.Outpoint != nil {
		return *x.Outpoint
	}
	return ""
}

func (x *GetSidechainCtipResponse) GetValue() uint64 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

func (x *GetSidechainCtipResponse) GetSequenceNumber() uint32 {
	if x != nil && x.SequenceNumber != nil {
		return *x.SequenceNumber
	}
	return 0
}

type ListSidechainsResponse_Sidechain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title         string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Nversion      uint32 `protobuf:"varint,3,opt,name=nversion,proto3" json:"nversion,omitempty"`
	Hashid1       string `protobuf:"bytes,4,opt,name=hashid1,proto3" json:"hashid1,omitempty"`
	Hashid2       string `protobuf:"bytes,5,opt,name=hashid2,proto3" json:"hashid2,omitempty"`
	Slot          int32  `protobuf:"varint,6,opt,name=slot,proto3" json:"slot,omitempty"`
	AmountSatoshi int64  `protobuf:"varint,7,opt,name=amount_satoshi,json=amountSatoshi,proto3" json:"amount_satoshi,omitempty"`
	ChaintipTxid  string `protobuf:"bytes,8,opt,name=chaintip_txid,json=chaintipTxid,proto3" json:"chaintip_txid,omitempty"`
}

func (x *ListSidechainsResponse_Sidechain) Reset() {
	*x = ListSidechainsResponse_Sidechain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drivechain_v1_drivechain_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSidechainsResponse_Sidechain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSidechainsResponse_Sidechain) ProtoMessage() {}

func (x *ListSidechainsResponse_Sidechain) ProtoReflect() protoreflect.Message {
	mi := &file_drivechain_v1_drivechain_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSidechainsResponse_Sidechain.ProtoReflect.Descriptor instead.
func (*ListSidechainsResponse_Sidechain) Descriptor() ([]byte, []int) {
	return file_drivechain_v1_drivechain_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListSidechainsResponse_Sidechain) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListSidechainsResponse_Sidechain) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ListSidechainsResponse_Sidechain) GetNversion() uint32 {
	if x != nil {
		return x.Nversion
	}
	return 0
}

func (x *ListSidechainsResponse_Sidechain) GetHashid1() string {
	if x != nil {
		return x.Hashid1
	}
	return ""
}

func (x *ListSidechainsResponse_Sidechain) GetHashid2() string {
	if x != nil {
		return x.Hashid2
	}
	return ""
}

func (x *ListSidechainsResponse_Sidechain) GetSlot() int32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *ListSidechainsResponse_Sidechain) GetAmountSatoshi() int64 {
	if x != nil {
		return x.AmountSatoshi
	}
	return 0
}

func (x *ListSidechainsResponse_Sidechain) GetChaintipTxid() string {
	if x != nil {
		return x.ChaintipTxid
	}
	return ""
}

var File_drivechain_v1_drivechain_proto protoreflect.FileDescriptor

var file_drivechain_v1_drivechain_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x22,
	0x17, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xdf, 0x02, 0x0a, 0x16, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x73, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x64, 0x65,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53,
	0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x0a, 0x73, 0x69, 0x64, 0x65, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x73, 0x1a, 0xf3, 0x01, 0x0a, 0x09, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x73, 0x68, 0x69, 0x64,
	0x31, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x73, 0x68, 0x69, 0x64, 0x31,
	0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x73, 0x68, 0x69, 0x64, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x68, 0x61, 0x73, 0x68, 0x69, 0x64, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c,
	0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x61,
	0x74, 0x6f, 0x73, 0x68, 0x69, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x74, 0x69,
	0x70, 0x5f, 0x74, 0x78, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x74, 0x69, 0x70, 0x54, 0x78, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x1d, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xc3, 0x01, 0x0a, 0x11,
	0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x76, 0x6f, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x41, 0x67,
	0x65, 0x22, 0x60, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x73, 0x22, 0x44, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x43, 0x74, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29,
	0x0a, 0x10, 0x73, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x73, 0x69, 0x64, 0x65, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xaf, 0x01, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x43, 0x74, 0x69, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x6f, 0x75, 0x74, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6f, 0x75, 0x74, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x0e, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0xe9, 0x01, 0x0a, 0x11,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5d, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x73, 0x12, 0x24, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69,
	0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x75, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x12, 0x2c, 0x2e, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x64,
	0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xcd, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x54, 0x77, 0x6f, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x73,
	0x61, 0x69, 0x6c, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_drivechain_v1_drivechain_proto_rawDescOnce sync.Once
	file_drivechain_v1_drivechain_proto_rawDescData = file_drivechain_v1_drivechain_proto_rawDesc
)

func file_drivechain_v1_drivechain_proto_rawDescGZIP() []byte {
	file_drivechain_v1_drivechain_proto_rawDescOnce.Do(func() {
		file_drivechain_v1_drivechain_proto_rawDescData = protoimpl.X.CompressGZIP(file_drivechain_v1_drivechain_proto_rawDescData)
	})
	return file_drivechain_v1_drivechain_proto_rawDescData
}

var file_drivechain_v1_drivechain_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_drivechain_v1_drivechain_proto_goTypes = []any{
	(*ListSidechainsRequest)(nil),            // 0: drivechain.v1.ListSidechainsRequest
	(*ListSidechainsResponse)(nil),           // 1: drivechain.v1.ListSidechainsResponse
	(*ListSidechainProposalsRequest)(nil),    // 2: drivechain.v1.ListSidechainProposalsRequest
	(*SidechainProposal)(nil),                // 3: drivechain.v1.SidechainProposal
	(*ListSidechainProposalsResponse)(nil),   // 4: drivechain.v1.ListSidechainProposalsResponse
	(*GetSidechainCtipRequest)(nil),          // 5: drivechain.v1.GetSidechainCtipRequest
	(*GetSidechainCtipResponse)(nil),         // 6: drivechain.v1.GetSidechainCtipResponse
	(*ListSidechainsResponse_Sidechain)(nil), // 7: drivechain.v1.ListSidechainsResponse.Sidechain
}
var file_drivechain_v1_drivechain_proto_depIdxs = []int32{
	7, // 0: drivechain.v1.ListSidechainsResponse.sidechains:type_name -> drivechain.v1.ListSidechainsResponse.Sidechain
	3, // 1: drivechain.v1.ListSidechainProposalsResponse.proposals:type_name -> drivechain.v1.SidechainProposal
	0, // 2: drivechain.v1.DrivechainService.ListSidechains:input_type -> drivechain.v1.ListSidechainsRequest
	2, // 3: drivechain.v1.DrivechainService.ListSidechainProposals:input_type -> drivechain.v1.ListSidechainProposalsRequest
	1, // 4: drivechain.v1.DrivechainService.ListSidechains:output_type -> drivechain.v1.ListSidechainsResponse
	4, // 5: drivechain.v1.DrivechainService.ListSidechainProposals:output_type -> drivechain.v1.ListSidechainProposalsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_drivechain_v1_drivechain_proto_init() }
func file_drivechain_v1_drivechain_proto_init() {
	if File_drivechain_v1_drivechain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_drivechain_v1_drivechain_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListSidechainsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drivechain_v1_drivechain_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListSidechainsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drivechain_v1_drivechain_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListSidechainProposalsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drivechain_v1_drivechain_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SidechainProposal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drivechain_v1_drivechain_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListSidechainProposalsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drivechain_v1_drivechain_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetSidechainCtipRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drivechain_v1_drivechain_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetSidechainCtipResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drivechain_v1_drivechain_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ListSidechainsResponse_Sidechain); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_drivechain_v1_drivechain_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_drivechain_v1_drivechain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_drivechain_v1_drivechain_proto_goTypes,
		DependencyIndexes: file_drivechain_v1_drivechain_proto_depIdxs,
		MessageInfos:      file_drivechain_v1_drivechain_proto_msgTypes,
	}.Build()
	File_drivechain_v1_drivechain_proto = out.File
	file_drivechain_v1_drivechain_proto_rawDesc = nil
	file_drivechain_v1_drivechain_proto_goTypes = nil
	file_drivechain_v1_drivechain_proto_depIdxs = nil
}
