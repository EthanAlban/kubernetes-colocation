package namespace

import (
	"colocation_backend/controllers"
	client "colocation_backend/pkg/client"
	"context"
	"github.com/wonderivan/logger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NamespaceController struct {
	controllers.BaseController
}

func (nsc *NamespaceController) GetAllNamespace() {
	kubeClient := client.KubeClient
	nslist, err := kubeClient.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error("list ns failed,err:", err)
		nsc.JsonResult(500, "list ns failed,err:"+err.Error(), nil)
	}
	nsc.JsonResult(200, "ok", nslist)
}
