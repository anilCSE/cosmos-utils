package targets

import (
	"log"
	"strings"

	"github.com/vitwit/cosmos-utils/relayer-alerter/alerting"
	"github.com/vitwit/cosmos-utils/relayer-alerter/config"
)

// SendTelegramAlert sends the alert to telegram chat
func SendTelegramAlert(msg string, cfg *config.Config) error {
	if strings.EqualFold(cfg.EnableAlerts, "yes") == true {
		if err := alerting.NewTelegramAlerter().Send(msg, cfg.Telegram.BotToken, cfg.Telegram.ChatID); err != nil {
			log.Printf("failed to send telegram alert: %v", err)
			return err
		}
	}
	return nil
}

// SendEmailAlert to send mail
func SendEmailAlert(msg string, cfg *config.Config) error {
	if err := alerting.NewEmailAlerter().Send(msg, cfg.SendGrid.SendGridAPIToken, cfg.SendGrid.ReceiverMailAddress); err != nil {
		log.Printf("failed to send email alert: %v", err)
		return err
	}
	return nil
}
