package g

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"quota_system/model"
)

var (
	DB            *gorm.DB
	RDS           *redis.Client
	ServiceConfig *model.BaseConfig
)
