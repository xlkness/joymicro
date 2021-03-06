package micro

import (
	pb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	//options "google.golang.org/genproto/googleapis/api/annotations"
	"joynova.com/joynova/joymicro/proto/protoc-gen-joymicro/generator"
)

func (g *joymicro) peerGenerateRegisterServiceHandler(file *generator.FileDescriptor,
	service *pb.ServiceDescriptorProto, index int) {
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
	// 驼峰服务Service Interface
	servAlias := servName + "Handler"
	servAliasInterface := servName + "HandlerInterface"

	// 封装服务
	//wrapServAlias := "wrap" + servAlias

	g.P("func Register", servAlias, "(s *server.ServicesManager, hdlr ", servAliasInterface, ") error {")
	//g.P("whdlr := &", wrapServAlias, "{")
	//g.P("h: hdlr,")
	//g.P("}")
	g.P("return s.RegisterOneService(serviceName, hdlr)")
	g.P("}")
	g.P()
}

func (g *joymicro) peerGenerateServiceNewAndRunHandler(file *generator.FileDescriptor,
	service *pb.ServiceDescriptorProto, index int, hasPeer2Peer bool, hashCHash bool) {
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
	// 驼峰服务Service Interface
	servAlias := servName + "Handler"
	servAliasInterface := servName + "HandlerInterface"

	// 首字母小写本地handler名
	servHandlerAliasUnexport := unexport(servName) + "HandlerLocal"

	// 封装服务
	//wrapServAlias := "wrap" + servAlias

	if hasPeer2Peer {
		g.P("func New", servAlias, "AndRun(nodeKey, listenAddr, exposeAddr string, etcdAddrs []string, handler ", servAliasInterface,
			", runErrHandleFun func(error), isLocal bool) (*server.ServicesManager, error) {")
	} else {
		g.P("func New", servAlias, "AndRun(listenAddr, exposeAddr string, etcdAddrs []string, handler ", servAliasInterface,
			", runErrHandleFun func(error), isLocal bool) (*server.ServicesManager, error) {")
	}

	g.P("if isLocal {")
	g.P(servHandlerAliasUnexport, " = handler")
	g.P("return nil, nil")
	g.P("} else {")
	if hasPeer2Peer {
		g.P("s, err := server.NewWithKey(nodeKey, exposeAddr, etcdAddrs)")
	} else {
		g.P("s, err := server.New(exposeAddr, etcdAddrs)")
	}
	g.P("if err != nil {")
	g.P("return nil, err")
	g.P("}")
	g.P()
	g.P("err = Register", servAlias, "(s, handler)")
	g.P("if err != nil {")
	g.P("return nil, err")
	g.P("}")
	g.P()
	g.P("go func(){")
	g.P("err = s.Run(listenAddr)")
	g.P("if err != nil {")
	g.P("runErrHandleFun(err)")
	g.P("}")
	g.P("}()")
	g.P()
	g.P("return s, nil")
	g.P("}")
	g.P("}")
	// g.P("func Register", servAlias, "(s *server.ServicesManager, hdlr ", servAliasInterface, ") error {")
	//g.P("whdlr := &", wrapServAlias, "{")
	//g.P("h: hdlr,")
	//g.P("}")
	// g.P("return s.RegisterOneService(serviceName, hdlr)")
	// g.P("}")
	// g.P()
}

func (g *joymicro) peerGenerateServiceHandlerInterface(file *generator.FileDescriptor,
	service *pb.ServiceDescriptorProto, index int) {
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
	servAlias := servName + "HandlerInterface"

	g.P("type ", servAlias, " interface {")
	for _, method := range service.Method {
		// 服务名
		//serviceNameService := serviceName + "Service"
		// 接收器名
		//receiver := "func (h *" + servAliasUnexport + ")"
		// 方法名
		methName := generator.CamelCase(method.GetName())
		// 协议请求参数
		methodInArg := g.typeName(method.GetInputType())
		// 协议响应参数
		methodOutArg := g.typeName(method.GetOutputType())

		g.P(methName, "(context.Context, *", methodInArg, ", *", methodOutArg, ") error")
	}
	g.P("}")
	g.P()
}
