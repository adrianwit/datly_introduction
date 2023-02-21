package executor

import (
	"context"
	"github.com/viant/datly/executor"
	"github.com/viant/datly/view"
	"github.com/viant/velty/est"
	"sync"
)

type Service struct {
	view.Resource
	executor *executor.Executor
}

func (r *Service) Init(ctx context.Context) error {
	return r.Resource.Init(context.Background())
}

func (r *Service) AddConnector(name string, driver string, dsn string, opts ...view.ConnectorOption) *view.Connector {
	connector := view.NewConnector(name, driver, dsn, opts...)
	r.AddConnectors(connector)
	return connector
}

func (r *Service) ensureExecutor() *executor.Executor {
	if r.executor != nil {
		return r.executor
	}

	r.executor = executor.New()
	return r.executor
}

func (r *Service) ExecWithSelectors(ctx context.Context, viewName string, selectors *view.Selectors) (*est.State, error) {
	aView, err := r.Resource.View(viewName)
	if err != nil {
		return nil, err
	}

	session, err := executor.NewSession(selectors, aView)
	if err != nil {
		return nil, err
	}

	execService := r.ensureExecutor()
	return session.State, execService.Exec(ctx, session)
}

func (r *Service) ExecSession(ctx context.Context, session *executor.Session) (*est.State, error) {
	execService := r.ensureExecutor()
	err := execService.Exec(ctx, session)
	return session.State, err
}

func (r *Service) newSelectors() *view.Selectors {
	return &view.Selectors{
		Index:   map[string]*view.Selector{},
		RWMutex: sync.RWMutex{},
	}
}
