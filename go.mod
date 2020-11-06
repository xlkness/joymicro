module joynova.com/joynova/joymicro

go 1.15

//replace google.golang.org/grpc => google.golang.org/grpc v1.29.0

//replace github.com/go-redis/redis/v8 v8.0.0-beta.9 => gopkg.in/redis.v8 v8.0.0-beta.9

//replace gopkg.in/redis.v8 v8.0.0-beta.9 => github.com/go-redis/redis/v8 v8.0.0-beta.9
replace google.golang.org/grpc => google.golang.org/grpc v1.29.0

require (
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/edwingeng/doublejump v0.0.0-20200219153503-7cfc0ed6e836
	github.com/gogo/protobuf v1.3.1
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0 // indirect
	github.com/jaegertracing/jaeger-client-go v2.25.0+incompatible // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/opentracing/opentracing-go v1.1.1-0.20190913142402-a7454ce5950e
	github.com/openzipkin/zipkin-go v0.2.5
	github.com/prometheus/client_golang v1.7.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/smallnest/rpcx v0.0.0-20200904040754-162b6648c998
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200122045848-3419fae592fc // indirect
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	go.uber.org/zap v1.13.0 // indirect
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899 // indirect
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f // indirect
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2 // indirect
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/appengine v1.6.1 // indirect
	google.golang.org/genproto v0.0.0-20200904004341-0bd0a958aa1d
	google.golang.org/grpc v1.31.0 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/ini.v1 v1.44.0 // indirect
)
