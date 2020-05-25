package main

import (
	"time"
)

/*var (
	Debug       bool      = false
	LogLevel    string    = "info"
	startUpTime time.Time = time.Now()
)*/

// note the difference from above statements
/*
var (
	Debug       bool
	LogLevel    = "info"
	startUpTime = time.Now()
)*/

func main() {
	//// note immediete multiple declaration
	//Debug, LogLevel, startUpTime := false, "info", time.Now()

	//Debug, LogLevel, startUpTime := getConfig() // from method, note hoe := is used
	//fmt.Println(Debug, LogLevel, startUpTime)

	//// Type only
	// var start, middle, end float32
	// fmt.Println(start, middle, end)
	// // Initial value mixed type
	// var name, left, right, top, bottom = "one", 1, 1.5, 2, 2.5
	// fmt.Println(name, left, right, top, bottom)
	// // works with functions also
	// var Debug, LogLevel, startUpTime = getConfig()
	// fmt.Println(Debug, LogLevel, startUpTime)

}

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}
