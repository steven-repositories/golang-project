package models

type Result[T any] struct {
	Data  *T
	Error error
}

func (result Result[T]) Successful() bool {
	return result.Error == nil
}

func Success[T any](data ...T) Result[T] {
	if len(data) > 0 {
		return Result[T]{
			Data: &data[0],
		}
	}

	return Result[T]{}
}

func Error[T any](err error) Result[T] {
	return Result[T]{
		Error: err,
	}
}
