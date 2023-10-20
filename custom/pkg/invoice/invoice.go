package invoice

import (
	"github.com/awitas/myapp/checksum"
	"github.com/viant/xdatly/types/core"
	"reflect"
	"time"
)

var PackageName = "invoice"

func init() {
	core.RegisterType(PackageName, "Entity", reflect.TypeOf(Entity{}), checksum.GeneratedTime)
	core.RegisterType(PackageName, "Invoice", reflect.TypeOf(Invoice{}), checksum.GeneratedTime)
}

type Entity struct {
	Entity []*Invoice `typeName:"Invoice"`
}

type Invoice struct {
	Id           int         `sqlx:"name=id,autoincrement,primaryKey,required"`
	CustomerName *string     `sqlx:"name=customer_name"   json:",omitempty"`
	CustomerId   *int        `sqlx:"name=customer_id"      json:",omitempty" validate:"required"`
	InvoiceDate  *time.Time  `sqlx:"name=invoice_date"     json:",omitempty" validate:"required"`
	Status       *int        `sqlx:"name=status"           json:",omitempty"`
	DueDate      *time.Time  `sqlx:"name=due_date"         json:",omitempty" validate:"required"`
	Total        *float64    `sqlx:"name=total_amount"     json:",omitempty" `
	DiscountCode *string     `sqlx:"name=discount_code,refTable=DISCOUNT,refColumn=code" json:",omitempty"`
	Created      *time.Time  `sqlx:"name=created"          json:",omitempty"`
	UserCreated  *int        `sqlx:"name=user_created,refTable=USER,refColumn=id" json:",omitempty"`
	Updated      *time.Time  `sqlx:"name=updated"          json:",omitempty"`
	UserUpdated  *int        `sqlx:"name=user_updated,refTable=USER,refColumn=id" json:",omitempty"`
	Items        []*Item     `typeName:"Items"   sqlx:"-"  datly:"relName=item,relColumn=id,refColumn=invoice_id,refTable=INVOICE_LIST_ITEM" sql:"SELECT * FROM INVOICE_LIST_ITEM"`
	Discount     *Discount   `typeName:"Discount" sqlx:"-" datly:"relName=discount,relColumn=discount_code,refColumn=code,refTable=DISCOUNT" sql:"SELECT * FROM DISCOUNT"`
	Has          *InvoiceHas `setMarker:"true"             typeName:"InvoiceHas" json:"-"  sqlx:"-" `
}

type Item struct {
	Id          int        `sqlx:"name=id,autoincrement,primaryKey,required"`
	InvoiceId   *int       `sqlx:"name=invoice_id,refTable=INVOICE,refColumn=id" json:",omitempty"`
	ProductId   *int       `sqlx:"name=product_id,refTable=PRODUCT,refColumn=id" json:",omitempty"`
	ProductName *string    `sqlx:"name=product_name" json:",omitempty"`
	Quantity    *int       `sqlx:"name=quantity" json:",omitempty"`
	Price       *float64   `sqlx:"name=price" json:",omitempty"`
	Total       *float64   `sqlx:"name=total" json:",omitempty"`
	Created     *time.Time `sqlx:"name=created" json:",omitempty"`
	UserCreated *int       `sqlx:"name=user_created,refTable=USER,refColumn=id" json:",omitempty"`
	Updated     *time.Time `sqlx:"name=updated" json:",omitempty"`
	UserUpdated *int       `sqlx:"name=user_updated,refTable=USER,refColumn=id" json:",omitempty"`
	Has         *ItemHas   `setMarker:"true" typeName:"ItemHas" json:"-"  sqlx:"-" `
}

type ItemHas struct {
	Id          bool
	InvoiceId   bool
	ProductId   bool
	ProductName bool
	Quantity    bool
	Price       bool
	Total       bool
	Created     bool
	UserCreated bool
	Updated     bool
	UserUpdated bool
}

type Product struct {
	Id     int         `sqlx:"name=ID,required"`
	Status int         `sqlx:"name=STATUS,required"`
	Has    *ProductHas `setMarker:"true" typeName:"ProductHas" json:"-"  sqlx:"-" `
}

type ProductHas struct {
	Id     bool
	Status bool
}

type Discount struct {
	Code      string       `sqlx:"name=code,primaryKey,required"`
	Pct       float64      `sqlx:"name=pct" json:",omitempty"`
	StartDate *time.Time   `sqlx:"name=start_date" json:",omitempty"`
	EndDate   *time.Time   `sqlx:"name=end_date" json:",omitempty"`
	Has       *DiscountHas `setMarker:"true" typeName:"DiscountHas" json:"-"  sqlx:"-" `
}

type DiscountHas struct {
	Code      bool
	StartDate bool
	EndDate   bool
}

type Acl struct {
	UserId   int     `sqlx:"name=USER_ID,required"`
	IsAdmin  int     `sqlx:"name=IS_ADMIN,required"`
	IsViewer int     `sqlx:"name=IS_VIEWER,required"`
	Has      *AclHas `setMarker:"true" typeName:"AclHas" json:"-"  sqlx:"-" `
}

type AclHas struct {
	UserId   bool
	IsAdmin  bool
	IsViewer bool
}

type Features struct {
	UserId         int          `sqlx:"name=USER_ID,required"`
	CanSetDiscount bool         `sqlx:"name=CAN_SET_DISCOUNT,required"`
	Has            *FeaturesHas `setMarker:"true" typeName:"FeaturesHas" json:"-"  sqlx:"-" `
}

type FeaturesHas struct {
	UserId         bool
	CanSetDiscount bool
}

type InvoiceHas struct {
	Id           bool
	CustomerName bool
	CustomerId   bool
	InvoiceDate  bool
	Status       bool
	DueDate      bool
	TotalAmount  bool
	DiscountCode bool
	Created      bool
	UserCreated  bool
	Updated      bool
	UserUpdated  bool
}
