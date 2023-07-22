package client

import (
	"GoSpace/config"
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/rs/zerolog/log"
)

type PostgreSQLClient struct {
	DB     *pg.DB
	Option *pg.Options
}

func (p *PostgreSQLClient) Init(conf *config.ConfPgSQL) error {
	p.Option = &pg.Options{
		Addr:     conf.Addr,
		User:     conf.User,
		Password: conf.Pwd,
		Database: conf.DB,
	}
	p.DB = pg.Connect(p.Option)
	err := p.DB.Ping(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("PostgreSQLClient.Init: handle connection error")
	} else {
		log.Info().Str("postgresql_dial_time_out", p.Option.DialTimeout.String()).
			Str("postgresql_read_time_out", p.Option.ReadTimeout.String()).
			Str("postgresql_write_time_out", p.Option.WriteTimeout.String()).
			Int("postgresql_max_retries", p.Option.MaxRetries).
			Bool("postgresql_retry_statement_timeout", p.Option.RetryStatementTimeout).
			Int("postgresql_pool_size", p.Option.PoolSize).
			Int("postgresql_pool_size", p.Option.PoolSize).
			Msg("PostgreSQLClient.Init: connection is successful")
	}
	return nil
}
