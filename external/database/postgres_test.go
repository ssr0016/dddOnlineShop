package database

import (
	"onlineShop/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
}

func TestConnectionPostgres(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, err := ConnectPostgres(config.Cfg.DB)
		require.Nil(t, err)
		require.NotNil(t, db)
	})

	t.Run("invalid password", func(t *testing.T) {
		cfg := config.Cfg.DB
		cfg.Password = "invalid password"
		db, err := ConnectPostgres(cfg)
		require.NotNil(t, err)
		require.Nil(t, db)
	})
}
