package main

import (
	_ "colocation_backend/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run("0.0.0.0:9000")
}
