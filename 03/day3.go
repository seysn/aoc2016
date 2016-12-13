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

	nb := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr := MakeArray(strings.Split(scanner.Text(), " "))
		if arr[0]+arr[1] > arr[2] && arr[1]+arr[2] > arr[0] && arr[0]+arr[2] > arr[1] {
			nb++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("There are %d possible triangles\n", nb)
}
