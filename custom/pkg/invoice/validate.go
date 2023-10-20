package invoice

import (
	"fmt"
	"github.com/awitas/myapp/shared"
	"github.com/viant/govalidator"
	"time"
)

func (i *Invoice) Validate(cur *Invoice, acl *Acl, products []*Product) *shared.Validation {
	info := shared.NewValidationInfo()
	if !i.ensureAuthorized(acl, info) {
		return info
	}
	Items(i.Items).Validate(info, products)
	isInsert := cur == nil
	if isInsert {
		i.validateForInsert(info)
	} else { //update validation
		i.validateForUpdate(info, cur)
	}
	return info
}

func (i *Invoice) ensureAuthorized(acl *Acl, info *shared.Validation) bool {
	return true
}

func (i *Invoice) validateForInsert(info *shared.Validation) {

}

func (i *Invoice) validateForUpdate(info *shared.Validation, cur *Invoice) {
	info.Validate(i, govalidator.WithShallow(true), govalidator.WithSetMarker())
	if i.Has.DueDate && i.DueDate != nil {
		if i.DueDate.Before(time.Now()) {
			info.Validation.AddViolation("DueDate", i.DueDate, "inThePast",
				fmt.Sprintf("dueDate can not be in the past %v", i.DueDate))
		}
	}
}
