// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: Calcpb/calc.proto

package calcpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CalcClient is the client API for Calc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalcClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	PrimeNum(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (Calc_PrimeNumClient, error)
	ComputeAvg(ctx context.Context, opts ...grpc.CallOption) (Calc_ComputeAvgClient, error)
	FindMaxNum(ctx context.Context, opts ...grpc.CallOption) (Calc_FindMaxNumClient, error)
}

type calcClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcClient(cc grpc.ClientConnInterface) CalcClient {
	return &calcClient{cc}
}

func (c *calcClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calc.Calc/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcClient) PrimeNum(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (Calc_PrimeNumClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calc_ServiceDesc.Streams[0], "/calc.Calc/PrimeNum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcPrimeNumClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Calc_PrimeNumClient interface {
	Recv() (*PrimeNumberResponse, error)
	grpc.ClientStream
}

type calcPrimeNumClient struct {
	grpc.ClientStream
}

func (x *calcPrimeNumClient) Recv() (*PrimeNumberResponse, error) {
	m := new(PrimeNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcClient) ComputeAvg(ctx context.Context, opts ...grpc.CallOption) (Calc_ComputeAvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calc_ServiceDesc.Streams[1], "/calc.Calc/ComputeAvg", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcComputeAvgClient{stream}
	return x, nil
}

type Calc_ComputeAvgClient interface {
	Send(*ComputeAvgRequest) error
	CloseAndRecv() (*ComputeAvgResponse, error)
	grpc.ClientStream
}

type calcComputeAvgClient struct {
	grpc.ClientStream
}

func (x *calcComputeAvgClient) Send(m *ComputeAvgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcComputeAvgClient) CloseAndRecv() (*ComputeAvgResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAvgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcClient) FindMaxNum(ctx context.Context, opts ...grpc.CallOption) (Calc_FindMaxNumClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calc_ServiceDesc.Streams[2], "/calc.Calc/FindMaxNum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcFindMaxNumClient{stream}
	return x, nil
}

type Calc_FindMaxNumClient interface {
	Send(*FindMaxNumRequest) error
	Recv() (*FindMaxNumResponse, error)
	grpc.ClientStream
}

type calcFindMaxNumClient struct {
	grpc.ClientStream
}

func (x *calcFindMaxNumClient) Send(m *FindMaxNumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcFindMaxNumClient) Recv() (*FindMaxNumResponse, error) {
	m := new(FindMaxNumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalcServer is the server API for Calc service.
// All implementations must embed UnimplementedCalcServer
// for forward compatibility
type CalcServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	PrimeNum(*PrimeNumberRequest, Calc_PrimeNumServer) error
	ComputeAvg(Calc_ComputeAvgServer) error
	FindMaxNum(Calc_FindMaxNumServer) error
	mustEmbedUnimplementedCalcServer()
}

// UnimplementedCalcServer must be embedded to have forward compatible implementations.
type UnimplementedCalcServer struct {
}

func (UnimplementedCalcServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalcServer) PrimeNum(*PrimeNumberRequest, Calc_PrimeNumServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNum not implemented")
}
func (UnimplementedCalcServer) ComputeAvg(Calc_ComputeAvgServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAvg not implemented")
}
func (UnimplementedCalcServer) FindMaxNum(Calc_FindMaxNumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaxNum not implemented")
}
func (UnimplementedCalcServer) mustEmbedUnimplementedCalcServer() {}

// UnsafeCalcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalcServer will
// result in compilation errors.
type UnsafeCalcServer interface {
	mustEmbedUnimplementedCalcServer()
}

func RegisterCalcServer(s grpc.ServiceRegistrar, srv CalcServer) {
	s.RegisterService(&Calc_ServiceDesc, srv)
}

func _Calc_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.Calc/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calc_PrimeNum_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalcServer).PrimeNum(m, &calcPrimeNumServer{stream})
}

type Calc_PrimeNumServer interface {
	Send(*PrimeNumberResponse) error
	grpc.ServerStream
}

type calcPrimeNumServer struct {
	grpc.ServerStream
}

func (x *calcPrimeNumServer) Send(m *PrimeNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Calc_ComputeAvg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServer).ComputeAvg(&calcComputeAvgServer{stream})
}

type Calc_ComputeAvgServer interface {
	SendAndClose(*ComputeAvgResponse) error
	Recv() (*ComputeAvgRequest, error)
	grpc.ServerStream
}

type calcComputeAvgServer struct {
	grpc.ServerStream
}

func (x *calcComputeAvgServer) SendAndClose(m *ComputeAvgResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcComputeAvgServer) Recv() (*ComputeAvgRequest, error) {
	m := new(ComputeAvgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Calc_FindMaxNum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServer).FindMaxNum(&calcFindMaxNumServer{stream})
}

type Calc_FindMaxNumServer interface {
	Send(*FindMaxNumResponse) error
	Recv() (*FindMaxNumRequest, error)
	grpc.ServerStream
}

type calcFindMaxNumServer struct {
	grpc.ServerStream
}

func (x *calcFindMaxNumServer) Send(m *FindMaxNumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcFindMaxNumServer) Recv() (*FindMaxNumRequest, error) {
	m := new(FindMaxNumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Calc_ServiceDesc is the grpc.ServiceDesc for Calc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calc.Calc",
	HandlerType: (*CalcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _Calc_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNum",
			Handler:       _Calc_PrimeNum_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAvg",
			Handler:       _Calc_ComputeAvg_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaxNum",
			Handler:       _Calc_FindMaxNum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "Calcpb/calc.proto",
}
