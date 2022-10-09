package virtualNode

import (
	"colocation_backend/controllers/virtualNode"
	"github.com/astaxie/beego"
)

func VirtualNodeRouter() {
	beego.Router("api/virtual_node/all", &virtualNode.VirtualNodeController{}, "get:GetAllVirtualNodes")
}
