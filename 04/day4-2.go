package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"bufio"
	"bytes"
	"regexp"
	"strconv"
	"unicode"
	"strings"
)

func ShiftRune(r rune) rune {
	if r == 'z' {
		return 'a'
	}
	return r + 1
}

func ShiftRoom(s string) {
	var buffer bytes.Buffer
	var shifted bytes.Buffer

	for _, c := range s {
		if unicode.IsNumber(c) {
			buffer.WriteRune(c)
		}
	}

	res, _ := strconv.Atoi(buffer.String())
	name := strings.Replace(strings.Split(s, fmt.Sprintf("%c", buffer.String()[0]))[0], "-", " ", -1)

	for _, c := range name {
		if c != ' ' {
			tmp := c
			for i := 0; i < res; i++ {
				tmp = ShiftRune(tmp)
			}
			shifted.WriteRune(tmp)
		} else {
			shifted.WriteRune(' ')
		}
	}

	matched, _ := regexp.MatchString("northpole.*", shifted.String())
	if matched {
		fmt.Printf("id[%d]: %s\n", res, shifted.String())
	}
}

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ShiftRoom(scanner.Text())
	}

	fmt.Printf("Time elapsed since start : %v\n", time.Since(start))
}
