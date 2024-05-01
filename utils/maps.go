/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   maps.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: Xu Huang <xuh15462@gmail.com>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/05/01 22:27:19 by Xu Huang          #+#    #+#             */
/*   Updated: 2024/05/01 22:28:58 by Xu Huang         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import (
	"fmt"
	"maps"

	cmap "github.com/orcaman/concurrent-map/v2"
)

type ThirdPlatform int32

const (
	ThirdPlatform_Ali ThirdPlatform = iota + 1
	ThirdPlatform_Didi
)

func (t ThirdPlatform) String() string {
	switch t {
	case ThirdPlatform_Ali:
		return "Ali"
	case ThirdPlatform_Didi:
		return "Didi"
	default:
		return "Unknown"
	}
}

func NewStringerMap[T MapComparable]() {}

func Test() {
	NewStringerMap[ThirdPlatform]()
	myMap := make(map[string]bool)
	youMap := maps.Clone(myMap)
	maps.DeleteFunc(youMap, func(s string, b bool) bool { return !youMap[s] })
	if maps.Equal(myMap, youMap) {
		fmt.Printf("%s", "equal mymap and you map")
	}
	maps.EqualFunc(myMap, youMap, func(b1, b2 bool) bool { return b1 == b2 })

	resultMap := cmap.New[string]()
	resultMap.Set("1", "11")
	resultMap.Get("1")
	resultMap.Items()
	resultMap.GetShard("1")
}
