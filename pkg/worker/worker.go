package worker

import (
	"github.com/PierreKieffer/pitop/platform"
	"github.com/PierreKieffer/pitop/platform/pi"
)

type Status struct {
	CPULoad float64
	CPUFreq float64
	Mem     float64
	Temp    float64
	Disk    []float64
	Net     float64
}

func Worker(platform *platform.Status) platform.Status {

	var status Status

	var tempr pi.Temp

	return &status
}
