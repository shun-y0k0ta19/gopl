// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/rm", db.rm)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	tmpl := `
	<table>
	{{range $item, $price := .}}
          <tr>
            <td>{{$item}}</td>
            <td>{{$price}}</td>
          </tr>
        {{end}}
    </table>
	`
	t := template.New("table")
	//t.Parse(tmpl)
	template.Must(t.Parse(tmpl))
	fmt.Println("template is must")
	t.Execute(w, db)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	} else {
		pstr := req.URL.Query().Get("price")
		price, err := strconv.ParseFloat(pstr, 32)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "price is invalid: %v\n", pstr)
		}
		db[item] = dollars(price)
		fmt.Fprintf(w, "%f\n", price)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "already exist: %q\n", item)
	} else {
		pstr := req.URL.Query().Get("price")
		price, err := strconv.ParseFloat(pstr, 32)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "price is invalid: %v\n", pstr)
		}
		db[item] = dollars(price)
		fmt.Fprintf(w, "%f\n", price)
	}
}

func (db database) rm(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	} else {
		delete(db, item)
		fmt.Fprintf(w, "delete %v\n", item)
	}
}

func sprintDataBase(db database) string {
	const (
		tbl = "<table>\n%v</table>"
		tr  = "<tr>\n%v</tr>\n"
		td  = "<td>%v</td>\n<td>%v</td>\n"
	)
	var items string
	for item, price := range db {
		tds := fmt.Sprintf(td, item, price)
		trs := fmt.Sprintf(tr, tds)
		items += trs
	}
	table := fmt.Sprintf(tbl, items)
	return table
}
