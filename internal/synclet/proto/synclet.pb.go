// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/synclet/synclet.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LogLevel int32

const (
	LogLevel_INFO    LogLevel = 0
	LogLevel_VERBOSE LogLevel = 1
	LogLevel_DEBUG   LogLevel = 2
)

var LogLevel_name = map[int32]string{
	0: "INFO",
	1: "VERBOSE",
	2: "DEBUG",
}

var LogLevel_value = map[string]int32{
	"INFO":    0,
	"VERBOSE": 1,
	"DEBUG":   2,
}

func (x LogLevel) String() string {
	return proto.EnumName(LogLevel_name, int32(x))
}

func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6f1080a1433c7af1, []int{0}
}

type GetContainerIdForPodRequest struct {
	PodId                string    `protobuf:"bytes,1,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty"`
	ImageId              string    `protobuf:"bytes,2,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	LogStyle             *LogStyle `protobuf:"bytes,3,opt,name=log_style,json=logStyle,proto3" json:"log_style,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetContainerIdForPodRequest) Reset()         { *m = GetContainerIdForPodRequest{} }
func (m *GetContainerIdForPodRequest) String() string { return proto.CompactTextString(m) }
func (*GetContainerIdForPodRequest) ProtoMessage()    {}
func (*GetContainerIdForPodRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f1080a1433c7af1, []int{0}
}

func (m *GetContainerIdForPodRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContainerIdForPodRequest.Unmarshal(m, b)
}
func (m *GetContainerIdForPodRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContainerIdForPodRequest.Marshal(b, m, deterministic)
}
func (m *GetContainerIdForPodRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContainerIdForPodRequest.Merge(m, src)
}
func (m *GetContainerIdForPodRequest) XXX_Size() int {
	return xxx_messageInfo_GetContainerIdForPodRequest.Size(m)
}
func (m *GetContainerIdForPodRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContainerIdForPodRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetContainerIdForPodRequest proto.InternalMessageInfo

func (m *GetContainerIdForPodRequest) GetPodId() string {
	if m != nil {
		return m.PodId
	}
	return ""
}

func (m *GetContainerIdForPodRequest) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *GetContainerIdForPodRequest) GetLogStyle() *LogStyle {
	if m != nil {
		return m.LogStyle
	}
	return nil
}

type GetContainerIdForPodReply struct {
	// Types that are valid to be assigned to Content:
	//	*GetContainerIdForPodReply_ContainerId
	//	*GetContainerIdForPodReply_Message
	Content              isGetContainerIdForPodReply_Content `protobuf_oneof:"content"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *GetContainerIdForPodReply) Reset()         { *m = GetContainerIdForPodReply{} }
func (m *GetContainerIdForPodReply) String() string { return proto.CompactTextString(m) }
func (*GetContainerIdForPodReply) ProtoMessage()    {}
func (*GetContainerIdForPodReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f1080a1433c7af1, []int{1}
}

func (m *GetContainerIdForPodReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContainerIdForPodReply.Unmarshal(m, b)
}
func (m *GetContainerIdForPodReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContainerIdForPodReply.Marshal(b, m, deterministic)
}
func (m *GetContainerIdForPodReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContainerIdForPodReply.Merge(m, src)
}
func (m *GetContainerIdForPodReply) XXX_Size() int {
	return xxx_messageInfo_GetContainerIdForPodReply.Size(m)
}
func (m *GetContainerIdForPodReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContainerIdForPodReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetContainerIdForPodReply proto.InternalMessageInfo

type isGetContainerIdForPodReply_Content interface {
	isGetContainerIdForPodReply_Content()
}

type GetContainerIdForPodReply_ContainerId struct {
	ContainerId string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3,oneof"`
}

type GetContainerIdForPodReply_Message struct {
	Message *LogMessage `protobuf:"bytes,2,opt,name=message,proto3,oneof"`
}

func (*GetContainerIdForPodReply_ContainerId) isGetContainerIdForPodReply_Content() {}

func (*GetContainerIdForPodReply_Message) isGetContainerIdForPodReply_Content() {}

func (m *GetContainerIdForPodReply) GetContent() isGetContainerIdForPodReply_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *GetContainerIdForPodReply) GetContainerId() string {
	if x, ok := m.GetContent().(*GetContainerIdForPodReply_ContainerId); ok {
		return x.ContainerId
	}
	return ""
}

func (m *GetContainerIdForPodReply) GetMessage() *LogMessage {
	if x, ok := m.GetContent().(*GetContainerIdForPodReply_Message); ok {
		return x.Message
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GetContainerIdForPodReply) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GetContainerIdForPodReply_OneofMarshaler, _GetContainerIdForPodReply_OneofUnmarshaler, _GetContainerIdForPodReply_OneofSizer, []interface{}{
		(*GetContainerIdForPodReply_ContainerId)(nil),
		(*GetContainerIdForPodReply_Message)(nil),
	}
}

func _GetContainerIdForPodReply_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GetContainerIdForPodReply)
	// content
	switch x := m.Content.(type) {
	case *GetContainerIdForPodReply_ContainerId:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ContainerId)
	case *GetContainerIdForPodReply_Message:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Message); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GetContainerIdForPodReply.Content has unexpected type %T", x)
	}
	return nil
}

func _GetContainerIdForPodReply_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GetContainerIdForPodReply)
	switch tag {
	case 1: // content.container_id
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Content = &GetContainerIdForPodReply_ContainerId{x}
		return true, err
	case 2: // content.message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LogMessage)
		err := b.DecodeMessage(msg)
		m.Content = &GetContainerIdForPodReply_Message{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GetContainerIdForPodReply_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GetContainerIdForPodReply)
	// content
	switch x := m.Content.(type) {
	case *GetContainerIdForPodReply_ContainerId:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ContainerId)))
		n += len(x.ContainerId)
	case *GetContainerIdForPodReply_Message:
		s := proto.Size(x.Message)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Cmd struct {
	Argv                 []string `protobuf:"bytes,1,rep,name=argv,proto3" json:"argv,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cmd) Reset()         { *m = Cmd{} }
func (m *Cmd) String() string { return proto.CompactTextString(m) }
func (*Cmd) ProtoMessage()    {}
func (*Cmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f1080a1433c7af1, []int{2}
}

func (m *Cmd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cmd.Unmarshal(m, b)
}
func (m *Cmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cmd.Marshal(b, m, deterministic)
}
func (m *Cmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cmd.Merge(m, src)
}
func (m *Cmd) XXX_Size() int {
	return xxx_messageInfo_Cmd.Size(m)
}
func (m *Cmd) XXX_DiscardUnknown() {
	xxx_messageInfo_Cmd.DiscardUnknown(m)
}

var xxx_messageInfo_Cmd proto.InternalMessageInfo

func (m *Cmd) GetArgv() []string {
	if m != nil {
		return m.Argv
	}
	return nil
}

type UpdateContainerRequest struct {
	ContainerId          string    `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	TarArchive           []byte    `protobuf:"bytes,2,opt,name=tar_archive,json=tarArchive,proto3" json:"tar_archive,omitempty"`
	FilesToDelete        []string  `protobuf:"bytes,3,rep,name=files_to_delete,json=filesToDelete,proto3" json:"files_to_delete,omitempty"`
	Commands             []*Cmd    `protobuf:"bytes,4,rep,name=commands,proto3" json:"commands,omitempty"`
	LogStyle             *LogStyle `protobuf:"bytes,5,opt,name=log_style,json=logStyle,proto3" json:"log_style,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateContainerRequest) Reset()         { *m = UpdateContainerRequest{} }
func (m *UpdateContainerRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateContainerRequest) ProtoMessage()    {}
func (*UpdateContainerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f1080a1433c7af1, []int{3}
}

func (m *UpdateContainerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateContainerRequest.Unmarshal(m, b)
}
func (m *UpdateContainerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateContainerRequest.Marshal(b, m, deterministic)
}
func (m *UpdateContainerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateContainerRequest.Merge(m, src)
}
func (m *UpdateContainerRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateContainerRequest.Size(m)
}
func (m *UpdateContainerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateContainerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateContainerRequest proto.InternalMessageInfo

func (m *UpdateContainerRequest) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *UpdateContainerRequest) GetTarArchive() []byte {
	if m != nil {
		return m.TarArchive
	}
	return nil
}

func (m *UpdateContainerRequest) GetFilesToDelete() []string {
	if m != nil {
		return m.FilesToDelete
	}
	return nil
}

func (m *UpdateContainerRequest) GetCommands() []*Cmd {
	if m != nil {
		return m.Commands
	}
	return nil
}

func (m *UpdateContainerRequest) GetLogStyle() *LogStyle {
	if m != nil {
		return m.LogStyle
	}
	return nil
}

type LogMessage struct {
	Level                LogLevel `protobuf:"varint,1,opt,name=level,proto3,enum=synclet.LogLevel" json:"level,omitempty"`
	Message              []byte   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogMessage) Reset()         { *m = LogMessage{} }
func (m *LogMessage) String() string { return proto.CompactTextString(m) }
func (*LogMessage) ProtoMessage()    {}
func (*LogMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f1080a1433c7af1, []int{4}
}

func (m *LogMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogMessage.Unmarshal(m, b)
}
func (m *LogMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogMessage.Marshal(b, m, deterministic)
}
func (m *LogMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogMessage.Merge(m, src)
}
func (m *LogMessage) XXX_Size() int {
	return xxx_messageInfo_LogMessage.Size(m)
}
func (m *LogMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LogMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LogMessage proto.InternalMessageInfo

func (m *LogMessage) GetLevel() LogLevel {
	if m != nil {
		return m.Level
	}
	return LogLevel_INFO
}

func (m *LogMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type UpdateContainerReply struct {
	LogMessage           *LogMessage `protobuf:"bytes,1,opt,name=log_message,json=logMessage,proto3" json:"log_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateContainerReply) Reset()         { *m = UpdateContainerReply{} }
func (m *UpdateContainerReply) String() string { return proto.CompactTextString(m) }
func (*UpdateContainerReply) ProtoMessage()    {}
func (*UpdateContainerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f1080a1433c7af1, []int{5}
}

func (m *UpdateContainerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateContainerReply.Unmarshal(m, b)
}
func (m *UpdateContainerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateContainerReply.Marshal(b, m, deterministic)
}
func (m *UpdateContainerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateContainerReply.Merge(m, src)
}
func (m *UpdateContainerReply) XXX_Size() int {
	return xxx_messageInfo_UpdateContainerReply.Size(m)
}
func (m *UpdateContainerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateContainerReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateContainerReply proto.InternalMessageInfo

func (m *UpdateContainerReply) GetLogMessage() *LogMessage {
	if m != nil {
		return m.LogMessage
	}
	return nil
}

type LogStyle struct {
	ColorsEnabled        bool     `protobuf:"varint,1,opt,name=colors_enabled,json=colorsEnabled,proto3" json:"colors_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogStyle) Reset()         { *m = LogStyle{} }
func (m *LogStyle) String() string { return proto.CompactTextString(m) }
func (*LogStyle) ProtoMessage()    {}
func (*LogStyle) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f1080a1433c7af1, []int{6}
}

func (m *LogStyle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogStyle.Unmarshal(m, b)
}
func (m *LogStyle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogStyle.Marshal(b, m, deterministic)
}
func (m *LogStyle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogStyle.Merge(m, src)
}
func (m *LogStyle) XXX_Size() int {
	return xxx_messageInfo_LogStyle.Size(m)
}
func (m *LogStyle) XXX_DiscardUnknown() {
	xxx_messageInfo_LogStyle.DiscardUnknown(m)
}

var xxx_messageInfo_LogStyle proto.InternalMessageInfo

func (m *LogStyle) GetColorsEnabled() bool {
	if m != nil {
		return m.ColorsEnabled
	}
	return false
}

func init() {
	proto.RegisterEnum("synclet.LogLevel", LogLevel_name, LogLevel_value)
	proto.RegisterType((*GetContainerIdForPodRequest)(nil), "synclet.GetContainerIdForPodRequest")
	proto.RegisterType((*GetContainerIdForPodReply)(nil), "synclet.GetContainerIdForPodReply")
	proto.RegisterType((*Cmd)(nil), "synclet.Cmd")
	proto.RegisterType((*UpdateContainerRequest)(nil), "synclet.UpdateContainerRequest")
	proto.RegisterType((*LogMessage)(nil), "synclet.LogMessage")
	proto.RegisterType((*UpdateContainerReply)(nil), "synclet.UpdateContainerReply")
	proto.RegisterType((*LogStyle)(nil), "synclet.LogStyle")
}

func init() { proto.RegisterFile("internal/synclet/synclet.proto", fileDescriptor_6f1080a1433c7af1) }

var fileDescriptor_6f1080a1433c7af1 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x9b, 0xb5, 0x5d, 0xd2, 0x93, 0x6e, 0x2b, 0x66, 0xa0, 0x76, 0x08, 0x56, 0xc2, 0xbf,
	0x0a, 0xa1, 0x06, 0x0a, 0x17, 0x08, 0xae, 0x68, 0xd7, 0x6d, 0x95, 0x0a, 0x45, 0x2e, 0xe5, 0x82,
	0x9b, 0xc8, 0x8d, 0x4d, 0x16, 0xc9, 0x89, 0x43, 0xe2, 0x15, 0x55, 0x42, 0xe2, 0xed, 0x78, 0x0c,
	0x9e, 0x05, 0xc5, 0x69, 0x32, 0xca, 0x3a, 0x76, 0x53, 0xdb, 0xe7, 0x3b, 0xa7, 0xe7, 0xe7, 0xcf,
	0x8e, 0xe1, 0x9e, 0x1f, 0x4a, 0x16, 0x87, 0x84, 0xdb, 0xc9, 0x32, 0x74, 0x39, 0x93, 0xf9, 0xd8,
	0x8d, 0x62, 0x21, 0x05, 0xd2, 0x57, 0x4b, 0xeb, 0x27, 0xdc, 0x39, 0x61, 0x72, 0x20, 0x42, 0x49,
	0xfc, 0x90, 0xc5, 0x23, 0x7a, 0x2c, 0xe2, 0x8f, 0x82, 0x62, 0xf6, 0xed, 0x9c, 0x25, 0x12, 0xdd,
	0x82, 0xed, 0x48, 0x50, 0xc7, 0xa7, 0x4d, 0xad, 0xad, 0x75, 0x6a, 0xb8, 0x1a, 0x09, 0x3a, 0xa2,
	0xa8, 0x05, 0x86, 0x1f, 0x10, 0x8f, 0xa5, 0xc2, 0x96, 0x12, 0x74, 0xb5, 0x1e, 0x51, 0xd4, 0x85,
	0x1a, 0x17, 0x9e, 0x93, 0xc8, 0x25, 0x67, 0xcd, 0x72, 0x5b, 0xeb, 0x98, 0xbd, 0x1b, 0xdd, 0xbc,
	0xf9, 0x58, 0x78, 0xd3, 0x54, 0xc0, 0x06, 0x5f, 0xcd, 0xac, 0x1f, 0xd0, 0xda, 0x0c, 0x10, 0xf1,
	0x25, 0x7a, 0x00, 0x75, 0x37, 0x57, 0x0a, 0x88, 0xd3, 0x12, 0x36, 0xdd, 0x8b, 0x7c, 0x64, 0x83,
	0x1e, 0xb0, 0x24, 0x21, 0x1e, 0x53, 0x2c, 0x66, 0xef, 0xe6, 0xdf, 0xfd, 0xde, 0x67, 0xd2, 0x69,
	0x09, 0xe7, 0x59, 0xfd, 0x1a, 0xe8, 0x69, 0x3d, 0x0b, 0xa5, 0xd5, 0x82, 0xf2, 0x20, 0xa0, 0x08,
	0x41, 0x85, 0xc4, 0xde, 0xa2, 0xa9, 0xb5, 0xcb, 0x9d, 0x1a, 0x56, 0x73, 0xeb, 0xb7, 0x06, 0xb7,
	0x67, 0x11, 0x25, 0x92, 0x15, 0x70, 0xb9, 0x2b, 0xf7, 0x37, 0x61, 0xad, 0x43, 0x1d, 0x82, 0x29,
	0x49, 0xec, 0x90, 0xd8, 0x3d, 0xf3, 0x17, 0x19, 0x58, 0x1d, 0x83, 0x24, 0xf1, 0xbb, 0x2c, 0x82,
	0x1e, 0xc3, 0xde, 0x57, 0x9f, 0xb3, 0xc4, 0x91, 0xc2, 0xa1, 0x8c, 0x33, 0x99, 0xba, 0x95, 0x76,
	0xdf, 0x51, 0xe1, 0x4f, 0xe2, 0x48, 0x05, 0x51, 0x07, 0x0c, 0x57, 0x04, 0x01, 0x09, 0x69, 0xd2,
	0xac, 0xb4, 0xcb, 0x1d, 0xb3, 0x57, 0x2f, 0xb6, 0x37, 0x08, 0x28, 0x2e, 0xd4, 0x75, 0xe7, 0xab,
	0xd7, 0x3b, 0x3f, 0x01, 0xb8, 0xf0, 0x07, 0x3d, 0x81, 0x2a, 0x67, 0x0b, 0xc6, 0xd5, 0x66, 0x76,
	0xd7, 0x2b, 0xc7, 0xa9, 0x80, 0x33, 0x1d, 0x35, 0xd7, 0xed, 0xae, 0x17, 0xbe, 0x5a, 0x63, 0xd8,
	0xbf, 0x64, 0x58, 0x7a, 0x8a, 0xaf, 0xc0, 0x4c, 0xc1, 0xf2, 0x2a, 0xed, 0xca, 0x43, 0xc2, 0xc0,
	0x8b, 0xb9, 0xf5, 0x02, 0x8c, 0x1c, 0x1a, 0x3d, 0x82, 0x5d, 0x57, 0x70, 0x11, 0x27, 0x0e, 0x0b,
	0xc9, 0x9c, 0xb3, 0xcc, 0x72, 0x03, 0xef, 0x64, 0xd1, 0x61, 0x16, 0x7c, 0xfa, 0x4c, 0x95, 0x28,
	0x5a, 0x64, 0x40, 0x65, 0xf4, 0xe1, 0x78, 0xd2, 0x28, 0x21, 0x13, 0xf4, 0xcf, 0x43, 0xdc, 0x9f,
	0x4c, 0x87, 0x0d, 0x0d, 0xd5, 0xa0, 0x7a, 0x34, 0xec, 0xcf, 0x4e, 0x1a, 0x5b, 0xbd, 0x5f, 0x1a,
	0xe8, 0xd3, 0x8c, 0x01, 0x51, 0xd8, 0xdf, 0x74, 0x0b, 0xd1, 0xc3, 0x82, 0xf2, 0x3f, 0x5f, 0xc9,
	0x81, 0x75, 0x4d, 0x56, 0xc4, 0x97, 0x56, 0xe9, 0xb9, 0x86, 0x66, 0xb0, 0xf7, 0x8f, 0x41, 0xe8,
	0xb0, 0x28, 0xdd, 0x7c, 0xd7, 0x0e, 0xee, 0x5e, 0x9d, 0xb0, 0xfa, 0xdb, 0xfe, 0x9b, 0x2f, 0xaf,
	0x3d, 0x5f, 0x9e, 0x9d, 0xcf, 0xbb, 0xae, 0x08, 0xec, 0xef, 0x7e, 0x48, 0x03, 0x9f, 0x73, 0x16,
	0x7a, 0xb6, 0xf4, 0xb9, 0xb4, 0x2f, 0x3d, 0x05, 0xea, 0x09, 0x78, 0xab, 0x7e, 0xe7, 0xdb, 0x6a,
	0x78, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x83, 0x66, 0xdc, 0x31, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SyncletClient is the client API for Synclet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SyncletClient interface {
	GetContainerIdForPod(ctx context.Context, in *GetContainerIdForPodRequest, opts ...grpc.CallOption) (Synclet_GetContainerIdForPodClient, error)
	// updates the specified container and then restarts it
	// (much functionality packed into one rpc to minimize latency)
	UpdateContainer(ctx context.Context, in *UpdateContainerRequest, opts ...grpc.CallOption) (Synclet_UpdateContainerClient, error)
}

type syncletClient struct {
	cc *grpc.ClientConn
}

func NewSyncletClient(cc *grpc.ClientConn) SyncletClient {
	return &syncletClient{cc}
}

func (c *syncletClient) GetContainerIdForPod(ctx context.Context, in *GetContainerIdForPodRequest, opts ...grpc.CallOption) (Synclet_GetContainerIdForPodClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Synclet_serviceDesc.Streams[0], "/synclet.Synclet/GetContainerIdForPod", opts...)
	if err != nil {
		return nil, err
	}
	x := &syncletGetContainerIdForPodClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Synclet_GetContainerIdForPodClient interface {
	Recv() (*GetContainerIdForPodReply, error)
	grpc.ClientStream
}

type syncletGetContainerIdForPodClient struct {
	grpc.ClientStream
}

func (x *syncletGetContainerIdForPodClient) Recv() (*GetContainerIdForPodReply, error) {
	m := new(GetContainerIdForPodReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *syncletClient) UpdateContainer(ctx context.Context, in *UpdateContainerRequest, opts ...grpc.CallOption) (Synclet_UpdateContainerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Synclet_serviceDesc.Streams[1], "/synclet.Synclet/UpdateContainer", opts...)
	if err != nil {
		return nil, err
	}
	x := &syncletUpdateContainerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Synclet_UpdateContainerClient interface {
	Recv() (*UpdateContainerReply, error)
	grpc.ClientStream
}

type syncletUpdateContainerClient struct {
	grpc.ClientStream
}

func (x *syncletUpdateContainerClient) Recv() (*UpdateContainerReply, error) {
	m := new(UpdateContainerReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SyncletServer is the server API for Synclet service.
type SyncletServer interface {
	GetContainerIdForPod(*GetContainerIdForPodRequest, Synclet_GetContainerIdForPodServer) error
	// updates the specified container and then restarts it
	// (much functionality packed into one rpc to minimize latency)
	UpdateContainer(*UpdateContainerRequest, Synclet_UpdateContainerServer) error
}

func RegisterSyncletServer(s *grpc.Server, srv SyncletServer) {
	s.RegisterService(&_Synclet_serviceDesc, srv)
}

func _Synclet_GetContainerIdForPod_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetContainerIdForPodRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SyncletServer).GetContainerIdForPod(m, &syncletGetContainerIdForPodServer{stream})
}

type Synclet_GetContainerIdForPodServer interface {
	Send(*GetContainerIdForPodReply) error
	grpc.ServerStream
}

type syncletGetContainerIdForPodServer struct {
	grpc.ServerStream
}

func (x *syncletGetContainerIdForPodServer) Send(m *GetContainerIdForPodReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Synclet_UpdateContainer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UpdateContainerRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SyncletServer).UpdateContainer(m, &syncletUpdateContainerServer{stream})
}

type Synclet_UpdateContainerServer interface {
	Send(*UpdateContainerReply) error
	grpc.ServerStream
}

type syncletUpdateContainerServer struct {
	grpc.ServerStream
}

func (x *syncletUpdateContainerServer) Send(m *UpdateContainerReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Synclet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "synclet.Synclet",
	HandlerType: (*SyncletServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetContainerIdForPod",
			Handler:       _Synclet_GetContainerIdForPod_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateContainer",
			Handler:       _Synclet_UpdateContainer_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/synclet/synclet.proto",
}
