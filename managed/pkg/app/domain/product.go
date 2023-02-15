package domain

import "time"

type Product struct {
	Id          int32      `sqlx:"name=ID"`
	Name        *string    `sqlx:"name=NAME"`
	VendorId    *int64     `sqlx:"name=VENDOR_ID"`
	Status      *int64     `sqlx:"name=STATUS"`
	Created     *time.Time `sqlx:"name=CREATED"`
	UserCreated *int64     `sqlx:"name=USER_CREATED"`
	Updated     *time.Time `sqlx:"name=UPDATED"`
	UserUpdated *int64     `sqlx:"name=USER_UPDATED"`
	Vendor      *Vendor
	Performance []*Performance
}

type Vendor struct {
	Id          int32      `sqlx:"name=ID"`
	Name        *string    `sqlx:"name=NAME"`
	AccountId   *int64     `sqlx:"name=ACCOUNT_ID"`
	Created     *time.Time `sqlx:"name=CREATED"`
	UserCreated *int64     `sqlx:"name=USER_CREATED"`
	Updated     *time.Time `sqlx:"name=UPDATED"`
	UserUpdated *int64     `sqlx:"name=USER_UPDATED"`
}

type Performance struct {
	LocationId *int     `sqlx:"name=location_id"`
	ProductId  *int     `sqlx:"name=product_id"`
	Quantity   *float64 `sqlx:"name=quantity"`
	Price      *float64 `sqlx:"name=price"`
}
