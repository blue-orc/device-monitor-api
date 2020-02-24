package controllers

import (
	"device-monitor-api/monitor"
	"fmt"
	"github.com/shirou/gopsutil/net"
	"device-monitor-api/utilities"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//InitNetController Initializes net endpoints
func InitNetController(r *mux.Router) {
	r.HandleFunc("/net", netHandler).Methods("POST")
	r.HandleFunc("/net/clear", netClearHandler).Methods("GET")
}

func netHandler(w http.ResponseWriter, r *http.Request) {
	var n []net.IOCountersStat
	err := utilities.ReadJsonHttpBody(r, &n)
	if err != nil {
		msg := "Net handler: " + err.Error()
		fmt.Println(msg)
		utilities.RespondBadRequest(w, msg)
	}
	if len(n) == 0 {
		msg := "Net handler: " + "recieved empty set of data"
		fmt.Println(msg)
		utilities.RespondBadRequest(w, msg)
	}
	ip := r.RemoteAddr
	i := strings.Index(ip, ":")
	ip = ip[0:i]
	monitor.UpdateNetMonitor(ip, n[0])
	utilities.RespondOK(w)
}

func netClearHandler(w http.ResponseWriter, r *http.Request) {
	monitor.ClearNetMonitor()
	utilities.RespondOK(w)
}
