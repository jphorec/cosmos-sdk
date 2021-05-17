// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package types

import (
	context "context"
	types "github.com/cosmos/cosmos-sdk/types"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Evidence queries evidence based on evidence hash.
	Evidence(ctx context.Context, in *QueryEvidenceRequest, opts ...grpc.CallOption) (*QueryEvidenceResponse, error)
	// AllEvidence queries all evidence.
	AllEvidence(ctx context.Context, in *QueryAllEvidenceRequest, opts ...grpc.CallOption) (*QueryAllEvidenceResponse, error)
}

type queryClient struct {
	cc           grpc.ClientConnInterface
	_Evidence    types.Invoker
	_AllEvidence types.Invoker
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc: cc}
}

func (c *queryClient) Evidence(ctx context.Context, in *QueryEvidenceRequest, opts ...grpc.CallOption) (*QueryEvidenceResponse, error) {
	if invoker := c._Evidence; invoker != nil {
		var out QueryEvidenceResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Evidence, err = invokerConn.Invoker("/cosmos.evidence.v1beta1.QueryEvidence")
		if err != nil {
			var out QueryEvidenceResponse
			err = c._Evidence(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryEvidenceResponse)
	err := c.cc.Invoke(ctx, "/cosmos.evidence.v1beta1.QueryEvidence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllEvidence(ctx context.Context, in *QueryAllEvidenceRequest, opts ...grpc.CallOption) (*QueryAllEvidenceResponse, error) {
	if invoker := c._AllEvidence; invoker != nil {
		var out QueryAllEvidenceResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._AllEvidence, err = invokerConn.Invoker("/cosmos.evidence.v1beta1.QueryAllEvidence")
		if err != nil {
			var out QueryAllEvidenceResponse
			err = c._AllEvidence(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryAllEvidenceResponse)
	err := c.cc.Invoke(ctx, "/cosmos.evidence.v1beta1.QueryAllEvidence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Evidence queries evidence based on evidence hash.
	Evidence(context.Context, *QueryEvidenceRequest) (*QueryEvidenceResponse, error)
	// AllEvidence queries all evidence.
	AllEvidence(context.Context, *QueryAllEvidenceRequest) (*QueryAllEvidenceResponse, error)
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Evidence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEvidenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Evidence(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.evidence.v1beta1.QueryEvidence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Evidence(types.UnwrapSDKContext(ctx), req.(*QueryEvidenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllEvidence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllEvidenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllEvidence(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.evidence.v1beta1.QueryAllEvidence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllEvidence(types.UnwrapSDKContext(ctx), req.(*QueryAllEvidenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.evidence.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Evidence",
			Handler:    _Query_Evidence_Handler,
		},
		{
			MethodName: "AllEvidence",
			Handler:    _Query_AllEvidence_Handler,
		},
	},
	Metadata: "cosmos/evidence/v1beta1/query.proto",
}

const (
	QueryEvidenceMethod    = "/cosmos.evidence.v1beta1.QueryEvidence"
	QueryAllEvidenceMethod = "/cosmos.evidence.v1beta1.QueryAllEvidence"
)