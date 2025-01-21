package dto

import "time"

type CreateDataSetReq struct {
	DsName   string         `json:"ds_name"`
	DsDesc   string         `json:"ds_desc"`
	SourceId int32          `json:"source_id"`
	Meta     []*DataSetData `json:"meta"`
	CreateBy string         `json:"create_by"`
}

type UpdateDataSetReq struct {
	Id     int            `json:"id"`
	DsName string         `json:"ds_name"`
	DsDesc string         `json:"ds_desc"`
	Meta   []*DataSetData `json:"meta"`
}

type DeleteDataSetReq struct {
	Ids []int32 `json:"ids"`
}

type QueryDataSetReq struct {
	PageSize int    `json:"page_size"`
	PageNum  int    `json:"page_num"`
	DsName   string `json:"ds_name"`
	DsType   string `json:"ds_type"`
}

type DataSetData struct {
	TableName   string       `json:"table_name"`
	Description string       `json:"description"`
	FieldMetas  []*FieldMeta `json:"field_metas"`
}

type FieldMeta struct {
	FieldName string `json:"field_name"`
	FieldDesc string `json:"field_desc"`
	FieldType string `json:"field_type"`
}

type DataSetResult struct {
	Id         int            `json:"id"`
	DsName     string         `json:"ds_name"`
	DsDesc     string         `json:"ds_desc"`
	DsType     string         `json:"ds_type"`
	SourceId   int32          `json:"source_id"`
	Meta       []*DataSetData `json:"meta"`
	CreateBy   string         `json:"create_by"`
	CreateTime time.Time      `json:"create_time"`
	UpdateTime time.Time      `json:"update_time"`
}

type QueryDataSetResp struct {
	PageSize int              `json:"page_size"`
	PageNum  int              `json:"page_num"`
	Count    int              `json:"count"`
	Data     []*DataSetResult `json:"data"`
}
