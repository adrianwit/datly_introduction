package product

import (
	"context"
	_ "embed"
	"github.com/viant/datly"
	"github.com/viant/datly/service/reader"
	"github.com/viant/datly/view"
	"github.com/viant/datly/view/state"
	"github.com/viant/demo/app/config"
	"github.com/viant/demo/app/domain"
	"reflect"
)

const (
	viewID    = "product"
	viewTable = "PRODUCT"
)

type Service struct {
	*datly.Service
	config *config.Config
}

func (s *Service) ByID(ctx context.Context, id int) (*domain.Product, error) {
	var result = make([]*domain.Product, 0)
	err := s.Read(ctx, viewID, &result, reader.WithCriteria("id = ?", id))
	if len(result) == 0 {
		return nil, err
	}
	return result[0], err

}
func (s *Service) List(ctx context.Context) ([]*domain.Product, error) {
	var result = make([]*domain.Product, 0)
	err := s.Read(ctx, viewID, &result)
	return result, err
}

func (s *Service) ListWithPeriod(ctx context.Context, period string) ([]*domain.Product, error) {
	var result = make([]*domain.Product, 0)
	err := s.Read(ctx, viewID, &result, reader.WithParameter("performance:period", period))
	return result, err
}

//go:embed tmpl/performance.vm
var performanceVQL string

func (s *Service) Init(ctx context.Context) error {
	demoConn, err := s.AddConnector(ctx, s.config.DemoDb.Name, s.config.DemoDb.Driver, s.config.DemoDb.DSN)
	if err != nil {
		return err
	}
	bqDevConn, err := s.AddConnector(ctx, s.config.BqDev.Name, s.config.BqDev.Driver, s.config.BqDev.DSN)
	if err != nil {
		return err
	}
	aView := s.initView(demoConn, bqDevConn)
	if err = s.AddViews(ctx, aView); err != nil {
		return err
	}
	return nil
}

func (s *Service) initView(demoConn *view.Connector, bqDevConn *view.Connector) *view.View {
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
						view.WithTemplateParameter(state.NewParameter("period",
							state.NewQueryLocation("period"),
							state.WithParameterType(reflect.TypeOf("")),
						)))),
					view.WithConnector(bqDevConn)))),
	)
	return aView
}

func New(cfg *config.Config) *Service {
	srv, _ := datly.New(context.Background())
	ret := &Service{config: cfg, Service: srv}
	return ret
}
