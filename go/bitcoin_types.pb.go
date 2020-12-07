// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitcoin_types.proto

package messages

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// *
// Structure representing Bitcoin transaction input
type BitcoinTransactionInput struct {
	AddressN *uint32 `protobuf:"varint,1,req,name=address_n,json=addressN" json:"address_n,omitempty"`
	PrevHash []byte  `protobuf:"bytes,2,req,name=prev_hash,json=prevHash" json:"prev_hash,omitempty"`
	// required uint32 index = 3;		// Index of output in previous transaction, which will be spended
	Value                *uint64  `protobuf:"varint,4,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BitcoinTransactionInput) Reset()         { *m = BitcoinTransactionInput{} }
func (m *BitcoinTransactionInput) String() string { return proto.CompactTextString(m) }
func (*BitcoinTransactionInput) ProtoMessage()    {}
func (*BitcoinTransactionInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_bitcoin_types_6b068667b421b045, []int{0}
}
func (m *BitcoinTransactionInput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BitcoinTransactionInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BitcoinTransactionInput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BitcoinTransactionInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BitcoinTransactionInput.Merge(dst, src)
}
func (m *BitcoinTransactionInput) XXX_Size() int {
	return m.Size()
}
func (m *BitcoinTransactionInput) XXX_DiscardUnknown() {
	xxx_messageInfo_BitcoinTransactionInput.DiscardUnknown(m)
}

var xxx_messageInfo_BitcoinTransactionInput proto.InternalMessageInfo

func (m *BitcoinTransactionInput) GetAddressN() uint32 {
	if m != nil && m.AddressN != nil {
		return *m.AddressN
	}
	return 0
}

func (m *BitcoinTransactionInput) GetPrevHash() []byte {
	if m != nil {
		return m.PrevHash
	}
	return nil
}

func (m *BitcoinTransactionInput) GetValue() uint64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

// *
// Structure representing transaction output
type BitcoinTransactionOutput struct {
	Address              *string  `protobuf:"bytes,1,req,name=address" json:"address,omitempty"`
	Coin                 *uint64  `protobuf:"varint,2,req,name=coin" json:"coin,omitempty"`
	AddressIndex         *uint32  `protobuf:"varint,3,opt,name=address_index,json=addressIndex" json:"address_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BitcoinTransactionOutput) Reset()         { *m = BitcoinTransactionOutput{} }
func (m *BitcoinTransactionOutput) String() string { return proto.CompactTextString(m) }
func (*BitcoinTransactionOutput) ProtoMessage()    {}
func (*BitcoinTransactionOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_bitcoin_types_6b068667b421b045, []int{1}
}
func (m *BitcoinTransactionOutput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BitcoinTransactionOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BitcoinTransactionOutput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BitcoinTransactionOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BitcoinTransactionOutput.Merge(dst, src)
}
func (m *BitcoinTransactionOutput) XXX_Size() int {
	return m.Size()
}
func (m *BitcoinTransactionOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_BitcoinTransactionOutput.DiscardUnknown(m)
}

var xxx_messageInfo_BitcoinTransactionOutput proto.InternalMessageInfo

func (m *BitcoinTransactionOutput) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *BitcoinTransactionOutput) GetCoin() uint64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *BitcoinTransactionOutput) GetAddressIndex() uint32 {
	if m != nil && m.AddressIndex != nil {
		return *m.AddressIndex
	}
	return 0
}

// *
// Structure representing transaction
type BitcoinTransactionType struct {
	InputsCnt            *uint32                     `protobuf:"varint,2,opt,name=inputs_cnt,json=inputsCnt" json:"inputs_cnt,omitempty"`
	Inputs               []*BitcoinTransactionInput  `protobuf:"bytes,3,rep,name=inputs" json:"inputs,omitempty"`
	OutputsCnt           *uint32                     `protobuf:"varint,4,opt,name=outputs_cnt,json=outputsCnt" json:"outputs_cnt,omitempty"`
	Outputs              []*BitcoinTransactionOutput `protobuf:"bytes,5,rep,name=outputs" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *BitcoinTransactionType) Reset()         { *m = BitcoinTransactionType{} }
func (m *BitcoinTransactionType) String() string { return proto.CompactTextString(m) }
func (*BitcoinTransactionType) ProtoMessage()    {}
func (*BitcoinTransactionType) Descriptor() ([]byte, []int) {
	return fileDescriptor_bitcoin_types_6b068667b421b045, []int{2}
}
func (m *BitcoinTransactionType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BitcoinTransactionType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BitcoinTransactionType.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BitcoinTransactionType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BitcoinTransactionType.Merge(dst, src)
}
func (m *BitcoinTransactionType) XXX_Size() int {
	return m.Size()
}
func (m *BitcoinTransactionType) XXX_DiscardUnknown() {
	xxx_messageInfo_BitcoinTransactionType.DiscardUnknown(m)
}

var xxx_messageInfo_BitcoinTransactionType proto.InternalMessageInfo

func (m *BitcoinTransactionType) GetInputsCnt() uint32 {
	if m != nil && m.InputsCnt != nil {
		return *m.InputsCnt
	}
	return 0
}

func (m *BitcoinTransactionType) GetInputs() []*BitcoinTransactionInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *BitcoinTransactionType) GetOutputsCnt() uint32 {
	if m != nil && m.OutputsCnt != nil {
		return *m.OutputsCnt
	}
	return 0
}

func (m *BitcoinTransactionType) GetOutputs() []*BitcoinTransactionOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func init() {
	proto.RegisterType((*BitcoinTransactionInput)(nil), "BitcoinTransactionInput")
	proto.RegisterType((*BitcoinTransactionOutput)(nil), "BitcoinTransactionOutput")
	proto.RegisterType((*BitcoinTransactionType)(nil), "BitcoinTransactionType")
}
func (m *BitcoinTransactionInput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BitcoinTransactionInput) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AddressN == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("address_n")
	} else {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBitcoinTypes(dAtA, i, uint64(*m.AddressN))
	}
	if m.PrevHash == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("prev_hash")
	} else {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBitcoinTypes(dAtA, i, uint64(len(m.PrevHash)))
		i += copy(dAtA[i:], m.PrevHash)
	}
	if m.Value != nil {
		dAtA[i] = 0x20
		i++
		i = encodeVarintBitcoinTypes(dAtA, i, uint64(*m.Value))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BitcoinTransactionOutput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BitcoinTransactionOutput) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Address == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("address")
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBitcoinTypes(dAtA, i, uint64(len(*m.Address)))
		i += copy(dAtA[i:], *m.Address)
	}
	if m.Coin == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("coin")
	} else {
		dAtA[i] = 0x10
		i++
		i = encodeVarintBitcoinTypes(dAtA, i, uint64(*m.Coin))
	}
	if m.AddressIndex != nil {
		dAtA[i] = 0x18
		i++
		i = encodeVarintBitcoinTypes(dAtA, i, uint64(*m.AddressIndex))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BitcoinTransactionType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BitcoinTransactionType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.InputsCnt != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintBitcoinTypes(dAtA, i, uint64(*m.InputsCnt))
	}
	if len(m.Inputs) > 0 {
		for _, msg := range m.Inputs {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintBitcoinTypes(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.OutputsCnt != nil {
		dAtA[i] = 0x20
		i++
		i = encodeVarintBitcoinTypes(dAtA, i, uint64(*m.OutputsCnt))
	}
	if len(m.Outputs) > 0 {
		for _, msg := range m.Outputs {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintBitcoinTypes(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintBitcoinTypes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BitcoinTransactionInput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AddressN != nil {
		n += 1 + sovBitcoinTypes(uint64(*m.AddressN))
	}
	if m.PrevHash != nil {
		l = len(m.PrevHash)
		n += 1 + l + sovBitcoinTypes(uint64(l))
	}
	if m.Value != nil {
		n += 1 + sovBitcoinTypes(uint64(*m.Value))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BitcoinTransactionOutput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Address != nil {
		l = len(*m.Address)
		n += 1 + l + sovBitcoinTypes(uint64(l))
	}
	if m.Coin != nil {
		n += 1 + sovBitcoinTypes(uint64(*m.Coin))
	}
	if m.AddressIndex != nil {
		n += 1 + sovBitcoinTypes(uint64(*m.AddressIndex))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BitcoinTransactionType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InputsCnt != nil {
		n += 1 + sovBitcoinTypes(uint64(*m.InputsCnt))
	}
	if len(m.Inputs) > 0 {
		for _, e := range m.Inputs {
			l = e.Size()
			n += 1 + l + sovBitcoinTypes(uint64(l))
		}
	}
	if m.OutputsCnt != nil {
		n += 1 + sovBitcoinTypes(uint64(*m.OutputsCnt))
	}
	if len(m.Outputs) > 0 {
		for _, e := range m.Outputs {
			l = e.Size()
			n += 1 + l + sovBitcoinTypes(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBitcoinTypes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBitcoinTypes(x uint64) (n int) {
	return sovBitcoinTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BitcoinTransactionInput) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBitcoinTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BitcoinTransactionInput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BitcoinTransactionInput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressN", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AddressN = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBitcoinTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevHash = append(m.PrevHash[:0], dAtA[iNdEx:postIndex]...)
			if m.PrevHash == nil {
				m.PrevHash = []byte{}
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &v
		default:
			iNdEx = preIndex
			skippy, err := skipBitcoinTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBitcoinTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("address_n")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("prev_hash")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BitcoinTransactionOutput) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBitcoinTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BitcoinTransactionOutput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BitcoinTransactionOutput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBitcoinTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Address = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Coin = &v
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressIndex", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AddressIndex = &v
		default:
			iNdEx = preIndex
			skippy, err := skipBitcoinTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBitcoinTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("address")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("coin")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BitcoinTransactionType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBitcoinTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BitcoinTransactionType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BitcoinTransactionType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InputsCnt", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InputsCnt = &v
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBitcoinTypes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Inputs = append(m.Inputs, &BitcoinTransactionInput{})
			if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutputsCnt", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OutputsCnt = &v
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Outputs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBitcoinTypes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Outputs = append(m.Outputs, &BitcoinTransactionOutput{})
			if err := m.Outputs[len(m.Outputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBitcoinTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBitcoinTypes
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
func skipBitcoinTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBitcoinTypes
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
					return 0, ErrIntOverflowBitcoinTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBitcoinTypes
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthBitcoinTypes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBitcoinTypes
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipBitcoinTypes(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthBitcoinTypes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBitcoinTypes   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("bitcoin_types.proto", fileDescriptor_bitcoin_types_6b068667b421b045) }

var fileDescriptor_bitcoin_types_6b068667b421b045 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xe5, 0x34, 0xfd, 0x9a, 0xde, 0x36, 0x52, 0xe5, 0x0f, 0x81, 0x11, 0x22, 0x44, 0x65,
	0xc9, 0x14, 0x21, 0x78, 0x83, 0x22, 0x24, 0xca, 0x00, 0x92, 0xe9, 0xc4, 0x12, 0xb9, 0xad, 0xa1,
	0x11, 0xad, 0x13, 0xc5, 0x4e, 0xa1, 0x4f, 0xc7, 0xca, 0xc8, 0x23, 0xa0, 0x3c, 0x09, 0xf2, 0x9f,
	0x4c, 0x55, 0xa7, 0xe4, 0xfc, 0x7c, 0xef, 0x39, 0xbe, 0xd7, 0xf0, 0x7f, 0x9e, 0xab, 0x45, 0x91,
	0x8b, 0x4c, 0xed, 0x4a, 0x2e, 0xd3, 0xb2, 0x2a, 0x54, 0x31, 0x2e, 0xe0, 0x64, 0x62, 0xf1, 0xac,
	0x62, 0x42, 0xb2, 0x85, 0xca, 0x0b, 0x31, 0x15, 0x65, 0xad, 0xf0, 0x19, 0xf4, 0xd9, 0x72, 0x59,
	0x71, 0x29, 0x33, 0x41, 0x50, 0xec, 0x25, 0x21, 0x0d, 0x1c, 0x78, 0xd4, 0x87, 0x65, 0xc5, 0xb7,
	0xd9, 0x8a, 0xc9, 0x15, 0xf1, 0x62, 0x2f, 0x19, 0xd2, 0x40, 0x83, 0x7b, 0x26, 0x57, 0xf8, 0x08,
	0xba, 0x5b, 0xb6, 0xae, 0x39, 0xf1, 0x63, 0x94, 0xf8, 0xd4, 0x8a, 0x07, 0x3f, 0xe8, 0x8e, 0x46,
	0xe3, 0x0d, 0x90, 0xfd, 0xc0, 0xa7, 0x5a, 0xe9, 0x44, 0x02, 0x3d, 0x17, 0x60, 0xf2, 0xfa, 0xb4,
	0x95, 0x18, 0x83, 0xaf, 0x5b, 0x4c, 0x92, 0x4f, 0xcd, 0x3f, 0xbe, 0x84, 0xb0, 0xbd, 0x5f, 0x2e,
	0x96, 0xfc, 0x93, 0x74, 0x62, 0x94, 0x84, 0x74, 0xe8, 0xe0, 0x54, 0xb3, 0xf1, 0x17, 0x82, 0xe3,
	0xfd, 0xbc, 0xd9, 0xae, 0xe4, 0xf8, 0x1c, 0x20, 0xd7, 0x83, 0xca, 0x6c, 0x21, 0x14, 0xf1, 0x4c,
	0x73, 0xdf, 0x92, 0x5b, 0xa1, 0xf0, 0x15, 0xfc, 0xb3, 0x82, 0x74, 0xe2, 0x4e, 0x32, 0xb8, 0x26,
	0xe9, 0x81, 0x45, 0x51, 0x57, 0x87, 0x2f, 0x60, 0x50, 0x98, 0x41, 0xac, 0xa3, 0x6f, 0x1c, 0xc1,
	0x21, 0x6d, 0x79, 0x03, 0x3d, 0xa7, 0x48, 0xd7, 0x78, 0x9e, 0xa6, 0x87, 0x76, 0x41, 0xdb, 0xca,
	0xc9, 0xdd, 0x77, 0x13, 0xa1, 0x9f, 0x26, 0x42, 0xbf, 0x4d, 0x84, 0x20, 0x12, 0x5c, 0xa5, 0xf2,
	0x7d, 0xa7, 0x9b, 0xf4, 0xf7, 0x83, 0xad, 0xd7, 0x5c, 0xd9, 0xc7, 0x9c, 0xd7, 0xaf, 0x93, 0xf0,
	0xb9, 0x65, 0x7a, 0xc6, 0x97, 0x60, 0xc3, 0xa5, 0x64, 0x6f, 0x5c, 0xfe, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x4f, 0x28, 0x69, 0x9e, 0xfe, 0x01, 0x00, 0x00,
}
