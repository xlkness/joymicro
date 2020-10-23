# joymicro

joymicro是使用rpcx框架，根据protobuf服务文件生成golang服务的库。 

##核心组件

### rpcx service

根目录的service&client封装了rpcx服务，使用**etcd**做注册中心，其中`client/selector`集成了两个过程调用的服务节点选择器，一个可以指定调用的服务节点名，用于点对点
通信;一个使用一致性hash，客户端给定hash key即可hash到唯一几个节点，用于请求顺序调用（需服务节点也能串行处理）。
注意：这两个插件不能同时使用！

### protoc-gen-joymicro

进入`proto/protoc-gen-joymicro`执行`go build`，生成的protoc-gen-joymicro可用于protoc编译。例如test.proto
定义了rpc服务，使用`protoc -I./ --proto_path=. --gogofaster_out=. --joymicro_out=. test.proto`即可生成test.pb.go的
协议结构描述文件和test.joymicro.pb.go的rpc服务描述文件（其中gogofaster_out插件非指定，可以任意替换）。

### 命令行调用

protoc-gen-joymicro插件生成的文件中，有一个"Json Handler for Test"的模块，根据服务proto生成
json参数的调用方法，内部其实还是将json转为protobuf结构的调用参去调用正式方法，所以"Json Handler"可以
用于测试调用。项目服务都定义好了之后可以编写可执行程序直接命令行给json数据调用方法。例如
`./xxx -registry=192.168.1.2:2382 -service=shop -method=Buy -d='{"id":1, "num":100, "money":200}'`

### todo  
- 分布式调用链追踪

### 没有什么*用的模板服务生成工具`tool/joymicro`  
- **编译协议生成插件**

`cd proto/protoc-gen-joymicro && go build`

将`protoc-gen-joymicro`复制到全局环境变量

- **编译joymicro工具**

`cd tool/joymicro && go build`

- **生成新服务**

`./joymicro new hello`

生成hello服务的协议和代码，按照输出日志操作一下即可运行服务器。

