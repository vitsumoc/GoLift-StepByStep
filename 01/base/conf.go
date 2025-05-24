package base

import (
	"os"

	"gopkg.in/yaml.v3"
)

// 表示配置结构
type Config struct {
	Http struct {
		Port int
	}
	Db struct {
		Path string
	}
}

// 配置全局对象
var Conf *Config

// 初始化配置函数
func InitConfig(path string) {
	Conf = &Config{}

	// 读取配置文件，获得配置内容
	fileData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// 将 yml 格式的文件内容转为配置对象
	err = yaml.Unmarshal(fileData, &Conf)
	if err != nil {
		panic(err)
	}
}
