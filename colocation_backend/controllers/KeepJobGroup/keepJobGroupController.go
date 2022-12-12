package keepJobGroup

import (
	"colocation_backend/controllers"
	client "colocation_backend/pkg/client"
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KeepJobGroupController struct {
	controllers.BaseController
}

func (kqc *KeepJobGroupController) GetAllKeepJobGroupNames() {
	keepClient := client.KeepClient
	queuelist, err := keepClient.InfraV1().KeepjobGroups(v1.NamespaceAll).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		kqc.JsonResult(500, "get keepqueue failed", err)
	}
	list := make([]string, len(queuelist.Items))
	for k, v := range queuelist.Items {
		list[k] = v.Name
	}
	kqc.JsonResult(200, "ok", list)
}
