//@Time : 2020/10/13 上午9:10
//@Author : bishisimo

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: protos/redux/redux.proto

package redux

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_redux_redux_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_redux_redux_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protos_redux_redux_proto_rawDescGZIP(), []int{0}
}

type Heat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host  string `protobuf:"bytes,1,opt,name=Host,proto3" json:"Host,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Port  int32  `protobuf:"varint,3,opt,name=Port,proto3" json:"Port,omitempty"`
	Event string `protobuf:"bytes,4,opt,name=Event,proto3" json:"Event,omitempty"`
}

func (x *Heat) Reset() {
	*x = Heat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_redux_redux_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heat) ProtoMessage() {}

func (x *Heat) ProtoReflect() protoreflect.Message {
	mi := &file_protos_redux_redux_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heat.ProtoReflect.Descriptor instead.
func (*Heat) Descriptor() ([]byte, []int) {
	return file_protos_redux_redux_proto_rawDescGZIP(), []int{1}
}

func (x *Heat) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Heat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Heat) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Heat) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_redux_redux_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_redux_redux_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_protos_redux_redux_proto_rawDescGZIP(), []int{2}
}

func (x *Reply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_redux_redux_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_protos_redux_redux_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_protos_redux_redux_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type DeviceKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Heat  *Heat `protobuf:"bytes,1,opt,name=Heat,proto3" json:"Heat,omitempty"`
	Index int32 `protobuf:"varint,2,opt,name=Index,proto3" json:"Index,omitempty"`
}

func (x *DeviceKey) Reset() {
	*x = DeviceKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_redux_redux_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceKey) ProtoMessage() {}

func (x *DeviceKey) ProtoReflect() protoreflect.Message {
	mi := &file_protos_redux_redux_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceKey.ProtoReflect.Descriptor instead.
func (*DeviceKey) Descriptor() ([]byte, []int) {
	return file_protos_redux_redux_proto_rawDescGZIP(), []int{4}
}

func (x *DeviceKey) GetHeat() *Heat {
	if x != nil {
		return x.Heat
	}
	return nil
}

func (x *DeviceKey) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key            *DeviceKey `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	StateCode      int32      `protobuf:"varint,2,opt,name=StateCode,proto3" json:"StateCode,omitempty"`
	LastActiveTime string     `protobuf:"bytes,3,opt,name=LastActiveTime,proto3" json:"LastActiveTime,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_redux_redux_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_protos_redux_redux_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_protos_redux_redux_proto_rawDescGZIP(), []int{5}
}

func (x *Device) GetKey() *DeviceKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Device) GetStateCode() int32 {
	if x != nil {
		return x.StateCode
	}
	return 0
}

func (x *Device) GetLastActiveTime() string {
	if x != nil {
		return x.LastActiveTime
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices []*Device `protobuf:"bytes,1,rep,name=Devices,proto3" json:"Devices,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_redux_redux_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_redux_redux_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_protos_redux_redux_proto_rawDescGZIP(), []int{6}
}

func (x *ListResponse) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

type TextMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromHost string `protobuf:"bytes,1,opt,name=FromHost,proto3" json:"FromHost,omitempty"`
	Text     string `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *TextMessage) Reset() {
	*x = TextMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_redux_redux_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextMessage) ProtoMessage() {}

func (x *TextMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_redux_redux_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextMessage.ProtoReflect.Descriptor instead.
func (*TextMessage) Descriptor() ([]byte, []int) {
	return file_protos_redux_redux_proto_rawDescGZIP(), []int{7}
}

func (x *TextMessage) GetFromHost() string {
	if x != nil {
		return x.FromHost
	}
	return ""
}

func (x *TextMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type FileMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`           //文件标识
	FromHost string `protobuf:"bytes,2,opt,name=FromHost,proto3" json:"FromHost,omitempty"` //发送方ip
	Num      int32  `protobuf:"varint,3,opt,name=Num,proto3" json:"Num,omitempty"`          //文件数量
	Path     string `protobuf:"bytes,4,opt,name=Path,proto3" json:"Path,omitempty"`         //当前区块
	Size     int32  `protobuf:"varint,5,opt,name=Size,proto3" json:"Size,omitempty"`        //总文件大小
	Data     []byte `protobuf:"bytes,6,opt,name=Data,proto3" json:"Data,omitempty"`         //数据
	Perm     int32  `protobuf:"varint,7,opt,name=Perm,proto3" json:"Perm,omitempty"`        //文件权限
}

func (x *FileMessage) Reset() {
	*x = FileMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_redux_redux_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMessage) ProtoMessage() {}

func (x *FileMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_redux_redux_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMessage.ProtoReflect.Descriptor instead.
func (*FileMessage) Descriptor() ([]byte, []int) {
	return file_protos_redux_redux_proto_rawDescGZIP(), []int{8}
}

func (x *FileMessage) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *FileMessage) GetFromHost() string {
	if x != nil {
		return x.FromHost
	}
	return ""
}

func (x *FileMessage) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *FileMessage) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileMessage) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FileMessage) GetPerm() int32 {
	if x != nil {
		return x.Perm
	}
	return 0
}

var File_protos_redux_redux_proto protoreflect.FileDescriptor

var file_protos_redux_redux_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x64, 0x75, 0x78, 0x2f, 0x72,
	0x65, 0x64, 0x75, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x58, 0x0a, 0x04, 0x48, 0x65, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x48,
	0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x1f, 0x0a,
	0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1d,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3c, 0x0a,
	0x09, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x04, 0x48, 0x65,
	0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x48, 0x65, 0x61, 0x74, 0x52,
	0x04, 0x48, 0x65, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x6c, 0x0a, 0x06, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x03,
	0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x4c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x4c, 0x61, 0x73, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x07, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x3d, 0x0a, 0x0b,
	0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46,
	0x72, 0x6f, 0x6d, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46,
	0x72, 0x6f, 0x6d, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x22, 0x9d, 0x01, 0x0a, 0x0b,
	0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x46, 0x72, 0x6f, 0x6d, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x46, 0x72, 0x6f, 0x6d, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x4e, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x65, 0x72, 0x6d, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x65, 0x72, 0x6d, 0x32, 0xc5, 0x02, 0x0a, 0x05,
	0x52, 0x65, 0x64, 0x75, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x42, 0x65, 0x61, 0x74, 0x48, 0x65, 0x61,
	0x74, 0x12, 0x05, 0x2e, 0x48, 0x65, 0x61, 0x74, 0x1a, 0x06, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x1d, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12,
	0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x06, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x23, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x06, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x0a, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x1a, 0x07, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65,
	0x78, 0x74, 0x12, 0x0c, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x06, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x0c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x06, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x28, 0x01, 0x12, 0x26, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x30, 0x01, 0x12, 0x16, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x06, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x06, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x19, 0x0a, 0x07,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x06, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x06, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65,
	0x64, 0x75, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_redux_redux_proto_rawDescOnce sync.Once
	file_protos_redux_redux_proto_rawDescData = file_protos_redux_redux_proto_rawDesc
)

func file_protos_redux_redux_proto_rawDescGZIP() []byte {
	file_protos_redux_redux_proto_rawDescOnce.Do(func() {
		file_protos_redux_redux_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_redux_redux_proto_rawDescData)
	})
	return file_protos_redux_redux_proto_rawDescData
}

var file_protos_redux_redux_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protos_redux_redux_proto_goTypes = []interface{}{
	(*Empty)(nil),        // 0: Empty
	(*Heat)(nil),         // 1: Heat
	(*Reply)(nil),        // 2: Reply
	(*Message)(nil),      // 3: Message
	(*DeviceKey)(nil),    // 4: DeviceKey
	(*Device)(nil),       // 5: Device
	(*ListResponse)(nil), // 6: ListResponse
	(*TextMessage)(nil),  // 7: TextMessage
	(*FileMessage)(nil),  // 8: FileMessage
}
var file_protos_redux_redux_proto_depIdxs = []int32{
	1,  // 0: DeviceKey.Heat:type_name -> Heat
	4,  // 1: Device.Key:type_name -> DeviceKey
	5,  // 2: ListResponse.Devices:type_name -> Device
	1,  // 3: Redux.BeatHeat:input_type -> Heat
	0,  // 4: Redux.CheckHealth:input_type -> Empty
	0,  // 5: Redux.ListDevice:input_type -> Empty
	4,  // 6: Redux.GetDevice:input_type -> DeviceKey
	7,  // 7: Redux.SendText:input_type -> TextMessage
	8,  // 8: Redux.SendFile:input_type -> FileMessage
	0,  // 9: Redux.GetMessageStream:input_type -> Empty
	0,  // 10: Redux.Stop:input_type -> Empty
	0,  // 11: Redux.Restart:input_type -> Empty
	0,  // 12: Redux.Snapshot:input_type -> Empty
	2,  // 13: Redux.BeatHeat:output_type -> Reply
	2,  // 14: Redux.CheckHealth:output_type -> Reply
	6,  // 15: Redux.ListDevice:output_type -> ListResponse
	5,  // 16: Redux.GetDevice:output_type -> Device
	2,  // 17: Redux.SendText:output_type -> Reply
	2,  // 18: Redux.SendFile:output_type -> Reply
	3,  // 19: Redux.GetMessageStream:output_type -> Message
	2,  // 20: Redux.Stop:output_type -> Reply
	2,  // 21: Redux.Restart:output_type -> Reply
	2,  // 22: Redux.Snapshot:output_type -> Reply
	13, // [13:23] is the sub-list for method output_type
	3,  // [3:13] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_protos_redux_redux_proto_init() }
func file_protos_redux_redux_proto_init() {
	if File_protos_redux_redux_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_redux_redux_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_protos_redux_redux_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heat); i {
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
		file_protos_redux_redux_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_protos_redux_redux_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_protos_redux_redux_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceKey); i {
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
		file_protos_redux_redux_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_protos_redux_redux_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_protos_redux_redux_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextMessage); i {
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
		file_protos_redux_redux_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMessage); i {
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
			RawDescriptor: file_protos_redux_redux_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_redux_redux_proto_goTypes,
		DependencyIndexes: file_protos_redux_redux_proto_depIdxs,
		MessageInfos:      file_protos_redux_redux_proto_msgTypes,
	}.Build()
	File_protos_redux_redux_proto = out.File
	file_protos_redux_redux_proto_rawDesc = nil
	file_protos_redux_redux_proto_goTypes = nil
	file_protos_redux_redux_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReduxClient is the client API for Redux service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReduxClient interface {
	//心跳
	BeatHeat(ctx context.Context, in *Heat, opts ...grpc.CallOption) (*Reply, error)
	//健康检查
	CheckHealth(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Reply, error)
	//展示在线列表;
	ListDevice(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListResponse, error)
	//传入非host获取host
	GetDevice(ctx context.Context, in *DeviceKey, opts ...grpc.CallOption) (*Device, error)
	//发送文本信息
	SendText(ctx context.Context, in *TextMessage, opts ...grpc.CallOption) (*Reply, error)
	//添加流式订阅,返回流式数据
	SendFile(ctx context.Context, opts ...grpc.CallOption) (Redux_SendFileClient, error)
	GetMessageStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Redux_GetMessageStreamClient, error)
	Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Reply, error)
	Restart(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Reply, error)
	Snapshot(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Reply, error)
}

type reduxClient struct {
	cc grpc.ClientConnInterface
}

func NewReduxClient(cc grpc.ClientConnInterface) ReduxClient {
	return &reduxClient{cc}
}

func (c *reduxClient) BeatHeat(ctx context.Context, in *Heat, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/Redux/BeatHeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reduxClient) CheckHealth(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/Redux/CheckHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reduxClient) ListDevice(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/Redux/ListDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reduxClient) GetDevice(ctx context.Context, in *DeviceKey, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/Redux/GetDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reduxClient) SendText(ctx context.Context, in *TextMessage, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/Redux/SendText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reduxClient) SendFile(ctx context.Context, opts ...grpc.CallOption) (Redux_SendFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Redux_serviceDesc.Streams[0], "/Redux/SendFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &reduxSendFileClient{stream}
	return x, nil
}

type Redux_SendFileClient interface {
	Send(*FileMessage) error
	CloseAndRecv() (*Reply, error)
	grpc.ClientStream
}

type reduxSendFileClient struct {
	grpc.ClientStream
}

func (x *reduxSendFileClient) Send(m *FileMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *reduxSendFileClient) CloseAndRecv() (*Reply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *reduxClient) GetMessageStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Redux_GetMessageStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Redux_serviceDesc.Streams[1], "/Redux/GetMessageStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &reduxGetMessageStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Redux_GetMessageStreamClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type reduxGetMessageStreamClient struct {
	grpc.ClientStream
}

func (x *reduxGetMessageStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *reduxClient) Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/Redux/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reduxClient) Restart(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/Redux/Restart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reduxClient) Snapshot(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/Redux/Snapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReduxServer is the server API for Redux service.
type ReduxServer interface {
	//心跳
	BeatHeat(context.Context, *Heat) (*Reply, error)
	//健康检查
	CheckHealth(context.Context, *Empty) (*Reply, error)
	//展示在线列表;
	ListDevice(context.Context, *Empty) (*ListResponse, error)
	//传入非host获取host
	GetDevice(context.Context, *DeviceKey) (*Device, error)
	//发送文本信息
	SendText(context.Context, *TextMessage) (*Reply, error)
	//添加流式订阅,返回流式数据
	SendFile(Redux_SendFileServer) error
	GetMessageStream(*Empty, Redux_GetMessageStreamServer) error
	Stop(context.Context, *Empty) (*Reply, error)
	Restart(context.Context, *Empty) (*Reply, error)
	Snapshot(context.Context, *Empty) (*Reply, error)
}

// UnimplementedReduxServer can be embedded to have forward compatible implementations.
type UnimplementedReduxServer struct {
}

func (*UnimplementedReduxServer) BeatHeat(context.Context, *Heat) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeatHeat not implemented")
}
func (*UnimplementedReduxServer) CheckHealth(context.Context, *Empty) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckHealth not implemented")
}
func (*UnimplementedReduxServer) ListDevice(context.Context, *Empty) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDevice not implemented")
}
func (*UnimplementedReduxServer) GetDevice(context.Context, *DeviceKey) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevice not implemented")
}
func (*UnimplementedReduxServer) SendText(context.Context, *TextMessage) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendText not implemented")
}
func (*UnimplementedReduxServer) SendFile(Redux_SendFileServer) error {
	return status.Errorf(codes.Unimplemented, "method SendFile not implemented")
}
func (*UnimplementedReduxServer) GetMessageStream(*Empty, Redux_GetMessageStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMessageStream not implemented")
}
func (*UnimplementedReduxServer) Stop(context.Context, *Empty) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (*UnimplementedReduxServer) Restart(context.Context, *Empty) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (*UnimplementedReduxServer) Snapshot(context.Context, *Empty) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Snapshot not implemented")
}

func RegisterReduxServer(s *grpc.Server, srv ReduxServer) {
	s.RegisterService(&_Redux_serviceDesc, srv)
}

func _Redux_BeatHeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Heat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReduxServer).BeatHeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Redux/BeatHeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReduxServer).BeatHeat(ctx, req.(*Heat))
	}
	return interceptor(ctx, in, info, handler)
}

func _Redux_CheckHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReduxServer).CheckHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Redux/CheckHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReduxServer).CheckHealth(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Redux_ListDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReduxServer).ListDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Redux/ListDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReduxServer).ListDevice(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Redux_GetDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReduxServer).GetDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Redux/GetDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReduxServer).GetDevice(ctx, req.(*DeviceKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Redux_SendText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReduxServer).SendText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Redux/SendText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReduxServer).SendText(ctx, req.(*TextMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Redux_SendFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ReduxServer).SendFile(&reduxSendFileServer{stream})
}

type Redux_SendFileServer interface {
	SendAndClose(*Reply) error
	Recv() (*FileMessage, error)
	grpc.ServerStream
}

type reduxSendFileServer struct {
	grpc.ServerStream
}

func (x *reduxSendFileServer) SendAndClose(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *reduxSendFileServer) Recv() (*FileMessage, error) {
	m := new(FileMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Redux_GetMessageStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReduxServer).GetMessageStream(m, &reduxGetMessageStreamServer{stream})
}

type Redux_GetMessageStreamServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type reduxGetMessageStreamServer struct {
	grpc.ServerStream
}

func (x *reduxGetMessageStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Redux_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReduxServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Redux/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReduxServer).Stop(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Redux_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReduxServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Redux/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReduxServer).Restart(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Redux_Snapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReduxServer).Snapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Redux/Snapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReduxServer).Snapshot(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Redux_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Redux",
	HandlerType: (*ReduxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BeatHeat",
			Handler:    _Redux_BeatHeat_Handler,
		},
		{
			MethodName: "CheckHealth",
			Handler:    _Redux_CheckHealth_Handler,
		},
		{
			MethodName: "ListDevice",
			Handler:    _Redux_ListDevice_Handler,
		},
		{
			MethodName: "GetDevice",
			Handler:    _Redux_GetDevice_Handler,
		},
		{
			MethodName: "SendText",
			Handler:    _Redux_SendText_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Redux_Stop_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _Redux_Restart_Handler,
		},
		{
			MethodName: "Snapshot",
			Handler:    _Redux_Snapshot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendFile",
			Handler:       _Redux_SendFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetMessageStream",
			Handler:       _Redux_GetMessageStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/redux/redux.proto",
}
