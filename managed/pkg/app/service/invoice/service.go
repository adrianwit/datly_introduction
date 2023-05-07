package invoice

import (
	"context"
	"github.com/viant/datly/executor"
	"github.com/viant/datly/reader"
	"github.com/viant/datly/view"
	"github.com/viant/demo/app/config"
	"github.com/viant/demo/app/domain"
	"reflect"
)

const (
	viewID    = "invoice"
	viewTable = "invoice"

	viewInsertID  = "invoice_insert"
	bodyParamName = "Invoice"
)

type Service struct {
	reader   *reader.Service
	config   *config.Config
	executor *executor.Executor
}

func (s *Service) ByID(ctx context.Context, id int) (*domain.Invoice, error) {
	var result = make([]*domain.Invoice, 0)
	err := s.reader.ReadInto(ctx, viewID, &result, reader.WithCriteria("id = ?", id))
	if len(result) == 0 {
		return nil, err
	}
	return result[0], err
}

func (s *Service) List(ctx context.Context) ([]*domain.Invoice, error) {
	var result = make([]*domain.Invoice, 0)
	err := s.reader.ReadInto(ctx, viewID, &result)
	return result, err
}

func (s *Service) Insert(ctx context.Context, invoices ...*Invoice) error {
	aView, err := s.reader.Resource.View(viewInsertID)
	if err != nil {
		return err
	}
	return s.executor.Execute(ctx, aView, executor.WithParameter(bodyParamName, invoices))
}

func (s *Service) Init(ctx context.Context) error {
	conn := s.reader.Resource.AddConnector(s.config.DemoDb.Name, s.config.DemoDb.Driver, s.config.DemoDb.DSN)
	invoiceView := view.NewView(viewID, viewTable,
		view.WithConnector(conn),
		view.WithCriteria("id"),
		view.WithViewType(reflect.TypeOf(&domain.Invoice{})),
		//This is optional since go struct defines datly rel/ref tag
		//view.WithOneToMany("Items", "id",
		//	view.NwReferenceView("", "invoice_id",
		//		view.NewView("items", "invoice_list_item", view.WithConnector(conn)))),
	)

	insertSQL := `$sequencer.Allocate("invoice", $Unsafe.Invoice, "Id")
$sequencer.Allocate("invoice_list_item", $Unsafe.Invoice, "Item/Id")
#foreach($recInvoice in $Unsafe.Invoice)
#if($recInvoice)
    $sql.Insert($recInvoice, "invoice");
      #foreach($recItem in $recInvoice.Item)
      #if($recItem)
          #set($recItem.InvoiceId = $recInvoice.Id)
          $sql.Insert($recItem, "invoice_list_item");
      #end
      #end
#end
#end`
	invoiceInsertView := view.NewView(viewInsertID, viewTable,
		view.WithViewKind(view.SQLExecMode),
		view.WithConnector(conn),
		view.WithCriteria("id"),
		view.WithTemplate(
			view.NewTemplate(
				insertSQL,
				view.WithTemplateParameter(view.NewParameter(
					bodyParamName,
					view.NewBodyLocation(""),
					view.WithParameterType(reflect.TypeOf([]*Invoice{})))),
			),
		),
	)
	s.reader.Resource.AddViews(invoiceView)
	s.reader.Resource.AddViews(invoiceInsertView)
	if err := s.reader.Resource.Init(ctx); err != nil {
		return err
	}
	return nil
}

func New(cfg *config.Config) *Service {
	ret := &Service{config: cfg}
	ret.reader = reader.New()
	ret.executor = executor.New()
	return ret
}
