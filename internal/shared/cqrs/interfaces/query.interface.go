package cqrsinterfaces

import "context"

type IQuery interface {
	QueryName() string
}

type IQueryHandler[T IQuery, R any] interface {
	Handle(ctx context.Context, query T) (R, error)
}
