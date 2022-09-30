package cluster_data

import (
	"encoding/csv"
	"encoding/json"
	"github.com/wonderivan/logger"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

var CLUSTE_USAGE_FILE string

type ClusterUsage struct {
	MachineId      string `json:"machine_id"`
	TimeStamp      int64  `json:"time_stamp"`
	CpuUtilPercent int64  `json:"cpu_util_percent"`
	MemUtilPercent int64  `json:"mem_util_percent"`
	MemGps         int64  `json:"mem_gps"`
	Mkpi           int64  `json:"mkpi"`
	NetIn          int64  `json:"net_in"`
	NetOut         int64  `json:"net_out"`
	DiskIoPercent  int64  `json:"disk_io_percent"`
}

func init() {
	CLUSTE_USAGE_FILE = "D:\\毕设\\集群数据\\alibaba_clusterdata_v2018\\machine_usage.csv"
}

func BuildUsageRecords() {
	metas := ReadClusterUsageFromCsv(CLUSTE_USAGE_FILE, 5000)
	jsonPath := "./m_1933.json"
	dumpMachineUsage(jsonPath, metas)
	sort.Slice(metas, func(i, j int) bool {
		return metas[i].TimeStamp < metas[j].TimeStamp
	})
	p := plot.New()
	p.Title.Text = "Get Started"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	points := make(plotter.XYs, 0)
	for _, meta := range metas {
		pt := plotter.XY{float64(meta.TimeStamp / 60), float64(meta.CpuUtilPercent)}
		points = append(points, pt)
	}
	err := plotutil.AddLinePoints(p, points)
	p.Save(30*vg.Inch, 8*vg.Inch, "points.png")
	if err != nil {
		return
	}

}

func ReadClusterUsageFromCsv(FileName string, expectLines int) []*ClusterUsage {
	fd, err := os.Open(FileName)
	if err != nil {
		logger.Error("Failed to open file", err)
		return []*ClusterUsage{}
	}
	defer fd.Close()
	reader := csv.NewReader(fd)
	// 对大文件进行一行一行的读取
	recordes := make([]*ClusterUsage, 0)
	for {
		row, err := reader.Read()
		line, _ := reader.FieldPos(1)
		if line == expectLines {
			break
		}
		if err != nil && err != io.EOF {
			logger.Error(err)
			return nil
		}
		if err == io.EOF {
			break
		}
		timeStamp, _ := strconv.Atoi(row[1])
		CpuUtilPercent, _ := strconv.Atoi(row[2])
		MemUtilPercent, _ := strconv.Atoi(row[3])
		MemGps, _ := strconv.Atoi(row[4])
		Mkpi, _ := strconv.Atoi(row[5])
		NetIn, _ := strconv.Atoi(row[6])
		NetOut, _ := strconv.Atoi(row[7])
		DiskIoPercent, _ := strconv.Atoi(row[8])
		if row[0] != "m_1932" {
			break
		}
		recordes = append(recordes, &ClusterUsage{
			MachineId:      row[0],
			TimeStamp:      int64(timeStamp) - int64(386640),
			CpuUtilPercent: int64(CpuUtilPercent),
			MemUtilPercent: int64(MemUtilPercent),
			MemGps:         int64(MemGps),
			Mkpi:           int64(Mkpi),
			NetIn:          int64(NetIn),
			NetOut:         int64(NetOut),
			DiskIoPercent:  int64(DiskIoPercent),
		})
	}
	return recordes
}

func dumpMachineUsage(jsonPath string, info []*ClusterUsage) {
	jsonFile, err := os.Create(jsonPath) // 创建 json 文件
	if err != nil {
		log.Printf("create json file %v error [ %v ]", jsonPath, err)
		return
	}
	defer jsonFile.Close()

	encode := json.NewEncoder(jsonFile) // 创建编码器
	err = encode.Encode(info)           // 编码
	if err != nil {
		log.Printf("encode error [ %v ]", err)
		return
	}
}
