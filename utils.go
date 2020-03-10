package main

import (
	"log"
)

func Debug(s string, a ...interface{}) {
	if debug {
		log.Printf(s, a...)
	}
}
