package product_test

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	_ "github.com/viant/bigquery"
	"github.com/viant/demo/app/config"
	"github.com/viant/demo/app/service/product"
	"github.com/viant/toolbox"
	"testing"
)

func TestServiceList(t *testing.T) {
	cfg := &config.Config{}
	cfg.InitTest()
	srv := product.New(cfg)
	err := srv.Init(context.Background())
	if !assert.Nil(t, err) {
		return
	}
	products, err := srv.ListWithPeriod(context.Background(), "today")
	if !assert.Nil(t, err) {
		return
	}
	assert.True(t, len(products) > 0)
	toolbox.DumpIndent(products, true)
}

func TestServiceByID(t *testing.T) {
	cfg := &config.Config{}
	cfg.InitTest()
	srv := product.New(cfg)
	err := srv.Init(context.Background())
	product, err := srv.ByID(context.Background(), 3)
	assert.Nil(t, err)
	toolbox.DumpIndent(product, true)
}
