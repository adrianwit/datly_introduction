package audience

import (
	"context"
	"github.com/viant/datly"
	"github.com/viant/datly/service/reader"
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
	*datly.Service
	config *config.Config
}

func (s *Service) ByID(ctx context.Context, id int) (*domain.Audience, error) {
	var result = make([]*domain.Audience, 0)
	err := s.Read(ctx, viewID, &result, reader.WithCriteria("id = ?", id))
	if len(result) == 0 {
		return nil, err
	}
	return result[0], err

}

func (s *Service) List(ctx context.Context) ([]*domain.Audience, error) {
	var result = make([]*domain.Audience, 0)
	err := s.Read(ctx, viewID, &result)
	return result, err
}

func (s *Service) Init(ctx context.Context) error {
	err := s.AddConnectors(ctx, view.NewConnector(s.config.DemoDb.Name, s.config.DemoDb.Driver, s.config.DemoDb.DSN))
	if err != nil {
		return err
	}
	aView := view.NewView(viewID, viewTable,
		view.WithCriteria("ID"),
		view.WithViewType(reflect.TypeOf(&domain.Audience{})),
		view.WithOneToMany("Deals", "deal_ids",
			view.NwReferenceView("ID", "id",
				view.NewView("deal", "DEAL"))),
	)
	return s.AddViews(ctx, aView)
}

func New(config *config.Config) *Service {
	ret := &Service{config: config}
	ret.Service, _ = datly.New(context.Background())
	return ret
}
