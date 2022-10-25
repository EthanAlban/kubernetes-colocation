package node_clients

import (
	"context"
	"github.com/wonderivan/logger"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	client := Client
	time.Sleep(10 * time.Second)
	nodesInCluster, err := client.InfraV1().VirtualNodes(v1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		logger.Fatal(err)
		return
	}
	logger.Debug(nodesInCluster)
}
