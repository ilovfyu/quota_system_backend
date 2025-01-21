package model

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	Database string `mapstructure:"database" json:"database"`
}

type RedisConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	DB   int    `mapstructure:"db" json:"db"`
}

type AppConfig struct {
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	Name        string `mapstructure:"name" json:"name"`
	Description string `mapstructure:"description" json:"description"`
}

type LoggerConfig struct {
	SavePath   string `mapstructure:"savepath" json:"savepath"`
	Prefix     string `mapstructure:"prefix" json:"prefix"`
	MaxSize    int    `mapstructure:"maxsize" json:"maxsize"`
	MaxBackups int    `mapstructure:"maxbackups" json:"maxbackups"`
	MaxAge     int    `mapstructure:"maxage" json:"maxage"`
	Compress   bool   `mapstructure:"compress" json:"compress"`
	Level      int8   `mapstructure:"level" json:"level"`
}

type BaseConfig struct {
	ServiceConfig *AppConfig    `mapstructure:"app" json:"app"`
	RdsConfig     *RedisConfig  `mapstructure:"redis" json:"redis"`
	DBConfig      *MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	LoggerConfig  *LoggerConfig `mapstructure:"logger" json:"logger"`
}
