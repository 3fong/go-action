package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// func main() {

// 	bigSlowOperation()
// }

func cyclePage(link func(url string) error) {
	for _, url := range os.Args[1:] {
		link(url)
	}
}

func getlinks(url string) error {
	fmt.Println("getlinks in process...")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return err
	}
	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		return err
	}
	forEachNode(doc, startElement, endElement)
	return nil
}

// ???html???link??
func title(url string) error {
	fmt.Println("title in process...")
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// Check Content-Type is HTML (e.g., "text/html;charset=utf-8").
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}

func paincTest() {
	var f func(int) int
	f(3)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func closure() {
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func() {
			fmt.Println(v)
		}()
		fmt.Println("out value:", v)
	}
	time.Sleep(time.Second * 2)
}

func innerFunc(f func() int) {
	// f = squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func tempDirs() []string {
	return []string{"c://a", "c://a/b", "c://a/b/c"}
}

func partVariable() {
	var rmdirs []func()

	for _, d := range tempDirs() {
		dir := d               // NOTE: necessary!
		os.MkdirAll(dir, 0755) // creates parent directories too
		fmt.Println(dir)
		rmdirs = append(rmdirs, func() {
			fmt.Println("inner", dir)
			os.RemoveAll(dir)
		})
	}
	// ...do some work…
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}
}

func formatStr() {
	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name)
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, "zhangsan=%s:", "args...")
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

func printFile() {
	for _, url := range os.Args[1:] {
		var b, err = ReadFile(url)
		if err != nil {
			fmt.Println("err: ", err)
		}
		fmt.Println(b)
	}
}

// ReadFile ??????
func ReadFile(filename string) ([]byte, error) {
	filename = "C:\\Users\\bula\\Desktop\\Shiro.md"
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses

	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
