package main

import (
	"flag"
	"log"
	"os"

	"samhofi.us/x/keybase"
	//"samhofi.us/x/keybase/types/chat1"
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

/*
func (b *bot) SetOptions() {
	channel := chat1.ChatChannel{
		Name:        "keybase_git",
		TopicName:   "general",
		MembersType: keybase.TEAM,
	}

	b.opts = keybase.RunOptions{
		FilterChannel: channel,
	}
}
*/

func (b *bot) Run(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.BoolVar(&debug, "debug", false, "enables command debugging")
	if err := flags.Parse(args[1:]); err != nil {
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
