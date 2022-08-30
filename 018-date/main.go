package main

import (
	"fmt"
	"log"

	"github.com/tyatyushkin/calendar"
)

func main() {
	event := calendar.Event{}
	err := event.SetTitle("Mom's birthday")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2023)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(7)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}
