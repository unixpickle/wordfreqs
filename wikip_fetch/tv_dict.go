package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/yhat/scrape"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type TVDictionaryFetcher struct {
	URLs []string
}

func (t *TVDictionaryFetcher) Fetch() ([]WordEntry, error) {
	var dict []WordEntry
	for _, u := range t.URLs {
		resp, err := http.Get(u)
		if err != nil {
			return nil, err
		}
		parsed, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, err
		}
		htmlErr := fmt.Errorf("unknown page structure at: %s", u)
		mainContent, ok := scrape.Find(parsed, scrape.ById("mw-content-text"))
		if !ok {
			return nil, htmlErr
		}
		table, ok := scrape.Find(mainContent, scrape.ByTag(atom.Table))
		if !ok {
			return nil, htmlErr
		}
		rows := scrape.FindAll(table, scrape.ByTag(atom.Tr))
		if len(rows) == 0 {
			return nil, htmlErr
		}
		rows = rows[1:]
		for _, row := range rows {
			tds := scrape.FindAll(row, scrape.ByTag(atom.Td))
			if len(tds) < 3 {
				return nil, htmlErr
			}

			// Needed for https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/12001-14000
			// and other evil pages.
			if strings.HasPrefix(scrape.Text(tds[0]), "---") {
				continue
			}

			rank, err := strconv.Atoi(firstField(scrape.Text(tds[0])))
			if err != nil {
				return nil, htmlErr
			}
			count, err := strconv.Atoi(firstField(scrape.Text(tds[2])))
			if err != nil {
				return nil, htmlErr
			}

			dict = append(dict, WordEntry{
				Word: firstWordField(scrape.Text(tds[1])),
				Rank: rank,
				Freq: float64(count),
			})
		}
	}
	return dict, nil
}

func firstField(s string) string {
	f := strings.Fields(s)
	if len(f) == 0 {
		return ""
	}
	return f[0]
}

func firstWordField(s string) string {
	f := strings.Fields(s)
	if len(f) == 0 {
		return ""
	} else if len(f) == 1 {
		return f[0]
	}
	if strings.HasPrefix(f[1], "'") {
		return f[0] + f[1]
	}
	return f[0]
}
