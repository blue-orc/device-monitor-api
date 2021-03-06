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

var nm map[string][]NetMonitorData

func NetMonitorInit() {
	nm = map[string][]NetMonitorData{}
}

func UpdateNetMonitor(ip string, data []net.IOCountersStat) {
	if val, ok := nm[ip]; !ok || len(nm[ip]) != len(data) {
		var ns []NetMonitorData
		for _, d := range data {
			var n NetMonitorData
			n.bytesRecvInit = d.BytesRecv
			n.packetsRecvInit = d.PacketsRecv
			n.bytesSentInit = d.BytesSent
			n.packetsSentInit = d.PacketsSent
			n.Country = ipcheck.GetIPCountry(ip)
			ns = append(ns, n)
		}
		nm[ip] = ns
	} else {
		for i, d := range data {
			val[i].BytesRecv = d.BytesRecv - val[i].bytesRecvInit
			val[i].PacketsRecv = d.PacketsRecv - val[i].packetsRecvInit
			val[i].BytesSent = d.BytesSent - val[i].bytesSentInit
			val[i].packetsSentInit = d.PacketsSent - val[i].packetsSentInit
		}
		nm[ip] = val
	}
}

func ClearNetMonitor() {
	nm = map[string][]NetMonitorData{}
}

func GetNetMonitor() map[string][]NetMonitorData {
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
