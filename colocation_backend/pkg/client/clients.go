package node_clients

import (
	clientset "github.com/keep-resources/pkg/generated/clientset/versioned"
	"github.com/wonderivan/logger"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var Config *restclient.Config
var KeepClient *clientset.Clientset
var KubeClient *kubernetes.Clientset

func init() {
	//config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	var err error
	//Config, err = clientcmd.BuildConfigFromFlags("", "cluster_config")
	Config, err = clientcmd.BuildConfigFromFlags("", "min_config")
	if err != nil {
		logger.Fatal(err)
	}
	//logger.Info(Config)
	KeepClient, err = clientset.NewForConfig(Config)
	KubeClient, err = kubernetes.NewForConfig(Config)
	if err != nil {
		logger.Error(err)
		return
	}
}
