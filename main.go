package main

import (
	"fmt"
)

func main() {
	type Name struct {
		s string
	}
	a := []*Name{
		&Name{s: "taha"},
		&Name{s: "neo"},
		&Name{s: "Nilo"},
	}
	v := struct {
		Ame []*Name
	}{
		Ame: a,
	}
	for _, name := range v.Ame {
		fmt.Println(name)
		name.s = "changed"
	}

	for _, name := range v.Ame {
		fmt.Println(name)
	}
}
