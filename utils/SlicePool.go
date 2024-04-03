package utils

import "sync"

type SlicePool struct {
	NewSlice func(size int) any // function to create a
	GetSize  func(any) int      // function to get the size of a slice

	pool    []sync.Pool // pool array
	sizeMap []int       // map size
}
