package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var sFlag = flag.String("s", "", "movie title to search for")
var iFlag = flag.String("i", "", "a valid IMDb ID (e.g. tt1285016)")
var tFlag = flag.String("t", "", "movie title")
var dFlag = flag.Bool("d", false, "download poster")

const (
	apiURL                = "https://www.omdbapi.com/?apikey=32b78491"
	posterDBDir           = "poster.db"
	posterFilePathPattern = posterDBDir + "/poster-%s.jpg"
)

type SearchResult struct {
	TotalResults string
	Response     string
	Search       []*SearchItem
}

type SearchItem struct {
	Title  string
	Year   string
	ImdbID string
	Type   string
	Poster string
}

type QueryResult struct {
	Title    string
	Year     string
	ImdbID   string
	Type     string
	Poster   string
	Response string
}

func main() {
	flag.Parse()
	if *sFlag != "" {
		page := 1
		for {
			sr, err := getSearchResult(*sFlag, page)
			if err != nil {
				log.Fatalf("getSearchResult failed: %v\n", err)
			}
			data, err := json.MarshalIndent(sr, "", "    ")
			if err != nil {
				log.Fatalf("JSON marshaling failed: %v\n", err)
			}
			fmt.Printf("%s\n", data)
			if totalResults, err := strconv.Atoi(sr.TotalResults); err == nil && totalResults > page*10 {
				fmt.Printf("This is page %d of %d. Proceed to the next page? [Y/N] ", page, (totalResults+9)/10)
				input := bufio.NewScanner(os.Stdin)
				if input.Scan() {
					text := input.Text()
					if text == "N" || text == "n" {
						break
					} else {
						page++
						continue
					}
				} else {
					break
				}
			} else {
				break
			}
		}
	} else if *iFlag != "" || *tFlag != "" {
		qr, err := getQueryResult(*iFlag, *tFlag)
		if err != nil {
			log.Fatalf("getQueryResult failed: %v\n", err)
		}
		data, err := json.MarshalIndent(qr, "", "    ")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %v\n", err)
		}
		fmt.Printf("%s\n", data)
		if *dFlag == true && qr.ImdbID != "" && qr.Poster != "" {
			fmt.Printf("Downloading poster to file %s...\n", posterFilePath(qr.ImdbID))
			if err := downloadPoster(qr.Poster, qr.ImdbID); err != nil {
				log.Fatalf("downloadPoster failed: %v\n", err)
			} else {
				fmt.Printf("Done.\n")
			}
		}
	} else {
		// .\e4.13.exe -s Guardians
		// .\e4.13.exe -i tt2015381 -d
		// .\e4.13.exe -t "Guardians of the Galaxy Vol. 2" -d
		fmt.Fprintln(os.Stderr, "usage: poster -s \"movie title to search for\" -t \"movie title\" -i \"a valid IMDb ID (e.g. tt1285016)\" -d")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func getSearchResult(s string, page int) (*SearchResult, error) {
	searchURL := fmt.Sprintf("%s&s=%s&page=%d", apiURL, url.QueryEscape(s), page)

	resp, err := http.Get(searchURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Get Failed: %s", resp.Status)
	}

	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var sr SearchResult
	err = json.Unmarshal(jsonBytes, &sr)
	if err != nil {
		return nil, err
	}

	return &sr, nil
}

func getQueryResult(i string, t string) (*QueryResult, error) {
	queryURL := apiURL
	if i != "" {
		queryURL = queryURL + "&i=" + url.QueryEscape(i)
	}
	if t != "" {
		queryURL = queryURL + "&t=" + url.QueryEscape(t)
	}

	resp, err := http.Get(queryURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Get Failed: %s", resp.Status)
	}

	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var qr QueryResult
	err = json.Unmarshal(jsonBytes, &qr)
	if err != nil {
		return nil, err
	}

	return &qr, nil
}

func downloadPoster(url string, imdbID string) error {
	if exists(posterFilePath(imdbID)) {
		return nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Get Failed: %s", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	createDBDirectoryIfNecessary()
	err = ioutil.WriteFile(posterFilePath(imdbID), data, 0666)
	if err != nil {
		return err
	}

	return nil
}

func exists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(fmt.Errorf("os.Stat: %v", err))
	}
	return true
}

func posterFilePath(imdbID string) string {
	return fmt.Sprintf(posterFilePathPattern, imdbID)
}

func createDBDirectoryIfNecessary() {
	fileInfo, err := os.Stat(posterDBDir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(posterDBDir, 0777)
			if err != nil {
				panic(err)
			}
			return
		}
	}

	if !fileInfo.IsDir() {
		panic(fmt.Errorf("%s is not directory", posterDBDir))
	}
}
