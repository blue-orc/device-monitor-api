package controllers

import (
	"device-monitor-api/monitor"
	"device-monitor-api/utilities"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//InitTrainingController Initializes gpu endpoints
func InitTrainingController(r *mux.Router) {
	r.HandleFunc("/training", trainingHandler).Methods("POST")
	r.HandleFunc("/training/clear", trainingClearHandler).Methods("GET")
}

func trainingHandler(w http.ResponseWriter, r *http.Request) {
	var t monitor.TrainingStatus
	err := utilities.ReadJsonHttpBody(r, &t)
	if err != nil {
		msg := "Training handler: " + err.Error()
		fmt.Println(msg)
		utilities.RespondBadRequest(w, msg)
	}
	ip := r.RemoteAddr
	i := strings.Index(ip, ":")
	ip = ip[0:i]
	monitor.UpdateTrainingMonitor(ip, t)
	utilities.RespondOK(w)
}

func trainingClearHandler(w http.ResponseWriter, r *http.Request) {
	monitor.ClearTrainingMonitor()
	utilities.RespondOK(w)
}
