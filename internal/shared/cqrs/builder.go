package cqrs

import cqrsinterfaces "lookerdevelopers/boilerplate/internal/shared/cqrs/interfaces"

type Setup struct {
	CommandBus cqrsinterfaces.ICommandBus
	QueryBus   cqrsinterfaces.IQueryBus
}

func NewCQRSSetup() *Setup {
	return &Setup{
		CommandBus: NewCommandBus(),
		QueryBus:   NewQueryBus(),
	}
}
