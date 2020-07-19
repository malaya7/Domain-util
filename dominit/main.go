package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var tlds = []string{"com", "net", "gov"}

const okChars = "wertyuioplkjhgfdsazxcvbnm1234567890-_"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newTxt []rune
		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}
			if !strings.ContainsRune(okChars, r) {
				continue
			}
			newTxt = append(newTxt, r)
		}
		fmt.Println(string(newTxt) + "." + tlds[rand.Intn(len(tlds))])
	}
}
