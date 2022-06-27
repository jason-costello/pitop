package interfaces

import (
	"github.com/PierreKieffer/pitop/disk"
)

//go:generate go run --mod=mod github.com/golang/mock/mockgen --source=./disk.go --destination=../mocks/disk.go --package=mocks

type DiskCollector interface {
	ExtractDiskUsage() []disk.DiskInfo
}
