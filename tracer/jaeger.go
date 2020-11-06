package tracer

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
)

func NewJaegerTracer(reporterServerAddr, serviceName string) (opentracing.Tracer, error) {
	cfg := &jaegerConfig.Configuration{
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst, //固定采样
			Param: 1,                       //1=全采样、0=不采样
		},

		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: reporterServerAddr,
		},

		ServiceName: serviceName,
	}

	tracer, _, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}

	opentracing.SetGlobalTracer(tracer)

	return tracer, nil
}
