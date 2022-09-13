package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v2"
)

var config *Configuration
var _cfgInit = sync.Once{}

const configFileName = "conf.yaml"

type Configuration struct {
	Debug                 bool   `yaml:"debug"`
	DbConnectionString    string `yaml:"dbstr"`
	RedisConnectionString string `yaml:"redis_str"`
	Port                  string `yaml:"port"`
	APIBaseURL            string `yaml:"api_base_url"`
	FileServerDir         string `yaml:"file_srv_dir"`
	FileServerBaseUrl     string `yaml:"file_srv_base_url"`
}

func GetConfig() Configuration {
	_cfgInit.Do(func() {
		filename, err := fileName()
		if err != nil {
			panic(err)
		}

		data, err := os.ReadFile(filename)
		if err != nil {
			panic("读取配置文件失败\t" + err.Error())
		}

		configuration := Configuration{}
		err = yaml.Unmarshal(data, &configuration)
		if err != nil {
			panic("解析配置文件失败\t" + err.Error())
		}

		config = &configuration
	})

	return *config
}

func fileName() (string, error) {
	AppPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	WorkPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var appConfigPath = filepath.Join(AppPath, configFileName)
	if _, err = os.Stat(appConfigPath); err != nil {
		appConfigPath = filepath.Join(WorkPath, configFileName)
		if _, err = os.Stat(appConfigPath); err != nil {
			return "", fmt.Errorf("配置文件不存在: %s", appConfigPath)
		}

		return appConfigPath, err
	}

	return appConfigPath, nil
}
