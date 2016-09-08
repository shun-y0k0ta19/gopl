// Copyright © 2016 "Shun Yokota" All rights reserved

package sexpr

import (
	"fmt"
	"io"
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
	//fmt.Printf("Marshal() = %s\n", data)

	dec := NewDecoder(bytes.NewBuffer(data))
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		switch tok.(type) {
		case StartList:
			t.Logf("(")
			fmt.Print("(")
		case EndList:
			t.Logf(")")
			fmt.Print(")")
		case Int:
			t.Logf("%s ", tok.(Int).str)
			fmt.Printf("%s ", tok.(Int).str)
		case Symbol:
			t.Logf("%s ", tok.(Symbol))
			fmt.Printf("%s ", tok.(Symbol))
		case String:
			t.Logf("%s ", tok.(String))
			fmt.Printf("%s ", tok.(String))
		}
	}
}
