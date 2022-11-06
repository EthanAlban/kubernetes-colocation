package keepQueues

import (
	v1 "github.com/keep-resources/pkg/apis/infra/v1"
	"sort"
)

func KeepQueueSorter(jobs []v1.KeepJob) []v1.KeepJob {
	sort.Slice(jobs, func(i, j int) bool {
		aJob, bJob := jobs[i], jobs[j]
		return KeepQueuedJobScorer(aJob) < KeepQueuedJobScorer(bJob)
	})
	return jobs
}
