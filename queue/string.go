package queue

import (
	"cmp"
	"net/http"
	"slices"

	"github.com/bytedance/sonic"
	"golang.org/x/exp/constraints"
)

type Fuck struct {
}

type Caler[T cmp.Ordered] interface {
	Cal(T) T
}

type Stringer[E comparable] interface {
	string() E
}

type List[T cmp.Ordered] []T

func repeated[T cmp.Ordered](list []T) []T {
	var newList []T
	maper := make(map[T]bool)
	for _, i := range list {
		if _, ok := maper[i]; ok {
			continue
		}
		maper[i] = true
	}
	return newList
}

type User struct {
	Name, UserId, Email, Department string
	Age, MobilePhone                int
}

func (u *User) GetUniqueKey() string {
	return u.UserId
}

func (u *User) String() string {
	if u == nil {
		return "<Nil>"
	}
	str, err := sonic.MarshalString(u)
	if err != nil {
		return "<Marshal Failed>"
	}
	return str
}

func Cal[T constraints.Integer](t T) T {
	return t * 100
}

// Duplicate 数组去重
func Duplicate[T any](array []T, equalFunc func(i, j T) bool) []T {
	var newArray []T
	for _, i := range array {
		if !slices.ContainsFunc(newArray, func(t T) bool { return equalFunc(t, i) }) {
			newArray = append(newArray, i)
		}
	}
	return newArray
}

type server struct {
}

// HttpReq Function
func (s *server) HttpReq(req *http.Request) {
	client := new(http.Client)
	client.Do(req)
}
