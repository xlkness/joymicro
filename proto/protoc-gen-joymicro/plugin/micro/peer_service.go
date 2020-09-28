package micro

import (
	//"google.golang.org/protobuf/types/descriptorpb"
	"strings"

	//"github.com/golang/protobuf/proto"
	pb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	//options "google.golang.org/genproto/googleapis/api/annotations"
	"joynova.com/joynova/joymicro/proto/protoc-gen-joymicro/generator"
)

func (g *joymicro) peerGenerateServiceInterface(file *generator.FileDescriptor,
	service *pb.ServiceDescriptorProto, index int, hasPeer2Peer bool) {
	// 大写的服务名
	origServName := service.GetName()
	// 小写的服务名
	//lowerServiceName := strings.ToLower(service.GetName())
	// 协议文件夹包名
	//serviceName := lowerServiceName
	//if pkg := file.GetPackage(); pkg != "" {
	//	serviceName = pkg
	//}
	// 驼峰服务名
	servName := generator.CamelCase(origServName)
	// 驼峰服务Service
	servAlias := servName + "Service"

	// strip suffix
	if strings.HasSuffix(servAlias, "ServiceService") {
		servAlias = strings.TrimSuffix(servAlias, "Service")
	}

	g.P("type ", servAlias+"Interface", " interface {")
	for _, method := range service.Method {
		// 服务名
		//serviceNameService := serviceName + "Service"
		// 接收器名
		//receiver := "func (c *" + serviceNameService + ")"
		// 方法名
		methName := generator.CamelCase(method.GetName())
		// 协议请求参数
		methodInArg := g.typeName(method.GetInputType())
		// 协议响应参数
		methodOutArg := g.typeName(method.GetOutputType())

		g.P(methName, "(context.Context, *", methodInArg, ") (*", methodOutArg, ", error)")
		if hasPeer2Peer {
			g.P(methName+"Peer", "(context.Context, string, *", methodInArg, ") (*", methodOutArg, ", error)")
		}
		g.P(methName+"All", "(context.Context, *", methodInArg, ") (*", methodOutArg, ", error)")
	}
	g.P("}")
	g.P()
}

func (g *joymicro) peerGenerateNewService(file *generator.FileDescriptor,
	service *pb.ServiceDescriptorProto, index int, hasPeer2Peer bool) {
	// 大写的服务名
	origServName := service.GetName()
	// 小写的服务名
	//lowerServiceName := strings.ToLower(service.GetName())
	// 协议文件夹包名
	//serviceName := lowerServiceName
	//if pkg := file.GetPackage(); pkg != "" {
	//	serviceName = pkg
	//}
	// 驼峰服务名
	servName := generator.CamelCase(origServName)
	// 驼峰服务Service
	servAlias := servName + "Service"

	// strip suffix
	if strings.HasSuffix(servAlias, "ServiceService") {
		servAlias = strings.TrimSuffix(servAlias, "Service")
	}

	// 首字母小写服务
	servAliasUnexport := unexport(servAlias)

	// 封装服务
	//wrapServAlias := "wrap" + servAlias

	g.P("func New", servAlias, "(etcdAddrs []string, timeout time.Duration, isPermanent bool) ", servAlias+"Interface", " {")
	g.P("c := client.New(serviceName, etcdAddrs, timeout, isPermanent)")
	if hasPeer2Peer {
		g.P("// 设置点对点选择器")
		g.P("c.SetSelector(&client.PeerSelector{})")
	}
	g.P("return &", servAliasUnexport, "{")
	g.P("c: c,")
	g.P("}")
	g.P("}")
	g.P()
}

func (g *joymicro) peerGenerateServiceUnexport(file *generator.FileDescriptor,
	service *pb.ServiceDescriptorProto, index int, hasPeer2Peer bool) {
	// 大写的服务名
	origServName := service.GetName()
	// 小写的服务名
	//lowerServiceName := strings.ToLower(service.GetName())
	// 协议文件夹包名
	//serviceName := lowerServiceName
	//if pkg := file.GetPackage(); pkg != "" {
	//	serviceName = pkg
	//}
	// 驼峰服务名
	servName := generator.CamelCase(origServName)
	// 驼峰服务Service
	servAlias := servName + "Service"

	// strip suffix
	if strings.HasSuffix(servAlias, "ServiceService") {
		servAlias = strings.TrimSuffix(servAlias, "Service")
	}

	// 首字母小写服务
	servAliasUnexport := unexport(servAlias)

	// 封装服务
	//wrapServAlias := "wrap" + servAlias

	g.P("type ", servAliasUnexport, " struct {")
	g.P("c *client.Service")
	g.P("}")

	for _, method := range service.Method {
		// 服务名
		//serviceNameService := serviceName + "Service"
		// 接收器名
		receiver := "func (c *" + servAliasUnexport + ")"
		// 方法名
		methName := generator.CamelCase(method.GetName())
		// 协议请求参数
		methodInArg := g.typeName(method.GetInputType())
		// 协议响应参数
		methodOutArg := g.typeName(method.GetOutputType())

		gf := func(callFun string, peer bool) {
			g.P("out := new(", methodOutArg, ")")
			if peer {
				g.P("ctx = context.WithValue(ctx, \"select_key\", peerKey)")
			}
			g.P("err := ", `c.c.`, callFun, `(ctx, "`, method.GetName(), `", in, out)`)
			g.P("if err != nil { ")
			g.P("return nil, err")
			g.P("}")
			g.P("return out, nil")
			g.P("}")
			g.P()
		}

		// 生成负载均衡调用方法
		g.P(receiver, methName, "(ctx context.Context, in *", methodInArg, ") (*", methodOutArg, ", error) {")
		gf("Call", false)

		// 生成点对点调用方法
		if hasPeer2Peer {
			g.P(receiver, methName+"Peer", "(ctx context.Context, peerKey string, in *", methodInArg, ") (*", methodOutArg, ", error) {")
			gf("Call", true)
		}

		// 生成广播调用方法
		g.P(receiver, methName+"All", "(ctx context.Context, in *", methodInArg, ") (*", methodOutArg, ", error) {")
		gf("CallAll", false)
	}
}

func (g *joymicro) peerGenerateWrapService(file *generator.FileDescriptor, service *pb.ServiceDescriptorProto, index int) {
	// 大写的服务名
	origServName := service.GetName()
	// 小写的服务名
	//lowerServiceName := strings.ToLower(service.GetName())
	// 协议文件夹包名
	//serviceName := lowerServiceName
	//if pkg := file.GetPackage(); pkg != "" {
	//	serviceName = pkg
	//}
	// 驼峰服务名
	servName := generator.CamelCase(origServName)
	// 驼峰服务Service
	servAlias := servName + "Service"

	// strip suffix
	if strings.HasSuffix(servAlias, "ServiceService") {
		servAlias = strings.TrimSuffix(servAlias, "Service")
	}

	// 首字母小写服务
	//servAliasUnexport := unexport(servAlias)

	// 封装服务
	wrapServAlias := "wrap" + servAlias

	g.P("type ", wrapServAlias, " struct {")
	g.P("c *", clientPkg, ".Service")
	g.P("}")
	g.P()
	for _, method := range service.Method {
		// 服务名
		//serviceNameService := serviceName + "Service"
		wrapServiceNameService := "wrap" + servAlias
		// 接收器名
		receiver := "func (c *" + wrapServiceNameService + ")"
		// 方法名
		methName := generator.CamelCase(method.GetName())
		// 协议请求参数
		methodInArg := g.typeName(method.GetInputType())
		// 协议响应参数
		methodOutArg := g.typeName(method.GetOutputType())
		wrapMethodInArg := "Wrap" + methodInArg

		// 生成封装参数
		g.P("type ", wrapMethodInArg, " struct {")
		g.P("key string")
		g.P("req *", methodInArg)
		g.P("}")

		// 生成负载均衡调用方法
		g.P(receiver, methName, "(ctx context.Context, in *", methodInArg, ") (*", methodOutArg, ", error) {")
		g.P("wrapIn := &", wrapMethodInArg, "{req: in}")
		//g.P(`req := c.c.Call(ctx, "`, reqMethod, `", in)`)
		g.P("out := new(", methodOutArg, ")")
		// TODO: Pass descExpr to Invoke.
		g.P("err := ", `c.c.Call(ctx, "`, method.GetName(), `", wrapIn, out)`)
		g.P("if err != nil { return nil, err }")
		g.P("return out, nil")
		g.P("}")
		g.P()

		// 生成点对点调用方法
		g.P(receiver, methName+"Peer", "(ctx context.Context, peerKey string, in *", methodInArg, ") (*", methodOutArg, ", error) {")
		g.P("wrapIn := &", wrapMethodInArg, "{key: peerKey, req: in}")
		//g.P(`req := c.c.Call(ctx, "`, reqMethod, `", in)`)
		g.P("out := new(", methodOutArg, ")")
		// TODO: Pass descExpr to Invoke.
		g.P("err := ", `c.c.Call(ctx, "`, method.GetName(), `", wrapIn, out)`)
		g.P("if err != nil { return nil, err }")
		g.P("return out, nil")
		g.P("}")
		g.P()

		// 生成广播调用方法
		g.P(receiver, methName+"All", "(ctx context.Context, in *", methodInArg, ") (*", methodOutArg, ", error) {")
		g.P("wrapIn := &", wrapMethodInArg, "{req: in}")
		//g.P(`req := c.c.Call(ctx, "`, reqMethod, `", in)`)
		g.P("out := new(", methodOutArg, ")")
		// TODO: Pass descExpr to Invoke.
		g.P("//todo 暂时不支持广播")
		g.P("err := ", `c.c.Call(ctx, "`, method.GetName(), `", wrapIn, out)`)
		g.P("if err != nil { return nil, err }")
		g.P("return out, nil")
		g.P("}")
		g.P()
	}
}

func (g *joymicro) peerGenerateTestService(file *generator.FileDescriptor,
	service *pb.ServiceDescriptorProto, index int, hasPeer2Peer bool) {
	// 大写的服务名
	origServName := service.GetName()
	// 小写的服务名
	//lowerServiceName := strings.ToLower(service.GetName())
	// 协议文件夹包名
	//serviceName := lowerServiceName
	//if pkg := file.GetPackage(); pkg != "" {
	//	serviceName = pkg
	//}
	// 驼峰服务名
	servName := generator.CamelCase(origServName)
	// 驼峰服务Service
	servAlias := servName + "JsonTestService"

	// strip suffix
	if strings.HasSuffix(servAlias, "ServiceService") {
		servAlias = strings.TrimSuffix(servAlias, "Service")
	}

	// 首字母小写服务
	//servAliasUnexport := unexport(servAlias)

	// 封装服务
	//wrapServAlias := "wrap" + servAlias

	g.P("//===============================================Json Handler for Test===============================================")
	g.P()
	g.P("func New", servAlias, "(etcdAddrs []string, timeout time.Duration, isPermanent bool) reflect.Type {")
	g.P("c := New", servName, "Service(etcdAddrs, timeout, isPermanent)")
	g.P("return reflect.TypeOf(&", servAlias, "{c: c})")
	g.P("}")
	g.P()

	g.P("type ", servAlias, " struct {")
	g.P("c ", servName, "ServiceInterface")
	g.P("}")

	for _, method := range service.Method {
		// 服务名
		//serviceNameService := serviceName + "Service"
		// 接收器名
		receiver := "func (c *" + servAlias + ")"
		// 方法名
		methName := generator.CamelCase(method.GetName())
		// 协议请求参数
		methodInArg := g.typeName(method.GetInputType())
		// 协议响应参数
		methodOutArg := g.typeName(method.GetOutputType())

		gf := func(method string, peer bool) {
			g.P("newIn := &", methodInArg, "{}")
			g.P("err := json.Unmarshal([]byte(in), newIn)")
			g.P("if err != nil {")
			g.P("return nil, err")
			g.P("}")
			g.P()

			if peer {
				g.P("return c.c.", method, `(ctx, peerKey, newIn)`)
			} else {
				g.P("return c.c.", method, `(ctx, newIn)`)
			}
			//g.P("if err != nil { ")
			//g.P("return nil, err")
			//g.P("}")
			//g.P("return out, nil")
			g.P("}")
			g.P()
		}

		// 生成负载均衡调用方法
		g.P(receiver, methName, "(ctx context.Context, in string) (*", methodOutArg, ", error) {")
		gf(methName, false)

		// 生成点对点调用方法
		if hasPeer2Peer {
			g.P(receiver, methName+"Peer", "(ctx context.Context, peerKey string, in string) (*", methodOutArg, ", error) {")
			gf(methName+"Peer", true)
		}

		// 生成广播调用方法
		g.P(receiver, methName+"All", "(ctx context.Context, in string) (*", methodOutArg, ", error) {")
		gf(methName+"All", false)
	}
}
