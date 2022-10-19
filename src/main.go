package main

import (
	"learnGoWithTests/src/mocking"
	"os"
	"time"
)

func main() {
	//CountDown Mocking
	sleeper := &mocking.ConfigurableSleeper{1 * time.Second, time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
