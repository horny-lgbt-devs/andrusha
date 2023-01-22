package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
)

type CommandBalance struct {
	bot *Bot
}

var (
	_ ken.SlashCommand = (*CommandBalance)(nil)
)

func (c *CommandBalance) Name() string {
	return "balance"
}

func (c *CommandBalance) Description() string {
	return "Просмотр баланса"
}

func (c *CommandBalance) Version() string {
	return "1.0.0"
}

func (c *CommandBalance) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *CommandBalance) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *CommandBalance) Run(ctx ken.Context) (err error) {

	selfuser := ctx.GetEvent().Member.User.ID
	guild := ctx.GetEvent().GuildID

	switch c.bot.Storage.ExistsGuildUser(guild, selfuser) {
	case true:
		_, _, bal, _ := c.bot.Storage.GetGuildUser(guild, selfuser)

		err = ctx.RespondEmbed(&discordgo.MessageEmbed{
			Title:       "Ваш баланс:",
			Description: fmt.Sprint(bal),
		})
		return
	case false:

		c.bot.Storage.CreateGuildUser(guild, selfuser, 0)
		_, _, bal, _ := c.bot.Storage.GetGuildUser(guild, selfuser)

		err = ctx.RespondEmbed(&discordgo.MessageEmbed{
			Title:       "Ваш баланс:",
			Description: fmt.Sprint(bal),
		})
		return
	}

	return
}
