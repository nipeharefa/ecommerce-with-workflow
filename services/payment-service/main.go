package main

import (
	"log"
	"strings"

	boyo "github.com/nipeharefa/boyo-service"
	"github.com/spf13/viper"
)

var cfg *viper.Viper

func init() {
	cfg = viper.New()

	cfg.AutomaticEnv()
	cfg.SetConfigType("yaml")
	replacer := strings.NewReplacer(".", "_")
	cfg.SetEnvKeyReplacer(replacer)

	cfg.SetConfigFile(`config.yaml`)
	err := cfg.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
}

func main() {

	svc := boyo.NewBoyoService("payment-service", cfg)

	doWorkflow()
	registerRouter(svc)

	svc.Run()

}

func registerRouter(svc boyo.Service) {

	r := svc.GetRouter()

	r.GET("/", Home)
}
