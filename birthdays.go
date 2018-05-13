package birthday_reminder

import (
	"time"
)

type Birthday struct {
	Name string
	Date time.Time
}

// var birthdays = []Birthday{}

// func loadBirthdays() {
// 	b, _ := persist.Get("birthdays")
// 	byt := []byte(b)
// 	json.Unmarshal(byt, &birthdays)
// }

// func addBirthday(name string, day int, month int) {
// 	t := time.Now()
// 	date := time.Date(0, time.Month(month), day, 0, 0, 0, 0, t.Location())

// 	b := Birthday{Name: name, Date: date}
// 	birthdays = append(birthdays, b)
// 	byt, _ := json.Marshal(birthdays)
// 	persist.Add("birthdays", string(byt))
// }

// func printBirthdays() {
// 	byt, _ := json.Marshal(birthdays)
// 	fmt.Printf("birthdays: %+v", string(byt))
// }
// func printBirthdaysInDays(days int) {
// 	tB := birthdaysInDays(days)
// 	byt, _ := json.Marshal(tB)
// 	fmt.Printf("todays birthdays: %+v", string(byt))
// }

// func birthdaysInDays(days int) []Birthday {
// 	t := time.Now()
// 	for d := 0; d < days; d++ {
// 		t = t.Add(time.Hour * 24)
// 	}
// 	todays := []Birthday{}
// 	tMonth := t.Month()
// 	tDay := t.Day()
// 	for _, birthday := range birthdays {
// 		bMonth := birthday.Date.Month()
// 		bDay := birthday.Date.Day()
// 		if bMonth == tMonth && bDay == tDay {
// 			todays = append(todays, birthday)
// 		}
// 	}
// 	return todays
// }
