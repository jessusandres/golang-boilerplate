package cqrs

import (
	"context"
	"fmt"
	llog "lookerdevelopers/boilerplate/internal/shared/slog"
	"lookerdevelopers/boilerplate/internal/shared/utils"
	"reflect"

	"lookerdevelopers/boilerplate/internal/shared/cqrs/interfaces"
)

type CommandBus struct {
	handlers map[string]any
}

func NewCommandBus() cqrsinterfaces.ICommandBus {
	return &CommandBus{
		handlers: make(map[string]any),
	}
}

func ExecuteCommand[T any](ctx context.Context, commandBus cqrsinterfaces.ICommandBus, command cqrsinterfaces.ICommand) (T, error) {
	var zero T

	result, err := commandBus.Execute(ctx, command)
	if err != nil {
		return zero, err
	}

	if result == nil {
		return zero, nil
	}

	typedResult, ok := result.(T)

	if !ok {
		return zero, fmt.Errorf("commands result type mismatch: expected %T, got %T", zero, result)
	}

	return typedResult, nil
}

func (b *CommandBus) Execute(ctx context.Context, command cqrsinterfaces.ICommand) (any, error) {
	commandName := command.CommandName()

	handler, exists := b.handlers[commandName]

	if !exists {
		return nil, fmt.Errorf("no handler registered for commands: %s", commandName)
	}

	handlerValue := reflect.ValueOf(handler)
	handlerType := reflect.TypeOf(handler)

	handleMethod, found := handlerType.MethodByName("Handle")

	if !found {
		return nil, fmt.Errorf("handler for commands %s does not have Handle method", commandName)
	}

	args := []reflect.Value{
		handlerValue,             // receiver
		reflect.ValueOf(ctx),     // context
		reflect.ValueOf(command), // commands
	}

	llog.Logger.Infoln("Executing handler for commands", commandName)

	results := handleMethod.Func.Call(args)

	if len(results) != 2 {
		return nil, fmt.Errorf("handler must return (result, error)")
	}

	var result any
	var err error

	// The Command handlers should return structs
	if !utils.IsZeroValue(results[1]) {
		err = results[1].Interface().(error)
	}

	if !utils.IsZeroValue(results[0]) {
		result = results[0].Interface()
	}

	return result, err
}

func (b *CommandBus) Register(commandName string, handler any) error {
	b.handlers[commandName] = handler

	return nil
}
