package dto

import (
	"time"
)

// 创建数据源
type CreateDataSourceReq struct {
	DsName  string                 `json:"name"`
	DsType  string                 `json:"type"`
	Desc    string                 `json:"desc"`
	Connect map[string]interface{} `json:"connect"`
}

// 删除数据源
type DeleteDataSourceReq struct {
	Ids []int32 `json:"id"`
}

// 查询数据源
type QueryDataSourceReq struct {
	DsName   string `json:"name"`
	PageSize int    `json:"page_size"`
	PageNum  int    `json:"page_num"`
}

type DataSourceData struct {
	Id         int32     `json:"id"`
	DsName     string    `json:"name"`
	DsType     string    `json:"type"`
	Desc       string    `json:"desc"`
	Connect    string    `json:"connect"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

type QueryDataSourceResp struct {
	Data     []*DataSourceData `json:"data"`
	PageSize int               `json:"pageSize"`
	PageNum  int               `json:"pageNum"`
	Count    int               `json:"count"`
}

// 更新数据源
type UpdateDataSourceReq struct {
	Id      int32                  `json:"id"`
	DsName  string                 `json:"name"`
	DsType  string                 `json:"type"`
	Desc    string                 `json:"desc"`
	Connect map[string]interface{} `json:"connect"`
}

// 测试数据源连接
type TestDataSourceReq struct {
	Id int32 `json:"id"`
}
