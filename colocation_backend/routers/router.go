package routers

import (
	"colocation_backend/routers/keepJob"
	"colocation_backend/routers/public"
	"colocation_backend/routers/user"
	"colocation_backend/routers/virtualNode"
)

func init() {
	public.PublicRouter()
	user.UserRouter()
	virtualNode.VirtualNodeRouter()
	keepJob.KeepJobRouter()
}
