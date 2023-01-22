package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
)

type CommandBalanceMember struct {
	bot *Bot
}

var (
	_ ken.SlashCommand = (*CommandBalanceMember)(nil)
)

func (c *CommandBalanceMember) Name() string {
	return "balance-member"
}

func (c *CommandBalanceMember) Description() string {
	return "Просмотр баланса нужного вам пользователя"
}

func (c *CommandBalanceMember) Version() string {
	return "1.0.0"
}

func (c *CommandBalanceMember) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *CommandBalanceMember) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionUser,
			Name:        "user",
			Required:    true,
			Description: "Пользователь, чей баланс вы хотите просмотреть",
		},
	}
}

func (c *CommandBalanceMember) Run(ctx ken.Context) (err error) {

	user := ctx.Options().GetByName("user")
	guild := ctx.GetEvent().GuildID

	switch c.bot.Storage.ExistsGuildUser(guild, user.UserValue(ctx).ID) {
	case true:

		_, _, bal, _ := c.bot.Storage.GetGuildUser(guild, user.UserValue(ctx).ID)

		err = ctx.RespondEmbed(&discordgo.MessageEmbed{
			Title:       "Баланс пользователя " + user.UserValue(ctx).Username + ":",
			Description: fmt.Sprint(bal),
		})

	case false:
		err = ctx.RespondEmbed(&discordgo.MessageEmbed{
			Title: "Пользователь не найден, либо его кошелёк не создан!",
		})
	}

	return
}
