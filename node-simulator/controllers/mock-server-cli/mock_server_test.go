package mock_server_cli

import (
	"context"
	"github.com/wonderivan/logger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientset "node-simulator/generated/node/clientset/versioned"
	"testing"
)

func TestGetVirtualNodeTimeStampUsage(t *testing.T) {
	usage, err := GetVirtualNodeTimeStampUsage("m_6")
	if err != nil {
		return
	}
	logger.Info(usage)
}

func TestNodeStatus(t *testing.T) {
	var config *restclient.Config
	//config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	config, err := clientcmd.BuildConfigFromFlags("", "cluster_config")
	client, err := clientset.NewForConfig(config)
	if err != nil {
		logger.Error(err)
		return
	}
	node, err := client.InfraV1().Nodes(metav1.NamespaceDefault).Get(context.Background(), "m-6", metav1.GetOptions{})
	logger.Error(node.Status)
}

func TestAcc(t *testing.T) {
	var config *restclient.Config
	//config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	config, err := clientcmd.BuildConfigFromFlags("", "cluster_config")
	client, err := clientset.NewForConfig(config)
	if err != nil {
		logger.Error(err)
		return
	}
	node, err := client.InfraV1().Nodes(metav1.NamespaceDefault).Get(context.Background(), "m-6", metav1.GetOptions{})
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(node.Status.CpuUsage)
	node.Status.CpuUsage = 97
	node.Spec.NodeName = "m-6"
	_, err = client.InfraV1().Nodes(metav1.NamespaceDefault).Update(context.Background(), node, metav1.UpdateOptions{})
	client.InfraV1().Nodes(node.Namespace).UpdateStatus(context.Background(), node, metav1.UpdateOptions{})
	if err != nil {
		logger.Error(err)
		return
	}
	node, err = client.InfraV1().Nodes(metav1.NamespaceDefault).Get(context.Background(), "m-6", metav1.GetOptions{})
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(node.Status)
}
