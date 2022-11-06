package tests

import (
	"fmt"
	"node-simulator/controllers/infra/keepQueues"
	"testing"
)

func TestScoreNormalize(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(keepQueues.NormalizeScore(int64(-i * 10)))
	}
}
