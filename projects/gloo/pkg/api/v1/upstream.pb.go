// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/upstream.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
// Upstreams represent destination for routing HTTP requests. Upstreams can be compared to
// [clusters](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/cds.proto) in Envoy terminology.
// Each upstream in Gloo has a type. Supported types include `static`, `kubernetes`, `aws`, `consul`, and more.
// Each upstream type is handled by a corresponding Gloo plugin.
type Upstream struct {
	// Type-specific configuration. Examples include static, kubernetes, and aws.
	// The type-specific config for the upstream is called a spec.
	UpstreamSpec *UpstreamSpec `protobuf:"bytes,2,opt,name=upstream_spec,json=upstreamSpec,proto3" json:"upstream_spec,omitempty"`
	// Status indicates the validation status of the resource. Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	// Upstreams and their configuration can be automatically by Gloo Discovery
	// if this upstream is created or modified by Discovery, metadata about the operation will be placed here.
	DiscoveryMetadata    *DiscoveryMetadata `protobuf:"bytes,8,opt,name=discovery_metadata,json=discoveryMetadata,proto3" json:"discovery_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Upstream) Reset()         { *m = Upstream{} }
func (m *Upstream) String() string { return proto.CompactTextString(m) }
func (*Upstream) ProtoMessage()    {}
func (*Upstream) Descriptor() ([]byte, []int) {
	return fileDescriptor_b74df493149f644d, []int{0}
}
func (m *Upstream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Upstream.Unmarshal(m, b)
}
func (m *Upstream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Upstream.Marshal(b, m, deterministic)
}
func (m *Upstream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Upstream.Merge(m, src)
}
func (m *Upstream) XXX_Size() int {
	return xxx_messageInfo_Upstream.Size(m)
}
func (m *Upstream) XXX_DiscardUnknown() {
	xxx_messageInfo_Upstream.DiscardUnknown(m)
}

var xxx_messageInfo_Upstream proto.InternalMessageInfo

func (m *Upstream) GetUpstreamSpec() *UpstreamSpec {
	if m != nil {
		return m.UpstreamSpec
	}
	return nil
}

func (m *Upstream) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Upstream) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Upstream) GetDiscoveryMetadata() *DiscoveryMetadata {
	if m != nil {
		return m.DiscoveryMetadata
	}
	return nil
}

// created by discovery services
type DiscoveryMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoveryMetadata) Reset()         { *m = DiscoveryMetadata{} }
func (m *DiscoveryMetadata) String() string { return proto.CompactTextString(m) }
func (*DiscoveryMetadata) ProtoMessage()    {}
func (*DiscoveryMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_b74df493149f644d, []int{1}
}
func (m *DiscoveryMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryMetadata.Unmarshal(m, b)
}
func (m *DiscoveryMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryMetadata.Marshal(b, m, deterministic)
}
func (m *DiscoveryMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryMetadata.Merge(m, src)
}
func (m *DiscoveryMetadata) XXX_Size() int {
	return xxx_messageInfo_DiscoveryMetadata.Size(m)
}
func (m *DiscoveryMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryMetadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Upstream)(nil), "gloo.solo.io.Upstream")
	proto.RegisterType((*DiscoveryMetadata)(nil), "gloo.solo.io.DiscoveryMetadata")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/upstream.proto", fileDescriptor_b74df493149f644d)
}

var fileDescriptor_b74df493149f644d = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0xd7, 0x4f,
	0xcf, 0xc9, 0xcf, 0xd7, 0x2f, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x29, 0x86, 0xf0, 0x12, 0x0b,
	0x32, 0xf5, 0xcb, 0x0c, 0xf5, 0x4b, 0x0b, 0x8a, 0x4b, 0x8a, 0x52, 0x13, 0x73, 0xf5, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0x78, 0x40, 0x72, 0x7a, 0x20, 0x6d, 0x7a, 0x99, 0xf9, 0x52, 0x22, 0xe9,
	0xf9, 0xe9, 0xf9, 0x60, 0x09, 0x7d, 0x10, 0x0b, 0xa2, 0x46, 0xca, 0x10, 0x8b, 0x05, 0x60, 0x3a,
	0x3b, 0xb3, 0x04, 0x66, 0x6c, 0x6e, 0x6a, 0x49, 0x62, 0x4a, 0x62, 0x49, 0x22, 0x54, 0x8b, 0x3e,
	0x11, 0x5a, 0x8a, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x49, 0xb0, 0x03, 0xc6, 0x87, 0x6a, 0xb1, 0x22,
	0xc9, 0xdf, 0x05, 0x39, 0xa5, 0xe9, 0x99, 0x79, 0x50, 0xeb, 0x94, 0xb6, 0x33, 0x71, 0x71, 0x84,
	0x42, 0x43, 0x42, 0xc8, 0x9e, 0x8b, 0x17, 0x16, 0x2a, 0xf1, 0xc5, 0x05, 0xa9, 0xc9, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x52, 0x7a, 0xc8, 0x61, 0xa3, 0x07, 0x53, 0x1e, 0x5c, 0x90, 0x9a,
	0x1c, 0xc4, 0x53, 0x8a, 0xc4, 0x13, 0x72, 0xe7, 0x62, 0x83, 0x78, 0x46, 0x82, 0x0d, 0xac, 0x53,
	0x44, 0x2f, 0x39, 0xbf, 0x28, 0x15, 0xae, 0x33, 0x18, 0x2c, 0xe7, 0x24, 0x79, 0xe2, 0x9e, 0x3c,
	0xc3, 0xa7, 0x7b, 0xf2, 0x82, 0x25, 0xa9, 0xc5, 0x25, 0x29, 0x99, 0x69, 0x69, 0x56, 0x4a, 0x99,
	0xe9, 0x79, 0xf9, 0x45, 0xa9, 0x4a, 0x41, 0x50, 0xed, 0x42, 0x16, 0x5c, 0x1c, 0xb0, 0x80, 0x94,
	0x60, 0x07, 0x1b, 0x25, 0x86, 0x6a, 0x94, 0x2f, 0x54, 0xd6, 0x89, 0x05, 0x64, 0x58, 0x10, 0x5c,
	0xb5, 0x90, 0x1f, 0x97, 0x50, 0x4a, 0x66, 0x71, 0x72, 0x7e, 0x59, 0x6a, 0x51, 0x65, 0x3c, 0xdc,
	0x0c, 0x0e, 0xb0, 0x19, 0xf2, 0xa8, 0x1e, 0x71, 0x81, 0xa9, 0x83, 0x19, 0x16, 0x24, 0x98, 0x82,
	0x2e, 0x64, 0x25, 0xde, 0xf4, 0x91, 0x85, 0x85, 0x8b, 0xa9, 0xb4, 0xb8, 0xe9, 0x23, 0x0b, 0xb7,
	0x10, 0x27, 0xcc, 0xbb, 0xc5, 0x4a, 0xc2, 0x5c, 0x82, 0x18, 0x06, 0x38, 0x99, 0xad, 0x78, 0x24,
	0xc7, 0x18, 0x65, 0x40, 0x5c, 0x84, 0x14, 0x64, 0xa7, 0x43, 0x23, 0x25, 0x89, 0x0d, 0x1c, 0x1b,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0x50, 0x38, 0xd4, 0xc3, 0x02, 0x00, 0x00,
}

func (this *Upstream) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream)
	if !ok {
		that2, ok := that.(Upstream)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.UpstreamSpec.Equal(that1.UpstreamSpec) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.DiscoveryMetadata.Equal(that1.DiscoveryMetadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DiscoveryMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DiscoveryMetadata)
	if !ok {
		that2, ok := that.(DiscoveryMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
