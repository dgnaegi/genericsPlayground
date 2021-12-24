package main

type Result[T any] struct {
	value T
	err   error
}

func ExecuteOnSuccess[T any, T2 any](result Result[T], function func(value T) Result[T2]) Result[T2] {
	if result.err != nil {
		return Result[T2]{
			err: result.err,
		}
	}
	return function(result.value)
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		value: value,
		err:   nil,
	}
}

func Err[T any](value T, err error) Result[T] {
	return Result[T]{
		value: value,
		err:   err,
	}
}
