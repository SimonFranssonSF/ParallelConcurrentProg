package main

import (
	"fmt"
	"time"
)

//Function that prints a message about current time and a text as often as the inputted delay allows.
func Remind(text string, delay time.Duration) {
	for {
		time.Sleep(delay)
		now := time.Now()
		nowFormatted := now.Format("15:04")
		fmt.Println("Klockan är", nowFormatted, "+", text)
	}

}

func main() {
	//Three goroutines with different delay and textstrings sent to function Remind
	go Remind("Dags att äta", 3*time.Hour)
	go Remind("Dags att arbeta", 8*time.Hour)
	go Remind("Dags att sova", 24*time.Hour)

	//Prevents the main from exiting early
	select {}
}
