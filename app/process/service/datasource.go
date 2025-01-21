package service

import (
	"context"
	"encoding/json"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"quota_system/app/process/dto"
	"quota_system/dal/model"
	"quota_system/dal/query"
	"quota_system/utils"
)

func CreateDataSourceService(req dto.CreateDataSourceReq) (err error) {
	err = query.Q.Transaction(func(tx *query.Query) error {
		q := tx.QuotaDatasourceInfo
		// password加密
		if _, ok := req.Connect["password"]; ok {
			password, err := utils.HashPassword(cast.ToString(req.Connect["password"]))
			if err != nil {
				return err
			}
			req.Connect["password"] = password
		}
		jsonStr, _ := json.Marshal(req.Connect)
		err = q.WithContext(context.Background()).Create(&model.QuotaDatasourceInfo{
			Name:    req.DsName,
			Type:    req.DsType,
			Desc:    req.Desc,
			Connect: string(jsonStr),
		})
		return err
	})
	return err
}

func DeleteDataSourceService(req dto.DeleteDataSourceReq) (err error) {
	err = query.Q.Transaction(func(tx *query.Query) error {
		q := tx.QuotaDatasourceInfo
		for _, record := range req.Ids {
			q.WithContext(context.Background()).Where(q.ID.Eq(record)).Update(q.IsDeleted, 1)
		}
		return nil
	})
	return err
}

func QueryDataSourceService(req dto.QueryDataSourceReq) (resp *dto.QueryDataSourceResp, err error) {
	m := query.QuotaDatasourceInfo
	q := m.WithContext(context.Background()).Where(m.IsDeleted.Eq(0))
	offset := (req.PageNum - 1) * req.PageSize
	result, count, err := q.FindByPage(offset, req.PageSize)
	if err != nil {
		zap.L().Error("query datasource failed", zap.Error(err))
		return nil, err
	}
	var datas []*dto.DataSourceData
	for _, record := range result {
		datas = append(datas, &dto.DataSourceData{
			Id:         record.ID,
			DsName:     record.Name,
			DsType:     record.Type,
			Desc:       record.Desc,
			Connect:    record.Connect,
			CreateTime: record.CreateTime,
			UpdateTime: record.UpdateTime,
		})
	}
	return &dto.QueryDataSourceResp{
		Data:     datas,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		Count:    int(count),
	}, nil
}

func UpdateDataSourceService(req dto.UpdateDataSourceReq) (err error) {
	err = query.Q.Transaction(func(tx *query.Query) error {
		q := tx.QuotaDatasourceInfo
		connect, err := json.Marshal(req.Connect)
		if err != nil {
			zap.L().Error("json marshal connect failed", zap.Error(err))
			return err
		}
		_, err = q.WithContext(context.Background()).Where(q.ID.Eq(req.Id)).Updates(&model.QuotaDatasourceInfo{
			Name:    req.DsName,
			Type:    req.DsType,
			Desc:    req.Desc,
			Connect: string(connect),
		})
		if err != nil {
			zap.L().Error("update datasource failed", zap.Error(err))
			return err
		}
		return nil
	})
	return err
}

// 测试数据源是否连通
func TestDataSourceService(req dto.TestDataSourceReq) (res bool, err error) {
	// 查询数据源
	q := query.QuotaDatasourceInfo
	datasourceInfo, err := q.WithContext(context.Background()).Where(q.ID.Eq(req.Id)).First()
	if err != nil {
		zap.L().Error("query datasource failed", zap.Error(err))
		return false, err
	}
	// 构造连接信息
	connectionMeta := utils.ConnectionConfig{
		Name:    datasourceInfo.Name,
		Connect: datasourceInfo.Connect,
		Type:    datasourceInfo.Type,
	}
	// 调用测试
	if ok, err := utils.TestConnection(connectionMeta); !ok {
		return false, err
	}
	return true, nil
}
