package platform

import (
	"context"

	"github.com/PierreKieffer/pitop/cpu"
	"github.com/PierreKieffer/pitop/disk"
	"github.com/PierreKieffer/pitop/mem"
	"github.com/PierreKieffer/pitop/net"
	"github.com/PierreKieffer/pitop/pkg/collector"
)

type platform struct {
	ctx       context.Context
	collector collector.Collector
}

type Status struct {
	CPULoad cpu.CPULoad
	CPUFreq cpu.CPUFreq
	MemStat mem.MemStat
	Temp    float64
	Disk    []disk.DiskInfo
	Net     net.NetStat
}

func New(ctx context.Context, c *collector.Collector) *platform {
	return &platform{
		ctx:       ctx,
		collector: c,
	}
}

func (p *platform) Collect(ctx context.Context) Status {
	cStatus := p.collector.Collect(ctx)
	status := Status{
		CPULoad: cStatus.CPULoad,
		CPUFreq: cStatus.CPUFreq,
		MemStat: cStatus.MemStat,
		Temp:    cStatus.Temp,
		Disk:    cStatus.Disk,
		Net:     cStatus.Net,
	}
	return status
}
