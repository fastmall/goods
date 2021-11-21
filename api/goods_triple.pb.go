// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.5
// - protoc             v3.8.0
// source: goods.proto

package api

import (
	context "context"
	protocol "dubbo.apache.org/dubbo-go/v3/protocol"
	dubbo3 "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	invocation "dubbo.apache.org/dubbo-go/v3/protocol/invocation"
	grpc_go "github.com/dubbogo/grpc-go"
	codes "github.com/dubbogo/grpc-go/codes"
	metadata "github.com/dubbogo/grpc-go/metadata"
	status "github.com/dubbogo/grpc-go/status"
	common "github.com/dubbogo/triple/pkg/common"
	constant "github.com/dubbogo/triple/pkg/common/constant"
	triple "github.com/dubbogo/triple/pkg/triple"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc_go.SupportPackageIsVersion7

// GoodsServiceClient is the client API for GoodsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsServiceClient interface {
	CreateGoodsItem(ctx context.Context, in *CreateGoodsItemRequest, opts ...grpc_go.CallOption) (*CreateGoodsItemResponse, common.ErrorWithAttachment)
}

type goodsServiceClient struct {
	cc *triple.TripleConn
}

type GoodsServiceClientImpl struct {
	CreateGoodsItem func(ctx context.Context, in *CreateGoodsItemRequest) (*CreateGoodsItemResponse, error)
}

func (c *GoodsServiceClientImpl) GetDubboStub(cc *triple.TripleConn) GoodsServiceClient {
	return NewGoodsServiceClient(cc)
}

func (c *GoodsServiceClientImpl) XXX_InterfaceName() string {
	return "com.jeongen.fastmall.goods.GoodsService"
}

func NewGoodsServiceClient(cc *triple.TripleConn) GoodsServiceClient {
	return &goodsServiceClient{cc}
}

func (c *goodsServiceClient) CreateGoodsItem(ctx context.Context, in *CreateGoodsItemRequest, opts ...grpc_go.CallOption) (*CreateGoodsItemResponse, common.ErrorWithAttachment) {
	out := new(CreateGoodsItemResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/CreateGoodsItem", in, out)
}

// GoodsServiceServer is the server API for GoodsService service.
// All implementations must embed UnimplementedGoodsServiceServer
// for forward compatibility
type GoodsServiceServer interface {
	CreateGoodsItem(context.Context, *CreateGoodsItemRequest) (*CreateGoodsItemResponse, error)
	mustEmbedUnimplementedGoodsServiceServer()
}

// UnimplementedGoodsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsServiceServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedGoodsServiceServer) CreateGoodsItem(context.Context, *CreateGoodsItemRequest) (*CreateGoodsItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGoodsItem not implemented")
}
func (s *UnimplementedGoodsServiceServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedGoodsServiceServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedGoodsServiceServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &GoodsService_ServiceDesc
}
func (s *UnimplementedGoodsServiceServer) XXX_InterfaceName() string {
	return "com.jeongen.fastmall.goods.GoodsService"
}

func (UnimplementedGoodsServiceServer) mustEmbedUnimplementedGoodsServiceServer() {}

// UnsafeGoodsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServiceServer will
// result in compilation errors.
type UnsafeGoodsServiceServer interface {
	mustEmbedUnimplementedGoodsServiceServer()
}

func RegisterGoodsServiceServer(s grpc_go.ServiceRegistrar, srv GoodsServiceServer) {
	s.RegisterService(&GoodsService_ServiceDesc, srv)
}

func _GoodsService_CreateGoodsItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodsItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("CreateGoodsItem", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

// GoodsService_ServiceDesc is the grpc_go.ServiceDesc for GoodsService service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoodsService_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "com.jeongen.fastmall.goods.GoodsService",
	HandlerType: (*GoodsServiceServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "CreateGoodsItem",
			Handler:    _GoodsService_CreateGoodsItem_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "goods.proto",
}
