package public

import (
	"colocation_backend/controllers/public"
	"github.com/astaxie/beego"
)

func PublicRouter() {
	beego.Router("api/utils/heartbeat", &public.PublicController{}, "get:Healthz")
}
