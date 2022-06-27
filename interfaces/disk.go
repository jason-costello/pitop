package interfaces

type Disk struct{}

type DiskCollector interface {
	ExtractDiskUsage() *[]DiskInfo
}

type DiskInfo struct {
	MountingPoint string
	Size          string
	Used          string
	Avail         string
	PercentUse    string
}
