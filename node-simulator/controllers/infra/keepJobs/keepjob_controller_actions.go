package keepJobs

import (
	"context"
	"github.com/wonderivan/logger"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels2 "k8s.io/apimachinery/pkg/labels"
	clients "node-simulator/controllers/infra/node-clients"
)

func (kjr *KeepJobReconciler) DeleteJobWithLabel(labelMap map[string]string) error {
	labels := labels2.SelectorFromSet(labelMap).String()
	if jobs, err := clients.KubeClient.BatchV1().Jobs(v1.NamespaceAll).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labels,
	}); err != nil {
		return err
	} else {
		for _, job := range jobs.Items {
			err = clients.KubeClient.BatchV1().Jobs(job.Namespace).Delete(context.TODO(), job.Name, metav1.DeleteOptions{})
			if err != nil {
				logger.Error(err)
				break
			}
		}
		return err
	}

}
