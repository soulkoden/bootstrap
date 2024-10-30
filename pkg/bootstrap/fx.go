package bootstrap

import (
	"strconv"

	"go.uber.org/fx"
)

func BindService[T any](group string, factory any) any {
	return fx.Annotate(factory, fx.As(new(T)), fx.ResultTags(groupToTag(group)))
}

func GetServices[T any](group string) any {
	return fx.Annotate(func(factories []T) []T {
		return factories
	}, fx.ParamTags(groupToTag(group)))
}

func Provide[T any](group string, factories ...any) fx.Option {
	options := make([]any, 0, len(factories)+1)
	for _, factory := range factories {
		options = append(options, BindService[T](group, factory))
	}

	return fx.Provide(append(options, GetServices[T](group))...)
}

func groupToTag(group string) string {
	return `group:` + strconv.Quote(group)
}
