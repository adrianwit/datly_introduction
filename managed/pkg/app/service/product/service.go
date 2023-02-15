package product

import (
	"context"
	_ "embed"
	dreader "github.com/viant/datly/reader"
	"github.com/viant/datly/view"
	"github.com/viant/demo/app/config"
	"github.com/viant/demo/app/domain"
	"github.com/viant/demo/app/service/reader"
	"reflect"
)

const (
	viewID    = "product"
	viewTable = "PRODUCT"
)

type Service struct {
	reader *reader.Service
	config *config.Config
}

func (s *Service) ByID(ctx context.Context, id int) (*domain.Product, error) {
	var result = make([]*domain.Product, 0)
	err := s.reader.ReadWithCriteria(ctx, viewID, &result, "id = ?", id)
	if len(result) == 0 {
		return nil, err
	}
	return result[0], err

}
func (s *Service) List(ctx context.Context) ([]*domain.Product, error) {
	var result = make([]*domain.Product, 0)
	err := s.reader.Read(ctx, viewID, &result)
	return result, err
}

func (s *Service) ListWithPeriod(ctx context.Context, period string) ([]*domain.Product, error) {
	var result = make([]*domain.Product, 0)
	err := s.reader.ReadWithSession(ctx, viewID, &result, func(aView *view.View, session *dreader.Session) error {
		return aView.SetParameter("performance:period", session.Selectors, period)
	})
	return result, err
}

//go:embed tmpl/performance.vm
var performanceVQL string

func (s *Service) Init(ctx context.Context) error {
	demoConn := s.reader.AddConnector(s.config.DemoDb.Name, s.config.DemoDb.Driver, s.config.DemoDb.DSN)
	bqDevConn := s.reader.AddConnector(s.config.BqDev.Name, s.config.BqDev.Driver, s.config.BqDev.DSN)

	aView := view.NewView(viewID, viewTable,
		view.WithConnector(demoConn),
		view.WithCriteria("ID"),
		view.WithViewType(reflect.TypeOf(&domain.Product{})),
		view.WithOneToOne("Vendor", "VENDOR_ID",
			view.NwReferenceView("ID", "ID",
				view.NewView("product_vendor", "VENDOR",
					view.WithConnector(demoConn)))),
		view.WithOneToMany("Performance", "ID",
			view.NwReferenceView("ProductId", "product_id",
				view.NewView("performance", "product_performance",
					view.WithTemplate(view.NewTemplate(performanceVQL,
						view.WithTemplateParameter(view.NewParameter("period",
							view.NewQueryLocation("period"),
							view.WithParameterType(reflect.TypeOf("")),
						)))),
					view.WithConnector(bqDevConn)))),
	)
	s.reader.AddViews(aView)
	return s.reader.Init(ctx)
}

func New(cfg *config.Config) *Service {
	ret := &Service{config: cfg}
	ret.reader = &reader.Service{}
	return ret
}
