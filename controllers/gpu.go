package controllers

import (
	"device-monitor-api/monitor"
	"fmt"
	"github.com/NVIDIA/gpu-monitoring-tools/bindings/go/nvml"
	"device-monitor-api/utilities"
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
	var g []nvml.DeviceStatus
	err := utilities.ReadJsonHttpBody(r, &g)
	if err != nil {
		msg := "GPU handler: " + err.Error()
		fmt.Println(msg)
		utilities.RespondBadRequest(w, msg)
	}
	ip := r.RemoteAddr
	i := strings.Index(ip, ":")
	ip = ip[0:i]
	monitor.UpdateGPUMonitor(ip, g)
	utilities.RespondOK(w)
}

func gpuClearHandler(w http.ResponseWriter, r *http.Request) {
	monitor.ClearCPUMonitor()
	utilities.RespondOK(w)
}
