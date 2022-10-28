package keepJobs

import (
	"context"
	"errors"
	infrav1 "github.com/keep-resources/pkg/apis/infra/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (kjr *KeepJobReconciler) SetKeepJobDefaultValues(obj *infrav1.KeepJob) error {
	var err error
	nilTime := metav1.Time{}
	if obj.Spec.CreatingTime == nilTime {
		obj.Spec.CreatingTime = metav1.Now()

	}
	if obj.Spec.Weight == 0 {
		obj.Spec.Weight = 1
	}
	err = kjr.Update(context.TODO(), obj)
	if err != nil {
		return errors.New("set default values for job:" + obj.Namespace + "/" + obj.Spec.JobName + " failed:" + err.Error())
	}
	return nil
}
