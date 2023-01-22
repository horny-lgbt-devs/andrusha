package database

import "fmt"

type GuildUser struct {
	GuildID string
	UserID  string
	Balance int64
}

func (db *Database) AddBalanceGuildUser(guildid, userid string, bal int64) error {

	if db.ExistsGuildUser(guildid, userid) {

		_, _, bln, _ := db.GetGuildUser(guildid, userid)
		err := db.UpdateGuildUser(guildid, userid, bal+bln)
		if err != nil {
			fmt.Println(err)
		}
	} else {

		db.CreateGuildUser(guildid, userid, bal)
	}

	return nil
}

func (db *Database) LeadersGuildUser(guildid string) error {

	err := db.Session.Limit(1).Where("guild_id = ?", guildid)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}

func (db *Database) UpdateGuildUser(guildid, userid string, bal int64) error {

	_, err := db.Session.Limit(1).Where("guild_id = ?", guildid).And("user_id = ?", userid).Update(&GuildUser{GuildID: guildid, UserID: userid, Balance: bal})
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (db *Database) CreateGuildUser(guildid, userid string, bal int64) error {

	_, err := db.Session.InsertOne(&GuildUser{GuildID: guildid, UserID: userid, Balance: bal})
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) UpdateOrCreateGuildUser(guildid, userid string, bal int64) error {

	user := &GuildUser{GuildID: guildid, UserID: userid}

	if ok, _ := db.Session.Get(user); !ok {

		err := db.CreateGuildUser(guildid, userid, bal)
		if err != nil {
			return err
		}

	} else {

		err := db.UpdateGuildUser(guildid, userid, bal)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *Database) GetGuildUser(guildid, userid string) (string, string, int64, error) {

	var us []GuildUser

	err := db.Session.Limit(1).Find(&us, &GuildUser{GuildID: guildid, UserID: userid})
	if err != nil {
		e := "err"
		return e, e, 0, err
	}

	return us[0].GuildID, us[0].UserID, us[0].Balance, nil
}

func (db *Database) ExistsGuildUser(guildid, userid string) bool {

	user := &GuildUser{GuildID: guildid, UserID: userid}

	if ok, _ := db.Session.Get(user); ok {
		return true
	}

	return false
}
