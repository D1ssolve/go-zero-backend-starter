package telegram

import (
	"context"
	"fmt"
	"strings"

	"github.com/D1ssolve/go-zero-backend-starter/internal/domain/note"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/zeromicro/go-zero/core/logx"
)

type Transport struct {
	bot     *bot.Bot
	service *note.Service
}

func New(token string, service *note.Service) (*Transport, error) {
	telegramBot, err := bot.New(token, bot.WithErrorsHandler(func(err error) {
		logx.Errorf("telegram update: %v", err)
	}))
	if err != nil {
		return nil, err
	}

	transport := &Transport{bot: telegramBot, service: service}
	telegramBot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, transport.start)
	telegramBot.RegisterHandler(bot.HandlerTypeMessageText, "/notes", bot.MatchTypeExact, transport.list)
	telegramBot.RegisterHandler(bot.HandlerTypeMessageText, "/add", bot.MatchTypePrefix, transport.add)
	return transport, nil
}

func (t *Transport) Run(ctx context.Context) {
	t.bot.Start(ctx)
}

func (t *Transport) start(ctx context.Context, telegramBot *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	_, err := telegramBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Ready. Use /add <text> and /notes.",
	})
	if err != nil {
		logx.WithContext(ctx).Errorf("send telegram message: %v", err)
	}
}

func (t *Transport) add(ctx context.Context, telegramBot *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	text := strings.TrimSpace(strings.TrimPrefix(update.Message.Text, "/add"))
	created, err := t.service.Create(ctx, text)
	message := "Created note #" + fmt.Sprint(created.ID)
	if err != nil {
		message = "Could not create note: " + err.Error()
	}
	if _, err := telegramBot.SendMessage(ctx, &bot.SendMessageParams{ChatID: update.Message.Chat.ID, Text: message}); err != nil {
		logx.WithContext(ctx).Errorf("send telegram message: %v", err)
	}
}

func (t *Transport) list(ctx context.Context, telegramBot *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	items, err := t.service.List(ctx, 10)
	message := "No notes yet."
	if err != nil {
		message = "Could not load notes."
	} else if len(items) > 0 {
		lines := make([]string, 0, len(items))
		for _, item := range items {
			lines = append(lines, fmt.Sprintf("#%d — %s", item.ID, item.Text))
		}
		message = strings.Join(lines, "\n")
	}
	if _, err := telegramBot.SendMessage(ctx, &bot.SendMessageParams{ChatID: update.Message.Chat.ID, Text: message}); err != nil {
		logx.WithContext(ctx).Errorf("send telegram message: %v", err)
	}
}
