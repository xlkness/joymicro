package main

import (
	"fmt"
	"joynova.com/joynova/joymicro/tool/joymicro/cmd"
	"os"
)

func Desc() string {
	return fmt.Sprintf("Description:\n") +
		fmt.Sprintf("   joymicro是一个微服务框架，根据proto3服务定义文件生成服务和对应调用客户端，\n") +
		fmt.Sprintf("   joymicro使用rpcx作为服务调用框架，使用etcd作为注册中心。\n\n") +
		fmt.Sprintf("   使用%v [command] -h|--help或者%v help [command]来获取帮助。\n\n", os.Args[0], os.Args[0])
}

func Usage() string {
	return fmt.Sprintf("Usage:\n") +
		fmt.Sprintf("   %v command [command options] [arguments...]\n\n", os.Args[0])
}

func Commands() string {
	str := fmt.Sprintf("Commands:\n")
	for _, v := range cmds {
		str += v.Desc()
	}

	return str
}

func printHelp() {
	fmt.Printf(Desc())
	fmt.Printf(Usage())
	fmt.Printf(Commands())
}

func checkArgs() {
	if len(os.Args) <= 2 {

	}
}

var cmds []Cmd

func init() {
	cmds = []Cmd{new(cmd.NewHandler)}
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}
	cmd := os.Args[1]
	switch cmd {
	case "-h", "--help":
		printHelp()
	case "help":
		if len(os.Args) == 3 {
			helpCmd := os.Args[2]
			for _, v := range cmds {
				if v.Name() == helpCmd {
					fmt.Printf(v.Help())
					return
				}
			}
		}

		printHelp()
	default:
		for _, v := range cmds {
			if v.Name() == cmd {
				err := v.CheckArgs(os.Args[2:]...)
				if err != nil {
					fmt.Printf("Command Check Argument Error:%v\n\n", err)
					fmt.Printf("==========================Help==========================\n")
					fmt.Printf(v.Help())
					return
				}
				v.Exec(os.Args[2:]...)
				return
			}
		}
		printHelp()
	}
}
