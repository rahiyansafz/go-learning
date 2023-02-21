package main

func PrintSlice[T any](s []T) {
	for _, v := range s {
		print(v)
	}
}
