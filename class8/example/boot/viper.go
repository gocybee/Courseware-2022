package boot

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	g "main/app/global"
	"os"
)

const (
	configEnv  = "ZHIHU_CONFIG_PATH"           // 预定义环境变量
	configFile = "manifest/config/config.yaml" // 预定义配置文件位置
)

func ViperSetup(path ...string) {
	var configPath string

	// 获取配置文件路径
	// 优先级: 参数 > 命令行 > 环境变量 > 默认值
	if len(path) != 0 {
		// 参数
		configPath = path[0]
	} else {
		// 命令行
		flag.StringVar(&configPath, "c", "", "set config path")
		flag.Parse()

		if configPath == "" {
			if configPath = os.Getenv(configEnv); configPath != "" {
				// 环境变量
			} else {
				// 默认值
				configPath = configFile
			}
		}
	}

	fmt.Printf("get config path: %s", configPath)

	v := viper.New()
	v.SetConfigFile(configPath) // 设置配置文件路径
	v.SetConfigType("yaml")     // 设置配置文件类型
	err := v.ReadInConfig()     // 读取配置文件
	if err != nil {
		panic(fmt.Errorf("get config file failed, err: %v", err))
	}

	if err = v.Unmarshal(&g.Config); err != nil {
		// 将配置文件反序列化到 Config 结构体
		panic(fmt.Errorf("unmarshal config failed, err: %v", err))
	}
}
