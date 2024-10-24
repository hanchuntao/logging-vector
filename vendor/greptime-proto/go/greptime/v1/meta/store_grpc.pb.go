// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: greptime/v1/meta/store.proto

package meta

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

// StoreClient is the client API for Store service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreClient interface {
	// Range gets the keys in the range from the key-value store.
	Range(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*RangeResponse, error)
	// Put puts the given key into the key-value store.
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	// BatchGet atomically get values by the given keys from the key-value store.
	BatchGet(ctx context.Context, in *BatchGetRequest, opts ...grpc.CallOption) (*BatchGetResponse, error)
	// BatchPut atomically puts the given keys into the key-value store.
	BatchPut(ctx context.Context, in *BatchPutRequest, opts ...grpc.CallOption) (*BatchPutResponse, error)
	// BatchDelete atomically deletes the given keys and its associating values
	// from the key-value store.
	BatchDelete(ctx context.Context, in *BatchDeleteRequest, opts ...grpc.CallOption) (*BatchDeleteResponse, error)
	// CompareAndPut atomically puts the value to the given updated
	// value if the current value == the expected value.
	CompareAndPut(ctx context.Context, in *CompareAndPutRequest, opts ...grpc.CallOption) (*CompareAndPutResponse, error)
	// DeleteRange deletes the given range from the key-value store.
	DeleteRange(ctx context.Context, in *DeleteRangeRequest, opts ...grpc.CallOption) (*DeleteRangeResponse, error)
	// MoveValue atomically renames the key to the given updated key.
	MoveValue(ctx context.Context, in *MoveValueRequest, opts ...grpc.CallOption) (*MoveValueResponse, error)
}

type storeClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreClient(cc grpc.ClientConnInterface) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) Range(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*RangeResponse, error) {
	out := new(RangeResponse)
	err := c.cc.Invoke(ctx, "/greptime.v1.meta.Store/Range", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/greptime.v1.meta.Store/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) BatchGet(ctx context.Context, in *BatchGetRequest, opts ...grpc.CallOption) (*BatchGetResponse, error) {
	out := new(BatchGetResponse)
	err := c.cc.Invoke(ctx, "/greptime.v1.meta.Store/BatchGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) BatchPut(ctx context.Context, in *BatchPutRequest, opts ...grpc.CallOption) (*BatchPutResponse, error) {
	out := new(BatchPutResponse)
	err := c.cc.Invoke(ctx, "/greptime.v1.meta.Store/BatchPut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) BatchDelete(ctx context.Context, in *BatchDeleteRequest, opts ...grpc.CallOption) (*BatchDeleteResponse, error) {
	out := new(BatchDeleteResponse)
	err := c.cc.Invoke(ctx, "/greptime.v1.meta.Store/BatchDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) CompareAndPut(ctx context.Context, in *CompareAndPutRequest, opts ...grpc.CallOption) (*CompareAndPutResponse, error) {
	out := new(CompareAndPutResponse)
	err := c.cc.Invoke(ctx, "/greptime.v1.meta.Store/CompareAndPut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) DeleteRange(ctx context.Context, in *DeleteRangeRequest, opts ...grpc.CallOption) (*DeleteRangeResponse, error) {
	out := new(DeleteRangeResponse)
	err := c.cc.Invoke(ctx, "/greptime.v1.meta.Store/DeleteRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) MoveValue(ctx context.Context, in *MoveValueRequest, opts ...grpc.CallOption) (*MoveValueResponse, error) {
	out := new(MoveValueResponse)
	err := c.cc.Invoke(ctx, "/greptime.v1.meta.Store/MoveValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
// All implementations must embed UnimplementedStoreServer
// for forward compatibility
type StoreServer interface {
	// Range gets the keys in the range from the key-value store.
	Range(context.Context, *RangeRequest) (*RangeResponse, error)
	// Put puts the given key into the key-value store.
	Put(context.Context, *PutRequest) (*PutResponse, error)
	// BatchGet atomically get values by the given keys from the key-value store.
	BatchGet(context.Context, *BatchGetRequest) (*BatchGetResponse, error)
	// BatchPut atomically puts the given keys into the key-value store.
	BatchPut(context.Context, *BatchPutRequest) (*BatchPutResponse, error)
	// BatchDelete atomically deletes the given keys and its associating values
	// from the key-value store.
	BatchDelete(context.Context, *BatchDeleteRequest) (*BatchDeleteResponse, error)
	// CompareAndPut atomically puts the value to the given updated
	// value if the current value == the expected value.
	CompareAndPut(context.Context, *CompareAndPutRequest) (*CompareAndPutResponse, error)
	// DeleteRange deletes the given range from the key-value store.
	DeleteRange(context.Context, *DeleteRangeRequest) (*DeleteRangeResponse, error)
	// MoveValue atomically renames the key to the given updated key.
	MoveValue(context.Context, *MoveValueRequest) (*MoveValueResponse, error)
	mustEmbedUnimplementedStoreServer()
}

// UnimplementedStoreServer must be embedded to have forward compatible implementations.
type UnimplementedStoreServer struct {
}

func (UnimplementedStoreServer) Range(context.Context, *RangeRequest) (*RangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Range not implemented")
}
func (UnimplementedStoreServer) Put(context.Context, *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedStoreServer) BatchGet(context.Context, *BatchGetRequest) (*BatchGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGet not implemented")
}
func (UnimplementedStoreServer) BatchPut(context.Context, *BatchPutRequest) (*BatchPutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchPut not implemented")
}
func (UnimplementedStoreServer) BatchDelete(context.Context, *BatchDeleteRequest) (*BatchDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDelete not implemented")
}
func (UnimplementedStoreServer) CompareAndPut(context.Context, *CompareAndPutRequest) (*CompareAndPutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompareAndPut not implemented")
}
func (UnimplementedStoreServer) DeleteRange(context.Context, *DeleteRangeRequest) (*DeleteRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRange not implemented")
}
func (UnimplementedStoreServer) MoveValue(context.Context, *MoveValueRequest) (*MoveValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveValue not implemented")
}
func (UnimplementedStoreServer) mustEmbedUnimplementedStoreServer() {}

// UnsafeStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreServer will
// result in compilation errors.
type UnsafeStoreServer interface {
	mustEmbedUnimplementedStoreServer()
}

func RegisterStoreServer(s grpc.ServiceRegistrar, srv StoreServer) {
	s.RegisterService(&Store_ServiceDesc, srv)
}

func _Store_Range_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Range(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greptime.v1.meta.Store/Range",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Range(ctx, req.(*RangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greptime.v1.meta.Store/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_BatchGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).BatchGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greptime.v1.meta.Store/BatchGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).BatchGet(ctx, req.(*BatchGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_BatchPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).BatchPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greptime.v1.meta.Store/BatchPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).BatchPut(ctx, req.(*BatchPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_BatchDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).BatchDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greptime.v1.meta.Store/BatchDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).BatchDelete(ctx, req.(*BatchDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_CompareAndPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompareAndPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).CompareAndPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greptime.v1.meta.Store/CompareAndPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).CompareAndPut(ctx, req.(*CompareAndPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_DeleteRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).DeleteRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greptime.v1.meta.Store/DeleteRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).DeleteRange(ctx, req.(*DeleteRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_MoveValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).MoveValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greptime.v1.meta.Store/MoveValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).MoveValue(ctx, req.(*MoveValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Store_ServiceDesc is the grpc.ServiceDesc for Store service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Store_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greptime.v1.meta.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Range",
			Handler:    _Store_Range_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Store_Put_Handler,
		},
		{
			MethodName: "BatchGet",
			Handler:    _Store_BatchGet_Handler,
		},
		{
			MethodName: "BatchPut",
			Handler:    _Store_BatchPut_Handler,
		},
		{
			MethodName: "BatchDelete",
			Handler:    _Store_BatchDelete_Handler,
		},
		{
			MethodName: "CompareAndPut",
			Handler:    _Store_CompareAndPut_Handler,
		},
		{
			MethodName: "DeleteRange",
			Handler:    _Store_DeleteRange_Handler,
		},
		{
			MethodName: "MoveValue",
			Handler:    _Store_MoveValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greptime/v1/meta/store.proto",
}
