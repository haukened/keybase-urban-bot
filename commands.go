package main

import (
	"samhofi.us/x/keybase/types/chat1"
)

func (b *bot) Ping(convid chat1.ConvIDStr, msgid chat1.MessageID) {
	Debug("Ping received in %s", convid)
	b.k.ReplyByConvID(convid, msgid, "Pong!")
}
