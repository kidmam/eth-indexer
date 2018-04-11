// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/maichain/eth-indexer/service/pb/block.proto

/*
	Package pb is a generated protocol buffer package.

	It is generated from these files:
		github.com/maichain/eth-indexer/service/pb/block.proto
		github.com/maichain/eth-indexer/service/pb/block_header.proto
		github.com/maichain/eth-indexer/service/pb/transaction.proto
		github.com/maichain/eth-indexer/service/pb/transaction_receipt.proto

	It has these top-level messages:
		BlockQueryRequest
		BlockQueryResponse
		BlockHeader
		Transaction
		TransactionQueryRequest
		TransactionQueryResponse
		TransactionReceipt
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BlockQueryRequest struct {
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *BlockQueryRequest) Reset()                    { *m = BlockQueryRequest{} }
func (*BlockQueryRequest) ProtoMessage()               {}
func (*BlockQueryRequest) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{0} }

type BlockQueryResponse struct {
	Hash         string                      `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Number       int64                       `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Nonce        []byte                      `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Transactions []*TransactionQueryResponse `protobuf:"bytes,4,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *BlockQueryResponse) Reset()                    { *m = BlockQueryResponse{} }
func (*BlockQueryResponse) ProtoMessage()               {}
func (*BlockQueryResponse) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{1} }

func init() {
	proto.RegisterType((*BlockQueryRequest)(nil), "pb.BlockQueryRequest")
	proto.RegisterType((*BlockQueryResponse)(nil), "pb.BlockQueryResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BlockService service

type BlockServiceClient interface {
	GetBlockByHash(ctx context.Context, in *BlockQueryRequest, opts ...grpc.CallOption) (*BlockQueryResponse, error)
}

type blockServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlockServiceClient(cc *grpc.ClientConn) BlockServiceClient {
	return &blockServiceClient{cc}
}

func (c *blockServiceClient) GetBlockByHash(ctx context.Context, in *BlockQueryRequest, opts ...grpc.CallOption) (*BlockQueryResponse, error) {
	out := new(BlockQueryResponse)
	err := grpc.Invoke(ctx, "/pb.BlockService/GetBlockByHash", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BlockService service

type BlockServiceServer interface {
	GetBlockByHash(context.Context, *BlockQueryRequest) (*BlockQueryResponse, error)
}

func RegisterBlockServiceServer(s *grpc.Server, srv BlockServiceServer) {
	s.RegisterService(&_BlockService_serviceDesc, srv)
}

func _BlockService_GetBlockByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockServiceServer).GetBlockByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BlockService/GetBlockByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockServiceServer).GetBlockByHash(ctx, req.(*BlockQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlockService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BlockService",
	HandlerType: (*BlockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlockByHash",
			Handler:    _BlockService_GetBlockByHash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/maichain/eth-indexer/service/pb/block.proto",
}

func (m *BlockQueryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockQueryRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Hash)))
		i += copy(dAtA[i:], m.Hash)
	}
	return i, nil
}

func (m *BlockQueryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockQueryResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Hash)))
		i += copy(dAtA[i:], m.Hash)
	}
	if m.Number != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintBlock(dAtA, i, uint64(m.Number))
	}
	if len(m.Nonce) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Nonce)))
		i += copy(dAtA[i:], m.Nonce)
	}
	if len(m.Transactions) > 0 {
		for _, msg := range m.Transactions {
			dAtA[i] = 0x22
			i++
			i = encodeVarintBlock(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintBlock(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BlockQueryRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	return n
}

func (m *BlockQueryResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Number != 0 {
		n += 1 + sovBlock(uint64(m.Number))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if len(m.Transactions) > 0 {
		for _, e := range m.Transactions {
			l = e.Size()
			n += 1 + l + sovBlock(uint64(l))
		}
	}
	return n
}

func sovBlock(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBlock(x uint64) (n int) {
	return sovBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *BlockQueryRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&BlockQueryRequest{`,
		`Hash:` + fmt.Sprintf("%v", this.Hash) + `,`,
		`}`,
	}, "")
	return s
}
func (this *BlockQueryResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&BlockQueryResponse{`,
		`Hash:` + fmt.Sprintf("%v", this.Hash) + `,`,
		`Number:` + fmt.Sprintf("%v", this.Number) + `,`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`Transactions:` + strings.Replace(fmt.Sprintf("%v", this.Transactions), "TransactionQueryResponse", "TransactionQueryResponse", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringBlock(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *BlockQueryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: BlockQueryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockQueryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlock
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
func (m *BlockQueryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: BlockQueryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockQueryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = append(m.Nonce[:0], dAtA[iNdEx:postIndex]...)
			if m.Nonce == nil {
				m.Nonce = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transactions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transactions = append(m.Transactions, &TransactionQueryResponse{})
			if err := m.Transactions[len(m.Transactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlock
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
func skipBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
				return 0, ErrInvalidLengthBlock
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBlock
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
				next, err := skipBlock(dAtA[start:])
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
	ErrInvalidLengthBlock = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlock   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/maichain/eth-indexer/service/pb/block.proto", fileDescriptorBlock)
}

var fileDescriptorBlock = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xb1, 0x4e, 0xf3, 0x30,
	0x18, 0xac, 0xdb, 0xfe, 0x95, 0x7e, 0x53, 0x21, 0xd5, 0x40, 0x15, 0xaa, 0xca, 0x8a, 0xba, 0x90,
	0xa5, 0xb1, 0x28, 0x12, 0x13, 0x03, 0xea, 0x02, 0x2b, 0x81, 0x89, 0xcd, 0x0e, 0x26, 0x09, 0xb4,
	0x76, 0x88, 0x9d, 0x8a, 0x0a, 0x21, 0x21, 0x5e, 0x81, 0xa5, 0x8f, 0xc3, 0xd8, 0x11, 0x89, 0x85,
	0x91, 0xa6, 0x3c, 0x08, 0x8a, 0x53, 0x89, 0x16, 0x58, 0xd8, 0xbe, 0x3b, 0xdd, 0x77, 0xdf, 0x9d,
	0x0d, 0xf7, 0x83, 0x48, 0x87, 0x29, 0x73, 0x7d, 0x39, 0x24, 0x43, 0x1a, 0xf9, 0x21, 0x8d, 0x04,
	0xe1, 0x3a, 0xec, 0x46, 0xe2, 0x82, 0xdf, 0xf2, 0x84, 0x28, 0x9e, 0x8c, 0x22, 0x9f, 0x93, 0x98,
	0x11, 0x36, 0x90, 0xfe, 0xb5, 0x1b, 0x27, 0x52, 0x4b, 0x54, 0x8e, 0x59, 0xab, 0x1d, 0x48, 0x19,
	0x0c, 0x38, 0xa1, 0x71, 0x44, 0xa8, 0x10, 0x52, 0x53, 0x1d, 0x49, 0xa1, 0x0a, 0x45, 0xab, 0xbb,
	0xe4, 0x1c, 0xc8, 0x40, 0x12, 0x43, 0xb3, 0xf4, 0xd2, 0x20, 0x03, 0xcc, 0xb4, 0x90, 0x1f, 0xfc,
	0x21, 0x88, 0x4e, 0xa8, 0x50, 0xd4, 0xcf, 0xaf, 0x15, 0xdb, 0x9d, 0x1d, 0xd8, 0xe8, 0xe7, 0xe9,
	0x4e, 0x52, 0x9e, 0x8c, 0x3d, 0x7e, 0x93, 0x72, 0xa5, 0x11, 0x82, 0xd5, 0x90, 0xaa, 0xd0, 0x02,
	0x36, 0x70, 0xfe, 0x7b, 0x66, 0xee, 0x4c, 0x00, 0x44, 0xcb, 0x4a, 0x15, 0x4b, 0xa1, 0xf8, 0x6f,
	0x52, 0xd4, 0x84, 0x35, 0x91, 0x0e, 0x19, 0x4f, 0xac, 0xb2, 0x0d, 0x9c, 0x8a, 0xb7, 0x40, 0x68,
	0x13, 0xfe, 0x13, 0x52, 0xf8, 0xdc, 0xaa, 0xd8, 0xc0, 0xa9, 0x7b, 0x05, 0x40, 0x87, 0xb0, 0xbe,
	0x14, 0x4b, 0x59, 0x55, 0xbb, 0xe2, 0xac, 0xf5, 0xda, 0x6e, 0xcc, 0xdc, 0xb3, 0x2f, 0x7e, 0xe5,
	0xaa, 0xb7, 0xb2, 0xd1, 0xbb, 0x82, 0x75, 0x93, 0xec, 0xb4, 0x28, 0x8a, 0xce, 0xe1, 0xfa, 0x11,
	0xd7, 0x86, 0xea, 0x8f, 0x8f, 0xf3, 0x44, 0x5b, 0xb9, 0xdb, 0x8f, 0x9e, 0xad, 0xe6, 0x77, 0xba,
	0xb0, 0xef, 0x6c, 0x3f, 0xbe, 0x7e, 0x3c, 0x95, 0x37, 0x50, 0x83, 0x8c, 0x76, 0x8b, 0xcf, 0x53,
	0xe4, 0x2e, 0xaf, 0x76, 0xdf, 0xb7, 0xa7, 0x33, 0x5c, 0x7a, 0x9b, 0xe1, 0xd2, 0x43, 0x86, 0xc1,
	0x34, 0xc3, 0xe0, 0x25, 0xc3, 0xe0, 0x3d, 0xc3, 0x60, 0x32, 0xc7, 0xa5, 0xe7, 0x39, 0x06, 0xac,
	0x66, 0x1e, 0x76, 0xef, 0x33, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x07, 0x2b, 0x09, 0x21, 0x02, 0x00,
	0x00,
}
