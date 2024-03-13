// Package err предоставляет набор функций для обработки ошибок.
package err

// Must делает любую функцию Mustable (либо вернет необходимое значение, либо запаникует)
func Must[T any](item T, err error) T {
	if err != nil {
		panic(err)
	}

	return item
}
