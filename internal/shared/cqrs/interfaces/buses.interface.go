package cqrsinterfaces

import "context"

type ICommandBus interface {
	Execute(ctx context.Context, command ICommand) (any, error)
	Register(commandName string, handler any) error
}

type IQueryBus interface {
	Execute(ctx context.Context, query IQuery) (any, error)
	Register(queryName string, handler any) error
}

type IEventBus interface {
	Publish(ctx context.Context, event IEvent) error
	Subscribe(eventName string, handler any) error
}
