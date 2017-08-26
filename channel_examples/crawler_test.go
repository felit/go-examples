package channel_examples

import (
	"testing"
	"os"
	"fmt"
)

func TestCrawler(t *testing.T) {
	worklist := make(chan []string)
	go func() { worklist <- os.Args[1:] }()
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
func crawl(url string) []string {
	fmt.Println(url)
	list := []string{}
	return list
}
