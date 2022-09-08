// This is testing file

package main

import (
	"fmt"
	mc "myCollections"
)

func main() {
	pr := mc.NewPair[int](1, 2)
	pr1 := pr.Copy()
	pr.Swap()
	println(pr.Sum())
	fmt.Println(pr, pr1)
	vec := mc.NewVector[uint]()
	var i uint
	for i = 0; i < 100; i++ {
		vec.PushBack(i)
	}
	vec.Reverse()
	vec.Sort()
	println(vec.ToString())

}
