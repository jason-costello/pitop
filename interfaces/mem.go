package interfaces

type MemStat struct {
	MemTotal  uint64
	MemFree   uint64
	MemUsage  float32
	SwapTotal uint64
	SwapFree  uint64
	SwapUsage float32
	Buffers   uint64
	Cached    uint64
}

type MemUsage struct {
}

type MemInfo interface {
	GetMemStats() *MemStat
}
