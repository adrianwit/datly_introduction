package domain

import "time"

type Trader struct {
	Id          int32      `sqlx:"name=id"`
	FirstName   *string    `sqlx:"name=first_name"`
	LastName    *string    `sqlx:"name=last_name"`
	Email       *string    `sqlx:"name=email"`
	PhoneNumber *string    `sqlx:"name=phone_number"`
	JoinDate    *time.Time `sqlx:"name=join_date"`
	Acl         *Acl
}

type Acl struct {
	UserId         int   `sqlx:"name=USER_ID"`
	IsReadonly     *bool `sqlx:"name=IS_READONLY"`
	CanUseFeature1 *bool `sqlx:"name=CAN_USE_FEATURE1"`
}
