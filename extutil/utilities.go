package extutil

import (
	"io/fs"
	"sort"
)

// returns true if x is found in a
func ContainsStrings(a []string, x string) bool {
	var i = sort.SearchStrings(a, x)
	return i < len(a) && a[i] == x
}

// returns true if x is found in a
func ContainsInts(a []int, x int) bool {
	var i = sort.SearchInts(a, x)
	return i < len(a) && a[i] == x
}

// returns true if x is found in a
func ContainsFloat64s(a []float64, x float64) bool {
	var i = sort.SearchFloat64s(a, x)
	return i < len(a) && a[i] == x
}

// panics if fs.sub fails
func MustSubFS(fsys fs.FS, dir string) fs.FS {
	fsub, err := fs.Sub(fsys, dir)
	if err != nil {
		panic(err)
	}

	return fsub
}
