package interfaces

//go:generate go run --mod=mod github.com/golang/mock/mockgen --source=./temp.go --destination=../mocks/temp.go --package=mocks

type TempCollector interface {
	ExtractTemp() float64
}
