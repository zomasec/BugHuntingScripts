package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

func countParameters(u *url.URL) int {
	return len(u.Query())
}

func filterURLs(urls []*url.URL) []*url.URL {
	countMap := make(map[int][]*url.URL)
	var maxCount int

	for _, u := range urls {
		count := countParameters(u)
		countMap[count] = append(countMap[count], u)
		if count > maxCount {
			maxCount = count
		}
	}

	return countMap[maxCount]
}

func parseURLs(input string) ([]*url.URL, error) {
	lines := strings.Split(input, "\n")
	var urls []*url.URL

	for _, line := range lines {
		if line != "" {
			parsedURL, err := url.Parse(line)
			if err != nil {
				return nil, err
			}
			urls = append(urls, parsedURL)
		}
	}

	return urls, nil
}

func main() {
	inputFile := flag.String("i", "", "Input file containing URLs")
	outputFile := flag.String("o", "", "Output file to save filtered URLs")
	flag.Parse()

	var input string

	if *inputFile != "" {
		content, err := ioutil.ReadFile(*inputFile)
		if err != nil {
			fmt.Println("Error reading input file:", err)
			return
		}
		input = string(content)
	} else {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			data, _ := ioutil.ReadAll(os.Stdin)
			input = string(data)
		}
	}

	urls, err := parseURLs(input)
	if err != nil {
		fmt.Println("Error parsing URLs:", err)
		return
	}

	filteredURLs := filterURLs(urls)

	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer file.Close()

		for _, u := range filteredURLs {
			file.WriteString(u.String() + "\n")
		}
	} else {
		for _, u := range filteredURLs {
			fmt.Println(u)
		}
	}
}
