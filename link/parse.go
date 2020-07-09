package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

// Parse receives a Reader representing a html, parses it and extract all anchors and their texts
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	anchors := links(doc)
	var links []Link
	for _, anchor := range anchors {
		links = append(links, buildLink(anchor))
	}

	return links, nil
}

func buildLink(n *html.Node) Link {
	var href string
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			href = attr.Val
			break
		}
	}

	return Link{href, getAllText(n)}
}

func links(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var anchors []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		anchors = append(anchors, links(c)...)
	}
	return anchors
}

func getAllText(n *html.Node) string {

	var text string
	if n.Type == html.TextNode && len(strings.TrimSpace(n.Data)) > 0 {
		return n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += getAllText(c)
	}

	return strings.Join(strings.Fields(text), " ")

}
