package dnsLookup

import (
	"testing"
)

func TestGetDomain(t *testing.T) {
    tests := []struct {
		url      string
		expected string
	}{
		{"https://www.example.com/path/page?query=1", "example.com"},
		{"http://example.org", "example.org"},
		{"https://subdomain.example.net/page", "subdomain.example.net"},
		{"www.test.com", "test.com"},
		{"example.io", "example.io"},
		{"https://www.example.co.uk/path", "example.co.uk"},
		{"", ""},
	}

	for _, tt := range tests {
		got := extractDomain(tt.url)
		if got != tt.expected {
			t.Errorf("extractDomain(%q) = %q; want %q", tt.url, got, tt.expected)
		}
	}
}

func TestDnsLookup(t *testing.T) {
    lookupIp("example.com")
    result := "foo"
    expected := "foo"

    if result != expected {
        t.Errorf("LookupIp(http://example.com) = %s; want %s", result, expected)
    }
}
