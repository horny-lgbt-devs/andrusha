// https://github.com/zekroTJA/ken/tree/master/middlewares/ratelimit
package andrusharatelimit

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
)

type Middleware struct {
	manager Manager
}

var _ ken.MiddlewareBefore = (*Middleware)(nil)

func New(manager ...Manager) *Middleware {
	var m Manager

	if len(manager) > 0 && manager[0] != nil {
		m = manager[0]
	} else {
		m = newInternalManager()
	}

	return &Middleware{m}
}

func (m *Middleware) Before(ctx *ken.Ctx) (next bool, err error) {
	c, ok := ctx.Command.(LimitedCommand)
	if !ok {
		return true, nil
	}

	var guildID string
	if c.IsLimiterGlobal() {
		guildID = "__global__"
	} else {
		var ch *discordgo.Channel
		ch, err = ctx.Channel()
		if err != nil {
			return
		}
		if ch.Type == discordgo.ChannelTypeDM || ch.Type == discordgo.ChannelTypeGroupDM {
			guildID = "__dm__"
		} else {
			guildID = ctx.GetEvent().GuildID
		}
	}

	limiter := m.manager.GetLimiter(ctx.Command, ctx.User().ID, guildID)
	if ok, next := limiter.Take(); !ok {
		err := ctx.RespondError(fmt.Sprintf(
			"\nПодождите %s перед повторным использованием команды!",
			next.Round(1*time.Second).String()), "<:blobfoxwhaaaat:1060626643992981635> Не так быстро!")
		return false, err
	}

	return true, nil
}
