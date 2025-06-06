package bootstrap

import (
	"github.com/stretchr/testify/assert"
	"kratos-gorm-example/gen/api/go/common/conf"
	"testing"
)

func TestNewConsulRegistry(t *testing.T) {
	var cfg conf.Registry
	cfg.Consul.Scheme = "http"
	cfg.Consul.Address = "127.0.0.1:8500"
	cfg.Consul.HealthCheck = false

	reg := NewConsulRegistry(&cfg)
	assert.Nil(t, reg)
}

func TestNewEtcdRegistry(t *testing.T) {
	var cfg conf.Registry
	cfg.Etcd.Endpoints = []string{"127.0.0.1:2379"}

	reg := NewEtcdRegistry(&cfg)
	assert.Nil(t, reg)
}

func TestNewNacosRegistry(t *testing.T) {
	var cfg conf.Registry
	cfg.Nacos.Address = "127.0.0.1"
	cfg.Nacos.Port = 8848

	reg := NewNacosRegistry(&cfg)
	assert.Nil(t, reg)
}

func TestNewZooKeeperRegistry(t *testing.T) {
	var cfg conf.Registry
	cfg.Zookeeper.Endpoints = []string{"127.0.0.1:2181"}

	reg := NewZooKeeperRegistry(&cfg)
	assert.Nil(t, reg)
}

func TestNewKubernetesRegistry(t *testing.T) {
	var cfg conf.Registry
	reg := NewKubernetesRegistry(&cfg)
	assert.Nil(t, reg)
}

func TestNewEurekaRegistry(t *testing.T) {
	var cfg conf.Registry
	cfg.Eureka.Endpoints = []string{"https://127.0.0.1:18761"}

	reg := NewEurekaRegistry(&cfg)
	assert.Nil(t, reg)
}

func TestNewPolarisRegistry(t *testing.T) {
	//var cfg conf.Registry
	//cfg.Polaris.Address = "127.0.0.1"
	//cfg.Polaris.Port = 8091
	//cfg.Polaris.InstanceCount = 5
	//cfg.Polaris.Namespace = "default"
	//cfg.Polaris.Service = "DiscoverEchoServer"
	//cfg.Polaris.Token = ""
	//
	//reg := NewPolarisRegistry(&cfg)
	//assert.Nil(t, reg)
}

func TestNewServicecombRegistry(t *testing.T) {
	var cfg conf.Registry
	cfg.Servicecomb.Endpoints = []string{"127.0.0.1:30100"}

	reg := NewServicecombRegistry(&cfg)
	assert.Nil(t, reg)
}
