package main

import (
	"fmt"
	"math"
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

	switch len(os.Args) {
	case 1: // hn
		printHelp()
	case 2: // hn 200
		points, err := strconv.Atoi(os.Args[1])
		if err != nil {
			Err("invalid value for points %q, points should be number", os.Args[1])
			os.Exit(1)
		}
		if points < 1 {
			points = 100
		}
		searchByPoints(points)
	case 3, 4: // hn 3 golang / hn 3 golang -a
		days, err := strconv.Atoi(os.Args[1])
		if err != nil {
			Err("invalid value for days %q, days should be number", os.Args[1])
			os.Exit(1)
		}
		if days < 1 {
			days = 1
		}
		topics := strings.Split(os.Args[2], ",")
		minPoints := 10
		if len(os.Args) == 4 {
			if os.Args[3] != "-a" {
				Err("invalid value for last param, it must be -a")
				os.Exit(1)
			}
			minPoints = 1
		}
		searchByTopics(days, topics, minPoints)
	default:
		Err("Error: Invalid input having extra values")
		printHelp()
	}
}

func printHelp() {
	Print("")
	Print("Usage-1: hn <days> <comma seperated search topics>")
	Print("\thn 3 golang		// search golang topic for last 3 days (with 10+ points)")
	Print("\thn 1 ai,llm		// search both ai and llm topics for last 1 days (with 10+ points)")
	Print("\thn 2 \"open source\"	// use dobule-quotes for search topic having multiple words")
	Print("")
	Print("Usage-2: hn <days> <comma seperated search topics> -a")
	Print("\thn 3 golang -a		// search golang topic for last 3 days (with 1+ points)")
	Print("")
	Print("Usage-3: hn <points>")
	Print("\thn 200			// search any news from last 2 days with 200 to 300 points")
	Print("\thn 500			// search any news from last 2 days with 500+ points")
}

func searchByPoints(points int) {
	days := 2
	now := time.Now()
	past := now.AddDate(0, 0, -days)

	var startPoints, endPoints int
	if points < 500 {
		startPoints = points
		endPoints = points + 100
		Info("searching from last %d days, between %d to %d points", days, startPoints, endPoints)
	} else {
		startPoints = points
		endPoints = math.MaxInt
		Info("searching from last %d days, %d+ points", days, startPoints)
	}

	fmt.Println()
	totalPosts := 0
	sorted := true
	for pageNo := 0; ; pageNo++ {
		// url=https://hn.algolia.com/api/v1/search_by_date?tags=story&page=0&numericFilters=created_at_i>1763489663,created_at_i<1763662463,points>200,points<300
		url := API_SEARCH_BY_DATE +
			fmt.Sprintf(`?tags=%s`, `story`) +
			fmt.Sprintf(`&page=%d`, pageNo) +
			fmt.Sprintf(`&numericFilters=%s`, url.QueryEscape(fmt.Sprintf(`created_at_i>%d,created_at_i<%d,points>%d,points<%d`, toSecond(past), toSecond(now), startPoints, endPoints)))

		items := searchCall(url)
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

func searchByTopics(days int, topics []string, minPoints int) {
	Info("searching from last %d days, for topics %v, with %d+ points", days, topics, minPoints)

	now := time.Now()
	past := now.AddDate(0, 0, -days)

	for _, topic := range topics {
		totalPosts := 0
		fmt.Println()
		fmt.Printf("** %s **\n", topic)
		for pageNo := 0; ; pageNo++ {
			// url=https://hn.algolia.com/api/v1/search_by_date?tags=story&page=0&query=golang&numericFilters=created_at_i>1763403219,created_at_i<1763662419
			url := API_SEARCH_BY_DATE +
				fmt.Sprintf(`?tags=%s`, `story`) +
				fmt.Sprintf(`&page=%d`, pageNo) +
				fmt.Sprintf(`&query=%s`, url.QueryEscape(topic)) +
				fmt.Sprintf(`&numericFilters=%s`, url.QueryEscape(fmt.Sprintf(`created_at_i>%d,created_at_i<%d`, toSecond(past), toSecond(now))))

			items := searchCall(url)
			if len(items) == 0 {
				break
			}

			sort.Slice(items, func(i, j int) bool {
				return items[i].Points > items[j].Points
			})

			printFooter := false
			for _, item := range items {
				if item.Points >= minPoints {
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

func toSecond(t time.Time) int64 {
	return t.UnixNano() / 1e9
}
