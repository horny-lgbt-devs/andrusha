package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
)

type CommandHelp struct{}

var (
	_ ken.SlashCommand = (*CommandHelp)(nil)
	_ ken.DmCapable    = (*CommandHelp)(nil)
)

func (c *CommandHelp) Name() string {
	return "help"
}

func (c *CommandHelp) Description() string {
	return "Список команд"
}

func (c *CommandHelp) Version() string {
	return "1.0.0"
}

func (c *CommandHelp) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *CommandHelp) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *CommandHelp) IsDmCapable() bool {
	return true
}

func (c *CommandHelp) Run(ctx ken.Context) (err error) {

	err = ctx.RespondEmbed(&discordgo.MessageEmbed{
		Title: "Данная команда ещё не готова",
		Image: &discordgo.MessageEmbedImage{
			URL: "https://media.discordapp.net/attachments/1013822339382771893/1063143914334330960/caption.gif",
		},
	})
	return
}
