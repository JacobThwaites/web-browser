package dnsLookup

import (
	"fmt"
	"net"
	"regexp"
)

func extractDomain(url string) string {
	re := regexp.MustCompile(`^(?:https?://)?(?:www\.)?([^/]+)`)
    match := re.FindStringSubmatch(url)

    if len(match) > 1 {
		return match[1]
    } else {
		return ""
    }
}

func lookupIp(url string) {
	ips, err := net.LookupIP(url)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, ip := range ips {
		fmt.Println("IP address:", ip)
	}
}