package sitemap

import (
	"bytes"
	"encoding/xml"
	"gophercises/link"
	"io"
	"io/ioutil"
	"net/http"
)

type Urlset struct {
	XMLName xml.Name `xml:"urlset"`
	URL     []URL    `xml:"url"`
}

type URL struct {
	Loc string `xml:"loc"`
}

// NewSitemap gets a html page, parses all its links and returns a XML file as sitemap format
func NewSitemap(url string) ([]byte, error) {

	reader, _ := readHTMLPage(url)
	links, err := link.Parse(reader)
	if err != nil {
		return nil, err
	}

	var xmlDoc *Urlset
	var urls []URL
	for _, link := range links {
		urls = append(urls, URL{link.Href})
		xmlDoc = &Urlset{URL: urls}
	}

	file, _ := xml.MarshalIndent(xmlDoc, " ", " ")

	return file, nil
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
