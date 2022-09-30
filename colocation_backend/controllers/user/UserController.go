package user

import "colocation_backend/controllers"

type UserController struct {
	controllers.BaseController
}

func (this *UserController) Login() {
	this.JsonResult(200, "ok", nil)
}
