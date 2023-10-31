package conf

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"os"
)

var App AppConf

type AppConf struct {
	DbUsername   string `mapstructure:"DB_USERNAME"`
	DbPassword   string `mapstructure:"DB_PASSWORD"`
	DbHostname   string `mapstructure:"DB_HOSTNAME"`
	DbPort       string `mapstructure:"DB_PORT"`
	DbName       string `mapstructure:"DB_NAME"`
	ServPort     string `mapstructure:"BACKEND_PORT"`
	ServHost     string `mapstructure:"BACKEND_HOSTNAME"`
	Domain       string `mapstructure:"DOMAIN_NAME"`
	FrontendPort string `mapstructure:"FRONTEND_PORT"`
	BotId        uuid.UUID
	BotName      string `mapstructure:"CHAT_BOT_NAME"`
}

func (a *AppConf) SetBotId(id uuid.UUID) {
	a.BotId = id
}

func LoadConfig() {
	v := viper.New()
	v.SetConfigName("app")
	v.AddConfigPath("conf/")
	//v.SetEnvPrefix("eurovision")
	//v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
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

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(res) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}
