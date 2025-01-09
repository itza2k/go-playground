package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// fetches url content
func fetchURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error fetching URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	body, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error parsing HTML: %v", err)
	}

	var sb strings.Builder
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			sb.WriteString(n.Data)
		}
		if n.FirstChild != nil {
			f(n.FirstChild)
		}
		if n.NextSibling != nil {
			f(n.NextSibling)
		}
	}
	f(body)

	return sb.String(), nil
}

func main() {
	url := "https://itza2k.github.io"
	content, err := fetchURL(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Fetched content:")
	fmt.Println(content)
}
