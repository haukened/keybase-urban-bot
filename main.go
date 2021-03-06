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
type bot struct {
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

// newBot returns a new empty bot
func newBot() *bot {
	var b bot
	b.k = keybase.NewKeybase()
	b.handlers = keybase.Handlers{}
	b.opts = keybase.RunOptions{}
	b.payments = make(map[stellar1.PaymentID]botReply)
	return &b
}

// this handles setting up command advertisements and aliases
func (b *bot) registerCommands() {
	opts := keybase.AdvertiseCommandsOptions{
		Advertisements: []chat1.AdvertiseCommandAPIParam{
			{
				Typ: "public",
				Commands: []chat1.UserBotCommandInput{
					{
						Name:        "urban",
						Description: "Perform an urbandictionary lookup",
						Usage:       "<word(s)>",
					},
				},
			},
		},
	}
	b.k.AdvertiseCommands(opts)
}

// run performs a proxy main function
func (b *bot) run(args []string) error {
	// parse the arguments
	err := parseArgs(args)
	if err != nil {
		return err
	}

	b.registerHandlers()
	b.registerCommands()

	log.Println("Starting...")
	b.k.Run(b.handlers, &b.opts)
	return nil
}

// main is a thin skeleton, proxied to Bot.Run()
func main() {
	b := newBot()
	if err := b.run(os.Args); err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
}
