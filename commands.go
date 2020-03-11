package main

import (
	"log"
	"strings"
	"time"

	"github.com/haukened/gourban"
	"samhofi.us/x/keybase/types/chat1"
)

// Ping sends Pong! as a reply
func (b *bot) ping(convid chat1.ConvIDStr) {
	Debug("Ping received in %s", convid)
	b.k.SendMessageByConvID(convid, "Pong!")
}

// Urban performs and returns an urbandictionary.com lookup
func (b *bot) urban(convid chat1.ConvIDStr, mid chat1.MessageID, message []string, membersType string) {
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
	if membersType == "impteamnative" {
		// its a PM so just respond
		b.k.SendMessageByConvID(convid, "`Definition:` %s\n`Usage:` %s", result.Definition, result.Example)
	} else {
		// its a team so send with a 10 min fuse in case its REALLY BAD
		dur := 10 * time.Minute
		b.k.SendEphemeralByConvID(convid, dur, "`Definition:` %s\n`Usage:` %s", result.Definition, result.Example)
	}
}
