package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

func LoadConfigFromToml(filePath string) error {
	config = NewDefaultConfig()
	// 读取Toml格式的配置
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from file error, path:%s, %s", filePath, err)
	}
	return nil
}

func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	err := env.Parse(config)
	if err != nil {
		return err
	}
	return nil
}
