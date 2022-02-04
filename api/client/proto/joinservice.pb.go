// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joinservice.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gravitational/teleport/api/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// RegisterUsingIAMMethodResponse is a stream response and will contain either a
// Challenge or signed Certs to join the cluster.
type RegisterUsingIAMMethodResponse struct {
	// Challenge is a crypto-random string that should be included in the signed
	// sts:GetCallerIdentity request.
	Challenge string `protobuf:"bytes,1,opt,name=Challenge,proto3" json:"Challenge,omitempty"`
	// Certs is the returned signed certs.
	Certs                *Certs   `protobuf:"bytes,2,opt,name=Certs,proto3" json:"Certs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterUsingIAMMethodResponse) Reset()         { *m = RegisterUsingIAMMethodResponse{} }
func (m *RegisterUsingIAMMethodResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterUsingIAMMethodResponse) ProtoMessage()    {}
func (*RegisterUsingIAMMethodResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eef71e659fffadc3, []int{0}
}
func (m *RegisterUsingIAMMethodResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterUsingIAMMethodResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterUsingIAMMethodResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterUsingIAMMethodResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterUsingIAMMethodResponse.Merge(m, src)
}
func (m *RegisterUsingIAMMethodResponse) XXX_Size() int {
	return m.Size()
}
func (m *RegisterUsingIAMMethodResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterUsingIAMMethodResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterUsingIAMMethodResponse proto.InternalMessageInfo

func (m *RegisterUsingIAMMethodResponse) GetChallenge() string {
	if m != nil {
		return m.Challenge
	}
	return ""
}

func (m *RegisterUsingIAMMethodResponse) GetCerts() *Certs {
	if m != nil {
		return m.Certs
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterUsingIAMMethodResponse)(nil), "proto.RegisterUsingIAMMethodResponse")
}

func init() { proto.RegisterFile("joinservice.proto", fileDescriptor_eef71e659fffadc3) }

var fileDescriptor_eef71e659fffadc3 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x89, 0x50, 0xa1, 0x69, 0x2f, 0xe6, 0x20, 0x65, 0x91, 0x65, 0x29, 0x08, 0x7b, 0xda,
	0x48, 0x3d, 0x78, 0xd6, 0x9e, 0x14, 0x7a, 0x89, 0xfa, 0x00, 0xd9, 0x3a, 0x64, 0x47, 0xd7, 0x4c,
	0x4c, 0xa6, 0x05, 0xdf, 0xd0, 0xa3, 0x8f, 0x20, 0xfb, 0x24, 0xc2, 0x46, 0x50, 0x41, 0x3c, 0x05,
	0xbe, 0x9f, 0xef, 0xcf, 0x3f, 0xf2, 0xe8, 0x91, 0xd0, 0x27, 0x88, 0x7b, 0xdc, 0x42, 0x13, 0x22,
	0x31, 0xa9, 0xc9, 0xf8, 0x14, 0x33, 0x7e, 0x0d, 0x90, 0x32, 0x2b, 0x2e, 0x1c, 0x72, 0xb7, 0x6b,
	0x9b, 0x2d, 0x3d, 0x6b, 0x17, 0xed, 0x1e, 0xd9, 0x32, 0x92, 0xb7, 0xbd, 0x66, 0xe8, 0x21, 0x50,
	0x64, 0x6d, 0x03, 0xea, 0x51, 0xd1, 0x3f, 0xc4, 0x65, 0x2b, 0x4b, 0x03, 0x0e, 0x13, 0x43, 0xbc,
	0x4f, 0xe8, 0xdd, 0xf5, 0xe5, 0x66, 0x03, 0xdc, 0xd1, 0x83, 0x81, 0x14, 0xc8, 0x27, 0x50, 0x27,
	0x72, 0xba, 0xee, 0x6c, 0xdf, 0x83, 0x77, 0xb0, 0x10, 0x95, 0xa8, 0xa7, 0xe6, 0x1b, 0xa8, 0xa5,
	0x9c, 0xac, 0x21, 0x72, 0x5a, 0x1c, 0x54, 0xa2, 0x9e, 0xad, 0xe6, 0xb9, 0xb6, 0x19, 0x99, 0xc9,
	0xd1, 0x8a, 0xe5, 0xec, 0x86, 0xd0, 0xdf, 0xe6, 0x2b, 0x14, 0xc8, 0xe3, 0xbf, 0xbf, 0x54, 0x55,
	0x93, 0xa7, 0xfd, 0x8a, 0xef, 0xe8, 0x09, 0xbc, 0x81, 0x97, 0x1d, 0x24, 0x2e, 0x4e, 0xbf, 0xfa,
	0xff, 0xdf, 0x5c, 0x8b, 0x33, 0x71, 0x35, 0x7f, 0x1b, 0x4a, 0xf1, 0x3e, 0x94, 0xe2, 0x63, 0x28,
	0x45, 0x7b, 0x38, 0x7a, 0xe7, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x61, 0xfd, 0x4d, 0x50,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JoinServiceClient is the client API for JoinService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JoinServiceClient interface {
	// RegisterUsingIAMMethod is used to register a new node to the cluster using
	// the IAM join method.
	RegisterUsingIAMMethod(ctx context.Context, opts ...grpc.CallOption) (JoinService_RegisterUsingIAMMethodClient, error)
}

type joinServiceClient struct {
	cc *grpc.ClientConn
}

func NewJoinServiceClient(cc *grpc.ClientConn) JoinServiceClient {
	return &joinServiceClient{cc}
}

func (c *joinServiceClient) RegisterUsingIAMMethod(ctx context.Context, opts ...grpc.CallOption) (JoinService_RegisterUsingIAMMethodClient, error) {
	stream, err := c.cc.NewStream(ctx, &_JoinService_serviceDesc.Streams[0], "/proto.JoinService/RegisterUsingIAMMethod", opts...)
	if err != nil {
		return nil, err
	}
	x := &joinServiceRegisterUsingIAMMethodClient{stream}
	return x, nil
}

type JoinService_RegisterUsingIAMMethodClient interface {
	Send(*types.RegisterUsingTokenRequest) error
	Recv() (*RegisterUsingIAMMethodResponse, error)
	grpc.ClientStream
}

type joinServiceRegisterUsingIAMMethodClient struct {
	grpc.ClientStream
}

func (x *joinServiceRegisterUsingIAMMethodClient) Send(m *types.RegisterUsingTokenRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *joinServiceRegisterUsingIAMMethodClient) Recv() (*RegisterUsingIAMMethodResponse, error) {
	m := new(RegisterUsingIAMMethodResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// JoinServiceServer is the server API for JoinService service.
type JoinServiceServer interface {
	// RegisterUsingIAMMethod is used to register a new node to the cluster using
	// the IAM join method.
	RegisterUsingIAMMethod(JoinService_RegisterUsingIAMMethodServer) error
}

// UnimplementedJoinServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJoinServiceServer struct {
}

func (*UnimplementedJoinServiceServer) RegisterUsingIAMMethod(srv JoinService_RegisterUsingIAMMethodServer) error {
	return status.Errorf(codes.Unimplemented, "method RegisterUsingIAMMethod not implemented")
}

func RegisterJoinServiceServer(s *grpc.Server, srv JoinServiceServer) {
	s.RegisterService(&_JoinService_serviceDesc, srv)
}

func _JoinService_RegisterUsingIAMMethod_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(JoinServiceServer).RegisterUsingIAMMethod(&joinServiceRegisterUsingIAMMethodServer{stream})
}

type JoinService_RegisterUsingIAMMethodServer interface {
	Send(*RegisterUsingIAMMethodResponse) error
	Recv() (*types.RegisterUsingTokenRequest, error)
	grpc.ServerStream
}

type joinServiceRegisterUsingIAMMethodServer struct {
	grpc.ServerStream
}

func (x *joinServiceRegisterUsingIAMMethodServer) Send(m *RegisterUsingIAMMethodResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *joinServiceRegisterUsingIAMMethodServer) Recv() (*types.RegisterUsingTokenRequest, error) {
	m := new(types.RegisterUsingTokenRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _JoinService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.JoinService",
	HandlerType: (*JoinServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterUsingIAMMethod",
			Handler:       _JoinService_RegisterUsingIAMMethod_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "joinservice.proto",
}

func (m *RegisterUsingIAMMethodResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterUsingIAMMethodResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterUsingIAMMethodResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Certs != nil {
		{
			size, err := m.Certs.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintJoinservice(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Challenge) > 0 {
		i -= len(m.Challenge)
		copy(dAtA[i:], m.Challenge)
		i = encodeVarintJoinservice(dAtA, i, uint64(len(m.Challenge)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintJoinservice(dAtA []byte, offset int, v uint64) int {
	offset -= sovJoinservice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RegisterUsingIAMMethodResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Challenge)
	if l > 0 {
		n += 1 + l + sovJoinservice(uint64(l))
	}
	if m.Certs != nil {
		l = m.Certs.Size()
		n += 1 + l + sovJoinservice(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovJoinservice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozJoinservice(x uint64) (n int) {
	return sovJoinservice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisterUsingIAMMethodResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJoinservice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterUsingIAMMethodResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterUsingIAMMethodResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Challenge", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoinservice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJoinservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJoinservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Challenge = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoinservice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoinservice
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJoinservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Certs == nil {
				m.Certs = &Certs{}
			}
			if err := m.Certs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJoinservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthJoinservice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipJoinservice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJoinservice
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowJoinservice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowJoinservice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthJoinservice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupJoinservice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthJoinservice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthJoinservice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJoinservice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupJoinservice = fmt.Errorf("proto: unexpected end of group")
)