package controllers

import (
	"device-monitor-api/monitor"
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"gpu-demonstration-api/utilities"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//InitDiskController Initializes disk endpoints
func InitDiskController(r *mux.Router) {
	r.HandleFunc("/disk", diskHandler).Methods("POST")
	r.HandleFunc("/disk/clear", diskClearHandler).Methods("GET")
}

func diskHandler(w http.ResponseWriter, r *http.Request) {
	var d map[string]disk.IOCountersStat
	err := utilities.ReadJsonHttpBody(r, &d)
	if err != nil {
		msg := "Disk handler: " + err.Error()
		fmt.Println(msg)
		utilities.RespondBadRequest(w, msg)
	}
	ip := r.RemoteAddr
	i := strings.Index(ip, ":")
	ip = ip[0:i]
	monitor.UpdateDiskMonitor(ip, d["sda"])
	utilities.RespondOK(w)
}

func diskClearHandler(w http.ResponseWriter, r *http.Request) {
	monitor.ClearDiskMonitor()
	utilities.RespondOK(w)
}
