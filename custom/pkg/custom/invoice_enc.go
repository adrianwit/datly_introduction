package invoice

import (
	"github.com/viant/xdatly/types/core"
	"github.com/viant/xdatly/types/custom/generated"
	"reflect"
	"time"
)

var PackageName = "invoice"

func init() {
	core.RegisterType(PackageName, "Invoice", reflect.TypeOf(Invoice{}), generated.GeneratedTime)
}

type Invoice struct {
	Id           int         `sqlx:"name=id,primaryKey"`
	CustomerName *string     `sqlx:"name=customer_name" json:",omitempty"`
	InvoiceDate  *time.Time  `sqlx:"name=invoice_date" json:",omitempty"`
	DueDate      *time.Time  `sqlx:"name=due_date" json:",omitempty"`
	TotalAmount  *float64    `sqlx:"name=total_amount" json:",omitempty"`
	Item         []*Item     `typeName:"Item" sqlx:"-"`
	Has          *InvoiceHas `presenceIndex:"true" typeName:"InvoiceHas" json:"-" sqlx:"presence=true"`
}

type Item struct {
	Id          int      `sqlx:"name=id,primaryKey"`
	InvoiceId   *int     `sqlx:"name=invoice_id" json:",omitempty"`
	ProductName *string  `sqlx:"name=product_name" json:",omitempty"`
	Quantity    *int     `sqlx:"name=quantity" json:",omitempty"`
	Price       *float64 `sqlx:"name=price" json:",omitempty"`
	Total       *float64 `sqlx:"name=total" json:",omitempty"`
	Has         *ItemHas `presenceIndex:"true" typeName:"ItemHas" json:"-" sqlx:"presence=true"`
}

type ItemHas struct {
	Id          bool
	InvoiceId   bool
	ProductName bool
	Quantity    bool
	Price       bool
	Total       bool
}

type InvoiceHas struct {
	Id           bool
	CustomerName bool
	InvoiceDate  bool
	DueDate      bool
	TotalAmount  bool
}
