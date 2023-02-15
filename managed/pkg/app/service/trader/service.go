package trader

import (
	"context"
	"github.com/viant/datly/view"
	"github.com/viant/demo/app/config"
	"github.com/viant/demo/app/domain"
	"github.com/viant/demo/app/service/reader"
	"reflect"
)

const (
	viewID    = "trader"
	viewTable = "trader"
)

type Service struct {
	reader *reader.Service
	config *config.Config
}

func (s *Service) ByID(ctx context.Context, id int) (*domain.Trader, error) {
	var result = make([]*domain.Trader, 0)
	err := s.reader.ReadWithCriteria(ctx, viewID, &result, "id = ?", id)
	if len(result) == 0 {
		return nil, err
	}
	return result[0], err

}

func (s *Service) List(ctx context.Context) ([]*domain.Trader, error) {
	var result = make([]*domain.Trader, 0)
	err := s.reader.Read(ctx, viewID, &result)
	return result, err
}

func (s *Service) Init(ctx context.Context) error {
	demoConn := s.reader.AddConnector(s.config.DemoDb.Name, s.config.DemoDb.Driver, s.config.DemoDb.DSN)
	aclConn := s.reader.AddConnector(s.config.AclDb.Name, s.config.AclDb.Driver, s.config.AclDb.DSN)

	aView := view.NewView(viewID, viewTable,
		view.WithConnector(demoConn),
		view.WithCriteria("id"),
		view.WithViewType(reflect.TypeOf(&domain.Trader{})),
		view.WithOneToOne("Acl", "id",
			view.NwReferenceView("UserId", "USER_ID",
				view.NewView("trader_acl", "USER_ACL",
					view.WithSQL(`SELECT
                          USER_ID,
                          ARRAY_EXISTS(ROLE, 'READ_ONLY') AS IS_READONLY,
                          ARRAY_EXISTS(PERMISSION, 'FEATURE1') AS CAN_USE_FEATURE1
                    FROM USER_ACL `),
					view.WithConnector(aclConn)))),
	)
	s.reader.AddViews(aView)
	return s.reader.Init(ctx)
}

func New(cfg *config.Config) *Service {
	ret := &Service{config: cfg}
	ret.reader = &reader.Service{}
	return ret
}
