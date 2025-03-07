// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: staticanalysis/v1/staticanalysis.proto

package v1

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

const (
	StaticAnalysis_GetStaticDbFiles_FullMethodName       = "/staticanalysis.v1.StaticAnalysis/GetStaticDbFiles"
	StaticAnalysis_AnalyzeProjectPath_FullMethodName     = "/staticanalysis.v1.StaticAnalysis/AnalyzeProjectPath"
	StaticAnalysis_AnalyzeDbFile_FullMethodName          = "/staticanalysis.v1.StaticAnalysis/AnalyzeDbFile"
	StaticAnalysis_GetHotFunctions_FullMethodName        = "/staticanalysis.v1.StaticAnalysis/GetHotFunctions"
	StaticAnalysis_GetFunctionAnalysis_FullMethodName    = "/staticanalysis.v1.StaticAnalysis/GetFunctionAnalysis"
	StaticAnalysis_GetFunctionCallGraph_FullMethodName   = "/staticanalysis.v1.StaticAnalysis/GetFunctionCallGraph"
	StaticAnalysis_ListGitLabRepositories_FullMethodName = "/staticanalysis.v1.StaticAnalysis/ListGitLabRepositories"
	StaticAnalysis_CloneGitLabRepository_FullMethodName  = "/staticanalysis.v1.StaticAnalysis/CloneGitLabRepository"
)

// StaticAnalysisClient is the client API for StaticAnalysis service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaticAnalysisClient interface {
	// 获取静态分析数据库文件列表
	GetStaticDbFiles(ctx context.Context, in *GetStaticDbFilesRequest, opts ...grpc.CallOption) (*GetStaticDbFilesResponse, error)
	// 分析指定路径的项目并生成callgraph
	AnalyzeProjectPath(ctx context.Context, in *AnalyzeProjectPathRequest, opts ...grpc.CallOption) (*AnalyzeProjectPathResponse, error)
	// 分析指定数据库文件
	AnalyzeDbFile(ctx context.Context, in *AnalyzeDbFileRequest, opts ...grpc.CallOption) (*AnalyzeDbFileResponse, error)
	// 获取热点函数分析数据
	GetHotFunctions(ctx context.Context, in *GetHotFunctionsReq, opts ...grpc.CallOption) (*GetHotFunctionsReply, error)
	// 获取函数调用关系分析
	GetFunctionAnalysis(ctx context.Context, in *GetFunctionAnalysisReq, opts ...grpc.CallOption) (*GetFunctionAnalysisReply, error)
	// 获取函数调用关系图
	GetFunctionCallGraph(ctx context.Context, in *GetFunctionCallGraphReq, opts ...grpc.CallOption) (*GetFunctionCallGraphReply, error)
	// 获取GitLab仓库列表
	ListGitLabRepositories(ctx context.Context, in *ListGitLabRepositoriesRequest, opts ...grpc.CallOption) (*ListGitLabRepositoriesResponse, error)
	// 克隆GitLab仓库
	CloneGitLabRepository(ctx context.Context, in *CloneGitLabRepositoryRequest, opts ...grpc.CallOption) (*CloneGitLabRepositoryResponse, error)
}

type staticAnalysisClient struct {
	cc grpc.ClientConnInterface
}

func NewStaticAnalysisClient(cc grpc.ClientConnInterface) StaticAnalysisClient {
	return &staticAnalysisClient{cc}
}

func (c *staticAnalysisClient) GetStaticDbFiles(ctx context.Context, in *GetStaticDbFilesRequest, opts ...grpc.CallOption) (*GetStaticDbFilesResponse, error) {
	out := new(GetStaticDbFilesResponse)
	err := c.cc.Invoke(ctx, StaticAnalysis_GetStaticDbFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticAnalysisClient) AnalyzeProjectPath(ctx context.Context, in *AnalyzeProjectPathRequest, opts ...grpc.CallOption) (*AnalyzeProjectPathResponse, error) {
	out := new(AnalyzeProjectPathResponse)
	err := c.cc.Invoke(ctx, StaticAnalysis_AnalyzeProjectPath_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticAnalysisClient) AnalyzeDbFile(ctx context.Context, in *AnalyzeDbFileRequest, opts ...grpc.CallOption) (*AnalyzeDbFileResponse, error) {
	out := new(AnalyzeDbFileResponse)
	err := c.cc.Invoke(ctx, StaticAnalysis_AnalyzeDbFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticAnalysisClient) GetHotFunctions(ctx context.Context, in *GetHotFunctionsReq, opts ...grpc.CallOption) (*GetHotFunctionsReply, error) {
	out := new(GetHotFunctionsReply)
	err := c.cc.Invoke(ctx, StaticAnalysis_GetHotFunctions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticAnalysisClient) GetFunctionAnalysis(ctx context.Context, in *GetFunctionAnalysisReq, opts ...grpc.CallOption) (*GetFunctionAnalysisReply, error) {
	out := new(GetFunctionAnalysisReply)
	err := c.cc.Invoke(ctx, StaticAnalysis_GetFunctionAnalysis_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticAnalysisClient) GetFunctionCallGraph(ctx context.Context, in *GetFunctionCallGraphReq, opts ...grpc.CallOption) (*GetFunctionCallGraphReply, error) {
	out := new(GetFunctionCallGraphReply)
	err := c.cc.Invoke(ctx, StaticAnalysis_GetFunctionCallGraph_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticAnalysisClient) ListGitLabRepositories(ctx context.Context, in *ListGitLabRepositoriesRequest, opts ...grpc.CallOption) (*ListGitLabRepositoriesResponse, error) {
	out := new(ListGitLabRepositoriesResponse)
	err := c.cc.Invoke(ctx, StaticAnalysis_ListGitLabRepositories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticAnalysisClient) CloneGitLabRepository(ctx context.Context, in *CloneGitLabRepositoryRequest, opts ...grpc.CallOption) (*CloneGitLabRepositoryResponse, error) {
	out := new(CloneGitLabRepositoryResponse)
	err := c.cc.Invoke(ctx, StaticAnalysis_CloneGitLabRepository_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaticAnalysisServer is the server API for StaticAnalysis service.
// All implementations must embed UnimplementedStaticAnalysisServer
// for forward compatibility
type StaticAnalysisServer interface {
	// 获取静态分析数据库文件列表
	GetStaticDbFiles(context.Context, *GetStaticDbFilesRequest) (*GetStaticDbFilesResponse, error)
	// 分析指定路径的项目并生成callgraph
	AnalyzeProjectPath(context.Context, *AnalyzeProjectPathRequest) (*AnalyzeProjectPathResponse, error)
	// 分析指定数据库文件
	AnalyzeDbFile(context.Context, *AnalyzeDbFileRequest) (*AnalyzeDbFileResponse, error)
	// 获取热点函数分析数据
	GetHotFunctions(context.Context, *GetHotFunctionsReq) (*GetHotFunctionsReply, error)
	// 获取函数调用关系分析
	GetFunctionAnalysis(context.Context, *GetFunctionAnalysisReq) (*GetFunctionAnalysisReply, error)
	// 获取函数调用关系图
	GetFunctionCallGraph(context.Context, *GetFunctionCallGraphReq) (*GetFunctionCallGraphReply, error)
	// 获取GitLab仓库列表
	ListGitLabRepositories(context.Context, *ListGitLabRepositoriesRequest) (*ListGitLabRepositoriesResponse, error)
	// 克隆GitLab仓库
	CloneGitLabRepository(context.Context, *CloneGitLabRepositoryRequest) (*CloneGitLabRepositoryResponse, error)
	mustEmbedUnimplementedStaticAnalysisServer()
}

// UnimplementedStaticAnalysisServer must be embedded to have forward compatible implementations.
type UnimplementedStaticAnalysisServer struct {
}

func (UnimplementedStaticAnalysisServer) GetStaticDbFiles(context.Context, *GetStaticDbFilesRequest) (*GetStaticDbFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStaticDbFiles not implemented")
}
func (UnimplementedStaticAnalysisServer) AnalyzeProjectPath(context.Context, *AnalyzeProjectPathRequest) (*AnalyzeProjectPathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnalyzeProjectPath not implemented")
}
func (UnimplementedStaticAnalysisServer) AnalyzeDbFile(context.Context, *AnalyzeDbFileRequest) (*AnalyzeDbFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnalyzeDbFile not implemented")
}
func (UnimplementedStaticAnalysisServer) GetHotFunctions(context.Context, *GetHotFunctionsReq) (*GetHotFunctionsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHotFunctions not implemented")
}
func (UnimplementedStaticAnalysisServer) GetFunctionAnalysis(context.Context, *GetFunctionAnalysisReq) (*GetFunctionAnalysisReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFunctionAnalysis not implemented")
}
func (UnimplementedStaticAnalysisServer) GetFunctionCallGraph(context.Context, *GetFunctionCallGraphReq) (*GetFunctionCallGraphReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFunctionCallGraph not implemented")
}
func (UnimplementedStaticAnalysisServer) ListGitLabRepositories(context.Context, *ListGitLabRepositoriesRequest) (*ListGitLabRepositoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGitLabRepositories not implemented")
}
func (UnimplementedStaticAnalysisServer) CloneGitLabRepository(context.Context, *CloneGitLabRepositoryRequest) (*CloneGitLabRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloneGitLabRepository not implemented")
}
func (UnimplementedStaticAnalysisServer) mustEmbedUnimplementedStaticAnalysisServer() {}

// UnsafeStaticAnalysisServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaticAnalysisServer will
// result in compilation errors.
type UnsafeStaticAnalysisServer interface {
	mustEmbedUnimplementedStaticAnalysisServer()
}

func RegisterStaticAnalysisServer(s grpc.ServiceRegistrar, srv StaticAnalysisServer) {
	s.RegisterService(&StaticAnalysis_ServiceDesc, srv)
}

func _StaticAnalysis_GetStaticDbFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStaticDbFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticAnalysisServer).GetStaticDbFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticAnalysis_GetStaticDbFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticAnalysisServer).GetStaticDbFiles(ctx, req.(*GetStaticDbFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticAnalysis_AnalyzeProjectPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzeProjectPathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticAnalysisServer).AnalyzeProjectPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticAnalysis_AnalyzeProjectPath_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticAnalysisServer).AnalyzeProjectPath(ctx, req.(*AnalyzeProjectPathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticAnalysis_AnalyzeDbFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzeDbFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticAnalysisServer).AnalyzeDbFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticAnalysis_AnalyzeDbFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticAnalysisServer).AnalyzeDbFile(ctx, req.(*AnalyzeDbFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticAnalysis_GetHotFunctions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotFunctionsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticAnalysisServer).GetHotFunctions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticAnalysis_GetHotFunctions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticAnalysisServer).GetHotFunctions(ctx, req.(*GetHotFunctionsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticAnalysis_GetFunctionAnalysis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFunctionAnalysisReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticAnalysisServer).GetFunctionAnalysis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticAnalysis_GetFunctionAnalysis_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticAnalysisServer).GetFunctionAnalysis(ctx, req.(*GetFunctionAnalysisReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticAnalysis_GetFunctionCallGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFunctionCallGraphReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticAnalysisServer).GetFunctionCallGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticAnalysis_GetFunctionCallGraph_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticAnalysisServer).GetFunctionCallGraph(ctx, req.(*GetFunctionCallGraphReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticAnalysis_ListGitLabRepositories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGitLabRepositoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticAnalysisServer).ListGitLabRepositories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticAnalysis_ListGitLabRepositories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticAnalysisServer).ListGitLabRepositories(ctx, req.(*ListGitLabRepositoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticAnalysis_CloneGitLabRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneGitLabRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticAnalysisServer).CloneGitLabRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticAnalysis_CloneGitLabRepository_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticAnalysisServer).CloneGitLabRepository(ctx, req.(*CloneGitLabRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StaticAnalysis_ServiceDesc is the grpc.ServiceDesc for StaticAnalysis service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StaticAnalysis_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "staticanalysis.v1.StaticAnalysis",
	HandlerType: (*StaticAnalysisServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStaticDbFiles",
			Handler:    _StaticAnalysis_GetStaticDbFiles_Handler,
		},
		{
			MethodName: "AnalyzeProjectPath",
			Handler:    _StaticAnalysis_AnalyzeProjectPath_Handler,
		},
		{
			MethodName: "AnalyzeDbFile",
			Handler:    _StaticAnalysis_AnalyzeDbFile_Handler,
		},
		{
			MethodName: "GetHotFunctions",
			Handler:    _StaticAnalysis_GetHotFunctions_Handler,
		},
		{
			MethodName: "GetFunctionAnalysis",
			Handler:    _StaticAnalysis_GetFunctionAnalysis_Handler,
		},
		{
			MethodName: "GetFunctionCallGraph",
			Handler:    _StaticAnalysis_GetFunctionCallGraph_Handler,
		},
		{
			MethodName: "ListGitLabRepositories",
			Handler:    _StaticAnalysis_ListGitLabRepositories_Handler,
		},
		{
			MethodName: "CloneGitLabRepository",
			Handler:    _StaticAnalysis_CloneGitLabRepository_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staticanalysis/v1/staticanalysis.proto",
}
