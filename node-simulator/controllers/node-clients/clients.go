package node_clients

import (
	"github.com/wonderivan/logger"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientset "node-simulator/generated/node/clientset/versioned"
)

var Config *restclient.Config
var Client *clientset.Clientset

func init() {
	//config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	var err error
	//Config, err = clientcmd.BuildConfigFromFlags("", "cluster_config")
	Config, err = clientcmd.BuildConfigFromFlags("", "min_config")
	if err != nil {
		Config, err = restclient.InClusterConfig()
		if err != nil {
			logger.Fatal(err)
			return
		}
	}
	//logger.Info(Config)
	Client, err = clientset.NewForConfig(Config)
	if err != nil {
		logger.Error(err)
		return
	}
}
