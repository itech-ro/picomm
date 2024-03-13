package picomm

import (
	"fmt"
	"sync"
	"time"
)

type (
	// Controller ...
	Controller struct {
		jobsChannel  chan Job
		jobsQueue    map[int]Job
		isJobRunning bool
	}

	// Job ...
	Job struct {
		Name      string    `json:"name"`
		PIN       int       `json:"pin"`
		Status    string    `json:"status"`
		Duration  int       `json:"duration"`
		StartTime time.Time `json:"startTime,omitempty"`
		EndTime   time.Time `json:"endTime,omitempty"`
	}
)

// NewController ...
func NewController(jobsChannel chan Job, jobsQueue map[int]Job) *Controller {
	return &Controller{
		jobsChannel: jobsChannel,
		jobsQueue:   jobsQueue,
	}
}

// IsJobRunning ...
func (c *Controller) IsJobRunning() bool {
	return c.isJobRunning
}

// StartJobs ...
func (c *Controller) StartJobs() {
	var m sync.Mutex

	m.Lock()
	defer m.Unlock()
	c.isJobRunning = true
}

// EndJobs ...
func (c *Controller) EndJobs() {
	var m sync.Mutex

	m.Lock()
	defer m.Unlock()
	c.isJobRunning = false
}

// GetJobs ...
func (c *Controller) GetJobs() []Job {
	jobs := make([]Job, 0)

	if len(c.jobsQueue) == 0 {
		return jobs
	}

	for _, job := range c.jobsQueue {
		jobs = append(jobs, job)
	}

	return jobs
}

// ProcessJobs ...
func (c *Controller) ProcessJobs(jobs []Job) ([]Job, error) {
	c.StartJobs()

	returnJobs := make([]Job, 0)
	delay := 0
	for _, job := range jobs {
		job.StartTime = time.Now().Add(time.Duration(delay) * time.Second)
		delay = delay + job.Duration
		job.EndTime = job.StartTime.Add(time.Duration(job.Duration) * time.Second)
		c.jobsQueue[job.PIN] = job
		fmt.Printf("Send job: %+v\n", job)
		returnJobs = append(returnJobs, job)
		c.jobsChannel <- job
	}
	return returnJobs, nil
}
