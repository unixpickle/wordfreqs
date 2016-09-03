package main

// A WordEntry is an entry in a frequency dictionary.
type WordEntry struct {
	Word string
	Rank int
	Freq float64
}

// A Fetcher is capable of fetching some online frequency
// dictionary.
type Fetcher interface {
	// Fetch fetches a frequency dictionary.
	// The result is sorted by rank.
	Fetch() ([]WordEntry, error)
}

var FetcherNames = []string{"TV"}

var Fetchers = map[string]Fetcher{
	"TV": &TVDictionaryFetcher{
		URLs: []string{
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/1-1000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/1001-2000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/2001-3000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/3001-4000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/4001-5000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/5001-6000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/6001-7000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/7001-8000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/8001-9000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/9001-10000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/10001-12000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/12001-14000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/14001-16000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/16001-18000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/18001-20000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/20001-22000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/22001-24000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/24001-26000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/26001-28000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/28001-30000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/30001-32000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/32001-34000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/34001-36000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/36001-38000",
			"https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/TV/2006/38001-40000",
		},
	},
}
