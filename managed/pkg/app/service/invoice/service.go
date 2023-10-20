package invoice

import (
	"context"
	_ "embed"
	"github.com/viant/datly"
	"github.com/viant/datly/service/executor"
	"github.com/viant/datly/service/reader"
	"github.com/viant/datly/view"
	"github.com/viant/datly/view/state"
	"github.com/viant/demo/app/config"
	"github.com/viant/demo/app/domain"
	"reflect"
)

const (
	viewID    = "invoice"
	viewTable = "invoice"

	viewInsertID     = "invoice_insert"
	invoiceParamName = "Invoice"
)

type Service struct {
	*datly.Service
	config *config.Config
}

func (s *Service) ByID(ctx context.Context, id int) (*domain.Invoice, error) {
	var result = make([]*domain.Invoice, 0)
	err := s.Read(ctx, viewID, &result, reader.WithCriteria("id = ?", id))
	if len(result) == 0 {
		return nil, err
	}
	return result[0], err
}
func (s *Service) List(ctx context.Context) ([]*domain.Invoice, error) {
	var result = make([]*domain.Invoice, 0)
	err := s.Read(ctx, viewID, &result)
	return result, err
}

func (s *Service) initReaderView(ctx context.Context) error {
	err := s.AddConnectors(ctx, view.NewConnector(s.config.DemoDb.Name, s.config.DemoDb.Driver, s.config.DemoDb.DSN))
	if err != nil {
		return err
	}
	invoiceView := view.NewView(viewID, viewTable,
		view.WithCriteria("id"),
		view.WithViewType(reflect.TypeOf(&domain.Invoice{})),
		view.WithOneToMany("Items", "id", //relation can be also defined with datly tag
			view.NwReferenceView("", "invoice_id",
				view.NewView("items", "invoice_list_item"))),
	)
	return s.AddViews(ctx, invoiceView)
}

func (s *Service) Init(ctx context.Context) error {
	if err := s.initReaderView(ctx); err != nil {
		return err
	}
	if err := s.initInsertView(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Service) Insert(ctx context.Context, invoices ...*Invoice) error {
	return s.Exec(ctx, viewInsertID, executor.WithParameter(invoiceParamName, invoices))
}

//go:embed tmpl/insert.vm
var insertSQL string

func (s *Service) initInsertView(ctx context.Context) error {
	err := s.AddConnectors(ctx, view.NewConnector(s.config.DemoDb.Name, s.config.DemoDb.Driver, s.config.DemoDb.DSN))
	if err != nil {
		return err
	}
	invoiceInsertView := view.NewView(viewInsertID, viewTable,
		view.WithViewKind(view.ModeExec),
		view.WithTemplate(
			view.NewTemplate(
				insertSQL,
				view.WithTemplateParameter(state.NewParameter(
					invoiceParamName,
					state.NewBodyLocation("Entity"),
					state.WithParameterType(reflect.TypeOf([]*Invoice{})))))))

	return s.AddViews(ctx, invoiceInsertView)
}

func New(cfg *config.Config) *Service {
	srv, _ := datly.New(context.Background())
	return &Service{config: cfg, Service: srv}
}
