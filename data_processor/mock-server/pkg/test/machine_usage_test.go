package test

import (
	"fmt"
	"mock-server/pkg/cluster-data"
	"testing"
)

func TestBuildUsageRecords(t *testing.T) {
	cluster_data.BuildUsageRecords()
}

func TestGetAllMachines(t *testing.T) {
	machines := cluster_data.GetAllMachines("D:\\毕设\\集群数据\\alibaba_clusterdata_v2018\\machine_usage.csv")
	fmt.Println(machines)
}
