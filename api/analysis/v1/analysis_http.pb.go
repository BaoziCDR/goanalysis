// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.26.1
// source: analysis/v1/analysis.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAnalysisAnalyzeDbFile = "/analysis.v1.Analysis/AnalyzeDbFile"
const OperationAnalysisAnalyzeProjectPath = "/analysis.v1.Analysis/AnalyzeProjectPath"
const OperationAnalysisCheckDatabase = "/analysis.v1.Analysis/CheckDatabase"
const OperationAnalysisGenerateImage = "/analysis.v1.Analysis/GenerateImage"
const OperationAnalysisGetAllFunctionName = "/analysis.v1.Analysis/GetAllFunctionName"
const OperationAnalysisGetAllGIDs = "/analysis.v1.Analysis/GetAllGIDs"
const OperationAnalysisGetAllParentFuncNames = "/analysis.v1.Analysis/GetAllParentFuncNames"
const OperationAnalysisGetAnalysis = "/analysis.v1.Analysis/GetAnalysis"
const OperationAnalysisGetAnalysisByGID = "/analysis.v1.Analysis/GetAnalysisByGID"
const OperationAnalysisGetChildFunctions = "/analysis.v1.Analysis/GetChildFunctions"
const OperationAnalysisGetFunctionAnalysis = "/analysis.v1.Analysis/GetFunctionAnalysis"
const OperationAnalysisGetFunctionCallGraph = "/analysis.v1.Analysis/GetFunctionCallGraph"
const OperationAnalysisGetGidsByFunctionName = "/analysis.v1.Analysis/GetGidsByFunctionName"
const OperationAnalysisGetGoroutineStats = "/analysis.v1.Analysis/GetGoroutineStats"
const OperationAnalysisGetHotFunctions = "/analysis.v1.Analysis/GetHotFunctions"
const OperationAnalysisGetParamsByID = "/analysis.v1.Analysis/GetParamsByID"
const OperationAnalysisGetStaticDbFiles = "/analysis.v1.Analysis/GetStaticDbFiles"
const OperationAnalysisGetTraceGraph = "/analysis.v1.Analysis/GetTraceGraph"
const OperationAnalysisGetTracesByParentFunc = "/analysis.v1.Analysis/GetTracesByParentFunc"
const OperationAnalysisVerifyProjectPath = "/analysis.v1.Analysis/VerifyProjectPath"

type AnalysisHTTPServer interface {
	// AnalyzeDbFile 分析指定数据库文件
	AnalyzeDbFile(context.Context, *AnalyzeDbFileRequest) (*AnalyzeDbFileResponse, error)
	// AnalyzeProjectPath 分析指定路径的项目并生成callgraph
	AnalyzeProjectPath(context.Context, *AnalyzeProjectPathRequest) (*AnalyzeProjectPathResponse, error)
	// CheckDatabase CheckDatabase checks if the trace database exists
	CheckDatabase(context.Context, *CheckDatabaseRequest) (*CheckDatabaseResponse, error)
	GenerateImage(context.Context, *GenerateImageReq) (*GenerateImageReply, error)
	GetAllFunctionName(context.Context, *GetAllFunctionNameReq) (*GetAllFunctionNameReply, error)
	GetAllGIDs(context.Context, *GetAllGIDsReq) (*GetAllGIDsReply, error)
	// GetAllParentFuncNames GetAllParentFuncNames 获取所有的父函数名称
	GetAllParentFuncNames(context.Context, *GetAllParentFuncNamesReq) (*GetAllParentFuncNamesReply, error)
	// GetAnalysis Sends a greeting
	GetAnalysis(context.Context, *AnalysisRequest) (*AnalysisReply, error)
	GetAnalysisByGID(context.Context, *AnalysisByGIDRequest) (*AnalysisByGIDReply, error)
	// GetChildFunctions GetChildFunctions 获取函数的子函数
	GetChildFunctions(context.Context, *GetChildFunctionsReq) (*GetChildFunctionsReply, error)
	// GetFunctionAnalysis GetFunctionAnalysis 获取函数调用关系分析
	GetFunctionAnalysis(context.Context, *GetFunctionAnalysisReq) (*GetFunctionAnalysisReply, error)
	// GetFunctionCallGraph GetFunctionCallGraph 获取函数调用关系图
	GetFunctionCallGraph(context.Context, *GetFunctionCallGraphReq) (*GetFunctionCallGraphReply, error)
	GetGidsByFunctionName(context.Context, *GetGidsByFunctionNameReq) (*GetGidsByFunctionNameReply, error)
	// GetGoroutineStats GetGoroutineStats 获取Goroutine统计信息
	GetGoroutineStats(context.Context, *GetGoroutineStatsReq) (*GetGoroutineStatsReply, error)
	// GetHotFunctions GetHotFunctions 获取热点函数分析数据
	GetHotFunctions(context.Context, *GetHotFunctionsReq) (*GetHotFunctionsReply, error)
	GetParamsByID(context.Context, *GetParamsByIDReq) (*GetParamsByIDReply, error)
	// GetStaticDbFiles 获取静态分析数据库文件列表
	GetStaticDbFiles(context.Context, *GetStaticDbFilesRequest) (*GetStaticDbFilesResponse, error)
	GetTraceGraph(context.Context, *GetTraceGraphReq) (*GetTraceGraphReply, error)
	// GetTracesByParentFunc GetTracesByParentFunc 根据父函数名称获取函数调用
	GetTracesByParentFunc(context.Context, *GetTracesByParentFuncReq) (*GetTracesByParentFuncReply, error)
	VerifyProjectPath(context.Context, *VerifyProjectPathReq) (*VerifyProjectPathReply, error)
}

func RegisterAnalysisHTTPServer(s *http.Server, srv AnalysisHTTPServer) {
	r := s.Route("/")
	r.GET("/analysis/{name}", _Analysis_GetAnalysis0_HTTP_Handler(srv))
	r.GET("/api/traces/{gid}", _Analysis_GetAnalysisByGID0_HTTP_Handler(srv))
	r.GET("/api/gids", _Analysis_GetAllGIDs0_HTTP_Handler(srv))
	r.GET("/api/params/{id}", _Analysis_GetParamsByID0_HTTP_Handler(srv))
	r.GET("/api/traces/{gid}/mermaid", _Analysis_GenerateImage0_HTTP_Handler(srv))
	r.GET("/api/functions", _Analysis_GetAllFunctionName0_HTTP_Handler(srv))
	r.POST("/api/gids/function", _Analysis_GetGidsByFunctionName0_HTTP_Handler(srv))
	r.POST("/api/verify/path", _Analysis_VerifyProjectPath0_HTTP_Handler(srv))
	r.GET("/api/check/db", _Analysis_CheckDatabase0_HTTP_Handler(srv))
	r.GET("/api/static/dbfiles", _Analysis_GetStaticDbFiles0_HTTP_Handler(srv))
	r.POST("/api/static/analyze/path", _Analysis_AnalyzeProjectPath0_HTTP_Handler(srv))
	r.POST("/api/static/analyze", _Analysis_AnalyzeDbFile0_HTTP_Handler(srv))
	r.GET("/api/traces/{gid}/graph", _Analysis_GetTraceGraph0_HTTP_Handler(srv))
	r.GET("/api/traces/parent/{parent_func}", _Analysis_GetTracesByParentFunc0_HTTP_Handler(srv))
	r.GET("/api/functions/parent", _Analysis_GetAllParentFuncNames0_HTTP_Handler(srv))
	r.GET("/api/functions/{func_name}/children", _Analysis_GetChildFunctions0_HTTP_Handler(srv))
	r.GET("/api/hot-functions", _Analysis_GetHotFunctions0_HTTP_Handler(srv))
	r.GET("/api/goroutine-stats", _Analysis_GetGoroutineStats0_HTTP_Handler(srv))
	r.POST("/api/function/analysis", _Analysis_GetFunctionAnalysis0_HTTP_Handler(srv))
	r.GET("/api/function/graph", _Analysis_GetFunctionCallGraph0_HTTP_Handler(srv))
	r.GET("/api/function/{function_name}/graph", _Analysis_GetFunctionCallGraph1_HTTP_Handler(srv))
}

func _Analysis_GetAnalysis0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AnalysisRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisGetAnalysis)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAnalysis(ctx, req.(*AnalysisRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AnalysisReply)
		return ctx.Result(200, reply)
	}
}

func _Analysis_GetAnalysisByGID0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AnalysisByGIDRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisGetAnalysisByGID)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAnalysisByGID(ctx, req.(*AnalysisByGIDRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AnalysisByGIDReply)
		return ctx.Result(200, reply)
	}
}

func _Analysis_GetAllGIDs0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAllGIDsReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisGetAllGIDs)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAllGIDs(ctx, req.(*GetAllGIDsReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAllGIDsReply)
		return ctx.Result(200, reply)
	}
}

func _Analysis_GetParamsByID0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetParamsByIDReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisGetParamsByID)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetParamsByID(ctx, req.(*GetParamsByIDReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetParamsByIDReply)
		return ctx.Result(200, reply)
	}
}

func _Analysis_GenerateImage0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GenerateImageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisGenerateImage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GenerateImage(ctx, req.(*GenerateImageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GenerateImageReply)
		return ctx.Result(200, reply)
	}
}

func _Analysis_GetAllFunctionName0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAllFunctionNameReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisGetAllFunctionName)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAllFunctionName(ctx, req.(*GetAllFunctionNameReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAllFunctionNameReply)
		return ctx.Result(200, reply)
	}
}

func _Analysis_GetGidsByFunctionName0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetGidsByFunctionNameReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisGetGidsByFunctionName)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetGidsByFunctionName(ctx, req.(*GetGidsByFunctionNameReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetGidsByFunctionNameReply)
		return ctx.Result(200, reply)
	}
}

func _Analysis_VerifyProjectPath0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in VerifyProjectPathReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisVerifyProjectPath)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.VerifyProjectPath(ctx, req.(*VerifyProjectPathReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VerifyProjectPathReply)
		return ctx.Result(200, reply)
	}
}

func _Analysis_CheckDatabase0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CheckDatabaseRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisCheckDatabase)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CheckDatabase(ctx, req.(*CheckDatabaseRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CheckDatabaseResponse)
		return ctx.Result(200, reply)
	}
}

func _Analysis_GetStaticDbFiles0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetStaticDbFilesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisGetStaticDbFiles)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetStaticDbFiles(ctx, req.(*GetStaticDbFilesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetStaticDbFilesResponse)
		return ctx.Result(200, reply)
	}
}

func _Analysis_AnalyzeProjectPath0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AnalyzeProjectPathRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisAnalyzeProjectPath)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AnalyzeProjectPath(ctx, req.(*AnalyzeProjectPathRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AnalyzeProjectPathResponse)
		return ctx.Result(200, reply)
	}
}

func _Analysis_AnalyzeDbFile0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AnalyzeDbFileRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisAnalyzeDbFile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AnalyzeDbFile(ctx, req.(*AnalyzeDbFileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AnalyzeDbFileResponse)
		return ctx.Result(200, reply)
	}
}

func _Analysis_GetTraceGraph0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTraceGraphReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisGetTraceGraph)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTraceGraph(ctx, req.(*GetTraceGraphReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTraceGraphReply)
		return ctx.Result(200, reply)
	}
}

func _Analysis_GetTracesByParentFunc0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTracesByParentFuncReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisGetTracesByParentFunc)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTracesByParentFunc(ctx, req.(*GetTracesByParentFuncReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTracesByParentFuncReply)
		return ctx.Result(200, reply)
	}
}

func _Analysis_GetAllParentFuncNames0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAllParentFuncNamesReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisGetAllParentFuncNames)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAllParentFuncNames(ctx, req.(*GetAllParentFuncNamesReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAllParentFuncNamesReply)
		return ctx.Result(200, reply)
	}
}

func _Analysis_GetChildFunctions0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetChildFunctionsReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisGetChildFunctions)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetChildFunctions(ctx, req.(*GetChildFunctionsReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetChildFunctionsReply)
		return ctx.Result(200, reply)
	}
}

func _Analysis_GetHotFunctions0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetHotFunctionsReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisGetHotFunctions)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetHotFunctions(ctx, req.(*GetHotFunctionsReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetHotFunctionsReply)
		return ctx.Result(200, reply)
	}
}

func _Analysis_GetGoroutineStats0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetGoroutineStatsReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisGetGoroutineStats)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetGoroutineStats(ctx, req.(*GetGoroutineStatsReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetGoroutineStatsReply)
		return ctx.Result(200, reply)
	}
}

func _Analysis_GetFunctionAnalysis0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetFunctionAnalysisReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisGetFunctionAnalysis)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFunctionAnalysis(ctx, req.(*GetFunctionAnalysisReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetFunctionAnalysisReply)
		return ctx.Result(200, reply)
	}
}

func _Analysis_GetFunctionCallGraph0_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetFunctionCallGraphReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisGetFunctionCallGraph)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFunctionCallGraph(ctx, req.(*GetFunctionCallGraphReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetFunctionCallGraphReply)
		return ctx.Result(200, reply)
	}
}

func _Analysis_GetFunctionCallGraph1_HTTP_Handler(srv AnalysisHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetFunctionCallGraphReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAnalysisGetFunctionCallGraph)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFunctionCallGraph(ctx, req.(*GetFunctionCallGraphReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetFunctionCallGraphReply)
		return ctx.Result(200, reply)
	}
}

type AnalysisHTTPClient interface {
	AnalyzeDbFile(ctx context.Context, req *AnalyzeDbFileRequest, opts ...http.CallOption) (rsp *AnalyzeDbFileResponse, err error)
	AnalyzeProjectPath(ctx context.Context, req *AnalyzeProjectPathRequest, opts ...http.CallOption) (rsp *AnalyzeProjectPathResponse, err error)
	CheckDatabase(ctx context.Context, req *CheckDatabaseRequest, opts ...http.CallOption) (rsp *CheckDatabaseResponse, err error)
	GenerateImage(ctx context.Context, req *GenerateImageReq, opts ...http.CallOption) (rsp *GenerateImageReply, err error)
	GetAllFunctionName(ctx context.Context, req *GetAllFunctionNameReq, opts ...http.CallOption) (rsp *GetAllFunctionNameReply, err error)
	GetAllGIDs(ctx context.Context, req *GetAllGIDsReq, opts ...http.CallOption) (rsp *GetAllGIDsReply, err error)
	GetAllParentFuncNames(ctx context.Context, req *GetAllParentFuncNamesReq, opts ...http.CallOption) (rsp *GetAllParentFuncNamesReply, err error)
	GetAnalysis(ctx context.Context, req *AnalysisRequest, opts ...http.CallOption) (rsp *AnalysisReply, err error)
	GetAnalysisByGID(ctx context.Context, req *AnalysisByGIDRequest, opts ...http.CallOption) (rsp *AnalysisByGIDReply, err error)
	GetChildFunctions(ctx context.Context, req *GetChildFunctionsReq, opts ...http.CallOption) (rsp *GetChildFunctionsReply, err error)
	GetFunctionAnalysis(ctx context.Context, req *GetFunctionAnalysisReq, opts ...http.CallOption) (rsp *GetFunctionAnalysisReply, err error)
	GetFunctionCallGraph(ctx context.Context, req *GetFunctionCallGraphReq, opts ...http.CallOption) (rsp *GetFunctionCallGraphReply, err error)
	GetGidsByFunctionName(ctx context.Context, req *GetGidsByFunctionNameReq, opts ...http.CallOption) (rsp *GetGidsByFunctionNameReply, err error)
	GetGoroutineStats(ctx context.Context, req *GetGoroutineStatsReq, opts ...http.CallOption) (rsp *GetGoroutineStatsReply, err error)
	GetHotFunctions(ctx context.Context, req *GetHotFunctionsReq, opts ...http.CallOption) (rsp *GetHotFunctionsReply, err error)
	GetParamsByID(ctx context.Context, req *GetParamsByIDReq, opts ...http.CallOption) (rsp *GetParamsByIDReply, err error)
	GetStaticDbFiles(ctx context.Context, req *GetStaticDbFilesRequest, opts ...http.CallOption) (rsp *GetStaticDbFilesResponse, err error)
	GetTraceGraph(ctx context.Context, req *GetTraceGraphReq, opts ...http.CallOption) (rsp *GetTraceGraphReply, err error)
	GetTracesByParentFunc(ctx context.Context, req *GetTracesByParentFuncReq, opts ...http.CallOption) (rsp *GetTracesByParentFuncReply, err error)
	VerifyProjectPath(ctx context.Context, req *VerifyProjectPathReq, opts ...http.CallOption) (rsp *VerifyProjectPathReply, err error)
}

type AnalysisHTTPClientImpl struct {
	cc *http.Client
}

func NewAnalysisHTTPClient(client *http.Client) AnalysisHTTPClient {
	return &AnalysisHTTPClientImpl{client}
}

func (c *AnalysisHTTPClientImpl) AnalyzeDbFile(ctx context.Context, in *AnalyzeDbFileRequest, opts ...http.CallOption) (*AnalyzeDbFileResponse, error) {
	var out AnalyzeDbFileResponse
	pattern := "/api/static/analyze"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAnalysisAnalyzeDbFile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) AnalyzeProjectPath(ctx context.Context, in *AnalyzeProjectPathRequest, opts ...http.CallOption) (*AnalyzeProjectPathResponse, error) {
	var out AnalyzeProjectPathResponse
	pattern := "/api/static/analyze/path"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAnalysisAnalyzeProjectPath))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) CheckDatabase(ctx context.Context, in *CheckDatabaseRequest, opts ...http.CallOption) (*CheckDatabaseResponse, error) {
	var out CheckDatabaseResponse
	pattern := "/api/check/db"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAnalysisCheckDatabase))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) GenerateImage(ctx context.Context, in *GenerateImageReq, opts ...http.CallOption) (*GenerateImageReply, error) {
	var out GenerateImageReply
	pattern := "/api/traces/{gid}/mermaid"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAnalysisGenerateImage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) GetAllFunctionName(ctx context.Context, in *GetAllFunctionNameReq, opts ...http.CallOption) (*GetAllFunctionNameReply, error) {
	var out GetAllFunctionNameReply
	pattern := "/api/functions"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAnalysisGetAllFunctionName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) GetAllGIDs(ctx context.Context, in *GetAllGIDsReq, opts ...http.CallOption) (*GetAllGIDsReply, error) {
	var out GetAllGIDsReply
	pattern := "/api/gids"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAnalysisGetAllGIDs))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) GetAllParentFuncNames(ctx context.Context, in *GetAllParentFuncNamesReq, opts ...http.CallOption) (*GetAllParentFuncNamesReply, error) {
	var out GetAllParentFuncNamesReply
	pattern := "/api/functions/parent"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAnalysisGetAllParentFuncNames))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) GetAnalysis(ctx context.Context, in *AnalysisRequest, opts ...http.CallOption) (*AnalysisReply, error) {
	var out AnalysisReply
	pattern := "/analysis/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAnalysisGetAnalysis))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) GetAnalysisByGID(ctx context.Context, in *AnalysisByGIDRequest, opts ...http.CallOption) (*AnalysisByGIDReply, error) {
	var out AnalysisByGIDReply
	pattern := "/api/traces/{gid}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAnalysisGetAnalysisByGID))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) GetChildFunctions(ctx context.Context, in *GetChildFunctionsReq, opts ...http.CallOption) (*GetChildFunctionsReply, error) {
	var out GetChildFunctionsReply
	pattern := "/api/functions/{func_name}/children"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAnalysisGetChildFunctions))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) GetFunctionAnalysis(ctx context.Context, in *GetFunctionAnalysisReq, opts ...http.CallOption) (*GetFunctionAnalysisReply, error) {
	var out GetFunctionAnalysisReply
	pattern := "/api/function/analysis"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAnalysisGetFunctionAnalysis))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) GetFunctionCallGraph(ctx context.Context, in *GetFunctionCallGraphReq, opts ...http.CallOption) (*GetFunctionCallGraphReply, error) {
	var out GetFunctionCallGraphReply
	pattern := "/api/function/{function_name}/graph"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAnalysisGetFunctionCallGraph))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) GetGidsByFunctionName(ctx context.Context, in *GetGidsByFunctionNameReq, opts ...http.CallOption) (*GetGidsByFunctionNameReply, error) {
	var out GetGidsByFunctionNameReply
	pattern := "/api/gids/function"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAnalysisGetGidsByFunctionName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) GetGoroutineStats(ctx context.Context, in *GetGoroutineStatsReq, opts ...http.CallOption) (*GetGoroutineStatsReply, error) {
	var out GetGoroutineStatsReply
	pattern := "/api/goroutine-stats"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAnalysisGetGoroutineStats))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) GetHotFunctions(ctx context.Context, in *GetHotFunctionsReq, opts ...http.CallOption) (*GetHotFunctionsReply, error) {
	var out GetHotFunctionsReply
	pattern := "/api/hot-functions"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAnalysisGetHotFunctions))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) GetParamsByID(ctx context.Context, in *GetParamsByIDReq, opts ...http.CallOption) (*GetParamsByIDReply, error) {
	var out GetParamsByIDReply
	pattern := "/api/params/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAnalysisGetParamsByID))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) GetStaticDbFiles(ctx context.Context, in *GetStaticDbFilesRequest, opts ...http.CallOption) (*GetStaticDbFilesResponse, error) {
	var out GetStaticDbFilesResponse
	pattern := "/api/static/dbfiles"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAnalysisGetStaticDbFiles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) GetTraceGraph(ctx context.Context, in *GetTraceGraphReq, opts ...http.CallOption) (*GetTraceGraphReply, error) {
	var out GetTraceGraphReply
	pattern := "/api/traces/{gid}/graph"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAnalysisGetTraceGraph))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) GetTracesByParentFunc(ctx context.Context, in *GetTracesByParentFuncReq, opts ...http.CallOption) (*GetTracesByParentFuncReply, error) {
	var out GetTracesByParentFuncReply
	pattern := "/api/traces/parent/{parent_func}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAnalysisGetTracesByParentFunc))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AnalysisHTTPClientImpl) VerifyProjectPath(ctx context.Context, in *VerifyProjectPathReq, opts ...http.CallOption) (*VerifyProjectPathReply, error) {
	var out VerifyProjectPathReply
	pattern := "/api/verify/path"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAnalysisVerifyProjectPath))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
