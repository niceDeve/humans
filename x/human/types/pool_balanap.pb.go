// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: human/pool_balanap.proto

package types

import (
	fmt "fmt"
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

type PoolBalanap struct {
	Index              string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	TransactionData    string `protobuf:"bytes,2,opt,name=transactionData,proto3" json:"transactionData,omitempty"`
	OriginChain        string `protobuf:"bytes,3,opt,name=originChain,proto3" json:"originChain,omitempty"`
	OriginAddress      string `protobuf:"bytes,4,opt,name=originAddress,proto3" json:"originAddress,omitempty"`
	TargetChain        string `protobuf:"bytes,5,opt,name=targetChain,proto3" json:"targetChain,omitempty"`
	TargetAddress      string `protobuf:"bytes,6,opt,name=targetAddress,proto3" json:"targetAddress,omitempty"`
	Amount             string `protobuf:"bytes,7,opt,name=amount,proto3" json:"amount,omitempty"`
	Time               string `protobuf:"bytes,8,opt,name=time,proto3" json:"time,omitempty"`
	Creator            string `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator,omitempty"`
	Status             string `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	ConfirmedBlockHash string `protobuf:"bytes,11,opt,name=confirmedBlockHash,proto3" json:"confirmedBlockHash,omitempty"`
	SignedKey          string `protobuf:"bytes,12,opt,name=signedKey,proto3" json:"signedKey,omitempty"`
	Fee                string `protobuf:"bytes,13,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *PoolBalanap) Reset()         { *m = PoolBalanap{} }
func (m *PoolBalanap) String() string { return proto.CompactTextString(m) }
func (*PoolBalanap) ProtoMessage()    {}
func (*PoolBalanap) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eda154650c31af7, []int{0}
}
func (m *PoolBalanap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolBalanap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolBalanap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolBalanap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolBalanap.Merge(m, src)
}
func (m *PoolBalanap) XXX_Size() int {
	return m.Size()
}
func (m *PoolBalanap) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolBalanap.DiscardUnknown(m)
}

var xxx_messageInfo_PoolBalanap proto.InternalMessageInfo

func (m *PoolBalanap) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *PoolBalanap) GetTransactionData() string {
	if m != nil {
		return m.TransactionData
	}
	return ""
}

func (m *PoolBalanap) GetOriginChain() string {
	if m != nil {
		return m.OriginChain
	}
	return ""
}

func (m *PoolBalanap) GetOriginAddress() string {
	if m != nil {
		return m.OriginAddress
	}
	return ""
}

func (m *PoolBalanap) GetTargetChain() string {
	if m != nil {
		return m.TargetChain
	}
	return ""
}

func (m *PoolBalanap) GetTargetAddress() string {
	if m != nil {
		return m.TargetAddress
	}
	return ""
}

func (m *PoolBalanap) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *PoolBalanap) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *PoolBalanap) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *PoolBalanap) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *PoolBalanap) GetConfirmedBlockHash() string {
	if m != nil {
		return m.ConfirmedBlockHash
	}
	return ""
}

func (m *PoolBalanap) GetSignedKey() string {
	if m != nil {
		return m.SignedKey
	}
	return ""
}

func (m *PoolBalanap) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func init() {
	proto.RegisterType((*PoolBalanap)(nil), "human.human.PoolBalanap")
}

func init() { proto.RegisterFile("human/pool_balanap.proto", fileDescriptor_6eda154650c31af7) }

var fileDescriptor_6eda154650c31af7 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbf, 0x4e, 0x32, 0x41,
	0x14, 0xc5, 0xd9, 0x8f, 0x7f, 0x1f, 0xb3, 0x12, 0xcd, 0xd5, 0x98, 0x5b, 0x98, 0x0d, 0x31, 0x16,
	0x34, 0x42, 0xe1, 0x13, 0x88, 0x16, 0x26, 0x36, 0xc6, 0xd2, 0xc6, 0x5c, 0x76, 0x07, 0x98, 0xb8,
	0x3b, 0xb3, 0x99, 0x19, 0x12, 0x78, 0x0b, 0x7b, 0x5f, 0xc8, 0x92, 0xd2, 0xd2, 0xc0, 0x8b, 0x98,
	0xbd, 0x23, 0x11, 0x8c, 0xcd, 0xee, 0x9c, 0xdf, 0x39, 0xf9, 0x35, 0x57, 0xe0, 0x6c, 0x5e, 0x90,
	0x1e, 0x96, 0xc6, 0xe4, 0xcf, 0x63, 0xca, 0x49, 0x53, 0x39, 0x28, 0xad, 0xf1, 0x06, 0x62, 0x6e,
	0x06, 0xfc, 0x3d, 0x7f, 0xab, 0x8b, 0xf8, 0xc1, 0x98, 0x7c, 0x14, 0x26, 0x70, 0x22, 0x9a, 0x4a,
	0x67, 0x72, 0x81, 0x51, 0x2f, 0xea, 0x77, 0x1e, 0x43, 0x80, 0xbe, 0x38, 0xf4, 0x96, 0xb4, 0xa3,
	0xd4, 0x2b, 0xa3, 0x6f, 0xc9, 0x13, 0xfe, 0xe3, 0xfe, 0x37, 0x86, 0x9e, 0x88, 0x8d, 0x55, 0x53,
	0xa5, 0x6f, 0x66, 0xa4, 0x34, 0xd6, 0x79, 0xb5, 0x8b, 0xe0, 0x42, 0x74, 0x43, 0xbc, 0xce, 0x32,
	0x2b, 0x9d, 0xc3, 0x06, 0x6f, 0xf6, 0x61, 0xe5, 0xf1, 0x64, 0xa7, 0xd2, 0x07, 0x4f, 0x33, 0x78,
	0x76, 0x50, 0xe5, 0x09, 0x71, 0xeb, 0x69, 0x05, 0xcf, 0x1e, 0x84, 0x53, 0xd1, 0xa2, 0xc2, 0xcc,
	0xb5, 0xc7, 0x36, 0xd7, 0xdf, 0x09, 0x40, 0x34, 0xbc, 0x2a, 0x24, 0xfe, 0x67, 0xca, 0x6f, 0x40,
	0xd1, 0x4e, 0xad, 0x24, 0x6f, 0x2c, 0x76, 0x18, 0x6f, 0x63, 0x65, 0x71, 0x9e, 0xfc, 0xdc, 0xa1,
	0x08, 0x96, 0x90, 0x60, 0x20, 0x20, 0x35, 0x7a, 0xa2, 0x6c, 0x21, 0xb3, 0x51, 0x6e, 0xd2, 0x97,
	0x3b, 0x72, 0x33, 0x8c, 0x79, 0xf3, 0x47, 0x03, 0x67, 0xa2, 0xe3, 0xd4, 0x54, 0xcb, 0xec, 0x5e,
	0x2e, 0xf1, 0x80, 0x67, 0x3f, 0x00, 0x8e, 0x44, 0x7d, 0x22, 0x25, 0x76, 0x99, 0x57, 0xcf, 0xd1,
	0xe5, 0xfb, 0x3a, 0x89, 0x56, 0xeb, 0x24, 0xfa, 0x5c, 0x27, 0xd1, 0xeb, 0x26, 0xa9, 0xad, 0x36,
	0x49, 0xed, 0x63, 0x93, 0xd4, 0x9e, 0x8e, 0xc3, 0x79, 0x17, 0xc3, 0xf0, 0xf7, 0xcb, 0x52, 0xba,
	0x71, 0x8b, 0x0f, 0x7c, 0xf5, 0x15, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x0f, 0xcc, 0x46, 0xfc, 0x01,
	0x00, 0x00,
}

func (m *PoolBalanap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolBalanap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolBalanap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fee) > 0 {
		i -= len(m.Fee)
		copy(dAtA[i:], m.Fee)
		i = encodeVarintPoolBalanap(dAtA, i, uint64(len(m.Fee)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.SignedKey) > 0 {
		i -= len(m.SignedKey)
		copy(dAtA[i:], m.SignedKey)
		i = encodeVarintPoolBalanap(dAtA, i, uint64(len(m.SignedKey)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.ConfirmedBlockHash) > 0 {
		i -= len(m.ConfirmedBlockHash)
		copy(dAtA[i:], m.ConfirmedBlockHash)
		i = encodeVarintPoolBalanap(dAtA, i, uint64(len(m.ConfirmedBlockHash)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintPoolBalanap(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPoolBalanap(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Time) > 0 {
		i -= len(m.Time)
		copy(dAtA[i:], m.Time)
		i = encodeVarintPoolBalanap(dAtA, i, uint64(len(m.Time)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintPoolBalanap(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.TargetAddress) > 0 {
		i -= len(m.TargetAddress)
		copy(dAtA[i:], m.TargetAddress)
		i = encodeVarintPoolBalanap(dAtA, i, uint64(len(m.TargetAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.TargetChain) > 0 {
		i -= len(m.TargetChain)
		copy(dAtA[i:], m.TargetChain)
		i = encodeVarintPoolBalanap(dAtA, i, uint64(len(m.TargetChain)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.OriginAddress) > 0 {
		i -= len(m.OriginAddress)
		copy(dAtA[i:], m.OriginAddress)
		i = encodeVarintPoolBalanap(dAtA, i, uint64(len(m.OriginAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OriginChain) > 0 {
		i -= len(m.OriginChain)
		copy(dAtA[i:], m.OriginChain)
		i = encodeVarintPoolBalanap(dAtA, i, uint64(len(m.OriginChain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TransactionData) > 0 {
		i -= len(m.TransactionData)
		copy(dAtA[i:], m.TransactionData)
		i = encodeVarintPoolBalanap(dAtA, i, uint64(len(m.TransactionData)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintPoolBalanap(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPoolBalanap(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoolBalanap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolBalanap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovPoolBalanap(uint64(l))
	}
	l = len(m.TransactionData)
	if l > 0 {
		n += 1 + l + sovPoolBalanap(uint64(l))
	}
	l = len(m.OriginChain)
	if l > 0 {
		n += 1 + l + sovPoolBalanap(uint64(l))
	}
	l = len(m.OriginAddress)
	if l > 0 {
		n += 1 + l + sovPoolBalanap(uint64(l))
	}
	l = len(m.TargetChain)
	if l > 0 {
		n += 1 + l + sovPoolBalanap(uint64(l))
	}
	l = len(m.TargetAddress)
	if l > 0 {
		n += 1 + l + sovPoolBalanap(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovPoolBalanap(uint64(l))
	}
	l = len(m.Time)
	if l > 0 {
		n += 1 + l + sovPoolBalanap(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPoolBalanap(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovPoolBalanap(uint64(l))
	}
	l = len(m.ConfirmedBlockHash)
	if l > 0 {
		n += 1 + l + sovPoolBalanap(uint64(l))
	}
	l = len(m.SignedKey)
	if l > 0 {
		n += 1 + l + sovPoolBalanap(uint64(l))
	}
	l = len(m.Fee)
	if l > 0 {
		n += 1 + l + sovPoolBalanap(uint64(l))
	}
	return n
}

func sovPoolBalanap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoolBalanap(x uint64) (n int) {
	return sovPoolBalanap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolBalanap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoolBalanap
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
			return fmt.Errorf("proto: PoolBalanap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolBalanap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolBalanap
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
				return ErrInvalidLengthPoolBalanap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolBalanap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolBalanap
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
				return ErrInvalidLengthPoolBalanap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolBalanap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransactionData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolBalanap
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
				return ErrInvalidLengthPoolBalanap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolBalanap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolBalanap
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
				return ErrInvalidLengthPoolBalanap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolBalanap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolBalanap
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
				return ErrInvalidLengthPoolBalanap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolBalanap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolBalanap
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
				return ErrInvalidLengthPoolBalanap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolBalanap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolBalanap
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
				return ErrInvalidLengthPoolBalanap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolBalanap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolBalanap
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
				return ErrInvalidLengthPoolBalanap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolBalanap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolBalanap
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
				return ErrInvalidLengthPoolBalanap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolBalanap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolBalanap
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
				return ErrInvalidLengthPoolBalanap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolBalanap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfirmedBlockHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolBalanap
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
				return ErrInvalidLengthPoolBalanap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolBalanap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConfirmedBlockHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolBalanap
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
				return ErrInvalidLengthPoolBalanap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolBalanap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignedKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolBalanap
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
				return ErrInvalidLengthPoolBalanap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolBalanap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPoolBalanap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoolBalanap
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
func skipPoolBalanap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoolBalanap
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
					return 0, ErrIntOverflowPoolBalanap
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
					return 0, ErrIntOverflowPoolBalanap
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
				return 0, ErrInvalidLengthPoolBalanap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoolBalanap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoolBalanap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoolBalanap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoolBalanap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoolBalanap = fmt.Errorf("proto: unexpected end of group")
)
