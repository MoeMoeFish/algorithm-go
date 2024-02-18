// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: paxoskv.proto

package paxoskv

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

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vi64 int64 `protobuf:"varint,1,opt,name=Vi64,proto3" json:"Vi64,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paxoskv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_paxoskv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_paxoskv_proto_rawDescGZIP(), []int{0}
}

func (x *Value) GetVi64() int64 {
	if x != nil {
		return x.Vi64
	}
	return 0
}

type BalloutNum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N          int64 `protobuf:"varint,1,opt,name=N,proto3" json:"N,omitempty"`
	ProposerId int64 `protobuf:"varint,2,opt,name=ProposerId,proto3" json:"ProposerId,omitempty"`
}

func (x *BalloutNum) Reset() {
	*x = BalloutNum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paxoskv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalloutNum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalloutNum) ProtoMessage() {}

func (x *BalloutNum) ProtoReflect() protoreflect.Message {
	mi := &file_paxoskv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalloutNum.ProtoReflect.Descriptor instead.
func (*BalloutNum) Descriptor() ([]byte, []int) {
	return file_paxoskv_proto_rawDescGZIP(), []int{1}
}

func (x *BalloutNum) GetN() int64 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *BalloutNum) GetProposerId() int64 {
	if x != nil {
		return x.ProposerId
	}
	return 0
}

type PaxosInstanceId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Ver int64  `protobuf:"varint,2,opt,name=Ver,proto3" json:"Ver,omitempty"`
}

func (x *PaxosInstanceId) Reset() {
	*x = PaxosInstanceId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paxoskv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaxosInstanceId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaxosInstanceId) ProtoMessage() {}

func (x *PaxosInstanceId) ProtoReflect() protoreflect.Message {
	mi := &file_paxoskv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaxosInstanceId.ProtoReflect.Descriptor instead.
func (*PaxosInstanceId) Descriptor() ([]byte, []int) {
	return file_paxoskv_proto_rawDescGZIP(), []int{2}
}

func (x *PaxosInstanceId) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PaxosInstanceId) GetVer() int64 {
	if x != nil {
		return x.Ver
	}
	return 0
}

type Proposer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  *PaxosInstanceId `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Bal *BalloutNum      `protobuf:"bytes,2,opt,name=Bal,proto3" json:"Bal,omitempty"`
	Val *Value           `protobuf:"bytes,3,opt,name=Val,proto3" json:"Val,omitempty"`
}

func (x *Proposer) Reset() {
	*x = Proposer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paxoskv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proposer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proposer) ProtoMessage() {}

func (x *Proposer) ProtoReflect() protoreflect.Message {
	mi := &file_paxoskv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proposer.ProtoReflect.Descriptor instead.
func (*Proposer) Descriptor() ([]byte, []int) {
	return file_paxoskv_proto_rawDescGZIP(), []int{3}
}

func (x *Proposer) GetId() *PaxosInstanceId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Proposer) GetBal() *BalloutNum {
	if x != nil {
		return x.Bal
	}
	return nil
}

func (x *Proposer) GetVal() *Value {
	if x != nil {
		return x.Val
	}
	return nil
}

type Acceptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastBal *BalloutNum `protobuf:"bytes,1,opt,name=LastBal,proto3" json:"LastBal,omitempty"`
	VBal    *BalloutNum `protobuf:"bytes,2,opt,name=VBal,proto3" json:"VBal,omitempty"`
	Val     *Value      `protobuf:"bytes,3,opt,name=Val,proto3" json:"Val,omitempty"`
}

func (x *Acceptor) Reset() {
	*x = Acceptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paxoskv_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Acceptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Acceptor) ProtoMessage() {}

func (x *Acceptor) ProtoReflect() protoreflect.Message {
	mi := &file_paxoskv_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Acceptor.ProtoReflect.Descriptor instead.
func (*Acceptor) Descriptor() ([]byte, []int) {
	return file_paxoskv_proto_rawDescGZIP(), []int{4}
}

func (x *Acceptor) GetLastBal() *BalloutNum {
	if x != nil {
		return x.LastBal
	}
	return nil
}

func (x *Acceptor) GetVBal() *BalloutNum {
	if x != nil {
		return x.VBal
	}
	return nil
}

func (x *Acceptor) GetVal() *Value {
	if x != nil {
		return x.Val
	}
	return nil
}

var File_paxoskv_proto protoreflect.FileDescriptor

var file_paxoskv_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x78, 0x6f, 0x73, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x61, 0x78, 0x6f, 0x73, 0x6b, 0x76, 0x22, 0x1b, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x56, 0x69, 0x36, 0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x56, 0x69, 0x36, 0x34, 0x22, 0x3a, 0x0a, 0x0a, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x75, 0x74,
	0x4e, 0x75, 0x6d, 0x12, 0x0c, 0x0a, 0x01, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01,
	0x4e, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x35, 0x0a, 0x0f, 0x50, 0x61, 0x78, 0x6f, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x56, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x56, 0x65, 0x72, 0x22, 0x7d, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x70, 0x61, 0x78, 0x6f, 0x73, 0x6b, 0x76, 0x2e, 0x50, 0x61, 0x78, 0x6f, 0x73,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x52, 0x02, 0x49, 0x64, 0x12, 0x25,
	0x0a, 0x03, 0x42, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61,
	0x78, 0x6f, 0x73, 0x6b, 0x76, 0x2e, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x4e, 0x75, 0x6d,
	0x52, 0x03, 0x42, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x03, 0x56, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x61, 0x78, 0x6f, 0x73, 0x6b, 0x76, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x03, 0x56, 0x61, 0x6c, 0x22, 0x84, 0x01, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x07, 0x4c, 0x61, 0x73, 0x74, 0x42, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61, 0x78, 0x6f, 0x73, 0x6b, 0x76, 0x2e,
	0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x4e, 0x75, 0x6d, 0x52, 0x07, 0x4c, 0x61, 0x73, 0x74,
	0x42, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x04, 0x56, 0x42, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61, 0x78, 0x6f, 0x73, 0x6b, 0x76, 0x2e, 0x42, 0x61, 0x6c, 0x6c,
	0x6f, 0x75, 0x74, 0x4e, 0x75, 0x6d, 0x52, 0x04, 0x56, 0x42, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x03,
	0x56, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x61, 0x78, 0x6f,
	0x73, 0x6b, 0x76, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x56, 0x61, 0x6c, 0x32, 0x6e,
	0x0a, 0x07, 0x50, 0x61, 0x78, 0x6f, 0x73, 0x4b, 0x56, 0x12, 0x31, 0x0a, 0x07, 0x50, 0x72, 0x65,
	0x70, 0x61, 0x72, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x61, 0x78, 0x6f, 0x73, 0x6b, 0x76, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x70, 0x61, 0x78, 0x6f, 0x73, 0x6b,
	0x76, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x61, 0x78, 0x6f, 0x73, 0x6b, 0x76,
	0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x70, 0x61, 0x78, 0x6f,
	0x73, 0x6b, 0x76, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x22, 0x00, 0x42, 0x20,
	0x5a, 0x1e, 0x6d, 0x6f, 0x65, 0x6d, 0x6f, 0x65, 0x66, 0x69, 0x73, 0x68, 0x2f, 0x61, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x6d, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x61, 0x78, 0x6f, 0x73, 0x6b, 0x76,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_paxoskv_proto_rawDescOnce sync.Once
	file_paxoskv_proto_rawDescData = file_paxoskv_proto_rawDesc
)

func file_paxoskv_proto_rawDescGZIP() []byte {
	file_paxoskv_proto_rawDescOnce.Do(func() {
		file_paxoskv_proto_rawDescData = protoimpl.X.CompressGZIP(file_paxoskv_proto_rawDescData)
	})
	return file_paxoskv_proto_rawDescData
}

var file_paxoskv_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_paxoskv_proto_goTypes = []interface{}{
	(*Value)(nil),           // 0: paxoskv.Value
	(*BalloutNum)(nil),      // 1: paxoskv.BalloutNum
	(*PaxosInstanceId)(nil), // 2: paxoskv.PaxosInstanceId
	(*Proposer)(nil),        // 3: paxoskv.Proposer
	(*Acceptor)(nil),        // 4: paxoskv.Acceptor
}
var file_paxoskv_proto_depIdxs = []int32{
	2, // 0: paxoskv.Proposer.Id:type_name -> paxoskv.PaxosInstanceId
	1, // 1: paxoskv.Proposer.Bal:type_name -> paxoskv.BalloutNum
	0, // 2: paxoskv.Proposer.Val:type_name -> paxoskv.Value
	1, // 3: paxoskv.Acceptor.LastBal:type_name -> paxoskv.BalloutNum
	1, // 4: paxoskv.Acceptor.VBal:type_name -> paxoskv.BalloutNum
	0, // 5: paxoskv.Acceptor.Val:type_name -> paxoskv.Value
	3, // 6: paxoskv.PaxosKV.Prepare:input_type -> paxoskv.Proposer
	3, // 7: paxoskv.PaxosKV.Accept:input_type -> paxoskv.Proposer
	4, // 8: paxoskv.PaxosKV.Prepare:output_type -> paxoskv.Acceptor
	4, // 9: paxoskv.PaxosKV.Accept:output_type -> paxoskv.Acceptor
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_paxoskv_proto_init() }
func file_paxoskv_proto_init() {
	if File_paxoskv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_paxoskv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_paxoskv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalloutNum); i {
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
		file_paxoskv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaxosInstanceId); i {
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
		file_paxoskv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proposer); i {
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
		file_paxoskv_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Acceptor); i {
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
			RawDescriptor: file_paxoskv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_paxoskv_proto_goTypes,
		DependencyIndexes: file_paxoskv_proto_depIdxs,
		MessageInfos:      file_paxoskv_proto_msgTypes,
	}.Build()
	File_paxoskv_proto = out.File
	file_paxoskv_proto_rawDesc = nil
	file_paxoskv_proto_goTypes = nil
	file_paxoskv_proto_depIdxs = nil
}
