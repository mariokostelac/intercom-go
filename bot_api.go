package intercom

import (
	"encoding/json"

	"github.com/mariokostelac/intercom-go/interfaces"
)

// BotRepository defines the interface for working with Bots through the API.
type BotRepository interface {
	save(*Bot) (Bot, error)
}

// BotAPI implements BotRepository
type BotAPI struct {
	httpClient interfaces.HTTPClient
}

func (api BotAPI) save(bot *Bot) (Bot, error) {
	data, err := api.httpClient.Post("/bots", bot)
	var savedBot Bot
	if err != nil {
		return savedBot, err
	}
	err = json.Unmarshal(data, &savedBot)
	return savedBot, err
}
