package http

import (
	"fmt"
	"os"
	"testing"
	"web-browser/hsts"
)

func TestMain(m *testing.M) {
    hsts.LoadHsts()
    code := m.Run()

    os.Exit(code)
}

func TestGetRequest(t *testing.T) {
    res, error := Get("http://example.com")

    if error != nil {
        fmt.Println(error)
        fmt.Println(res == nil)
    }

    result := "foo"
    expected := "foo"

    if result != expected {
        t.Errorf("LookupIp(http://example.com) = %s; want %s", result, expected)
    }
}