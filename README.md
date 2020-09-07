
- **编译协议生成插件**

`cd proto/protoc-gen-joymicro && go build`

将`protoc-gen-joymicro`复制到全局环境变量

- **编译joymicro工具**

`cd tool/joymicro && go build`

- **生成新服务**

`./joymicro new hello`

生成hello服务的协议和代码，按照输出日志操作一下即可运行服务器。

