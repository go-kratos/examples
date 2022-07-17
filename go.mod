module github.com/go-kratos/examples

go 1.16

require (
	entgo.io/ent v0.9.0
	github.com/BurntSushi/toml v0.3.1
	github.com/envoyproxy/protoc-gen-validate v0.6.2
	github.com/gin-gonic/gin v1.7.7
	github.com/go-kratos/gin v0.1.0
	github.com/go-kratos/kratos/contrib/config/apollo/v2 v2.0.0-20220309025117-4387085047b9
	github.com/go-kratos/kratos/contrib/metrics/prometheus/v2 v2.0.0-20220309025117-4387085047b9
	github.com/go-kratos/kratos/contrib/opensergo/v2 v2.0.0-20220517145828-c8c870b77f70
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20220309025117-4387085047b9
	github.com/go-kratos/kratos/contrib/registry/discovery/v2 v2.0.0-20220309025117-4387085047b9
	github.com/go-kratos/kratos/contrib/registry/etcd/v2 v2.0.0-20220309025117-4387085047b9
	github.com/go-kratos/kratos/contrib/registry/eureka/v2 v2.0.0-20220330020930-99a0646acb98
	github.com/go-kratos/kratos/contrib/registry/kubernetes/v2 v2.0.0-20220714125901-1ab3d8f02840
	github.com/go-kratos/kratos/contrib/registry/nacos/v2 v2.0.0-20220309025117-4387085047b9
	github.com/go-kratos/kratos/contrib/registry/zookeeper/v2 v2.0.0-20220309025117-4387085047b9
	github.com/go-kratos/kratos/v2 v2.4.0
	github.com/go-kratos/swagger-api v1.0.0
	github.com/go-redis/redis/extra/redisotel v0.3.0
	github.com/go-redis/redis/v8 v8.11.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang-jwt/jwt/v4 v4.4.1
	github.com/google/wire v0.5.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	github.com/hashicorp/consul/api v1.12.0
	github.com/labstack/echo/v4 v4.6.3
	github.com/nacos-group/nacos-sdk-go v1.0.9
	github.com/nicksnyder/go-i18n/v2 v2.1.2
	github.com/prometheus/client_golang v1.12.0
	github.com/segmentio/kafka-go v0.4.27
	github.com/sirupsen/logrus v1.8.1
	github.com/soheilhy/cmux v0.1.4
	go.etcd.io/etcd/client/v3 v3.5.1
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/exporters/jaeger v1.3.0
	go.opentelemetry.io/otel/sdk v1.7.0
	go.opentelemetry.io/otel/trace v1.7.0
	go.uber.org/zap v1.19.0
	golang.org/x/text v0.3.7
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd
	google.golang.org/grpc v1.46.2
	google.golang.org/grpc/examples v0.0.0-20220105183818-2fb1ac854b20 // indirect
	google.golang.org/protobuf v1.28.0
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.2
	k8s.io/client-go v0.24.3
)
