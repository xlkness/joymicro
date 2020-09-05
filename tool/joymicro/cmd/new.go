package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type NewHandler struct {
}

func (nh *NewHandler) Name() string {
	return "new"
}

func (nh *NewHandler) Desc() string {
	return fmt.Sprintf("   %v\t创建新的微服务\n", nh.Name())
}

func (nh *NewHandler) Help() string {
	return fmt.Sprintf("Desc:\n") +
		fmt.Sprintf("   %v 创建新的微服务\n\n", nh.Name()) +
		fmt.Sprintf("Usage:\n") +
		fmt.Sprintf("   %v new <service>\n\n", os.Args[0]) +
		fmt.Sprintf("Example:\n") +
		fmt.Sprintf("   '%v new shop && cd shop'\n", os.Args[0])
}

func (nh *NewHandler) CheckArgs(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("请输入一个服务名")
	}
	if len(args) > 1 {
		return fmt.Errorf("参数数量错误")
	}

	_, err := os.Stat(args[0])
	if err == nil || os.IsExist(err) {
		return fmt.Errorf("目录%v已存在", args[0])
	}

	return nil
}

func (nh *NewHandler) Exec(args ...string) {
	servicePath := args[0]
	nh.newServiceDir(servicePath)
}

func (nh *NewHandler) newServiceDir(servicePath string) {
	err := os.MkdirAll(servicePath, os.ModePerm)
	if err != nil {
		fmt.Printf("创建文件夹%v错误:%v\n", servicePath, err)
		os.Exit(1)
	}
	fmt.Printf("==> create dir %v\n", servicePath)

	nh.outputHandlerDir(servicePath)

	nh.outputProtoDir(servicePath)

	nh.outputMainFile(servicePath)

	fmt.Printf("\n\n")
	tree(servicePath, 1, true)
}

func (nh *NewHandler) outputHandlerDir(servicePath string) {
	service := filepath.Base(servicePath)
	handlerDir := servicePath + "/handler"
	err := os.MkdirAll(handlerDir, os.ModePerm)
	if err != nil {
		fmt.Printf("创建文件夹%v错误:%v\n", handlerDir, err)
		os.Exit(1)
	}
	fmt.Printf("==> create dir %v\n", handlerDir)

	nh.newHandler(service, handlerDir)
}

func (nh *NewHandler) outputProtoDir(servicePath string) {
	service := filepath.Base(servicePath)
	handlerDir := servicePath + "/proto"
	err := os.MkdirAll(handlerDir, os.ModePerm)
	if err != nil {
		fmt.Printf("创建文件夹%v错误:%v\n", handlerDir, err)
		os.Exit(1)
	}
	fmt.Printf("==> create dir %v\n", handlerDir)

	nh.newProto(service, handlerDir)
}

func (nh *NewHandler) outputMainFile(service string) {
	fileName := service + "/main.go"
	fd, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE,os.ModePerm)
	if err != nil {
		fmt.Printf("创建文件 %v 错误:%v\n", fileName, err)
		os.Exit(1)
	}
	defer fd.Close()

	fmt.Printf("==> create file %v\n", fileName)

	mfd := &myfd{fd}
	mfd.out("package main")
	mfd.out()
	mfd.out()
	mfd.out("import (")
	mfd.out("\t\"joynova.com/joynova/joymicro/service\"")
	mfd.out("\t\"shop/proto\"")
	mfd.out("\t\"shop/handler\"")
	mfd.out(")")
	mfd.out()
	mfd.out()

	mfd.out("func main() {")
	mfd.out("\ts, err := service.New(\":8888\", []string{\"127.0.0.1:2382\"})")
	mfd.out("\t\tif err != nil {")
	mfd.out("\t\tpanic(err)")
	mfd.out("\t}")

	protoService := []rune(service)
	protoService[0] -= 32
	mfd.out("\terr = proto.Register", string(protoService), "Handler(s, new(handler.ShopHandler))")
	mfd.out("\tif err != nil {")
	mfd.out("\t\tpanic(err)")
	mfd.out("\t}")
	mfd.out()
	//mfd.out("\tfmt.Printf(\"start service ...\\n\")")
	mfd.out("\terr = s.Run()")
	mfd.out("\tif err != nil {")
	mfd.out("\t\tpanic(err)")
	mfd.out("\t}")
	mfd.out("}")
}

func (nh *NewHandler) newHandler(service, handlerDir string) {
	fileName := handlerDir + "/handler.go"
	fd, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE,os.ModePerm)
	if err != nil {
		fmt.Printf("创建文件 %v 错误:%v\n", fileName, err)
		os.Exit(1)
	}
	defer fd.Close()

	mfd := &myfd{fd}

	mfd.out("package handler")
	mfd.out()
	mfd.out()

	protoService := []rune(service)
	protoService[0] -= 32
	Service := string(protoService)

	mfd.out("import (")
	mfd.out("\t\"context\"")
	mfd.out("\t\"", service, "/proto\"")
	mfd.out(")")
	mfd.out()

	mfd.out("type ", Service, "Handler struct {")
	mfd.out("}")
	mfd.out()
	mfd.out("func (s *", Service, "Handler) Hello(ctx context.Context, req *proto.Request, res *proto.Response) error {")
	mfd.out("\treturn nil")
	mfd.out("}")
}

func (nh *NewHandler) newProto(service, protoDir string) {
	fileName := protoDir + "/" + service + ".proto"
	fd, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE,os.ModePerm)
	if err != nil {
		fmt.Printf("创建文件 %v 错误:%v\n", fileName, err)
		os.Exit(1)
	}
	defer fd.Close()

	mfd := &myfd{fd}

	mfd.out("syntax = \"proto3\";")
	mfd.out()
	mfd.out()

	mfd.out("package ", "proto", ";")
	mfd.out()

	protoService := []rune(service)
	protoService[0] -= 32
	mfd.out("service ", string(protoService), " {")
	mfd.out("\t rpc Hello(Request) returns (Response) {};")
	mfd.out("}")
	mfd.out()
	mfd.out()

	mfd.out("message Request {")
	mfd.out("\tstring name = 1;")
	mfd.out("}")
	mfd.out()
	mfd.out()
	mfd.out("message Response {")
	mfd.out("\tint32 errCode = 1;")
	mfd.out("\tint32 msg =2 ;")
	mfd.out("}")
}

func tree(dir string, deepth int, prelast bool) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("查看目录%v所有文件错误:%v\n", dir, err)
		os.Exit(1)
	}

	fmt.Printf("%v\n", filepath.Base(dir))
	for i, v := range files {
		if i == len(files) - 1 {
			treeGraph(deepth, prelast, true)
			prelast = true
		} else {
			treeGraph(deepth, prelast, false)
			prelast = false
		}
		if v.IsDir() {
			tree(dir + "/" + v.Name() , deepth+1, prelast)
		} else {
			fmt.Printf("%v\n", filepath.Base(v.Name()))
		}
	}
}

func treeGraph(deepth int, prelast, last bool) {
	if deepth == 1 {
		if last {
			fmt.Printf("└── ")
		} else {
			fmt.Printf("│── ")
		}
	} else {
		if prelast && last {
			for i := 0; i < deepth - 1; i++ {
				fmt.Printf("    ")
			}
			fmt.Printf("└── ")
		} else if !prelast && last {
			fmt.Printf("│   ")
			for i := 0; i < deepth - 2; i++ {
				fmt.Printf(" ")
			}
			fmt.Printf("└── ")
		} else if !prelast && !last {
			fmt.Printf("│   ")
			for i := 0; i < deepth - 2; i++ {
				fmt.Printf(" ")
			}
			fmt.Printf("│── ")
		} else if prelast && !last {
			for i := 0; i < deepth - 1; i++ {
				fmt.Printf("    ")
			}
			fmt.Printf("│── ")
		}
	}
}

type myfd struct {
	fd *os.File
}

func (mf *myfd) out(args...interface{}) {
	str := fmt.Sprint(args...)
	if str != "" {
		mf.fd.WriteString(str)
	}
	mf.fd.WriteString("\n")
}