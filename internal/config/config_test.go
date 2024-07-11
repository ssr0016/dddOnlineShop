package config

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoandConfig(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		filename := "../../cmd/api/config.yaml"
		err := LoadConfig(filename)

		require.Nil(t, err)
		log.Printf("config: %v", Cfg)
	})

	t.Run("file not exist", func(t *testing.T) {
		filename := "config.yaml"
		err := LoadConfig(filename)

		require.NotNil(t, err)
		log.Printf("%v\n", err)
	})
}
