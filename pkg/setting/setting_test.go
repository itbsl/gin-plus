package setting

import "testing"

func TestInit(t *testing.T) {
	Init("./../../config/config.dev.yaml")
	if Config.AppConfig == nil {
		t.Fatalf("配置文件读取成功，反序列化失败")
	}
	if Config.AppConfig.Name != "gin-plus" {
		t.Fatalf("配置文件读取单元测试失败,读取的值为：%v,目标值为: gin-plus", Config.AppConfig.Name)
	}
	if Config.ServerConfig.Port != 8080 {
		t.Fatalf("配置文件读取单元测试失败,读取的值为：%v,目标值为: gin-plus", Config.ServerConfig.Port)
	}
	if Config.LogConfig.Filename != "app" {
		t.Fatalf("配置文件读取单元测试失败,读取的值为：%v,目标值为: gin-plus", Config.LogConfig.Filename)
	}
	t.Logf("配置文件初始化正确")
}
