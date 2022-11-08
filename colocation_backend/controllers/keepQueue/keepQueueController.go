package keepQueue

import (
	"colocation_backend/controllers"
	client "colocation_backend/pkg/client"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KeepQueueController struct {
	controllers.BaseController
}

func (kqc *KeepQueueController) GetAllKeepQueueNames() {
	keepClient := client.KeepClient
	queuelist, err := keepClient.InfraV1().KeepQueues().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		kqc.JsonResult(500, "get keepqueue failed", err)
	}
	list := make([]string, len(queuelist.Items))
	for k, v := range queuelist.Items {
		list[k] = v.Spec.QueueName
	}
	kqc.JsonResult(200, "ok", list)
}
