package conf

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

var App AppConf

type AppConf struct {
	DbUsername string `mapstructure:"DB_USERNAME"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbHostname string `mapstructure:"DB_HOSTNAME"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`
	ServPort   string `mapstructure:"BACKEND_PORT"`
	ServHost   string `mapstructure:"BACKEND_HOSTNAME"`
	HttpProto  string `mapstructure:"HTTP_PROTOCOL"`
	Domain     string `mapstructure:"DOMAIN_NAME"`
	BotId      uuid.UUID
	BotName    string `mapstructure:"CHAT_BOT_NAME"`
	Assets     string `mapstructure:"ASSET_DIR"`
}

func (a *AppConf) SetBotId(id uuid.UUID) {
	a.BotId = id
}

func LoadConfig() {
	v := viper.New()
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
}
