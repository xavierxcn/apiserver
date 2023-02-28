package task

import (
	"github.com/robfig/cron"
	"github.com/xavierxcn/apiserver/internal/task/tasks/auto_update"
	"github.com/xavierxcn/apiserver/pkg/log"
)

func Run() {
	c := cron.New()

	err := c.AddFunc("@every 5s", auto_update.Run)
	if err != nil {
		log.Fatalw("AddFunc error", "err", err)
	}

	c.Start()

	log.Info("run task")
}
