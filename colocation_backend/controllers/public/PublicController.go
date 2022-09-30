package public

import "colocation_backend/controllers"

type PublicController struct {
	controllers.BaseController
}

func (this *PublicController) Healthz() {
	this.JsonResult(200, "ok", nil)
}
