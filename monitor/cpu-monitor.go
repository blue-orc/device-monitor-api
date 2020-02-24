package monitor

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

type CPUStatus struct {
	Memory     *mem.VirtualMemoryStat
	CPUPercent []float64
}

var cm map[string]CPUStatus

func CPUMonitorInit() {
	cm = map[string]CPUStatus{}
}

func UpdateCPUMonitor(ip string, data CPUStatus) {
	cm[ip] = data
}

func GetCPUMonitor() map[string]CPUStatus {
	return cm
}

func ClearCPUMonitor() {
	cm = map[string]CPUStatus{}
}

func GetCPUMonitorJSON() (bytes []byte) {
	bytes, err := json.Marshal(GetCPUMonitor())
	if err != nil {
		fmt.Println("GetCPUMonitorJSON: " + err.Error())
		return
	}
	return
}
