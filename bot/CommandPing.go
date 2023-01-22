package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
)

type CommandPing struct{}

var (
	_ ken.SlashCommand = (*CommandPing)(nil)
	_ ken.DmCapable    = (*CommandPing)(nil)
)

func (c *CommandPing) Name() string {
	return "ping"
}

func (c *CommandPing) Description() string {
	return "Пинг понг!! жесть!!!!!!"
}

func (c *CommandPing) Version() string {
	return "1.0.0"
}

func (c *CommandPing) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *CommandPing) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionBoolean,
			Name:        "ping",
			Required:    true,
			Description: "Опция!!!! жееесть!!!!",
		},
	}
}

func (c *CommandPing) IsDmCapable() bool {
	return true
}

func (c *CommandPing) Run(ctx ken.Context) (err error) {
	val := ctx.Options().GetByName("ping").BoolValue()

	msg := "Понг!!!"
	if val {
		msg = "Пинг???"
	}

	err = ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})
	return
}
