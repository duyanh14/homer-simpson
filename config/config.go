package config

import (
	"bytes"
	"strings"

	"github.com/spf13/viper"
)

//  ssh -N -L 3336:127.0.0.1:3306 ducnp09081998@85.10.205.173
var yamlConfig = []byte(`
db_migration_data: false
cors_local: true
http_address: 8080
postgres:
  username: postgres
  password: postgres
  database: postgres
  host: localhost
  port: 5432
admin_dev:
  - ducnp
  - dev
ignore_authen:
  - Login
  - Register
jwt:
  private_key: ""
  public_key: ""
  signing_method: RS512
  is_refresh_token: false
  refresh_token_expire: 216000
  long_token_expire: 216000
  short_token_expire: 300
  issuer: "simpson"
redis_setting:
  addrs:
  - redis-14409.c1.ap-southeast-1-1.ec2.cloud.redislabs.com:14409
  password: ZW2IrPDs7DboDQgVok952YhXFO5ezzmK
  database: 0 #redisshop
logger:
  mode: development
  disable_caller: false
  disable_stacktrace: false
  encoding: json
  level: debug
`)

type (
	// Config fsdf
	Config struct {
		CORSLocal       bool     `yaml:"cors_local" mapstructure:"cors_local"`
		DbMigrationdata bool     `yaml:"db_migration_data" mapstructure:"db_migration_data"`
		HTTPAddress     int      `yaml:"http_address" mapstructure:"http_address"`
		Postgres        Postgres `yaml:"postgres" mapstructure:"postgres"`
		Admindev        []string `yaml:"admin_dev" mapstructure:"admin_dev"`
		JWT             Jwt      `yaml:"jwt" mapstructure:"jwt"`
		IgnoreAuthen    []string `yaml:"ignore_authen" mapstructure:"ignore_authen"`
		Redis           Redis    `yaml:"redis_setting" mapstructure:"redis_setting"`
		Logger          Logger   `yaml:"logger" mapstructure:"logger"`
	}

	Logger struct {
		Mode              string `yaml:"mode" mapstructure:"mode"`
		DisableCaller     bool   `yaml:"disable_caller" mapstructure:"disable_caller"`
		DisableStacktrace bool   `yaml:"disable_stacktrace" mapstructure:"disable_stacktrace"`
		Encoding          string `yaml:"encoding" mapstructure:"encoding"`
		Level             string `yaml:"level" mapstructure:"level"`
	}

	Jwt struct {
		IsRefreshToken         bool   `yaml:"is_refresh_token" mapstructure:"is_refresh_token"`
		SigningMethod          string `yaml:"signing_method" mapstructure:"signing_method"`
		PrivateKey             string `yaml:"private_key" mapstructure:"private_key"`
		PublicKey              string `yaml:"public_key" mapstructure:"public_key"`
		RefreshTokenExpireTime uint32 `yaml:"refresh_token_expire" mapstructure:"refresh_token_expire"`
		LongTokenExpireTime    uint32 `yaml:"long_token_expire" mapstructure:"long_token_expire"`
		ShortTokenExpireTime   uint32 `yaml:"short_token_expire" mapstructure:"short_token_expire"`
		Issuer                 string `yaml:"issuer" mapstructure:"issuer"`
	}

	// Postgres
	Postgres struct {
		Username    string `yaml:"username" mapstructure:"username"`
		Password    string `yaml:"password" mapstructure:"password"`
		Database    string `yaml:"database" mapstructure:"database"`
		Host        string `yaml:"host" mapstructure:"host"`
		Port        int    `yaml:"port" mapstructure:"port"`
		Migrate     bool   `yaml:"migrate" mapstructure:"migrate"`
		MaxIdleConn int    `yaml:"max_idle_conn" mapstructure:"max_idle_conn"`
		MaxOpenConn int    `yaml:"max_open_conn" mapstructure:"max_open_conn"`
		MaxLifeTime int    `yaml:"max_life_time" mapstructure:"max_life_time"` // hour
	}

	Redis struct {
		Addrs               []string
		Password            string `yaml:"password" mapstructure:"password"`
		Database            int    `yaml:"database" mapstructure:"database"`
		PoolSize            int    `yaml:"pool_size" mapstructure:"pool_size"`
		DialTimeoutSeconds  int    `yaml:"dial_timeout_seconds" mapstructure:"dial_timeout_seconds"`
		ReadTimeoutSeconds  int    `yaml:"read_timeout_seconds" mapstructure:"read_timeout_seconds"`
		WriteTimeoutSeconds int    `yaml:"write_timeout_seconds" mapstructure:"write_timeout_seconds"`
		IdleTimeoutSeconds  int    `yaml:"idle_timeout_seconds" mapstructure:"idle_timeout_seconds"`
	}
)

// LoadConfig is func load config for app
func LoadConfig() (*Config, error) {
	var conf = &Config{}
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(yamlConfig))
	if err != nil {
		return nil, err
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "__"))
	viper.AutomaticEnv()
	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
