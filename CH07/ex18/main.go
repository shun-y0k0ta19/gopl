// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

//Node is node interface
type Node interface{}

//CharData is character data interface
type CharData string

//Element is element interface
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	root := Element{Type: xml.Name{Space: "root", Local: "root"}, Attr: *new([]xml.Attr), Children: *new([]Node)}
	stack := []Element{root}
	for {
		//for i := 0; i < 30; i++ {
		tok, err := dec.Token()
		if err == io.EOF {
			fmt.Println(stack[0].Type.Local)
			//root = stack[0].Children[0].(Element)
			//root = stack[0]
			fmt.Println("Parse done!")
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		last := len(stack) - 1
		switch tok := tok.(type) {
		case xml.StartElement:
			//fmt.Printf("len(stack): %d\n", len(stack))
			//fmt.Println(tok.Name.Local)
			child := Element{Type: tok.Name, Attr: tok.Attr, Children: *new([]Node)}
			//	fmt.Printf("child: %s\n", child)
			//	fmt.Printf("stack[%d].children: %s\n", last, stack[last].Children)
			stack = append(stack, child)
			//fmt.Printf("sstack[0].c[0]: %v\n", stack[0].Children[0])
			//fmt.Println("stack[0]:")
			showChild(stack[0])
		case xml.EndElement:
			stack[last-1].Children = append(stack[last-1].Children, stack[last])
			stack = stack[:last]
			root = stack[len(stack)-1]
		case xml.CharData:
			//fmt.Println(string(tok.Copy()))
			stack[last].Children = append(stack[last].Children, tok.Copy())
		}
		fmt.Println()
	}
	fmt.Println("\nParse done!")
	showChild(root)
	//	showChild(stack[0])
}

func showChild(node Node) {
	switch nn := node.(type) {
	case Element:
		fmt.Printf("<%s>\n", nn.Type.Local)
		for _, c := range nn.Children {
			showChild(c)
		}
	case CharData:
		fmt.Printf("[%s]\n", nn)
	}
}
