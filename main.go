package main

import (
	"fmt"

	"date-calculation/date"
)

func main() {
	y := date.NewYear(2022)
	forthm := y.GetMonth(4)

	fmt.Printf("%s, %s\n", y, forthm)
}
