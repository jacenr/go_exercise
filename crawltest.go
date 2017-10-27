package main

import (
	// "bytes"
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	// "io"
	"io/ioutil"
	"log"
	"net/http"
	urlp "net/url"
	"os"
	"path"
	"strings"
)

var suffix string

// var buf *bytes.Buffer = bytes.NewBuffer([]byte{})
// var buf bytes.Buffer

func main() {
	urlstruc, err := urlp.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	urlhost := urlstruc.Host
	urlslc := strings.Split(urlhost, ".")
	l := len(urlslc)
	// fmt.Println(urlslc[l-2:])
	suffix = strings.Join(urlslc[l-2:], ".")

	// if os.MkdirAll(, os.ModeDir) != nil {
	// 	log.Fatal(err)
	// }

	// Crawl the web breadthfirst,
	// starting from the commandline arguments.
	breadthFirst(crawl, os.Args[1:])
}

func crawl(url string) []string {
	// fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			item = strings.TrimRight(item, "/")
			if !strings.HasPrefix(item, "http") {
				continue
			}
			urlstruc, err := urlp.Parse(item)
			if err != nil {
				log.Fatal(err)
			}
			urlhost := urlstruc.Host
			if !strings.HasSuffix(urlhost, suffix) {
				continue
			}
			fmt.Println(item)
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func savepage(url string, doc []byte) error {
	// defer body.Close()
	urlstruc, err := urlp.Parse(url)
	if err != nil {
		log.Fatal(err)
	}
	urlhost := urlstruc.Host
	// fmt.Println("host:" + urlhost)
	urlpath := strings.TrimRight(urlstruc.Path, "/")
	// fmt.Println("path:" + urlpath)
	dirpath := urlhost + "/" + path.Dir(urlpath)
	// fmt.Println("dirpath:" + dirpath)
	filename := path.Base(urlpath)
	if urlpath == "" {
		dirpath = urlhost
		filename = urlhost
	}
	// fmt.Println("filename" + dirpath + filename)
	// filename := path.Base(urlpath)
	if os.MkdirAll(dirpath, os.ModeDir) != nil {
		// log.Fatal("Can't create dir: " + dirpath)
		// return error("Can't create dir: " + dirpath)
		fmt.Errorf("%s%s", "Can't create dir: ", dirpath)
	}
	// fmt.Println(*body)
	// var doc []byte
	// if n, err := body.Read(doc); err != nil && err != io.EOF {
	// 	// log.Fatal(err)
	// 	fmt.Println("1")
	// 	fmt.Println(n)
	// 	fmt.Println(doc)
	// 	return err
	// }
	// doc, err := ioutil.ReadAll(body)
	// body = ioutil.NopCloser(bytes.NewBuffer(doc))
	// fmt.Println(doc)
	// var doc []byte = []byte("test")
	if err != nil {
		return err
	}
	// doc := buf.Bytes()
	f, err := os.Create(dirpath + "/" + filename)
	defer f.Close()
	if err != nil {
		// log.Fatal(err)
		fmt.Println("2")
		return err
	}
	if _, err := f.Write(doc); err != nil {
		fmt.Println("3")
		// log.Fatal(err)
		return err
	}
	if err := f.Sync(); err != nil {
		// log.Fatal(err)
		fmt.Println("4")
		return err
	}
	// fmt.Println("ok")
	return nil
}

func Extract(url string) ([]string, error) {
	// urlstrc := urlp.Parse(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// var body2 io.ReadCloser
	// buf.Reset()
	// io.Copy(buf, resp.Body)
	// resp.Write(buf)
	// ioutil.NopCloser(resp.Body)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	if err := savepage(url, body); err != nil {
		log.Fatalf("%v%v", fmt.Errorf("%s", "Save Page error: "), err)
	}

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	doc, err := html.Parse(resp.Body)

	body = []byte{}

	// fmt.Println(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		// fmt.Println(n.Data)
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
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
