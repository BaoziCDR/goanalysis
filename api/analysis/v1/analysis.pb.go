// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0--rc1
// source: analysis/v1/analysis.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetGidsByFunctionNameReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FunctionName  string                 `protobuf:"bytes,1,opt,name=functionName,proto3" json:"functionName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetGidsByFunctionNameReq) Reset() {
	*x = GetGidsByFunctionNameReq{}
	mi := &file_analysis_v1_analysis_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGidsByFunctionNameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGidsByFunctionNameReq) ProtoMessage() {}

func (x *GetGidsByFunctionNameReq) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_v1_analysis_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGidsByFunctionNameReq.ProtoReflect.Descriptor instead.
func (*GetGidsByFunctionNameReq) Descriptor() ([]byte, []int) {
	return file_analysis_v1_analysis_proto_rawDescGZIP(), []int{0}
}

func (x *GetGidsByFunctionNameReq) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

type GetGidsByFunctionNameReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Gids          []string               `protobuf:"bytes,1,rep,name=gids,proto3" json:"gids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetGidsByFunctionNameReply) Reset() {
	*x = GetGidsByFunctionNameReply{}
	mi := &file_analysis_v1_analysis_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGidsByFunctionNameReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGidsByFunctionNameReply) ProtoMessage() {}

func (x *GetGidsByFunctionNameReply) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_v1_analysis_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGidsByFunctionNameReply.ProtoReflect.Descriptor instead.
func (*GetGidsByFunctionNameReply) Descriptor() ([]byte, []int) {
	return file_analysis_v1_analysis_proto_rawDescGZIP(), []int{1}
}

func (x *GetGidsByFunctionNameReply) GetGids() []string {
	if x != nil {
		return x.Gids
	}
	return nil
}

type GetAllFunctionNameReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllFunctionNameReq) Reset() {
	*x = GetAllFunctionNameReq{}
	mi := &file_analysis_v1_analysis_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllFunctionNameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllFunctionNameReq) ProtoMessage() {}

func (x *GetAllFunctionNameReq) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_v1_analysis_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllFunctionNameReq.ProtoReflect.Descriptor instead.
func (*GetAllFunctionNameReq) Descriptor() ([]byte, []int) {
	return file_analysis_v1_analysis_proto_rawDescGZIP(), []int{2}
}

type GetAllFunctionNameReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FunctionNames []string               `protobuf:"bytes,1,rep,name=functionNames,proto3" json:"functionNames,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllFunctionNameReply) Reset() {
	*x = GetAllFunctionNameReply{}
	mi := &file_analysis_v1_analysis_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllFunctionNameReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllFunctionNameReply) ProtoMessage() {}

func (x *GetAllFunctionNameReply) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_v1_analysis_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllFunctionNameReply.ProtoReflect.Descriptor instead.
func (*GetAllFunctionNameReply) Descriptor() ([]byte, []int) {
	return file_analysis_v1_analysis_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllFunctionNameReply) GetFunctionNames() []string {
	if x != nil {
		return x.FunctionNames
	}
	return nil
}

type GenerateImageReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Gid           string                 `protobuf:"bytes,1,opt,name=gid,proto3" json:"gid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateImageReq) Reset() {
	*x = GenerateImageReq{}
	mi := &file_analysis_v1_analysis_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateImageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateImageReq) ProtoMessage() {}

func (x *GenerateImageReq) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_v1_analysis_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateImageReq.ProtoReflect.Descriptor instead.
func (*GenerateImageReq) Descriptor() ([]byte, []int) {
	return file_analysis_v1_analysis_proto_rawDescGZIP(), []int{4}
}

func (x *GenerateImageReq) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

type GenerateImageReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Image         string                 `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateImageReply) Reset() {
	*x = GenerateImageReply{}
	mi := &file_analysis_v1_analysis_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateImageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateImageReply) ProtoMessage() {}

func (x *GenerateImageReply) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_v1_analysis_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateImageReply.ProtoReflect.Descriptor instead.
func (*GenerateImageReply) Descriptor() ([]byte, []int) {
	return file_analysis_v1_analysis_proto_rawDescGZIP(), []int{5}
}

func (x *GenerateImageReply) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

// The request message containing the user's name.
type AnalysisRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnalysisRequest) Reset() {
	*x = AnalysisRequest{}
	mi := &file_analysis_v1_analysis_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnalysisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisRequest) ProtoMessage() {}

func (x *AnalysisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_v1_analysis_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisRequest.ProtoReflect.Descriptor instead.
func (*AnalysisRequest) Descriptor() ([]byte, []int) {
	return file_analysis_v1_analysis_proto_rawDescGZIP(), []int{6}
}

func (x *AnalysisRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type AnalysisReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnalysisReply) Reset() {
	*x = AnalysisReply{}
	mi := &file_analysis_v1_analysis_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnalysisReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisReply) ProtoMessage() {}

func (x *AnalysisReply) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_v1_analysis_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisReply.ProtoReflect.Descriptor instead.
func (*AnalysisReply) Descriptor() ([]byte, []int) {
	return file_analysis_v1_analysis_proto_rawDescGZIP(), []int{7}
}

func (x *AnalysisReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AnalysisByGIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Gid           string                 `protobuf:"bytes,1,opt,name=gid,proto3" json:"gid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnalysisByGIDRequest) Reset() {
	*x = AnalysisByGIDRequest{}
	mi := &file_analysis_v1_analysis_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnalysisByGIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisByGIDRequest) ProtoMessage() {}

func (x *AnalysisByGIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_v1_analysis_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisByGIDRequest.ProtoReflect.Descriptor instead.
func (*AnalysisByGIDRequest) Descriptor() ([]byte, []int) {
	return file_analysis_v1_analysis_proto_rawDescGZIP(), []int{8}
}

func (x *AnalysisByGIDRequest) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

type TraceParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pos           int32                  `protobuf:"varint,1,opt,name=pos,proto3" json:"pos,omitempty"`
	Param         string                 `protobuf:"bytes,2,opt,name=param,proto3" json:"param,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TraceParams) Reset() {
	*x = TraceParams{}
	mi := &file_analysis_v1_analysis_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TraceParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceParams) ProtoMessage() {}

func (x *TraceParams) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_v1_analysis_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceParams.ProtoReflect.Descriptor instead.
func (*TraceParams) Descriptor() ([]byte, []int) {
	return file_analysis_v1_analysis_proto_rawDescGZIP(), []int{9}
}

func (x *TraceParams) GetPos() int32 {
	if x != nil {
		return x.Pos
	}
	return 0
}

func (x *TraceParams) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

type AnalysisByGIDReply struct {
	state         protoimpl.MessageState          `protogen:"open.v1"`
	TraceData     []*AnalysisByGIDReply_TraceData `protobuf:"bytes,1,rep,name=trace_data,json=traceData,proto3" json:"trace_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnalysisByGIDReply) Reset() {
	*x = AnalysisByGIDReply{}
	mi := &file_analysis_v1_analysis_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnalysisByGIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisByGIDReply) ProtoMessage() {}

func (x *AnalysisByGIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_v1_analysis_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisByGIDReply.ProtoReflect.Descriptor instead.
func (*AnalysisByGIDReply) Descriptor() ([]byte, []int) {
	return file_analysis_v1_analysis_proto_rawDescGZIP(), []int{10}
}

func (x *AnalysisByGIDReply) GetTraceData() []*AnalysisByGIDReply_TraceData {
	if x != nil {
		return x.TraceData
	}
	return nil
}

type GetAllGIDsReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllGIDsReq) Reset() {
	*x = GetAllGIDsReq{}
	mi := &file_analysis_v1_analysis_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllGIDsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllGIDsReq) ProtoMessage() {}

func (x *GetAllGIDsReq) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_v1_analysis_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllGIDsReq.ProtoReflect.Descriptor instead.
func (*GetAllGIDsReq) Descriptor() ([]byte, []int) {
	return file_analysis_v1_analysis_proto_rawDescGZIP(), []int{11}
}

type GetAllGIDsReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Gids          []uint64               `protobuf:"varint,1,rep,packed,name=gids,proto3" json:"gids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllGIDsReply) Reset() {
	*x = GetAllGIDsReply{}
	mi := &file_analysis_v1_analysis_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllGIDsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllGIDsReply) ProtoMessage() {}

func (x *GetAllGIDsReply) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_v1_analysis_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllGIDsReply.ProtoReflect.Descriptor instead.
func (*GetAllGIDsReply) Descriptor() ([]byte, []int) {
	return file_analysis_v1_analysis_proto_rawDescGZIP(), []int{12}
}

func (x *GetAllGIDsReply) GetGids() []uint64 {
	if x != nil {
		return x.Gids
	}
	return nil
}

type GetParamsByIDReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetParamsByIDReq) Reset() {
	*x = GetParamsByIDReq{}
	mi := &file_analysis_v1_analysis_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetParamsByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetParamsByIDReq) ProtoMessage() {}

func (x *GetParamsByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_v1_analysis_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetParamsByIDReq.ProtoReflect.Descriptor instead.
func (*GetParamsByIDReq) Descriptor() ([]byte, []int) {
	return file_analysis_v1_analysis_proto_rawDescGZIP(), []int{13}
}

func (x *GetParamsByIDReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetParamsByIDReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Params        []*TraceParams         `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetParamsByIDReply) Reset() {
	*x = GetParamsByIDReply{}
	mi := &file_analysis_v1_analysis_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetParamsByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetParamsByIDReply) ProtoMessage() {}

func (x *GetParamsByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_v1_analysis_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetParamsByIDReply.ProtoReflect.Descriptor instead.
func (*GetParamsByIDReply) Descriptor() ([]byte, []int) {
	return file_analysis_v1_analysis_proto_rawDescGZIP(), []int{14}
}

func (x *GetParamsByIDReply) GetParams() []*TraceParams {
	if x != nil {
		return x.Params
	}
	return nil
}

type AnalysisByGIDReply_TraceData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Gid           int32                  `protobuf:"varint,3,opt,name=gid,proto3" json:"gid,omitempty"`
	Indent        int32                  `protobuf:"varint,4,opt,name=indent,proto3" json:"indent,omitempty"`
	Params        []*TraceParams         `protobuf:"bytes,5,rep,name=params,proto3" json:"params,omitempty"`
	ParamCount    int32                  `protobuf:"varint,6,opt,name=paramCount,proto3" json:"paramCount,omitempty"`
	TimeCost      string                 `protobuf:"bytes,7,opt,name=timeCost,proto3" json:"timeCost,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnalysisByGIDReply_TraceData) Reset() {
	*x = AnalysisByGIDReply_TraceData{}
	mi := &file_analysis_v1_analysis_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnalysisByGIDReply_TraceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisByGIDReply_TraceData) ProtoMessage() {}

func (x *AnalysisByGIDReply_TraceData) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_v1_analysis_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisByGIDReply_TraceData.ProtoReflect.Descriptor instead.
func (*AnalysisByGIDReply_TraceData) Descriptor() ([]byte, []int) {
	return file_analysis_v1_analysis_proto_rawDescGZIP(), []int{10, 0}
}

func (x *AnalysisByGIDReply_TraceData) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AnalysisByGIDReply_TraceData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AnalysisByGIDReply_TraceData) GetGid() int32 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *AnalysisByGIDReply_TraceData) GetIndent() int32 {
	if x != nil {
		return x.Indent
	}
	return 0
}

func (x *AnalysisByGIDReply_TraceData) GetParams() []*TraceParams {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *AnalysisByGIDReply_TraceData) GetParamCount() int32 {
	if x != nil {
		return x.ParamCount
	}
	return 0
}

func (x *AnalysisByGIDReply_TraceData) GetTimeCost() string {
	if x != nil {
		return x.TimeCost
	}
	return ""
}

var File_analysis_v1_analysis_proto protoreflect.FileDescriptor

var file_analysis_v1_analysis_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x47, 0x69,
	0x64, 0x73, 0x42, 0x79, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x47, 0x69,
	0x64, 0x73, 0x42, 0x79, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x67, 0x69, 0x64, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x22, 0x3f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x24, 0x0a,
	0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x22, 0x24, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x12, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x0f, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0d,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x14, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x42, 0x79, 0x47, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x69,
	0x64, 0x22, 0x35, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70,
	0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x22, 0xa8, 0x02, 0x0a, 0x12, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x42, 0x79, 0x47, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x48, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x42, 0x79, 0x47, 0x49, 0x44, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0xc7, 0x01, 0x0a, 0x09, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x67,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69,
	0x6e, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x43,
	0x6f, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x43,
	0x6f, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x49, 0x44,
	0x73, 0x52, 0x65, 0x71, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x49,
	0x44, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x04, 0x67, 0x69, 0x64, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x46, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x32, 0x9b, 0x06, 0x0a, 0x08, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x12, 0x61, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x12, 0x1c, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x71, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x42, 0x79, 0x47, 0x49, 0x44, 0x12, 0x21, 0x2e, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x42, 0x79, 0x47, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x42, 0x79, 0x47, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x67, 0x69, 0x64, 0x7d, 0x12, 0x59, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x47, 0x49, 0x44, 0x73, 0x12, 0x1a, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x49, 0x44,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x49, 0x44, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x69, 0x64, 0x73, 0x12, 0x69, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1d, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x72, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x1d, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x1f, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x67, 0x69, 0x64, 0x7d, 0x2f, 0x6d, 0x65, 0x72,
	0x6d, 0x61, 0x69, 0x64, 0x12, 0x76, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x24,
	0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x86, 0x01, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x47, 0x69, 0x64, 0x73, 0x42, 0x79, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x69, 0x64, 0x73, 0x42, 0x79, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47,
	0x69, 0x64, 0x73, 0x42, 0x79, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01,
	0x2a, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x69, 0x64, 0x73, 0x2f, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x2f, 0x67, 0x6f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_analysis_v1_analysis_proto_rawDescOnce sync.Once
	file_analysis_v1_analysis_proto_rawDescData []byte
)

func file_analysis_v1_analysis_proto_rawDescGZIP() []byte {
	file_analysis_v1_analysis_proto_rawDescOnce.Do(func() {
		file_analysis_v1_analysis_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_analysis_v1_analysis_proto_rawDesc), len(file_analysis_v1_analysis_proto_rawDesc)))
	})
	return file_analysis_v1_analysis_proto_rawDescData
}

var file_analysis_v1_analysis_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_analysis_v1_analysis_proto_goTypes = []any{
	(*GetGidsByFunctionNameReq)(nil),     // 0: analysis.v1.GetGidsByFunctionNameReq
	(*GetGidsByFunctionNameReply)(nil),   // 1: analysis.v1.GetGidsByFunctionNameReply
	(*GetAllFunctionNameReq)(nil),        // 2: analysis.v1.GetAllFunctionNameReq
	(*GetAllFunctionNameReply)(nil),      // 3: analysis.v1.GetAllFunctionNameReply
	(*GenerateImageReq)(nil),             // 4: analysis.v1.GenerateImageReq
	(*GenerateImageReply)(nil),           // 5: analysis.v1.GenerateImageReply
	(*AnalysisRequest)(nil),              // 6: analysis.v1.AnalysisRequest
	(*AnalysisReply)(nil),                // 7: analysis.v1.AnalysisReply
	(*AnalysisByGIDRequest)(nil),         // 8: analysis.v1.AnalysisByGIDRequest
	(*TraceParams)(nil),                  // 9: analysis.v1.TraceParams
	(*AnalysisByGIDReply)(nil),           // 10: analysis.v1.AnalysisByGIDReply
	(*GetAllGIDsReq)(nil),                // 11: analysis.v1.GetAllGIDsReq
	(*GetAllGIDsReply)(nil),              // 12: analysis.v1.GetAllGIDsReply
	(*GetParamsByIDReq)(nil),             // 13: analysis.v1.GetParamsByIDReq
	(*GetParamsByIDReply)(nil),           // 14: analysis.v1.GetParamsByIDReply
	(*AnalysisByGIDReply_TraceData)(nil), // 15: analysis.v1.AnalysisByGIDReply.TraceData
}
var file_analysis_v1_analysis_proto_depIdxs = []int32{
	15, // 0: analysis.v1.AnalysisByGIDReply.trace_data:type_name -> analysis.v1.AnalysisByGIDReply.TraceData
	9,  // 1: analysis.v1.GetParamsByIDReply.params:type_name -> analysis.v1.TraceParams
	9,  // 2: analysis.v1.AnalysisByGIDReply.TraceData.params:type_name -> analysis.v1.TraceParams
	6,  // 3: analysis.v1.Analysis.GetAnalysis:input_type -> analysis.v1.AnalysisRequest
	8,  // 4: analysis.v1.Analysis.GetAnalysisByGID:input_type -> analysis.v1.AnalysisByGIDRequest
	11, // 5: analysis.v1.Analysis.GetAllGIDs:input_type -> analysis.v1.GetAllGIDsReq
	13, // 6: analysis.v1.Analysis.GetParamsByID:input_type -> analysis.v1.GetParamsByIDReq
	4,  // 7: analysis.v1.Analysis.GenerateImage:input_type -> analysis.v1.GenerateImageReq
	2,  // 8: analysis.v1.Analysis.GetAllFunctionName:input_type -> analysis.v1.GetAllFunctionNameReq
	0,  // 9: analysis.v1.Analysis.GetGidsByFunctionName:input_type -> analysis.v1.GetGidsByFunctionNameReq
	7,  // 10: analysis.v1.Analysis.GetAnalysis:output_type -> analysis.v1.AnalysisReply
	10, // 11: analysis.v1.Analysis.GetAnalysisByGID:output_type -> analysis.v1.AnalysisByGIDReply
	12, // 12: analysis.v1.Analysis.GetAllGIDs:output_type -> analysis.v1.GetAllGIDsReply
	14, // 13: analysis.v1.Analysis.GetParamsByID:output_type -> analysis.v1.GetParamsByIDReply
	5,  // 14: analysis.v1.Analysis.GenerateImage:output_type -> analysis.v1.GenerateImageReply
	3,  // 15: analysis.v1.Analysis.GetAllFunctionName:output_type -> analysis.v1.GetAllFunctionNameReply
	1,  // 16: analysis.v1.Analysis.GetGidsByFunctionName:output_type -> analysis.v1.GetGidsByFunctionNameReply
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_analysis_v1_analysis_proto_init() }
func file_analysis_v1_analysis_proto_init() {
	if File_analysis_v1_analysis_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_analysis_v1_analysis_proto_rawDesc), len(file_analysis_v1_analysis_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_analysis_v1_analysis_proto_goTypes,
		DependencyIndexes: file_analysis_v1_analysis_proto_depIdxs,
		MessageInfos:      file_analysis_v1_analysis_proto_msgTypes,
	}.Build()
	File_analysis_v1_analysis_proto = out.File
	file_analysis_v1_analysis_proto_goTypes = nil
	file_analysis_v1_analysis_proto_depIdxs = nil
}
