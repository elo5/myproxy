package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

var (
	// 配置 文件 路径
	configPath string
)

// Config Config
type Config struct {
	ListenAddr string `json:"listen"`
	RemoteAddr string `json:"remote"`
	Password   string `json:"password"`
}

func init() {
	home, _ := homedir.Dir()
	// 默认的配置文件名称
	configFileName := ".myproxy.json"
	// 如果用户有传配置文件，就使用用户传入的配置文件
	if len(os.Args) == 2 {
		configFileName = os.Args[1]
	}
	configPath = path.Join(home, configFileName)
}

// SaveConfig 保存config
func (config *Config) SaveConfig() {
	configJSON, _ := json.MarshalIndent(config, "", "	")
	err := ioutil.WriteFile(configPath, configJSON, 0644)
	if err != nil {
		// fmt.Errorf("保存配置到文件 %s 出错: %s", configPath, err)
		log.Printf("保存配置到文件 %s 出错: %s \n", configPath, err)
	}
	log.Printf("保存配置到文件 %s 成功\n", configPath)
}

// ReadConfig Read Config
func (config *Config) ReadConfig() {
	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		log.Printf("从文件 %s 中读取配置\n", configPath)
		file, err := os.Open(configPath)
		if err != nil {
			log.Fatalf("打开配置文件 %s 出错:%s  \n", configPath, err)
		}
		defer file.Close()

		err = json.NewDecoder(file).Decode(config)
		if err != nil {
			log.Fatalf("格式不合法的 JSON 配置文件:\n%s  ", file.Name())
		}
	}
}
