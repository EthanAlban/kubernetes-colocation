package routers

import (
	"colocation_backend/routers/public"
	"colocation_backend/routers/user"
)

func init() {
	public.PublicRouter()
	user.UserRouter()
}
