package cqrsinterfaces

import "context"

type IEvent interface {
	EventName() string
	OccurredOn() string
	AggregateID() string
}

type IEventHandler[T IEvent] interface {
	Handle(ctx context.Context, event T) error
}
