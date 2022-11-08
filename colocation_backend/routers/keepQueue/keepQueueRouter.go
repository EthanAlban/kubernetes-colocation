package keepQueue

import (
	"colocation_backend/controllers/keepQueue"
	"github.com/astaxie/beego"
)

func KeepQueueRouter() {
	beego.Router("api/keep_queue/all", &keepQueue.KeepQueueController{}, "get:GetAllKeepQueueNames")
}
