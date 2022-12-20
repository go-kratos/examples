module kratos-realtimemap

go 1.18

require (
	github.com/go-kratos/kratos/contrib/config/consul/v2 v2.0.0-20221220065744-a017ab09576f
	github.com/go-kratos/kratos/contrib/config/etcd/v2 v2.0.0-20221220065744-a017ab09576f
	github.com/go-kratos/kratos/contrib/config/nacos/v2 v2.0.0-20221220065744-a017ab09576f
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20221220065744-a017ab09576f
	github.com/go-kratos/kratos/v2 v2.5.3
	github.com/go-kratos/swagger-api v1.0.1
	github.com/google/subcommands v1.0.1
	github.com/google/wire v0.5.0
	github.com/gorilla/handlers v1.5.1
	github.com/hashicorp/consul/api v1.18.0
	github.com/kellydunn/golang-geo v0.7.0
	github.com/nacos-group/nacos-sdk-go v1.1.3
	github.com/tx7do/kratos-transport v1.0.5
	github.com/tx7do/kratos-transport/transport/kafka v0.0.0-20221220110912-9db1b79d385c
	github.com/tx7do/kratos-transport/transport/mqtt v0.0.0-20221220110310-6e536d718b0a
	github.com/tx7do/kratos-transport/transport/websocket v0.0.0-20221220110912-9db1b79d385c
	go.etcd.io/etcd/client/v3 v3.5.6
	go.opentelemetry.io/otel v1.11.2
	go.opentelemetry.io/otel/exporters/jaeger v1.11.2
	go.opentelemetry.io/otel/sdk v1.11.2
	golang.org/x/tools v0.1.12
	google.golang.org/genproto v0.0.0-20221207170731-23e4bf6bdc37
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/aliyun/alibaba-cloud-sdk-go v1.62.91 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/eclipse/paho.mqtt.golang v1.4.2 // indirect
	github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-kratos/grpc-gateway/v2 v2.5.1-0.20210811062259-c92d36e434b1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.4.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.13 // indirect
	github.com/kylelemons/go-gypsy v1.0.0 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/opentracing/opentracing-go v1.2.1-0.20220228012449-10b1cf09e00b // indirect
	github.com/openzipkin/zipkin-go v0.4.1 // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rakyll/statik v0.1.7 // indirect
	github.com/segmentio/kafka-go v0.4.38 // indirect
	github.com/tx7do/kratos-transport/broker/kafka v0.0.0-20221220110912-9db1b79d385c // indirect
	github.com/tx7do/kratos-transport/broker/mqtt v0.0.0-20221220110310-6e536d718b0a // indirect
	github.com/xdg/scram v1.0.5 // indirect
	github.com/xdg/stringprep v1.0.3 // indirect
	github.com/ziutek/mymysql v1.5.4 // indirect
	go.etcd.io/etcd/api/v3 v3.5.6 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.6 // indirect
	go.opentelemetry.io/otel/exporters/zipkin v1.11.2 // indirect
	go.opentelemetry.io/otel/trace v1.11.2 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/crypto v0.4.0 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
