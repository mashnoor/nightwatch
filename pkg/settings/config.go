package settings

import (
	"flag"
	"github.com/mashnoor/nightwatch/pkg/strcts"
	"gopkg.in/yaml.v2"
	"os"
)

var (
	SystemAppConfig strcts.AppConfig
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readConfigFile() string {
	var configPath string
	flag.StringVar(&configPath, "config", "", "a string var")
	flag.Parse()

	if configPath == "" {
		panic("config file location must be provided")
	}

	dat, err := os.ReadFile(configPath)
	check(err)

	return string(dat)
}

func LoadAppConfig() {
	config := readConfigFile()
	//fmt.Println(config)
	err := yaml.Unmarshal([]byte(config), &SystemAppConfig)
	check(err)
	//validateConfig()
}
