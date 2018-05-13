package main

import (
	"time"
)

func initNoon() chan bool {
	t := time.Now()
	n := time.Date(t.Year(), t.Month(), t.Day(), 12, 0, 0, 0, t.Location())
	d := n.Sub(t)
	if d < 0 {
		n = n.Add(24 * time.Hour)
		d = n.Sub(t)
	}
	task := make(chan bool)
	go beginTask(d, 24*time.Hour, task)
	return task
}

func initEveryMinute() chan bool {
	t := time.Now()
	n := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()+1, 0, 0, t.Location())
	d := n.Sub(t)
	task := make(chan bool)
	go beginTask(d, time.Minute, task)
	return task
}

func beginTask(timeTilFirst time.Duration, period time.Duration, task chan bool) {
	d := timeTilFirst
	for {
		time.Sleep(d)
		d = period
		task <- true
	}
}
