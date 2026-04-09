package main

import (
	"flag"
	"log"
	"strings"
)

func main() {
	src := flag.String("src", "https://raw.githubusercontent.com/tdlib/td/refs/heads/master/td/generate/scheme/td_api.tl", "Path to TL file or URL")
	version := flag.String("version", "", "TDLib version")
	commit := flag.String("commit", "", "TDLib commit hash")
	out := flag.String("out", "tdlib.json", "Output JSON file path")
	flag.Parse()

	var data *TDLibJSON
	var err error

	if strings.HasPrefix(*src, "https://") {
		log.Printf("Fetching TL schema from %s...", *src)
		data, err = FetchAndParseTL(*src)
	} else {
		log.Printf("Reading TL schema from %s...", *src)
		data, err = ParseTLFromFile(*src)
	}

	if err != nil {
		log.Fatalf("Failed to parse TL: %v", err)
	}

	data.Version = *version
	data.Commit = *commit

	log.Println("TDLib options...")
	options, err := getOptions()
	if err != nil {
		log.Printf("Warning: Failed to get options: %v", err)
	} else {
		data.Options = options
		log.Printf("Fetched %d options.", len(options))
	}

	log.Printf("Saving JSON to %s...", *out)
	if err := SaveTDLibJSON(data, *out); err != nil {
		log.Fatalf("Failed to save JSON: %v", err)
	}
	log.Println("Done.")
}
