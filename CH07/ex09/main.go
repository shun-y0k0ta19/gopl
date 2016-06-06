// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"
)

//Track is information of a song in CD
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

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!+printTracks
func printTracks(tracks []*Track) {
	fmt.Println(sprintTracks(tracks))
}

//!-printTracks

//!+sprintTracks
func sprintTracks(tracks []*Track) string {
	const (
		tbl = "<table>\n%v</table>"
		tr  = "<tr>\n%v</tr>\n"
		td  = "<td>%v</td>\n<td>%v</td>\n<td>%v</td>\n<td>%v</td>\n<td>%v</td>\n"
	)
	headtd := fmt.Sprintf(td, "Title", "Artist", "Album", "Year", "Length")
	headtr := fmt.Sprintf(tr, headtd)
	vals := headtr
	for _, t := range tracks {
		tds := fmt.Sprintf(td, t.Title, t.Artist, t.Album, t.Year, t.Length)
		trs := fmt.Sprintf(tr, tds)
		vals += trs
	}
	table := fmt.Sprintf(tbl, vals)
	return table
}

//!-sprintTracks

//!+artistcode
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-artistcode

//!+albumcode
type byAlbum []*Track

func (x byAlbum) Len() int           { return len(x) }
func (x byAlbum) Less(i, j int) bool { return x[i].Album < x[j].Album }
func (x byAlbum) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-albumcode

//!+yearcode
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-yearcode

//!+titlecode
type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-titlecode

//!+lengthcode
type byLength []*Track

func (x byLength) Len() int           { return len(x) }
func (x byLength) Less(i, j int) bool { return x[i].Length < x[j].Length }
func (x byLength) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-lengthcode

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	switch query.Get("sortkey") {
	case "artist":
		fmt.Println("byArtist:")
		sort.Sort(byArtist(tracks))
	case "title":
		fmt.Println("byTitle:")
		sort.Sort(byTitle(tracks))
	case "album":
		fmt.Println("byAlbum:")
		sort.Sort(byAlbum(tracks))
	case "year":
		fmt.Println("byYear:")
		sort.Sort(byYear(tracks))
	case "length":
		fmt.Println("byLength:")
		sort.Sort(byLength(tracks))

	}
	//fmt.Println("byArtist:")
	//sort.Sort(byArtist(tracks))
	printTracks(tracks)
	fmt.Fprint(w, sprintTracks(tracks))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
