package main

import (
	"fmt"
	"time"
)

func TimeModule() {
	date := time.Now()

	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
	fmt.Println(date.Hour())
	fmt.Println(date.Minute())
	fmt.Println(date.Second())
	fmt.Println(date.Nanosecond())

	utc := time.Date(2022, 1, 1, 0, 0, 12, 12, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02" // * layout adalah bawaan dari package time
	parse, _ := time.Parse(layout, "2021-11-18")
	fmt.Println(parse)

	local := date.Local()
	fmt.Println(local)

	utcLocal := time.Date(2009, time.July, 30, 10, 11, 0, 0, time.UTC)
	fmt.Println(utcLocal)

	parseLocal, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(parseLocal)
}
