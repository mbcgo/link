package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	htmlFile := flag.String("file", "ex1.html", "path of HTML file to parse")
	flag.Parse()

	f, err := os.Open(*htmlFile)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to read file: %s", htmlFile))
	}
	node, err := html.Parse(f)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to parse html file %s", htmlFile))
	}

	DFS(node, nil)
}

func DFS(node *html.Node, data *string) {
	if node == nil {
		return
	}

	if node.DataAtom == atom.A {
		if data == nil {
			PrintLinkContents(node)
		}
	} else if data != nil && node.Data != node.DataAtom.String() {
		*data += strings.TrimSpace(node.Data)
	}
	DFS(node.FirstChild, data)
	DFS(node.NextSibling, data)
}

func PrintLinkContents(node *html.Node) {
	// Get Href contents
	var href string
	for _, a := range node.Attr {
		if a.Key == "href" {
			href = a.Val
		}
	}

	// Get child contents
	var data string
	DFS(node.FirstChild, &data)

	fmt.Printf("Link{\n    Href: \"%s\",\n    Text: \"%s\",\n}\n", href, data)
}
