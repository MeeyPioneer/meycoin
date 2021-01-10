// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PeerRole int32

const (
	PeerRole_LegacyVersion PeerRole = 0
	PeerRole_Producer      PeerRole = 1
	PeerRole_Watcher       PeerRole = 2
	PeerRole_Agent         PeerRole = 3
)

var PeerRole_name = map[int32]string{
	0: "LegacyVersion",
	1: "Producer",
	2: "Watcher",
	3: "Agent",
}
var PeerRole_value = map[string]int32{
	"LegacyVersion": 0,
	"Producer":      1,
	"Watcher":       2,
	"Agent":         3,
}

func (x PeerRole) String() string {
	return proto.EnumName(PeerRole_name, int32(x))
}
func (PeerRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_node_2c4eb40676311241, []int{0}
}

// PeerAddress contains static information of peer and addresses to connect peer
type PeerAddress struct {
	// @Deprecated advertised address and port will be in addresses field in meycoin v2.
	// address is string representation of ip address or domain name.
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// @Deprecated
	Port                 uint32   `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	PeerID               []byte   `protobuf:"bytes,3,opt,name=peerID,proto3" json:"peerID,omitempty"`
	Role                 PeerRole `protobuf:"varint,4,opt,name=role,enum=types.PeerRole" json:"role,omitempty"`
	Version              string   `protobuf:"bytes,5,opt,name=version" json:"version,omitempty"`
	Addresses            []string `protobuf:"bytes,6,rep,name=addresses" json:"addresses,omitempty"`
	ProducerIDs          [][]byte `protobuf:"bytes,7,rep,name=producerIDs,proto3" json:"producerIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerAddress) Reset()         { *m = PeerAddress{} }
func (m *PeerAddress) String() string { return proto.CompactTextString(m) }
func (*PeerAddress) ProtoMessage()    {}
func (*PeerAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_2c4eb40676311241, []int{0}
}
func (m *PeerAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerAddress.Unmarshal(m, b)
}
func (m *PeerAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerAddress.Marshal(b, m, deterministic)
}
func (dst *PeerAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerAddress.Merge(dst, src)
}
func (m *PeerAddress) XXX_Size() int {
	return xxx_messageInfo_PeerAddress.Size(m)
}
func (m *PeerAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerAddress.DiscardUnknown(m)
}

var xxx_messageInfo_PeerAddress proto.InternalMessageInfo

func (m *PeerAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PeerAddress) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *PeerAddress) GetPeerID() []byte {
	if m != nil {
		return m.PeerID
	}
	return nil
}

func (m *PeerAddress) GetRole() PeerRole {
	if m != nil {
		return m.Role
	}
	return PeerRole_LegacyVersion
}

func (m *PeerAddress) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *PeerAddress) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *PeerAddress) GetProducerIDs() [][]byte {
	if m != nil {
		return m.ProducerIDs
	}
	return nil
}

type AgentCertificate struct {
	CertVersion uint32 `protobuf:"varint,1,opt,name=certVersion" json:"certVersion,omitempty"`
	BPID        []byte `protobuf:"bytes,2,opt,name=BPID,proto3" json:"BPID,omitempty"`
	BPPubKey    []byte `protobuf:"bytes,3,opt,name=BPPubKey,proto3" json:"BPPubKey,omitempty"`
	// CreateTime is the number of nanoseconds elapsed since January 1, 1970 UTC
	CreateTime int64 `protobuf:"varint,4,opt,name=createTime" json:"createTime,omitempty"`
	// CreateTime is the number of nanoseconds elapsed since January 1, 1970 UTC
	ExpireTime           int64    `protobuf:"varint,5,opt,name=expireTime" json:"expireTime,omitempty"`
	AgentID              []byte   `protobuf:"bytes,6,opt,name=agentID,proto3" json:"agentID,omitempty"`
	AgentAddress         [][]byte `protobuf:"bytes,7,rep,name=AgentAddress,proto3" json:"AgentAddress,omitempty"`
	Signature            []byte   `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentCertificate) Reset()         { *m = AgentCertificate{} }
func (m *AgentCertificate) String() string { return proto.CompactTextString(m) }
func (*AgentCertificate) ProtoMessage()    {}
func (*AgentCertificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_2c4eb40676311241, []int{1}
}
func (m *AgentCertificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentCertificate.Unmarshal(m, b)
}
func (m *AgentCertificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentCertificate.Marshal(b, m, deterministic)
}
func (dst *AgentCertificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentCertificate.Merge(dst, src)
}
func (m *AgentCertificate) XXX_Size() int {
	return xxx_messageInfo_AgentCertificate.Size(m)
}
func (m *AgentCertificate) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentCertificate.DiscardUnknown(m)
}

var xxx_messageInfo_AgentCertificate proto.InternalMessageInfo

func (m *AgentCertificate) GetCertVersion() uint32 {
	if m != nil {
		return m.CertVersion
	}
	return 0
}

func (m *AgentCertificate) GetBPID() []byte {
	if m != nil {
		return m.BPID
	}
	return nil
}

func (m *AgentCertificate) GetBPPubKey() []byte {
	if m != nil {
		return m.BPPubKey
	}
	return nil
}

func (m *AgentCertificate) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *AgentCertificate) GetExpireTime() int64 {
	if m != nil {
		return m.ExpireTime
	}
	return 0
}

func (m *AgentCertificate) GetAgentID() []byte {
	if m != nil {
		return m.AgentID
	}
	return nil
}

func (m *AgentCertificate) GetAgentAddress() [][]byte {
	if m != nil {
		return m.AgentAddress
	}
	return nil
}

func (m *AgentCertificate) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*PeerAddress)(nil), "types.PeerAddress")
	proto.RegisterType((*AgentCertificate)(nil), "types.AgentCertificate")
	proto.RegisterEnum("types.PeerRole", PeerRole_name, PeerRole_value)
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_node_2c4eb40676311241) }

var fileDescriptor_node_2c4eb40676311241 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x4f, 0x4f, 0xe3, 0x30,
	0x10, 0xc5, 0xd7, 0x4d, 0x93, 0x26, 0xd3, 0x74, 0x37, 0x3b, 0x87, 0x95, 0xb5, 0x42, 0x28, 0x2a,
	0x97, 0x88, 0x43, 0x0f, 0xf0, 0x09, 0xfa, 0xe7, 0x52, 0xc1, 0x21, 0xb2, 0x10, 0x9c, 0xd3, 0x64,
	0x28, 0x91, 0x4a, 0x1c, 0x39, 0x2e, 0xa2, 0x37, 0x3e, 0x23, 0x9f, 0x08, 0xd9, 0x71, 0x69, 0xb9,
	0xcd, 0x7b, 0x23, 0xe7, 0xe7, 0xf7, 0x62, 0x80, 0x46, 0x56, 0x34, 0x6b, 0x95, 0xd4, 0x12, 0x7d,
	0x7d, 0x68, 0xa9, 0x9b, 0x7e, 0x32, 0x18, 0xe7, 0x44, 0x6a, 0x5e, 0x55, 0x8a, 0xba, 0x0e, 0x39,
	0x8c, 0x8a, 0x7e, 0xe4, 0x2c, 0x65, 0x59, 0x24, 0x8e, 0x12, 0x11, 0x86, 0xad, 0x54, 0x9a, 0x0f,
	0x52, 0x96, 0x4d, 0x84, 0x9d, 0xf1, 0x1f, 0x04, 0x2d, 0x91, 0x5a, 0xaf, 0xb8, 0x97, 0xb2, 0x2c,
	0x16, 0x4e, 0xe1, 0x15, 0x0c, 0x95, 0xdc, 0x11, 0x1f, 0xa6, 0x2c, 0xfb, 0x7d, 0xf3, 0x67, 0x66,
	0x59, 0x33, 0xc3, 0x11, 0x72, 0x47, 0xc2, 0x2e, 0x0d, 0xea, 0x8d, 0x54, 0x57, 0xcb, 0x86, 0xfb,
	0x3d, 0xca, 0x49, 0xbc, 0x80, 0xc8, 0x51, 0xa9, 0xe3, 0x41, 0xea, 0x65, 0x91, 0x38, 0x19, 0x98,
	0xc2, 0xb8, 0x55, 0xb2, 0xda, 0x97, 0x06, 0xd5, 0xf1, 0x51, 0xea, 0x65, 0xb1, 0x38, 0xb7, 0xa6,
	0x1f, 0x03, 0x48, 0xe6, 0x5b, 0x6a, 0xf4, 0x92, 0x94, 0xae, 0x9f, 0xeb, 0xb2, 0xd0, 0x64, 0x8e,
	0x95, 0xa4, 0xf4, 0xa3, 0x43, 0x32, 0x1b, 0xe3, 0xdc, 0x32, 0x09, 0x17, 0xf9, 0x7a, 0x65, 0x13,
	0xc6, 0xc2, 0xce, 0xf8, 0x1f, 0xc2, 0x45, 0x9e, 0xef, 0x37, 0x77, 0x74, 0x70, 0x19, 0xbf, 0x35,
	0x5e, 0x02, 0x94, 0x8a, 0x0a, 0x4d, 0x0f, 0xf5, 0x6b, 0x9f, 0xd5, 0x13, 0x67, 0x8e, 0xd9, 0xd3,
	0x7b, 0x5b, 0xab, 0x7e, 0xef, 0xf7, 0xfb, 0x93, 0x63, 0xbb, 0x36, 0xb7, 0x5c, 0xaf, 0x78, 0x60,
	0x3f, 0x7d, 0x94, 0x38, 0x85, 0xd8, 0xde, 0xdf, 0xfd, 0x15, 0x97, 0xf1, 0x87, 0x67, 0x4a, 0xea,
	0xea, 0x6d, 0x53, 0xe8, 0xbd, 0x22, 0x1e, 0xda, 0xf3, 0x27, 0xe3, 0x7a, 0x09, 0xe1, 0xb1, 0x6e,
	0xfc, 0x0b, 0x93, 0x7b, 0xda, 0x16, 0xe5, 0xc1, 0x05, 0x4d, 0x7e, 0x61, 0x0c, 0x61, 0xee, 0x0a,
	0x4b, 0x18, 0x8e, 0x61, 0xf4, 0x54, 0xe8, 0xf2, 0x85, 0x54, 0x32, 0xc0, 0x08, 0x7c, 0xcb, 0x49,
	0xbc, 0x4d, 0x60, 0x9f, 0xca, 0xed, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xb6, 0x20, 0x9b,
	0x38, 0x02, 0x00, 0x00,
}
