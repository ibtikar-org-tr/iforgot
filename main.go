package main

import (
	"github.com/ibtikar-org-tr/iforgot/initializers"
	"github.com/ibtikar-org-tr/iforgot/lib"
	"github.com/ibtikar-org-tr/iforgot/routers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()
	initializers.SrvInit()
	lib.SetupLogOutput()
}

func main() {
	routers.AllRouters()
}
