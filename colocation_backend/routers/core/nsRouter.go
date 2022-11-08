package core

import (
	"colocation_backend/controllers/core/namespace"
	"github.com/astaxie/beego"
)

func NamespaceRouter() {
	beego.Router("api/core/namespace/all", &namespace.NamespaceController{}, "get:GetAllNamespace")
}
