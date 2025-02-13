package initialize

import (
	_ "github.com/go-sql-driver/mysql"
	"goBlog/lib/config/model"
	"goBlog/lib/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

// GormMysql 初始化Mysql数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //配置日志级别，打印出所有的sql
	}); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		// 连接最大存活时间
		sqlDB.SetConnMaxLifetime(time.Minute * 3)
		//空闲连接最大存活时间
		sqlDB.SetConnMaxIdleTime(time.Minute * 1)
		err = sqlDB.Ping()
		if err != nil {
			log.Println("数据库无法连接")
			_ = sqlDB.Close()
			panic(err)
		}
		return db
	}
}

// GormMysqlByConfig 初始化Mysql数据库用过传入配置
func GormMysqlByConfig(m model.Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //配置日志级别，打印出所有的sql
	}); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		// 连接最大存活时间
		sqlDB.SetConnMaxLifetime(time.Minute * 3)
		//空闲连接最大存活时间
		sqlDB.SetConnMaxIdleTime(time.Minute * 1)
		err = sqlDB.Ping()
		if err != nil {
			log.Println("数据库无法连接")
			_ = sqlDB.Close()
			panic(err)
		}
		return db
	}
}
