package config

import (
	"encoding/json"
	"log"
	"os"
)

var config Config

// Config struct to hold the configuration data
type Config struct {
	Nodes []Node `json:"nodes"`
}

type Node struct {
	ChainId   string   `json:"chain_id"`
	ChainName string   `json:"chain_name"`
	RpcAddr   []string `json:"rpc_addr"`
}

func init() {
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
}

func GetConfig() *Config {
	return &config
}
