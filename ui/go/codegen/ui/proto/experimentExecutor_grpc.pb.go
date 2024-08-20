// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: experiment/experimentExecutor.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ExperimentExecutor_ExecuteExperiment_FullMethodName = "/experiment.ExperimentExecutor/ExecuteExperiment"
)

// ExperimentExecutorClient is the client API for ExperimentExecutor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The experiment executor service definition.
type ExperimentExecutorClient interface {
	// executes an experiment
	ExecuteExperiment(ctx context.Context, in *ExecuteExperimentRequest, opts ...grpc.CallOption) (*ExecuteExperimentResponse, error)
}

type experimentExecutorClient struct {
	cc grpc.ClientConnInterface
}

func NewExperimentExecutorClient(cc grpc.ClientConnInterface) ExperimentExecutorClient {
	return &experimentExecutorClient{cc}
}

func (c *experimentExecutorClient) ExecuteExperiment(ctx context.Context, in *ExecuteExperimentRequest, opts ...grpc.CallOption) (*ExecuteExperimentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExecuteExperimentResponse)
	err := c.cc.Invoke(ctx, ExperimentExecutor_ExecuteExperiment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExperimentExecutorServer is the server API for ExperimentExecutor service.
// All implementations must embed UnimplementedExperimentExecutorServer
// for forward compatibility
//
// The experiment executor service definition.
type ExperimentExecutorServer interface {
	// executes an experiment
	ExecuteExperiment(context.Context, *ExecuteExperimentRequest) (*ExecuteExperimentResponse, error)
	mustEmbedUnimplementedExperimentExecutorServer()
}

// UnimplementedExperimentExecutorServer must be embedded to have forward compatible implementations.
type UnimplementedExperimentExecutorServer struct {
}

func (UnimplementedExperimentExecutorServer) ExecuteExperiment(context.Context, *ExecuteExperimentRequest) (*ExecuteExperimentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteExperiment not implemented")
}
func (UnimplementedExperimentExecutorServer) mustEmbedUnimplementedExperimentExecutorServer() {}

// UnsafeExperimentExecutorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExperimentExecutorServer will
// result in compilation errors.
type UnsafeExperimentExecutorServer interface {
	mustEmbedUnimplementedExperimentExecutorServer()
}

func RegisterExperimentExecutorServer(s grpc.ServiceRegistrar, srv ExperimentExecutorServer) {
	s.RegisterService(&ExperimentExecutor_ServiceDesc, srv)
}

func _ExperimentExecutor_ExecuteExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteExperimentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentExecutorServer).ExecuteExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExperimentExecutor_ExecuteExperiment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentExecutorServer).ExecuteExperiment(ctx, req.(*ExecuteExperimentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExperimentExecutor_ServiceDesc is the grpc.ServiceDesc for ExperimentExecutor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExperimentExecutor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "experiment.ExperimentExecutor",
	HandlerType: (*ExperimentExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteExperiment",
			Handler:    _ExperimentExecutor_ExecuteExperiment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "experiment/experimentExecutor.proto",
}
