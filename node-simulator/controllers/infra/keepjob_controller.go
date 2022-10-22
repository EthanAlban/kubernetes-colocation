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
	kapps "k8s.io/api/apps/v1"
	kbatch "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"node-simulator/controllers/infra/keepJobs"
	node_clients "node-simulator/controllers/infra/node-clients"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
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
func (r *KeepJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var keepJob infrav1.KeepJob
	_ = log.FromContext(ctx)
	obj := &infrav1.KeepJob{}
	deployment := &kapps.Deployment{}
	err := controllerutil.SetControllerReference(&keepJob, deployment, r.Scheme)
	if err != nil {
		return ctrl.Result{}, err
	}
	// keepjob不存在了 删除对应的job
	if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
		// delete keepJob
		logger.Debug("keepjob ", req.NamespacedName, "not found")
		foundJob := &kbatch.Job{}
		if err := r.Get(ctx, types.NamespacedName{Name: req.Name, Namespace: req.Namespace}, foundJob); err == nil {
			errDel := node_clients.KubeClient.BatchV1().Jobs(foundJob.Namespace).Delete(ctx, foundJob.Name, metav1.DeleteOptions{})
			if errDel != nil {
				logger.Error(errDel)
			}
		} else {
			logger.Error(err)
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	} else {
		// 找到了就拉起对应的keepjob的job
		foundJob := &kbatch.Job{}
		if err = r.Get(ctx, types.NamespacedName{Name: obj.Spec.JobName, Namespace: obj.Namespace}, foundJob); err != nil && errors.IsNotFound(err) {
			errCreate := keepJobs.CreateNewJobForKeepJob(obj)
			logger.Debug("start reconcile new keepjob ", obj.Spec.JobName)
			if errCreate != nil {
				logger.Error(errCreate)
			}
		}
	}
	if err != nil {
		logger.Error(err)
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KeepJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrav1.KeepJob{}).
		Owns(&kapps.Deployment{}).
		Complete(r)
}
