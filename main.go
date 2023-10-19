package main

import (
	"github.com/luoruofeng/24voice/src"
)

const (
	BotToken         = "5757148150:AAFGDMCzu98iHr8dEmEnWqd-41GIZqkwnZc"
	ChannelName      = "@twentyfourvoice"
	AudioFilePath    = "demo/demo.mp3"
	ImageFilePath    = "demo/demo.jpg"
	TrojanServerAddr = "tw03.yuncyo.xyz"
	TrojanServerPort = 443
	TrojanPassword   = "bb49d5c2-e693-3705-b422-d3de2fd675cd"
	TrojanSNI        = "tw03zhuan.kulime.space"
	ProxyURL         = "http://127.0.0.1:7890" // 请替换成 Clash 代理的实际地址和端口
)

func main() {
	src.NewTgSrv()

	// 获取机器人最后一次响应的对话ID
	// for {
	// 	time.Sleep(time.Second)
	// 	fmt.Println(ts.LatestChatId)

	// 	if ts.LatestChatId != 0 {
	// 		// 发送MP3
	// 		fileId, _ := ts.UploadFile(ts.LatestChatId, "demo/demo.jpg")

	// 		// 分享图片
	// 		mes, err := ts.ShareFile(ts.LatestChatId, "jpg", fileId)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 		fmt.Println((*mes.Photo)[0].FileID)
	// 		break
	// 	}
	// }

	// 发送消息
	// ts.Send("这是一条来自 Bot 的消息")

	// 在这里执行主程序逻辑，或者休眠主 Goroutine 以保持程序运行
	select {}
}
