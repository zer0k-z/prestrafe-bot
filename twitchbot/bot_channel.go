package twitchbot

import (
	"fmt"

	"github.com/gempir/go-twitch-irc"

	"prestrafe-bot/config"
)

type botChannel struct {
	name              string
	verificationToken string
	commands          []ChatCommand
	channelSink       ChatMessageSink
}

func newChannel(client *twitch.Client, config *config.ChannelConfig) *botChannel {
	commands := []ChatCommand{
		CreateWrCommandBuilder().
			WithConfig(config.GetCommandConfig("*")).
			WithConfig(config.GetCommandConfig("wr")).
			Build(),
	}

	return &botChannel{
		name:              config.Name,
		verificationToken: config.VerificationToken,
		commands:          commands,
		channelSink: func(format string, a ...interface{}) {
			client.Say(config.Name, fmt.Sprintf(format, a...))
		},
	}
}

func (c *botChannel) handle(user *twitch.User, message *twitch.Message) {
	for _, command := range c.commands {
		if command.TryHandle(user, message, c.channelSink) {
			return
		}
	}
}
