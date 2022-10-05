package routers

import (
	"github.com/astaxie/beego"
	"mock-server/controllers/mockServer"
)

func init() {
	beego.Router("/query", &mockServer.MockServerController{}, "post:PublishUsage")
}
