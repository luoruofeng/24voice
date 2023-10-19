package src

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const (
	// BotToken         = "5757148150:AAFGDMCzu98iHr8dEmEnWqd-41GIZqkwnZc"
	// ChannelName      = "@twentyfourvoice"
	// AudioFilePath    = "demo/demo.mp3"
	// ImageFilePath    = "demo/demo.jpg"
	// TrojanServerAddr = "tw03.yuncyo.xyz"
	// TrojanServerPort = 443
	// TrojanPassword   = "bb49d5c2-e693-3705-b422-d3de2fd675cd"
	// TrojanSNI        = "tw03zhuan.kulime.space"
	// ProxyURL         = "http://127.0.0.1:7890" // 请替换成 Clash 代理的实际地址和端口

	BotToken         = ""
	ChannelName      = ""
	AudioFilePath    = ""
	ImageFilePath    = ""
	TrojanServerAddr = ""
	TrojanServerPort = 443
	TrojanPassword   = ""
	TrojanSNI        = ""
	ProxyURL         = "" // 请替换成 Clash 代理的实际地址和端口
)

type TgSrv struct {
	Bot          *tgbotapi.BotAPI
	ChannelName  string
	ProxyURL     string
	BotToken     string
	LatestChatId int64
}

func createBot(proxyURL, botToken string) *tgbotapi.BotAPI {

	// 设置 Clash 代理地址和端口
	proxyAddr, err := url.Parse(proxyURL)
	if err != nil {
		fmt.Println("Error parsing proxy URL:", err)
		os.Exit(1)
	}

	// 创建一个自定义的 HTTP Transport，以便使用代理
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyAddr),
	}

	// 创建一个 HTTP 客户端，使用自定义 Transport
	client := &http.Client{
		Transport: transport,
		Timeout:   120 * time.Second,
	}

	// 创建 Telegram Bot API 客户端
	bot, err := tgbotapi.NewBotAPIWithClient(botToken, client)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	return bot
}

func NewTgSrv() *TgSrv {
	srv := &TgSrv{
		Bot:         createBot(ProxyURL, BotToken),
		ProxyURL:    ProxyURL,
		BotToken:    BotToken,
		ChannelName: ChannelName,
	}
	go srv.Updates()
	return srv
}

// 上传文件
func (s *TgSrv) UploadFile(chatID int64, filePath string) (string, error) {

	// 打开音频文件
	audioFilePath := filePath
	audioFile, err := os.Open(audioFilePath)
	if err != nil {
		log.Panic(err)
	}
	defer audioFile.Close()

	// 创建音频消息配置
	fileBytes, err := FileToBytes(filePath)
	if err != nil {
		log.Panic(err)
	}
	fileConfig := tgbotapi.FileBytes{
		Name:  filepath.Base(filePath),
		Bytes: fileBytes,
	}

	suffix := GetFileExtensionWithoutDot(filePath)
	if suffix == "mp3" || suffix == "ogg" || suffix == "flac" || suffix == "wav" || suffix == "m4a" {
		// 上传音频文件到 Telegram 服务器
		msg := tgbotapi.NewAudioUpload(chatID, fileConfig)
		msg.Caption = "收到新的音频文件"
		respMes, err := s.Bot.Send(msg)
		if err != nil {
			return "", err
		}
		log.Println("Audio file uploaded successfully.", respMes.Chat, respMes.Audio)
		if respMes.Audio.FileID != "" {
			return respMes.Audio.FileID, nil
		}
	} else if suffix == "jpeg" || suffix == "jpg" || suffix == "png" || suffix == "gif" || suffix == "bmp" || suffix == "tiff" || suffix == "svg" {
		// 上传音频文件到 Telegram 服务器
		msg := tgbotapi.NewPhotoUpload(chatID, fileConfig)
		msg.Caption = "收到新的图片文件"
		respMes, err := s.Bot.Send(msg)
		if err != nil {
			return "", err
		}
		log.Println("Photo file uploaded successfully.", respMes.Chat, respMes.Photo)
		if (*respMes.Photo)[0].FileID != "" {
			return (*respMes.Photo)[0].FileID, nil
		}
	} else if suffix == "pdf" || suffix == "docx" || suffix == "doc" || suffix == "xlsx" || suffix == "xls" || suffix == "ppt" || suffix == "pptx" || suffix == "txt" || suffix == "csv" || suffix == "rtf" || suffix == "html" || suffix == "markdown" {
		// 上传文件到 Telegram 服务器
		msg := tgbotapi.NewDocumentUpload(chatID, fileConfig)
		msg.Caption = "收到新的文件"
		respMes, err := s.Bot.Send(msg)
		if err != nil {
			return "", err
		}
		log.Println("Document file uploaded successfully.", respMes.Chat, respMes.Document)
		if respMes.Document.FileID != "" {
			return respMes.Document.FileID, nil
		}
	} else {
		return "", errors.New("文件后缀不支持")
	}

	return "", nil
}

// 分享文件
func (s *TgSrv) ShareFile(chatId int64, suffix string, fileId string) (*tgbotapi.Message, error) {
	if suffix == "mp3" || suffix == "ogg" || suffix == "flac" || suffix == "wav" || suffix == "m4a" {
		// 构建音频流链接消息
		audioMsg := tgbotapi.NewAudioShare(chatId, fileId)
		audioMsg.Caption = "分享新的音乐"
		// 发送音频流链接消息
		result, err := s.Bot.Send(audioMsg)
		if err != nil {
			return nil, err
		}
		return &result, nil
	} else if suffix == "jpeg" || suffix == "jpg" || suffix == "png" || suffix == "gif" || suffix == "bmp" || suffix == "tiff" || suffix == "svg" {
		//TODO
		// 构建音频流链接消息
		imageMsg := tgbotapi.NewPhotoShare(chatId, fileId)
		imageMsg.Caption = "分享新的图片"
		// 发送音频流链接消息
		result, err := s.Bot.Send(imageMsg)
		if err != nil {
			return nil, err
		}
		return &result, nil
	} else if suffix == "pdf" || suffix == "docx" || suffix == "doc" || suffix == "xlsx" || suffix == "xls" || suffix == "ppt" || suffix == "pptx" || suffix == "txt" || suffix == "csv" || suffix == "rtf" || suffix == "html" || suffix == "markdown" {
		//TODO
		return nil, errors.New("补全代码")
	} else {
		return nil, errors.New("文件后缀不支持")
	}
}

func ListFilesInFolder(folderPath string) ([]string, error) {
	var fileList []string

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileList = append(fileList, info.Name())
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileList, nil
}

// 更新消息
func (s *TgSrv) Updates() {
	// 设置更新超时时间
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 120
	// 获取更新通道
	updates, err := s.Bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	// 处理接收到的消息
	for update := range updates {
		if update.ChannelPost != nil {

			if strings.HasPrefix(update.ChannelPost.Text, "/books") {
				fileList, err := ListFilesInFolder("demo/")
				if err != nil {
					msg := tgbotapi.NewMessage(update.ChannelPost.Chat.ID, "无法获取文件列表")
					s.Bot.Send(msg)
					continue
				}

				var responseText string
				for _, file := range fileList {
					responseText += file + "\n"
				}

				msg := tgbotapi.NewMessage(update.ChannelPost.Chat.ID, responseText)
				s.Bot.Send(msg)
			}
		}

		// 检查 update.ChannelPost 是否为空来获取频道消息的内容
		if update.ChannelPost != nil {
			chatID := update.ChannelPost.Chat.ID
			s.LatestChatId = chatID
			// 根据不同属性的存在与否来确定消息类型
			if update.ChannelPost.Text != "" {
				// 文本消息，直接回复相同的内容
				// messageText := update.ChannelPost.Text
				// msg := tgbotapi.NewMessage(chatID, messageText)
				// _, err := s.Bot.Send(msg)
				// if err != nil {
				// 	log.Println("Error sending reply:", err)
				// }
			} else if update.ChannelPost.Photo != nil {
				// 图片消息，直接回复相同的图片
				// 此处添加处理图片消息的逻辑
			} else if update.ChannelPost.Video != nil {
				// 视频消息，直接回复相同的视频
				// 此处添加处理视频消息的逻辑
			} else if update.ChannelPost.Audio != nil {
				// 音频消息，直接回复相同的音频
				// 此处添加处理音频消息的逻辑
			} else {
				// 未知类型的消息
				log.Println("Unknown message type")
			}
		}
	}
}

func (s *TgSrv) Send(message string) error {
	msg := tgbotapi.NewMessageToChannel(s.ChannelName, message)
	_, err := s.Bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
