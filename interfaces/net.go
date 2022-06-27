package interfaces

type NetStat struct {
	BytesRecv      uint64
	BytesSent      uint64
	TotalBytesRecv uint64
	TotalBytesSent uint64
}

type NetInfo interface {
	ComputeNetStats() *NetStat
}
