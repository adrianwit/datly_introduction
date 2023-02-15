package reader

import (
	"context"
	"github.com/viant/datly/reader"
	"github.com/viant/datly/view"
)

type Service struct {
	view.Resource
	reader *reader.Service
}

func (r *Service) Init(ctx context.Context) error {
	return r.Resource.Init(context.Background())
}

func (r *Service) AddConnector(name string, driver string, dsn string, opts ...view.ConnectorOption) *view.Connector {
	connector := view.NewConnector(name, driver, dsn, opts...)
	r.AddConnectors(connector)
	return connector
}

func (r *Service) ensureReader() *reader.Service {
	if r.reader != nil {
		return r.reader
	}
	r.reader = reader.New()
	return r.reader
}

func (r *Service) ReadWithCriteria(ctx context.Context, viewName string, dest interface{}, criteria string, params ...interface{}) error {
	aView, err := r.Resource.View(viewName)
	if err != nil {
		return err
	}
	readerService := r.ensureReader()
	session := reader.NewSession(dest, aView)
	session.AddCriteria(aView, criteria, params...)
	return readerService.Read(ctx, session)
}

func (r *Service) Read(ctx context.Context, viewName string, dest interface{}) error {
	aView, err := r.Resource.View(viewName)
	if err != nil {
		return err
	}
	readerService := r.ensureReader()
	session := reader.NewSession(dest, aView)
	return readerService.Read(ctx, session)
}

func (r *Service) ReadWithSession(ctx context.Context, viewName string, dest interface{}, setSession func(aView *view.View, session *reader.Session) error) error {
	aView, err := r.Resource.View(viewName)
	if err != nil {
		return err
	}
	readerService := r.ensureReader()
	session := reader.NewSession(dest, aView)
	if err = setSession(aView, session); err != nil {
		return err
	}
	return readerService.Read(ctx, session)
}
