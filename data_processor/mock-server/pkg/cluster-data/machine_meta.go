package cluster_data

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/wonderivan/logger"
	"io"
	"os"
	"strconv"
)

var CLUSTE_META_FILE string

type ClusterMeta struct {
	MachineId      string `json:"machine_id"`
	TimeStamp      int64  `json:"time_stamp"`
	Failuredomain1 int64  `json:"failure_domain_1"`
	Failuredomain2 string `json:"failure_domain_2"`
	CpuNum         int64  `json:"cpu_num"`
	MemSize        int64  `json:"mem_size"`
	Status         string `json:"status"`
}

func init() {
	CLUSTE_META_FILE = "D:\\毕设\\集群数据\\alibaba_clusterdata_v2018\\machine_meta.csv"
	CLUSTE_USAGE_FILE = "D:\\毕设\\集群数据\\alibaba_clusterdata_v2018\\machine_usage.csv"
}

func BuildFirstThousandsRecords() {
	metas := ReadClusterMetaFromCsv(CLUSTE_META_FILE)
	for _, meta := range metas {
		js, err := json.Marshal(meta)
		if err != nil {
			logger.Error(err)
			continue
		}
		fmt.Println(string(js))
	}

}

func ReadClusterMetaFromCsv(FileName string) []*ClusterMeta {
	fd, err := os.Open(FileName)
	if err != nil {
		logger.Error("Failed to open file", err)
		return []*ClusterMeta{}
	}
	defer fd.Close()
	reader := csv.NewReader(fd)
	// 对大文件进行一行一行的读取
	recordes := make([]*ClusterMeta, 0)
	for {
		row, err := reader.Read()
		//line, col := reader.FieldPos(len(row))
		if err != nil && err != io.EOF {
			logger.Error(err)
			return nil
		}
		if err == io.EOF {
			break
		}
		timeStamp, _ := strconv.Atoi(row[1])
		failureDomain1, _ := strconv.Atoi(row[2])
		cpuNum, _ := strconv.Atoi(row[4])
		memSize, _ := strconv.Atoi(row[5])
		recordes = append(recordes, &ClusterMeta{
			MachineId:      row[0],
			TimeStamp:      int64(timeStamp),
			Failuredomain1: int64(failureDomain1),
			Failuredomain2: row[3],
			CpuNum:         int64(cpuNum),
			MemSize:        int64(memSize),
			Status:         row[6],
		})
	}
	return recordes
}
