package audience

import (
	"context"
	"github.com/viant/datly/reader"
	"github.com/viant/datly/view"
	"github.com/viant/demo/app/config"
	"github.com/viant/demo/app/domain"
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
	err := s.reader.ReadInto(ctx, viewID, &result, reader.WithCriteria("id = ?", id))
	if len(result) == 0 {
		return nil, err
	}
	return result[0], err

}

func (s *Service) List(ctx context.Context) ([]*domain.Audience, error) {
	var result = make([]*domain.Audience, 0)
	err := s.reader.ReadInto(ctx, viewID, &result)
	return result, err
}

func (s *Service) Init(ctx context.Context) error {
	conn := s.reader.Resource.AddConnector(s.config.DemoDb.Name, s.config.DemoDb.Driver, s.config.DemoDb.DSN)
	aView := view.NewView(viewID, viewTable,
		view.WithConnector(conn),
		view.WithCriteria("ID"),
		view.WithViewType(reflect.TypeOf(&domain.Audience{})),
		view.WithOneToMany("Deals", "deal_ids",
			view.NwReferenceView("ID", "id",
				view.NewView("deal", "DEAL", view.WithConnector(conn)))),
	)
	s.reader.Resource.AddViews(aView)
	return s.reader.Resource.Init(ctx)
}

func New(cfg *config.Config) *Service {
	ret := &Service{config: cfg}
	ret.reader = &reader.Service{}
	return ret
}
