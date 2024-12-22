package main

import (
	"github.com/ibtikar-org-tr/iforgot/initializers"
	"github.com/ibtikar-org-tr/iforgot/routers"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	routers.AllRouters()
}
