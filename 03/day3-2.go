package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func MakeArray(s []string) []int {
	var r []int
	for _, str := range s {
		if str != "" {
			if i, err := strconv.Atoi(str); err == nil {
				r = append(r, i)
			}
		}
	}
	return r
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	arr := [3][]int{}
	nb, i := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr[i%3] = MakeArray(strings.Split(scanner.Text(), " "))
		i++
		if i > 2 {
			i = 0
			for j := 0; j < 3; j++ {
				if arr[0][j]+arr[1][j] > arr[2][j] && arr[1][j]+arr[2][j] > arr[0][j] && arr[0][j]+arr[2][j] > arr[1][j] {
					nb++
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("There are %d possible triangles\n", nb)
}
