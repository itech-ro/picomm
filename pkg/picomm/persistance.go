package picomm

import (
	"encoding/json"
	"os"
)

type (
	// Persistance ...
	Persistance struct {
		filePath string
	}
)

// NewPersistance ...
func NewPersistance(filePath string) *Persistance {
	return &Persistance{
		filePath: filePath,
	}
}

// Read ...
func (p *Persistance) Read() ([]Job, error) {
	var jobs []Job

	data, err := os.ReadFile(p.filePath)
	if err != nil {
		return jobs, err
	}

	err = json.Unmarshal(data, &jobs)
	return jobs, err
}

// Store ...
func (p *Persistance) Store(jobs []Job) error {
	jobsJSON, err := json.Marshal(jobs)
	if err != nil {
		return err
	}

	return os.WriteFile(p.filePath, jobsJSON, 0644)
}

// Clear ...
func (p *Persistance) Clear() error {
	return os.Remove(p.filePath)
}
