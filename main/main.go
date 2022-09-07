package main

import (
	"myCollections"
)

func main() {
	vector := myCollections.NewVector[int]()
	for i := 0; i <= 10; i++ {
		vector.PushBack(i)
	}
	vector.Reverse()
	println(vector.ToString())
	vector.Sort()
	vector.DeleteBack()
	println(vector.ToString())
}
