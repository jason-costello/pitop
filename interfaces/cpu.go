package interfaces

import (
	"github.com/PierreKieffer/pitop/cpu"
)

//go:generate go run --mod=mod github.com/golang/mock/mockgen --source=./cpu.go --destination=../mocks/cpu.go --package=mocks

type CPUCollector interface {
	ComputeCPULoad() cpu.CPULoad
	GetCoresStats() cpu.CPUStat
	ExtractCPUFrequency() cpu.CPUFreq
}
