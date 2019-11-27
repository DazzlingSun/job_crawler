// Code generated by protoc-gen-go. DO NOT EDIT.
// source: liverpc.proto

package liverpc // import "github.com/DazzlingSun/monitorService/src/basic/net/rpc/liverpc"

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

type Header struct {
	// APP_NAME.xxxxx , when separated by dot,
	// the first part is always app_name, the rest is undefined
	Caller      string `protobuf:"bytes,1,opt,name=caller,proto3" json:"caller,omitempty"`
	Uid         int64  `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Platform    string `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform,omitempty"`
	Src         string `protobuf:"bytes,4,opt,name=src,proto3" json:"src,omitempty"`
	TraceId     string `protobuf:"bytes,5,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	UserIp      string `protobuf:"bytes,7,opt,name=user_ip,json=userIp,proto3" json:"user_ip,omitempty"`
	SourceGroup string `protobuf:"bytes,8,opt,name=source_group,json=sourceGroup,proto3" json:"source_group,omitempty"`
	Buvid       string `protobuf:"bytes,9,opt,name=buvid,proto3" json:"buvid,omitempty"`
	// session data, format is http query
	// such as access_token=abc&SESS_DATA=def
	Sessdata2            string   `protobuf:"bytes,10,opt,name=sessdata2,proto3" json:"sessdata2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_liverpc_376c41dd15148bd6, []int{0}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (dst *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(dst, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetCaller() string {
	if m != nil {
		return m.Caller
	}
	return ""
}

func (m *Header) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Header) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *Header) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *Header) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *Header) GetUserIp() string {
	if m != nil {
		return m.UserIp
	}
	return ""
}

func (m *Header) GetSourceGroup() string {
	if m != nil {
		return m.SourceGroup
	}
	return ""
}

func (m *Header) GetBuvid() string {
	if m != nil {
		return m.Buvid
	}
	return ""
}

func (m *Header) GetSessdata2() string {
	if m != nil {
		return m.Sessdata2
	}
	return ""
}

// http is inside the protocol body
// {"body":..., "header":..., "http":...}
// this is used when a proxy forward a http request to a rpc request
type HTTP struct {
	IsHttps              int32             `protobuf:"varint,1,opt,name=is_https,json=isHttps,proto3" json:"is_https,omitempty"`
	Body                 string            `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Cookie               map[string]string `protobuf:"bytes,3,rep,name=cookie,proto3" json:"cookie,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Header               map[string]string `protobuf:"bytes,4,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Uri                  string            `protobuf:"bytes,5,opt,name=uri,proto3" json:"uri,omitempty"`
	Method               string            `protobuf:"bytes,6,opt,name=method,proto3" json:"method,omitempty"`
	Protocol             string            `protobuf:"bytes,7,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HTTP) Reset()         { *m = HTTP{} }
func (m *HTTP) String() string { return proto.CompactTextString(m) }
func (*HTTP) ProtoMessage()    {}
func (*HTTP) Descriptor() ([]byte, []int) {
	return fileDescriptor_liverpc_376c41dd15148bd6, []int{1}
}
func (m *HTTP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTP.Unmarshal(m, b)
}
func (m *HTTP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTP.Marshal(b, m, deterministic)
}
func (dst *HTTP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTP.Merge(dst, src)
}
func (m *HTTP) XXX_Size() int {
	return xxx_messageInfo_HTTP.Size(m)
}
func (m *HTTP) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTP.DiscardUnknown(m)
}

var xxx_messageInfo_HTTP proto.InternalMessageInfo

func (m *HTTP) GetIsHttps() int32 {
	if m != nil {
		return m.IsHttps
	}
	return 0
}

func (m *HTTP) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *HTTP) GetCookie() map[string]string {
	if m != nil {
		return m.Cookie
	}
	return nil
}

func (m *HTTP) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HTTP) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *HTTP) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *HTTP) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func init() {
	proto.RegisterType((*Header)(nil), "liverpc.Header")
	proto.RegisterType((*HTTP)(nil), "liverpc.HTTP")
	proto.RegisterMapType((map[string]string)(nil), "liverpc.HTTP.CookieEntry")
	proto.RegisterMapType((map[string]string)(nil), "liverpc.HTTP.HeaderEntry")
}

func init() { proto.RegisterFile("liverpc.proto", fileDescriptor_liverpc_376c41dd15148bd6) }

var fileDescriptor_liverpc_376c41dd15148bd6 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0xce, 0xd3, 0x30,
	0x10, 0xc7, 0x95, 0x26, 0x4d, 0x1a, 0x17, 0x24, 0x64, 0x21, 0xf0, 0xf7, 0x89, 0xa1, 0x2d, 0x4b,
	0x17, 0x1a, 0x51, 0x16, 0x60, 0x04, 0x21, 0xda, 0x0d, 0x45, 0x9d, 0x58, 0x22, 0xc7, 0x36, 0xad,
	0xd5, 0xa4, 0xb6, 0x6c, 0xa7, 0x52, 0x9e, 0x94, 0x87, 0xe0, 0x25, 0xd0, 0xd9, 0xa6, 0x74, 0x60,
	0x61, 0xbb, 0xdf, 0xdd, 0xfd, 0xed, 0xbb, 0xff, 0xa1, 0xa7, 0x9d, 0xbc, 0x0a, 0xa3, 0xd9, 0x46,
	0x1b, 0xe5, 0x14, 0x2e, 0x22, 0xae, 0x7e, 0x25, 0x28, 0xdf, 0x09, 0xca, 0x85, 0xc1, 0x2f, 0x50,
	0xce, 0x68, 0xd7, 0x09, 0x43, 0x92, 0x45, 0xb2, 0x2e, 0xeb, 0x48, 0xf8, 0x19, 0x4a, 0x07, 0xc9,
	0xc9, 0x64, 0x91, 0xac, 0xd3, 0x1a, 0x42, 0xfc, 0x88, 0x66, 0xba, 0xa3, 0xee, 0x87, 0x32, 0x3d,
	0x49, 0x7d, 0xef, 0x8d, 0xa1, 0xdb, 0x1a, 0x46, 0x32, 0x9f, 0x86, 0x10, 0x3f, 0xa0, 0x99, 0x33,
	0x94, 0x89, 0x46, 0x72, 0x32, 0xf5, 0xe9, 0xc2, 0xf3, 0x9e, 0xe3, 0x97, 0xa8, 0x18, 0xac, 0x30,
	0x8d, 0xd4, 0xa4, 0x08, 0x7f, 0x02, 0xee, 0x35, 0x5e, 0xa2, 0x27, 0x56, 0x0d, 0x86, 0x89, 0xe6,
	0x68, 0xd4, 0xa0, 0xc9, 0xcc, 0x57, 0xe7, 0x21, 0xf7, 0x15, 0x52, 0xf8, 0x39, 0x9a, 0xb6, 0xc3,
	0x55, 0x72, 0x52, 0xfa, 0x5a, 0x00, 0xfc, 0x0a, 0x95, 0x56, 0x58, 0xcb, 0xa9, 0xa3, 0x5b, 0x82,
	0x7c, 0xe5, 0x6f, 0x62, 0xf5, 0x73, 0x82, 0xb2, 0xdd, 0xe1, 0xf0, 0x0d, 0x66, 0x92, 0xb6, 0x39,
	0x39, 0xa7, 0xad, 0xdf, 0x76, 0x5a, 0x17, 0xd2, 0xee, 0x00, 0x31, 0x46, 0x59, 0xab, 0xf8, 0xe8,
	0xf7, 0x2d, 0x6b, 0x1f, 0xe3, 0xb7, 0x28, 0x67, 0x4a, 0x9d, 0xa5, 0x20, 0xe9, 0x22, 0x5d, 0xcf,
	0xb7, 0x0f, 0x9b, 0x3f, 0x76, 0xc2, 0x6b, 0x9b, 0xcf, 0xbe, 0xf6, 0xe5, 0xe2, 0xcc, 0x58, 0xc7,
	0x46, 0x90, 0x9c, 0xbc, 0xaf, 0x24, 0xfb, 0x97, 0x24, 0x78, 0x1e, 0x25, 0xa1, 0xd1, 0x1b, 0x6d,
	0x64, 0xf4, 0x08, 0x42, 0x38, 0x49, 0x2f, 0xdc, 0x49, 0x71, 0x92, 0x07, 0x7b, 0x02, 0xf9, 0x03,
	0xc0, 0x1d, 0x99, 0xea, 0xa2, 0x71, 0x37, 0x7e, 0xfc, 0x80, 0xe6, 0x77, 0xf3, 0xc0, 0xa3, 0x67,
	0x31, 0xc6, 0x93, 0x42, 0x08, 0xc6, 0x5d, 0x69, 0x37, 0x88, 0xb8, 0x61, 0x80, 0x8f, 0x93, 0xf7,
	0x09, 0x48, 0xef, 0xe6, 0xfa, 0x1f, 0xe9, 0xa7, 0xd7, 0xdf, 0x97, 0x47, 0xf5, 0x86, 0xa9, 0xbe,
	0x57, 0x97, 0xaa, 0x93, 0xad, 0xa1, 0x66, 0xac, 0x2e, 0xc2, 0x55, 0x46, 0xb3, 0x2a, 0x6e, 0xde,
	0xe6, 0x7e, 0xc8, 0x77, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x68, 0x55, 0x93, 0x8d, 0x02,
	0x00, 0x00,
}
