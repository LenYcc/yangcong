/**
 * @Author: dengmingcong
 * @Description:
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2022/01/25 6:38 下午
 */

package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

const DEFAULT_CONFIG_PATH = "/Users/dengmingcong/go/src/yangcong/config/config_json"

var globalConfig *Config
var once sync.Once

type configInterface interface{}

type configMap map[string]interface{}

type Config struct {
	configs map[string]*configMap
}

func GetConfig(path string) *Config{
	once.Do(func() {
		if path == "" {
			path = DEFAULT_CONFIG_PATH
		}
		globalConfig = loadConfig(path)
	})
	return globalConfig
}

func loadConfig(path string) *Config{
	config := Config{configs: make(map[string]*configMap)}
	configDir, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, fileInfo := range configDir {
		if !fileInfo.IsDir() {
			config.configs[fileInfo.Name()] = &configMap{}
			var file *os.File
			file, err = os.Open(path + "/" + fileInfo.Name())
			if err != nil {
				panic(err)
			}
			jsonReader := bufio.NewReader(file)
			decoder := json.NewDecoder(jsonReader)
			err = decoder.Decode(config.configs[fileInfo.Name()])
			if err != nil {
				fmt.Println("Decoder failed", err.Error())
			} else {
				fmt.Println("Decoder success")
				fmt.Println(config.configs[fileInfo.Name()])
			}
		}else{
			panic(fmt.Errorf("Path is not dir,please select another config dir,path[%v]", path))
		}
	}
	return &config
}

