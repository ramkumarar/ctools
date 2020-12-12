package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

//ConsulConfig stores consul configuration
type ConsulConfig []struct {
	Space       string `json:"space"`
	Environment string `json:"environment"`
	Readkey     string `json:"readkey"`
	Writekey    string `json:"writekey"`
}

//EnvConsulKeyMap maps environment string to consul keys
type EnvConsulKeyMap map[string]ConsulKey

//ServiceConsulSpaceMap maps services to Consul spaces
type ServiceConsulSpaceMap map[string]string

//ConsulKey stores consul configuration
type ConsulKey struct {
	Readkey  string `json:"readkey"`
	Writekey string `json:"writekey"`
}

const (
	envNSWithSubEnvAndBrand = "/%s/%s/%s/"
	envNSWithdBrand         = "/%s/%s/"
)

var envConsulKeyMap = make(EnvConsulKeyMap)

//GetSpace returns consul space fo a given service
func GetSpace(service string) (string, error) {

	file, _ := ioutil.ReadFile("servicespacemapping.json")
	serviceConsulSpaceMap := ServiceConsulSpaceMap{}

	_ = json.Unmarshal([]byte(file), &serviceConsulSpaceMap)

	value, ok := serviceConsulSpaceMap[service]

	if !ok {
		return "", errors.New("Space for service not configured")
	}
	return value, nil
}

//GetConsulKey returns consul keys fo a given space and environment
func GetConsulKey(space string, environment string) (ConsulKey, error) {
	file, _ := ioutil.ReadFile("consulconfig.json")
	consulConfig := ConsulConfig{}
	_ = json.Unmarshal([]byte(file), &consulConfig)

	for _, config := range consulConfig {
		key := fmt.Sprintf("%s-%s", config.Space, config.Environment)
		value := ConsulKey{
			Readkey:  config.Readkey,
			Writekey: config.Writekey,
		}
		envConsulKeyMap[key] = value
	}

	key := fmt.Sprintf("%s-%s", space, environment)

	value, ok := envConsulKeyMap[key]

	if !ok {
		return ConsulKey{}, errors.New("Consul Keys for the space not configured")
	}
	return value, nil
}

func getEnvironmentNameSpace(environment string, subenvironment string, brand string) string {

	if brand == "grp" {
		envNamespace := fmt.Sprintf("/%s/%s/", environment, subenvironment)
		return envNamespace
	}

	if environment == "nft" {
		envNamespace := fmt.Sprintf("/%s/%s/", environment, brand)
		return envNamespace
	}

	envNamespace := fmt.Sprintf("/%s/%s/%s/", environment, subenvironment, brand)
	return envNamespace

}
