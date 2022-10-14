// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/validator-preference/v1beta1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Params defines the parameters for the module.
type Params struct {
	ValsetCreationFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=valset_creation_fee,json=valsetCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"valset_creation_fee" yaml:"valset_creation_fee"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_db06f71db3b2b0f5, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetValsetCreationFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ValsetCreationFee
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "osmosis.validatorpreference.v1beta1.Params")
}

func init() {
	proto.RegisterFile("osmosis/validator-preference/v1beta1/params.proto", fileDescriptor_db06f71db3b2b0f5)
}

var fileDescriptor_db06f71db3b2b0f5 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0x21, 0x75, 0x28, 0x13, 0x85, 0x81, 0x66, 0x70, 0x50, 0x58, 0xba, 0xc4, 0x56,
	0xca, 0xc6, 0x84, 0x5a, 0x89, 0x11, 0x21, 0xc6, 0x2e, 0x91, 0x93, 0x5e, 0x83, 0x45, 0x12, 0x47,
	0xb6, 0x89, 0xe8, 0x5b, 0xb0, 0xb2, 0x33, 0xf1, 0x24, 0x1d, 0x3b, 0x32, 0x15, 0x94, 0xbc, 0x01,
	0x4f, 0x80, 0x6a, 0x9b, 0x96, 0xa1, 0x53, 0x72, 0x3a, 0x7f, 0xdf, 0xfd, 0xba, 0xeb, 0xc7, 0x42,
	0x95, 0x42, 0x71, 0x45, 0x1b, 0x56, 0xf0, 0x39, 0xd3, 0x42, 0x46, 0xb5, 0x84, 0x05, 0x48, 0xa8,
	0x32, 0xa0, 0x4d, 0x9c, 0x82, 0x66, 0x31, 0xad, 0x99, 0x64, 0xa5, 0x22, 0xb5, 0x14, 0x5a, 0x0c,
	0x2e, 0x1d, 0x42, 0x76, 0xc8, 0x9e, 0x20, 0x8e, 0xf0, 0xcf, 0x72, 0x91, 0x0b, 0xf3, 0x9e, 0x6e,
	0xff, 0x2c, 0xea, 0x0f, 0x33, 0xc3, 0x26, 0xb6, 0x61, 0x0b, 0xd7, 0xc2, 0xb6, 0xa2, 0x29, 0x53,
	0xfb, 0xb9, 0x99, 0xe0, 0x95, 0xed, 0x87, 0xef, 0xa8, 0xdf, 0xbb, 0x37, 0x31, 0x06, 0x6f, 0xa8,
	0x7f, 0xda, 0xb0, 0x42, 0x81, 0x4e, 0x32, 0x09, 0x4c, 0x73, 0x51, 0x25, 0x0b, 0x80, 0x73, 0x74,
	0x71, 0x34, 0x3a, 0x1e, 0x0f, 0x89, 0xf3, 0x6e, 0x4d, 0x7f, 0x79, 0xc8, 0x54, 0xf0, 0x6a, 0x72,
	0xb7, 0xda, 0x04, 0xde, 0xcf, 0x26, 0xf0, 0x97, 0xac, 0x2c, 0xae, 0xc3, 0x03, 0x8e, 0xf0, 0xe3,
	0x2b, 0x18, 0xe5, 0x5c, 0x3f, 0x3e, 0xa7, 0x24, 0x13, 0xa5, 0x8b, 0xe8, 0x3e, 0x91, 0x9a, 0x3f,
	0x51, 0xbd, 0xac, 0x41, 0x19, 0x9d, 0x7a, 0x38, 0xb1, 0x86, 0xa9, 0x13, 0xdc, 0x02, 0x4c, 0x66,
	0xab, 0x16, 0xa3, 0x75, 0x8b, 0xd1, 0x77, 0x8b, 0xd1, 0x6b, 0x87, 0xbd, 0x75, 0x87, 0xbd, 0xcf,
	0x0e, 0x7b, 0xb3, 0x9b, 0x7f, 0x5a, 0xb7, 0xc1, 0xa8, 0x60, 0xa9, 0xa2, 0xbb, 0x0b, 0xc4, 0x63,
	0xfa, 0x72, 0xf8, 0x0e, 0x66, 0x68, 0xda, 0x33, 0x9b, 0xb8, 0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x55, 0x97, 0x65, 0x82, 0xb4, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValsetCreationFee) > 0 {
		for iNdEx := len(m.ValsetCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValsetCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ValsetCreationFee) > 0 {
		for _, e := range m.ValsetCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValsetCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValsetCreationFee = append(m.ValsetCreationFee, types.Coin{})
			if err := m.ValsetCreationFee[len(m.ValsetCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)