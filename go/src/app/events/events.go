package events

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/process"
)

// ProcessEvent is the event dispatched to notify
// something happened to the monitored process
type ProcessEvent struct {
	Name    string
	Data    interface{}
	Message string
}

const MEMORY = "MEM"
const CPU = "CPU"
const STARTED = "STARTED"
const DONE = "DONE"
const FAILED = "FAILED"
const LOG = "LOG"

func CPUEEvent(percent float64) ProcessEvent {
	cpuAsString := fmt.Sprint(percent)
	return ProcessEvent{Name: CPU, Data: cpuAsString, Message: "CPU %"}
}

func MemoryEvent(memoryStats *process.MemoryInfoStat) ProcessEvent {
	memAsString, _ := json.Marshal(memoryStats)
	return ProcessEvent{Name: MEMORY, Data: string(memAsString), Message: "Memory stats"}
}

func StartedProcessEvent(pid int) ProcessEvent {
	return ProcessEvent{Name: STARTED, Data: pid}
}

func SuccessEvent(output []byte) ProcessEvent {
	return ProcessEvent{Name: DONE, Data: "Success", Message: string(output)}
}

func FailedEvent(err error) ProcessEvent {
	return ProcessEvent{Name: FAILED, Data: "Failed", Message: err.Error()}
}

func LogEvent(message string) ProcessEvent {
	return ProcessEvent{Name: LOG, Data: "", Message: message}
}

func LogEventf(message string, args interface{}) ProcessEvent {
	msg := fmt.Sprintf(message, args)
	return ProcessEvent{Name: LOG, Data: "", Message: msg}
}

func LogEventData(message string, extraData interface{}) ProcessEvent {
	data, _ := json.Marshal(extraData)
	return ProcessEvent{Name: LOG, Data: data, Message: message}
}
