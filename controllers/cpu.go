package controllers

import (
	"device-monitor-api/monitor"
	"fmt"
	"gpu-demonstration-api/utilities"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//InitCPUController Initializes cpu endpoints
func InitCPUController(r *mux.Router) {
	r.HandleFunc("/cpu", cpuHandler).Methods("POST")
	r.HandleFunc("/cpu/clear", cpuClearHandler).Methods("GET")
}

func cpuHandler(w http.ResponseWriter, r *http.Request) {
	var c monitor.CPUStatus
	err := utilities.ReadJsonHttpBody(r, &c)
	if err != nil {
		msg := "CPU handler: " + err.Error()
		fmt.Println(msg)
		utilities.RespondBadRequest(w, msg)
	}
	ip := r.RemoteAddr
	i := strings.Index(ip, ":")
	ip = ip[0:i]
	monitor.UpdateCPUMonitor(ip, c)
	utilities.RespondOK(w)
}

func cpuClearHandler(w http.ResponseWriter, r *http.Request) {
	monitor.ClearCPUMonitor()
	utilities.RespondOK(w)
}
