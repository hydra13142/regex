package main

import (
	"fmt"
	"github.com/hydra13142/regex"
)

func main() {
	reg := regex.Compile(`\s+|[-+*/]|\d+(?:\.\d+)?(?:[eE][-+]?\d+)?|\(#{0}+?\)`)
	for _, u := range reg.FindAllStringSubmatch("1+(2-3)*(4/5.33)*((6-7)+8)") {
		for _, v := range u {
			fmt.Print(v, " ")
		}
		fmt.Println("")
	}
}
