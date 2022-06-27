package darwin

type Temperature struct{}

type Temp struct{}

func (t *Temp) ExtractTemp() float64 {
	return 0.0
}
