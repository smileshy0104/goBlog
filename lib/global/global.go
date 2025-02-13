package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"goBlog/lib/config/model"
	"gorm.io/gorm"
)

var (
	GVA_DB            *gorm.DB
	GVA_CONFIG        model.ConfigModel
	GVA_VP            *viper.Viper
	GVA_ACTIVE_DBNAME *string
	GVA_REDIS         redis.UniversalClient
	GVA_OUTPUT        string
)
