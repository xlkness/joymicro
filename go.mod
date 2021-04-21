module joynova.com/joynova/joymicro

go 1.15

//replace google.golang.org/grpc => google.golang.org/grpc v1.29.0

//replace github.com/go-redis/redis/v8 v8.0.0-beta.9 => gopkg.in/redis.v8 v8.0.0-beta.9

//replace gopkg.in/redis.v8 v8.0.0-beta.9 => github.com/go-redis/redis/v8 v8.0.0-beta.9
replace google.golang.org/grpc => google.golang.org/grpc v1.29.0

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/edwingeng/doublejump v0.0.0-20200219153503-7cfc0ed6e836
	github.com/gogo/protobuf v1.3.1
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.4.3
	github.com/opentracing/opentracing-go v1.1.1-0.20190913142402-a7454ce5950e
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/rpcxio/rpcx-etcd v0.0.0-20210408131901-67fd64750268 // indirect
	github.com/smallnest/rpcx v1.6.2
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	github.com/xtaci/lossyconn v0.0.0-20200209145036-adba10fffc37 // indirect
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f // indirect
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
	google.golang.org/genproto v0.0.0-20200904004341-0bd0a958aa1d
	google.golang.org/protobuf v1.25.0
	gopkg.in/ini.v1 v1.44.0 // indirect
)
