package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"samhofi.us/x/keybase"
	"samhofi.us/x/keybase/types/chat1"
	"samhofi.us/x/keybase/types/stellar1"
)

// RegisterHandlers is called by main to map these handler funcs to events
func (b *Bot) RegisterHandlers() {
	chat := b.chatHandler
	conv := b.convHandler
	wallet := b.walletHandler
	err := b.errHandler

	b.handlers = keybase.Handlers{
		ChatHandler:         &chat,
		ConversationHandler: &conv,
		WalletHandler:       &wallet,
		ErrorHandler:        &err,
	}
}

// chatHandler should handle all messages coming from the chat
func (b *Bot) chatHandler(m chat1.MsgSummary) {
	// only handle text, we don't really care about attachments
	if m.Content.TypeName != "text" {
		return
	}
	// only for debugging
	if m.Content.Text.Payments != nil {
		Debug(p(m.Content.Text.Payments))
	}
	// if the message is @myusername just perform the default function
	if strings.HasPrefix(m.Content.Text.Body, fmt.Sprintf("@%s", b.k.Username)) {
		words := strings.Fields(m.Content.Text.Body)
		b.Urban(m.ConvID, m.Id, words, m.Channel.MembersType)
	}
	// its a command for me, iterate through extended commands
	if strings.HasPrefix(m.Content.Text.Body, "!") {
		// break up the message into words
		words := strings.Fields(m.Content.Text.Body)
		// strip the ! from the first word, and lowercase to derive the command
		thisCommand := strings.ToLower(strings.Replace(words[0], "!", "", 1))
		// decide if this is askind for extended commands
		switch thisCommand {
		case "ping":
			b.Ping(m.ConvID)
		case "urban":
			fallthrough
		case "urbandictionary":
			b.Urban(m.ConvID, m.Id, words, m.Channel.MembersType)
		default:
			return
		}
	}
}

// handle conversations (this fires when a new conversation is initiated)
// i.e. when someone opens a conversation to you but hasn't sent a message yet
func (b *Bot) convHandler(m chat1.ConvSummary) {
	log.Println("---[ conv ]---")
	log.Println(p(m))
}

// this handles wallet events, like when someone send you money in chat
func (b *Bot) walletHandler(m stellar1.PaymentDetailsLocal) {
	if m.Summary.StatusSimplified > 0 {
		log.Printf("%s Payment of %s Received from %s txn %s !!!!\n", m.Summary.StatusDescription, m.Summary.AmountDescription, m.Summary.FromUsername, m.Summary.Id)
	}
}

// this handles all errors returned from the keybase binary
func (b *Bot) errHandler(m error) {
	log.Println("---[ error ]---")
	log.Println(p(m))
}

// this JSON pretty prints errors and debug
func p(b interface{}) string {
	s, _ := json.MarshalIndent(b, "", "  ")
	return string(s)
}
