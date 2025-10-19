package cqrs

import (
	"context"
	"fmt"
	"lookerdevelopers/boilerplate/internal/shared/cqrs/interfaces"
	"lookerdevelopers/boilerplate/internal/shared/slog"
	"lookerdevelopers/boilerplate/internal/shared/utils"
	"reflect"
)

type QueryBus struct {
	handlers map[string]any
}

func NewQueryBus() cqrsinterfaces.IQueryBus {
	return &QueryBus{
		handlers: make(map[string]any),
	}
}

func ExecuteQuery[T any](ctx context.Context, queryBus cqrsinterfaces.IQueryBus, query cqrsinterfaces.IQuery) (T, error) {
	var zero T

	result, err := queryBus.Execute(ctx, query)
	if err != nil {
		return zero, err
	}

	if result == nil {
		return zero, nil
	}

	typedResult, ok := result.(T)

	if !ok {
		return zero, fmt.Errorf("queries result type mismatch: expected %T, got %T", zero, result)
	}

	return typedResult, nil
}

func (b *QueryBus) Execute(ctx context.Context, query cqrsinterfaces.IQuery) (any, error) {
	queryName := query.QueryName()

	handler, exists := b.handlers[queryName]

	if !exists {
		return nil, fmt.Errorf("no handler registered for queries: %s", queryName)
	}

	handlerValue := reflect.ValueOf(handler)
	handlerType := reflect.TypeOf(handler)

	handleMethod, found := handlerType.MethodByName("Handle")

	if !found {
		return nil, fmt.Errorf("handler for queries %s does not have Handle method", queryName)
	}

	args := []reflect.Value{
		handlerValue,           // receiver
		reflect.ValueOf(ctx),   // context
		reflect.ValueOf(query), // queries
	}

	slog.Logger.Infoln("Executing handler for commands", queryName)

	results := handleMethod.Func.Call(args)

	if len(results) != 2 {
		return nil, fmt.Errorf("handler must return (result, error)")
	}

	var result any
	var err error

	// We need pointers to know if the result is nil or not
	//if !results[0].IsNil() {
	//	result = results[0].Interface()
	//}
	//
	//if !results[1].IsNil() {
	//	err = results[1].Interface().(error)
	//}

	// The Query handlers should return structs
	if !utils.IsZeroValue(results[1]) {
		err = results[1].Interface().(error)
	}

	if !utils.IsZeroValue(results[0]) {
		result = results[0].Interface()
	}

	return result, err
}

func (b *QueryBus) Register(queryName string, handler interface{}) error {
	b.handlers[queryName] = handler
	return nil
}
