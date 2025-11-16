package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	API_SEARCH_BY_DATE = "https://hn.algolia.com/api/v1/search_by_date"
	API_SEARCH         = "https://hn.algolia.com/api/v1/search"
)

func main() {
	fmt.Printf("Hello from HackerNews ")
	PrintHyperLink(`https://news.ycombinator.com/news`, "(link)")

	var err error

	if len(os.Args) <= 1 {
		Print("")
		Print("Usage: hn <days> <comma seperated search terms>")
		Print("\thn 3 golang		// search golang topic for last 3 days")
		Print("\thn 2 ai,llm		// search ai and llm topics for last 2 days")
		Print("\thn 5 \"open source\"	// use dobule-quotes for search terms with multiple words")
		Print("")
		Print("Usage: hn <points>")
		Print("\thn 200			// search any news with 200-300 points")
		Print("\thn 500			// search any news with 500+ points")
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

	sorted := true
	if len(topics) > 0 {
		searchByTopics(days, topics, sorted)
	} else {
		points := days
		searchByPoints(points)
	}
}

func searchByPoints(points int) {
	days := 2
	now := time.Now()
	past := now.AddDate(0, 0, -days)
	Info("searching for %d days, from %s to %s date", days, DateToString(past), DateToString(now))

	// https://hn.algolia.com/api/v1/search_by_date?tags=story&page=0&numericFilters=created_at_i%3E1745692181,created_at_i%3C1745864981,points%3E500,points%3C1000

	// var searchUrl string
	var startPoints, endPoints int
	if points < 500 {
		// searchUrl = fmt.Sprintf(`%s?%s`, API_SEARCH_BY_DATE, `tags=story&numericFilters=created_at_i>%d,created_at_i<%d,points>=%d,points<=%d&page=%d`)
		startPoints = points
		endPoints = points + 100
	} else {
		// searchUrl = fmt.Sprintf(`%s?%s`, API_SEARCH_BY_DATE, `tags=story&numericFilters=created_at_i>%d,created_at_i<%d,points>=%d,points<=%d&page=%d`)
		startPoints = points
		endPoints = points + 500
	}

	fmt.Println()
	totalPosts := 0
	sorted := true
	for pageNo := 0; ; pageNo++ {
		// urlPath := fmt.Sprintf(searchUrl, toSecond(past), toSecond(now), startPoints, endPoints, pageNo)

		urlPath := API_SEARCH_BY_DATE +
			fmt.Sprintf(`?tags=%s`, `story`) +
			fmt.Sprintf(`&page=%d`, pageNo) +
			fmt.Sprintf(`&numericFilters=%s`, url.QueryEscape(fmt.Sprintf(`created_at_i>%d,created_at_i<%d,points>%d,points<%d`, toSecond(past), toSecond(now), startPoints, endPoints)))

		items := call(urlPath)
		if len(items) == 0 {
			break
		}

		if sorted {
			sort.Slice(items, func(i, j int) bool {
				return items[i].Points > items[j].Points
			})
		}

		for _, item := range items {
			fmt.Printf("(%d) %s ", item.Points, item.Title)
			url := item.Url
			if url == "" {
				url = fmt.Sprintf(`https://news.ycombinator.com/item?id=%s`, item.ObjectID)
			}
			PrintHyperLink(url, "(link)")
		}
		totalPosts += len(items)
		fmt.Printf("------ page=%d, posts=%d ------\n", (pageNo + 1), totalPosts)
		if len(items) < 20 {
			break
		}
	}
}

func searchByTopics(days int, topics []string, sorted bool) {
	now := time.Now()
	past := now.AddDate(0, 0, -days)
	Info("searching for %d days, from %s to %s date", days, DateToString(past), DateToString(now))

	// searchUrl := fmt.Sprintf(`%s?%s%s`, API_SEARCH_BY_DATE, `tags=story&page=%d&`, url.QueryEscape(`query="%s"&numericFilters=created_at_i>%d,created_at_i<%d`))

	for _, topic := range topics {
		totalPosts := 0
		fmt.Println()
		fmt.Printf("** %s **\n", topic)
		for pageNo := 0; ; pageNo++ {
			// url := fmt.Sprintf(searchUrl, url.QueryEscape(topic), pageNo, toSecond(past), toSecond(now))
			urlPath := API_SEARCH_BY_DATE +
				fmt.Sprintf(`?tags=%s`, `story`) +
				fmt.Sprintf(`&page=%d`, pageNo) +
				fmt.Sprintf(`&query=%s`, url.QueryEscape(topic)) +
				fmt.Sprintf(`&numericFilters=%s`, url.QueryEscape(fmt.Sprintf(`created_at_i>%d,created_at_i<%d`, toSecond(past), toSecond(now))))

			items := call(urlPath)
			if len(items) == 0 {
				break
			}

			if sorted {
				sort.Slice(items, func(i, j int) bool {
					return items[i].Points > items[j].Points
				})
			}

			printFooter := false
			for _, item := range items {
				if days == 1 {
					if item.Points >= 20 {
						printItem(item)
						printFooter = true
					}
				} else if days == 2 {
					if item.Points >= 10 {
						printItem(item)
						printFooter = true
					}
				} else {
					printItem(item)
					printFooter = true
				}
			}

			totalPosts += len(items)

			if printFooter {
				fmt.Printf("------ page=%d, posts=%d ------\n", (pageNo + 1), totalPosts)
			}

			if len(items) < 20 {
				break
			}
		}
	}

}

func printItem(item *item) {
	fmt.Printf("(%d) %s ", item.Points, item.Title)
	url := item.Url
	if url == "" {
		url = fmt.Sprintf(`https://news.ycombinator.com/item?id=%s`, item.ObjectID)
	}
	PrintHyperLink(url, "(link)")
}

func call(urlPath string) []*item {
	// urlPath = `https://hn.algolia.com/api/v1/search_by_date?tags=story&query=%22java%22&numericFilters=created_at_i%3E1745692181,created_at_i%3C1745864981&page=0`

	// fmt.Println("url=", urlPath)
	// urlPart := strings.Split(urlPath, `?`)
	// urlPath = fmt.Sprintf(`%s?%s`, urlPart[0], url.QueryEscape(urlPart[1]))
	// fmt.Println("url=", urlPath)

	// fmt.Println("url=", urlPath)
	// urlPath = `https://hn.algolia.com/api/v1/search_by_date?tags=story&query=%22java%22&numericFilters=created_at_i%3E1745692181,created_at_i%3C1745864981&page=0`
	// urlPath = `https://hn.algolia.com/api/v1/search_by_date?tags=story&page=0&query=%22java%22&numericFilters=created_at_i%3E1745695548,created_at_i%3C1745868348`
	// fmt.Println("url=", urlPath)

	client := http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Get(urlPath)
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
