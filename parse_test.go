package link

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var testData = []struct {
	path  string
	links []Link
}{
	{
		path: "ex1.html",
		links: []Link{
			{"/other-page", "A link to another page"},
		},
	},
	{
		path: "ex2.html",
		links: []Link{
			{"https://www.twitter.com/joncalhoun", "Check me out on twitter"},
			{"https://github.com/gophercises", "Gophercises is on Github!"},
		},
	},
	{
		path: "ex3.html",
		links: []Link{
			{"#", "Login"},
			{"/lost", "Lost? Need help?"},
			{"https://twitter.com/marcusolsson", "@marcusolsson"},
		},
	},
	{
		path: "ex4.html",
		links: []Link{
			{"/dog-cat", "dog cat"},
		},
	},
}

func TestParse(t *testing.T) {
	for _, td := range testData {
		f, err := os.Open(filepath.Join("testdata", td.path))
		if err != nil {
			t.Fatal(err)
		}
		links, err := Parse(f)
		if err != nil {
			t.Errorf("Parse(%q) = %v", f.Name(), err)
		}
		if !cmp.Equal(links, td.links) {
			t.Errorf("Parse(%q) = %v, want %v", f.Name(), links, td.links)
		}
	}
}
