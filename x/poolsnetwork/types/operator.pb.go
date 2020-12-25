// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poolsnetwork/v1beta/operator.proto

package types

import (
	fmt "fmt"
	github_com_bloxapp_pools_network_shared_types "github.com/bloxapp/pools-network/shared/types"
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

// Operator represents a pools network operator, not to be confused with validator which is an entity local to the Tendermint protocol.
// An Operator has the responsibility of executing various tasks within the pools network, post collateral and so on.
type Operator struct {
	Id               uint64                                                         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EthereumAddress  github_com_bloxapp_pools_network_shared_types.EthereumAddress  `protobuf:"bytes,2,opt,name=ethereum_address,json=ethereumAddress,proto3,customtype=github.com/bloxapp/pools-network/shared/types.EthereumAddress" json:"ethereum_address"`
	ConsensusAddress github_com_bloxapp_pools_network_shared_types.ConsensusAddress `protobuf:"bytes,3,opt,name=consensus_address,json=consensusAddress,proto3,customtype=github.com/bloxapp/pools-network/shared/types.ConsensusAddress" json:"consensus_address"`
	EthStake         uint64                                                         `protobuf:"varint,4,opt,name=eth_stake,json=ethStake,proto3" json:"eth_stake,omitempty"`
	CdtBalance       uint64                                                         `protobuf:"varint,5,opt,name=cdt_balance,json=cdtBalance,proto3" json:"cdt_balance,omitempty"`
}

func (m *Operator) Reset()         { *m = Operator{} }
func (m *Operator) String() string { return proto.CompactTextString(m) }
func (*Operator) ProtoMessage()    {}
func (*Operator) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0e60acd720822a7, []int{0}
}
func (m *Operator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Operator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Operator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Operator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operator.Merge(m, src)
}
func (m *Operator) XXX_Size() int {
	return m.Size()
}
func (m *Operator) XXX_DiscardUnknown() {
	xxx_messageInfo_Operator.DiscardUnknown(m)
}

var xxx_messageInfo_Operator proto.InternalMessageInfo

func (m *Operator) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Operator) GetEthStake() uint64 {
	if m != nil {
		return m.EthStake
	}
	return 0
}

func (m *Operator) GetCdtBalance() uint64 {
	if m != nil {
		return m.CdtBalance
	}
	return 0
}

type UpdateOperator struct {
	Nonce              uint64                                                          `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ConsensusAddress   *github_com_bloxapp_pools_network_shared_types.ConsensusAddress `protobuf:"bytes,2,opt,name=consensus_address,json=consensusAddress,proto3,customtype=github.com/bloxapp/pools-network/shared/types.ConsensusAddress" json:"consensus_address,omitempty"`
	NewEthereumAddress *github_com_bloxapp_pools_network_shared_types.EthereumAddress  `protobuf:"bytes,3,opt,name=new_ethereum_address,json=newEthereumAddress,proto3,customtype=github.com/bloxapp/pools-network/shared/types.EthereumAddress" json:"new_ethereum_address,omitempty"`
	NewEthStake        uint64                                                          `protobuf:"varint,4,opt,name=new_eth_stake,json=newEthStake,proto3" json:"new_eth_stake,omitempty"`
	Exit               bool                                                            `protobuf:"varint,5,opt,name=exit,proto3" json:"exit,omitempty"`
}

func (m *UpdateOperator) Reset()         { *m = UpdateOperator{} }
func (m *UpdateOperator) String() string { return proto.CompactTextString(m) }
func (*UpdateOperator) ProtoMessage()    {}
func (*UpdateOperator) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0e60acd720822a7, []int{1}
}
func (m *UpdateOperator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateOperator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateOperator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateOperator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateOperator.Merge(m, src)
}
func (m *UpdateOperator) XXX_Size() int {
	return m.Size()
}
func (m *UpdateOperator) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateOperator.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateOperator proto.InternalMessageInfo

func (m *UpdateOperator) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *UpdateOperator) GetNewEthStake() uint64 {
	if m != nil {
		return m.NewEthStake
	}
	return 0
}

func (m *UpdateOperator) GetExit() bool {
	if m != nil {
		return m.Exit
	}
	return false
}

func init() {
	proto.RegisterType((*Operator)(nil), "poolsnetwork.v1beta1.Operator")
	proto.RegisterType((*UpdateOperator)(nil), "poolsnetwork.v1beta1.UpdateOperator")
}

func init() {
	proto.RegisterFile("poolsnetwork/v1beta/operator.proto", fileDescriptor_a0e60acd720822a7)
}

var fileDescriptor_a0e60acd720822a7 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x41, 0x6f, 0xda, 0x30,
	0x14, 0x4e, 0x32, 0x98, 0x98, 0xd9, 0x18, 0xb3, 0x72, 0x88, 0x36, 0x29, 0xa0, 0x9c, 0xb8, 0x2c,
	0x11, 0x9a, 0x76, 0xdc, 0x24, 0x32, 0xb1, 0x2b, 0x12, 0xd3, 0x2e, 0xbd, 0x44, 0x4e, 0xfc, 0x44,
	0x22, 0x20, 0xb6, 0x62, 0x53, 0xe8, 0xbf, 0xe8, 0x3f, 0xea, 0xa5, 0x07, 0x8e, 0x1c, 0xab, 0x1e,
	0x50, 0x05, 0x7f, 0xa4, 0xc2, 0x0e, 0xa8, 0xd0, 0x4a, 0x55, 0x55, 0xf5, 0xf6, 0xde, 0xe7, 0xcf,
	0xef, 0x7b, 0xef, 0x7b, 0x36, 0xf2, 0x38, 0x63, 0x13, 0x91, 0x83, 0x9c, 0xb3, 0x62, 0x1c, 0x9c,
	0x77, 0x63, 0x90, 0x24, 0x60, 0x1c, 0x0a, 0x22, 0x59, 0xe1, 0xf3, 0x82, 0x49, 0x86, 0xed, 0x87,
	0x1c, 0x5f, 0x73, 0xba, 0x5f, 0xed, 0x11, 0x1b, 0x31, 0x45, 0x08, 0x76, 0x91, 0xe6, 0x7a, 0x57,
	0x16, 0xaa, 0x0d, 0xca, 0xeb, 0xb8, 0x81, 0xac, 0x8c, 0x3a, 0x66, 0xdb, 0xec, 0x54, 0x86, 0x56,
	0x46, 0x31, 0x47, 0x4d, 0x90, 0x29, 0x14, 0x30, 0x9b, 0x46, 0x84, 0xd2, 0x02, 0x84, 0x70, 0xac,
	0xb6, 0xd9, 0xf9, 0x18, 0xf6, 0x97, 0xeb, 0x96, 0x71, 0xbb, 0x6e, 0xfd, 0x1a, 0x65, 0x32, 0x9d,
	0xc5, 0x7e, 0xc2, 0xa6, 0x41, 0x3c, 0x61, 0x0b, 0xc2, 0x79, 0xa0, 0xd4, 0xbf, 0xef, 0x5b, 0x14,
	0x29, 0x29, 0x80, 0x06, 0xf2, 0x82, 0x83, 0xf0, 0xfb, 0x65, 0xb5, 0x9e, 0x2e, 0x36, 0xfc, 0x0c,
	0xc7, 0x00, 0x16, 0xe8, 0x4b, 0xc2, 0x72, 0x01, 0xb9, 0x98, 0x89, 0x83, 0xe4, 0x3b, 0x25, 0xf9,
	0xb7, 0x94, 0xfc, 0xfd, 0x32, 0xc9, 0x3f, 0xfb, 0x72, 0x7b, 0xcd, 0x66, 0x72, 0x82, 0xe0, 0x6f,
	0xe8, 0x03, 0xc8, 0x34, 0x12, 0x92, 0x8c, 0xc1, 0xa9, 0xa8, 0xe9, 0x6b, 0x20, 0xd3, 0x7f, 0xbb,
	0x1c, 0xb7, 0x50, 0x3d, 0xa1, 0x32, 0x8a, 0xc9, 0x84, 0xe4, 0x09, 0x38, 0x55, 0x75, 0x8c, 0x12,
	0x2a, 0x43, 0x8d, 0x78, 0xd7, 0x16, 0x6a, 0xfc, 0xe7, 0x94, 0x48, 0x38, 0xf8, 0x68, 0xa3, 0x6a,
	0xce, 0x76, 0x6c, 0x6d, 0xa5, 0x4e, 0x30, 0x7b, 0x6a, 0x36, 0x6d, 0x67, 0xf8, 0x26, 0x73, 0x09,
	0x64, 0xe7, 0x30, 0x8f, 0x1e, 0xad, 0x50, 0xfb, 0xd9, 0x7b, 0xfd, 0xfa, 0x70, 0x0e, 0xf3, 0x13,
	0x0c, 0x7b, 0xe8, 0x53, 0x29, 0x7a, 0x64, 0x68, 0x5d, 0x53, 0xb5, 0xa7, 0x18, 0x55, 0x60, 0x91,
	0x49, 0x65, 0x66, 0x6d, 0xa8, 0xe2, 0x70, 0xb0, 0xdc, 0xb8, 0xe6, 0x6a, 0xe3, 0x9a, 0x77, 0x1b,
	0xd7, 0xbc, 0xdc, 0xba, 0xc6, 0x6a, 0xeb, 0x1a, 0x37, 0x5b, 0xd7, 0x38, 0xfb, 0xf9, 0x6c, 0x93,
	0x8b, 0xe0, 0xe8, 0x57, 0xa8, 0x66, 0xe3, 0xf7, 0xea, 0x81, 0xff, 0xb8, 0x0f, 0x00, 0x00, 0xff,
	0xff, 0xd6, 0xd2, 0x49, 0x67, 0x32, 0x03, 0x00, 0x00,
}

func (m *Operator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Operator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Operator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CdtBalance != 0 {
		i = encodeVarintOperator(dAtA, i, uint64(m.CdtBalance))
		i--
		dAtA[i] = 0x28
	}
	if m.EthStake != 0 {
		i = encodeVarintOperator(dAtA, i, uint64(m.EthStake))
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.ConsensusAddress.Size()
		i -= size
		if _, err := m.ConsensusAddress.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOperator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.EthereumAddress.Size()
		i -= size
		if _, err := m.EthereumAddress.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOperator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Id != 0 {
		i = encodeVarintOperator(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UpdateOperator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateOperator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateOperator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Exit {
		i--
		if m.Exit {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.NewEthStake != 0 {
		i = encodeVarintOperator(dAtA, i, uint64(m.NewEthStake))
		i--
		dAtA[i] = 0x20
	}
	if m.NewEthereumAddress != nil {
		{
			size := m.NewEthereumAddress.Size()
			i -= size
			if _, err := m.NewEthereumAddress.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintOperator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.ConsensusAddress != nil {
		{
			size := m.ConsensusAddress.Size()
			i -= size
			if _, err := m.ConsensusAddress.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintOperator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Nonce != 0 {
		i = encodeVarintOperator(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOperator(dAtA []byte, offset int, v uint64) int {
	offset -= sovOperator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Operator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovOperator(uint64(m.Id))
	}
	l = m.EthereumAddress.Size()
	n += 1 + l + sovOperator(uint64(l))
	l = m.ConsensusAddress.Size()
	n += 1 + l + sovOperator(uint64(l))
	if m.EthStake != 0 {
		n += 1 + sovOperator(uint64(m.EthStake))
	}
	if m.CdtBalance != 0 {
		n += 1 + sovOperator(uint64(m.CdtBalance))
	}
	return n
}

func (m *UpdateOperator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovOperator(uint64(m.Nonce))
	}
	if m.ConsensusAddress != nil {
		l = m.ConsensusAddress.Size()
		n += 1 + l + sovOperator(uint64(l))
	}
	if m.NewEthereumAddress != nil {
		l = m.NewEthereumAddress.Size()
		n += 1 + l + sovOperator(uint64(l))
	}
	if m.NewEthStake != 0 {
		n += 1 + sovOperator(uint64(m.NewEthStake))
	}
	if m.Exit {
		n += 2
	}
	return n
}

func sovOperator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOperator(x uint64) (n int) {
	return sovOperator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Operator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperator
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
			return fmt.Errorf("proto: Operator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Operator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthereumAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EthereumAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ConsensusAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthStake", wireType)
			}
			m.EthStake = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EthStake |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CdtBalance", wireType)
			}
			m.CdtBalance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CdtBalance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOperator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOperator
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOperator
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
func (m *UpdateOperator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperator
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
			return fmt.Errorf("proto: UpdateOperator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateOperator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_bloxapp_pools_network_shared_types.ConsensusAddress
			m.ConsensusAddress = &v
			if err := m.ConsensusAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewEthereumAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_bloxapp_pools_network_shared_types.EthereumAddress
			m.NewEthereumAddress = &v
			if err := m.NewEthereumAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewEthStake", wireType)
			}
			m.NewEthStake = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewEthStake |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exit", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Exit = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipOperator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOperator
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOperator
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
func skipOperator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOperator
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
					return 0, ErrIntOverflowOperator
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
					return 0, ErrIntOverflowOperator
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
				return 0, ErrInvalidLengthOperator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOperator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOperator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOperator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOperator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOperator = fmt.Errorf("proto: unexpected end of group")
)
