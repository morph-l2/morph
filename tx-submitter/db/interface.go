package db

// Database defines the interface for database operations
type Database interface {
	GetString(key string) (string, error)
	PutString(key, val string) error
	GetFloat(key string) (float64, error)
	PutFloat(key string, val float64) error
	Close() error
}
