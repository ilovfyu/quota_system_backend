package common

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"quota_system/dal/query"
	"quota_system/global"
)

func InitMySQL() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		g.ServiceConfig.DBConfig.Username,
		g.ServiceConfig.DBConfig.Password,
		g.ServiceConfig.DBConfig.Host,
		g.ServiceConfig.DBConfig.Port,
		g.ServiceConfig.DBConfig.Database,
	)
	dsnWithZone := dsn + "&loc=Asia%2FShanghai"
	db, err := gorm.Open(mysql.Open(dsnWithZone), &gorm.Config{})
	if err != nil {
		zap.L().Error("init mysql failed", zap.Error(err))
		panic(err)
	}
	g.DB = db
	query.SetDefault(db)
}
