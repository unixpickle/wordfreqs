package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "list_name output_file")
		fmt.Fprintln(os.Stderr, "Available lists:")
		for _, name := range FetcherNames {
			fmt.Fprintln(os.Stderr, " -", name)
		}
		fmt.Fprintln(os.Stderr)
		os.Exit(1)
	}

	fetcher := Fetchers[os.Args[1]]
	if fetcher == nil {
		fmt.Fprintln(os.Stderr, "Unknown list:", os.Args[1])
		os.Exit(1)
	}

	results, err := fetcher.Fetch()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fetch failed:", err)
		os.Exit(1)
	}

	outFile, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to write output:", err)
		os.Exit(1)
	}
	defer outFile.Close()
	csvWriter := csv.NewWriter(outFile)

	for _, res := range results {
		err = csvWriter.Write([]string{res.Word, strconv.Itoa(res.Rank),
			strconv.FormatFloat(res.Freq, 'f', -1, 64)})
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to write output:", err)
			os.Exit(1)
		}
	}
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to write output:", err)
		os.Exit(1)
	}
}
