package worker

import (
	"sync"

	"github.com/PierreKieffer/pitop/interfaces"
	"github.com/PierreKieffer/pitop/platform/pi"
)

type Status struct {
	CPULoad *interfaces.CPULoad
	CPUFreq *interfaces.CPUFreq
	Mem     *interfaces.MemStat
	Temp    *float64
	Disk    *[]interfaces.DiskInfo
	Net     *interfaces.NetStat
}

func Worker() *Status {

	var status Status

	var tempr pi.Temp

	var wg sync.WaitGroup
	wg.Add(6)

	// CPU Load
	go func() {
		defer wg.Done()
		status.CPULoad = tempr.ComputeCPULoad()
	}()

	// CPU Frequency
	go func() {
		defer wg.Done()
		status.CPUFreq = tempr.ExtractCPUFrequency()
	}()

	// Mem Stats
	go func() {
		defer wg.Done()
		status.Mem = tempr.GetMemStats()
	}()

	// Temp
	go func() {
		defer wg.Done()
		status.Temp = tempr.ExtractTemp()
		// status.Temp = temp.ExtractTemp()
	}()

	// Disk
	go func() {
		defer wg.Done()
		status.Disk = tempr.ExtractDiskUsage()
	}()

	// Net
	go func() {
		defer wg.Done()
		status.Net = tempr.ComputeNetStats()
	}()
	wg.Wait()

	return &status
}
