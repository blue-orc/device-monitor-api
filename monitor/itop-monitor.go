package monitor

import (
	"fmt"
	"github.com/orcaman/concurrent-map"
)

type IftopMonitor struct {
	BytesReceivedRate float64
	TotalReceived     float64
}

var im cmap.ConcurrentMap

func IftopMonitorInit() {
	im = cmap.New()
}

func UpdateIftopMonitor(ip string, data IftopMonitor) {
	im.Set(ip, data)
}

func ClearIftopMonitor() {
	im = cmap.New()
}

func GetIftopMonitorJSON() (bytes []byte) {
	bytes, err := im.MarshalJSON()
	if err != nil {
		fmt.Println("GetIftopMonitor: " + err.Error())
		return
	}
	return
}
