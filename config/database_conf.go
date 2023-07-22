package config

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"os"
)

type ConfDatabase struct {
	ConfPgSQL *ConfPgSQL `yaml:"postgre_sql"`
}

type ConfPgSQL struct {
	Addr string `yaml:"addr"`
	User string `yaml:"user"`
	Pwd  string `yaml:"pwd"`
	DB   string `yaml:"db"`
}

func (c *ConfDatabase) Init(filePath string) error {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(fileBytes, c); err != nil {
		return err
	}
	log.Info().Str("postgresql_addr", c.ConfPgSQL.Addr).
		Str("postgresql_user", c.ConfPgSQL.User).
		Str("postgresql_db", c.ConfPgSQL.DB).
		Msg("config.ConfDatabase.Init: Unmarshal success")
	return nil
}
