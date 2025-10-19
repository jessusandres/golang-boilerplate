package cqrsinterfaces

import "context"

type ICommand interface {
	CommandName() string
}

type ICommandHandler[T ICommand, R any] interface {
	Handle(ctx context.Context, command T) (R, error)
}
