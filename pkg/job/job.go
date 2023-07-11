package job

import (
	"github.com/go-co-op/gocron"
	"time"
)

type Task struct {
	*gocron.Scheduler
}

func New() *Task {
	return &Task{
		gocron.NewScheduler(time.UTC),
	}
}
func (t *Task) Stop() {
	// todo implement me
	t.Stop()
}
