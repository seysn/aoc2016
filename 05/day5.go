package main

import (
	"fmt"
	"bytes"
	"regexp"
	"crypto/md5"
	"encoding/hex"
)

// TODO : Try the challenge with goroutines
func main() {
	var buffer bytes.Buffer
	var validHash = regexp.MustCompile(`^0{5}.*$`)
	count := 0
	input := "ugkcyxxp"

	for len(buffer.String()) < 8 {
		tmp := md5.Sum([]byte(fmt.Sprintf("%s%d", input, count)))
		if str := hex.EncodeToString(tmp[:]); validHash.MatchString(str) {
			buffer.WriteByte(str[5])
		}
		count++
	}

	fmt.Println(buffer.String())
}
