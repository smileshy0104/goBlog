package initialize

import (
	"goBlog/lib/global"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		global.GVA_ACTIVE_DBNAME = &global.GVA_CONFIG.Mysql.Dbname
		return GormMysql()
	default:
		global.GVA_ACTIVE_DBNAME = &global.GVA_CONFIG.Mysql.Dbname
		return GormMysql()
	}
}
