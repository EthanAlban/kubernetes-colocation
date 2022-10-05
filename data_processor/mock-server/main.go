package main

import (
	"github.com/astaxie/beego"
	_ "mock-server/routers"
)

func main() {
	beego.Run("0.0.0.0:8080")
}
