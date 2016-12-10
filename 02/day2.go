package main

import (
	"fmt"
	"io/ioutil"
)

type Position struct {
	x, y int
}

var keypad = [][]int {
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

var d = map[byte]Position {
	'U': Position{ 0, -1},
	'D': Position{ 0,  1},
	'L': Position{-1,  0},
	'R': Position{ 1,  0},
}

func Move(p Position, b byte) Position {
	res := Position{p.x + d[b].x, p.y + d[b].y}
	if res.x >= 0 && res.x <= 2 && res.y >= 0 && res.y <= 2 {
		return res
	}
	return p
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	curr := Position{1, 1}

	if err != nil {
		panic(err)
	}
	
	for _, c := range dat {
		if c == '\n' {
			fmt.Printf("%d", GetValue(curr));
		} else {
			curr = Move(curr, c)
		}
	}
}
