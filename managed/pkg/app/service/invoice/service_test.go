package invoice_test

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/viant/demo/app/config"
	"github.com/viant/demo/app/service/invoice"
	"github.com/viant/toolbox"
	"testing"
)

func TestServiceList(t *testing.T) {
	cfg := &config.Config{}
	cfg.InitTest()
	srv := invoice.New(cfg)
	err := srv.Init(context.Background())
	invoices, err := srv.List(context.Background())
	assert.Nil(t, err)
	toolbox.DumpIndent(invoices, true)
}

func TestServiceByID(t *testing.T) {
	cfg := &config.Config{}
	cfg.InitTest()
	srv := invoice.New(cfg)
	err := srv.Init(context.Background())
	invoice, err := srv.ByID(context.Background(), 3)
	assert.Nil(t, err)
	toolbox.DumpIndent(invoice, true)
}
