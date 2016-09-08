// Copyright © 2016 "Shun Yokota" All rights reserved

package sexpr

import (
	"fmt"
	"testing"

	"bytes"

	"gopl.io/ch12/sexpr"
)

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

	// Encode it
	data, err := sexpr.Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)
	fmt.Printf("Marshal() = %s\n", data)

	dec := NewDecoder(bytes.NewBuffer(data))
	// Decode it
	var movie Movie
	if err := Unmarshal(dec, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)
	fmt.Printf("Unmarshal() = %+v\n", movie)

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
