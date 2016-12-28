package main

import (
	"fmt"
	"regexp"
	"crypto/md5"
	"encoding/hex"
)

func PrintTab(t []byte) {
	for _, b := range t {
		if b == 0 {
			fmt.Printf("_")
		} else {
			fmt.Printf("%c", b)
		}
	}
	fmt.Println()
}

func IsFull(t []byte) bool {
	for _, b := range t {
		if b == 0 {
			return false
		}
	}
	return true
}

func main() {
	var buffer = make([]byte, 8)
	var validHash = regexp.MustCompile(`^0{5}.*$`)
	count := 0
	input := "ugkcyxxp"
	
	for !IsFull(buffer) {
		tmp := md5.Sum([]byte(fmt.Sprintf("%s%d", input, count)))
		if str := hex.EncodeToString(tmp[:]); validHash.MatchString(str) {
			if str[5] >= '0' && str[5] <= '7' && buffer[str[5] - '0'] == 0 {
				fmt.Printf("%s\n", str)
				buffer[str[5] - '0'] = str[6]
				PrintTab(buffer)
			}
		}
		count++
	}
}
