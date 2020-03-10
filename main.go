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
