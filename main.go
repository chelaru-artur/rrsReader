package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const ItemsFromFeed int = 10

var readers = []string{
	"http://feeds.feedburner.com/Rabotamd-IT?format=xml",
	"http://rss.nytimes.com/services/xml/rss/nyt/Europe.xml"}

func main() {
	allItems := make([]Item, 0)
	for _, url := range readers {
		items, err := getFeed(url)
		if err == nil {
			allItems = append(allItems, items[:ItemsFromFeed]...)
		}
	}
	ParseToHtml(allItems)
	log.Printf("index.html generated with %d items", len(allItems))
}

func getFeed(url string) ([]Item, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	feed, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	result := ParseRss(feed)
	return result.ItemList, nil
}
