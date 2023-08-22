// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http3a is an e-commerce server that registers the /list and /price
// endpoints by calling (*http.ServeMux).HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"text/template"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	//!+main
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/listHtml", db.listHtml)
	//!-main
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]int

var dbmutex sync.Mutex

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: $%d\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "$%d\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if _, ok := db[item]; ok {
		if nprice, nerr := strconv.Atoi(price); nerr != nil || nprice <= 0 {
			w.WriteHeader(http.StatusNotAcceptable) // 406
			fmt.Fprintf(w, "price of %q is not allowed: %q\n", item, price)
		} else {
			dbmutex.Lock()
			fmt.Fprintf(w, "update price of %q from $%d to $%d\n", item, db[item], nprice)
			db[item] = nprice
			dbmutex.Unlock()
		}
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

var dbTable = template.Must(template.New("dbTable").Parse(`
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
		<h1>Items</h1>
		<table>
		  <thead>
				<tr>
					<th>Name</th>
					<th>Price</th>
				</tr>
			</thead>
			<tbody>
				{{range .}}
				<tr>
					<td>{{.Name}}</td>
					<td>{{.Price}}</td>
				</tr>
				{{end}}
			</tbody>
		</table>
	</body>
</html>
`))

type databaseStruct struct {
	Name  string
	Price string
}

func (db database) listHtml(w http.ResponseWriter, req *http.Request) {
	var ds []*databaseStruct
	var names []string
	for item := range db {
		names = append(names, item)
	}
	sort.Strings(names)
	for _, item := range names {
		ds = append(ds, &databaseStruct{item, strconv.Itoa(db[item])})
	}
	if err := dbTable.Execute(w, ds); err != nil {
		log.Fatal(err)
	}
}

/*
//!+handlerfunc
package http

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
//!-handlerfunc
*/
