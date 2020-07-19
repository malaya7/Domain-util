package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherw = "*"

var trans = []string{
	otherw,
	otherw + " app",
	otherw + " site",
	otherw + " time",
	"get " + otherw,
	"let's " + otherw,
	otherw + " hq",
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := trans[rand.Intn(len(trans))]
		fmt.Println(strings.Replace(t, otherw, s.Text(), -1))
	}
}
