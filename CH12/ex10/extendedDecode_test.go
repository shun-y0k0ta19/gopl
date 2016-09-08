// Copyright © 2016 "Shun Yokota" All rights reserved

package sexpr

import (
	"fmt"
	"testing"

	"golang_training/CH12/ex03"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func init() {
	var movie Movie
	RegisterType(movie)
	RegisterType(movie.Actor)
	RegisterType(movie.Oscars)
	RegisterType(movie.Sequel)
	RegisterType(movie.Subtitle)
	RegisterType(movie.Title)
	RegisterType(movie.Year)
}

func TestDecode(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	fmt.Printf("------------------encode and decode Movie:Start--------------------\n")
	// Encode it
	data, err := sexpr.Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal():\n %s\n\n", data)
	fmt.Printf("Marshal(): \n %s\n\n", data)

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal(): \n %+v\n\n", movie)
	fmt.Printf("Unmarshal():\n %+v\n", movie)
	fmt.Printf("------------------encode and decode Movie:End--------------------\n")

	fmt.Printf("\n\n")
}

type testStruct struct {
	F         float64
	B         bool
	Interface interface{}
}

func TestExtendedDecode(t *testing.T) {
	ts := testStruct{F: 0.01, B: true, Interface: []int{1, 2, 3}}
	fmt.Printf("------------------encode and decode ts:Start--------------------\n")
	testExtendedDecode(ts, t)
	fmt.Printf("------------------encode and decode ts:End--------------------\n")
	fmt.Printf("\n\n")

	ts = testStruct{F: 0.01, B: true, Interface: testStruct{F: 0.3, B: false, Interface: 5}}
	RegisterType(5)
	fmt.Printf("------------------encode and decode ts:Start--------------------\n")
	testExtendedDecode(ts, t)
	fmt.Printf("------------------encode and decode ts:End--------------------\n")
	fmt.Printf("\n\n")

}

func testExtendedDecode(ts testStruct, t *testing.T) {
	// Register types
	RegisterType(ts.B)
	RegisterType(ts.F)
	RegisterType(ts.Interface)

	data, err := sexpr.Marshal(ts)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)
	fmt.Printf("Marshal() = %s\n", data)
	var result testStruct
	if err := Unmarshal(data, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", result)
	fmt.Printf("Unmarshal() = %+v\n", result)

	tsm, err := sexpr.Marshal(ts)
	if err != nil {
		t.Error(err)
	}
	resm, err := sexpr.Marshal(result)
	if err != nil {
		t.Error(err)
	}

	if string(tsm) != string(resm) {
		t.Errorf("want: %v\nresult: %v\n", ts, result)
	}
}

/*
func Test(t *testing.T) {
	data := "(string:fasf, sfsd.\nasn = 100)"
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader([]byte(data)))
	lexShow(lex)
	fmt.Printf("text: %s\n", lex.text())
	fmt.Printf("text: %s\n", lex.text())
	for i := 0; i < 13; i++ {
		lexShow(lex)
	}

	const src = `{"Hoge1":[1,2,3],"Hoge2":["a","b","c"]}`
	lex.scan.Init(bytes.NewReader([]byte(src)))
	lexShow(lex)
	fmt.Printf("text: %s\n", lex.text())
	fmt.Printf("text: %s\n", lex.text())
	for i := 0; i < 13; i++ {
		lexShow(lex)
	}

}

func lexShow(lex *lexer) {
	lex.next()
	fmt.Printf("token: '%[1]c',%[1]d\n", lex.token)
	if lex.token == scanner.Ident {
		fmt.Println("Ident dayo")
	}
	fmt.Printf("text: %s\n", lex.text())
}
*/
