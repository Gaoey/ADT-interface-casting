package main

import (
	"encoding/json"
	"fmt"
)

type A1 struct {
	X string      `json:"a"`
	Y int         `json:"b"`
	Z interface{} `json:"c"`
}

func (a A1) Casting(bb interface{}) Origin {
	// convert map to json
	jsonString, _ := json.Marshal(a.Z)
	fmt.Println(string(jsonString))

	// convert json to B1
	_, ok := bb.(B1)
	if ok {
		s := B1{}
		json.Unmarshal(jsonString, &s)
		return &s
	}

	// convert json to C1
	_, ok = bb.(C1)
	if ok {
		s := C1{}
		json.Unmarshal(jsonString, &s)
		return &s
	}

	return nil
}

type Origin interface {
	Shit() Origin
}

type B1 struct {
	X string `json:"a"`
	Y int    `json:"b"`
}

func (b B1) Shit() Origin {
	return b
}

type C1 struct {
	X1 string `json:"a1"`
	Y1 int    `json:"b2"`
	Z1 string `json:"c3"`
}

func (c C1) Shit() Origin {
	return c
}

func main() {
	tst1 := B1{X: "222", Y: 2}
	tst2 := A1{X: "111", Y: 1, Z: tst1}
	// Boom publish

	// Yeah .. Consume
	res, _ := json.Marshal(tst2)

	var e A1
	if err := json.Unmarshal(res, &e); err != nil {
		fmt.Printf("%#v\n", err)
		return
	}

	s := e.Casting(B1{})

	fmt.Printf("%#v\n", s)
}
