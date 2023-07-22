package client

import (
	"GoSpace/config"
	"github.com/rs/zerolog/log"
)

type Manager struct {
	PgClient PostgreSQLClient
}

var ClientManager Manager

func (m *Manager) InitClient() error {
	if err := m.PgClient.Init(config.ConfManager.GetPgSqlConf()); err != nil {
		log.Fatal().Err(err).Msg("client.Manager.InitClient failed")
		return err
	}
	return nil
}
