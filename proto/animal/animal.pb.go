// Code generated by protoc-gen-go. DO NOT EDIT.
// source: animal.proto

package animal

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Animal struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32    `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	VesselId             string   `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Animal) Reset()         { *m = Animal{} }
func (m *Animal) String() string { return proto.CompactTextString(m) }
func (*Animal) ProtoMessage()    {}
func (*Animal) Descriptor() ([]byte, []int) {
	return fileDescriptor_900753348fe370c9, []int{0}
}

func (m *Animal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Animal.Unmarshal(m, b)
}
func (m *Animal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Animal.Marshal(b, m, deterministic)
}
func (m *Animal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Animal.Merge(m, src)
}
func (m *Animal) XXX_Size() int {
	return xxx_messageInfo_Animal.Size(m)
}
func (m *Animal) XXX_DiscardUnknown() {
	xxx_messageInfo_Animal.DiscardUnknown(m)
}

var xxx_messageInfo_Animal proto.InternalMessageInfo

func (m *Animal) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Animal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Animal) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Animal) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_900753348fe370c9, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type Response struct {
	Created              bool     `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Animal               *Animal  `protobuf:"bytes,2,opt,name=animal,proto3" json:"animal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_900753348fe370c9, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetAnimal() *Animal {
	if m != nil {
		return m.Animal
	}
	return nil
}

type Test struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_900753348fe370c9, []int{3}
}

func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (m *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(m, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Animal)(nil), "animal.Animal")
	proto.RegisterType((*GetRequest)(nil), "animal.GetRequest")
	proto.RegisterType((*Response)(nil), "animal.Response")
	proto.RegisterType((*Test)(nil), "animal.Test")
}

func init() { proto.RegisterFile("animal.proto", fileDescriptor_900753348fe370c9) }

var fileDescriptor_900753348fe370c9 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x6d, 0x75, 0x63, 0x77, 0x76, 0x5d, 0x64, 0x0e, 0x12, 0xf4, 0x52, 0x72, 0x90, 0x3d,
	0xed, 0x61, 0xf5, 0x05, 0x3c, 0x88, 0x08, 0x9e, 0xa2, 0x77, 0xa9, 0xcd, 0xa0, 0x03, 0x6b, 0x53,
	0x33, 0xb1, 0xbe, 0xbe, 0x90, 0x34, 0xe8, 0x2d, 0xff, 0x37, 0x0c, 0xf9, 0xfe, 0x81, 0x75, 0x37,
	0xf0, 0x67, 0x77, 0xd8, 0x8d, 0xc1, 0x47, 0x8f, 0x2a, 0x27, 0xe3, 0x41, 0xdd, 0xa5, 0x17, 0x6e,
	0xa0, 0x66, 0xa7, 0xab, 0xb6, 0xda, 0x2e, 0x6d, 0xcd, 0x0e, 0x5b, 0x58, 0x39, 0x92, 0x3e, 0xf0,
	0x18, 0xd9, 0x0f, 0xba, 0x4e, 0x83, 0xff, 0x08, 0x2f, 0x40, 0xfd, 0x10, 0xbf, 0x7f, 0x44, 0x7d,
	0xdc, 0x56, 0xdb, 0x85, 0x9d, 0x13, 0x5e, 0xc1, 0x72, 0x22, 0x11, 0x3a, 0xbc, 0xb2, 0xd3, 0x8b,
	0xb4, 0xd7, 0x64, 0xf0, 0xe8, 0xcc, 0x1a, 0xe0, 0x81, 0xa2, 0xa5, 0xaf, 0x6f, 0x92, 0x68, 0x9e,
	0xa0, 0xb1, 0x24, 0xa3, 0x1f, 0x84, 0x50, 0xc3, 0x69, 0x1f, 0xa8, 0x8b, 0x94, 0x2d, 0x1a, 0x5b,
	0x22, 0x5e, 0xc3, 0xac, 0x9b, 0x2c, 0x56, 0xfb, 0xcd, 0x6e, 0xee, 0x92, 0xd5, 0x6d, 0x29, 0xa3,
	0xe0, 0xe4, 0x85, 0x24, 0xee, 0xef, 0xe1, 0x2c, 0x4f, 0x9e, 0x29, 0x4c, 0xdc, 0x13, 0xde, 0xa6,
	0x4f, 0x33, 0x13, 0xc4, 0xb2, 0xfe, 0x27, 0x72, 0x79, 0x5e, 0x58, 0xd1, 0x31, 0x47, 0x6f, 0x2a,
	0x9d, 0xea, 0xe6, 0x37, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xba, 0xdf, 0x63, 0x3a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AnimalServiceClient is the client API for AnimalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnimalServiceClient interface {
	GetAnimals(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error)
}

type animalServiceClient struct {
	cc *grpc.ClientConn
}

func NewAnimalServiceClient(cc *grpc.ClientConn) AnimalServiceClient {
	return &animalServiceClient{cc}
}

func (c *animalServiceClient) GetAnimals(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/animal.AnimalService/GetAnimals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnimalServiceServer is the server API for AnimalService service.
type AnimalServiceServer interface {
	GetAnimals(context.Context, *GetRequest) (*Response, error)
}

// UnimplementedAnimalServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAnimalServiceServer struct {
}

func (*UnimplementedAnimalServiceServer) GetAnimals(ctx context.Context, req *GetRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnimals not implemented")
}

func RegisterAnimalServiceServer(s *grpc.Server, srv AnimalServiceServer) {
	s.RegisterService(&_AnimalService_serviceDesc, srv)
}

func _AnimalService_GetAnimals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimalServiceServer).GetAnimals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/animal.AnimalService/GetAnimals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimalServiceServer).GetAnimals(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AnimalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "animal.AnimalService",
	HandlerType: (*AnimalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAnimals",
			Handler:    _AnimalService_GetAnimals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "animal.proto",
}
