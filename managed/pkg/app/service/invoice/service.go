package invoice

import (
	"context"
	"github.com/viant/datly/view"
	"github.com/viant/demo/app/config"
	"github.com/viant/demo/app/domain"
	"github.com/viant/demo/app/service/reader"
	"reflect"
)

const (
	viewID    = "invoice"
	viewTable = "invoice"
)

type Service struct {
	reader *reader.Service
	config *config.Config
}

func (s *Service) ByID(ctx context.Context, id int) (*domain.Invoice, error) {
	var result = make([]*domain.Invoice, 0)
	err := s.reader.ReadWithCriteria(ctx, viewID, &result, "id = ?", id)
	if len(result) == 0 {
		return nil, err
	}
	return result[0], err

}

func (s *Service) List(ctx context.Context) ([]*domain.Invoice, error) {
	var result = make([]*domain.Invoice, 0)
	err := s.reader.Read(ctx, viewID, &result)
	return result, err
}

func (s *Service) Init(ctx context.Context) error {
	conn := s.reader.AddConnector(s.config.DemoDb.Name, s.config.DemoDb.Driver, s.config.DemoDb.DSN)
	invoiceView := view.NewView(viewID, viewTable,
		view.WithConnector(conn),
		view.WithCriteria("id"),
		view.WithViewType(reflect.TypeOf(&domain.Invoice{})),
		view.WithOneToMany("Items", "id",
			view.NwReferenceView("", "invoice_id",
				view.NewView("items", "invoice_list_item", view.WithConnector(conn)))),
	)
	s.reader.AddViews(invoiceView)
	return s.reader.Init(ctx)
}

func New(cfg *config.Config) *Service {
	ret := &Service{config: cfg}
	ret.reader = &reader.Service{}
	return ret
}
