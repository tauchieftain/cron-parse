package main

import (
	cronParse "cronParse/src"
	"fmt"
	"time"
)

func main() {
	expr := "0 5 * * * ?"
	p := cronParse.NewParse(expr)
	fmt.Println(p.NextExecTime(time.Now()))
}
