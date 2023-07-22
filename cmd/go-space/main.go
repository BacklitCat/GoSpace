package main

import (
	"GoSpace/config"
	"GoSpace/internal/client"
	"github.com/rs/zerolog/log"
)

func main() {
	start()
}

func start() {
	coreInit()

}

func coreInit() {
	if err := config.ConfManager.InitConf("./deployments"); err != nil {
		log.Fatal().Err(err).Msg("coreInit: config.ConfManager.InitConf() failed")
		return
	}

	if err := client.ClientManager.InitClient(); err != nil {
		log.Fatal().Err(err).Msg("coreInit: client.ClientManager.InitClient() failed")
		return
	}
}
