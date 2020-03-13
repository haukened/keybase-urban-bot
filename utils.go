package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"strconv"
)

// Debug provides printing only when --debug flag is set or BOT_DEBUG env var is set
func Debug(s string, a ...interface{}) {
	if debug {
		log.Printf(s, a...)
	}
}

// parseArgs parses command line and environment args and sets globals
func parseArgs(args []string) error {
	// first check for command line flags
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.BoolVar(&debug, "debug", false, "enables command debugging")
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	// then check the env vars
	envDebug := os.Getenv("BOT_DEBUG")
	if envDebug != "" {
		ret, err := strconv.ParseBool(envDebug)
		if err != nil {
			return err
		}

		// if flag was false but env is true, set debug
		if debug == false && ret == true {
			debug = true
		}
	}
	if debug {
		log.Println("Debugging enabled.")
	}
	return nil
}

// this JSON pretty prints errors and debug
func p(b interface{}) string {
	s, _ := json.MarshalIndent(b, "", "  ")
	return string(s)
}
