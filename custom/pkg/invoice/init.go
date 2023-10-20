package invoice

import (
	"github.com/awitas/myapp/shared"
	"time"
)

const (
	DefaultDueDuration = time.Hour * 24 * 30 //30 days
)

func (i *Invoice) Init(cur *Invoice, acl *Acl, features *Features, discount *Discount) bool {
	i.Discount = discount
	isInsert := cur == nil
	now := shared.TimePtr(time.Now())
	i.excludeInactiveFeature(features)
	if isInsert {
		i.initialiseForInsert(now, acl)
	} else {
		i.initialiseForUpdate(now, acl, cur)
	}
	Items(i.Items).Init(now, acl)
	return true
}

func (i *Invoice) excludeInactiveFeature(features *Features) {
	if !features.CanSetDiscount {
		i.DiscountCode = nil
		i.Has.DiscountCode = false
		i.Discount = nil
	}
}

func (i *Invoice) initialiseForInsert(now *time.Time, acl *Acl) {
	i.UserCreated = &acl.UserId
	i.Created = now
	if i.Status == nil {
		i.Status = shared.IntPtr(1)
	}
	if i.DueDate == nil {
		i.DueDate = shared.TimePtr(now.Add(DefaultDueDuration))
	}
	i.Total = shared.Float64Ptr(Items(i.Items).Total(i.Discount))
}

func (i *Invoice) initialiseForUpdate(now *time.Time, acl *Acl, cur *Invoice) {
	i.UserUpdated = &acl.UserId
	i.Has.UserUpdated = true
	i.Updated = now
	i.Has.Updated = true
	items := Items(i.Items).Merge(cur.Items)
	i.Total = shared.Float64Ptr(items.Total(i.Discount))
}
