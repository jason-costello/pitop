package interfaces

type CPUCollector interface {
	ComputeCPULoad() *CPULoad
	GetCoresStats() *CPUStat
}

type CoreStat struct {
	CoreId    string
	User      uint64
	Nice      uint64
	System    uint64
	Idle      uint64
	IOWait    uint64
	IRQ       uint64
	SoftIRQ   uint64
	Steal     uint64
	Guest     uint64
	GuestNice uint64
}

type CPUStat struct {
	CPU  *CoreStat // total
	CPU0 *CoreStat
	CPU1 *CoreStat
	CPU2 *CoreStat
	CPU3 *CoreStat
}

type CPULoad struct {
	CPU  float32 // total %
	CPU0 float32
	CPU1 float32
	CPU2 float32
	CPU3 float32
}

type CPUFreq struct {
	Freq uint64
}
