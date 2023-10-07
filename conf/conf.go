package conf

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

var (
	Db     dbConf
	Server serverConf
	Bot    botConf
	main   mainConf
)

type (
	dbConf struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Hostname string `mapstructure:"hostname"`
		Port     string `mapstructure:"port"`
		Name     string `mapstructure:"name"`
	}

	serverConf struct {
		Port string `mapstructure:"port"`
		Url  string `mapstructure:"url"`
	}

	botConf struct {
		Id   uuid.UUID `mapstructure:"id"`
		Name string    `mapstructure:"name"`
	}

	mainConf struct {
		Db   dbConf     `mapstructure:"db"`
		Serv serverConf `mapstructure:"server"`
		Bot  botConf    `mapstructure:"bot"`
	}
)

func (b *botConf) SetId(id uuid.UUID) {
	viper.Set("bot.id", id.String())
	b.Id = id
}

func LoadConfig() {
	viper.SetConfigName("defaults")
	viper.AddConfigPath("conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&main)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshalling config: %w", err))
	}

	Db = main.Db
	Server = main.Serv
	Bot = main.Bot
}
