package mysql

import (
	"gin-plus/config"
	"testing"
)

func TestInit(t *testing.T) {
	conf := &config.MySQLConfig{
		Host:              "172.16.114.128",
		Port:              3306,
		Database:          "blog",
		Username:          "root",
		Password:          "root",
		TablePrefix:       "blog_",
		SingularTable:     true,
		Charset:           "utf8mb4",
		DefaultStringSize: 171,
		ParseTime:         true,
		MaxIdleConns:      10,
		MaxOpenConns:      100,
		ConnMaxLifetime:   60,
	}
	err := Init(conf)
	if err != nil {
		t.Fatalf("数据库初始化连接失败，原因为: %v\n", err)
	}
	t.Logf("MySQL初始化连接配置正确")
}
