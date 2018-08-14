package get

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
)

type Page struct {
	Title       string
	Description string
}

func isDescription(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "name" && attr.Val == "description" {
			return true
		}
	}
	return false
}

func Get(url string) (*Page, error) {
	var page Page
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	reader := io.Reader(resp.Body)
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}

	var f func(node *html.Node)
	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "title" {
			page.Title = node.FirstChild.Data
		} else if node.Type == html.ElementNode && node.Data == "meta" {
			if isDescription(node.Attr) {
				for _, attr := range node.Attr {
					if attr.Key == "content" {
						page.Description = attr.Val
					}
				}
			}
		} else {
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				f(child)
			}
		}
	}
	f(doc)
	return &page, nil
}

func main() {
	p, err := Get("http://voyagegroup.com")
	if err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "Get: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%#v", p)
}
