// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethereum_messages.proto

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
// Request: Generate a Ethereum address from a seed, device sends back the address in a Success message
// @next Failure
// @next ResponseEthereumAddress
type EthereumAddress struct {
	AddressN             *uint32  `protobuf:"varint,1,req,name=address_n,json=addressN" json:"address_n,omitempty"`
	StartIndex           *uint32  `protobuf:"varint,2,opt,name=start_index,json=startIndex" json:"start_index,omitempty"`
	ConfirmAddress       *bool    `protobuf:"varint,3,opt,name=confirm_address,json=confirmAddress" json:"confirm_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EthereumAddress) Reset()         { *m = EthereumAddress{} }
func (m *EthereumAddress) String() string { return proto.CompactTextString(m) }
func (*EthereumAddress) ProtoMessage()    {}
func (*EthereumAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethereum_messages_8fe0d75077b972c5, []int{0}
}
func (m *EthereumAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthereumAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthereumAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *EthereumAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthereumAddress.Merge(dst, src)
}
func (m *EthereumAddress) XXX_Size() int {
	return m.Size()
}
func (m *EthereumAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_EthereumAddress.DiscardUnknown(m)
}

var xxx_messageInfo_EthereumAddress proto.InternalMessageInfo

func (m *EthereumAddress) GetAddressN() uint32 {
	if m != nil && m.AddressN != nil {
		return *m.AddressN
	}
	return 0
}

func (m *EthereumAddress) GetStartIndex() uint32 {
	if m != nil && m.StartIndex != nil {
		return *m.StartIndex
	}
	return 0
}

func (m *EthereumAddress) GetConfirmAddress() bool {
	if m != nil && m.ConfirmAddress != nil {
		return *m.ConfirmAddress
	}
	return false
}

// *
// Response: Return the generated Ethereum address
// @prev EthereumAddress
type ResponseEthereumAddress struct {
	Addresses            [][]byte `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseEthereumAddress) Reset()         { *m = ResponseEthereumAddress{} }
func (m *ResponseEthereumAddress) String() string { return proto.CompactTextString(m) }
func (*ResponseEthereumAddress) ProtoMessage()    {}
func (*ResponseEthereumAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethereum_messages_8fe0d75077b972c5, []int{1}
}
func (m *ResponseEthereumAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseEthereumAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseEthereumAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ResponseEthereumAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseEthereumAddress.Merge(dst, src)
}
func (m *ResponseEthereumAddress) XXX_Size() int {
	return m.Size()
}
func (m *ResponseEthereumAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseEthereumAddress.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseEthereumAddress proto.InternalMessageInfo

func (m *ResponseEthereumAddress) GetAddresses() [][]byte {
	if m != nil {
		return m.Addresses
	}
	return nil
}

// *
// Request: Send a part of data to Skywallet
// @next Failure
// @next TxRequest
type EthereumTxAck struct {
	Tx                   *EthereumTransactionType `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *EthereumTxAck) Reset()         { *m = EthereumTxAck{} }
func (m *EthereumTxAck) String() string { return proto.CompactTextString(m) }
func (*EthereumTxAck) ProtoMessage()    {}
func (*EthereumTxAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethereum_messages_8fe0d75077b972c5, []int{2}
}
func (m *EthereumTxAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthereumTxAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthereumTxAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *EthereumTxAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthereumTxAck.Merge(dst, src)
}
func (m *EthereumTxAck) XXX_Size() int {
	return m.Size()
}
func (m *EthereumTxAck) XXX_DiscardUnknown() {
	xxx_messageInfo_EthereumTxAck.DiscardUnknown(m)
}

var xxx_messageInfo_EthereumTxAck proto.InternalMessageInfo

func (m *EthereumTxAck) GetTx() *EthereumTransactionType {
	if m != nil {
		return m.Tx
	}
	return nil
}

func init() {
	proto.RegisterType((*EthereumAddress)(nil), "EthereumAddress")
	proto.RegisterType((*ResponseEthereumAddress)(nil), "ResponseEthereumAddress")
	proto.RegisterType((*EthereumTxAck)(nil), "EthereumTxAck")
}
func (m *EthereumAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthereumAddress) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AddressN == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("address_n")
	} else {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEthereumMessages(dAtA, i, uint64(*m.AddressN))
	}
	if m.StartIndex != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEthereumMessages(dAtA, i, uint64(*m.StartIndex))
	}
	if m.ConfirmAddress != nil {
		dAtA[i] = 0x18
		i++
		if *m.ConfirmAddress {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ResponseEthereumAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseEthereumAddress) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for _, b := range m.Addresses {
			dAtA[i] = 0xa
			i++
			i = encodeVarintEthereumMessages(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *EthereumTxAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthereumTxAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Tx != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEthereumMessages(dAtA, i, uint64(m.Tx.Size()))
		n1, err := m.Tx.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintEthereumMessages(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EthereumAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AddressN != nil {
		n += 1 + sovEthereumMessages(uint64(*m.AddressN))
	}
	if m.StartIndex != nil {
		n += 1 + sovEthereumMessages(uint64(*m.StartIndex))
	}
	if m.ConfirmAddress != nil {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResponseEthereumAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for _, b := range m.Addresses {
			l = len(b)
			n += 1 + l + sovEthereumMessages(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EthereumTxAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovEthereumMessages(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEthereumMessages(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEthereumMessages(x uint64) (n int) {
	return sovEthereumMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EthereumAddress) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEthereumMessages
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
			return fmt.Errorf("proto: EthereumAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthereumAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressN", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthereumMessages
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartIndex", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthereumMessages
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
			m.StartIndex = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfirmAddress", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthereumMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.ConfirmAddress = &b
		default:
			iNdEx = preIndex
			skippy, err := skipEthereumMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEthereumMessages
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

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseEthereumAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEthereumMessages
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
			return fmt.Errorf("proto: ResponseEthereumAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseEthereumAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthereumMessages
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
				return ErrInvalidLengthEthereumMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, make([]byte, postIndex-iNdEx))
			copy(m.Addresses[len(m.Addresses)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEthereumMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEthereumMessages
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
func (m *EthereumTxAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEthereumMessages
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
			return fmt.Errorf("proto: EthereumTxAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthereumTxAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthereumMessages
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
				return ErrInvalidLengthEthereumMessages
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tx == nil {
				m.Tx = &EthereumTransactionType{}
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEthereumMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEthereumMessages
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
func skipEthereumMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEthereumMessages
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
					return 0, ErrIntOverflowEthereumMessages
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
					return 0, ErrIntOverflowEthereumMessages
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
				return 0, ErrInvalidLengthEthereumMessages
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEthereumMessages
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
				next, err := skipEthereumMessages(dAtA[start:])
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
	ErrInvalidLengthEthereumMessages = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEthereumMessages   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("ethereum_messages.proto", fileDescriptor_ethereum_messages_8fe0d75077b972c5)
}

var fileDescriptor_ethereum_messages_8fe0d75077b972c5 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd0, 0xb1, 0x4e, 0xf3, 0x30,
	0x10, 0x07, 0x70, 0x39, 0x5d, 0xf2, 0xb9, 0x5f, 0xa8, 0x14, 0x21, 0xd5, 0x02, 0x14, 0xa2, 0x2c,
	0x78, 0xca, 0xc0, 0x82, 0x18, 0x5b, 0xa9, 0x03, 0x0b, 0x83, 0xe9, 0xc4, 0x12, 0x99, 0xe4, 0x0a,
	0x51, 0x1b, 0x3b, 0xf2, 0x5d, 0x45, 0xf2, 0x86, 0x8c, 0x3c, 0x02, 0xca, 0x93, 0xa0, 0xd0, 0x24,
	0x48, 0x4c, 0xb6, 0x7e, 0xa7, 0xb3, 0xff, 0xfa, 0xf3, 0x25, 0xd0, 0x1b, 0x38, 0x38, 0x56, 0x59,
	0x05, 0x88, 0xfa, 0x15, 0x30, 0xad, 0x9d, 0x25, 0x7b, 0x71, 0x3e, 0x0d, 0xa8, 0xad, 0x47, 0x4d,
	0x1a, 0xbe, 0xd8, 0x0c, 0xbe, 0x2a, 0x0a, 0x07, 0x88, 0xe1, 0x25, 0xff, 0xa7, 0x4f, 0xd7, 0xcc,
	0x08, 0x16, 0x7b, 0x32, 0x50, 0xfe, 0x00, 0x8f, 0xe1, 0x35, 0x9f, 0x23, 0x69, 0x47, 0x59, 0x69,
	0x0a, 0x68, 0x84, 0x17, 0x33, 0x19, 0x28, 0xfe, 0x43, 0x0f, 0xbd, 0x84, 0x37, 0x7c, 0x91, 0x5b,
	0xb3, 0x2b, 0x5d, 0x95, 0x0d, 0x4b, 0x62, 0x16, 0x33, 0xe9, 0xab, 0xb3, 0x81, 0x87, 0x6f, 0x92,
	0x3b, 0xbe, 0x54, 0x80, 0xb5, 0x35, 0x08, 0x7f, 0x13, 0x5c, 0x4d, 0x09, 0x00, 0x05, 0x8b, 0x67,
	0xf2, 0xbf, 0xfa, 0x85, 0xe4, 0x9e, 0x07, 0xe3, 0xc2, 0xb6, 0x59, 0xe5, 0xfb, 0x50, 0x72, 0x8f,
	0x1a, 0xc1, 0x62, 0x26, 0xe7, 0xb7, 0x22, 0x9d, 0x66, 0x4e, 0x1b, 0xd4, 0x39, 0x95, 0xd6, 0x6c,
	0xdb, 0x1a, 0x94, 0x47, 0xcd, 0x7a, 0xf3, 0xd1, 0x45, 0xec, 0xb3, 0x8b, 0xd8, 0x57, 0x17, 0x31,
	0x1e, 0x19, 0xa0, 0x14, 0xf7, 0x6d, 0x6e, 0x4b, 0xd3, 0x9f, 0xef, 0xfa, 0x70, 0x00, 0x3a, 0x15,
	0xf3, 0x72, 0xdc, 0xad, 0x83, 0xa7, 0xd1, 0xfa, 0x07, 0x9e, 0xfd, 0xb1, 0xd0, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x95, 0xe0, 0xd0, 0x2e, 0x64, 0x01, 0x00, 0x00,
}
