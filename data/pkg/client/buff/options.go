package buff

type FilterFunc[T any] func(T) (bool, error)
