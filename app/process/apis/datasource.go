package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"quota_system/app/process/dto"
	"quota_system/app/process/service"
	"quota_system/common"
)

func CreateDataSourceController(c *gin.Context) {
	var req dto.CreateDataSourceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("create data source with invalid param",
			zap.String("method", "CreateDataSourceController"),
			zap.Error(err),
		)
		common.Failed(c, common.BadRequestError)
		return
	}

	fmt.Printf("%+v\n", req)

	err := service.CreateDataSourceService(req)
	if err != nil {
		zap.L().Error("create data source service failed",
			zap.String("method", "CreateDataSourceController"),
			zap.Error(err),
		)
		common.Failed(c, common.IntervalError)
		return
	}
	common.Ok(c)
}

func DeleteDataSourceController(c *gin.Context) {
	var req dto.DeleteDataSourceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("create data source with invalid param",
			zap.String("method", "DeleteDataSourceController"),
			zap.Error(err),
		)
		common.Failed(c, common.BadRequestError)
		return
	}
	err := service.DeleteDataSourceService(req)
	if err != nil {
		zap.L().Error("create data source service failed",
			zap.String("method", "CreateDataSourceController"),
			zap.Error(err),
		)
		common.Failed(c, common.IntervalError)
		return
	}
	common.Ok(c)
}

func GetDataSourceListController(c *gin.Context) {
	var req dto.QueryDataSourceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("create data source with invalid param",
			zap.String("method", "GetDataSourceListController"),
			zap.Error(err),
		)
		common.Failed(c, common.BadRequestError)
		return
	}
	resp, err := service.QueryDataSourceService(req)
	if err != nil {
		zap.L().Error("create data source service failed",
			zap.String("method", "GetDataSourceListController"),
			zap.Error(err),
		)
		common.Failed(c, common.IntervalError)
		return
	}
	common.OkWithData(c, resp)
}

func UpdateDataSourceController(c *gin.Context) {
	var req dto.UpdateDataSourceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("create data source with invalid param",
			zap.String("method", "UpdateDataSourceController"),
			zap.Error(err),
		)
		common.Failed(c, common.BadRequestError)
		return
	}
	err := service.UpdateDataSourceService(req)
	if err != nil {
		zap.L().Error("create data source service failed",
			zap.String("method", "UpdateDataSourceController"),
			zap.Error(err),
		)
		common.Failed(c, common.IntervalError)
		return
	}
	common.Ok(c)
}

// 连通性测试
func ConnectionTestDataSourceController(c *gin.Context) {
	var req dto.TestDataSourceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("create data source with invalid param",
			zap.String("method", "UpdateDataSourceController"),
			zap.Error(err),
		)
		common.Failed(c, common.BadRequestError)
		return
	}
	ping, err := service.TestDataSourceService(req)
	if err != nil {
		zap.L().Error("create data source service failed",
			zap.String("method", "TestDataSourceController"),
			zap.Error(err),
		)
		common.Failed(c, common.IntervalError)
		return
	}
	if !ping {
		common.Failed(c, common.ConnectionDBError)
	}
	common.Ok(c)
}
