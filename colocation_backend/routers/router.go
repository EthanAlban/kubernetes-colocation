package routers

import (
	"colocation_backend/routers/core"
	"colocation_backend/routers/keepJob"
	KeepJobGroup "colocation_backend/routers/keepJobGroup"
	"colocation_backend/routers/keepQueue"
	"colocation_backend/routers/public"
	"colocation_backend/routers/user"
	"colocation_backend/routers/virtualNode"
)

func init() {
	public.PublicRouter()
	user.UserRouter()
	virtualNode.VirtualNodeRouter()
	keepJob.KeepJobRouter()
	keepQueue.KeepQueueRouter()
	core.NamespaceRouter()
	KeepJobGroup.KeepJobGroupRouter()
}
