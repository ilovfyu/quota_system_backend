package service

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"quota_system/app/process/dto"
	"quota_system/dal/model"
	"quota_system/dal/query"
)

func CreateDataSetService(req dto.CreateDataSetReq) (err error) {
	q := query.QuotaDatasetInfo
	jsonStr, _ := json.Marshal(req.Meta)
	// 根据数据源查询数据集类型
	s := query.QuotaDatasourceInfo
	sourceInfo, err := s.WithContext(context.Background()).Where(s.ID.Eq(req.SourceId)).First()
	if err != nil {
		zap.L().Error("query datasource failed", zap.Int32("sourceId", req.SourceId), zap.Error(err))
		return err
	}
	err = q.WithContext(context.Background()).Create(&model.QuotaDatasetInfo{
		DsName:   req.DsName,
		DsType:   sourceInfo.Type,
		DsDesc:   req.DsDesc,
		SourceID: req.SourceId,
		Meta:     string(jsonStr),
		CreateBy: "ilovfyu",
	})
	if err != nil {
		zap.L().Error("CreateDataSetService failed", zap.Error(err))
		return err
	}
	return nil
}

func QueryDataSetService(req dto.QueryDataSetReq) (resp *dto.QueryDataSetResp, err error) {
	m := query.QuotaDatasetInfo
	q := m.WithContext(context.Background()).Where(m.IsDeleted.Eq(0))
	if len(req.DsName) > 0 {
		q.Where(m.DsName.Like("%" + req.DsName + "%"))
	}
	if len(req.DsType) > 0 {
		q.Where(m.DsType.Like("%" + req.DsType + "%"))
	}
	offset := (req.PageNum - 1) * req.PageSize
	result, count, err := q.FindByPage(offset, req.PageSize)
	if err != nil {
		zap.L().Error("QueryDataSetService failed", zap.Error(err))
		return nil, err
	}
	var datas []*dto.DataSetResult
	for _, v := range result {
		var meta []*dto.DataSetData
		err := json.Unmarshal([]byte(v.Meta), &meta)
		if err != nil {
			zap.L().Error("QueryDataSetService failed", zap.Error(err))
			return nil, err
		}
		datas = append(datas, &dto.DataSetResult{
			Id:         int(v.ID),
			DsName:     v.DsName,
			DsType:     v.DsType,
			DsDesc:     v.DsDesc,
			SourceId:   v.SourceID,
			Meta:       meta,
			CreateBy:   v.CreateBy,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		})
	}
	return &dto.QueryDataSetResp{
		Data:     datas,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Count:    int(count),
	}, nil
}

func DeleteDataSetService(req dto.DeleteDataSetReq) (err error) {
	err = query.Q.Transaction(func(tx *query.Query) error {
		q := query.QuotaDatasetInfo
		for _, id := range req.Ids {
			_, err = q.WithContext(context.Background()).Where(q.ID.Eq(id)).Update(q.IsDeleted, 1)
			if err != nil {
				zap.L().Error("DeleteDataSetService failed", zap.Error(err))
				return err
			}
		}
		return nil
	})
	return
}

func UpdateDataSetService(req dto.UpdateDataSetReq) (err error) {
	q := query.QuotaDatasetInfo
	datasetInfo, err := q.WithContext(context.Background()).Where(q.ID.Eq(int32(req.Id))).First()
	if err != nil {
		zap.L().Error("UpdateDataSetService failed", zap.Error(err))
		return
	}
	if datasetInfo.ID <= 0 {
		zap.L().Error("UpdateDataSetService failed", zap.Error(err))
		return
	}
	// 更新
	err = query.Q.Transaction(func(tx *query.Query) error {
		jsonStr, _ := json.Marshal(req.Meta)
		_, err = q.WithContext(context.Background()).Where(q.ID.Eq(int32(req.Id))).Where(q.IsDeleted.Eq(0)).Updates(&model.QuotaDatasetInfo{
			DsName: req.DsName,
			DsDesc: req.DsDesc,
			Meta:   string(jsonStr),
		})
		return err
	})
	return
}
