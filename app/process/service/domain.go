package service

import (
	"context"
	"go.uber.org/zap"
	"quota_system/app/process/dto"
	"quota_system/dal/model"
	"quota_system/dal/query"
	"quota_system/utils"
	"strings"
)

func CreateDomainSerivce(req dto.CreateDomainReq) (err error) {
	// 创建业务流程
	q := query.QuotaBucInfo

	guid := utils.GeneratorGuid()

	// 组装buctags
	err = q.WithContext(context.Background()).Create(&model.QuotaBucInfo{
		GUID:      guid,
		BucName:   req.BucName,
		BucDesc:   req.BucDesc,
		BucStatus: int32(req.BucStatus),
		DsID:      int32(req.DsId),
		BucTags:   strings.Join(req.BucTags, ","),
	})
	return
}

func FindDomainService(req dto.FindDomainReq) (err error, resp *dto.FindDomainResp) {
	m := query.QuotaBucInfo
	q := m.WithContext(context.Background()).Where(m.BucStatus.Eq(0)).Where(m.IsDeleted.Eq(0))
	offset := (req.PageNum - 1) * req.PageSize
	result, count, err := q.FindByPage(offset, req.PageSize)
	if err != nil {
		zap.L().Error("FindDomainService, FindByPage failed", zap.Error(err))
		return err, nil
	}
	var datas []*dto.FindDomainData
	for _, v := range result {
		datas = append(datas, &dto.FindDomainData{
			Id:         int(v.ID),
			Guid:       v.GUID,
			BucName:    v.BucName,
			BucDesc:    v.BucDesc,
			BucStatus:  int(v.BucStatus),
			BucTags:    strings.Split(v.BucTags, ","),
			DsId:       int(v.DsID),
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		})
	}
	return nil, &dto.FindDomainResp{
		Data:     datas,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		Count:    int(count),
	}
}

func DeleteDomainService(req dto.DeleteDomainReq) (err error) {
	err = query.Q.Transaction(func(tx *query.Query) error {
		m := tx.QuotaBucInfo
		// 查找是否存在
		md, err := m.WithContext(context.Background()).Where(m.GUID.Eq(req.Guid)).Find()
		if err != nil {
			zap.L().Error("DeleteDomainService,query failed", zap.Error(err))
			return err
		}
		if len(md) == 0 {
			zap.L().Error("DeleteDomainService, Find buc failed", zap.Error(err))
			return err
		}
		if _, err = m.WithContext(context.Background()).Where(m.GUID.Eq(req.Guid)).Update(m.IsDeleted, 1); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		zap.L().Error("DeleteDomainService failed", zap.Error(err))
		return err
	}
	return nil
}

func UpdateDomainService(req dto.UpdateDomainReq) (err error) {
	err = query.Q.Transaction(func(tx *query.Query) error {
		q := tx.QuotaBucInfo
		_, err = q.WithContext(context.Background()).Where(q.GUID.Eq(req.Guid)).Updates(&model.QuotaBucInfo{
			BucName:   req.BucName,
			BucDesc:   req.BucDesc,
			BucStatus: int32(req.BucStatus),
			BucTags:   strings.Join(req.BucTags, ","),
			DsID:      int32(req.DsId),
		})
		return err
	})
	if err != nil {
		zap.L().Error("UpdateDomainService failed", zap.Error(err))
		return err
	}
	return nil
}
