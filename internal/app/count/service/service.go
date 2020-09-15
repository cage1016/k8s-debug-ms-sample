package service

import (
	"context"

	"github.com/go-kit/kit/log"

	"github.com/cage1016/ms-sample/internal/app/count/model"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(CountService) CountService

// Service describes a service that adds things together
// Implement yor service methods methods.
// e.x: Foo(ctx context.Context, s string)(rs string, err error)
type CountService interface {
	// [method=post,expose=true]
	Tic(ctx context.Context) (err error)
	// [method=get,expose=true]
	Tac(ctx context.Context) (res int64, err error)
}

// the concrete implementation of service interface
type stubCountService struct {
	logger log.Logger
	repo   model.CountRespository
}

// New return a new instance of the service.
// If you want to add service middleware this is the place to put them.
func New(repo model.CountRespository, logger log.Logger) (s CountService) {
	var svc CountService
	{
		svc = &stubCountService{
			repo: repo,
			logger: logger,
		}
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

// Implement the business logic of Tic
func (co *stubCountService) Tic(ctx context.Context) (err error) {
	return co.repo.Add(ctx)
}

// Implement the business logic of Tac
func (co *stubCountService) Tac(ctx context.Context) (res int64, err error) {
	return co.repo.Get(ctx)
}
