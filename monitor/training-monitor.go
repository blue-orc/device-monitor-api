package monitor

import (
	"encoding/json"
	"fmt"
)

type TrainingStatus struct {
	CurrentFile       string
	Step              string
	TrainingScript    string
	BatchSize         string
	NumberOfWorkers   string
	NumberOfFiles     string
	CurrentFileIndex  string
	CurrentImageIndex string
	ImagesPerFile     string
	Loss              string
	Status            string
	CurrentEpoch      string
	Epochs            string
	Layers            string
	Depth             string
	LearningRate      string
}

var tm map[string]TrainingStatus

func TrainingMonitorInit() {
	tm = map[string]TrainingStatus{}
}

func UpdateTrainingMonitor(ip string, data TrainingStatus) {
	tm[ip] = data
}

func GetTrainingMonitor() map[string]TrainingStatus {
	return tm
}

func ClearTrainingMonitor() {
	tm = map[string]TrainingStatus{}
}

func GetTrainingMonitorJSON() (bytes []byte) {
	bytes, err := json.Marshal(GetTrainingMonitor())
	if err != nil {
		fmt.Println("GetTrainingMonitorJSON: " + err.Error())
		return
	}
	return
}
