package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

// Config 全局变量, 用来保存程序所有配置信息
var Config = new(Configure)

func Init(filePath string) {
	//1.指定配置文件路径
	if filePath != "" {
		viper.SetConfigFile(filePath)
	} else {
		viper.SetConfigFile("./config/config.dev.yaml")
	}
	//2.读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		//信息读取失败
		log.Fatalf("viper.ReadInConfig() failed: %v\n", err)
	}
	//3.将读取到的信息反序列化到全局变量Config
	if err = viper.Unmarshal(Config); err != nil {
		log.Fatalf("viper.Unmarshal() failed: %v\n", err)
	}
	//4.监听配置文件，当配置文件发生修改后立即更新配置文件信息到全局变量Config
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件修改了...\n")
		if err = viper.Unmarshal(Config); err != nil {
			log.Fatalf("viper.Unmarshal() failed: %v\n", err)
		}
	})
}

type Configure struct {
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
