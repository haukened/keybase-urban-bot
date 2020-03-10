package main

import (
	"log"
)

func Debug(s string, a ...interface{}) {
	if debug {
		log.Printf(s, a...)
	}
}

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
