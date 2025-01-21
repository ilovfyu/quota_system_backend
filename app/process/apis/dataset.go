package apis

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"quota_system/app/process/dto"
	"quota_system/app/process/service"
	"quota_system/common"
)

func CreateDataSetController(c *gin.Context) {
	var req dto.CreateDataSetReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("create data source with invalid param",
			zap.String("method", "CreateDataSetController"),
			zap.Error(err),
		)
		common.Failed(c, common.BadRequestError)
		return
	}
	if err := service.CreateDataSetService(req); err != nil {
		zap.L().Error("create data source service failed",
			zap.String("method", "CreateDataSetController"),
			zap.Error(err),
		)
		common.Failed(c, common.IntervalError)
		return
	}
	common.Ok(c)
}

func DeleteDataSetController(c *gin.Context) {
	var req dto.DeleteDataSetReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("delete data set with invalid param",
			zap.String("method", "DeleteDataSetController"),
			zap.Error(err),
		)
		common.Failed(c, common.BadRequestError)
		return
	}
	err := service.DeleteDataSetService(req)
	if err != nil {

	}
	common.Ok(c)
}

func QueryDataSetController(c *gin.Context) {
	var req dto.QueryDataSetReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("create data set with invalid param",
			zap.String("method", "QueryDataSetController"),
			zap.Error(err),
		)
		common.Failed(c, common.BadRequestError)
		return
	}
	resp, err := service.QueryDataSetService(req)
	if err != nil {
		zap.L().Error("query dataset service failed",
			zap.String("method", "QueryDataSetController"),
			zap.Error(err),
		)
		common.Failed(c, common.IntervalError)
		return
	}
	common.OkWithData(c, resp)
}

func UpdateDataSetController(c *gin.Context) {
	var req dto.UpdateDataSetReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("update dataset invalid param",
			zap.String("method", "UpdateDataSetController"),
			zap.Error(err),
		)
		common.Failed(c, common.BadRequestError)
		return
	}
	err := service.UpdateDataSetService(req)
	if err != nil {
		zap.L().Error("update dataset service failed",
			zap.String("method", "UpdateDataSetController"),
			zap.Error(err),
		)
		common.Failed(c, common.IntervalError)
		return
	}
	common.Ok(c)
}
