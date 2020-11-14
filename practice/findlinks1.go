package practice

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/html"
)

var text string = `0.0.0.0   account.jetbrains.com
0.0.0.0	www.jetbrains.com
59.110.226.9 www.liu.com
5.110.226.9		www.niu.com
::1        localhost`

// func main() {
// 	regexp.Compile()

// }

func runFindlinks() {
	for _, url := range os.Args[1:] {
		links, err := Findlinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Findlinks: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func Findlinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	return visit(nil, doc), nil
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func Fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

func runFetch() {
	for _, url := range os.Args[1:] {
		fetch2(url)
	}
}

func fetch2(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer f.Close()
	n, err = io.Copy(f, resp.Body)
	return local, n, err
}
