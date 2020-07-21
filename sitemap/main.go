package sitemap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"gophercises/link"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Urlset struct {
	XMLName xml.Name `xml:"urlset"`
	URL     []URL    `xml:"url"`
}

type URL struct {
	Loc string `xml:"loc"`
}

// NewSitemap gets a html page, parses all its links and returns a XML file as sitemap format
func NewSitemap(urlPath string) ([]byte, error) {

	resp, err := http.Get("https://gophercises.com")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	links, err := link.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}

	for _, href := range hrefs {
		fmt.Println(href)
	}

	return nil, nil
}

func readHTMLPage(url string) (io.Reader, error) {
	resp, err := http.Get("https://www.calhoun.io/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return bytes.NewReader(body), nil
}
