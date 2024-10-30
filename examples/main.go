package main

import (
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/soulkoden/bootstrap/pkg/bootstrap"
)

func main() {
	bootstrap.MustLoadEnv()
	zap.ReplaceGlobals(bootstrap.MustConfigureLogger())

	fx.New(
		// configure logging
		fx.Supply(zap.L()), // or not globally usage: fx.Supply(bootstrap.MustConfigureLogger())
		bootstrap.ZapLogger,

		// service groups usage
		bootstrap.Provide[marker]("svc", newService1, newService2),
		//// Same as:
		//fx.Provide(
		//  bootstrap.BindService[marker]("svc", newService1),
		//  bootstrap.BindService[marker]("svc", newService2),
		//  bootstrap.GetServices[marker]("svc"),
		//),

		fx.Invoke(func(services []marker, shutdowner fx.Shutdowner) error {
			for _, s := range services {
				fmt.Printf("service instance: %T implements marker\n", s)
			}

			return shutdowner.Shutdown()
		}),
	).Run()
}

type marker interface{}

type service1 struct{ marker }

func newService1() *service1 { return &service1{} }

type service2 struct{ marker }

func newService2() *service2 { return &service2{} }
