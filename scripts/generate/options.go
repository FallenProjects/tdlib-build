package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

var (
	rowRegex = regexp.MustCompile(`(?s)<tr>\s*<td>(.*?)</td>\s*<td>(.*?)</td>\s*<td>(.*?)</td>\s*<td>(.*?)</td>\s*</tr>`)
	tagRegex = regexp.MustCompile(`<[^>]*>`)
)

func getOptions() (map[string]*OptionDef, error) {
	url := "https://core.telegram.org/tdlib/options"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch options: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch options: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read options body: %v", err)
	}

	content := string(body)

	startIndex := strings.Index(content, "list-of-options-supported-by-tdlib")
	if startIndex == -1 {
		return nil, fmt.Errorf("could not find options list in HTML")
	}
	content = content[startIndex:]

	options := make(map[string]*OptionDef)
	matches := rowRegex.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		name := cleanHTML(match[1])
		typeName := cleanHTML(match[2])
		writableStr := cleanHTML(match[3])
		description := cleanHTML(match[4])

		// Skip header
		if name == "Name" || name == "" {
			continue
		}

		options[name] = &OptionDef{
			Type:        mapType(typeName),
			Writable:    mapWritable(writableStr),
			Description: description,
		}
	}

	return options, nil
}

func cleanHTML(s string) string {
	s = tagRegex.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, "&amp;", "&")
	s = strings.ReplaceAll(s, "&lt;", "<")
	s = strings.ReplaceAll(s, "&gt;", ">")
	s = strings.ReplaceAll(s, "&quot;", "\"")
	s = strings.ReplaceAll(s, "&#39;", "'")
	return strings.TrimSpace(s)
}

func mapType(t string) string {
	switch t {
	case "Integer":
		return "int64"
	case "Boolean":
		return "Bool"
	case "String":
		return "string"
	default:
		return t
	}
}

func mapWritable(w string) bool {
	return strings.ToLower(w) == "yes"
}
