package monitor

import (
	"device-monitor-api/ipcheck"
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/net"
)

type NetMonitorData struct {
	BytesSent       uint64
	bytesSentInit   uint64
	PacketsSent     uint64
	packetsSentInit uint64
	BytesRecv       uint64
	bytesRecvInit   uint64
	PacketsRecv     uint64
	packetsRecvInit uint64
	Country         string
}

var nm map[string]NetMonitorData

func NetMonitorInit() {
	nm = map[string]NetMonitorData{}
}

func UpdateNetMonitor(ip string, data net.IOCountersStat) {
	if val, ok := nm[ip]; !ok {
		var n NetMonitorData
		n.bytesRecvInit = data.BytesRecv
		n.packetsRecvInit = data.PacketsRecv
		n.bytesSentInit = data.BytesSent
		n.packetsSentInit = data.PacketsSent
		n.Country = ipcheck.GetIPCountry(ip)
		nm[ip] = n
	} else {
		val.BytesRecv = data.BytesRecv - val.bytesRecvInit
		val.PacketsRecv = data.PacketsRecv - val.packetsRecvInit
		val.BytesSent = data.BytesSent - val.bytesSentInit
		val.packetsSentInit = data.PacketsSent - val.packetsSentInit
		nm[ip] = val
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
