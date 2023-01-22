package bot

import (
	"andrusha/andrusharatelimit"
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
)

type CommandWork struct {
	bot *Bot
}

var (
	_ ken.SlashCommand                 = (*CommandWork)(nil)
	_ andrusharatelimit.LimitedCommand = (*CommandWork)(nil)
)

func (c *CommandWork) Name() string {
	return "work"
}

func (c *CommandWork) Description() string {
	return "Один из способов заработать валюту"
}

func (c *CommandWork) Version() string {
	return "1.0.0"
}

func (c *CommandWork) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *CommandWork) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *CommandWork) LimiterBurst() int {
	return 1
}

func (c *CommandWork) LimiterRestoration() time.Duration {
	return 3 * time.Hour
}

func (c *CommandWork) IsLimiterGlobal() bool {
	return false
}

func (c *CommandWork) Run(ctx ken.Context) (err error) {

	selfuser := ctx.GetEvent().Member.User.ID
	guild := ctx.GetEvent().GuildID

	randbal := c.randint(100, 600)

	c.bot.Storage.AddBalanceGuildUser(guild, selfuser, int64(randbal))

	err = ctx.RespondEmbed(&discordgo.MessageEmbed{
		Title: "Вы заработали " + fmt.Sprint(randbal) + " ДЕНЕГ",
	})

	return
}

func (c *CommandWork) randint(min, max int) int {

	rand.Seed(time.Now().UnixNano())

	res := rand.Intn(max - min + 1)

	return res
}
