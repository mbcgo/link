package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Link struct {
	Href string
	Text string
}

var links []Link

func Parse(r io.Reader) ([]Link, error) {
	node, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	links = links[:0]
	DFS(node, nil)
	return links, nil
}

func DFS(node *html.Node, data *string) {
	if node == nil {
		return
	}

	if node.DataAtom == atom.A {
		if data == nil {
			links = append(links, newLink(node))
		}
	} else if data != nil && node.Data != node.DataAtom.String() && node.Type != html.CommentNode {
		*data += strings.TrimSpace(node.Data)
	}
	DFS(node.FirstChild, data)
	DFS(node.NextSibling, data)
}

func newLink(node *html.Node) Link {
	// Get Href contents
	var link Link
	for _, a := range node.Attr {
		if a.Key == "href" {
			link.Href = a.Val
		}
	}

	// Get child contents
	DFS(node.FirstChild, &link.Text)
	return link
}

func Print(links []Link) {
	for _, link := range links {
		fmt.Printf("Link{\n    Href: \"%s\",\n    Text: \"%s\",\n}\n", link.Href, link.Text)
	}
}
