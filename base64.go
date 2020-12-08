package main

import (
	"fmt"
	"strconv"
	"strings"
)

func padStart(origin string, length int, pad string) string {
	r := []rune(origin)
	n := length - len(r)
	if n <= 0 {
		return origin
	}

	padRune := []rune(pad)
	padLen := len(padRune)
	quo := n / padLen
	rem := n % padLen

	return strings.Repeat(pad, quo) + string(padRune[0:rem]) + origin
}

func strToBin(str string) []string {
	hexes := []byte(str)
	len := len(hexes)
	bins := make([]string, len)
	for i, v := range hexes {
		bins[i] = strconv.FormatInt(int64(v), 2)
	}
	return bins
}

func main() {
	fmt.Println(strToBin("あいうえお"))
	fmt.Println(strToBin("abcde"))
}
