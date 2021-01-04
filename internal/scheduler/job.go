package scheduler

// base job
type job struct {
	key      string
	caller interface{}
}

// Job is what you do for living.
type Job interface {
	check() error

	State() State
	Next()
}
