package user

import (
	"colocation_backend/controllers/user"
)
import "github.com/astaxie/beego"

func UserRouter() {
	beego.Router("api/usr/login", &user.UserController{}, "post:Login")
}
