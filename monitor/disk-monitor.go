package monitor

import (
	"device-monitor-api/ipcheck"
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/disk"
)

type DiskMonitorData struct {
	ReadCount      uint64
	readCountInit  uint64
	WriteCount     uint64
	writeCountInit uint64
	ReadBytes      uint64
	readBytesInit  uint64
	WriteBytes     uint64
	writeBytesInit uint64
	Country        string
}

var dm map[string]DiskMonitorData

func DiskMonitorInit() {
	dm = map[string]DiskMonitorData{}
}

func UpdateDiskMonitor(ip string, data disk.IOCountersStat) {
	if val, ok := dm[ip]; !ok {
		var d DiskMonitorData
		d.readBytesInit = data.ReadBytes
		d.readCountInit = data.ReadCount
		d.writeBytesInit = data.WriteBytes
		d.writeCountInit = data.WriteCount
		d.Country = ipcheck.GetIPCountry(ip)
		dm[ip] = d
	} else {
		val.ReadBytes = data.ReadBytes - val.readCountInit
		val.ReadCount = data.ReadCount - val.readCountInit
		val.WriteBytes = data.WriteBytes - val.writeCountInit
		val.WriteCount = data.WriteCount - val.writeCountInit
		dm[ip] = val
	}
}

func ClearDiskMonitor() {
	dm = map[string]DiskMonitorData{}
}

func GetDiskMonitor() map[string]DiskMonitorData {
	return dm
}

func GetDiskMonitorJSON() (bytes []byte) {
	bytes, err := json.Marshal(GetDiskMonitor())
	if err != nil {
		fmt.Println("GetDiskMonitorJSON: " + err.Error())
		return
	}
	return
}
