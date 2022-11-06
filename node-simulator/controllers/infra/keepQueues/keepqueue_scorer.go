package keepQueues

import (
	v1 "github.com/keep-resources/pkg/apis/infra/v1"
	"math"
)

// KeepQueuedJobScorer score a job in the queue to support it being sorted
func KeepQueuedJobScorer(job v1.KeepJob) float64 {
	rawScore := int64(0)
	// the later a job arrived the lower socre will get
	timeScore := -job.Spec.CreatingTime.Time.Unix() * ScoreWeightArriveTime
	rawScore += timeScore
	// the higher priority the more score will get
	priorityScore := job.Spec.Priority * job.Spec.Weight
	rawScore += int64(priorityScore)
	// todo other socre policy to add
	return NormalizeScore(rawScore)
}

// NormalizeScore 归一化
// 需要注意的是如果想映射的区间为[-0.5,0.5]，则数据都应该大于等于0，小于0的数据将被映射到[-1,0]区间上，而并非所有数据标准化的结果都映射到[0,1]区间上
func NormalizeScore(rawScore int64) float64 {
	return math.Atan(float64(rawScore)/float64(1000)) / math.Pi
}
