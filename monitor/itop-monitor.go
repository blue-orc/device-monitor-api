package monitor

import (
	"encoding/json"
	"fmt"
)

type IftopMonitor struct {
	BytesReceivedRate float64
	TotalReceived     float64
}

var im map[string]IftopMonitor

func IftopMonitorInit() {
	im = map[string]IftopMonitor{}
}

func UpdateIftopMonitor(ip string, data IftopMonitor) {
	im[ip] = data
}

func ClearIftopMonitor() {
	im = map[string]IftopMonitor{}
}

func GetIftopMonitor() map[string]IftopMonitor {
	return im
}

func GetIftopMonitorJSON() (bytes []byte) {
	bytes, err := json.Marshal(GetIftopMonitor())
	if err != nil {
		fmt.Println("GetIftopMonitor: " + err.Error())
		return
	}
	return
}
