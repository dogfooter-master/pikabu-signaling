package service

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

type HostsConfig struct {
	GrpcHosts string `mapstructure:"grpc_hosts"`
}
type ServerConfig struct {
	Signaling HostsConfig `mapstructure:"signaling"`
}

var serverHostConfig ServerConfig

func init() {
	if err := LoadConfig(); err != nil {
		panic(err)
	}

	viper.GetStringMap("server")
	_ = viper.UnmarshalKey("server", &serverHostConfig)

	//viper.Debug()
}
func LoadConfig() (err error) {
	viper.SetConfigFile(os.Getenv("PIKABU_HOME") + "/config/config.json")
	if err = viper.ReadInConfig(); err != nil {
		viper.SetConfigFile(os.Getenv("PIKABU_HOME") + "/pikabu-signaling" + "/config/config.json")
		if err = viper.ReadInConfig(); err != nil {
			return
		}
		return
	}
	return
}
func GetConfigServerGrpc() string {
	if strings.Contains(serverHostConfig.Signaling.GrpcHosts, "PORT") {
		port := "17092"
		hosts := strings.Replace(serverHostConfig.Signaling.GrpcHosts, "PORT", port, 1)
		return hosts
	} else {
		return serverHostConfig.Signaling.GrpcHosts
	}
}