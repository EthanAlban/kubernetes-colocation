package virtualNode

import (
	"colocation_backend/controllers"
	"context"
	"github.com/wonderivan/logger"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
import client "colocation_backend/pkg/client"

type VirtualNodeController struct {
	controllers.BaseController
}

func (vnc *VirtualNodeController) GetAllVirtualNodes() {
	nodeClient := client.Client
	nodeClient.InfraV1()
	nodes, err := nodeClient.InfraV1().VirtualNodes(v1.NamespaceAll).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error(err)
		vnc.JsonResult(500, err.Error(), nil)
	}
	vnc.JsonResult(200, "ok", nodes)
}
