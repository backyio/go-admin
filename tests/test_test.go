package tests

import (
	"testing"
	"time"

	"github.com/backyio/go-admin/modules/config"
	"github.com/backyio/go-admin/tests/frameworks/buffalo"
)

func TestBlackBoxTestSuitOfBuiltInTables(t *testing.T) {

	BlackBoxTestSuitOfBuiltInTables(t, buffalo.NewHandler, config.DatabaseList{
		"default": {
			Host:            "127.0.0.1",
			Port:            "26257",
			User:            "root",
			Pwd:             "",
			Name:            "go-admin-test",
			MaxIdleConns:    50,
			MaxOpenConns:    150,
			ConnMaxLifetime: time.Hour,
			ConnMaxIdleTime: 0,
			Driver:          config.DriverPostgresql,
		},
	})
}
