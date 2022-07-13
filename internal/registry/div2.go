package registry

import (
	"context"
	"fmt"
	"simpson/config"
	"simpson/internal/helper/logger"

	"go.uber.org/fx"
)

// var Module = fx.Provide(provideGormDB)

// func provideGormDB() (*gorm.DB, error) {
// 	uri := viper.GetString("MYSQL_URI")
// 	fmt.Println(uri)
// 	return nil, nil
// }

// var Module = fx.Provide(provideMongoDBClient)

// const defaultTimeout = 10 * time.Secondfunc

// func provideMongoDBClient(lifecycle fx.Lifecycle) (*mongo.Client, error) {
// 	mongoDBURI := viper.GetString(env.MongoURI)
// 	client, err := db.GetDBConnection(mongoDBURI)
// 	if err != nil {
// 		return nil, err
// 	}
// 	lifecycle.Append(fx.Hook{
// 		OnStart: func(ctx context.Context) error {
// 			ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
// 			defer cancel()
// 			return client.Connect(ctx)
// 		},
// 		OnStop: func(ctx context.Context) error {
// 			return client.Disconnect(ctx)
// 		},
// 	})
// 	return client, nil
// }

func invoke() {
	fmt.Println("invoke")
}
func BuidlContainerV2(ctx context.Context) error {
	app := fx.New(
		fx.Provide(
			config.LoadConfig,
		),
		fx.Invoke(
			logger.Newlogger,
			invoke,
		),
	)
	app.Run()
	return nil
}
