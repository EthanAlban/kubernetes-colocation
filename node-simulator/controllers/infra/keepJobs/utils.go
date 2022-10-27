package keepJobs

import (
	"github.com/google/uuid"
	v1 "github.com/keep-resources/pkg/apis/infra/v1"
)

func GenerateUniqueJobLable(job *v1.KeepJob) (string, string, error) {
	key, val := "", ""
	namespacedName := job.Spec.JobName + "/" + job.Spec.JobName
	rand, err := uuid.FromBytes([]byte(namespacedName))
	if err != nil {
		return "", "", err
	}
	key = namespacedName
	val = namespacedName + "-" + rand.String()
	return key, val, nil
}
