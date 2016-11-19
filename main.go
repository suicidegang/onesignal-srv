package main

import (
	"log"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/suicidegang/onesignal-srv/handler"
	proto "github.com/suicidegang/onesignal-srv/proto/onesignal"
)

func main() {
	service := micro.NewService(
		micro.Name("sg.micro.srv.onesignal"),
		micro.Version("0.1"),
		micro.Flags(
			cli.StringFlag{
				Name:   "api_key",
				EnvVar: "API_KEY",
				Usage:  "Onesignal auth API key",
			},
			cli.StringFlag{
				Name:   "app_id",
				EnvVar: "APP_ID",
				Usage:  "Onesignal default app ID",
			},
		),
		micro.Action(func(c *cli.Context) {
			if len(c.String("api_key")) > 0 {
				handler.ApiKey = c.String("api_key")
			}

			if len(c.String("app_id")) > 0 {
				handler.AppID = c.String("app_id")
			}
		}),
		micro.BeforeStart(func() error {
			log.Printf("[sg.micro.srv.onesignal] Starting service...")
			log.Printf("[sg.micro.srv.onesignal] API Key: %s", handler.ApiKey)
			log.Printf("[sg.micro.srv.onesignal] APP Id: %s", handler.AppID)
			return nil
		}),
	)

	service.Init()
	proto.RegisterOnesignalHandler(service.Server(), new(handler.Onesignal))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
