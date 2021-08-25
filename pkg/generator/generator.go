package generator

// Generator is the interface that all data generators must satisfy
type Generator interface {
	Next() interface{}
}
