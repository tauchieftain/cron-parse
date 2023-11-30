package main

import (
	"cronParse/cron"
	"fmt"
	"time"
)

func main() {
	expr := "0 5 * * * ?"
	p := cron.NewParse(expr)
	fmt.Println(p.NextExecTime(time.Now()))
}
