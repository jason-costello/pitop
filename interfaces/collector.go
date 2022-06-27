package interfaces

import (
	"context"
)

type Collector interface {
	Collect(ctx context.Context) interface{}
}
