package plugins

import (
	"context"
	"errors"
	infrav1 "github.com/keep-resources/pkg/apis/infra/v1"
	"github.com/wonderivan/logger"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	client2 "sigs.k8s.io/controller-runtime/pkg/client"
)

const Name = "keep"

type KeepScheduler struct {
	handle framework.Handle
	cache  cache.Cache
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
	virtualNodeList := &infrav1.VirtualNodeList{}
	err := ks.cache.List(context.TODO(), virtualNodeList, &client2.ListOptions{})
	if err != nil {
		logger.Error(err)
		return framework.NewStatus(framework.Unschedulable, "can't get the virtualNodes from cache")
	}
	logger.Debug(virtualNodeList)
	//client := clients.Client
	//nodesInCluster, err := client.InfraV1().VirtualNodes(v1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//logger.Debug("find virtualNodes:", nodesInCluster)
	return framework.NewStatus(framework.Success, "")
}

func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	mgrConfig := ctrl.GetConfigOrDie()
	mgrConfig.QPS = 1000
	mgrConfig.Burst = 1000
	if err := infrav1.AddToScheme(scheme); err != nil {
		logger.Fatal(err)
	}
	mgr, err := ctrl.NewManager(mgrConfig, ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: "",
		LeaderElection:     false,
		Port:               9443,
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	go func() {
		if err = mgr.Start(ctrl.SetupSignalHandler()); err != nil {
			logger.Error(err)
			panic(err)
		}
	}()
	infraCache := mgr.GetCache()
	if infraCache.WaitForCacheSync(context.TODO()) {
		logger.Debug("infra chache successed")
		return &KeepScheduler{
			handle: h,
			cache:  infraCache,
		}, nil
	} else {
		err = errors.New("infra cache sync failed")
		logger.Error(err)
		return &KeepScheduler{}, err
	}
}
