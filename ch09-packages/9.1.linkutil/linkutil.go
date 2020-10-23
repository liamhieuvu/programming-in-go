package linkutil

import (
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

var hrefRx *regexp.Regexp

func init() {
	hrefRx = regexp.MustCompile(`<a[^>]+href=['"]?([^'">]+)['"]?`)
}

func LinksFromURL(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return LinksFromReader(resp.Body)
}

func LinksFromReader(reader io.Reader) ([]string, error) {
	html, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	uniqueLinks := make(map[string]struct{})
	for _, submatch := range hrefRx.FindAllSubmatch(html, -1) {
		uniqueLinks[string(submatch[1])] = struct{}{}
	}

	links := make([]string, 0, len(uniqueLinks))
	for link := range uniqueLinks {
		links = append(links, link)
	}
	return links, nil
}
