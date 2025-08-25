package main

import (
	"fmt"
	"net"

	"web-browser/dns"
	"web-browser/hsts"
)

func main() {
	hsts.LoadHsts()
	url := "https://example.com"
	ipAddress := dns.LookupIp(url)
	fmt.Println("IP Used: " + ipAddress.String())

	// TODO: implement HSTS and logic for deciding whether to use HTTP or HTTPS
	// HTTPS_PORT := 443

	conn, err := net.Dial("tcp", ipAddress.String())
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer conn.Close()

}