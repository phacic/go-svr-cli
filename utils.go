package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

// FatalOnErr logs error with log.Fatal
func FatalOnErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// PanicOnErr logs error with log.Panic
func PanicOnErr(e error) {
	if e != nil {
		log.Panic(e)
	}
}

// Logger is for logging across the app
var Logger = &logrus.Logger{
	Out:       os.Stdout,
	Level:     logrus.InfoLevel,
	Formatter: &logrus.TextFormatter{},
}
