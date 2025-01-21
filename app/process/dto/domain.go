package dto

import "time"

// 新增
type CreateDomainReq struct {
	BucName   string   `json:"buc_name"`
	BucDesc   string   `json:"buc_desc"`
	BucStatus int      `json:"buc_status"`
	BucTags   []string `json:"buc_tags"`
	DsId      int      `json:"ds_id"`
}

// 更新
type UpdateDomainReq struct {
	Guid      string   `json:"guid"`
	BucName   string   `json:"buc_name"`
	BucDesc   string   `json:"buc_desc"`
	BucStatus int      `json:"buc_status"`
	BucTags   []string `json:"buc_tags"`
	DsId      int      `json:"ds_id"`
}

// 查询
type FindDomainReq struct {
	PageSize int    `json:"page_size"`
	PageNum  int    `json:"page_num"`
	BucName  string `json:"buc_name"`
}

type FindDomainData struct {
	Id         int       `json:"id"`
	Guid       string    `json:"guid"`
	BucName    string    `json:"buc_name"`
	BucDesc    string    `json:"buc_desc"`
	BucStatus  int       `json:"buc_status"`
	BucTags    []string  `json:"buc_tags"`
	DsId       int       `json:"ds_id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type FindDomainResp struct {
	Data     []*FindDomainData `json:"data"`
	PageNum  int               `json:"page_num"`
	PageSize int               `json:"page_size"`
	Count    int               `json:"count"`
}

// 删除
type DeleteDomainReq struct {
	Guid string `json:"guid"`
}
