package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	args := flag.Args()
	fmt.Println(args)

	if len(args) < 1 {
		fmt.Println("Please specify start page")
		os.Exit(1)
	}
	retrieve(args[0])
}

func retrieve(uri string) {
	resp, err := http.Get(uri)

	if err != nil {
		fmt.Println("Get error is:", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	reader := strings.NewReader(string(body))
	root, err := html.Parse(reader)

	if err != nil {
		fmt.Println("error message is:", err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Println("Get tag")
			fmt.Println("  Data is:", n.Data)
			fmt.Println("  Attributes are:")
			for _, a := range n.Attr {
				// fmt.Println("    Key:", a.Key)
				fmt.Println("    Value:", a.Val)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(root)
}
