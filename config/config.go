package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config represents the configuration
type ConfigServer struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
}

// Config represents the configuration
type ConfigGateway struct {
	Apis struct {
		Name        string   `json:"name"`
		ListenPath  string   `json:"listenpath"`
		UpstreamURL string   `json:"upstreamurl"`
		Active      bool     `json:"active"`
		Plugins     []string `json:"plugins"`
	} `json:"apis"`
	Blacklist []string `json:"blacklist"`
}

// Configuration is the actual configuration for the project
var ServerConfiguration ConfigServer
var GatewayConfiguration ConfigGateway

// Parse takes the path of a configuration and makes it to an actual Config
func Parse(serverPath, gatewayPath string) error {

	file, err := os.Open(serverPath)
	if err != nil {
		return err
	}

	err = json.NewDecoder(file).Decode(&ServerConfiguration)
	if err != nil {
		return err
	}

	// file, err = os.Open(gatewayPath)
	// if err != nil {
	// 	return err
	// }

	// err = json.NewDecoder(file).Decode(&GatewayConfiguration)
	// if err != nil {
	// 	return err
	// }

	log.Println("Successfully read configuration at: %s and %s", serverPath, gatewayPath)
	return nil
}
