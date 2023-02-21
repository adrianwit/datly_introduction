package invoice

import (
	"context"
	"github.com/viant/datly/executor"
	"github.com/viant/datly/view"
	"github.com/viant/demo/app/config"
	"github.com/viant/demo/app/domain"
	mExecutor "github.com/viant/demo/app/service/executor"
	"github.com/viant/demo/app/service/reader"
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
	executor *mExecutor.Service
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

func (s *Service) Insert(ctx context.Context, invoices ...*Invoice) error {
	aView, err := s.executor.View(viewInsertID)
	if err != nil {
		return err
	}

	selectors := &view.Selectors{Index: map[string]*view.Selector{}}
	if err = aView.SetParameter(bodyParamName, selectors, invoices); err != nil {
		return err
	}

	session, err := executor.NewSession(selectors, aView)
	if err != nil {
		return err
	}

	_, err = s.executor.ExecSession(ctx, session)
	return err
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

	s.reader.AddViews(invoiceView)
	s.executor.AddViews(invoiceInsertView)
	if err := s.reader.Init(ctx); err != nil {
		return err
	}

	if err := s.executor.Init(ctx); err != nil {
		return err
	}

	return nil
}

func New(cfg *config.Config) *Service {
	ret := &Service{config: cfg}
	ret.reader = &reader.Service{}
	ret.executor = &mExecutor.Service{}
	return ret
}
