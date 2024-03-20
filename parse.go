package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents an HTML link.
type Link struct {
	// Href is a hyperlink reference.
	Href string
	// Text is an optional link text.
	Text string
}

func (l Link) String() string {
	return fmt.Sprintf("Link{%q, %q}", l.Href, l.Text)
}

// Parse parses links from the HTML document.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return processLinks(doc), nil
}

func processLinks(n *html.Node) []Link {
	var links []Link

	var fRec func(n *html.Node)
	fRec = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			var l Link
			for _, a := range n.Attr {
				if a.Key == "href" {
					l.Href = a.Val
					break
				}
			}
			l.Text = text(n)
			links = append(links, l)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			fRec(c)
		}
	}
	fRec(n)
	return links
}

func text(n *html.Node) string {
	var str strings.Builder

	var fRec func(n *html.Node)
	fRec = func(n *html.Node) {
		if n.Type == html.TextNode {
			str.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			fRec(c)
		}
	}
	fRec(n)
	return strings.Join(strings.Fields(str.String()), " ")
}
