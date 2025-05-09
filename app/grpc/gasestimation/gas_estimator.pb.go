// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: celestia/core/v1/gas_estimation/gas_estimator.proto

package gasestimation

import (
	context "context"
	encoding_binary "encoding/binary"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// TxPriority is the priority level of the requested gas price.
// The following priority levels are defined:
// - High Priority: The gas price is the price at the start of the top 10% of
// transactions’ gas prices from the last 5 blocks.
// - Medium Priority: The gas price is the mean of all gas prices from the last
// 5 blocks.
// - Low Priority: The gas price is the value at the end of the lowest 10% of
// gas prices from the last 5 blocks.
// - Unspecified Priority (default): This is equivalent to the Medium priority,
// using the mean of all gas prices from the last 5 blocks.
type TxPriority int32

const (
	// TX_PRIORITY_UNSPECIFIED none priority, the default priority level, which is
	// equivalent to the TX_PRIORITY_MEDIUM priority.
	TxPriority_TX_PRIORITY_UNSPECIFIED TxPriority = 0
	// TX_PRIORITY_LOW low priority.
	TxPriority_TX_PRIORITY_LOW TxPriority = 1
	// TX_PRIORITY_MEDIUM medium priority.
	TxPriority_TX_PRIORITY_MEDIUM TxPriority = 2
	// TX_PRIORITY_HIGH high priority.
	TxPriority_TX_PRIORITY_HIGH TxPriority = 3
)

var TxPriority_name = map[int32]string{
	0: "TX_PRIORITY_UNSPECIFIED",
	1: "TX_PRIORITY_LOW",
	2: "TX_PRIORITY_MEDIUM",
	3: "TX_PRIORITY_HIGH",
}

var TxPriority_value = map[string]int32{
	"TX_PRIORITY_UNSPECIFIED": 0,
	"TX_PRIORITY_LOW":         1,
	"TX_PRIORITY_MEDIUM":      2,
	"TX_PRIORITY_HIGH":        3,
}

func (x TxPriority) String() string {
	return proto.EnumName(TxPriority_name, int32(x))
}

func (TxPriority) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67d02876d749b9cc, []int{0}
}

// EstimateGasPriceRequest the request to estimate the gas price of the network.
// Takes a priority enum to define the priority level.
type EstimateGasPriceRequest struct {
	TxPriority TxPriority `protobuf:"varint,1,opt,name=tx_priority,json=txPriority,proto3,enum=celestia.core.v1.gas_estimation.TxPriority" json:"tx_priority,omitempty"`
}

func (m *EstimateGasPriceRequest) Reset()         { *m = EstimateGasPriceRequest{} }
func (m *EstimateGasPriceRequest) String() string { return proto.CompactTextString(m) }
func (*EstimateGasPriceRequest) ProtoMessage()    {}
func (*EstimateGasPriceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67d02876d749b9cc, []int{0}
}
func (m *EstimateGasPriceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EstimateGasPriceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EstimateGasPriceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EstimateGasPriceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstimateGasPriceRequest.Merge(m, src)
}
func (m *EstimateGasPriceRequest) XXX_Size() int {
	return m.Size()
}
func (m *EstimateGasPriceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EstimateGasPriceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EstimateGasPriceRequest proto.InternalMessageInfo

func (m *EstimateGasPriceRequest) GetTxPriority() TxPriority {
	if m != nil {
		return m.TxPriority
	}
	return TxPriority_TX_PRIORITY_UNSPECIFIED
}

// EstimateGasPriceResponse the response of the gas price estimation.
type EstimateGasPriceResponse struct {
	EstimatedGasPrice float64 `protobuf:"fixed64,1,opt,name=estimated_gas_price,json=estimatedGasPrice,proto3" json:"estimated_gas_price,omitempty"`
}

func (m *EstimateGasPriceResponse) Reset()         { *m = EstimateGasPriceResponse{} }
func (m *EstimateGasPriceResponse) String() string { return proto.CompactTextString(m) }
func (*EstimateGasPriceResponse) ProtoMessage()    {}
func (*EstimateGasPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67d02876d749b9cc, []int{1}
}
func (m *EstimateGasPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EstimateGasPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EstimateGasPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EstimateGasPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstimateGasPriceResponse.Merge(m, src)
}
func (m *EstimateGasPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *EstimateGasPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EstimateGasPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EstimateGasPriceResponse proto.InternalMessageInfo

func (m *EstimateGasPriceResponse) GetEstimatedGasPrice() float64 {
	if m != nil {
		return m.EstimatedGasPrice
	}
	return 0
}

// EstimateGasPriceAndUsageRequest the request to estimate the gas price of the
// network and also the gas used for the provided transaction.
type EstimateGasPriceAndUsageRequest struct {
	TxPriority TxPriority `protobuf:"varint,1,opt,name=tx_priority,json=txPriority,proto3,enum=celestia.core.v1.gas_estimation.TxPriority" json:"tx_priority,omitempty"`
	TxBytes    []byte     `protobuf:"bytes,2,opt,name=tx_bytes,json=txBytes,proto3" json:"tx_bytes,omitempty"`
}

func (m *EstimateGasPriceAndUsageRequest) Reset()         { *m = EstimateGasPriceAndUsageRequest{} }
func (m *EstimateGasPriceAndUsageRequest) String() string { return proto.CompactTextString(m) }
func (*EstimateGasPriceAndUsageRequest) ProtoMessage()    {}
func (*EstimateGasPriceAndUsageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67d02876d749b9cc, []int{2}
}
func (m *EstimateGasPriceAndUsageRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EstimateGasPriceAndUsageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EstimateGasPriceAndUsageRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EstimateGasPriceAndUsageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstimateGasPriceAndUsageRequest.Merge(m, src)
}
func (m *EstimateGasPriceAndUsageRequest) XXX_Size() int {
	return m.Size()
}
func (m *EstimateGasPriceAndUsageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EstimateGasPriceAndUsageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EstimateGasPriceAndUsageRequest proto.InternalMessageInfo

func (m *EstimateGasPriceAndUsageRequest) GetTxPriority() TxPriority {
	if m != nil {
		return m.TxPriority
	}
	return TxPriority_TX_PRIORITY_UNSPECIFIED
}

func (m *EstimateGasPriceAndUsageRequest) GetTxBytes() []byte {
	if m != nil {
		return m.TxBytes
	}
	return nil
}

// EstimateGasPriceAndUsageResponse the response of the gas price and used
// estimation.
type EstimateGasPriceAndUsageResponse struct {
	EstimatedGasPrice float64 `protobuf:"fixed64,1,opt,name=estimated_gas_price,json=estimatedGasPrice,proto3" json:"estimated_gas_price,omitempty"`
	EstimatedGasUsed  uint64  `protobuf:"varint,2,opt,name=estimated_gas_used,json=estimatedGasUsed,proto3" json:"estimated_gas_used,omitempty"`
}

func (m *EstimateGasPriceAndUsageResponse) Reset()         { *m = EstimateGasPriceAndUsageResponse{} }
func (m *EstimateGasPriceAndUsageResponse) String() string { return proto.CompactTextString(m) }
func (*EstimateGasPriceAndUsageResponse) ProtoMessage()    {}
func (*EstimateGasPriceAndUsageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67d02876d749b9cc, []int{3}
}
func (m *EstimateGasPriceAndUsageResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EstimateGasPriceAndUsageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EstimateGasPriceAndUsageResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EstimateGasPriceAndUsageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstimateGasPriceAndUsageResponse.Merge(m, src)
}
func (m *EstimateGasPriceAndUsageResponse) XXX_Size() int {
	return m.Size()
}
func (m *EstimateGasPriceAndUsageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EstimateGasPriceAndUsageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EstimateGasPriceAndUsageResponse proto.InternalMessageInfo

func (m *EstimateGasPriceAndUsageResponse) GetEstimatedGasPrice() float64 {
	if m != nil {
		return m.EstimatedGasPrice
	}
	return 0
}

func (m *EstimateGasPriceAndUsageResponse) GetEstimatedGasUsed() uint64 {
	if m != nil {
		return m.EstimatedGasUsed
	}
	return 0
}

func init() {
	proto.RegisterEnum("celestia.core.v1.gas_estimation.TxPriority", TxPriority_name, TxPriority_value)
	proto.RegisterType((*EstimateGasPriceRequest)(nil), "celestia.core.v1.gas_estimation.EstimateGasPriceRequest")
	proto.RegisterType((*EstimateGasPriceResponse)(nil), "celestia.core.v1.gas_estimation.EstimateGasPriceResponse")
	proto.RegisterType((*EstimateGasPriceAndUsageRequest)(nil), "celestia.core.v1.gas_estimation.EstimateGasPriceAndUsageRequest")
	proto.RegisterType((*EstimateGasPriceAndUsageResponse)(nil), "celestia.core.v1.gas_estimation.EstimateGasPriceAndUsageResponse")
}

func init() {
	proto.RegisterFile("celestia/core/v1/gas_estimation/gas_estimator.proto", fileDescriptor_67d02876d749b9cc)
}

var fileDescriptor_67d02876d749b9cc = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0x4e, 0xcd, 0x49,
	0x2d, 0x2e, 0xc9, 0x4c, 0xd4, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x33, 0xd4, 0x4f, 0x4f, 0x2c,
	0x8e, 0x07, 0x89, 0xe4, 0x26, 0x96, 0x64, 0xe6, 0xe7, 0x21, 0x73, 0xf3, 0x8b, 0xf4, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0xe4, 0x61, 0x9a, 0xf4, 0x40, 0x9a, 0xf4, 0xca, 0x0c, 0xf5, 0x50, 0x35,
	0x29, 0xa5, 0x73, 0x89, 0xbb, 0x42, 0x78, 0xa9, 0xee, 0x89, 0xc5, 0x01, 0x45, 0x99, 0xc9, 0xa9,
	0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x3e, 0x5c, 0xdc, 0x25, 0x15, 0xf1, 0x05, 0x45,
	0x99, 0xf9, 0x45, 0x99, 0x25, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0xda, 0x7a, 0x04,
	0x4c, 0xd4, 0x0b, 0xa9, 0x08, 0x80, 0x6a, 0x09, 0xe2, 0x2a, 0x81, 0xb3, 0x95, 0xbc, 0xb8, 0x24,
	0x30, 0x2d, 0x2a, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0xd2, 0xe3, 0x12, 0x86, 0x1a, 0x90, 0x9a,
	0x12, 0x0f, 0x32, 0xae, 0x00, 0x24, 0x0d, 0xb6, 0x91, 0x31, 0x48, 0x10, 0x2e, 0x05, 0xd3, 0xa7,
	0xd4, 0xc5, 0xc8, 0x25, 0x8f, 0x6e, 0x98, 0x63, 0x5e, 0x4a, 0x68, 0x71, 0x62, 0x3a, 0x6d, 0x5c,
	0x2f, 0x24, 0xc9, 0xc5, 0x51, 0x52, 0x11, 0x9f, 0x54, 0x59, 0x92, 0x5a, 0x2c, 0xc1, 0xa4, 0xc0,
	0xa8, 0xc1, 0x13, 0xc4, 0x5e, 0x52, 0xe1, 0x04, 0xe2, 0x2a, 0x35, 0x30, 0x72, 0x29, 0xe0, 0x76,
	0x0c, 0x79, 0x3e, 0x14, 0xd2, 0xe1, 0x12, 0x42, 0x55, 0x5f, 0x5a, 0x9c, 0x9a, 0x02, 0xb6, 0x99,
	0x25, 0x48, 0x00, 0x59, 0x79, 0x68, 0x71, 0x6a, 0x8a, 0x56, 0x0e, 0x17, 0x17, 0xc2, 0xdd, 0x42,
	0xd2, 0x5c, 0xe2, 0x21, 0x11, 0xf1, 0x01, 0x41, 0x9e, 0xfe, 0x41, 0x9e, 0x21, 0x91, 0xf1, 0xa1,
	0x7e, 0xc1, 0x01, 0xae, 0xce, 0x9e, 0x6e, 0x9e, 0xae, 0x2e, 0x02, 0x0c, 0x42, 0xc2, 0x5c, 0xfc,
	0xc8, 0x92, 0x3e, 0xfe, 0xe1, 0x02, 0x8c, 0x42, 0x62, 0x5c, 0x42, 0xc8, 0x82, 0xbe, 0xae, 0x2e,
	0x9e, 0xa1, 0xbe, 0x02, 0x4c, 0x42, 0x22, 0x5c, 0x02, 0xc8, 0xe2, 0x1e, 0x9e, 0xee, 0x1e, 0x02,
	0xcc, 0x46, 0xfb, 0x98, 0xb8, 0x78, 0xdc, 0x13, 0x8b, 0x5d, 0x61, 0x49, 0x4d, 0xa8, 0x93, 0x91,
	0x4b, 0x00, 0x3d, 0x04, 0x84, 0x2c, 0x08, 0x06, 0x35, 0x8e, 0x74, 0x27, 0x65, 0x49, 0x86, 0x4e,
	0x48, 0x30, 0x2b, 0x31, 0x08, 0x2d, 0x64, 0xc4, 0x4c, 0x67, 0xb0, 0xd8, 0x10, 0x72, 0x20, 0xd9,
	0x64, 0xb4, 0x54, 0x25, 0xe5, 0x48, 0x81, 0x09, 0x30, 0x37, 0x3a, 0x85, 0x9c, 0x78, 0x24, 0xc7,
	0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c,
	0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x55, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72,
	0x7e, 0xae, 0x3e, 0xcc, 0xa2, 0xfc, 0xa2, 0x74, 0x38, 0x5b, 0x37, 0xb1, 0xa0, 0x40, 0x1f, 0x84,
	0xd3, 0x8b, 0x0a, 0x92, 0x41, 0x19, 0x1e, 0x61, 0x71, 0x12, 0x1b, 0x38, 0xc7, 0x1b, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x47, 0x63, 0x14, 0xb9, 0x28, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GasEstimatorClient is the client API for GasEstimator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GasEstimatorClient interface {
	// estimateGasPrice takes a transaction priority and estimates the gas price
	// based on the gas prices of the transactions in the last five blocks. If no
	// transaction is found in the last five blocks, return the network min gas
	// price. It's up to the light client to set the gas price in this case to the
	// minimum gas price set by that node.
	EstimateGasPrice(ctx context.Context, in *EstimateGasPriceRequest, opts ...grpc.CallOption) (*EstimateGasPriceResponse, error)
	// EstimateGasPriceAndUsage takes a transaction priority and a transaction
	// bytes and estimates the gas price and the gas used for that transaction.
	// The gas price estimation is based on the gas prices of the transactions in
	// the last five blocks. If no transaction is found in the last five blocks,
	// return the network min gas price. It's up to the light client to set the
	// gas price in this case to the minimum gas price set by that node. The gas
	// used is estimated using the state machine simulation.
	EstimateGasPriceAndUsage(ctx context.Context, in *EstimateGasPriceAndUsageRequest, opts ...grpc.CallOption) (*EstimateGasPriceAndUsageResponse, error)
}

type gasEstimatorClient struct {
	cc grpc1.ClientConn
}

func NewGasEstimatorClient(cc grpc1.ClientConn) GasEstimatorClient {
	return &gasEstimatorClient{cc}
}

func (c *gasEstimatorClient) EstimateGasPrice(ctx context.Context, in *EstimateGasPriceRequest, opts ...grpc.CallOption) (*EstimateGasPriceResponse, error) {
	out := new(EstimateGasPriceResponse)
	err := c.cc.Invoke(ctx, "/celestia.core.v1.gas_estimation.GasEstimator/EstimateGasPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gasEstimatorClient) EstimateGasPriceAndUsage(ctx context.Context, in *EstimateGasPriceAndUsageRequest, opts ...grpc.CallOption) (*EstimateGasPriceAndUsageResponse, error) {
	out := new(EstimateGasPriceAndUsageResponse)
	err := c.cc.Invoke(ctx, "/celestia.core.v1.gas_estimation.GasEstimator/EstimateGasPriceAndUsage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GasEstimatorServer is the server API for GasEstimator service.
type GasEstimatorServer interface {
	// estimateGasPrice takes a transaction priority and estimates the gas price
	// based on the gas prices of the transactions in the last five blocks. If no
	// transaction is found in the last five blocks, return the network min gas
	// price. It's up to the light client to set the gas price in this case to the
	// minimum gas price set by that node.
	EstimateGasPrice(context.Context, *EstimateGasPriceRequest) (*EstimateGasPriceResponse, error)
	// EstimateGasPriceAndUsage takes a transaction priority and a transaction
	// bytes and estimates the gas price and the gas used for that transaction.
	// The gas price estimation is based on the gas prices of the transactions in
	// the last five blocks. If no transaction is found in the last five blocks,
	// return the network min gas price. It's up to the light client to set the
	// gas price in this case to the minimum gas price set by that node. The gas
	// used is estimated using the state machine simulation.
	EstimateGasPriceAndUsage(context.Context, *EstimateGasPriceAndUsageRequest) (*EstimateGasPriceAndUsageResponse, error)
}

// UnimplementedGasEstimatorServer can be embedded to have forward compatible implementations.
type UnimplementedGasEstimatorServer struct {
}

func (*UnimplementedGasEstimatorServer) EstimateGasPrice(ctx context.Context, req *EstimateGasPriceRequest) (*EstimateGasPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstimateGasPrice not implemented")
}
func (*UnimplementedGasEstimatorServer) EstimateGasPriceAndUsage(ctx context.Context, req *EstimateGasPriceAndUsageRequest) (*EstimateGasPriceAndUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstimateGasPriceAndUsage not implemented")
}

func RegisterGasEstimatorServer(s grpc1.Server, srv GasEstimatorServer) {
	s.RegisterService(&_GasEstimator_serviceDesc, srv)
}

func _GasEstimator_EstimateGasPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstimateGasPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GasEstimatorServer).EstimateGasPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/celestia.core.v1.gas_estimation.GasEstimator/EstimateGasPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GasEstimatorServer).EstimateGasPrice(ctx, req.(*EstimateGasPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GasEstimator_EstimateGasPriceAndUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstimateGasPriceAndUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GasEstimatorServer).EstimateGasPriceAndUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/celestia.core.v1.gas_estimation.GasEstimator/EstimateGasPriceAndUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GasEstimatorServer).EstimateGasPriceAndUsage(ctx, req.(*EstimateGasPriceAndUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var GasEstimator_serviceDesc = _GasEstimator_serviceDesc
var _GasEstimator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "celestia.core.v1.gas_estimation.GasEstimator",
	HandlerType: (*GasEstimatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EstimateGasPrice",
			Handler:    _GasEstimator_EstimateGasPrice_Handler,
		},
		{
			MethodName: "EstimateGasPriceAndUsage",
			Handler:    _GasEstimator_EstimateGasPriceAndUsage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "celestia/core/v1/gas_estimation/gas_estimator.proto",
}

func (m *EstimateGasPriceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EstimateGasPriceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EstimateGasPriceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TxPriority != 0 {
		i = encodeVarintGasEstimator(dAtA, i, uint64(m.TxPriority))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EstimateGasPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EstimateGasPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EstimateGasPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EstimatedGasPrice != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.EstimatedGasPrice))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func (m *EstimateGasPriceAndUsageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EstimateGasPriceAndUsageRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EstimateGasPriceAndUsageRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxBytes) > 0 {
		i -= len(m.TxBytes)
		copy(dAtA[i:], m.TxBytes)
		i = encodeVarintGasEstimator(dAtA, i, uint64(len(m.TxBytes)))
		i--
		dAtA[i] = 0x12
	}
	if m.TxPriority != 0 {
		i = encodeVarintGasEstimator(dAtA, i, uint64(m.TxPriority))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EstimateGasPriceAndUsageResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EstimateGasPriceAndUsageResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EstimateGasPriceAndUsageResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EstimatedGasUsed != 0 {
		i = encodeVarintGasEstimator(dAtA, i, uint64(m.EstimatedGasUsed))
		i--
		dAtA[i] = 0x10
	}
	if m.EstimatedGasPrice != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.EstimatedGasPrice))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func encodeVarintGasEstimator(dAtA []byte, offset int, v uint64) int {
	offset -= sovGasEstimator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EstimateGasPriceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TxPriority != 0 {
		n += 1 + sovGasEstimator(uint64(m.TxPriority))
	}
	return n
}

func (m *EstimateGasPriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EstimatedGasPrice != 0 {
		n += 9
	}
	return n
}

func (m *EstimateGasPriceAndUsageRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TxPriority != 0 {
		n += 1 + sovGasEstimator(uint64(m.TxPriority))
	}
	l = len(m.TxBytes)
	if l > 0 {
		n += 1 + l + sovGasEstimator(uint64(l))
	}
	return n
}

func (m *EstimateGasPriceAndUsageResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EstimatedGasPrice != 0 {
		n += 9
	}
	if m.EstimatedGasUsed != 0 {
		n += 1 + sovGasEstimator(uint64(m.EstimatedGasUsed))
	}
	return n
}

func sovGasEstimator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGasEstimator(x uint64) (n int) {
	return sovGasEstimator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EstimateGasPriceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGasEstimator
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
			return fmt.Errorf("proto: EstimateGasPriceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EstimateGasPriceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxPriority", wireType)
			}
			m.TxPriority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGasEstimator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxPriority |= TxPriority(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGasEstimator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGasEstimator
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
func (m *EstimateGasPriceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGasEstimator
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
			return fmt.Errorf("proto: EstimateGasPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EstimateGasPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field EstimatedGasPrice", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.EstimatedGasPrice = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipGasEstimator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGasEstimator
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
func (m *EstimateGasPriceAndUsageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGasEstimator
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
			return fmt.Errorf("proto: EstimateGasPriceAndUsageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EstimateGasPriceAndUsageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxPriority", wireType)
			}
			m.TxPriority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGasEstimator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxPriority |= TxPriority(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGasEstimator
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
				return ErrInvalidLengthGasEstimator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGasEstimator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxBytes = append(m.TxBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.TxBytes == nil {
				m.TxBytes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGasEstimator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGasEstimator
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
func (m *EstimateGasPriceAndUsageResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGasEstimator
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
			return fmt.Errorf("proto: EstimateGasPriceAndUsageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EstimateGasPriceAndUsageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field EstimatedGasPrice", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.EstimatedGasPrice = float64(math.Float64frombits(v))
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EstimatedGasUsed", wireType)
			}
			m.EstimatedGasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGasEstimator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EstimatedGasUsed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGasEstimator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGasEstimator
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
func skipGasEstimator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGasEstimator
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
					return 0, ErrIntOverflowGasEstimator
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
					return 0, ErrIntOverflowGasEstimator
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
				return 0, ErrInvalidLengthGasEstimator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGasEstimator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGasEstimator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGasEstimator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGasEstimator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGasEstimator = fmt.Errorf("proto: unexpected end of group")
)
