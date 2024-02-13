// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: account/account_list_as_claim.proto

package account

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

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_list_as_claim_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_list_as_claim_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_account_account_list_as_claim_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AccountListAsClaimResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *AccountListAsClaimResult) Reset() {
	*x = AccountListAsClaimResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_list_as_claim_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountListAsClaimResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountListAsClaimResult) ProtoMessage() {}

func (x *AccountListAsClaimResult) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_list_as_claim_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountListAsClaimResult.ProtoReflect.Descriptor instead.
func (*AccountListAsClaimResult) Descriptor() ([]byte, []int) {
	return file_account_account_list_as_claim_proto_rawDescGZIP(), []int{1}
}

func (x *AccountListAsClaimResult) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

type AccountListAsClaimRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *AccountListAsClaimRequest) Reset() {
	*x = AccountListAsClaimRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_list_as_claim_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountListAsClaimRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountListAsClaimRequest) ProtoMessage() {}

func (x *AccountListAsClaimRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_list_as_claim_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountListAsClaimRequest.ProtoReflect.Descriptor instead.
func (*AccountListAsClaimRequest) Descriptor() ([]byte, []int) {
	return file_account_account_list_as_claim_proto_rawDescGZIP(), []int{2}
}

func (x *AccountListAsClaimRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_account_account_list_as_claim_proto protoreflect.FileDescriptor

var file_account_account_list_as_claim_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x61, 0x73, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2d,
	0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a,
	0x18, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x34, 0x0a, 0x19, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0x6a, 0x0a,
	0x12, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x43, 0x6c, 0x61,
	0x69, 0x6d, 0x12, 0x22, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x75, 0x72, 0x69, 0x73, 0x74, 0x69, 0x6b,
	0x72, 0x6f, 0x74, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_account_list_as_claim_proto_rawDescOnce sync.Once
	file_account_account_list_as_claim_proto_rawDescData = file_account_account_list_as_claim_proto_rawDesc
)

func file_account_account_list_as_claim_proto_rawDescGZIP() []byte {
	file_account_account_list_as_claim_proto_rawDescOnce.Do(func() {
		file_account_account_list_as_claim_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_account_list_as_claim_proto_rawDescData)
	})
	return file_account_account_list_as_claim_proto_rawDescData
}

var file_account_account_list_as_claim_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_account_account_list_as_claim_proto_goTypes = []interface{}{
	(*Account)(nil),                   // 0: account.Account
	(*AccountListAsClaimResult)(nil),  // 1: account.AccountListAsClaimResult
	(*AccountListAsClaimRequest)(nil), // 2: account.AccountListAsClaimRequest
}
var file_account_account_list_as_claim_proto_depIdxs = []int32{
	0, // 0: account.AccountListAsClaimResult.accounts:type_name -> account.Account
	2, // 1: account.AccountListService.ListAsClaim:input_type -> account.AccountListAsClaimRequest
	1, // 2: account.AccountListService.ListAsClaim:output_type -> account.AccountListAsClaimResult
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_account_account_list_as_claim_proto_init() }
func file_account_account_list_as_claim_proto_init() {
	if File_account_account_list_as_claim_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_account_list_as_claim_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_account_account_list_as_claim_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountListAsClaimResult); i {
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
		file_account_account_list_as_claim_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountListAsClaimRequest); i {
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
			RawDescriptor: file_account_account_list_as_claim_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_account_list_as_claim_proto_goTypes,
		DependencyIndexes: file_account_account_list_as_claim_proto_depIdxs,
		MessageInfos:      file_account_account_list_as_claim_proto_msgTypes,
	}.Build()
	File_account_account_list_as_claim_proto = out.File
	file_account_account_list_as_claim_proto_rawDesc = nil
	file_account_account_list_as_claim_proto_goTypes = nil
	file_account_account_list_as_claim_proto_depIdxs = nil
}
