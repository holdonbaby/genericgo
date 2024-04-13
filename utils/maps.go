package utils

import (
	"fmt"
	"maps"
)

func Test() {
	myMap := make(map[string]bool)
	youMap := maps.Clone(myMap)
	maps.DeleteFunc(youMap, func(s string, b bool) bool { return !youMap[s] })
	if maps.Equal(myMap, youMap) {
		fmt.Printf("%s", "equal mymap and you map")
	}
	maps.EqualFunc(myMap, youMap, func(b1, b2 bool) bool { return b1 == b2 })

}
