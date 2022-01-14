package setting

import (
	"fmt"
	"gin-plus/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

func Init(filePath string) {
	//1.指定配置文件路径
	//(1)如果指定了配置文件，优先使用显式指定的值
	//(2)如果没有显式指定配置文件，通过系统环境变量获取当前所处的运行模式(通过设置测试和线上运行环境变量能使得线上线下配置隔离)
	//不用修改任何代码而且线上和线下的配置文件隔离开
	//(3)如果既没有显式指定配置文件路径，也没有设置系统环境变量，则从默认的配置文件路径找
	if filePath != "" {
		viper.SetConfigFile(filePath)
	} else if os.Getenv("gin_plus_mode") != "" {
		//也可以直接使用viper封装的方法
		//viper.AutomaticEnv()
		//mode := viper.GetString("gin_plus_mode")
		mode := os.Getenv("gin_plus_mode")
		viper.SetConfigFile("./config/config." + mode + ".yaml")
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
	if err = viper.Unmarshal(global.Config); err != nil {
		log.Fatalf("viper.Unmarshal() failed: %v\n", err)
	}
	//4.监听配置文件，当配置文件发生修改后立即更新配置文件信息到全局变量Config
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件%s修改了...\n", in.Name)
		if err = viper.Unmarshal(global.Config); err != nil {
			log.Fatalf("viper.Unmarshal() failed: %v\n", err)
		}
	})
}
