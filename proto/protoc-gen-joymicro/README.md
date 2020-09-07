# protoc-gen-joymicro

go build

将编译出来的protoc-gen-joymicro放到系统变量目录。

编译协议：
protoc --proto_path=. --go_out=. --joymicro_out=. /dir/path/xxx.proto

则go_out插件会生成pb文件，joymicro_out插件会生成服务定义pb文件.
