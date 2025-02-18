package ptr

func Addr[T any](t T) *T { return &t }

func ReturnNonNil[T any](input *T) T {
	if input == nil {
		var defaultValue T
		return defaultValue
	}
	return *input
}
