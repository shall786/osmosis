// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/valset-pref/v1beta1/query.proto

package queryproto

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/osmosis-labs/osmosis/v12/x/valset-pref/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request type for UserValidatorPreferences.
type QueryUserValidatorPreferences struct {
	// user account address
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (m *QueryUserValidatorPreferences) Reset()         { *m = QueryUserValidatorPreferences{} }
func (m *QueryUserValidatorPreferences) String() string { return proto.CompactTextString(m) }
func (*QueryUserValidatorPreferences) ProtoMessage()    {}
func (*QueryUserValidatorPreferences) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ffbeb4123fe56ae, []int{0}
}
func (m *QueryUserValidatorPreferences) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUserValidatorPreferences) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUserValidatorPreferences.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUserValidatorPreferences) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserValidatorPreferences.Merge(m, src)
}
func (m *QueryUserValidatorPreferences) XXX_Size() int {
	return m.Size()
}
func (m *QueryUserValidatorPreferences) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserValidatorPreferences.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserValidatorPreferences proto.InternalMessageInfo

// Response type the QueryUserValidatorPreferences query request
type QueryUserValidatorPreferenceResponse struct {
	Preferences []types.ValidatorPreference `protobuf:"bytes,1,rep,name=preferences,proto3" json:"preferences"`
}

func (m *QueryUserValidatorPreferenceResponse) Reset()         { *m = QueryUserValidatorPreferenceResponse{} }
func (m *QueryUserValidatorPreferenceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUserValidatorPreferenceResponse) ProtoMessage()    {}
func (*QueryUserValidatorPreferenceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ffbeb4123fe56ae, []int{1}
}
func (m *QueryUserValidatorPreferenceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUserValidatorPreferenceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUserValidatorPreferenceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUserValidatorPreferenceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserValidatorPreferenceResponse.Merge(m, src)
}
func (m *QueryUserValidatorPreferenceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryUserValidatorPreferenceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserValidatorPreferenceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserValidatorPreferenceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryUserValidatorPreferences)(nil), "osmosis.validatorpreference.v1beta1.QueryUserValidatorPreferences")
	proto.RegisterType((*QueryUserValidatorPreferenceResponse)(nil), "osmosis.validatorpreference.v1beta1.QueryUserValidatorPreferenceResponse")
}

func init() {
	proto.RegisterFile("osmosis/valset-pref/v1beta1/query.proto", fileDescriptor_9ffbeb4123fe56ae)
}

var fileDescriptor_9ffbeb4123fe56ae = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcf, 0x2f, 0xce, 0xcd,
	0x2f, 0xce, 0x2c, 0xd6, 0x2f, 0x4b, 0xcc, 0x29, 0x4e, 0x2d, 0xd1, 0x2d, 0x28, 0x4a, 0x4d, 0xd3,
	0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x52, 0x86, 0x2a, 0xd4, 0x2b, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9,
	0x2f, 0x02, 0x29, 0x4d, 0x2d, 0x4a, 0xcd, 0x4b, 0x4e, 0xd5, 0x83, 0x6a, 0x90, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x07, 0xab, 0xd7, 0x07, 0xb1, 0x20, 0x5a, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73,
	0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3,
	0x8a, 0xa1, 0xb2, 0x78, 0x5d, 0x50, 0x5c, 0x92, 0x58, 0x92, 0x0a, 0x51, 0xa8, 0x64, 0xcc, 0x25,
	0x1b, 0x08, 0x72, 0x50, 0x68, 0x71, 0x6a, 0x51, 0x18, 0xcc, 0x11, 0x01, 0x70, 0x47, 0x14, 0x0b,
	0x09, 0x71, 0xb1, 0x94, 0x16, 0xa7, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9,
	0x4a, 0x1d, 0x8c, 0x5c, 0x2a, 0xf8, 0x74, 0x05, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a,
	0x25, 0x70, 0x71, 0x23, 0x3c, 0x54, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0x64, 0xa1, 0x47,
	0x84, 0xaf, 0xf5, 0xb0, 0x18, 0xeb, 0xc4, 0x72, 0xe2, 0x9e, 0x3c, 0x43, 0x10, 0xb2, 0x91, 0x46,
	0x0f, 0x18, 0xb9, 0x58, 0xc1, 0x4e, 0x11, 0xba, 0xc6, 0xc8, 0x25, 0x81, 0xd3, 0x17, 0x4e, 0x44,
	0xd9, 0x89, 0x37, 0x24, 0xa4, 0x3c, 0x29, 0x36, 0x03, 0x16, 0x2e, 0x4a, 0xda, 0x4d, 0x97, 0x9f,
	0x4c, 0x66, 0x52, 0x15, 0x52, 0xd6, 0xc7, 0x17, 0x4f, 0xd5, 0xa0, 0xc0, 0xae, 0x75, 0x4a, 0x3c,
	0xf1, 0x50, 0x8e, 0xe1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63,
	0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x9c, 0xd3,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0x61, 0x86, 0xe9, 0xe6, 0x24, 0x26, 0x15,
	0x23, 0x4c, 0x36, 0x34, 0xd2, 0xaf, 0x40, 0x31, 0x3f, 0x39, 0x27, 0x33, 0x35, 0xaf, 0x04, 0x92,
	0x10, 0xc1, 0xa9, 0x20, 0x89, 0x0d, 0x4c, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x88, 0x49,
	0x0d, 0xe4, 0xb9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Returns the list of ValidatorPreferences for the user.
	UserValidatorPreferences(ctx context.Context, in *QueryUserValidatorPreferences, opts ...grpc.CallOption) (*QueryUserValidatorPreferenceResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) UserValidatorPreferences(ctx context.Context, in *QueryUserValidatorPreferences, opts ...grpc.CallOption) (*QueryUserValidatorPreferenceResponse, error) {
	out := new(QueryUserValidatorPreferenceResponse)
	err := c.cc.Invoke(ctx, "/osmosis.validatorpreference.v1beta1.Query/UserValidatorPreferences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Returns the list of ValidatorPreferences for the user.
	UserValidatorPreferences(context.Context, *QueryUserValidatorPreferences) (*QueryUserValidatorPreferenceResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) UserValidatorPreferences(ctx context.Context, req *QueryUserValidatorPreferences) (*QueryUserValidatorPreferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserValidatorPreferences not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_UserValidatorPreferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUserValidatorPreferences)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UserValidatorPreferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.validatorpreference.v1beta1.Query/UserValidatorPreferences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UserValidatorPreferences(ctx, req.(*QueryUserValidatorPreferences))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.validatorpreference.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserValidatorPreferences",
			Handler:    _Query_UserValidatorPreferences_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/valset-pref/v1beta1/query.proto",
}

func (m *QueryUserValidatorPreferences) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUserValidatorPreferences) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUserValidatorPreferences) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryUserValidatorPreferenceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUserValidatorPreferenceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUserValidatorPreferenceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Preferences) > 0 {
		for iNdEx := len(m.Preferences) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Preferences[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryUserValidatorPreferences) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryUserValidatorPreferenceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Preferences) > 0 {
		for _, e := range m.Preferences {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryUserValidatorPreferences) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryUserValidatorPreferences: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUserValidatorPreferences: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryUserValidatorPreferenceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryUserValidatorPreferenceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUserValidatorPreferenceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Preferences", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Preferences = append(m.Preferences, types.ValidatorPreference{})
			if err := m.Preferences[len(m.Preferences)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)