package main

import (
	"fmt"
	"github.com/hydra13142/regex"
)

func main() {
	reg := regex.Compile(`(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})(?::(\d{1,5}))?`)
	for _, s := range reg.FindAllString("127.0.0.1:8087 255.33.1.2:12345  217.169.209.2:6666  192.227.139.106:7808") {
		fmt.Printf("%-24s\n", s[0])
	}
}
