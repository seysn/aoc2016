package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"bufio"
	"bytes"
	"unicode"
	"strconv"
)

func IsRealRoom(s string, result chan int) {
	m := make(map[byte]int)
	var checksum_idx int
	var buffer bytes.Buffer

	for i, c := range s {
		if c == '[' {
			checksum_idx = i + 1
			break
		} else if c == '-' {
			continue
		} else if unicode.IsNumber(c) {
			buffer.WriteRune(c)
		} else {
			m[byte(c)]++
		}
	}

	for i := checksum_idx; i < checksum_idx + 5; i++ {
		if m[s[i]] < m[s[i + 1]] || (m[s[i]] == m[s[i + 1]] && s[i] >= s[i + 1]) {
			result <- 0
			return
		}
	}

	res, _ := strconv.Atoi(buffer.String())
	result <- res
}

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nbRooms := 0
	result := make(chan int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		go IsRealRoom(scanner.Text(), result)
		nbRooms++
	}

	sum := 0
	for i := 0; i < nbRooms; i++ {
		sum += <-result
	}

	fmt.Printf("The sum of the sector IDs of the real rooms is %d\n", sum)
	fmt.Printf("Time elapsed since start : %v\n", time.Since(start))
}
