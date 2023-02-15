package trader_test

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/viant/demo/app/config"
	"github.com/viant/demo/app/service/trader"
	_ "github.com/viant/dyndb"
	"github.com/viant/toolbox"
	"testing"
)

func TestServiceList(t *testing.T) {
	cfg := &config.Config{}
	cfg.InitTest()
	srv := trader.New(cfg)
	err := srv.Init(context.Background())
	traders, err := srv.List(context.Background())
	assert.Nil(t, err)
	toolbox.DumpIndent(traders, true)
}

func TestServiceByID(t *testing.T) {
	cfg := &config.Config{}
	cfg.InitTest()
	srv := trader.New(cfg)
	err := srv.Init(context.Background())
	trader, err := srv.ByID(context.Background(), 3)
	assert.Nil(t, err)
	toolbox.DumpIndent(trader, true)
}
