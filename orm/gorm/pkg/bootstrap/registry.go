package bootstrap

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"path/filepath"

	// etcd
	etcdKratos "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	etcdClient "go.etcd.io/etcd/client/v3"

	// consul
	consulKratos "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	consulClient "github.com/hashicorp/consul/api"

	// eureka
	eurekaKratos "github.com/go-kratos/kratos/contrib/registry/eureka/v2"

	// nacos
	nacosKratos "github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	nacosClients "github.com/nacos-group/nacos-sdk-go/clients"
	nacosConstant "github.com/nacos-group/nacos-sdk-go/common/constant"
	nacosVo "github.com/nacos-group/nacos-sdk-go/vo"

	// zookeeper
	zookeeperKratos "github.com/go-kratos/kratos/contrib/registry/zookeeper/v2"
	"github.com/go-zookeeper/zk"

	// kubernetes
	k8sRegistry "github.com/go-kratos/kratos/contrib/registry/kubernetes/v2"
	k8s "k8s.io/client-go/kubernetes"
	k8sRest "k8s.io/client-go/rest"
	k8sTools "k8s.io/client-go/tools/clientcmd"
	k8sUtil "k8s.io/client-go/util/homedir"

	// servicecomb
	servicecombClient "github.com/go-chassis/sc-client"
	servicecombKratos "github.com/go-kratos/kratos/contrib/registry/servicecomb/v2"

	"kratos-gorm-example/gen/api/go/common/conf"
)

type RegistryType string

const (
	RegistryTypeConsul      RegistryType = "consul"
	RegistryTypeEtcd        RegistryType = "etcd"
	RegistryTypeZooKeeper   RegistryType = "zookeeper"
	RegistryTypeNacos       RegistryType = "nacos"
	RegistryTypeKubernetes  RegistryType = "kubernetes"
	RegistryTypeEureka      RegistryType = "eureka"
	RegistryTypePolaris     RegistryType = "polaris"
	RegistryTypeServicecomb RegistryType = "servicecomb"
)

// NewRegistry 创建一个注册客户端
func NewRegistry(cfg *conf.Registry) registry.Registrar {
	if cfg == nil {
		return nil
	}

	switch RegistryType(cfg.Type) {
	case RegistryTypeConsul:
		return NewConsulRegistry(cfg)
	case RegistryTypeEtcd:
		return NewEtcdRegistry(cfg)
	case RegistryTypeZooKeeper:
		return NewZooKeeperRegistry(cfg)
	case RegistryTypeNacos:
		return NewNacosRegistry(cfg)
	case RegistryTypeKubernetes:
		return NewKubernetesRegistry(cfg)
	case RegistryTypeEureka:
		return NewEurekaRegistry(cfg)
	case RegistryTypePolaris:
		return nil
	case RegistryTypeServicecomb:
		return NewServicecombRegistry(cfg)
	}

	return nil
}

// NewDiscovery 创建一个发现客户端
func NewDiscovery(cfg *conf.Registry) registry.Discovery {
	if cfg == nil {
		return nil
	}

	switch RegistryType(cfg.Type) {
	case RegistryTypeConsul:
		return NewConsulRegistry(cfg)
	case RegistryTypeEtcd:
		return NewEtcdRegistry(cfg)
	case RegistryTypeZooKeeper:
		return NewZooKeeperRegistry(cfg)
	case RegistryTypeNacos:
		return NewNacosRegistry(cfg)
	case RegistryTypeKubernetes:
		return NewKubernetesRegistry(cfg)
	case RegistryTypeEureka:
		return NewEurekaRegistry(cfg)
	case RegistryTypePolaris:
		return nil
	case RegistryTypeServicecomb:
		return NewServicecombRegistry(cfg)
	}

	return nil
}

// NewConsulRegistry 创建一个注册发现客户端 - Consul
func NewConsulRegistry(c *conf.Registry) *consulKratos.Registry {
	cfg := consulClient.DefaultConfig()
	cfg.Address = c.Consul.Address
	cfg.Scheme = c.Consul.Scheme

	var cli *consulClient.Client
	var err error
	if cli, err = consulClient.NewClient(cfg); err != nil {
		log.Fatal(err)
	}

	reg := consulKratos.New(cli, consulKratos.WithHealthCheck(c.Consul.HealthCheck))

	return reg
}

// NewEtcdRegistry 创建一个注册发现客户端 - Etcd
func NewEtcdRegistry(c *conf.Registry) *etcdKratos.Registry {
	cfg := etcdClient.Config{
		Endpoints: c.Etcd.Endpoints,
	}

	var err error
	var cli *etcdClient.Client
	if cli, err = etcdClient.New(cfg); err != nil {
		log.Fatal(err)
	}

	reg := etcdKratos.New(cli)

	return reg
}

// NewZooKeeperRegistry 创建一个注册发现客户端 - ZooKeeper
func NewZooKeeperRegistry(c *conf.Registry) *zookeeperKratos.Registry {
	conn, _, err := zk.Connect(c.Zookeeper.Endpoints, c.Zookeeper.Timeout.AsDuration())
	if err != nil {
		log.Fatal(err)
	}

	reg := zookeeperKratos.New(conn)
	if err != nil {
		log.Fatal(err)
	}

	return reg
}

// NewNacosRegistry 创建一个注册发现客户端 - Nacos
func NewNacosRegistry(c *conf.Registry) *nacosKratos.Registry {
	srvConf := []nacosConstant.ServerConfig{
		*nacosConstant.NewServerConfig(c.Nacos.Address, c.Nacos.Port),
	}

	cliConf := nacosConstant.ClientConfig{
		NamespaceId:          c.Nacos.NamespaceId,
		TimeoutMs:            uint64(c.Nacos.Timeout.AsDuration().Milliseconds()), // http请求超时时间，单位毫秒
		BeatInterval:         c.Nacos.BeatInterval.AsDuration().Milliseconds(),    // 心跳间隔时间，单位毫秒
		UpdateThreadNum:      int(c.Nacos.UpdateThreadNum),                        // 更新服务的线程数
		LogLevel:             c.Nacos.LogLevel,
		CacheDir:             c.Nacos.CacheDir,             // 缓存目录
		LogDir:               c.Nacos.LogDir,               // 日志目录
		NotLoadCacheAtStart:  c.Nacos.NotLoadCacheAtStart,  // 在启动时不读取本地缓存数据，true--不读取，false--读取
		UpdateCacheWhenEmpty: c.Nacos.UpdateCacheWhenEmpty, // 当服务列表为空时是否更新本地缓存，true--更新,false--不更新
	}

	cli, err := nacosClients.NewNamingClient(
		nacosVo.NacosClientParam{
			ClientConfig:  &cliConf,
			ServerConfigs: srvConf,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	reg := nacosKratos.New(cli)

	return reg
}

// NewKubernetesRegistry 创建一个注册发现客户端 - Kubernetes
func NewKubernetesRegistry(_ *conf.Registry) *k8sRegistry.Registry {
	restConfig, err := k8sRest.InClusterConfig()
	if err != nil {
		home := k8sUtil.HomeDir()
		kubeConfig := filepath.Join(home, ".kube", "config")
		restConfig, err = k8sTools.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			log.Fatal(err)
			return nil
		}
	}

	clientSet, err := k8s.NewForConfig(restConfig)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	reg := k8sRegistry.NewRegistry(clientSet)

	return reg
}

// NewEurekaRegistry 创建一个注册发现客户端 - Eureka
func NewEurekaRegistry(c *conf.Registry) *eurekaKratos.Registry {
	var opts []eurekaKratos.Option
	opts = append(opts, eurekaKratos.WithHeartbeat(c.Eureka.HeartbeatInterval.AsDuration()))
	opts = append(opts, eurekaKratos.WithRefresh(c.Eureka.RefreshInterval.AsDuration()))
	opts = append(opts, eurekaKratos.WithEurekaPath(c.Eureka.Path))

	var err error
	var reg *eurekaKratos.Registry
	if reg, err = eurekaKratos.New(c.Eureka.Endpoints, opts...); err != nil {
		log.Fatal(err)
	}

	return reg
}

// NewPolarisRegistry 创建一个注册发现客户端 - Polaris
//func NewPolarisRegistry(c *conf.Registry) *polarisKratos.Registry {
//	var err error
//
//	var consumer polarisApi.ConsumerAPI
//	if consumer, err = polarisApi.NewConsumerAPI(); err != nil {
//		log.Fatalf("fail to create consumerAPI, err is %v", err)
//	}
//
//	var provider polarisApi.ProviderAPI
//	provider = polarisApi.NewProviderAPIByContext(consumer.SDKContext())
//
//	log.Infof("start to register instances, count %d", c.Polaris.InstanceCount)
//
//	var resp *polarisModel.InstanceRegisterResponse
//	for i := 0; i < (int)(c.Polaris.InstanceCount); i++ {
//		registerRequest := &polarisApi.InstanceRegisterRequest{}
//		registerRequest.Service = c.Polaris.Service
//		registerRequest.Namespace = c.Polaris.Namespace
//		registerRequest.Host = c.Polaris.Address
//		registerRequest.Port = (int)(c.Polaris.Port) + i
//		registerRequest.ServiceToken = c.Polaris.Token
//		registerRequest.SetHealthy(true)
//		if resp, err = provider.RegisterInstance(registerRequest); err != nil {
//			log.Fatalf("fail to register instance %d, err is %v", i, err)
//		} else {
//			log.Infof("register instance %d response: instanceId %s", i, resp.InstanceID)
//		}
//	}
//
//	reg := polarisKratos.NewRegistry(provider, consumer)
//
//	return reg
//}

// NewServicecombRegistry 创建一个注册发现客户端 - Servicecomb
func NewServicecombRegistry(c *conf.Registry) *servicecombKratos.Registry {
	cfg := servicecombClient.Options{
		Endpoints: c.Servicecomb.Endpoints,
	}

	var cli *servicecombClient.Client
	var err error
	if cli, err = servicecombClient.NewClient(cfg); err != nil {
		log.Fatal(err)
	}

	reg := servicecombKratos.NewRegistry(cli)

	return reg
}
