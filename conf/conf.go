package conf

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"log"
)

var App AppConf

var v = viper.New()

type AppConf struct {
	DbUsername  string `mapstructure:"DB_USERNAME"`
	DbPassword  string `mapstructure:"DB_PASSWORD"`
	DbHostname  string `mapstructure:"DB_HOSTNAME"`
	DbPort      string `mapstructure:"DB_PORT"`
	DbName      string `mapstructure:"DB_NAME"`
	ServPort    string `mapstructure:"BACKEND_PORT"`
	ServHost    string `mapstructure:"BACKEND_HOSTNAME"`
	HttpProto   string `mapstructure:"HTTP_PROTOCOL"`
	Domain      string `mapstructure:"DOMAIN_NAME"`
	BotId       uuid.UUID
	BotIdString string `mapstructure:"BOT_ID"`
	BotName     string `mapstructure:"CHAT_BOT_NAME"`
	Assets      string `mapstructure:"ASSET_DIR"`
}

func (a *AppConf) SetBotId(id uuid.UUID) {
	a.BotId = id
	v.Set("BOT_ID", id)
	err := v.WriteConfig()
	if err != nil {
		log.Fatal("Failed to write bot id to config.", err)
	}
}

func LoadConfig() {
	v.SetConfigName("app")
	v.AddConfigPath("conf/")

	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = v.Unmarshal(&App)
	fmt.Printf("%+v", App)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshalling config: %w", err))
	}

	if App.BotIdString != "" {
		App.BotId, err = uuid.Parse(App.BotIdString)
		if err != nil {
			log.Fatalln("Failed to parse bot id.", err)
		}
	}
}
