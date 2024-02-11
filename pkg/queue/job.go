package queue

import "time"

type Job struct {
	ID        string
	Payload   string
	ExecuteAt time.Time
}

func EnqueueDelayedJob(jobs chan<- Job, job Job) {
	timeToExecute := time.Until(job.ExecuteAt)
	timer := time.NewTimer(timeToExecute)

	go func() {
		<-timer.C
		jobs <- job
	}()
}
