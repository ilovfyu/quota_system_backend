package app

import (
	"github.com/gin-gonic/gin"
	"quota_system/app/process/apis"
)

func InitService(e *gin.Engine) {

	// business operation
	buc_router := e.Group("/api/v1/buc")

	buc_router.POST("", apis.CreateDomainController)
	buc_router.GET("", apis.QueryDomainController)
	buc_router.DELETE("", apis.DeleteDomainController)

	// datasource operation
	datasource_router := e.Group("/api/v1/datasource")
	datasource_router.POST("", apis.CreateDataSourceController)
	datasource_router.GET("", apis.GetDataSourceListController)
	datasource_router.DELETE("", apis.DeleteDataSourceController)
	datasource_router.PUT("", apis.UpdateDataSourceController)
	datasource_router.GET("/test", apis.ConnectionTestDataSourceController)

	// dataset operation
	dataset_router := e.Group("/api/v1/dataset")
	dataset_router.POST("", apis.CreateDataSetController)
	dataset_router.GET("", apis.QueryDataSetController)
	dataset_router.DELETE("", apis.DeleteDataSetController)
	dataset_router.PUT("", apis.UpdateDataSetController)
}
