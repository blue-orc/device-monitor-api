package controllers

import (
	"device-monitor-api/monitor"
	"device-monitor-api/utilities"
	"fmt"
	"github.com/NVIDIA/gpu-monitoring-tools/bindings/go/nvml"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//InitGPUController Initializes gpu endpoints
func InitGPUController(r *mux.Router) {
	r.HandleFunc("/gpu", gpuHandler).Methods("POST")
	r.HandleFunc("/gpu/clear", gpuClearHandler).Methods("GET")
}

func gpuHandler(w http.ResponseWriter, r *http.Request) {
	var g nvml.DeviceStatus
	var ga []nvml.DeviceStatus
	ga = append(ga, g)
	err := utilities.ReadJsonHttpBody(r, &g)
	if err != nil {
		msg := "GPU handler: " + err.Error()
		fmt.Println(msg)
		utilities.RespondBadRequest(w, msg)
	}
	fmt.Println(g)
	ip := r.RemoteAddr
	i := strings.Index(ip, ":")
	ip = ip[0:i]
	monitor.UpdateGPUMonitor(ip, ga)
	utilities.RespondOK(w)
}

func gpuClearHandler(w http.ResponseWriter, r *http.Request) {
	monitor.ClearCPUMonitor()
	utilities.RespondOK(w)
}
