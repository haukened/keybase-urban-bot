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

func (b *bot) RegisterHandlers() {
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

func (b *bot) chatHandler(m chat1.MsgSummary) {
	if m.Content.TypeName != "text" {
		return
	}

	Debug(m.Content.TypeName)

	if strings.HasPrefix(m.Content.Text.Body, fmt.Sprintf("@%s", b.k.Username)) {
		// message is @me so do my function
		words := strings.Fields(m.Content.Text.Body)
		b.Urban(m.ConvID, m.Id, words, m.Channel.MembersType)
	}

	if strings.HasPrefix(m.Content.Text.Body, "!") {
		// its a command for me, iterate through extended commands
		words := strings.Fields(m.Content.Text.Body)
		thisCommand := strings.ToLower(strings.Replace(words[0], "!", "", 1))
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

func (b *bot) convHandler(m chat1.ConvSummary) {
	log.Println("---[ conv ]---")
	log.Println(p(m))
}

func (b *bot) walletHandler(m stellar1.PaymentDetailsLocal) {
	if m.Summary.StatusSimplified == 3 {
		log.Printf("Payment of %s Received from %s!!!!\n", m.Summary.AmountDescription, m.Summary.FromUsername)
	}
}

func (b *bot) errHandler(m error) {
	log.Println("---[ error ]---")
	log.Println(p(m))
}

func p(b interface{}) string {
	s, _ := json.MarshalIndent(b, "", "  ")
	return string(s)
}
