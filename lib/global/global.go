package global

import (
	"github.com/spf13/viper"
	"goBlog/lib/config/model"
)

var (
	GVA_CONFIG model.ConfigModel
	GVA_VP     *viper.Viper
)
