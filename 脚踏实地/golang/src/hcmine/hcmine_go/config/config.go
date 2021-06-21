package config

import "hcmine/hcmine_go/util"

type Config struct {
	SyscallOn    int
	ConntrackOn  int
	SuppressPort []int
	//Logging logp.Logging

	ClientType string           `yaml:"client_type"`
	MqttConfig MqttClientConfig `yaml:"mqtt_client"`
	EsConfig   EsClientConfig   `yaml:"elasticsearch_client"`

	KubeClusterId string `yaml:"kube_clusterid"`
	KubeAuthType  string `yaml:"kubeauth_type"`
	KubeConfigDir string `yaml:"kubeconfig_dir"`
}

type MqttClientConfig struct {
	Server    []string `yaml:"server"`
	Clientid  string   `yaml:"clientid"`
	FileStore string   `yaml:"filestore"`
	UserName  string   `yaml:"username"`
	Password  string   `yaml:"password"`
}

func (c *MqttClientConfig) GenerateClientId(prefix string) {
	randString := util.RandStringBytes(5)
	c.Clientid = prefix + "-" + randString
}

type EsClientConfig struct {
	Server   []string `yaml:"server"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
}

// Config Singleton
var ConfigSingleton Config
