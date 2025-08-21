package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"

	"github.com/tidwall/gjson"
)

var hstsList map[string]EntryInfo

func fetchHsts() {
    url := "https://raw.githubusercontent.com/chromium/chromium/refs/heads/master/net/http/transport_security_state_static.json"

    // Fetch the data
    resp, err := http.Get(url)
    if err != nil {
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

type EntryInfo struct {
    Mode              string
	Policy            string
	IncludeSubdomains bool
    Pins              string
}

func parseEntries(data []byte) map[string]EntryInfo {
	results := gjson.GetBytes(data, "entries")
    entries := make(map[string]EntryInfo)

	results.ForEach(func(_, value gjson.Result) bool {
        name := value.Get("name").String()
        info := EntryInfo{
            Mode: value.Get("mode").String(),
            Policy: value.Get("policy").String(),
            IncludeSubdomains: value.Get("include_subdomains").Bool(),
            Pins: value.Get("pins").String(),
        }

        entries[name] = info
		return true // keep iterating
	})

	return entries
}

func loadHsts() {
    data, err := os.ReadFile("entries.json")
	if err != nil {
		panic(err)
	}

    hstsList = parseEntries(data)
}

func GetHstsByDomain(domain string) (EntryInfo, bool) {
    e, ok := hstsList[domain]
	return e, ok
}

func main() {
    fetchHsts()
    loadHsts()
}
