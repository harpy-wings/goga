package goga

type DB interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte) error
}
