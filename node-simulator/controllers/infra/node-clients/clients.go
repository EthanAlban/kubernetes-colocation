package node_clients

import (
	"github.com/keep-resources/pkg/generated/clientset/versioned"
	"github.com/wonderivan/logger"
	"k8s.io/client-go/kubernetes"

	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var Config *restclient.Config
var Client *versioned.Clientset
var KubeClient *kubernetes.Clientset

func init() {
	//config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	var err error
	//Config, err = clientcmd.BuildConfigFromFlags("", "cluster_config")
	//Config, err = clientcmd.BuildConfigFromFlags("", "min_config")
	Config, err = clientcmd.BuildConfigFromFlags("", "actual_config")
	if err != nil {
		Config, err = restclient.InClusterConfig()
		if err != nil {
			logger.Fatal(err)
			return
		}
	}
	//logger.Info(Config)
	Client, err = versioned.NewForConfig(Config)
	KubeClient, err = kubernetes.NewForConfig(Config)
	if err != nil {
		logger.Error(err)
		return
	}
}
