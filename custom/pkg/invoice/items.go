package invoice

import (
	"github.com/adrianwit/datly_introduction/custom/pkg/shared"
	"time"
)

type Items []*Item

func (i Items) Init(now *time.Time, act *Acl) {

}

func (i Items) Total(discount *Discount) float64 {
	return 0
}

func (i Items) Merge(items []*Item) Items {
	return i
}

func (i Items) Validate(info *shared.Validation, products []*Product) {

}
