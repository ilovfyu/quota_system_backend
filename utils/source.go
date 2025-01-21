package utils

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"time"
)

type ConnectionConfig struct {
	Name    string
	Type    string
	Connect string
}

func TestConnection(config ConnectionConfig) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	switch config.Type {
	case "mysql":
		return config.ConnectionMySQL(ctx)
	case "postgres":
		config.ConnectionPostgres(config.Connect)
	case "doris":
		config.ConnectionDoris(config.Connect)
	case "clickhouse":
		config.ConnectionClickHouse(config.Connect)
	default:
		return false, fmt.Errorf("unknown connection type: %s", config.Type)
	}
	return false, nil
}

func (c ConnectionConfig) ConnectionMySQL(ctx context.Context) (bool, error) {
	cm := cast.ToStringMap(c.Connect)
	fmt.Println(cm["port"])
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cm["username"], cm["password"], cm["host"], cast.ToString(cm["port"]), cm["db"])
	fmt.Println(dsn)
	db, err := sql.Open(c.Type, dsn)
	if err != nil {
		zap.L().Error("failed to connect to mysql", zap.Error(err))
		return false, err
	}
	defer db.Close()
	err = db.PingContext(ctx)
	if err != nil {
		zap.L().Error("failed to ping to mysql", zap.Error(err))
		return false, err
	}
	return true, nil
}

func (c ConnectionConfig) ConnectionPostgres(connect string) {

}

func (c ConnectionConfig) ConnectionClickHouse(connect string) {

}

func (c ConnectionConfig) ConnectionDoris(connect string) {

}
