package mock_server_cli

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/wonderivan/logger"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	v12 "node-simulator/apis/node/v1"
	node_clients "node-simulator/controllers/node-clients"
	"time"
)

type MockServerResponse struct {
	Data struct {
		MachineId      string `json:"machine_id"`
		TimeStamp      int    `json:"time_stamp"`
		CpuUtilPercent int    `json:"cpu_util_percent"`
		MemUtilPercent int    `json:"mem_util_percent"`
		MemGps         int    `json:"mem_gps"`
		Mkpi           int    `json:"mkpi"`
		NetIn          int    `json:"net_in"`
		NetOut         int    `json:"net_out"`
		DiskIoPercent  int    `json:"disk_io_percent"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Message string `json:"message"`
}

// StartMockServer 启动动态更新节点用量的server
func StartMockServer() {
	client := node_clients.Client
	go func() {
		for {
			time.Sleep(10 * time.Second)
			nodesInCluster, err := client.InfraV1().Nodes(v1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
			if err != nil {
				logger.Fatal(err)
				return
			}
			logger.Info("Request for new usage...")
			for _, node := range nodesInCluster.Items {
				go func(node v12.Node) {
					usage, err := GetVirtualNodeTimeStampUsage(node.Spec.NodeName)
					if err != nil {
						logger.Error(err)
						return
					}
					node.Status.MemUsage = usage.Data.MemUtilPercent
					node.Status.CpuUsage = usage.Data.CpuUtilPercent
					node.Status.DiskUsage = usage.Data.DiskIoPercent

					_, err = client.InfraV1().Nodes(node.Namespace).UpdateStatus(context.Background(), &node, metav1.UpdateOptions{})
					if err != nil {
						logger.Error(err)
					}
				}(node)
			}
		}
	}()
}

func GetVirtualNodeTimeStampUsage(nodename string) (*MockServerResponse, error) {
	logger.Debug("query for node: " + nodename)
	if nodename == "" {
		return nil, errors.New("nodename is empty,please check")
	}
	url := "http://10.106.49.175:8080/query?nodename=" + nodename
	//url := "http:/keep-mock-server.keep-colocation-mock-server:31880/query?nodename=" + nodename
	//url := "http://172.17.14.232:30768/query?nodename=" + nodename
	method := "POST"
	//  mock-server.keep-colocation-mock-server:31880/query
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Api-Key", "<API Key>")

	res, err := client.Do(req)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	resp := &MockServerResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return resp, nil
}
