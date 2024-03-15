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
		persistance  *Persistance
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
func NewController(jobsChannel chan Job, persistance *Persistance, jobsQueue map[int]Job) *Controller {
	return &Controller{
		jobsChannel: jobsChannel,
		jobsQueue:   jobsQueue,
		persistance: persistance,
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

	c.persistance.Clear()
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

	c.persistance.Store(returnJobs)

	return returnJobs, nil
}

// Init ...
func (c *Controller) Init() error {
	var jobs []Job
	persistanceJobs, err := c.persistance.Read()

	if err != nil {
		return err
	}

	for _, job := range persistanceJobs {
		if job.EndTime.Before(time.Now()) {
			continue
		}

		duration := 0
		if job.StartTime.After(time.Now()) {
			duration = job.Duration
		} else {
			diff := job.EndTime.Sub(time.Now())
			duration = int(diff.Seconds())
		}

		job.Duration = duration
		jobs = append(jobs, job)
	}

	if len(jobs) == 0 {
		return nil
	}

	_, err = c.ProcessJobs(jobs)
	return err
}
