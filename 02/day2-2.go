package main

import (
	"fmt"
	"io/ioutil"
)

type Position struct {
	x, y int
}

var keypad = [][]byte {
	{'#', '#', '1', '#', '#'},
	{'#', '2', '3', '4', '#'},
	{'5', '6', '7', '8', '9'},
	{'#', 'A', 'B', 'C', '#'},
	{'#', '#', 'C', '#', '#'},
}

var d = map[byte]Position {
	'U': Position{ 0, -1},
	'D': Position{ 0,  1},
	'L': Position{-1,  0},
	'R': Position{ 1,  0},
}

func GetValue(p Position) byte {
	return keypad[p.y][p.x]
}

func Move(p Position, b byte) Position {
	res := Position{p.x + d[b].x, p.y + d[b].y}
	if res.x >= 0 && res.x <= 4 && res.y >= 0 && res.y <= 4 {
		if GetValue(res) == '#' {
			return p
		}
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
			fmt.Printf("%c", GetValue(curr));
		} else {
			curr = Move(curr, c)
		}
	}
}
