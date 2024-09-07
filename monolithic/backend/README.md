# GO后端

## API文档

### Swagger UI

- [Admin Swagger UI](http://localhost:7788/docs/)

### openapi.yaml

- [Admin openapi.yaml](http://localhost:7788/docs/openapi.yaml)

## Buf.Build使用

使用[buf.build](https://buf.build/)进行Protobuf API的工程化构建。

相关命令行工具和插件的具体安装方法请参见：[Kratos微服务框架API工程化指南](https://juejin.cn/post/7191095845096259641)

在`backend`根目录下执行命令：

### 更新buf.lock

```bash
buf mod update
```

### 生成GO代码

```bash
buf generate
```

### 生成OpenAPI v3文档

```bash
buf generate --path api/admin/service/v1 --template api/admin/service/v1/buf.openapi.gen.yaml
```

## Make构建

请在`app/{服务名}/service`下执行：

### 初始化开发环境

```bash
make init
```

### 生成API的go代码

```bash
make api
```

### 生成API的OpenAPI v3 文档

```bash
make openapi
```

### 生成ent代码

```bash
make ent
```

### 生成wire代码

```bash
make wire
```

### 构建程序

```bash
make build
```

### 调试运行

```bash
make run
```

### 构建Docker镜像

```bash
make docker
```
