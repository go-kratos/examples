module kratos-casbin

go 1.16

require (
	github.com/go-kratos/kratos/contrib/config/consul/v2 v2.0.0-20220301141459-ed6ab7caf9ca
	github.com/go-kratos/kratos/contrib/config/etcd/v2 v2.0.0-20220301141459-ed6ab7caf9ca
	github.com/go-kratos/kratos/contrib/config/nacos/v2 v2.0.0-20220301141459-ed6ab7caf9ca
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20211119091424-ef3322ec0764
	github.com/go-kratos/kratos/v2 v2.2.0
	github.com/google/wire v0.5.0
	github.com/hashicorp/consul/api v1.11.0
	github.com/nacos-group/nacos-sdk-go v1.1.0
	go.etcd.io/etcd/client/v3 v3.5.2
	go.opentelemetry.io/otel v1.3.0
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0
	go.opentelemetry.io/otel/sdk v1.3.0
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/armon/go-metrics v0.3.10 // indirect
	github.com/casbin/casbin/v2 v2.42.0
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-kratos/swagger-api v1.0.1
	github.com/golang-jwt/jwt/v4 v4.2.0
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/gorilla/handlers v1.5.1
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.0.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/serf v0.9.6 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mitchellh/mapstructure v1.4.2 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/tx7do/kratos-casbin v0.0.0-20220317124747-964d93859d7d
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	google.golang.org/genproto v0.0.0-20220310185008-1973136f34c6
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/ini.v1 v1.51.0 // indirect
)
