package picomm

import (
	"fmt"
	"time"
)

type (
	// Runner ...
	Runner struct {
		Jobs        []Job
		jobsChannel chan Job
		jobsQueue   map[int]Job
	}
)

// NewRunner ...
func NewRunner(jobsChannel chan Job, jobsQueue map[int]Job) *Runner {
	return &Runner{
		jobsChannel: jobsChannel,
		jobsQueue:   jobsQueue,
	}
}

// RunJobs ...
func (r *Runner) RunJobs() {
	for job := range r.jobsChannel {
		job.Status = "PROCESSING"
		r.jobsQueue[job.PIN] = job
		fmt.Printf("Start job: %+v\n", job)
		time.Sleep(time.Duration(job.Duration) * time.Second)
		job.Status = "DONE"
		r.jobsQueue[job.PIN] = job
		fmt.Printf("End job: %+v\n", job)
	}
}
