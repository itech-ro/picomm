package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/itech-ro/picomm/pkg/picomm"
)

func main() {
	cfg, err := picomm.GetConfig("config", ".")
	if err != nil {
		log.Panicln(err)
	}

	jobsChannel := make(chan picomm.Job, 10)

	jobsQueue := make(map[int]picomm.Job)

	persistance := picomm.NewPersistance("jobs.dat")

	controller := picomm.NewController(jobsChannel, persistance, jobsQueue)
	controller.Init()

	runner := picomm.NewRunner(jobsChannel, jobsQueue)

	go runner.RunJobs()

	router := mux.NewRouter()
	router.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		picomm.HandleJobs(w, r, cfg, controller)
	}).Methods("POST")

	router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		picomm.HandleStatus(w, r, cfg, controller)
	}).Methods("GET")

	http.Handle("/", router)

	log.Println(" - http server:", cfg.GetString("http.address"))
	log.Fatal(http.ListenAndServe(cfg.GetString("http.address"), nil))
}
