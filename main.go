package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/nurali-techie/hn/config"
)

func main() {
	fmt.Printf("Hello from HackerNews\n")

	var err error

	if len(os.Args) <= 1 {
		Print("")
		Print("Usage: hn <days> <comma seperated search terms>")
		Print("Example:")
		Print("\thn 5 golang")
		Print("\thn 3 devops,java")
		Print("\thn 2 \"open source\"")
		Print("\thn 7 \"cloud native,microservices\"")
		os.Exit(0)
	}

	var days int
	if len(os.Args) > 1 {
		days, err = strconv.Atoi(os.Args[1])
		if err != nil {
			Err("invalid days param %q", os.Args[1])
			os.Exit(1)
		}
		if days < 1 {
			Err("days param should be greater than 0")
			os.Exit(1)

		}
	}

	var topics []string
	if len(os.Args) > 2 {
		query := os.Args[2]
		if query != "" {
			topics = strings.Split(query, ",")
		}
	}

	search(days, topics)
}

func search(days int, topics []string) {
	now := time.Now()
	past := now.AddDate(0, 0, -days)
	Info("searching for %d days, from %s to %s date", days, DateToString(past), DateToString(now))

	searchUrl := fmt.Sprintf("%s?%s", config.API_SEARCH_BY_DATE, "tags=story&query=%s&numericFilters=created_at_i>%d,created_at_i<%d")

	for _, topic := range topics {
		url := fmt.Sprintf(searchUrl, url.QueryEscape(topic), toSecond(past), toSecond(now))
		items := call(url)
		fmt.Println()
		fmt.Printf("** %s **\n", topic)
		for _, item := range items {
			fmt.Printf("(%d) %s ", item.Points, item.Title)
			PrintHyperLink(item.Url, "(link)")
		}
	}

}

func call(url string) []*item {
	client := http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Get(url)
	if err != nil {
		Err("search failed with error, %v", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		Err("search failed with error, %s", resp.Status)
		os.Exit(1)
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Err("read search result failed with error, %v", err)
		os.Exit(1)
	}

	items, err := parse(content)
	if err != nil {
		Err("parse search result failed with error, %v", err)
		os.Exit(1)
	}

	return items
}

func toSecond(t time.Time) int64 {
	return t.UnixNano() / 1e9
}
