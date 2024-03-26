package main

import (
	"fmt"
	"sort"
)

func main() {

	l := list{
		{title: "moby dick", price: 10, released: toTimestamp(118281600)},
		{title: "odyssey", price: 15, released: toTimestamp("733622400")},
		{title: "hobbit", price: 25},
		//{title: "rubik's cube", price: 5},
		//{title: "minecraft", price: 20},
		//{title: "tetris", price: 5},
		//{title: "yoda", price: 150},
	}

	//sort.Sort(l)

	sort.Sort(sort.Reverse(&byRelease{l}))

	l.discount(.5)
	fmt.Print(l)

}
