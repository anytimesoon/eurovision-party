package conf

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"log"
)

var App AppConf

var v = viper.New()

type AppConf struct {
	DbPath      string `mapstructure:"DB_PATH"`
	ServHost    string `mapstructure:"BACKEND_HOST"`
	HttpProto   string `mapstructure:"HTTP_PROTOCOL"`
	Domain      string `mapstructure:"DOMAIN_NAME"`
	BotId       uuid.UUID
	BotIdString string `mapstructure:"BOT_ID"`
	BotName     string `mapstructure:"CHAT_BOT_NAME"`
	Assets      string `mapstructure:"ASSET_DIR"`
	Secret      string `mapstructure:"SECRET"`
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

	v.SetDefault("DB_PATH", "storage/")
	v.SetDefault("BACKEND_HOST", "localhost:8080")
	v.SetDefault("HTTP_PROTOCOL", "http://")
	v.SetDefault("CHAT_BOT_NAME", "Eurobot")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if v.GetString("SECRET") == "" {
		bytes := make([]byte, 32)
		if _, err := rand.Read(bytes); err != nil {
			panic(err)
		}
		hexString := hex.EncodeToString(bytes)
		v.Set("SECRET", hexString)
		err := v.WriteConfig()
		if err != nil {
			log.Fatal("Failed to write bot id to config.", err)
		}
	}

	err = v.Unmarshal(&App)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshalling config: %w", err))
	}

	log.Println("Domain name is:", App.Domain)

	if App.BotIdString != "" {
		App.BotId, err = uuid.Parse(App.BotIdString)
		if err != nil {
			log.Fatalln("Failed to parse bot id.", err)
		}
	}
}
