package main

import (
	"log"
	"os"

	"samhofi.us/x/keybase"
	"samhofi.us/x/keybase/types/chat1"
	"samhofi.us/x/keybase/types/stellar1"
)

// this global controls debug printing
var debug bool

// Bot holds the necessary information for the bot to work.
type Bot struct {
	k        *keybase.Keybase
	handlers keybase.Handlers
	opts     keybase.RunOptions
	payments map[stellar1.PaymentID]botReply
}

// hold reply information when needed
type botReply struct {
	convID chat1.ConvIDStr
	msgID  chat1.MessageID
}

// NewBot returns a new empty bot
func NewBot() *Bot {
	var b Bot
	b.k = keybase.NewKeybase()
	b.handlers = keybase.Handlers{}
	b.opts = keybase.RunOptions{}
	b.payments = make(map[stellar1.PaymentID]botReply)
	return &b
}

// Run performs a proxy main function
func (b *Bot) Run(args []string) error {
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

// main is a thin skeleton, proxied to Bot.Run()
func main() {
	b := NewBot()
	if err := b.Run(os.Args); err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
}
