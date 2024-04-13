package utils

import (
	"fmt"
	"slices"
)

func SliceFunc() {
	var slice []string = []string{"func", "456", "func", "123", "456", "789"}
	fmt.Printf("%s", slices.Compact(slice))
	slices.Sort(slice)
	fmt.Printf("%s", slices.Compact(slice))
	fmt.Printf("%s", slices.Max(slice))
}
