package interfaces

type Temp struct{}

type Temperature interface {
	ExtractTemp() *float64
}
