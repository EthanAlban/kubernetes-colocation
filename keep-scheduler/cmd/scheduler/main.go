package main

import (
	"fmt"
	"github.com/wonderivan/logger"
	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
	"math/rand"
	"os"
	"sigs.k8s.io/scheduler-plugins/pkg/plugins"
	"time"
)

func main() {
	logger.Debug("Scheduler starting...")
	rand.Seed(time.Now().UTC().UnixNano())
	command := app.NewSchedulerCommand(
		app.WithPlugin(plugins.Name, plugins.New),
	)

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

}
