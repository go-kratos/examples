# kratos-casbin

本项目主要示例了两个中间件的使用方法：
- 使用JWT来实现用户登陆认证。
- 如何使用Casbin来实现RBAC权限管理。

## [AuthN & AuthZ](http://technoponder.com/authentication-authorization/)

在此之前，我们需要搞清楚两个概念：  

- AuthN – Authentication is establishing the your identity.
- AuthZ – Authorization is establishing your privilege

用中文简言之：

- AuthN 系统主要用于认证 （Authentication），决定谁访问了系统。
- AuthZ 系统主要用于授权 （Authorization），决定访问者具有什么样的权限。

在这个示例中：

- AuthN 系统是用户登录系统，并且获取用户的身份信息，由JWT来实现。
- AuthZ 系统是根据用户的身份信息，决定用户具有什么权限，由Casbin来实现。

在这个示例中，我们将使用三个角色：

- admin：管理员角色，具有所有权限，显示所有标签页。
- moderator：普通用户角色，具有部分权限，无权限访问User标签页，无法看到Admin标签页。
- user：普通用户角色，具有一些权限，无权限访问User标签页，无法看到Admin和Moderator标签页。

我简化了这个示例，角色只用于限定用户显示的标签页而已，实际上我演示上只是用了用户名去限定用户的权限，而非角色。

## [什么是Casbin？](https://casbin.org/docs/zh-CN/overview)

Casbin是一个强大的、高效的开源访问控制框架，其权限管理机制支持多种访问控制模型。目前这个框架的生态已经发展的越来越好了。提供了各种语言的类库，自定义的权限模型语言，以及模型编辑器。

### Casbin 可以：

1. 支持自定义请求的格式，默认的请求格式为`{subject, object, action}`。
2. 具有访问控制模型model和策略policy两个核心概念。
3. 支持RBAC中的多层角色继承，不止主体可以有角色，资源也可以具有角色。
4. 支持内置的超级用户 例如：`root` 或 `administrator`。超级用户可以执行任何操作而无需显式的权限声明。
5. 支持多种内置的操作符，如 `keyMatch`，方便对路径式的资源进行管理，如 `/foo/bar` 可以映射到 `/foo*`

### Casbin 不能：

1. 身份认证 authentication（即验证用户的用户名和密码），Casbin 只负责访问控制。应该有其他专门的组件负责身份认证，然后由 Casbin 进行访问控制，二者是相互配合的关系。
2. 管理用户列表或角色列表。 Casbin 认为由项目自身来管理用户、角色列表更为合适， 用户通常有他们的密码，但是 Casbin 的设计思想并不是把它作为一个存储密码的容器。 而是存储RBAC方案中用户和角色之间的映射关系。

## [什么是JWT？](https://www.jianshu.com/p/576dbf44b2ae)

Json web token (JWT), 是为了在网络应用环境间传递声明而执行的一种基于JSON的开放标准（ (RFC 7519).该token被设计为紧凑且安全的，特别适用于分布式站点的单点登录（SSO）场景。JWT的声明一般被用来在身份提供者和服务提供者间传递被认证的用户身份信息，以便于从资源服务器获取资源，也可以增加一些额外的其它业务逻辑所必须的声明信息，该token也可直接被用于认证，也可被加密。

### [JWT需要注意的点](https://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html)

1. JWT 默认是不加密，但也是可以加密的。生成原始 Token 以后，可以用密钥再加密一次。
2. JWT 不加密的情况下，不能将秘密数据写入 JWT。
3. JWT 不仅可以用于认证，也可以用于交换信息。有效使用 JWT，可以降低服务器查询数据库的次数。
4. JWT 的最大缺点是，由于服务器不保存 session 状态，因此无法在使用过程中废止某个 token，或者更改 token 的权限。也就是说，一旦 JWT 签发了，在到期之前就会始终有效，除非服务器部署额外的逻辑。
5. JWT 本身包含了认证信息，一旦泄露，任何人都可以获得该令牌的所有权限。为了减少盗用，JWT 的有效期应该设置得比较短。对于一些比较重要的权限，使用时应该再次对用户进行认证。
6. 为了减少盗用，JWT 不应该使用 HTTP 协议明码传输，要使用 HTTPS 协议传输。

## 涵盖的技术点

- 使用Kratos开发微服务
- 使用Kratos的BFF与网页端交互
- 使用Kratos的JWT中间件进行登陆认证(Authentication)
- 使用了我自己实现了Kratos的 [Casbin中间件](https://github.com/tx7do/kratos-casbin) 进行API访问权限验证(Authorization)

## 技术栈

- [Kratos](https://go-kratos.dev/)
- [Consul](https://www.consul.io/)
- [Jaeger](https://www.jaegertracing.io/)
- [JWT](https://jwt.io/)
- [Casbin](https://casbin.org/)

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
前台的访问地址: <http://localhost:8081/>

## 参考资料

- [玩转 Kubernetes 权限控制 AuthN/Z](https://caicloud.io/blog/5a1b869275d69a0681e19585)
- [React Typescript Authentication example with Hooks, Axios and Rest API](https://reactjsexample.com/react-typescript-authentication-example-with-hooks-axios-and-rest-api/)
- [JSON Web Token 入门教程](https://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html)
- [AuthN & AuthZ](http://technoponder.com/authentication-authorization/)
- [Using TypeScript with React](https://www.digitalocean.com/community/tutorials/react-typescript-with-react)
- [How To Set Up a New TypeScript Project](https://www.digitalocean.com/community/tutorials/typescript-new-project)