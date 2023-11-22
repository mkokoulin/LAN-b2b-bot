package main

import (
	"context"
	"lan_b2b_bot/internal/config"
	"lan_b2b_bot/internal/services"
	"log"
	"reflect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
		return
	}

	gc, err := services.NewGoogleClient(ctx, cfg.Google.Scope)
	if err != nil {
		log.Fatal(err)
		return
	}

	requestsSheets, err := services.NewRequestsSheets(ctx, gc, cfg.GoogleSheets.Requests.SpreadsheetId, cfg.GoogleSheets.Requests.ReadRange)
	if err != nil {
		log.Fatal(err)
		return
	}

	builderSheets, err := services.NewBuilderSheets(ctx, gc, cfg.GoogleSheets.Builder.SpreadsheetId, cfg.GoogleSheets.Builder.ReadRange)
	if err != nil {
		log.Fatal(err)
		return
	}


	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60 * 3

	updates := bot.GetUpdatesChan(u)

	currentStep := 0
	finalStepIsShown := false

	request := []string{}
	steps := []string{}

	for update := range updates {
		if update.Message == nil {
            continue
        }

		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String || reflect.TypeOf(update.Message.Text).Kind() == reflect.Int && update.Message.Text != "" { // If we got a message
			if update.Message.Text == "/start" {
				request = []string{}
				currentStep = 0
				finalStepIsShown = false
			}

			if currentStep == 0 {
				steps, err = builderSheets.GetTexts(ctx)
				if err != nil {
					log.Fatal(err)
					return
				}

				request = append(request, update.Message.Chat.FirstName)
				request = append(request, update.Message.Chat.LastName)
				request = append(request, update.Message.Chat.UserName)
			}
			
			if currentStep >= 1 {
				request = append(request, update.Message.Text)
			}

			if !finalStepIsShown {
				if currentStep <= len(steps) - 1 {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, steps[currentStep])
					bot.Send(msg)
		
					currentStep++
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Благодарим за проявленный интерес! Мы скоро свяжемся с вами")
					bot.Send(msg)
	
					finalStepIsShown = true;

					_, err := requestsSheets.CreateRequest(ctx, request)
					if err != nil {
						log.Panic(err)
						return
					}
				}
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Наш менеджер скоро свяжется с вами. Если хотите отправить еще одни запрос введите команду /start")
				bot.Send(msg)
			}
		}
	}
}