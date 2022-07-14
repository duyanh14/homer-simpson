package registry

import (
	"context"
	"sync"

	"github.com/sarulabs/di"
)

var (
	buildOnce sync.Once
	container di.Container
	builder   *di.Builder
)

const (
	APIDI = "APIDI"
)

func buildAPI() error {
	defs := []di.Def{}
	apiDI := di.Def{
		Name:  APIDI,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return nil, nil
		},
		Close: func(obj interface{}) error {
			return nil
		},
	}
	defs = append(defs, apiDI)

	//

	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}

func usecaseBuild() error {
	return nil
}

func adapterBuild() error {
	return nil
}

func helperBuild() error {
	return nil
}

func serviceBuild() error {
	return nil
}

func configBuild() error {
	return nil
}

func BuildContainer(ctx context.Context) {
	buildOnce.Do(func() {
		builder, err := di.NewBuilder()
		if err != nil {
			panic(err)
		}

		if err := buildAPI(); err != nil {
			panic(err)
		}

		if err := configBuild(); err != nil {
			panic(err)
		}

		if err := usecaseBuild(); err != nil {
			panic(err)
		}

		if err := serviceBuild(); err != nil {
			panic(err)
		}

		if err := helperBuild(); err != nil {
			panic(err)
		}
		if err := adapterBuild(); err != nil {
			panic(err)
		}

		container = builder.Build()
	})
}
