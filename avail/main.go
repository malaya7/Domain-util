package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func exists(domain string) (bool, error) {
	const whoisServer string = "com.whois-servers.net"
	conn, err := net.Dial("tcp", whoisServer+":43")
	if err != nil {
		return false, err
	}
	defer conn.Close()
	conn.Write([]byte(domain + "\r\n"))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		// fmt.Println("Scanning:" + text)
		if strings.Contains(strings.ToLower(text), "no match") {
			return false, nil
		}
	}
	return true, nil
}

var ok = map[bool]string{true: "✓ Yes", false: "X NO"}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		domain := s.Text()
		here, err := exists(domain)
		if err != nil {
			log.Fatalln(err)
		}
		//fmt.Print(here)
		fmt.Println(domain + ", Available: " + ok[!here])
		time.Sleep(1 * time.Second)
	}
}
