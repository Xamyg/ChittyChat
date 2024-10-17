// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.20.3
// source: chittychat/chitty_chat.proto

package chitty_chat

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

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text      string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	User      string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	mi := &file_chittychat_chitty_chat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_chittychat_chitty_chat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_chittychat_chitty_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Chat) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Chat) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Chat) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

var File_chittychat_chitty_chat_proto protoreflect.FileDescriptor

var file_chittychat_chitty_chat_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x69,
	0x74, 0x74, 0x79, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x22, 0x4c, 0x0a, 0x04, 0x43, 0x68,
	0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0x44, 0x0a, 0x0a, 0x43, 0x68, 0x69, 0x74,
	0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x10, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x20,
	0x5a, 0x1e, 0x48, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x33, 0x2f, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79,
	0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x5f, 0x63, 0x68, 0x61, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chittychat_chitty_chat_proto_rawDescOnce sync.Once
	file_chittychat_chitty_chat_proto_rawDescData = file_chittychat_chitty_chat_proto_rawDesc
)

func file_chittychat_chitty_chat_proto_rawDescGZIP() []byte {
	file_chittychat_chitty_chat_proto_rawDescOnce.Do(func() {
		file_chittychat_chitty_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chittychat_chitty_chat_proto_rawDescData)
	})
	return file_chittychat_chitty_chat_proto_rawDescData
}

var file_chittychat_chitty_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chittychat_chitty_chat_proto_goTypes = []any{
	(*Chat)(nil), // 0: chittychat.Chat
}
var file_chittychat_chitty_chat_proto_depIdxs = []int32{
	0, // 0: chittychat.ChittyChat.ChatStream:input_type -> chittychat.Chat
	0, // 1: chittychat.ChittyChat.ChatStream:output_type -> chittychat.Chat
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chittychat_chitty_chat_proto_init() }
func file_chittychat_chitty_chat_proto_init() {
	if File_chittychat_chitty_chat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chittychat_chitty_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chittychat_chitty_chat_proto_goTypes,
		DependencyIndexes: file_chittychat_chitty_chat_proto_depIdxs,
		MessageInfos:      file_chittychat_chitty_chat_proto_msgTypes,
	}.Build()
	File_chittychat_chitty_chat_proto = out.File
	file_chittychat_chitty_chat_proto_rawDesc = nil
	file_chittychat_chitty_chat_proto_goTypes = nil
	file_chittychat_chitty_chat_proto_depIdxs = nil
}
