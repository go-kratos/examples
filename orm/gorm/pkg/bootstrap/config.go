package bootstrap

import (
	"os"
	"path/filepath"
	"strings"

	"google.golang.org/grpc"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"

	// file
	fileKratos "github.com/go-kratos/kratos/v2/config/file"

	// etcd
	etcdKratos "github.com/go-kratos/kratos/contrib/config/etcd/v2"
	etcdClient "go.etcd.io/etcd/client/v3"

	// consul
	consulKratos "github.com/go-kratos/kratos/contrib/config/consul/v2"
	consulApi "github.com/hashicorp/consul/api"

	// nacos
	nacosKratos "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	nacosClients "github.com/nacos-group/nacos-sdk-go/clients"
	nacosConstant "github.com/nacos-group/nacos-sdk-go/common/constant"
	nacosVo "github.com/nacos-group/nacos-sdk-go/vo"

	// apollo
	apolloKratos "github.com/go-kratos/kratos/contrib/config/apollo/v2"

	// kubernetes
	k8sKratos "github.com/go-kratos/kratos/contrib/config/kubernetes/v2"
	k8sUtil "k8s.io/client-go/util/homedir"

	"kratos-gorm-example/gen/api/go/common/conf"
)

const remoteConfigSourceConfigFile = "remote.yaml"

// NewConfigProvider 创建一个配置
func NewConfigProvider(configPath string) config.Config {
	err, rc := LoadRemoteConfigSourceConfigs(configPath)
	if err != nil {
		log.Error("LoadRemoteConfigSourceConfigs: ", err.Error())
	}
	if rc != nil {
		return config.New(
			config.WithSource(
				NewFileConfigSource(configPath),
				NewRemoteConfigSource(rc),
			),
		)
	} else {
		return config.New(
			config.WithSource(
				NewFileConfigSource(configPath),
			),
		)
	}
}

// LoadBootstrapConfig 加载程序引导配置
func LoadBootstrapConfig(configPath string) *conf.Bootstrap {
	cfg := NewConfigProvider(configPath)
	if err := cfg.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := cfg.Scan(&bc); err != nil {
		panic(err)
	}

	if bc.Server == nil {
		bc.Server = &conf.Server{}
		_ = cfg.Scan(&bc.Server)
	}

	if bc.Client == nil {
		bc.Client = &conf.Client{}
		_ = cfg.Scan(&bc.Client)
	}

	if bc.Data == nil {
		bc.Data = &conf.Data{}
		_ = cfg.Scan(&bc.Data)
	}

	if bc.Trace == nil {
		bc.Trace = &conf.Tracer{}
		_ = cfg.Scan(&bc.Trace)
	}

	if bc.Logger == nil {
		bc.Logger = &conf.Logger{}
		_ = cfg.Scan(&bc.Logger)
	}

	if bc.Registry == nil {
		bc.Registry = &conf.Registry{}
		_ = cfg.Scan(&bc.Registry)
	}

	if bc.Oss == nil {
		bc.Oss = &conf.OSS{}
		_ = cfg.Scan(&bc.Oss)
	}

	return &bc
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// LoadRemoteConfigSourceConfigs 加载远程配置源的本地配置
func LoadRemoteConfigSourceConfigs(configPath string) (error, *conf.RemoteConfig) {
	configPath = configPath + "/" + remoteConfigSourceConfigFile
	if !pathExists(configPath) {
		return nil, nil
	}

	cfg := config.New(
		config.WithSource(
			NewFileConfigSource(configPath),
		),
	)
	defer func(cfg config.Config) {
		err := cfg.Close()
		if err != nil {
			panic(err)
		}
	}(cfg)

	var err error

	if err = cfg.Load(); err != nil {
		return err, nil
	}

	var rc conf.Bootstrap
	if err = cfg.Scan(&rc); err != nil {
		return err, nil
	}

	return nil, rc.Config
}

type ConfigType string

const (
	ConfigTypeLocalFile  ConfigType = "file"
	ConfigTypeNacos      ConfigType = "nacos"
	ConfigTypeConsul     ConfigType = "consul"
	ConfigTypeEtcd       ConfigType = "etcd"
	ConfigTypeApollo     ConfigType = "apollo"
	ConfigTypeKubernetes ConfigType = "kubernetes"
	ConfigTypePolaris    ConfigType = "polaris"
)

// NewRemoteConfigSource 创建一个远程配置源
func NewRemoteConfigSource(c *conf.RemoteConfig) config.Source {
	switch ConfigType(c.Type) {
	default:
		fallthrough
	case ConfigTypeLocalFile:
		return nil
	case ConfigTypeNacos:
		return NewNacosConfigSource(c)
	case ConfigTypeConsul:
		return NewConsulConfigSource(c)
	case ConfigTypeEtcd:
		return NewEtcdConfigSource(c)
	case ConfigTypeApollo:
		return NewApolloConfigSource(c)
	case ConfigTypeKubernetes:
		return NewKubernetesConfigSource(c)
	case ConfigTypePolaris:
		return NewPolarisConfigSource(c)
	}
}

// getConfigKey 获取合法的配置名
func getConfigKey(configKey string, useBackslash bool) string {
	if useBackslash {
		return strings.Replace(configKey, `.`, `/`, -1)
	} else {
		return configKey
	}
}

// NewFileConfigSource 创建一个本地文件配置源
func NewFileConfigSource(filePath string) config.Source {
	return fileKratos.NewSource(filePath)
}

// NewNacosConfigSource 创建一个远程配置源 - Nacos
func NewNacosConfigSource(c *conf.RemoteConfig) config.Source {
	srvConf := []nacosConstant.ServerConfig{
		*nacosConstant.NewServerConfig(c.Nacos.Address, c.Nacos.Port),
	}

	cliConf := nacosConstant.ClientConfig{
		TimeoutMs:            10 * 1000, // http请求超时时间，单位毫秒
		BeatInterval:         5 * 1000,  // 心跳间隔时间，单位毫秒
		UpdateThreadNum:      20,        // 更新服务的线程数
		LogLevel:             "debug",
		CacheDir:             "../../configs/cache", // 缓存目录
		LogDir:               "../../configs/log",   // 日志目录
		NotLoadCacheAtStart:  true,                  // 在启动时不读取本地缓存数据，true--不读取，false--读取
		UpdateCacheWhenEmpty: true,                  // 当服务列表为空时是否更新本地缓存，true--更新,false--不更新
	}

	nacosClient, err := nacosClients.NewConfigClient(
		nacosVo.NacosClientParam{
			ClientConfig:  &cliConf,
			ServerConfigs: srvConf,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return nacosKratos.NewConfigSource(nacosClient,
		nacosKratos.WithGroup(getConfigKey(c.Nacos.Key, false)),
		nacosKratos.WithDataID("bootstrap.yaml"),
	)
}

// NewEtcdConfigSource 创建一个远程配置源 - Etcd
func NewEtcdConfigSource(c *conf.RemoteConfig) config.Source {
	cfg := etcdClient.Config{
		Endpoints:   c.Etcd.Endpoints,
		DialTimeout: c.Etcd.Timeout.AsDuration(),
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	}

	cli, err := etcdClient.New(cfg)
	if err != nil {
		panic(err)
	}

	source, err := etcdKratos.New(cli, etcdKratos.WithPath(getConfigKey(c.Etcd.Key, true)))
	if err != nil {
		log.Fatal(err)
	}

	return source
}

// NewConsulConfigSource 创建一个远程配置源 - Consul
func NewConsulConfigSource(c *conf.RemoteConfig) config.Source {
	cfg := consulApi.DefaultConfig()
	cfg.Address = c.Consul.Address
	cfg.Scheme = c.Consul.Scheme

	cli, err := consulApi.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	source, err := consulKratos.New(cli,
		consulKratos.WithPath(getConfigKey(c.Consul.Key, true)),
	)
	if err != nil {
		log.Fatal(err)
	}

	return source
}

// NewApolloConfigSource 创建一个远程配置源 - Apollo
func NewApolloConfigSource(c *conf.RemoteConfig) config.Source {
	source := apolloKratos.NewSource(
		apolloKratos.WithAppID(c.Apollo.AppId),
		apolloKratos.WithCluster(c.Apollo.Cluster),
		apolloKratos.WithEndpoint(c.Apollo.Endpoint),
		apolloKratos.WithNamespace(c.Apollo.Namespace),
		apolloKratos.WithSecret(c.Apollo.Secret),
		apolloKratos.WithEnableBackup(),
	)
	return source
}

// NewKubernetesConfigSource 创建一个远程配置源 - Kubernetes
func NewKubernetesConfigSource(c *conf.RemoteConfig) config.Source {
	source := k8sKratos.NewSource(
		k8sKratos.Namespace(c.Kubernetes.Namespace),
		k8sKratos.LabelSelector(""),
		k8sKratos.KubeConfig(filepath.Join(k8sUtil.HomeDir(), ".kube", "config")),
	)
	return source
}

// NewPolarisConfigSource 创建一个远程配置源 - Polaris
func NewPolarisConfigSource(_ *conf.RemoteConfig) config.Source {
	//configApi, err := polarisApi.NewConfigAPI()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//var opts []polarisKratos.Option
	//opts = append(opts, polarisKratos.WithNamespace("default"))
	//opts = append(opts, polarisKratos.WithFileGroup("default"))
	//opts = append(opts, polarisKratos.WithFileName("default.yaml"))
	//
	//source, err := polarisKratos.New(configApi, opts...)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//return source
	return nil
}
