package setting

import (
	"gin-plus/global"
	"testing"
)

func TestInit(t *testing.T) {
	Init("./setting_test.yaml")
	if global.Config.AppConfig == nil {
		t.Fatalf("配置文件读取成功，反序列化失败")
	}
	if global.Config.AppConfig.Name != "gin-plus" {
		t.Fatalf("配置文件读取单元测试失败,读取的值为：%v,目标值为: gin-plus", global.Config.AppConfig.Name)
	}
	if global.Config.ServerConfig.Port != 8080 {
		t.Fatalf("配置文件读取单元测试失败,读取的值为：%v,目标值为: gin-plus", global.Config.ServerConfig.Port)
	}
	if global.Config.LogConfig.Filename != "app" {
		t.Fatalf("配置文件读取单元测试失败,读取的值为：%v,目标值为: gin-plus", global.Config.LogConfig.Filename)
	}
	t.Logf("配置文件初始化正确")
}
