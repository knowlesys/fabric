// Code generated by protoc-gen-go.
// source: common/common.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/common.proto
	common/configuration.proto

It has these top-level messages:
	Header
	ChainHeader
	SignatureHeader
	Payload
	Envelope
	Block
	BlockHeader
	BlockData
	BlockMetadata
	ConfigurationEnvelope
	SignedConfigurationItem
	ConfigurationItem
	ConfigurationSignature
	Policy
	SignaturePolicyEnvelope
	SignaturePolicy
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// These status codes are intended to resemble selected HTTP status codes
type Status int32

const (
	Status_UNKNOWN               Status = 0
	Status_SUCCESS               Status = 200
	Status_BAD_REQUEST           Status = 400
	Status_FORBIDDEN             Status = 403
	Status_NOT_FOUND             Status = 404
	Status_INTERNAL_SERVER_ERROR Status = 500
	Status_SERVICE_UNAVAILABLE   Status = 503
)

var Status_name = map[int32]string{
	0:   "UNKNOWN",
	200: "SUCCESS",
	400: "BAD_REQUEST",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	500: "INTERNAL_SERVER_ERROR",
	503: "SERVICE_UNAVAILABLE",
}
var Status_value = map[string]int32{
	"UNKNOWN":               0,
	"SUCCESS":               200,
	"BAD_REQUEST":           400,
	"FORBIDDEN":             403,
	"NOT_FOUND":             404,
	"INTERNAL_SERVER_ERROR": 500,
	"SERVICE_UNAVAILABLE":   503,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HeaderType int32

const (
	HeaderType_MESSAGE                   HeaderType = 0
	HeaderType_CONFIGURATION_TRANSACTION HeaderType = 1
	HeaderType_CONFIGURATION_ITEM        HeaderType = 2
	HeaderType_ENDORSER_TRANSACTION      HeaderType = 3
)

var HeaderType_name = map[int32]string{
	0: "MESSAGE",
	1: "CONFIGURATION_TRANSACTION",
	2: "CONFIGURATION_ITEM",
	3: "ENDORSER_TRANSACTION",
}
var HeaderType_value = map[string]int32{
	"MESSAGE":                   0,
	"CONFIGURATION_TRANSACTION": 1,
	"CONFIGURATION_ITEM":        2,
	"ENDORSER_TRANSACTION":      3,
}

func (x HeaderType) String() string {
	return proto.EnumName(HeaderType_name, int32(x))
}
func (HeaderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Header struct {
	ChainHeader     *ChainHeader     `protobuf:"bytes,1,opt,name=chainHeader" json:"chainHeader,omitempty"`
	SignatureHeader *SignatureHeader `protobuf:"bytes,2,opt,name=signatureHeader" json:"signatureHeader,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Header) GetChainHeader() *ChainHeader {
	if m != nil {
		return m.ChainHeader
	}
	return nil
}

func (m *Header) GetSignatureHeader() *SignatureHeader {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

// Header is a generic replay prevention and identity message to include in a signed payload
type ChainHeader struct {
	Type int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	// Timestamp is the local time when the message was created
	// by the sender
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	// Identifier of the chain this message is bound for
	ChainID []byte `protobuf:"bytes,4,opt,name=chainID,proto3" json:"chainID,omitempty"`
}

func (m *ChainHeader) Reset()                    { *m = ChainHeader{} }
func (m *ChainHeader) String() string            { return proto.CompactTextString(m) }
func (*ChainHeader) ProtoMessage()               {}
func (*ChainHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChainHeader) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type SignatureHeader struct {
	// Creator of the message, specified as a certificate chain
	Creator []byte `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// Arbitrary number that may only be used once. Can be used to detect replay attacks.
	Nonce []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// The epoch in which this header was generated, where epoch is defined based on block height
	Epoch uint64 `protobuf:"varint,3,opt,name=epoch" json:"epoch,omitempty"`
}

func (m *SignatureHeader) Reset()                    { *m = SignatureHeader{} }
func (m *SignatureHeader) String() string            { return proto.CompactTextString(m) }
func (*SignatureHeader) ProtoMessage()               {}
func (*SignatureHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Payload is the message contents (and header to allow for signing)
type Payload struct {
	// Header is included to provide identity and prevent replay
	Header *Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// Data, the encoding of which is defined by the type in the header
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Payload) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

// Envelope wraps a Payload with a signature so that the message may be authenticated
type Envelope struct {
	// A marshaled Payload
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// A signature by the creator specified in the Payload header
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// This is finalized block structure to be shared among the orderer and peer
// Note that the BlockHeader chains to the previous BlockHeader, and the BlockData hash is embedded
// in the BlockHeader.  This makes it natural and obvious that the Data is included in the hash, but
// the Metadata is not.
type Block struct {
	Header   *BlockHeader   `protobuf:"bytes,1,opt,name=Header" json:"Header,omitempty"`
	Data     *BlockData     `protobuf:"bytes,2,opt,name=Data" json:"Data,omitempty"`
	Metadata *BlockMetadata `protobuf:"bytes,3,opt,name=Metadata" json:"Metadata,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetData() *BlockData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Block) GetMetadata() *BlockMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type BlockHeader struct {
	Number       uint64 `protobuf:"varint,1,opt,name=Number" json:"Number,omitempty"`
	PreviousHash []byte `protobuf:"bytes,2,opt,name=PreviousHash,proto3" json:"PreviousHash,omitempty"`
	DataHash     []byte `protobuf:"bytes,3,opt,name=DataHash,proto3" json:"DataHash,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type BlockData struct {
	Data [][]byte `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (m *BlockData) Reset()                    { *m = BlockData{} }
func (m *BlockData) String() string            { return proto.CompactTextString(m) }
func (*BlockData) ProtoMessage()               {}
func (*BlockData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type BlockMetadata struct {
	Metadata [][]byte `protobuf:"bytes,1,rep,name=Metadata,proto3" json:"Metadata,omitempty"`
}

func (m *BlockMetadata) Reset()                    { *m = BlockMetadata{} }
func (m *BlockMetadata) String() string            { return proto.CompactTextString(m) }
func (*BlockMetadata) ProtoMessage()               {}
func (*BlockMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*Header)(nil), "common.Header")
	proto.RegisterType((*ChainHeader)(nil), "common.ChainHeader")
	proto.RegisterType((*SignatureHeader)(nil), "common.SignatureHeader")
	proto.RegisterType((*Payload)(nil), "common.Payload")
	proto.RegisterType((*Envelope)(nil), "common.Envelope")
	proto.RegisterType((*Block)(nil), "common.Block")
	proto.RegisterType((*BlockHeader)(nil), "common.BlockHeader")
	proto.RegisterType((*BlockData)(nil), "common.BlockData")
	proto.RegisterType((*BlockMetadata)(nil), "common.BlockMetadata")
	proto.RegisterEnum("common.Status", Status_name, Status_value)
	proto.RegisterEnum("common.HeaderType", HeaderType_name, HeaderType_value)
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 667 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xad, 0xeb, 0xfc, 0x34, 0xd7, 0xf9, 0x5a, 0x7f, 0xd3, 0x1f, 0x42, 0x04, 0x6a, 0x65, 0x09,
	0x54, 0xb5, 0x22, 0x11, 0x45, 0x48, 0x6c, 0x9d, 0x64, 0xda, 0x5a, 0xb4, 0xe3, 0x32, 0x76, 0x5a,
	0x89, 0x4d, 0xe4, 0x24, 0xd3, 0x24, 0x90, 0x64, 0x22, 0xdb, 0xa9, 0xd4, 0x2d, 0x5b, 0x24, 0x40,
	0x82, 0x87, 0xe2, 0x81, 0x90, 0xd8, 0xa2, 0x99, 0xb1, 0xdd, 0xb8, 0xab, 0xcc, 0xb9, 0xe7, 0xdc,
	0xeb, 0x73, 0xcf, 0xc4, 0x86, 0xed, 0x01, 0x9f, 0xcd, 0xf8, 0xbc, 0xa9, 0x7e, 0x1a, 0x8b, 0x90,
	0xc7, 0x1c, 0x95, 0x14, 0xaa, 0xef, 0x8f, 0x38, 0x1f, 0x4d, 0x59, 0x53, 0x56, 0xfb, 0xcb, 0xdb,
	0x66, 0x3c, 0x99, 0xb1, 0x28, 0x0e, 0x66, 0x0b, 0x25, 0xb4, 0xbe, 0x68, 0x50, 0x3a, 0x67, 0xc1,
	0x90, 0x85, 0xe8, 0x2d, 0x18, 0x83, 0x71, 0x30, 0x99, 0x2b, 0x58, 0xd3, 0x0e, 0xb4, 0x43, 0xe3,
	0x64, 0xbb, 0x91, 0xcc, 0x6d, 0x3f, 0x50, 0x74, 0x55, 0x87, 0x6c, 0xd8, 0x8a, 0x26, 0xa3, 0x79,
	0x10, 0x2f, 0x43, 0x96, 0xb4, 0xae, 0xcb, 0xd6, 0x27, 0x69, 0xab, 0x97, 0xa7, 0xe9, 0x63, 0xbd,
	0xf5, 0x5d, 0x03, 0x63, 0x65, 0x3e, 0x42, 0x50, 0x88, 0xef, 0x17, 0x4c, 0x5a, 0x28, 0x52, 0x79,
	0x46, 0x35, 0x28, 0xdf, 0xb1, 0x30, 0x9a, 0xf0, 0xb9, 0x1c, 0x5f, 0xa4, 0x29, 0x44, 0xef, 0xa0,
	0x92, 0x6d, 0x55, 0xd3, 0xe5, 0xa3, 0xeb, 0x0d, 0xb5, 0x77, 0x23, 0xdd, 0xbb, 0xe1, 0xa7, 0x0a,
	0xfa, 0x20, 0x16, 0x33, 0xe5, 0x26, 0x4e, 0xa7, 0x56, 0x38, 0xd0, 0x0e, 0xab, 0x34, 0x85, 0xd6,
	0x0d, 0x6c, 0x3d, 0x72, 0x2d, 0xc5, 0x21, 0x0b, 0x62, 0xae, 0xa2, 0x11, 0x62, 0x05, 0xd1, 0x0e,
	0x14, 0xe7, 0x7c, 0x3e, 0x60, 0xd2, 0x58, 0x95, 0x2a, 0x20, 0xaa, 0x6c, 0xc1, 0x07, 0x63, 0x69,
	0xa9, 0x40, 0x15, 0xb0, 0x30, 0x94, 0xaf, 0x82, 0xfb, 0x29, 0x0f, 0x86, 0xe8, 0x25, 0x94, 0xc6,
	0xab, 0x51, 0x6f, 0xa6, 0x79, 0x25, 0x31, 0x25, 0xac, 0x48, 0x63, 0x18, 0xc4, 0x41, 0x32, 0x5d,
	0x9e, 0xad, 0x16, 0x6c, 0xe0, 0xf9, 0x1d, 0x9b, 0x72, 0x95, 0xcc, 0x42, 0x8d, 0x4c, 0x8d, 0x25,
	0x10, 0x3d, 0x83, 0x4a, 0x16, 0x75, 0xd2, 0xfe, 0x50, 0xb0, 0xbe, 0x69, 0x50, 0x6c, 0x4d, 0xf9,
	0xe0, 0x33, 0x3a, 0x4e, 0xff, 0x03, 0x8f, 0x2f, 0x5d, 0xd2, 0xa9, 0x9d, 0x24, 0x87, 0x17, 0x50,
	0xe8, 0xa4, 0x76, 0x8c, 0x93, 0xff, 0x73, 0x52, 0x41, 0x50, 0x49, 0xa3, 0xd7, 0xb0, 0x71, 0xc9,
	0xe2, 0x40, 0x3a, 0x57, 0x97, 0xb2, 0x9b, 0x93, 0xa6, 0x24, 0xcd, 0x64, 0x16, 0x03, 0x63, 0xe5,
	0x81, 0x68, 0x0f, 0x4a, 0x64, 0x39, 0xeb, 0x27, 0xae, 0x0a, 0x34, 0x41, 0xc8, 0x82, 0xea, 0x55,
	0xc8, 0xee, 0x26, 0x7c, 0x19, 0x9d, 0x07, 0xd1, 0x38, 0x59, 0x2c, 0x57, 0x43, 0x75, 0xd8, 0x10,
	0x2e, 0x24, 0xaf, 0x4b, 0x3e, 0xc3, 0xd6, 0x3e, 0x54, 0x32, 0xb3, 0x22, 0x5c, 0xb9, 0x8d, 0x76,
	0xa0, 0x8b, 0x70, 0xc5, 0xd9, 0x3a, 0x86, 0xff, 0x72, 0x16, 0xc5, 0xb4, 0x6c, 0x17, 0x25, 0xcc,
	0xf0, 0xd1, 0x57, 0x0d, 0x4a, 0x5e, 0x1c, 0xc4, 0xcb, 0x08, 0x19, 0x50, 0xee, 0x92, 0xf7, 0xc4,
	0xbd, 0x21, 0xe6, 0x1a, 0xaa, 0x42, 0xd9, 0xeb, 0xb6, 0xdb, 0xd8, 0xf3, 0xcc, 0xdf, 0x1a, 0x32,
	0xc1, 0x68, 0xd9, 0x9d, 0x1e, 0xc5, 0x1f, 0xba, 0xd8, 0xf3, 0xcd, 0x1f, 0x3a, 0xda, 0x84, 0xca,
	0xa9, 0x4b, 0x5b, 0x4e, 0xa7, 0x83, 0x89, 0xf9, 0x53, 0x62, 0xe2, 0xfa, 0xbd, 0x53, 0xb7, 0x4b,
	0x3a, 0xe6, 0x2f, 0x1d, 0xd5, 0x61, 0xd7, 0x21, 0x3e, 0xa6, 0xc4, 0xbe, 0xe8, 0x79, 0x98, 0x5e,
	0x63, 0xda, 0xc3, 0x94, 0xba, 0xd4, 0xfc, 0xa3, 0xa3, 0x1a, 0x6c, 0x8b, 0x92, 0xd3, 0xc6, 0xbd,
	0x2e, 0xb1, 0xaf, 0x6d, 0xe7, 0xc2, 0x6e, 0x5d, 0x60, 0xf3, 0xaf, 0x7e, 0xf4, 0x09, 0x40, 0xa5,
	0xe7, 0x8b, 0x77, 0xc6, 0x80, 0xf2, 0x25, 0xf6, 0x3c, 0xfb, 0x0c, 0x9b, 0x6b, 0xe8, 0x39, 0x3c,
	0x6d, 0xbb, 0xe4, 0xd4, 0x39, 0xeb, 0x52, 0xdb, 0x77, 0x5c, 0xd2, 0xf3, 0xa9, 0x4d, 0x3c, 0xbb,
	0x2d, 0xce, 0xa6, 0x86, 0xf6, 0x00, 0xe5, 0x69, 0xc7, 0xc7, 0x97, 0xe6, 0x3a, 0xaa, 0xc1, 0x0e,
	0x26, 0x1d, 0x97, 0x7a, 0x98, 0xe6, 0x3a, 0xf4, 0xd6, 0xab, 0x8f, 0xc7, 0xa3, 0x49, 0x3c, 0x5e,
	0xf6, 0xc5, 0xbd, 0x36, 0xc7, 0xf7, 0x0b, 0x16, 0x4e, 0xd9, 0x70, 0xc4, 0xc2, 0xe6, 0x6d, 0xd0,
	0x0f, 0x27, 0x03, 0xf5, 0xd1, 0x89, 0x92, 0x0f, 0x53, 0xbf, 0x24, 0xe1, 0x9b, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x0e, 0x0d, 0x13, 0x2b, 0xb0, 0x04, 0x00, 0x00,
}