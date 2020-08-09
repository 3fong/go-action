package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

func main() {
	// for _, url := range os.Args[1:] {
    //    findlinks(url)
	// }
	formatStr()
}

func findlinks(url string) {
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
	forEachNode(doc,startElement,endElement)
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

func closure(){
	s := []string{"a", "b", "c"}                             
    for _, v := range s { 
        go func() {
            fmt.Println(v)
		}()     
		fmt.Println("out value:",v)            
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
	return []string{"c://a","c://a/b","c://a/b/c"}
}

func partVariable(){
	var rmdirs []func()

	for _, d := range tempDirs() {
		dir := d // NOTE: necessary!
		os.MkdirAll(dir, 0755) // creates parent directories too
		fmt.Println(dir)
		rmdirs = append(rmdirs, func() {
			fmt.Println("inner",dir)
			os.RemoveAll(dir)
		})
	}
	// ...do some work…
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}
}

func formatStr(){
	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name)
}

func errorf(linenum int, format string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
    fmt.Fprintf(os.Stderr, "zhangsan=%s:", "args...")
    fmt.Fprintf(os.Stderr, format, args...)
    fmt.Fprintln(os.Stderr)
}
