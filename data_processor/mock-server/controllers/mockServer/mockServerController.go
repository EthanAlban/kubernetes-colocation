package mockServer

import (
	"encoding/json"
	"fmt"
	"github.com/wonderivan/logger"
	"mock-server/controllers"
	cluster_data "mock-server/pkg/cluster-data"
	"os"
	"time"
)

type MockServerController struct {
	controllers.BaseController
}

var startTimeStamp time.Time
var machineMap map[string][]*cluster_data.ClusterUsage

func init() {
	startTimeStamp = time.Now()
	machineMap = make(map[string][]*cluster_data.ClusterUsage)
}

// PublishUsage 按时间戳给虚拟节点更新用量信息
func (ms *MockServerController) PublishUsage() {
	nodeName := ms.Ctx.Input.Query("nodename")
	targetSec := int64(time.Now().Sub(startTimeStamp).Seconds())
	if _, ok := machineMap[nodeName]; !ok { // 做应用级别的缓存
		fmt.Println(os.Getwd())
		jsonPath := "./controllers/mockServer/machines/" + nodeName + ".json"

		jsonFile, err := os.Open(jsonPath)
		if err != nil {
			logger.Error("open json file %v error [ %v ]", jsonPath, err)
			return
		}
		defer jsonFile.Close()
		var metas []*cluster_data.ClusterUsage
		decoder := json.NewDecoder(jsonFile)
		err = decoder.Decode(&metas)
		machineMap[nodeName] = metas
	}
	var res *cluster_data.ClusterUsage
	metas := machineMap[nodeName]
	for _, meta := range metas {
		if meta.TimeStamp > targetSec {
			res = meta
			break
		}
	}
	if res == nil {
		res = metas[0]
		startTimeStamp = time.Now()
	}
	ms.JsonResult(200, "ok", res)
}
