// Copyright © 2016 "Shun Yokota" All rights reserved

package sexpr

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	/*
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
	*/
	// Encode it
	//data, err := Marshal(strangelove)
	data, err := Marshal(0.0123)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)
	fmt.Printf("Marshal() = %s\n", data)

	data, err = Marshal(5 + 3.3i)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)
	fmt.Printf("Marshal() = %s\n", data)

	data, err = Marshal(make(chan<- string))
	if err == nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", err)
	fmt.Printf("Marshal() = %s\n", err)

	f := encode
	data, err = Marshal(f)
	if err == nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", err)
	fmt.Printf("Marshal() = %s\n", err)

	type hogeif interface {
	}
	type ifStruct struct {
		hogeif
	}
	type intS []int
	hg := ifStruct{intS{1, 2, 3}}
	data, err = Marshal(hg)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)
	fmt.Printf("Marshal() = %s\n", data)

}
