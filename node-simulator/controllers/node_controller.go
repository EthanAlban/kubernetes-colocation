/*
Copyright 2022 ethan.

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

package controllers

import (
	"context"
	"fmt"
	"github.com/wonderivan/logger"
	"k8s.io/apimachinery/pkg/runtime"
	infrav1 "node-simulator/apis/node/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// NodeReconciler reconciles a Node object
type NodeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=infra.keep.cn,resources=nodes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=infra.keep.cn,resources=nodes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=infra.keep.cn,resources=nodes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Node object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *NodeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	logger.Info("object", req.NamespacedName)
	// your logic here

	// 1. Print Spec.Detail and Status.Created in log
	obj := &infrav1.Node{}
	if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
		_ = fmt.Errorf("couldn't find object:%s", req.String())
	} else {
		//打印Detail和Created
		logger.Info("Successfully get detail", "Detail", obj.Spec.Detail)
		logger.Info("", "Created", obj.Status.Created)
	}
	// 2. Change Created
	if !obj.Status.Created {
		obj.Status.Created = true
		err := r.Status().Update(ctx, obj)
		if err != nil {
			logger.Error(err)
			return ctrl.Result{}, err
		}
	}
	//if obj.Status.Created{
	//	obj.Status.CpuUsage = 80
	//	obj.Status.MemUsage = 90
	//	r.Status().Update(ctx, obj)
	//}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrav1.Node{}).
		Complete(r)
}
