package invoice_test

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/viant/demo/app/config"
	"github.com/viant/demo/app/service/invoice"
	_ "github.com/viant/sqlx/metadata/product/mysql"
	"github.com/viant/toolbox"
	"testing"
)

func TestServiceList(t *testing.T) {
	cfg := &config.Config{}
	cfg.InitTest()
	srv := invoice.New(cfg)
	err := srv.Init(context.Background())
	if !assert.Nil(t, err) {
		return
	}
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

func TestServiceInsert(t *testing.T) {
	cfg := &config.Config{}
	cfg.InitTest()
	srv := invoice.New(cfg)
	err := srv.Init(context.Background())
	if !assert.Nil(t, err) {
		return
	}
	name := "John"
	totalAmount := 125.5
	invoices := []*invoice.Invoice{
		{
			CustomerName: &name,
			TotalAmount:  &totalAmount,
			Has:          &invoice.InvoiceHas{},
		},
	}

	err = srv.Insert(context.Background(), invoices...)
	assert.Nil(t, err)

	for _, anInvoice := range invoices {
		invoice, err := srv.ByID(context.Background(), anInvoice.Id)
		assert.Nil(t, err)
		toolbox.DumpIndent(invoice, true)
	}
}
