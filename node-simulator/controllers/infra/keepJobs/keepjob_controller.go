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

package keepJobs

import (
	"context"
	infrav1 "github.com/keep-resources/pkg/apis/infra/v1"
	"github.com/wonderivan/logger"
	kapps "k8s.io/api/apps/v1"
	kbatch "k8s.io/api/batch/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	KeepClients "node-simulator/controllers/infra/node-clients"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"time"
)

// KeepJobReconciler reconciles a KeepJob object
type KeepJobReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=infra.keep.cn,resources=keepjobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=infra.keep.cn,resources=keepjobs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=infra.keep.cn,resources=keepjobs/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KeepJob object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (kjr *KeepJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = log.FromContext(ctx)
	var err error
	logger.Debug("start new round keepjob reconcile")
	obj := &infrav1.KeepJob{}
	// keepjob不存在了 删除对应的job
	if err := kjr.Get(ctx, req.NamespacedName, obj); err != nil {
		// delete keepJob
		labelKey, labelVal, err := GenerateUniqueJobLable(obj)
		logger.Debug("keepjob ", req.NamespacedName, "not found,start to delete jobs with label: "+labelKey+":"+labelVal)
		foundJob := &kbatch.Job{}
		if err := kjr.Get(ctx, types.NamespacedName{Name: req.Name, Namespace: req.Namespace}, foundJob); err == nil {
			if err != nil {
				logger.Error(err)
				return ctrl.Result{RequeueAfter: 5 * time.Second}, err
			}
			errDel := kjr.DeleteJobWithLabel(map[string]string{labelKey: labelVal})
			if errDel != nil {
				logger.Error(errDel)
			}
		} else {
			logger.Error("could not find job:", req.NamespacedName, err)
		}
		return ctrl.Result{}, err
	} else {
		// 找到了就将任务放到声明的queue里边
		//先找到对应的queue，如果没有就算了
		//KeepClients.Client.InfraV1().
		if queue, err := KeepClients.Client.InfraV1().KeepQueues().Get(context.TODO(), obj.Spec.JobQueueName, v12.GetOptions{}); err != nil {
			logger.Warn("job declared keepQueue:", obj.Spec.JobQueueName, " but not found,", err)
			return ctrl.Result{RequeueAfter: 5 * time.Second}, err
		} else {
			// 找到了对应的queue就将keepjob加入改queue
			queue.Spec.OwnJobs = append(queue.Spec.OwnJobs, *obj)
			_, err = KeepClients.Client.InfraV1().KeepQueues().Update(context.TODO(), queue, v12.UpdateOptions{})
			if err != nil {
				logger.Error(err)
				return ctrl.Result{}, err
			}
		}
	}
	return ctrl.Result{}, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *KeepJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrav1.KeepJob{}).
		Owns(&kapps.Deployment{}).
		Complete(r)
}
