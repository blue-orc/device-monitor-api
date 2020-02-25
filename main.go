package main

import (
	"device-monitor-api/constants"
	"device-monitor-api/controllers"
	"device-monitor-api/monitor"
	"device-monitor-api/text-reader"
	"device-monitor-api/websocket"
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var addr = flag.String("addr", ":9002", "websocket service address")

func main() {
	monitor.DiskMonitorInit()
	monitor.NetMonitorInit()
	monitor.CPUMonitorInit()
	monitor.GPUMonitorInit()
	monitor.TrainingMonitorInit()

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("main: %s", err.Error())
	}

	constants.AmericanCIDR = reader.Read(wd + "/america.txt")
	constants.JapanCIDR = reader.Read(wd + "/japan.txt")
	fmt.Println(constants.AmericanCIDR[0])
	fmt.Println(len(constants.AmericanCIDR))
	mux := mux.NewRouter()
	initializeControllers(mux)
	go func() {
		fmt.Println("API starting")
		err := http.ListenAndServe(":9001",
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}))(mux))
		if err != nil {
			log.Fatal("API ListenAndServe: ", err)
		}
	}()
	flag.Parse()
	hub := websocket.NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, w, r)
	})

	fmt.Println("Starting websocket server")
	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("WS ListenAndServe: ", err)
	}

	fmt.Println("closing")
}

func initializeControllers(r *mux.Router) {
	controllers.InitCPUController(r)
	controllers.InitDiskController(r)
	controllers.InitGPUController(r)
	controllers.InitStatusController(r)
	controllers.InitTrainingController(r)
	controllers.InitNetController(r)
}
