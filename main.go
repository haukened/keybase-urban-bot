package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"samhofi.us/x/keybase"
)

var debug bool

type bot struct {
	k        *keybase.Keybase
	handlers keybase.Handlers
	opts     keybase.RunOptions
}

func NewBot() *bot {
	var b bot
	b.k = keybase.NewKeybase()
	b.handlers = keybase.Handlers{}
	b.opts = keybase.RunOptions{}
	return &b
}

func (b *bot) Run(args []string) error {
	// parse the arguments
	err := parseArgs(args)
	if err != nil {
		return err
	}
	//b.SetOptions()
	b.RegisterHandlers()

	log.Println("Starting...")
	b.k.Run(b.handlers, &b.opts)
	return nil
}

func main() {
	b := NewBot()
	if err := b.Run(os.Args); err != nil {
		log.Printf("%s", err)
		os.Exit(1)
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
