package audience_test

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/viant/demo/app/config"
	"github.com/viant/demo/app/service/audience"
	"github.com/viant/toolbox"
	"testing"
)

func TestServiceList(t *testing.T) {
	cfg := &config.Config{}
	cfg.InitTest()
	srv := audience.New(cfg)
	err := srv.Init(context.Background())
	items, err := srv.List(context.Background())
	if !assert.Nil(t, err) {
		return
	}
	toolbox.DumpIndent(items, true)
}

func TestServiceByID(t *testing.T) {
	cfg := &config.Config{}
	cfg.InitTest()
	srv := audience.New(cfg)
	err := srv.Init(context.Background())
	item, err := srv.ByID(context.Background(), 3)
	if !assert.Nil(t, err) {
		return
	}
	toolbox.DumpIndent(item, true)
}
