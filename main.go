package main

import (
	"memorial01/conf"
	"memorial01/routers"
)

func main() {
	conf.Init()
	r := routers.NewRouter()
	_ = r.Run(":9090")
}
