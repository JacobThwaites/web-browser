package dns

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

func LookupIp(url string) net.IP {
	domain := extractDomain(url)
	ips, err := net.LookupIP(domain)

	if err != nil {
		fmt.Println("Error:", err)
	}

	last := ips[len(ips)-1]
	return last
}