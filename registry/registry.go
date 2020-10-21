package registry

var NameSpace = ""

var DefaultBaseDir = "joymicro/services"

func SetNameSpace(ns string) {
	NameSpace = ns
}

func SetDefaultBaseDir(dir string) {
	DefaultBaseDir = dir
}
