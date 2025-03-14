package service

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/toheart/functrace"
	v1 "github.com/toheart/goanalysis/api/analysis/v1"
	"github.com/toheart/goanalysis/internal/biz/analysis"
	"github.com/toheart/goanalysis/internal/biz/entity"
	"github.com/toheart/goanalysis/internal/biz/rewrite"
	"google.golang.org/grpc"
)

// GreeterService is a greeter service.
type AnalysisService struct {
	v1.UnimplementedAnalysisServer

	uc  *analysis.AnalysisBiz
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewAnalysisService(uc *analysis.AnalysisBiz, logger log.Logger) *AnalysisService {
	return &AnalysisService{uc: uc, log: log.NewHelper(logger)}
}

func (a *AnalysisService) RegisterGrpc(svr *grpc.Server) {
	v1.RegisterAnalysisServer(svr, a)
}

func (a *AnalysisService) RegisterHttp(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return v1.RegisterAnalysisHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

// SayHello implements helloworld.GreeterServer.
func (a *AnalysisService) GetAnalysis(ctx context.Context, in *v1.AnalysisRequest) (*v1.AnalysisReply, error) {
	return &v1.AnalysisReply{Message: "Hello " + in.Name}, nil
}

func (a *AnalysisService) GetAnalysisByGID(ctx context.Context, in *v1.AnalysisByGIDRequest) (*v1.AnalysisByGIDReply, error) {
	traces, err := a.uc.GetTracesByGID(in.Dbpath, in.Gid)
	if err != nil {
		return nil, err
	}

	reply := &v1.AnalysisByGIDReply{}
	for _, trace := range traces {
		traceData := &v1.AnalysisByGIDReply_TraceData{
			Id:             int32(trace.ID),
			Name:           trace.Name,
			Gid:            int32(trace.GID),
			Indent:         int32(trace.Indent),
			ParamCount:     int32(len(trace.Params)),
			TimeCost:       trace.TimeCost,
			ParentFuncname: trace.ParentFuncName,
		}

		reply.TraceData = append(reply.TraceData, traceData)
	}
	return reply, nil
}

func (a *AnalysisService) GetAllGIDs(ctx context.Context, in *v1.GetAllGIDsReq) (*v1.GetAllGIDsReply, error) {
	page := in.Page
	limit := in.Limit
	includeMetrics := in.IncludeMetrics
	dbpath := in.Dbpath
	reply := &v1.GetAllGIDsReply{}
	gids, err := a.uc.GetAllGIDs(dbpath, int(page), int(limit))
	if err != nil {
		return nil, err
	}

	for _, gid := range gids {
		initialFunc, err := a.uc.GetInitialFunc(dbpath, gid)
		if err != nil {
			return nil, err
		}

		body := &v1.GetAllGIDsReply_Body{
			Gid:         gid,
			InitialFunc: initialFunc,
		}

		// 如果需要包含调用深度和执行时间
		if includeMetrics {
			// 获取调用深度
			depth, err := a.uc.GetGoroutineCallDepth(dbpath, gid)
			if err != nil {
				// 如果获取失败，使用默认值
				body.Depth = 0
			} else {
				body.Depth = int32(depth)
			}

			// 获取执行时间
			execTime, err := a.uc.GetGoroutineExecutionTime(dbpath, gid)
			if err != nil {
				// 如果获取失败，使用默认值
				body.ExecutionTime = "N/A"
			} else {
				body.ExecutionTime = execTime
			}
		}

		reply.Body = append(reply.Body, body)
	}

	// 获取总数
	total, err := a.uc.GetTotalGIDs(dbpath)
	if err != nil {
		return nil, err
	}
	reply.Total = int32(total)

	return reply, nil
}

func (a *AnalysisService) GetParamsByID(ctx context.Context, in *v1.GetParamsByIDReq) (*v1.GetParamsByIDReply, error) {
	params, err := a.uc.GetParamsByID(in.Dbpath, in.Id)
	if err != nil {
		return nil, err
	}

	reply := &v1.GetParamsByIDReply{}
	for _, param := range params {
		reply.Params = append(reply.Params, &v1.TraceParams{
			Pos:   int32(param.Pos),
			Param: param.Param,
		})
	}
	return reply, nil
}

func (a *AnalysisService) GenerateImage(ctx context.Context, in *v1.GenerateImageReq) (*v1.GenerateImageReply, error) {
	traces, err := a.uc.GetTracesByGID(in.Dbpath, in.Gid)
	if err != nil {
		return nil, err
	}

	// 构建Mermaid图表
	var mermaidText strings.Builder
	mermaidText.WriteString("graph TD\n")

	// 使用栈来跟踪调用关系
	stack := make([]functrace.TraceData, 0)

	for _, trace := range traces {
		// 根据缩进级别调整栈
		for len(stack) > 0 && stack[len(stack)-1].Indent >= trace.Indent {
			stack = stack[:len(stack)-1] // 弹出栈顶元素
		}

		// 添加节点
		nodeID := fmt.Sprintf("n%d", trace.ID)
		nodeName := sanitizeMermaidText(trace.Name)
		mermaidText.WriteString(fmt.Sprintf("    %s[\"%s\"]\n", nodeID, nodeName))

		// 如果有父节点，添加边
		if len(stack) > 0 {
			parentID := fmt.Sprintf("n%d", stack[len(stack)-1].ID)
			mermaidText.WriteString(fmt.Sprintf("    %s --> |%s| %s\n", parentID, trace.TimeCost, nodeID))
		}

		// 将当前节点压入栈
		stack = append(stack, trace)
	}

	return &v1.GenerateImageReply{
		Image: mermaidText.String(),
	}, nil
}

func sanitizeMermaidText(text string) string {
	// 替换可能导致Mermaid语法问题的字符
	text = strings.ReplaceAll(text, "\"", "'")
	return text
}

func (a *AnalysisService) GetAllFunctionName(ctx context.Context, in *v1.GetAllFunctionNameReq) (*v1.GetAllFunctionNameReply, error) {
	functionNames, err := a.uc.GetAllFunctionName(in.Dbpath)
	if err != nil {
		return nil, err
	}
	return &v1.GetAllFunctionNameReply{FunctionNames: functionNames}, nil
}

func (a *AnalysisService) GetGidsByFunctionName(ctx context.Context, in *v1.GetGidsByFunctionNameReq) (*v1.GetGidsByFunctionNameReply, error) {
	// 获取函数名称
	functionName := in.FunctionName
	includeMetrics := in.IncludeMetrics
	dbpath := in.Path // 使用 Path 字段作为 dbpath

	// 获取所有GID
	allGids, err := a.uc.GetAllGIDs(dbpath, 0, 1000) // 假设最多1000个GID
	if err != nil {
		return nil, err
	}

	// 过滤包含指定函数的GID
	var matchingGids []uint64
	for _, gid := range allGids {
		traces, err := a.uc.GetTracesByGID(dbpath, strconv.FormatUint(gid, 10))
		if err != nil {
			continue
		}

		// 检查是否包含指定函数
		for _, trace := range traces {
			// 简化函数名称进行比较
			simplifiedName := getLastSegment(removeParentheses(trace.Name))
			if simplifiedName == functionName {
				matchingGids = append(matchingGids, gid)
				break
			}
		}
	}

	// 构建响应
	reply := &v1.GetGidsByFunctionNameReply{}
	for _, gid := range matchingGids {
		initialFunc, err := a.uc.GetInitialFunc(dbpath, gid)
		if err != nil {
			continue
		}

		body := &v1.GetGidsByFunctionNameReply_Body{
			Gid:         gid,
			InitialFunc: initialFunc,
		}

		// 如果需要包含调用深度和执行时间
		if includeMetrics {
			// 获取调用深度
			depth, err := a.uc.GetGoroutineCallDepth(dbpath, gid)
			if err != nil {
				// 如果获取失败，使用默认值
				body.Depth = 0
			} else {
				body.Depth = int32(depth)
			}

			// 获取执行时间
			execTime, err := a.uc.GetGoroutineExecutionTime(dbpath, gid)
			if err != nil {
				// 如果获取失败，使用默认值
				body.ExecutionTime = "N/A"
			} else {
				body.ExecutionTime = execTime
			}
		}

		reply.Body = append(reply.Body, body)
	}

	// 获取总数
	reply.Total = int32(len(matchingGids))

	return reply, nil
}

func getLastSegment(name string) string {
	parts := strings.Split(name, ".")
	return parts[len(parts)-1]
}

func removeParentheses(name string) string {
	idx := strings.Index(name, "(")
	if idx > 0 {
		return name[:idx]
	}
	return name
}

func (s *AnalysisService) VerifyProjectPath(ctx context.Context, in *v1.VerifyProjectPathReq) (*v1.VerifyProjectPathReply, error) {
	verified := s.uc.VerifyProjectPath(in.Path)
	return &v1.VerifyProjectPathReply{Verified: verified}, nil
}

// 以下是需要保留的动态分析相关方法，但需要移除静态分析相关的方法
// 保留的方法：
// - GetTraceGraph
// - GetTracesByParentFunc
// - GetAllParentFuncNames
// - GetChildFunctions
// - GetHotFunctions
// - GetGoroutineStats
// - GetFunctionAnalysis
// - GetFunctionCallGraph

// 定义图形数据的结构
type GraphNode struct {
	ID        string `json:"id"`        // 节点唯一标识
	Name      string `json:"name"`      // 函数名称
	CallCount int    `json:"callCount"` // 调用次数
	Package   string `json:"package"`   // 包名
	TimeCost  string `json:"timeCost"`  // 执行耗时
}

type GraphEdge struct {
	Source string `json:"source"` // 源节点ID
	Target string `json:"target"` // 目标节点ID
	Label  string `json:"label"`  // 边标签（通常是耗时）
	Count  int    `json:"count"`  // 调用次数
}

type GraphData struct {
	Nodes []GraphNode `json:"nodes"`
	Edges []GraphEdge `json:"edges"`
}

func (a *AnalysisService) GetTraceGraph(ctx context.Context, in *v1.GetTraceGraphReq) (*v1.GetTraceGraphReply, error) {
	// 记录请求信息
	a.log.WithContext(ctx).Infof("GetTraceGraph request received: gid=%s, dbpath=%s", in.Gid, in.Dbpath)

	// 参数验证
	if in.Gid == "" {
		a.log.WithContext(ctx).Errorf("GetTraceGraph failed: gid is empty")
		return nil, errors.New("gid is required")
	}

	if in.Dbpath == "" {
		a.log.WithContext(ctx).Errorf("GetTraceGraph failed: dbpath is empty")
		return nil, errors.New("dbpath is required")
	}

	// 获取指定GID的调用跟踪数据
	traces, err := a.uc.GetTracesByGID(in.Dbpath, in.Gid)
	if err != nil {
		a.log.WithContext(ctx).Errorf("Failed to get traces for GID %s: %v", in.Gid, err)
		return nil, fmt.Errorf("failed to get traces: %w", err)
	}

	// 检查是否有数据
	if len(traces) == 0 {
		a.log.WithContext(ctx).Warnf("No trace data found for GID %s", in.Gid)
		// 返回空结果而不是错误
		return &v1.GetTraceGraphReply{
			Nodes: []*v1.GraphNode{},
			Edges: []*v1.GraphEdge{},
		}, nil
	}

	// 构建图形数据
	a.log.WithContext(ctx).Infof("Building graph from %d traces", len(traces))
	graphData := buildGraphFromTraces(traces)

	// 记录节点和边的数量
	a.log.WithContext(ctx).Infof("Graph built with %d nodes and %d edges",
		len(graphData.Nodes), len(graphData.Edges))

	// 转换为protobuf类型并返回
	return &v1.GetTraceGraphReply{
		Nodes: convertToProtoNodes(graphData.Nodes),
		Edges: convertToProtoEdges(graphData.Edges),
	}, nil
}

func buildGraphFromTraces(traces []functrace.TraceData) *GraphData {
	// 初始化图数据结构
	graphData := &GraphData{
		Nodes: []GraphNode{},
		Edges: []GraphEdge{},
	}

	// 如果没有跟踪数据，直接返回空图
	if len(traces) == 0 {
		return graphData
	}

	// 使用map来跟踪已添加的节点，避免重复
	nodeMap := make(map[string]bool)

	// 跟踪函数调用次数
	callCounts := make(map[string]int)

	// 使用map来存储节点信息，便于后续构建
	nodeInfoMap := make(map[string]functrace.TraceData)

	// 使用栈来跟踪调用关系
	stack := make([]functrace.TraceData, 0, 32) // 预分配一定容量以提高性能

	// 第一遍遍历：构建调用关系和统计调用次数
	for _, trace := range traces {
		// 根据缩进级别调整栈
		for len(stack) > 0 && stack[len(stack)-1].Indent >= trace.Indent {
			stack = stack[:len(stack)-1] // 弹出栈顶元素
		}

		// 生成节点ID
		nodeID := fmt.Sprintf("n%d", trace.ID)

		// 记录节点信息
		nodeInfoMap[nodeID] = trace

		// 统计调用次数
		if !nodeMap[nodeID] {
			nodeMap[nodeID] = true
			callCounts[trace.Name]++
		}

		// 如果有父节点，添加边
		if len(stack) > 0 {
			parentID := fmt.Sprintf("n%d", stack[len(stack)-1].ID)

			// 添加边
			graphData.Edges = append(graphData.Edges, GraphEdge{
				Source: parentID,
				Target: nodeID,
				Label:  trace.TimeCost,
				Count:  callCounts[trace.Name],
			})
		}

		// 将当前节点压入栈
		stack = append(stack, trace)
	}

	// 第二遍：构建节点列表
	// 预分配节点数组以提高性能
	graphData.Nodes = make([]GraphNode, 0, len(nodeMap))

	// 按照ID顺序添加节点，使图形布局更稳定
	nodeIDs := make([]string, 0, len(nodeMap))
	for id := range nodeMap {
		nodeIDs = append(nodeIDs, id)
	}
	sort.Strings(nodeIDs)

	for _, id := range nodeIDs {
		if trace, ok := nodeInfoMap[id]; ok {
			// 提取包名
			packageName := extractPackageName(trace.Name)

			graphData.Nodes = append(graphData.Nodes, GraphNode{
				ID:        id,
				Name:      trace.Name,
				CallCount: callCounts[trace.Name],
				Package:   packageName,
				TimeCost:  trace.TimeCost,
			})
		}
	}

	return graphData
}

// 从函数名中提取包名
func extractPackageName(funcName string) string {
	parts := strings.Split(funcName, ".")
	if len(parts) <= 1 {
		return "main" // 默认包名
	}

	// 如果是形如 github.com/user/repo/pkg.Func 的格式
	return strings.Join(parts[:len(parts)-1], ".")
}

// 转换为protobuf类型的辅助函数
func convertToProtoNodes(nodes []GraphNode) []*v1.GraphNode {
	protoNodes := make([]*v1.GraphNode, len(nodes))
	for i, node := range nodes {
		protoNodes[i] = &v1.GraphNode{
			Id:        node.ID,
			Name:      node.Name,
			CallCount: int32(node.CallCount),
			// 暂时注释掉不存在的字段，等proto文件更新后再启用
			// Package:   node.Package,
			// TimeCost:  node.TimeCost,
		}
	}
	return protoNodes
}

func convertToProtoEdges(edges []GraphEdge) []*v1.GraphEdge {
	protoEdges := make([]*v1.GraphEdge, len(edges))
	for i, edge := range edges {
		protoEdges[i] = &v1.GraphEdge{
			Source: edge.Source,
			Target: edge.Target,
			Label:  edge.Label,
			// 暂时注释掉不存在的字段，等proto文件更新后再启用
			// Count:  int32(edge.Count),
		}
	}
	return protoEdges
}

// GetTracesByParentFunc 根据父函数名称获取函数调用
func (a *AnalysisService) GetTracesByParentFunc(ctx context.Context, in *v1.GetTracesByParentFuncReq) (*v1.GetTracesByParentFuncReply, error) {
	traces, err := a.uc.GetTracesByParentFunc(in.Dbpath, in.ParentFunc)
	if err != nil {
		return nil, err
	}

	reply := &v1.GetTracesByParentFuncReply{}
	for _, trace := range traces {
		traceData := &v1.GetTracesByParentFuncReply_TraceData{
			Id:             int32(trace.ID),
			Name:           trace.Name,
			Gid:            int32(trace.GID),
			Indent:         int32(trace.Indent),
			ParamCount:     int32(len(trace.Params)),
			TimeCost:       trace.TimeCost,
			ParentFuncname: trace.ParentFuncName,
		}

		reply.TraceData = append(reply.TraceData, traceData)
	}
	return reply, nil
}

// GetAllParentFuncNames 获取所有的父函数名称
func (a *AnalysisService) GetAllParentFuncNames(ctx context.Context, in *v1.GetAllParentFuncNamesReq) (*v1.GetAllParentFuncNamesReply, error) {
	parentFuncNames, err := a.uc.GetAllParentFuncNames(in.Dbpath)
	if err != nil {
		return nil, err
	}
	return &v1.GetAllParentFuncNamesReply{
		ParentFuncNames: parentFuncNames,
	}, nil
}

// GetChildFunctions 获取函数的子函数
func (a *AnalysisService) GetChildFunctions(ctx context.Context, in *v1.GetChildFunctionsReq) (*v1.GetChildFunctionsReply, error) {
	childFunctions, err := a.uc.GetChildFunctions(in.Dbpath, in.FuncName)
	if err != nil {
		return nil, err
	}
	return &v1.GetChildFunctionsReply{
		ChildFunctions: childFunctions,
	}, nil
}

// GetHotFunctions 获取热点函数分析数据
func (a *AnalysisService) GetHotFunctions(ctx context.Context, in *v1.GetHotFunctionsReq) (*v1.GetHotFunctionsReply, error) {
	hotFunctions, err := a.uc.GetHotFunctions(in.Dbpath, in.SortBy)
	if err != nil {
		return nil, err
	}

	reply := &v1.GetHotFunctionsReply{}
	for _, hf := range hotFunctions {
		hotFunc := &v1.GetHotFunctionsReply_HotFunction{
			Name:      hf.Name,
			Package:   hf.Package,
			CallCount: int32(hf.CallCount),
			TotalTime: hf.TotalTime,
			AvgTime:   hf.AvgTime,
		}
		reply.Functions = append(reply.Functions, hotFunc)
	}
	return reply, nil
}

// GetGoroutineStats 获取Goroutine统计信息
func (a *AnalysisService) GetGoroutineStats(ctx context.Context, in *v1.GetGoroutineStatsReq) (*v1.GetGoroutineStatsReply, error) {
	stats, err := a.uc.GetGoroutineStats(in.Dbpath)
	if err != nil {
		return nil, err
	}
	return &v1.GetGoroutineStatsReply{
		Active:   int32(stats.Active),
		AvgTime:  stats.AvgTime,
		MaxDepth: int32(stats.MaxDepth),
	}, nil
}

// GetFunctionAnalysis 获取函数调用关系分析
func (a *AnalysisService) GetFunctionAnalysis(ctx context.Context, in *v1.GetFunctionAnalysisReq) (*v1.GetFunctionAnalysisReply, error) {
	functionNodes, err := a.uc.GetFunctionAnalysis(in.Path, in.FunctionName, in.Type)
	if err != nil {
		return nil, err
	}

	reply := &v1.GetFunctionAnalysisReply{}
	reply.CallData = convertToProtoFunctionNodes(functionNodes)
	return reply, nil
}

// 转换为protobuf类型的辅助函数
func convertToProtoFunctionNodes(nodes []entity.FunctionNode) []*v1.GetFunctionAnalysisReply_FunctionNode {
	protoNodes := make([]*v1.GetFunctionAnalysisReply_FunctionNode, len(nodes))
	for i, node := range nodes {
		protoNode := &v1.GetFunctionAnalysisReply_FunctionNode{
			Id:        node.ID,
			Name:      node.Name,
			Package:   node.Package,
			CallCount: int32(node.CallCount),
			AvgTime:   node.AvgTime,
		}

		if len(node.Children) > 0 {
			protoNode.Children = convertToProtoFunctionNodes(node.Children)
		}

		protoNodes[i] = protoNode
	}
	return protoNodes
}

// GetFunctionCallGraph 获取函数调用关系图
func (a *AnalysisService) GetFunctionCallGraph(ctx context.Context, in *v1.GetFunctionCallGraphReq) (*v1.GetFunctionCallGraphReply, error) {
	callGraph, err := a.uc.GetFunctionCallGraph(in.Dbpath, in.FunctionName, int(in.Depth), in.Direction)
	if err != nil {
		return nil, err
	}

	reply := &v1.GetFunctionCallGraphReply{}

	// 转换节点
	for _, node := range callGraph.Nodes {
		protoNode := &v1.GetFunctionCallGraphReply_GraphNode{
			Id:        node.ID,
			Name:      node.Name,
			Package:   node.Package,
			CallCount: int32(node.CallCount),
			AvgTime:   node.AvgTime,
			NodeType:  node.NodeType,
		}
		reply.Nodes = append(reply.Nodes, protoNode)
	}

	// 转换边
	for _, edge := range callGraph.Edges {
		protoEdge := &v1.GetFunctionCallGraphReply_GraphEdge{
			Source:   edge.Source,
			Target:   edge.Target,
			Label:    edge.Label,
			EdgeType: edge.EdgeType,
		}
		reply.Edges = append(reply.Edges, protoEdge)
	}

	return reply, nil
}

// InstrumentProject 实现插桩功能
func (a *AnalysisService) InstrumentProject(ctx context.Context, in *v1.InstrumentProjectReq) (*v1.InstrumentProjectReply, error) {
	a.log.Infof("Instrumenting project at path: %s", in.Path)

	if in.Path == "" {
		return &v1.InstrumentProjectReply{
			Success: false,
			Message: "项目路径不能为空",
		}, nil
	}

	// 调用rewrite包进行插桩
	defer func() {
		if r := recover(); r != nil {
			a.log.Errorf("Panic during instrumentation: %v", r)
		}
	}()

	// 执行插桩操作
	rewrite.RewriteDir(in.Path)

	return &v1.InstrumentProjectReply{
		Success: true,
		Message: "项目插桩成功，现在可以运行您的程序进行分析",
	}, nil
}
