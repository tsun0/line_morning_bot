package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"

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

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowUTC := time.Now()
	nowJST := nowUTC.In(jst)
	noon := time.Date(nowJST.Year(), nowJST.Month(), nowJST.Day(), 12, 0, 0, 0, jst)
	log.Println(nowJST)
	log.Println(noon)

	if nowJST.Before(noon) {
		text := linebot.NewTextMessage("おはようございます")
		text2 := linebot.NewTextMessage("今日も1日がんばりましょう！")
		sticker := linebot.NewStickerMessage("2", randomSticker())
		messages = append(messages, text, text2, sticker)
		log.Println(messages)
		return messages
	}

	text := linebot.NewTextMessage("お疲れ様です")
	text2 := linebot.NewTextMessage("疲れたときは休憩しましょう！")
	sticker := linebot.NewStickerMessage("2", randomSticker())
	messages = append(messages, text, text2, sticker)
	log.Println(messages)
	return messages

}

func randomSticker() string {

	max, min := 179, 140
	return strconv.Itoa(rand.Intn(max-min) + min)

}
