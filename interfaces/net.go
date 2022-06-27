package interfaces

import (
	"github.com/PierreKieffer/pitop/net"
)

//go:generate go run --mod=mod github.com/golang/mock/mockgen --source=./net.go --destination=../mocks/net.go --package=mocks

type NetCollector interface {
	ComputeNetStats() net.NetStat
}
