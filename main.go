package main

import "github.com/holdonbaby/genericgo/queue"

func main() {
	q := new(queue.Queue[string])
	q.Push("爱情三十六计")
	q.Push("就像一场游戏")
	q.Push("just do it")
	q.HttpRequest()
}
