package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"goBlog/lib/config/model"
	"gorm.io/gorm"
	"log"
)

var (
	GVA_DB            *gorm.DB
	GVA_CONFIG        model.ConfigModel
	GVA_VP            *viper.Viper
	GVA_ACTIVE_DBNAME *string
	GVA_REDIS         redis.UniversalClient
	GVA_OUTPUT        string
	GVA_Logger        *log.Logger // 日志记录器

)
