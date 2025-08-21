package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func fetchHsts() {
    url := "https://raw.githubusercontent.com/chromium/chromium/refs/heads/master/net/http/transport_security_state_static.json"

    // Fetch the data
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("asdfasdf")
        panic(err)
    }
    defer resp.Body.Close()

    // Create file
    file, err := os.OpenFile("entries.json", os.O_WRONLY|os.O_TRUNC, 0644)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(resp.Body)

    for scanner.Scan() {
        line := scanner.Text()

        if strings.Contains(line, "//") {
            continue
        }

        _, err := file.WriteString(line + "\n")
        if err != nil {
            panic(err)
        }
    }
}

func main() {
    fetchHsts()
}
