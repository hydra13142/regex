package main

import (
	"fmt"
	"github.com/hydra13142/regex"
)

func main() {
	reg := regex.Compile(`(?<!\d)(?:0|1\d{0,2}|2(?:[0-4]\d?|5[0-5]?|[6-9])?|[3-9]\d?)\.#{1}\.#{1}\.#{1}(?::\d{1,5})?`)
	for _, s := range reg.FindAllStringSubmatch("127.0.0.1:8087 255.33.1.2:12345  217.169.209.2:6666  192.227.139.106:7808") {
		fmt.Printf("%-24s", s[0])
		for _, x := range s[1:] {
			fmt.Printf("%-8s", x)
		}
		fmt.Println("")
	}
}
