package collector

import (
	"context"

	"github.com/PierreKieffer/pitop/cpu"
	"github.com/PierreKieffer/pitop/disk"
	"github.com/PierreKieffer/pitop/interfaces"
	"github.com/PierreKieffer/pitop/mem"
	"github.com/PierreKieffer/pitop/net"
)

type Collector struct {
	ctx  context.Context
	cpu  interfaces.CPUCollector
	disk interfaces.DiskCollector
	mem  interfaces.MemCollector
	net  interfaces.NetCollector
	temp interfaces.TempCollector
}
type Status struct {
	CPULoad cpu.CPULoad
	CPUFreq cpu.CPUFreq
	MemStat mem.MemStat
	Temp    float64
	Disk    []disk.DiskInfo
	Net     net.NetStat
}

func New(ctx context.Context, c interfaces.CPUCollector, d interfaces.DiskCollector, m interfaces.MemCollector, n interfaces.NetCollector, t interfaces.TempCollector) *Collector {
	return &Collector{
		ctx:  ctx,
		cpu:  c,
		disk: d,
		mem:  m,
		net:  n,
		temp: t,
	}
}

func (c *Collector) Collect(ctx context.Context) Status {
	return Status{}
}
