# kratos-realtimemap

## 实时公共交通地图

源自于 [Proto.Actor](https://proto.actor/) 的一个实时地图的实例 [realtimemap-go](https://github.com/asynkron/realtimemap-go.git)

它有一个在线的演示可看: <https://realtimemap.skyrise.cloud/>

它之前与网页端的实时数据传输用的是微软的SignalR,我这里则直接使用了Websocket.

后端使用MQTT收取了一个开放的试验数据,它背后是推的一个公共交通的开放标准 [HFP API](https://digitransit.fi/en/developers/apis/4-realtime-api/vehicle-positions/) ,数据的覆盖面挺全的,该有的差不多都有了.另外谷歌也有一个基于Protobuf的开放API [GTFS](https://developers.google.com/transit/gtfs-realtime/reference).

后端收取到了mqtt数据之后,正常来说是要推到kafka里面去的,然后其他的微服务按需消费:该入库的入库,该实时推给客户端的推给客户端,该缓存的缓存.但是,我偷了个懒,一切从简了.总之,入库没有入,推kafka也没推,缓存数据也只是缓存在服务的内存里面.没有缓存到Redis去.

另外还有一个功能,我没有去具体实现,正常来说,设备会很多,在地图上面不可能全部的展示出来,这时候,前端地图就需要选择性的接收交通工具的实时数据,不然得爆.它是怎么做到这个事情的呢,当用户在前端切换地图的视窗的时候(放大,缩小,平移),会把它的视窗的大小和坐标推给后端,后端再根据这个视窗去裁剪,把不在视窗范围的交通工具给剪除,只留下视窗之内的交通工具.

前端调用数据,使用了Restfull和Websocket结合的方式,实时的遥测坐标数据使用Websocket来推,而相对比较静态的数据,比如,地理围栏,车辆属性信息,则使用了Restfull拉取.当然了,当车辆没有实时的遥测数据的时候,其缓存的历史数据,也是可以使用Restfull来拉取的,相对来说,这也是比较静态的数据.

- **注意**  
**今天测试发现,mqtt接收数据接收一段时间就自动断掉了,我还以为是我这边出问题了,后来做了一些测试才发现,是对方限制了使用,限制是根据ClientID进行的验证.**

## 涵盖的技术点

- 使用Kratos开发微服务
- 使用Kratos的BFF与网页端通讯
- Kratos与MQTT的融合使用
- Kratos与Websocket的融合使用

## 技术栈

- [Kratos](https://go-kratos.dev/)
- [Consul](https://www.consul.io/)
- [Jaeger](https://www.jaegertracing.io/)
- [MQTT](https://mqtt.org/)
- [Websocket](https://entgo.io/)
- [VUE](https://vuejs.org/)

## Docker部署开发服务器

### Consul

```shell
docker pull bitnami/consul:latest

docker run -itd \
    --name consul-server-standalone \
    -p 8300:8300 \
    -p 8500:8500 \
    -p 8600:8600/udp \
    -e CONSUL_BIND_INTERFACE='eth0' \
    -e CONSUL_AGENT_MODE=server \
    -e CONSUL_ENABLE_UI=true \
    -e CONSUL_BOOTSTRAP_EXPECT=1 \
    -e CONSUL_CLIENT_LAN_ADDRESS=0.0.0.0 \
    bitnami/consul:latest
```

### Jaeger

```shell
docker pull jaegertracing/all-in-one:latest

docker run -d \
    --name jaeger \
    -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
    -p 5775:5775/udp \
    -p 6831:6831/udp \
    -p 6832:6832/udp \
    -p 5778:5778 \
    -p 16686:16686 \
    -p 14268:14268 \
    -p 14250:14250 \
    -p 9411:9411 \
    jaegertracing/all-in-one:latest
```

## 测试

Swagger-UI的访问地址: <http://localhost:8800/q/swagger-ui>  
前台的访问地址: <http://localhost:8080/>

## 参考资料

- [GTFS Realtime Reference](https://developers.google.com/transit/gtfs-realtime/reference)
- [High-frequency positioning](https://digitransit.fi/en/developers/apis/4-realtime-api/vehicle-positions/)
