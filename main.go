package main

import (
	"device-monitor-api/controllers"
	"device-monitor-api/websocket"
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":9002", "websocket service address")

func main() {
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
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("WS ListenAndServe: ", err)
	}

	fmt.Println("closing")
}

func initializeControllers(r *mux.Router) {
	controllers.InitDiskController(r)
	controllers.InitStatusController(r)
	controllers.InitNetController(r)
}
