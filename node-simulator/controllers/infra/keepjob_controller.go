/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package infra

import (
	"context"
	infrav1 "github.com/keep-resources/pkg/apis/infra/v1"
	"github.com/wonderivan/logger"
	v12 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	node_clients "node-simulator/controllers/infra/node-clients"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// KeepJobReconciler reconciles a KeepJob object
type KeepJobReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=infra.keep.cn,resources=keepjobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=infra.keep.cn,resources=keepjobs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=infra.keep.cn,resources=keepjobs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KeepJob object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *KeepJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	obj := &infrav1.KeepJob{}
	if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
		logger.Debug("keepjob ", req.NamespacedName, " not found")
	} else {
		replicas := int32(obj.Spec.Replica)
		jobSelector := metav1.LabelSelector{
			MatchLabels: map[string]string{"app": obj.Spec.JobName},
		}
		_, err := node_clients.KubeClient.AppsV1().Deployments(req.Namespace).Create(ctx, &v12.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      obj.Spec.JobName,
				Namespace: obj.Namespace,
				Labels:    map[string]string{"app": obj.Spec.JobName},
			},
			Spec: v12.DeploymentSpec{
				Replicas: &replicas,
				Selector: &jobSelector,
				Template: v1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Name:      obj.Spec.JobName,
						Namespace: obj.Spec.Namespace,
						Labels:    map[string]string{"app": obj.Spec.JobName},
					},
					Spec: v1.PodSpec{
						Containers: []v1.Container{v1.Container{
							Name:  obj.Spec.JobName + "-" + obj.Spec.Image,
							Image: obj.Spec.Image,
						}},
						RestartPolicy: v1.RestartPolicyAlways,
					},
				},
			},
		}, metav1.CreateOptions{})
		if err != nil {
			logger.Error(err)
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KeepJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrav1.KeepJob{}).
		Complete(r)
}
