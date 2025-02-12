package model

type SystemConfig struct {
	DbType          string `mapstructure:"db-type" json:"db-type" yaml:"db-type"` // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}
