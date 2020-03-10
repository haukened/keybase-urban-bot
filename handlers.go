package main

import (
	"encoding/json"
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

	if strings.HasPrefix(m.Content.Text.Body, "!ping") {
		b.Ping(m.ConvID, m.Id)
	}
}

func (b *bot) convHandler(m chat1.ConvSummary) {
	log.Println("---[ conv ]---")
	log.Println(p(m))
}

func (b *bot) walletHandler(m stellar1.PaymentDetailsLocal) {
	log.Println("---[ wallet ]---")
	log.Println(p(m))
}

func (b *bot) errHandler(m error) {
	log.Println("---[ error ]---")
	log.Println(p(m))
}

func p(b interface{}) string {
	s, _ := json.MarshalIndent(b, "", "  ")
	return string(s)
}
