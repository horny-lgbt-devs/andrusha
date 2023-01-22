package bot

import (
	"andrusha/andrusharatelimit"
	"andrusha/database"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/zekrotja/ken"

	"github.com/zekrotja/ken/store"
	"xorm.io/xorm"
)

type Configuration struct {
	Token    string `required:"true"`
	Owner    string `required:"true"`
	EmbedCol int    `default:"0x6ba9cf" required:"true"`
	EmbedErr int    `default:"0xbd1515" required:"true"`
}

type Bot struct {
	config  Configuration
	session *discordgo.Session
	Sword   *ken.Ken
	Storage *database.Database
}

func New(config Configuration) (bot *Bot, err error) {

	bot = &Bot{
		config: config,
	}

	bot.session, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create new Discord session")
	}

	bot.Sword, err = ken.New(bot.session, ken.Options{
		CommandStore:   store.NewDefault(),
		OnCommandError: bot.commandErrorHandler,
		EmbedColors: ken.EmbedColors{
			Default: config.EmbedCol,
			Error:   config.EmbedErr,
		},
	})

	bot.RegisterCommands()

	bot.Sword.Session().AddHandler(bot.Ready)

	bot.session.Open()

	engine, errr := xorm.NewEngine("sqlite3", "./database/database.db")
	if errr != nil {
		fmt.Println(errr)
	}

	bot.Storage = database.Open(engine)

	return
}

func (bot Bot) Close() {

	log.Println("Смэрть")
	bot.session.Close()
}

func (b *Bot) RegisterCommands() (err error) {

	b.Sword.RegisterMiddlewares(andrusharatelimit.New())

	var cmdmap = map[string]ken.Command{
		"balance":        &CommandBalance{b},
		"balance-member": &CommandBalanceMember{b},
		"help":           &CommandHelp{},
		"ping":           &CommandPing{},
		"work":           &CommandWork{b},
	}

	for key := range cmdmap {
		err := b.Sword.RegisterCommands(cmdmap[key])
		if err != nil {
			return errors.Wrap(err, "Failed to register commands")
		}
		fmt.Println(key)
	}

	return
}

func (bot Bot) Ready(s *discordgo.Session, r *discordgo.Ready) {

	log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
}

func (bot Bot) commandErrorHandler(err error, ctx *ken.Ctx) {

	ctx.Defer()

	if err == ken.ErrNotDMCapable {
		ctx.FollowUpError("Эту команду нельзя использовать в личных сообщениях", "")
		return
	}

	ctx.FollowUpError(
		fmt.Sprintf(":(\n```\n%s\n```", err.Error()),
		"Произошла ошибка при выполнении команды")
}
