package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mbcgo/link/link"
)

func main() {
	htmlFile := flag.String("file", "ex1.html", "path of HTML file to parse")
	flag.Parse()

	if *htmlFile != "" {
		r, err := os.Open(*htmlFile)
		if err != nil {
			log.Fatal(fmt.Sprintf("failed to read file: %s", htmlFile))
		}
		links, err := link.Parse(r)
		if err != nil {
			log.Fatal("Failed to parse links")
		}
		link.Print(links)
	}
}
