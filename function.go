package main

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

func generateJoinMessage() []linebot.Message {

	var messages []linebot.Message
	text := linebot.NewTextMessage("Thank you for comming this room!")
	log.Println(text.Text)
	sticker := linebot.NewStickerMessage("2", "144")
	log.Println(sticker.StickerID)
	messages = append(messages, text, sticker)
	log.Println(messages)
	return messages

}

func generateBeaconMessage() []linebot.Message {

	var messages []linebot.Message
	text := linebot.NewTextMessage("おはようございます")
	text2 := linebot.NewTextMessage("今日も1日がんばりましょう！")
	log.Println(text.Text)
	sticker := linebot.NewStickerMessage("2", randomSticker())
	log.Println(sticker.StickerID)
	messages = append(messages, text, text2, sticker)
	log.Println(messages)
	return messages

}

func randomSticker() string {

	max, min := 179, 140
	return strconv.Itoa(rand.Intn(max-min) + min)

}
