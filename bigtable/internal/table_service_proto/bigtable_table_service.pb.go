// Code generated by protoc-gen-go.
// source: google.golang.org/cloud/bigtable/internal/table_service_proto/bigtable_table_service.proto
// DO NOT EDIT!

package google_bigtable_admin_table_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_bigtable_admin_table_v11 "google.golang.org/cloud/bigtable/internal/table_data_proto"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for BigtableTableService service

type BigtableTableServiceClient interface {
	// Creates a new table, to be served from a specified cluster.
	// The table can be created with a full set of initial column families,
	// specified in the request.
	CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*google_bigtable_admin_table_v11.Table, error)
	// Lists the names of all tables served from a specified cluster.
	ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error)
	// Gets the schema of the specified table, including its column families.
	GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*google_bigtable_admin_table_v11.Table, error)
	// Permanently deletes a specified table and all of its data.
	DeleteTable(ctx context.Context, in *DeleteTableRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// Changes the name of a specified table.
	// Cannot be used to move tables between clusters, zones, or projects.
	RenameTable(ctx context.Context, in *RenameTableRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// Creates a new column family within a specified table.
	CreateColumnFamily(ctx context.Context, in *CreateColumnFamilyRequest, opts ...grpc.CallOption) (*google_bigtable_admin_table_v11.ColumnFamily, error)
	// Changes the configuration of a specified column family.
	UpdateColumnFamily(ctx context.Context, in *google_bigtable_admin_table_v11.ColumnFamily, opts ...grpc.CallOption) (*google_bigtable_admin_table_v11.ColumnFamily, error)
	// Permanently deletes a specified column family and all of its data.
	DeleteColumnFamily(ctx context.Context, in *DeleteColumnFamilyRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// Delete all rows in a table corresponding to a particular prefix
	BulkDeleteRows(ctx context.Context, in *BulkDeleteRowsRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type bigtableTableServiceClient struct {
	cc *grpc.ClientConn
}

func NewBigtableTableServiceClient(cc *grpc.ClientConn) BigtableTableServiceClient {
	return &bigtableTableServiceClient{cc}
}

func (c *bigtableTableServiceClient) CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*google_bigtable_admin_table_v11.Table, error) {
	out := new(google_bigtable_admin_table_v11.Table)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/CreateTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error) {
	out := new(ListTablesResponse)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/ListTables", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*google_bigtable_admin_table_v11.Table, error) {
	out := new(google_bigtable_admin_table_v11.Table)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/GetTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) DeleteTable(ctx context.Context, in *DeleteTableRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/DeleteTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) RenameTable(ctx context.Context, in *RenameTableRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/RenameTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) CreateColumnFamily(ctx context.Context, in *CreateColumnFamilyRequest, opts ...grpc.CallOption) (*google_bigtable_admin_table_v11.ColumnFamily, error) {
	out := new(google_bigtable_admin_table_v11.ColumnFamily)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/CreateColumnFamily", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) UpdateColumnFamily(ctx context.Context, in *google_bigtable_admin_table_v11.ColumnFamily, opts ...grpc.CallOption) (*google_bigtable_admin_table_v11.ColumnFamily, error) {
	out := new(google_bigtable_admin_table_v11.ColumnFamily)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/UpdateColumnFamily", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) DeleteColumnFamily(ctx context.Context, in *DeleteColumnFamilyRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/DeleteColumnFamily", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) BulkDeleteRows(ctx context.Context, in *BulkDeleteRowsRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/BulkDeleteRows", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BigtableTableService service

type BigtableTableServiceServer interface {
	// Creates a new table, to be served from a specified cluster.
	// The table can be created with a full set of initial column families,
	// specified in the request.
	CreateTable(context.Context, *CreateTableRequest) (*google_bigtable_admin_table_v11.Table, error)
	// Lists the names of all tables served from a specified cluster.
	ListTables(context.Context, *ListTablesRequest) (*ListTablesResponse, error)
	// Gets the schema of the specified table, including its column families.
	GetTable(context.Context, *GetTableRequest) (*google_bigtable_admin_table_v11.Table, error)
	// Permanently deletes a specified table and all of its data.
	DeleteTable(context.Context, *DeleteTableRequest) (*google_protobuf1.Empty, error)
	// Changes the name of a specified table.
	// Cannot be used to move tables between clusters, zones, or projects.
	RenameTable(context.Context, *RenameTableRequest) (*google_protobuf1.Empty, error)
	// Creates a new column family within a specified table.
	CreateColumnFamily(context.Context, *CreateColumnFamilyRequest) (*google_bigtable_admin_table_v11.ColumnFamily, error)
	// Changes the configuration of a specified column family.
	UpdateColumnFamily(context.Context, *google_bigtable_admin_table_v11.ColumnFamily) (*google_bigtable_admin_table_v11.ColumnFamily, error)
	// Permanently deletes a specified column family and all of its data.
	DeleteColumnFamily(context.Context, *DeleteColumnFamilyRequest) (*google_protobuf1.Empty, error)
	// Delete all rows in a table corresponding to a particular prefix
	BulkDeleteRows(context.Context, *BulkDeleteRowsRequest) (*google_protobuf1.Empty, error)
}

func RegisterBigtableTableServiceServer(s *grpc.Server, srv BigtableTableServiceServer) {
	s.RegisterService(&_BigtableTableService_serviceDesc, srv)
}

func _BigtableTableService_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/CreateTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).CreateTable(ctx, req.(*CreateTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_ListTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).ListTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/ListTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).ListTables(ctx, req.(*ListTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_GetTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).GetTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/GetTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).GetTable(ctx, req.(*GetTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_DeleteTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).DeleteTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/DeleteTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).DeleteTable(ctx, req.(*DeleteTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_RenameTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).RenameTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/RenameTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).RenameTable(ctx, req.(*RenameTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_CreateColumnFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateColumnFamilyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).CreateColumnFamily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/CreateColumnFamily",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).CreateColumnFamily(ctx, req.(*CreateColumnFamilyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_UpdateColumnFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_bigtable_admin_table_v11.ColumnFamily)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).UpdateColumnFamily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/UpdateColumnFamily",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).UpdateColumnFamily(ctx, req.(*google_bigtable_admin_table_v11.ColumnFamily))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_DeleteColumnFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteColumnFamilyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).DeleteColumnFamily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/DeleteColumnFamily",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).DeleteColumnFamily(ctx, req.(*DeleteColumnFamilyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_BulkDeleteRows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkDeleteRowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).BulkDeleteRows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/BulkDeleteRows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).BulkDeleteRows(ctx, req.(*BulkDeleteRowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BigtableTableService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.bigtable.admin.table.v1.BigtableTableService",
	HandlerType: (*BigtableTableServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTable",
			Handler:    _BigtableTableService_CreateTable_Handler,
		},
		{
			MethodName: "ListTables",
			Handler:    _BigtableTableService_ListTables_Handler,
		},
		{
			MethodName: "GetTable",
			Handler:    _BigtableTableService_GetTable_Handler,
		},
		{
			MethodName: "DeleteTable",
			Handler:    _BigtableTableService_DeleteTable_Handler,
		},
		{
			MethodName: "RenameTable",
			Handler:    _BigtableTableService_RenameTable_Handler,
		},
		{
			MethodName: "CreateColumnFamily",
			Handler:    _BigtableTableService_CreateColumnFamily_Handler,
		},
		{
			MethodName: "UpdateColumnFamily",
			Handler:    _BigtableTableService_UpdateColumnFamily_Handler,
		},
		{
			MethodName: "DeleteColumnFamily",
			Handler:    _BigtableTableService_DeleteColumnFamily_Handler,
		},
		{
			MethodName: "BulkDeleteRows",
			Handler:    _BigtableTableService_BulkDeleteRows_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() {
	proto.RegisterFile("google.golang.org/cloud/bigtable/internal/table_service_proto/bigtable_table_service.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x94, 0x3f, 0x4f, 0xeb, 0x30,
	0x14, 0xc5, 0xdf, 0x5b, 0xde, 0x7b, 0x72, 0xa5, 0x37, 0x58, 0x88, 0x21, 0x48, 0x0c, 0x95, 0xd8,
	0x90, 0xa3, 0x16, 0x31, 0xb0, 0xa6, 0xfc, 0x59, 0x18, 0xaa, 0x52, 0x16, 0x18, 0x22, 0x27, 0xb9,
	0x58, 0x06, 0xff, 0x09, 0xb1, 0x53, 0xd4, 0x89, 0x2f, 0xca, 0x87, 0x81, 0xc4, 0x35, 0xa4, 0x50,
	0xe1, 0x66, 0x60, 0x89, 0x62, 0xdf, 0x73, 0xce, 0xcf, 0xbe, 0x37, 0x0a, 0xba, 0x61, 0x5a, 0x33,
	0x01, 0x84, 0x69, 0x41, 0x15, 0x23, 0xba, 0x62, 0x71, 0x2e, 0x74, 0x5d, 0xc4, 0x19, 0x67, 0x96,
	0x66, 0x02, 0x62, 0xae, 0x2c, 0x54, 0x8a, 0x8a, 0xb8, 0x5d, 0xa6, 0x06, 0xaa, 0x05, 0xcf, 0x21,
	0x2d, 0x2b, 0x6d, 0xf5, 0xbb, 0x2a, 0x5d, 0x2b, 0x92, 0xb6, 0x88, 0xf7, 0x57, 0xd9, 0x5e, 0x44,
	0x68, 0x21, 0xb9, 0x22, 0xee, 0x7d, 0x31, 0x8a, 0xe6, 0x7d, 0xd9, 0x05, 0xb5, 0x74, 0x33, 0xb8,
	0xa9, 0x38, 0x6a, 0x94, 0xff, 0xc4, 0x8d, 0x52, 0x09, 0xc6, 0x50, 0x06, 0x66, 0x05, 0xd9, 0x73,
	0x90, 0xb8, 0x5d, 0x65, 0xf5, 0x5d, 0x0c, 0xb2, 0xb4, 0x4b, 0x57, 0x1c, 0xbf, 0xfc, 0x45, 0x3b,
	0xc9, 0x2a, 0x66, 0xde, 0x3c, 0xae, 0x5c, 0x08, 0xbe, 0x47, 0x83, 0x49, 0x05, 0xd4, 0xba, 0x5d,
	0x3c, 0x26, 0xdf, 0x37, 0x88, 0x74, 0xc4, 0x33, 0x78, 0xac, 0xc1, 0xd8, 0xe8, 0x20, 0xe4, 0x69,
	0xd5, 0xc3, 0x5f, 0xb8, 0x46, 0xe8, 0x92, 0x1b, 0xdb, 0x2e, 0x0d, 0x1e, 0x85, 0x6c, 0x1f, 0x5a,
	0x4f, 0x1a, 0xf7, 0xb1, 0x98, 0x52, 0x2b, 0xd3, 0x60, 0x0b, 0xf4, 0xef, 0x02, 0xdc, 0x36, 0x8e,
	0x43, 0x09, 0x5e, 0xd9, 0xfb, 0x72, 0xb7, 0x68, 0x70, 0x0a, 0x02, 0xb6, 0x6e, 0x64, 0x47, 0xec,
	0x59, 0xbb, 0xde, 0xe3, 0x47, 0x48, 0xce, 0x9a, 0x11, 0xba, 0xf0, 0x19, 0x28, 0x2a, 0xb7, 0x0d,
	0xef, 0x88, 0xc3, 0xe1, 0xcf, 0x08, 0xbb, 0xa9, 0x4e, 0xb4, 0xa8, 0xa5, 0x3a, 0xa7, 0x92, 0x8b,
	0x25, 0x3e, 0xd9, 0xee, 0x4b, 0xe8, 0x7a, 0x3c, 0xea, 0x30, 0x68, 0xed, 0x98, 0xde, 0x0e, 0x50,
	0x21, 0x7c, 0x5d, 0x16, 0x9f, 0x0f, 0xd0, 0x2b, 0xa5, 0x37, 0x93, 0x23, 0xec, 0x26, 0xd0, 0xef,
	0xd2, 0x5f, 0x3d, 0xe1, 0xfe, 0x52, 0xf4, 0x3f, 0xa9, 0xc5, 0x83, 0xb3, 0xce, 0xf4, 0x93, 0xc1,
	0xc7, 0x21, 0xcc, 0xba, 0x3e, 0x88, 0x48, 0x12, 0x34, 0xcc, 0xb5, 0x0c, 0xa4, 0x26, 0xd1, 0xa6,
	0x3f, 0x80, 0x99, 0x36, 0x61, 0xd3, 0xdf, 0xd9, 0x9f, 0x36, 0xf5, 0xe8, 0x35, 0x00, 0x00, 0xff,
	0xff, 0x61, 0xcc, 0xfb, 0x30, 0x7f, 0x05, 0x00, 0x00,
}
