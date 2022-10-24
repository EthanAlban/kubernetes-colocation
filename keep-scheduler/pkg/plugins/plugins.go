package plugins

import (
	"context"
	"github.com/wonderivan/logger"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const Name = "keep"

type KeepScheduler struct {
	handle framework.Handle
}

var (
	//_ framework.QueueSortPlugin = &Yoda{}
	_ framework.FilterPlugin = &KeepScheduler{}
	//_ framework.PreScorePlugin  = &Yoda{}
	//_ framework.ScorePlugin     = &Yoda{}
	//_ framework.ScoreExtensions = &Yoda{}

	scheme = runtime.NewScheme()
)

func (ks *KeepScheduler) Name() string {
	return Name
}

func (ks *KeepScheduler) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	logger.Debug("filter pod: %v, node: %v", pod.Name, nodeInfo.Node().Name)
	return framework.NewStatus(framework.Success, "")
}

func New(_ runtime.Object, _ framework.Handle) (framework.Plugin, error) {
	return &KeepScheduler{}, nil
}
