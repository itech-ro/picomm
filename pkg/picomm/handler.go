package picomm

import (
	"encoding/json"
	"io"
	"net/http"
)

type (
	// JobsRequest ...
	JobsRequest struct {
		Jobs []Job `json:"jobs"`
	}
)

// HandleJobs ...
func HandleJobs(
	w http.ResponseWriter,
	r *http.Request,
	cfg Config,
	controller *Controller,
) {
	var jobs JobsRequest

	if controller.IsJobRunning() {
		http.Error(w, "there are running jobs", http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "cannot read request body", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &jobs)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	returnJobs, err := controller.ProcessJobs(jobs.Jobs)
	if err != nil {
		http.Error(w, "cannot start the jobs", http.StatusInternalServerError)
		return
	}

	jsonJobs, err := json.Marshal(returnJobs)
	if err != nil {
		http.Error(w, "cannot output the jobs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonJobs)
}

// HandleStatus ...
func HandleStatus(
	w http.ResponseWriter,
	r *http.Request,
	cfg Config,
	controller *Controller,
) {
	jobs := controller.GetJobs()
	json.NewEncoder(w).Encode(jobs)
}
