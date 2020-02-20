package monitor

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/net"
)

type NetMonitorData struct {
	BytesRecv       uint64
	bytesRecvInit   uint64
	PacketsRecv     uint64
	packetsRecvInit uint64
}

var nm map[string]NetMonitorData

func UpdateNetMonitor(ip string, data net.IOCountersStat) {
	if val, ok := nm[ip]; !ok {
		var n NetMonitorData
		n.bytesRecvInit = data.BytesRecv
		n.packetsRecvInit = data.PacketsRecv
		nm[ip] = n
	} else {
		val.BytesRecv = data.BytesRecv - val.bytesRecvInit
		val.PacketsRecv = data.PacketsRecv - val.packetsRecvInit
	}
}

func ClearNetMonitor() {
	nm = map[string]NetMonitorData{}
}

func GetNetMonitor() map[string]NetMonitorData {
	return nm
}

func GetNetMonitorJSON() (bytes []byte) {
	bytes, err := json.Marshal(GetNetMonitor())
	if err != nil {
		fmt.Println("GetNetMonitorJSON: " + err.Error())
		return
	}
	return
}
