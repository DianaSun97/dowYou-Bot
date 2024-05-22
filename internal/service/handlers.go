package service

import (
	"bytes"
	"context"
	"example.com/mod/internal/domain"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"io"
	"net/url"
)

func (s *Service) matchState(state domain.State) bot.MatchFunc {
	return func(u *models.Update) bool {
		if u.Message == nil {
			return false
		}
		currentState, ok := s.states[u.Message.Chat.ID]
		if !ok {
			return false
		}
		return currentState == state
	}
}

func (s *Service) handlerStart(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Use buttons for downloads video",
		ReplyMarkup: models.ReplyKeyboardMarkup{
			ResizeKeyboard: true,
			Keyboard: [][]models.KeyboardButton{
				{
					{Text: "Download video"},
				},
				{
					{Text: "📚 Course"},
					{Text: "📊 Somthing else"},
				},
			},
		},
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func (s *Service) handlerAskLinkYTVideo(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Add link from YouTube:",
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	s.states[update.Message.Chat.ID] = domain.StateWaitingVideo
}

func (s *Service) handlerDownloadYTVideo(ctx context.Context, b *bot.Bot, update *models.Update) {
	parsedURL, err := url.Parse(update.Message.Text)
	if err != nil {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Not valid URL",
		})
		return
	}

	msg, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Wait....",
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	video, err := s.repo.GetVideo(parsedURL.String())
	if err != nil {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Error video download",
		})
		return
	}

	_, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    msg.Chat.ID,
		MessageID: msg.ID,
		Text:      fmt.Sprintf("Video %s ready for dowload", video.Title),
		ReplyMarkup: models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{Text: "Download", CallbackData: "download_video:" + parsedURL.String()},
				},
			},
		},
	})
}

func (s *Service) handlerUploadYTVideo(ctx context.Context, b *bot.Bot, update *models.Update) {
	videoURL := update.CallbackQuery.Data[len("download_video:"):]
	video, err := s.repo.GetVideo(videoURL)
	if err != nil {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Text:   "Error video download",
		})
		return
	}
	formats := video.Formats
	if len(formats) == 0 {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Text:   "Error video",
		})
		return
	}

	reader, _, err := s.repo.DownloadVideo(video, &formats[0])
	if err != nil {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Text:   "Error video download",
		})
		return
	}

	go func() {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Text:   "Start video download",
		})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		writer := bytes.Buffer{}
		_, err = io.Copy(&writer, reader)
		if err != nil && err != io.EOF {
			fmt.Println("Error saving video to file:", err)
			return
		}

		_, err = b.SendDocument(ctx, &bot.SendDocumentParams{
			ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Document: &models.InputFileUpload{
				Data:     &writer,
				Filename: video.ID + ".mp4",
			},
		})
		if err != nil {
			fmt.Println("Error sending video:", err)
			return
		}

		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Text:   "Video successfully downloaded!",
		})

	}()

}
