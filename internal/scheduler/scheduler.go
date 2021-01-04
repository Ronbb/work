package scheduler

import (
	"context"

	"github.com/ronbb/ioc"
)

type scheduler struct {
	container ioc.Container
}

// Scheduler is a job runner.
type Scheduler interface {
	Run() error
	RunAsync(ctx context.Context)

	Hook(hooks... Hook)
	Job(jobs... Job) 
}
