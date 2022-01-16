package config

import "time"

type Config struct {
	*AppConfig      `mapstructure:"app"`
	*ServerConfig   `mapstructure:"server"`
	*LogConfig      `mapstructure:"log"`
	*MySQLConfig    `mapstructure:"mysql"`
	*RedisConfig    `mapstructure:"redis"`
	*RabbitMQConfig `mapstructure:"rabbitmq"`
}

type AppConfig struct {
	Name            string        `mapstructure:"name"`
	Mode            string        `mapstructure:"mode"`
	URL             string        `mapstructure:"url"`
	Version         string        `mapstructure:"version"`
	StartTime       string        `mapstructure:"startTime"`
	MachineId       int64         `mapstructure:"machineId"`
	JWTSecret       string        `mapstructure:"jwtSecret"`
	JWTTokenExpired time.Duration `mapstructure:"jwtTokenExpired"`
}

type ServerConfig struct {
	Port         int           `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
}

type LogConfig struct {
	Filename   string `mapstructure:"filename"`
	Ext        string `mapstructure:"ext"`
	Level      string `mapstructure:"level"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxAge     int    `mapstructure:"maxAge"`
	MaxBackups int    `mapstructure:"maxBackups"`
	SavePath   string `mapstructure:"savePath"`
}

type MySQLConfig struct {
	Host              string        `mapstructure:"host"`
	Port              int           `mapstructure:"port"`
	Database          string        `mapstructure:"database"`
	Username          string        `mapstructure:"username"`
	Password          string        `mapstructure:"password"`
	TablePrefix       string        `mapstructure:"tablePrefix"`
	SingularTable     bool          `mapstructure:"singularTable"`
	Charset           string        `mapstructure:"charset"`
	DefaultStringSize uint          `mapstructure:"defaultStringSize"`
	ParseTime         bool          `mapstructure:"parseTime"`
	MaxIdleConns      int           `mapstructure:"maxIdleConns"`
	MaxOpenConns      int           `mapstructure:"maxOpenConns"`
	ConnMaxLifetime   time.Duration `mapstructure:"connMaxLifetime"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"poolSize"`
}

type RabbitMQConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
