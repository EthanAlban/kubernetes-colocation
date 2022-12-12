package KeepJobGroup

import (
	"colocation_backend/controllers/KeepJobGroup"
	"github.com/astaxie/beego"
)

func KeepJobGroupRouter() {
	beego.Router("api/keepjob_group/all", &keepJobGroup.KeepJobGroupController{}, "get:GetAllKeepJobGroupNames")
}
