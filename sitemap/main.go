package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"gophercises/link"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	xmlns string = "http://www.sitemaps.org/schemas/sitemap/0.9"
)

type Url struct {
	Loc string `xml:"loc"`
}

type SiteMap struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Urlset  []Url    `xml:"url"`
}

func main() {
	urlFlag := flag.String("url", "http://gophercises.com", "the url to build sitemap for")
	maxDepthFlag := flag.Int("depth", 10, "maximum number of links deep to traverse")

	flag.Parse()

	pages := bfs(*urlFlag, *maxDepthFlag)

	toXml := SiteMap{
		Xmlns: xmlns,
	}

	for _, page := range pages {
		toXml.Urlset = append(toXml.Urlset, Url{page})
	}

	wr := os.Stdout
	fmt.Fprint(wr, xml.Header)
	enc := xml.NewEncoder(wr)
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}

}

func bfs(urlStr string, depth int) []string {
	visited := make(map[string]struct{})
	var queue []string
	var next_queue []string

	queue = append(queue, urlStr)
	visited[urlStr] = struct{}{}

	for i := 0; i <= depth; i++ {
		for _, url := range queue {
			for _, next_url := range get(url) {
				if _, ok := visited[next_url]; ok {
					continue
				}
				visited[next_url] = struct{}{}
				next_queue = append(next_queue, next_url)
			}

		}

		queue = next_queue
		next_queue = []string{}

	}
	ret := make([]string, 0, len(visited))
	for url := range visited {
		ret = append(ret, url)
	}

	return ret
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}

	return filter(hrefs(resp.Body, baseUrl.String()), withPrefix(baseUrl.String()))
}

func hrefs(r io.Reader, base string) []string {
	var ret []string

	links, _ := link.Parse(r)
	for _, l := range links {
		var url string
		switch {
		case strings.HasPrefix(l.Href, "/"):
			url = base + l.Href
		case strings.HasPrefix(l.Href, "http"):
			url = l.Href
		default:
			continue
		}
		ret = append(ret, strings.TrimRight(url, "/"))
	}
	return ret
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, l := range links {
		if keepFn(l) {
			ret = append(ret, l)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(s string) bool {
		return strings.HasPrefix(s, pfx)
	}
}
