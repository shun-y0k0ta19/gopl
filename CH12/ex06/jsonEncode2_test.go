// Copyright © 2016 "Shun Yokota" All rights reserved

package sexpr

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"testing"
)

func Test(t *testing.T) {

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
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	jdata, err := json.Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	fmt.Printf("Marshal() = \n%s\n", data)
	fmt.Printf("json.Marshal() = \n%s\n", jdata)
	if string(data) != string(jdata) {
		//t.Errorf("data: %s\njdata: %s", string(data), string(jdata))
	}

	testMarshal(0.012, t)
	testMarshal(1, t)
	testMarshal(true, t)
	testMarshal(false, t)
	testMarshal("test", t)
	i := 1
	testMarshal(&i, t)
	testMarshal([3]int{1, 2, 3}, t)
	testMarshal([]int{1, 2, 3}, t)
	testMarshal([]string{"1", "2", "3"}, t)
	type stru struct {
		I int
		S string
	}
	testMarshal(stru{1, "test"}, t)
	testMarshalMap(map[string]string{
		"Dr. Strangelove":            "Peter Sellers",
		"Grp. Capt. Lionel Mandrake": "Peter Sellers",
		"Pres. Merkin Muffley":       "Peter Sellers",
		"Gen. Buck Turgidson":        "George C. Scott",
		"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
		`Maj. T.J. "King" Kong`:      "Slim Pickens",
	}, t)

	type hogeif interface {
	}
	type ifStruct struct {
		Hoge1 hogeif
		Hoge2 hogeif
	}
	testMarshal(ifStruct{[]int{1, 2, 3}, []string{"a", "b", "c"}}, t)

	testMarshal(0, t)

	type zeroStruct struct {
		I   int
		F   float32
		B   bool
		S   string
		Ptr *int
		Arr [5]int
		Sli []string
		Ifs ifStruct
		M   map[string]int
		Hif hogeif
	}
	testMarshalZeroVal(zeroStruct{}, t)

}

func testMarshal(arg interface{}, t *testing.T) {
	data, err := Marshal(arg)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	jdata, err := json.Marshal(arg)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	fmt.Printf("     Marshal() = %s\n", data)
	fmt.Printf("json.Marshal() = %s\n", jdata)
	if string(data) != string(jdata) {
		t.Errorf("\n data: %s\njdata: %s", string(data), string(jdata))
	}
}

func testMarshalMap(arg interface{}, t *testing.T) {
	data, err := Marshal(arg)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	jdata, err := json.Marshal(arg)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	fmt.Printf("     Marshal() = %s\n", data)
	fmt.Printf("json.Marshal() = %s\n", jdata)
	sd := string(data)[1 : len(string(data))-1]
	sjd := string(jdata)[1 : len(string(jdata))-1]
	sds := strings.Split(sd, ",")
	sjds := strings.Split(sjd, ",")
	sort.Strings(sds)
	sort.Strings(sjds)
	for i, s := range sds {
		if s != sjds[i] {
			t.Errorf("\n data: %s\njdata: %s", string(data), string(jdata))
			return
		}
	}
}

func testMarshalZeroVal(arg interface{}, t *testing.T) {
	data, err := Marshal(arg)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	fmt.Printf("     Marshal() = %s\n", data)
	if string(data) != "{}" {
		t.Errorf("\n data: %s\n", string(data))
	}
}
