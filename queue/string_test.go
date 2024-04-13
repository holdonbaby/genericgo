package queue

import (
	"fmt"
	"slices"
	"testing"
)

func TestSchemeItemInArray(t *testing.T) {
	u1 := new(User)
	u1.UserId = "123456"
	u2 := new(User)
	u2.UserId = "457645"
	u3 := new(User)
	u3.UserId = "4575687"
	u4 := new(User)
	u4.UserId = "4575687"
	u5 := new(User)
	u5.UserId = "4575687"
	var array []*User
	if !slices.ContainsFunc(array, func(u *User) bool { return u.GetUniqueKey() == u4.GetUniqueKey() }) {
		array = append(array, u4)
	}
	if !slices.ContainsFunc(array, func(u *User) bool { return u.GetUniqueKey() == u3.GetUniqueKey() }) {
		array = append(array, u3)
	}
	if !slices.ContainsFunc(array, func(u *User) bool { return u.GetUniqueKey() == u2.GetUniqueKey() }) {
		array = append(array, u2)
	}
	if !slices.ContainsFunc(array, func(u *User) bool { return u.GetUniqueKey() == u1.GetUniqueKey() }) {
		array = append(array, u1)
	}
	if slices.ContainsFunc(array, func(u *User) bool { return u.GetUniqueKey() == u1.GetUniqueKey() }) {
		array = append(array, u1)
	}
	fmt.Printf("%s \n", array)

	// slices.SortFunc(array, func(a, b *User) int {
	// 	if a.GetUniqueKey() < b.GetUniqueKey() {
	// 		return 1
	// 	} else if a.GetUniqueKey() < b.GetUniqueKey() {
	// 		return -1
	// 	}
	// 	return 0
	// })
	// fmt.Printf("%s", slices.CompactFunc(array, func(u1, u2 *User) bool { return u1.GetUniqueKey() == u2.GetUniqueKey() }))

	fmt.Printf("%s", Duplicate(array, func(i, j *User) bool { return i.GetUniqueKey() == j.GetUniqueKey() }))
}
