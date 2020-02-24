package monitor

import (
	"encoding/json"
	"fmt"
	"github.com/NVIDIA/gpu-monitoring-tools/bindings/go/nvml"
)

var gm map[string][]nvml.DeviceStatus

func GPUMonitorInit() {
	gm = map[string][]nvml.DeviceStatus{}
}

func UpdateGPUMonitor(ip string, data []nvml.DeviceStatus) {
	gm[ip] = data
}

func GetGPUMonitor() map[string][]nvml.DeviceStatus {
	return gm
}

func ClearGPUMonitor() {
	gm = map[string][]nvml.DeviceStatus{}
}

func GetGPUMonitorJSON() (bytes []byte) {
	bytes, err := json.Marshal(GetGPUMonitor())
	if err != nil {
		fmt.Println("GetGPUMonitorJSON: " + err.Error())
		return
	}
	return
}
