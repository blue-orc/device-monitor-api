package controllers

import (
	"device-monitor-api/monitor"
	"device-monitor-api/utilities"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//InitIftopController Initializes net endpoints
func InitIftopController(r *mux.Router) {
	r.HandleFunc("/iftop", iftopHandler).Methods("POST")
	r.HandleFunc("/iftop/clear", iftopClearHandler).Methods("GET")
}

func iftopHandler(w http.ResponseWriter, r *http.Request) {
	var im monitor.IftopMonitor
	err := utilities.ReadJsonHttpBody(r, &im)
	fmt.Println(im)
	if err != nil {
		msg := "Iftop handler: " + err.Error()
		fmt.Println(msg)
		utilities.RespondBadRequest(w, msg)
	}
	ip := r.RemoteAddr
	i := strings.Index(ip, ":")
	ip = ip[0:i]
	monitor.UpdateIftopMonitor(ip, im)
	utilities.RespondOK(w)
}

func iftopClearHandler(w http.ResponseWriter, r *http.Request) {
	monitor.ClearIftopMonitor()
	utilities.RespondOK(w)
}
