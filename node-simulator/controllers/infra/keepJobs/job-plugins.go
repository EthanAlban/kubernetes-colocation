package keepJobs

import (
	"context"
	infrav1 "github.com/keep-resources/pkg/apis/infra/v1"
	"github.com/wonderivan/logger"
	jobv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	node_clients "node-simulator/controllers/infra/node-clients"
)

func CreateNewJobForKeepJob(obj *infrav1.KeepJob) error {
	replicas := int32(obj.Spec.Replica)
	// keepjob has no deployments
	logger.Debug("create job:", obj.Spec.JobName, " for keepJob:", obj.Name)
	_, errCreate := node_clients.KubeClient.BatchV1().Jobs(obj.Spec.Namespace).Create(context.TODO(), &jobv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:            obj.Name,
			Namespace:       obj.Spec.Namespace,
			Labels:          map[string]string{"app": obj.Spec.JobName},
			OwnerReferences: nil,
		},
		Spec: jobv1.JobSpec{
			Parallelism: &replicas,
			Completions: nil,
			Selector:    nil,
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:            obj.Spec.JobName,
					Namespace:       obj.Spec.Namespace,
					Labels:          map[string]string{"app": obj.Spec.JobName},
					OwnerReferences: nil,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{v1.Container{
						Name:  obj.Spec.JobName + "-" + obj.Spec.Image,
						Image: obj.Spec.Image,
					}},
					RestartPolicy: v1.RestartPolicyOnFailure,
				},
			},
		},
	}, metav1.CreateOptions{})
	return errCreate
}
