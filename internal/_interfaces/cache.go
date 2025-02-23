package _interfaces

type Cache[T any] interface {
	Get(key string) (*T, bool)
	Save(key string, value T)
}
