module kratos-casbin

go 1.16

require (
	github.com/go-kratos/kratos/contrib/config/consul/v2 v2.0.0-20220301141459-ed6ab7caf9ca
	github.com/go-kratos/kratos/contrib/config/etcd/v2 v2.0.0-20220301141459-ed6ab7caf9ca
	github.com/go-kratos/kratos/contrib/config/nacos/v2 v2.0.0-20220301141459-ed6ab7caf9ca
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20211119091424-ef3322ec0764
	github.com/go-kratos/kratos/v2 v2.7.2
	github.com/google/wire v0.5.0
	github.com/hashicorp/consul/api v1.11.0
	github.com/nacos-group/nacos-sdk-go v1.1.0
	go.etcd.io/etcd/client/v3 v3.5.2
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0
	go.opentelemetry.io/otel/sdk v1.16.0
	google.golang.org/grpc v1.58.2
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/armon/go-metrics v0.3.10 // indirect
	github.com/casbin/casbin/v2 v2.84.0
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-kratos/swagger-api v1.0.1
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/gorilla/handlers v1.5.1
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.0.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/serf v0.9.6 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mitchellh/mapstructure v1.4.2 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/tx7do/kratos-casbin v0.0.0-20240311042652-fbaa11f94322
	google.golang.org/genproto/googleapis/api v0.0.0-20230803162519-f966b187b2e5
	gopkg.in/ini.v1 v1.51.0 // indirect
)
