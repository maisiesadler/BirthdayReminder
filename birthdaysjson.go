package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

type Birthday2 struct {
	Name     string
	Birthday string
}

func readJson() []Birthday2 {
	var birthdays = []Birthday2{}
	byt, _ := ioutil.ReadFile("/birthday/birthdays.json")
	// byt := []byte(b)
	json.Unmarshal(byt, &birthdays)
	fmt.Printf("%v", birthdays)

	return birthdays
}

func writeJson() {
	var birthdays = []Birthday2{}
	b := Birthday2{Name: "Maisie", Birthday: "2802"}
	birthdays = append(birthdays, b)
	byt, _ := json.Marshal(birthdays)
	ioutil.WriteFile("./birthdays.json", byt, 0644)
}

func getBirthdaysInDays(birthdays []Birthday2, days int) []Birthday2 {
	t := time.Now()
	for d := 0; d < days; d++ {
		t = t.Add(time.Hour * 24)
	}
	todays := []Birthday2{}
	tMonth := int(t.Month())
	tDay := t.Day()
	for _, birthday := range birthdays {
		bDay, _ := strconv.Atoi(birthday.Birthday[0:2])
		bMonth, _ := strconv.Atoi(birthday.Birthday[2:4])
		if bMonth == tMonth && bDay == tDay {
			todays = append(todays, birthday)
		}
	}
	return todays
}

func getBirthdayReminders() []string {
	reminders := []string{}

	birthdays := readJson()
	oneWeek := getBirthdaysInDays(birthdays, 6)
	tomorrows := getBirthdaysInDays(birthdays, 1)
	todays := getBirthdaysInDays(birthdays, 0)

	for _, person := range oneWeek {
		reminders = append(reminders, "It's "+person.Name+"'s birthday in one week!")
	}
	for _, person := range tomorrows {
		reminders = append(reminders, "It's "+person.Name+"'s birthday tomorrow!")
	}
	for _, person := range todays {
		reminders = append(reminders, "It's "+person.Name+"'s birthday today!")
	}
	return reminders
}
