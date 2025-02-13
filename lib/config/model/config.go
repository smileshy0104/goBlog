package model

// 相关配置结构体
type ConfigModel struct {
	JWT    JWT          `json:"jwt" yaml:"jwt"`
	Viewer Viewer       `json:"viewer" yaml:"viewer"`
	System SystemConfig `json:"system" yaml:"system"`
	Mysql  Mysql        `json:"mysql" yaml:"mysql"`
	Redis  Redis        `json:"redis" yaml:"redis"`
}
