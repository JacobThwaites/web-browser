package http

import (
	"bufio"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"web-browser/dns"
	"web-browser/hsts"
	htmlparser "web-browser/htmlParser"
)

type Response struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0
}

func Get(url string) ([]byte, error) {
	domain := dns.ExtractDomain(url)
	ipAddr := dns.LookupIp(url)

	port := "80"
	useTLS := false

	if _, ok := hsts.GetHstsByDomain(domain); ok {
		port = "443"
		useTLS = true
	}

	address := ipAddr.String() + ":" + port

	var conn net.Conn
	var err error

	if useTLS {
		// Use TLS with correct SNI
		conn, err = tls.Dial("tcp", address, &tls.Config{ServerName: domain})
	} else {
		conn, err = net.Dial("tcp", address)
	}

	if err != nil {
		return nil, fmt.Errorf("dial error: %w", err)
	}

	defer conn.Close()

	// Write request
	req := "GET / HTTP/1.1\r\n" +
		"Host: " + domain + "\r\n" +
		"Connection: close\r\n" +
		"\r\n"

	_, err = conn.Write([]byte(req))
	if err != nil {
		return nil, fmt.Errorf("write error: %w", err)
	}

	// Read response
	var resp strings.Builder
	reader := bufio.NewReader(conn)
	for {
		chunk, err := reader.ReadString('\n')
		resp.WriteString(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("read error: %w", err)
		}
	}

	ParseHttp([]byte(resp.String()))

	return []byte(resp.String()), nil
}

func splitHeaderBody(html []byte) (string, string, error) {
	input := string(html)
    parts := strings.SplitN(input, "\r\n\r\n", 2)
    if len(parts) < 2 {
        parts = strings.SplitN(input, "\n\n", 2)
    }
    if len(parts) < 2 {
        return "", "", errors.New("invalid HTTP response: no header/body split")
    }
    return parts[0], parts[1], nil
}

func ParseHttp(input []byte) (htmlparser.DomElement, error) {
	rawHeaders, body, err := splitHeaderBody(input)

	if err != nil {
		fmt.Println(rawHeaders)
		return htmlparser.DomElement{}, err
	}

	domTree, err := htmlparser.ParseHTML([]byte(body))

	if err != nil {
		fmt.Println(err)
	}

	return domTree, err
}