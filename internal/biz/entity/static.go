package entity

import "time"

// 数据库文件结构
type StaticDbFile struct {
	Path       string    `json:"path"`
	Name       string    `json:"name"`
	Size       int64     `json:"size"`
	CreateTime time.Time `json:"createTime"`
}

// 分析结果结构
type StaticAnalysisResult struct {
	TotalFunctions      int `json:"totalFunctions"`
	TotalCalls          int `json:"totalCalls"`
	TotalPackages       int `json:"totalPackages"`
	PackageDependencies []struct {
		Source string `json:"source"`
		Target string `json:"target"`
		Count  int    `json:"count"`
	} `json:"packageDependencies"`
	HotFunctions []struct {
		Name      string `json:"name"`
		CallCount int    `json:"callCount"`
	} `json:"hotFunctions"`
}
