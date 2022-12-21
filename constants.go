package main

import (
	"fmt"
)

const s string = "constant"

func main() {
	fmt.Println(s)
	const n = 20
	const d = n / 4
	fmt.Println(d)
}
