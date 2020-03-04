package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("primitive", strings.Fields("-i tmp/mtn.jpg -o mtest.png -n 50 -m 6")...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}