package tracer

import (
	"context"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/smallnest/rpcx/share"
)

type JaegerOpenTracingClientPlugin struct{}

func (p *JaegerOpenTracingClientPlugin) PreCall(ctx context.Context, servicePath, serviceMethod string, args interface{}) error {
	var span1 opentracing.Span

	// if it is called in rpc service in case that a service calls antoher service,
	// we uses the span in the service context as the parent span.
	parentSpan := ctx.Value(share.OpentracingSpanServerKey)
	if parentSpan != nil {
		//fmt.Printf("parent:%+v\n", parentSpan)
		span1 = opentracing.StartSpan(
			"rpc.client."+servicePath+"."+serviceMethod,
			opentracing.ChildOf(parentSpan.(opentracing.Span).Context()))
	} else {
		wireContext, err := share.GetSpanContextFromContext(ctx)
		if err == nil && wireContext != nil { //try to parse span from request
			span1 = opentracing.StartSpan(
				"rpc.client."+servicePath+"."+serviceMethod, ext.RPCServerOption(wireContext))
		} else { // parse span from context or create root context
			span1, _ = opentracing.StartSpanFromContext(ctx,
				"rpc.client."+servicePath+"."+serviceMethod)
		}
	}

	if rpcxContext, ok := ctx.(*share.Context); ok {
		rpcxContext.SetValue(share.OpentracingSpanClientKey, span1)
	}

	return nil
}
func (p *JaegerOpenTracingClientPlugin) PostCall(ctx context.Context, servicePath, serviceMethod string, args interface{}, reply interface{}, err error) error {
	if rpcxContext, ok := ctx.(*share.Context); ok {
		span1 := rpcxContext.Value(share.OpentracingSpanClientKey)
		if span1 != nil {
			span1.(opentracing.Span).Finish()
		}
	}
	return nil
}
