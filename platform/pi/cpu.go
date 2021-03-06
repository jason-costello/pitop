package pi

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PierreKieffer/pitop/cpu"
	"github.com/PierreKieffer/pitop/pkg/utils"
)

func (t *Temp) ComputeCPULoad() *cpu.CPULoad {

	// Extract stats
	prevExtract := t.GetCoresStats()
	time.Sleep(time.Second)
	extract := t.GetCoresStats()

	// Compute Usage
	var cpuLoad cpu.CPULoad

	var wg sync.WaitGroup
	wg.Add(5)

	// cpu
	go func() {
		defer wg.Done()
		cpuLoad.CPU = computeCoreLoad(extract.CPU, prevExtract.CPU)
	}()

	// cpu0
	go func() {
		defer wg.Done()
		cpuLoad.CPU0 = computeCoreLoad(extract.CPU0, prevExtract.CPU0)
	}()

	// cpu1
	go func() {
		defer wg.Done()
		cpuLoad.CPU1 = computeCoreLoad(extract.CPU1, prevExtract.CPU1)
	}()

	// cpu2
	go func() {
		defer wg.Done()
		cpuLoad.CPU2 = computeCoreLoad(extract.CPU2, prevExtract.CPU2)
	}()

	// cpu3
	go func() {
		defer wg.Done()
		cpuLoad.CPU3 = computeCoreLoad(extract.CPU3, prevExtract.CPU3)
	}()
	wg.Wait()

	return &cpuLoad

}

func (t *Temp) GetCoresStats() *cpu.CPUStat {
	/*
		Method to parse /proc/stat file and extract each stats for each core
	*/

	var cpuStat cpu.CPUStat

	procStatBytes, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(procStatBytes), "\n")

	for _, statLine := range dataSlice {
		statSlice := utils.FormatStatSlice(strings.Split(statLine, " "))

		if len(statSlice) > 0 {
			if statSlice[0] != "" && statSlice[0][:3] == "cpu" {
				var coreStat cpu.CoreStat
				coreStat.CoreId = statSlice[0]
				coreStat.User, _ = strconv.ParseUint(statSlice[1], 10, 64)
				coreStat.Nice, _ = strconv.ParseUint(statSlice[2], 10, 64)
				coreStat.System, _ = strconv.ParseUint(statSlice[3], 10, 64)
				coreStat.Idle, _ = strconv.ParseUint(statSlice[4], 10, 64)
				coreStat.IOWait, _ = strconv.ParseUint(statSlice[5], 10, 64)
				coreStat.IRQ, _ = strconv.ParseUint(statSlice[6], 10, 64)
				coreStat.SoftIRQ, _ = strconv.ParseUint(statSlice[7], 10, 64)
				coreStat.Steal, _ = strconv.ParseUint(statSlice[8], 10, 64)
				coreStat.Guest, _ = strconv.ParseUint(statSlice[9], 10, 64)
				coreStat.GuestNice, _ = strconv.ParseUint(statSlice[10], 10, 64)

				switch statSlice[0] {
				case "cpu":
					cpuStat.CPU = &coreStat
				case "cpu0":
					cpuStat.CPU0 = &coreStat
				case "cpu1":
					cpuStat.CPU1 = &coreStat
				case "cpu2":
					cpuStat.CPU2 = &coreStat
				case "cpu3":
					cpuStat.CPU3 = &coreStat
				}

			}
		}
	}
	return &cpuStat
}

func computeCoreLoad(currentStat, previousStat *cpu.CoreStat) float32 {

	/*
	   user    nice   system  idle      iowait irq   softirq  steal  guest  guest_nice
	   cpu  74608   2520   24433   1117073   6176   4054  0        0      0      0
	   *
	   *    Idle = idle + iowait
	   *    Load = user + nice + system + irq + softirq + steal
	   *    Total = Idle + Load
	   *
	   *    DiffTotal = Total_t - Total_t-1
	   *    DiffIdle = Idle_t - Idle_t-1
	   *    percentage = ( DiffTotal - DiffIdle ) / DiffTotal
	   *
	   *
	*/

	PreviousIdle := previousStat.Idle + previousStat.IOWait
	PreviousLoad := previousStat.User + previousStat.Nice + previousStat.System + previousStat.IRQ + previousStat.SoftIRQ + previousStat.Steal
	PreviousTotal := PreviousIdle + PreviousLoad

	Idle := currentStat.Idle + currentStat.IOWait
	Load := currentStat.User + currentStat.Nice + currentStat.System + currentStat.IRQ + currentStat.SoftIRQ + currentStat.Steal
	Total := Idle + Load

	DiffTotal := Total - PreviousTotal
	DiffIdle := Idle - PreviousIdle

	CPULoadPercentage := (float32(DiffTotal) - float32(DiffIdle)) / float32(DiffTotal) * 100

	return CPULoadPercentage
}

func (t *Temp) ExtractCPUFrequency() cpu.CPUFreq {
	/*
		cat /proc/cpuinfo | grep "MHz"
	*/

	var cpuFreq cpu.CPUFreq
	cpuInfoBytes, err := ioutil.ReadFile("/sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	freq, _ := strconv.ParseUint(strings.Split(string(cpuInfoBytes), "\n")[0], 10, 64)
	cpuFreq.Freq = freq / 1000
	return cpuFreq
}

func ExtractCPUInfoFrequency() {
	/*
		cat /proc/cpuinfo | grep "MHz"
	*/

	cpuInfoBytes, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(cpuInfoBytes), "\n")
	for _, infoLine := range dataSlice {
		info := utils.FormatStatSlice(strings.Split(infoLine, " "))

		if len(info) > 1 && info[0] == "cpu" && info[1][:3] == "MHz" {
			fmt.Println(info[2])
		}
	}

}

func ExtractLSCPU() {

	cmd := "lscpu | grep -i mhz"
	run := exec.Command("bash", "-c", cmd)
	stdout, err := run.Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outputLines := strings.Split(string(stdout), "\n")

	for _, l := range outputLines {
		fmt.Println(l)
	}

}
