package interfaces

import (
	"github.com/PierreKieffer/pitop/mem"
)

//go:generate go run --mod=mod github.com/golang/mock/mockgen --source=./mem.go --destination=../mocks/mem.go --package=mocks

type MemCollector interface {
	ExtractMemStats() mem.MemStat
}
