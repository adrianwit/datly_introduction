package audience

import (
	"context"
	"github.com/viant/datly/view"
	"github.com/viant/demo/app/config"
	"github.com/viant/demo/app/domain"
	"github.com/viant/demo/app/service/reader"
	"reflect"
)

const (
	viewID    = "audience"
	viewTable = "AUDIENCE"
)

type Service struct {
	reader *reader.Service
	config *config.Config
}

func (s *Service) ByID(ctx context.Context, id int) (*domain.Audience, error) {
	var result = make([]*domain.Audience, 0)
	err := s.reader.ReadWithCriteria(ctx, viewID, &result, "id = ?", id)
	if len(result) == 0 {
		return nil, err
	}
	return result[0], err

}

func (s *Service) List(ctx context.Context) ([]*domain.Audience, error) {
	var result = make([]*domain.Audience, 0)
	err := s.reader.Read(ctx, viewID, &result)
	return result, err
}

func (s *Service) Init(ctx context.Context) error {
	conn := s.reader.AddConnector(s.config.DemoDb.Name, s.config.DemoDb.Driver, s.config.DemoDb.DSN)
	invoiceView := view.NewView(viewID, viewTable,
		view.WithConnector(conn),
		view.WithCriteria("ID"),
		view.WithViewType(reflect.TypeOf(&domain.Audience{})),
		view.WithOneToMany("Deals", "deal_ids",
			view.NwReferenceView("ID", "id",
				view.NewView("deal", "DEAL", view.WithConnector(conn)))),
	)
	s.reader.AddViews(invoiceView)
	return s.reader.Init(ctx)
}

func New(cfg *config.Config) *Service {
	ret := &Service{config: cfg}
	ret.reader = &reader.Service{}
	return ret
}
