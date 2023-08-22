// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
	"time"
)

// !+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

var trackTable = template.Must(template.New("trackTable").Parse(`
<!DOCTYPE html>
<html lang="en">
	<head>
	<meta charset="utf-8">
	<style type="text/css">
		table {
			border-collapse: collapse;
			border-spacing: 0px;
		}
		table, th, td {
			padding: 5px;
			border: 1px solid black;
		}
	</style>
	</head>
	<body>
		<h1>Tracks</h1>
		<table>
		  <thead>
				<tr>
					<th><a href="/?sort=Title">Title</a></th>
					<th><a href="/?sort=Year">Year</a></th>
					<th><a href="/?sort=Length">Length</a></th>
					<th><a href="/?sort=Artist">Artist</a></th>
					<th><a href="/?sort=Album">Album</a></th>
				</tr>
			</thead>
			<tbody>
				{{range .}}
				<tr>
					<td>{{.Title}}</td>
					<td>{{.Year}}</td>
					<td>{{.Length}}</td>
					<td>{{.Artist}}</td>
					<td>{{.Album}}</td>
				</tr>
				{{end}}
			</tbody>
		</table>
	</body>
</html>
`))

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.RequestURI == "/favicon.ico" {
		return
	}
	sortBy := request.FormValue("sort")
	recentColumns := []string{sortBy}
	fmt.Println("Custom Recently Clicked Columns:", recentColumns)
	//!+customcall
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		for _, recentColumn := range recentColumns {
			switch recentColumn {
			case "Title":
				if x.Title != y.Title {
					return x.Title < y.Title
				}
			case "Year":
				if x.Year != y.Year {
					return x.Year < y.Year
				}
			case "Length":
				if x.Length != y.Length {
					return x.Length < y.Length
				}
			case "Artist":
				if x.Artist != y.Artist {
					return x.Artist < y.Artist
				}
			case "Album":
				if x.Album != y.Album {
					return x.Album < y.Album
				}
			}
		}
		return false
	}})
	//!-customcall
	printTracks(responseWriter, tracks)
}

func printTracks(writer io.Writer, tracks []*Track) {
	if err := trackTable.Execute(writer, tracks); err != nil {
		log.Fatal(err)
	}
}

// !+customcode
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

//!-customcode
