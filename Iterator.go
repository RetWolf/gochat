package gochat

// Iterator defines a generic iterator interface which structs can then implement
type Iterator interface {
	Next() bool
	Get() interface{}
}
