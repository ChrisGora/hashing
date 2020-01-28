package hashing

type HashMap interface {
	Put(key string, value interface{}) error
	Get(key string, value interface{}) error
	ToString() string
}
