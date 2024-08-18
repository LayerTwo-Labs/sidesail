// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: drivechain/v1/drivechain.proto

package drivechainv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfirmedSatoshi uint64 `protobuf:"varint,1,opt,name=confirmed_satoshi,json=confirmedSatoshi,proto3" json:"confirmed_satoshi,omitempty"`
	PendingSatoshi   uint64 `protobuf:"varint,2,opt,name=pending_satoshi,json=pendingSatoshi,proto3" json:"pending_satoshi,omitempty"`
}

func (x *GetBalanceResponse) Reset() {
	*x = GetBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drivechain_v1_drivechain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceResponse) ProtoMessage() {}

func (x *GetBalanceResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return file_drivechain_v1_drivechain_proto_rawDescGZIP(), []int{0}
}

func (x *GetBalanceResponse) GetConfirmedSatoshi() uint64 {
	if x != nil {
		return x.ConfirmedSatoshi
	}
	return 0
}

func (x *GetBalanceResponse) GetPendingSatoshi() uint64 {
	if x != nil {
		return x.PendingSatoshi
	}
	return 0
}

type ListTransactionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *ListTransactionsResponse) Reset() {
	*x = ListTransactionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drivechain_v1_drivechain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTransactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransactionsResponse) ProtoMessage() {}

func (x *ListTransactionsResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListTransactionsResponse.ProtoReflect.Descriptor instead.
func (*ListTransactionsResponse) Descriptor() ([]byte, []int) {
	return file_drivechain_v1_drivechain_proto_rawDescGZIP(), []int{1}
}

func (x *ListTransactionsResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type Confirmation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height    uint32                 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Confirmation) Reset() {
	*x = Confirmation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drivechain_v1_drivechain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmation) ProtoMessage() {}

func (x *Confirmation) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Confirmation.ProtoReflect.Descriptor instead.
func (*Confirmation) Descriptor() ([]byte, []int) {
	return file_drivechain_v1_drivechain_proto_rawDescGZIP(), []int{2}
}

func (x *Confirmation) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Confirmation) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txid             string        `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	FeeSatoshi       uint64        `protobuf:"varint,2,opt,name=fee_satoshi,json=feeSatoshi,proto3" json:"fee_satoshi,omitempty"`
	ReceivedSatoshi  uint64        `protobuf:"varint,3,opt,name=received_satoshi,json=receivedSatoshi,proto3" json:"received_satoshi,omitempty"`
	SentSatoshi      uint64        `protobuf:"varint,4,opt,name=sent_satoshi,json=sentSatoshi,proto3" json:"sent_satoshi,omitempty"`
	ConfirmationTime *Confirmation `protobuf:"bytes,5,opt,name=confirmation_time,json=confirmationTime,proto3" json:"confirmation_time,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drivechain_v1_drivechain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_drivechain_v1_drivechain_proto_rawDescGZIP(), []int{3}
}

func (x *Transaction) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *Transaction) GetFeeSatoshi() uint64 {
	if x != nil {
		return x.FeeSatoshi
	}
	return 0
}

func (x *Transaction) GetReceivedSatoshi() uint64 {
	if x != nil {
		return x.ReceivedSatoshi
	}
	return 0
}

func (x *Transaction) GetSentSatoshi() uint64 {
	if x != nil {
		return x.SentSatoshi
	}
	return 0
}

func (x *Transaction) GetConfirmationTime() *Confirmation {
	if x != nil {
		return x.ConfirmationTime
	}
	return nil
}

type ListUnconfirmedTransactionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnconfirmedTransactions []*UnconfirmedTransaction `protobuf:"bytes,1,rep,name=unconfirmed_transactions,json=unconfirmedTransactions,proto3" json:"unconfirmed_transactions,omitempty"`
}

func (x *ListUnconfirmedTransactionsResponse) Reset() {
	*x = ListUnconfirmedTransactionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drivechain_v1_drivechain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUnconfirmedTransactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUnconfirmedTransactionsResponse) ProtoMessage() {}

func (x *ListUnconfirmedTransactionsResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListUnconfirmedTransactionsResponse.ProtoReflect.Descriptor instead.
func (*ListUnconfirmedTransactionsResponse) Descriptor() ([]byte, []int) {
	return file_drivechain_v1_drivechain_proto_rawDescGZIP(), []int{4}
}

func (x *ListUnconfirmedTransactionsResponse) GetUnconfirmedTransactions() []*UnconfirmedTransaction {
	if x != nil {
		return x.UnconfirmedTransactions
	}
	return nil
}

type UnconfirmedTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VirtualSize uint32                 `protobuf:"varint,1,opt,name=virtual_size,json=virtualSize,proto3" json:"virtual_size,omitempty"`
	Weight      uint32                 `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	Time        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Txid        string                 `protobuf:"bytes,4,opt,name=txid,proto3" json:"txid,omitempty"`
	FeeSatoshi  uint64                 `protobuf:"varint,5,opt,name=fee_satoshi,json=feeSatoshi,proto3" json:"fee_satoshi,omitempty"`
}

func (x *UnconfirmedTransaction) Reset() {
	*x = UnconfirmedTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drivechain_v1_drivechain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnconfirmedTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnconfirmedTransaction) ProtoMessage() {}

func (x *UnconfirmedTransaction) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UnconfirmedTransaction.ProtoReflect.Descriptor instead.
func (*UnconfirmedTransaction) Descriptor() ([]byte, []int) {
	return file_drivechain_v1_drivechain_proto_rawDescGZIP(), []int{5}
}

func (x *UnconfirmedTransaction) GetVirtualSize() uint32 {
	if x != nil {
		return x.VirtualSize
	}
	return 0
}

func (x *UnconfirmedTransaction) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *UnconfirmedTransaction) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *UnconfirmedTransaction) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *UnconfirmedTransaction) GetFeeSatoshi() uint64 {
	if x != nil {
		return x.FeeSatoshi
	}
	return 0
}

var File_drivechain_v1_drivechain_proto protoreflect.FileDescriptor

var file_drivechain_v1_drivechain_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64,
	0x5f, 0x73, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x53, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69,
	0x12, 0x27, 0x0a, 0x0f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x61, 0x74, 0x6f,
	0x73, 0x68, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x70, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x53, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x22, 0x5a, 0x0a, 0x18, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x60, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x38, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xda, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66,
	0x65, 0x65, 0x5f, 0x73, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x66, 0x65, 0x65, 0x53, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x12, 0x29, 0x0a, 0x10,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x73, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x53, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73,
	0x65, 0x6e, 0x74, 0x53, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x12, 0x48, 0x0a, 0x11, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x23, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x6e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x18,
	0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x17, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb8,
	0x01, 0x0a, 0x16, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x65, 0x65, 0x5f,
	0x73, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x66,
	0x65, 0x65, 0x53, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x32, 0x9c, 0x02, 0x0a, 0x11, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x47, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x27, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a,
	0x1b, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x32, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xcd, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0f,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x54, 0x77, 0x6f, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x73, 0x69, 0x64, 0x65,
	0x73, 0x61, 0x69, 0x6c, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_drivechain_v1_drivechain_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_drivechain_v1_drivechain_proto_goTypes = []any{
	(*GetBalanceResponse)(nil),                  // 0: drivechain.v1.GetBalanceResponse
	(*ListTransactionsResponse)(nil),            // 1: drivechain.v1.ListTransactionsResponse
	(*Confirmation)(nil),                        // 2: drivechain.v1.Confirmation
	(*Transaction)(nil),                         // 3: drivechain.v1.Transaction
	(*ListUnconfirmedTransactionsResponse)(nil), // 4: drivechain.v1.ListUnconfirmedTransactionsResponse
	(*UnconfirmedTransaction)(nil),              // 5: drivechain.v1.UnconfirmedTransaction
	(*timestamppb.Timestamp)(nil),               // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                       // 7: google.protobuf.Empty
}
var file_drivechain_v1_drivechain_proto_depIdxs = []int32{
	3, // 0: drivechain.v1.ListTransactionsResponse.transactions:type_name -> drivechain.v1.Transaction
	6, // 1: drivechain.v1.Confirmation.timestamp:type_name -> google.protobuf.Timestamp
	2, // 2: drivechain.v1.Transaction.confirmation_time:type_name -> drivechain.v1.Confirmation
	5, // 3: drivechain.v1.ListUnconfirmedTransactionsResponse.unconfirmed_transactions:type_name -> drivechain.v1.UnconfirmedTransaction
	6, // 4: drivechain.v1.UnconfirmedTransaction.time:type_name -> google.protobuf.Timestamp
	7, // 5: drivechain.v1.DrivechainService.GetBalance:input_type -> google.protobuf.Empty
	7, // 6: drivechain.v1.DrivechainService.ListTransactions:input_type -> google.protobuf.Empty
	7, // 7: drivechain.v1.DrivechainService.ListUnconfirmedTransactions:input_type -> google.protobuf.Empty
	0, // 8: drivechain.v1.DrivechainService.GetBalance:output_type -> drivechain.v1.GetBalanceResponse
	1, // 9: drivechain.v1.DrivechainService.ListTransactions:output_type -> drivechain.v1.ListTransactionsResponse
	4, // 10: drivechain.v1.DrivechainService.ListUnconfirmedTransactions:output_type -> drivechain.v1.ListUnconfirmedTransactionsResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_drivechain_v1_drivechain_proto_init() }
func file_drivechain_v1_drivechain_proto_init() {
	if File_drivechain_v1_drivechain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_drivechain_v1_drivechain_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetBalanceResponse); i {
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
			switch v := v.(*ListTransactionsResponse); i {
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
			switch v := v.(*Confirmation); i {
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
			switch v := v.(*Transaction); i {
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
			switch v := v.(*ListUnconfirmedTransactionsResponse); i {
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
			switch v := v.(*UnconfirmedTransaction); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_drivechain_v1_drivechain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
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
