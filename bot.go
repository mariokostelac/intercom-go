package intercom

import (
	"encoding/json"
	"fmt"
)

// BotService handles interactions with the API through a BotRepository.
type BotService struct {
	Repository BotRepository
}

// Save a new Bot, or update an existing one.
func (svc *BotService) Save(bot *Bot) (Bot, error) {
	return svc.Repository.save(bot)
}

// Bot represents a bot in Intercom.
type Bot struct {
	ID     json.Number `json:"id"`
	Name   string      `json:"name"`
	Avatar UserAvatar  `json:"avatar"`
}

// MessageAddress gets the address for a Bot in order to send a message on behalf of them
func (b Bot) MessageAddress() MessageAddress {
	return MessageAddress{
		Type: "bot",
		ID:   b.ID.String(),
	}
}

func (b Bot) String() string {
	return fmt.Sprintf("[intercom] Bot { id: %s, name: %s }", b.ID, b.Name)
}
