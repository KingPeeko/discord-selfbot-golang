package selfbot

import (
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Session *discordgo.Session
	User    *discordgo.User
}

func createBot(config Config) (Bot, error) {
	session, err := discordgo.New(config.Token)

	if err != nil {
		return Bot{}, err
	}

	err = session.Open()
	if err != nil {
		return Bot{}, err
	}

	user, err := session.User("@me")
	if err != nil {
		return Bot{}, err
	}

	bot := Bot{
		Session: session,
		User:    user,
	}

	return bot, nil
}
