package redis

import (
	"gin-plus/config"
	"testing"
)

func TestInit(t *testing.T) {
	conf := &config.RedisConfig{
		Host:     "172.16.114.128",
		Port:     6379,
		Password: "",
		DB:       0,
		PoolSize: 0,
	}
	if err := Init(conf); err != nil {
		t.Fatalf("Redis初始化连接失败: %v\n", err)
	}
	t.Logf("Redis初始化连接成功")
}
