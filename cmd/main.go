package main

import (
	"fmt"
	"github.com/KuldeepNITS09/cronparser/parser"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cron-parser \"<cron-string>\"")
		return
	}

	cronString := os.Args[1]
	parser := &parser.DefaultCronParser{}
	cron, err := parser.Parse(cronString)
	if err != nil {
		fmt.Println("Error parsing cron string:", err)
		return
	}

	fmt.Printf("%-14s %s\n", "minute", cron.Minute)
	fmt.Printf("%-14s %s\n", "hour", cron.Hour)
	fmt.Printf("%-14s %s\n", "day of month", cron.DayOfMonth)
	fmt.Printf("%-14s %s\n", "month", cron.Month)
	fmt.Printf("%-14s %s\n", "day of week", cron.DayOfWeek)
	fmt.Printf("%-14s %s\n", "command", cron.Command)
}
