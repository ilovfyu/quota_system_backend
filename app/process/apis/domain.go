package apis

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"quota_system/app/process/dto"
	"quota_system/app/process/service"
	"quota_system/common"
)

func CreateDomainController(c *gin.Context) {
	var cbs dto.CreateDomainReq
	if err := c.ShouldBindJSON(&cbs); err != nil {
		zap.L().Error("request params parse failed, err: ",
			zap.String("method", "CreateDomainController"),
			zap.Error(err),
		)
		common.Failed(c, common.BadRequestError)
		return
	}
	err := service.CreateDomainSerivce(cbs)
	if err != nil {
		zap.L().Error("create buc domain service failed, err: ",
			zap.String("method", "CreateDomainController"),
			zap.Error(err),
		)
		common.Failed(c, common.IntervalError)
		return
	}
	common.Ok(c)
}

func QueryDomainController(c *gin.Context) {
	var req dto.FindDomainReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("request params parse failed, err: ",
			zap.String("method", "QueryDomainController"),
			zap.Error(err),
		)
		common.Failed(c, common.BadRequestError)
		return
	}
	err, resp := service.FindDomainService(req)
	if err != nil {
		zap.L().Error("query buc domain service failed, err: ",
			zap.String("method", "QueryDomainController"),
			zap.Error(err),
		)
		common.Failed(c, common.IntervalError)
		return
	}
	common.OkWithData(c, resp)
}

func DeleteDomainController(c *gin.Context) {
	var req dto.DeleteDomainReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("create buc domain service failed, err: ",
			zap.String("method", "DeleteDomainController"),
			zap.Error(err),
		)
		common.Failed(c, common.IntervalError)
		return
	}
	err := service.DeleteDomainService(req)
	if err != nil {
		zap.L().Error("query buc domain service failed, err: ",
			zap.String("method", "DeleteDomainController"),
			zap.Error(err),
		)
		common.Failed(c, common.IntervalError)
		return
	}
	common.Ok(c)
}

func UpdateDomainController(c *gin.Context) {
	var req dto.UpdateDomainReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("create buc domain service failed, err: ",
			zap.String("method", "UpdateDomainController"),
			zap.Error(err),
		)
		common.Failed(c, common.IntervalError)
		return
	}
	err := service.UpdateDomainService(req)
	if err != nil {
		zap.L().Error("query buc domain service failed, err: ",
			zap.String("method", "DeleteDomainController"),
			zap.Error(err),
		)
		common.Failed(c, common.IntervalError)
		return
	}
	common.Ok(c)
}
