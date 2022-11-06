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

package keepQueues

import (
	"context"
	infrav1 "github.com/keep-resources/pkg/apis/infra/v1"
	"github.com/wonderivan/logger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	KeepClients "node-simulator/controllers/infra/node-clients"
	//
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// KeepQueueReconciler reconciles a KeepQueue object
type KeepQueueReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=infra.keep.cn,resources=keepqueues,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=infra.keep.cn,resources=keepqueues/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=infra.keep.cn,resources=keepqueues/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps,resources=deployments/status,verbs=get

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KeepQueue object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (kqr *KeepQueueReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	queue := &infrav1.KeepQueue{}
	// keepjob不存在了 删除对应的job
	if err := kqr.Get(ctx, req.NamespacedName, queue); err != nil {
		logger.Warn("keep queue:", req.NamespacedName, " has been deleted")
		return ctrl.Result{}, err
	}
	queue.Spec.OwnJobs = KeepQueueSorter(queue.Spec.OwnJobs)
	//
	_, err := KeepClients.Client.InfraV1().KeepQueues().Update(context.TODO(), queue, metav1.UpdateOptions{})
	logger.Debug("job:", req.NamespacedName, ",joined queue:", queue.Name)
	if err != nil {
		logger.Error(err)
		return ctrl.Result{}, err
	}
	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KeepQueueReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrav1.KeepQueue{}).
		Complete(r)
}
