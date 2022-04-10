package etc

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Config ServerConfig

type ServerConfig struct {
	Server Server `yaml:"server"`
	Log    Logger `yaml:"logger"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
}

type Server struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	HttpPort string `yaml:"httpport"`
	GrpcPort string `json:"grpcport"`
}

type Logger struct {
	Logdir   string `yaml:"logdir"`
	FileName string `yaml:"filename"`
	Level    string `yaml:"level"`
	Format   string `yaml:"format"` // text or json
}

type Mysql struct {
	IP           string `yaml:"ip"`
	Port         string `yaml:"port"`
	Database     string `yaml:"database"`
	Charset      string `yaml:"charset"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	MaxIdleConns int    `yaml:"maxidleconns"`
	MaxOpenConns int    `yaml:"maxopenconns"`
}

type Redis struct {
	Address      string `yaml:"address"`
	Password     string `yaml:"password"`
	Db           int    `yaml:"db"`
	PoolSize     int    `yaml:"poolsize"`
	MinIdleConns int    `yaml:"minidleconns"`
}

func InitConfig(path string) {
	configPath := "conf/config.yaml"
	if path != "" {
		configPath = path
	}

	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configFile, &Config)
	if err != nil {
		panic(err)
	}
}
