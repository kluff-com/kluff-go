// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: selling.proto

package selling

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

// SellingClient is the client API for Selling service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SellingClient interface {
	// demo call to check javascript sdk package fixme/removeme later
	GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error)
	GetMakeQuotation(ctx context.Context, in *MakeQuotationRequest, opts ...grpc.CallOption) (*MakeQuotationResponse, error)
	GetCustomerGroupDetails(ctx context.Context, in *CustomerGroupDetailRequest, opts ...grpc.CallOption) (*CustomerGroupDetailResponses, error)
	GetLoyaltyPrograms(ctx context.Context, in *GetLoyaltyProgramsRequest, opts ...grpc.CallOption) (*GetLoyaltyProgramsResponse, error)
	GetCustomerPrimaryContact(ctx context.Context, in *GetCustomerPrimaryContactRequest, opts ...grpc.CallOption) (*GetCustomerPrimaryContactResponse, error)
	SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error)
}

type sellingClient struct {
	cc grpc.ClientConnInterface
}

func NewSellingClient(cc grpc.ClientConnInterface) SellingClient {
	return &sellingClient{cc}
}

func (c *sellingClient) GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error) {
	out := new(Feature)
	err := c.cc.Invoke(ctx, "/selling.Selling/GetFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellingClient) GetMakeQuotation(ctx context.Context, in *MakeQuotationRequest, opts ...grpc.CallOption) (*MakeQuotationResponse, error) {
	out := new(MakeQuotationResponse)
	err := c.cc.Invoke(ctx, "/selling.Selling/GetMakeQuotation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellingClient) GetCustomerGroupDetails(ctx context.Context, in *CustomerGroupDetailRequest, opts ...grpc.CallOption) (*CustomerGroupDetailResponses, error) {
	out := new(CustomerGroupDetailResponses)
	err := c.cc.Invoke(ctx, "/selling.Selling/GetCustomerGroupDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellingClient) GetLoyaltyPrograms(ctx context.Context, in *GetLoyaltyProgramsRequest, opts ...grpc.CallOption) (*GetLoyaltyProgramsResponse, error) {
	out := new(GetLoyaltyProgramsResponse)
	err := c.cc.Invoke(ctx, "/selling.Selling/GetLoyaltyPrograms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellingClient) GetCustomerPrimaryContact(ctx context.Context, in *GetCustomerPrimaryContactRequest, opts ...grpc.CallOption) (*GetCustomerPrimaryContactResponse, error) {
	out := new(GetCustomerPrimaryContactResponse)
	err := c.cc.Invoke(ctx, "/selling.Selling/GetCustomerPrimaryContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellingClient) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error) {
	out := new(SendEmailResponse)
	err := c.cc.Invoke(ctx, "/selling.Selling/SendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SellingServer is the server API for Selling service.
// All implementations must embed UnimplementedSellingServer
// for forward compatibility
type SellingServer interface {
	// demo call to check javascript sdk package fixme/removeme later
	GetFeature(context.Context, *Point) (*Feature, error)
	GetMakeQuotation(context.Context, *MakeQuotationRequest) (*MakeQuotationResponse, error)
	GetCustomerGroupDetails(context.Context, *CustomerGroupDetailRequest) (*CustomerGroupDetailResponses, error)
	GetLoyaltyPrograms(context.Context, *GetLoyaltyProgramsRequest) (*GetLoyaltyProgramsResponse, error)
	GetCustomerPrimaryContact(context.Context, *GetCustomerPrimaryContactRequest) (*GetCustomerPrimaryContactResponse, error)
	SendEmail(context.Context, *SendEmailRequest) (*SendEmailResponse, error)
	mustEmbedUnimplementedSellingServer()
}

// UnimplementedSellingServer must be embedded to have forward compatible implementations.
type UnimplementedSellingServer struct {
}

func (UnimplementedSellingServer) GetFeature(context.Context, *Point) (*Feature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeature not implemented")
}
func (UnimplementedSellingServer) GetMakeQuotation(context.Context, *MakeQuotationRequest) (*MakeQuotationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMakeQuotation not implemented")
}
func (UnimplementedSellingServer) GetCustomerGroupDetails(context.Context, *CustomerGroupDetailRequest) (*CustomerGroupDetailResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerGroupDetails not implemented")
}
func (UnimplementedSellingServer) GetLoyaltyPrograms(context.Context, *GetLoyaltyProgramsRequest) (*GetLoyaltyProgramsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoyaltyPrograms not implemented")
}
func (UnimplementedSellingServer) GetCustomerPrimaryContact(context.Context, *GetCustomerPrimaryContactRequest) (*GetCustomerPrimaryContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerPrimaryContact not implemented")
}
func (UnimplementedSellingServer) SendEmail(context.Context, *SendEmailRequest) (*SendEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (UnimplementedSellingServer) mustEmbedUnimplementedSellingServer() {}

// UnsafeSellingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SellingServer will
// result in compilation errors.
type UnsafeSellingServer interface {
	mustEmbedUnimplementedSellingServer()
}

func RegisterSellingServer(s grpc.ServiceRegistrar, srv SellingServer) {
	s.RegisterService(&Selling_ServiceDesc, srv)
}

func _Selling_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellingServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selling.Selling/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellingServer).GetFeature(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _Selling_GetMakeQuotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeQuotationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellingServer).GetMakeQuotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selling.Selling/GetMakeQuotation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellingServer).GetMakeQuotation(ctx, req.(*MakeQuotationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Selling_GetCustomerGroupDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerGroupDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellingServer).GetCustomerGroupDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selling.Selling/GetCustomerGroupDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellingServer).GetCustomerGroupDetails(ctx, req.(*CustomerGroupDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Selling_GetLoyaltyPrograms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoyaltyProgramsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellingServer).GetLoyaltyPrograms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selling.Selling/GetLoyaltyPrograms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellingServer).GetLoyaltyPrograms(ctx, req.(*GetLoyaltyProgramsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Selling_GetCustomerPrimaryContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerPrimaryContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellingServer).GetCustomerPrimaryContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selling.Selling/GetCustomerPrimaryContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellingServer).GetCustomerPrimaryContact(ctx, req.(*GetCustomerPrimaryContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Selling_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellingServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selling.Selling/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellingServer).SendEmail(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Selling_ServiceDesc is the grpc.ServiceDesc for Selling service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Selling_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "selling.Selling",
	HandlerType: (*SellingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _Selling_GetFeature_Handler,
		},
		{
			MethodName: "GetMakeQuotation",
			Handler:    _Selling_GetMakeQuotation_Handler,
		},
		{
			MethodName: "GetCustomerGroupDetails",
			Handler:    _Selling_GetCustomerGroupDetails_Handler,
		},
		{
			MethodName: "GetLoyaltyPrograms",
			Handler:    _Selling_GetLoyaltyPrograms_Handler,
		},
		{
			MethodName: "GetCustomerPrimaryContact",
			Handler:    _Selling_GetCustomerPrimaryContact_Handler,
		},
		{
			MethodName: "SendEmail",
			Handler:    _Selling_SendEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "selling.proto",
}
