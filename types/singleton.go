package types

func Singleton[T any]() func() *T {
	var instance T
	return SingletonFrom[T](&instance)
}

func SingletonFrom[T any](inst *T) func() *T {
	return func() *T {
		return inst
	}
}
