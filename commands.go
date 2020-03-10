package main

import (
	"log"
	"strings"
	"time"

	"github.com/haukened/gourban"
	"samhofi.us/x/keybase/types/chat1"
)

func (b *bot) Ping(convid chat1.ConvIDStr) {
	Debug("Ping received in %s", convid)
	b.k.SendMessageByConvID(convid, "Pong!")
}

func (b *bot) Urban(convid chat1.ConvIDStr, mid chat1.MessageID, message []string) {
	Debug("Urban received in %s", convid)
	if len(message[1:]) == 0 {
		// no arguments to command
		return
	}
	queryText := strings.Join(message[1:], " ")
	result, err := gourban.Top(queryText)
	if err != nil {
		b.k.ReactByConvID(convid, mid, "Error.")
		log.Printf("%s\n", err)
	}
	dur := 10 * time.Minute
	b.k.SendEphemeralByConvID(convid, dur, "`Definition:` %s\n`Usage:` %s", result.Definition, result.Example)
}
