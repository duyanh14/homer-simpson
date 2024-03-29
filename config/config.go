package config

import (
	"bytes"
	"strings"

	"github.com/spf13/viper"
)

var yamlConfig = []byte(`
db_migration_data: false
cors_local: true
http_address: 11001
postgres:
  username: postgres
  password: postgres
  database: postgres
  host: localhost
  port: 5432
  migrate: true
  is_debug: true
admin_dev:
  - ducnp
  - dev
ignore_authen:
  - /user/register
  - /user/login
  - api/v1/state
jwt:
  private_key: "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS1FJQkFBS0NBZ0VBcVdjSWlweU1Wbm83RURTVnZiK20ydUpDSHV3S3pQRHI0ZU5vMHdsRU1qTGVZaHdQCnBzTUFtVUZ2dGg4bytzSVBlRDVIdmFMSUlkNFZvSnowcDVxM3FtUjE5eWI4cUlMYjJzWE95RDNCeXUvZmVsRnoKenQyWFFsZEpMRHFyQmRsZlRDL0VGZEdNb093enhzeldtWVhVV0dtRVRZenZpRFNaWXpYOWpvMnRWNVU4UEhEUgoyanZDNDdUTW1HakZsQWtHdmRMZXp3M0dPQjE2ZE1PbFNoK1Y3THBSL2krbmpnQk1IaTJHUFpDUXRFcTM4OTVmCk83c21OcWo3WjFSNHZlSXBPME5RL3BzQnA5akNpZGswQmRNcFlQTkp5TVFRQ3EvZFdXckxOWjl0V1JOL2RubFUKRDRBcWJ2bHRWaGNNTTkrTERBYm9XQ1RudmU4cGxRM2pVMksxZGREUEx5aU9DQ3REVStOeE5ZWVVqSlJLZTJKdgpsUUNFY2tkd0pBa3J2TmlESE5ObFU3UVdTb1hvZ3VBeFp4UVZsNWdnajEyeW9MVHpocWx2RVI1ZW9SdUxJUEF2Cm80V0U3a0FUM0pMSmN0b25LRy82ZTY1Wkx4Mk11K3JJUm5OSFNjOUNmZHkrYXg0K2RkQlliUTA2Z1lKVCtLNkgKMEJaOEQyK0RrNytxa2tWOGVsMjlxYUNlbkxmT2pwK1dWbVZWcHVwLzNOMnBBbnlwRloyQVJNS3FyUXZUaForRQp2UXM5ajZtdUxtWlJMWmVQZjRoUU1XRHRwNVhZWXAyMTh6VlZ3RUpXVGdTSHM1WVlhT21UeEJ6RHdBQWxvbDJPCmNkZENYVkpRdkQvMEJjZk5oTHp5TktCb2pkVWc2Y2xsZGdvcHpVZmVNU0s4ZFhFbzlhUnFUWCtoa3ZNQ0F3RUEKQVFLQ0FnQkNzYWZ4T3BVZDd0MldFYjNxaUt2ZWZUbmZCcnR4MkdCbC9LeG1lYjE1cmdGVHl3eDJjVXVwQ3RRNgpkYVlpbzRpSHBLMkdBYktkY3R3ODhjNTR0amNkUkpRVkRDcHBCYXdkUThlWG1valRwd0RySFdwa3hXVXAzMGVWCnZEdGRBQ20reXdkK041UVFQeTViVklZZ3gwbkV6VFlqZWh0TUdybFJFQTZWc3ozTnY1bUppamwwbkJhc281L3gKV1Q2QWU5VkZvdUhKenZoOVcwMjg3dElDanREdDlLVVhXQmUxbG5SaFpUYjJ3aFhXaFp4S3pQL2RlbDhmSWdSbwp1QjhaVXhNOG05SWxpMDFwa2JOMXBpZ1BLd1dkZit0S3ViWS9IdUljdENqeElhVmhib1crVlNJOVFBNTB4b3EvClhDK1h0S2ovd1RCMGs0TVJpamxxejU3eVUyaHhKOXgxcHN3NnYyZ2lXcjdXS2dsMm5PbFdNRU5LdHAvMXRDRnUKT2JvYVk1eUlYcEp4TDF3TDN4clJYNStzbzViQTBTdUZkU1hYWFcwMW1WdmV2Y2poUHMrQWJURnFuTktBK2dNagoydk4rMG1iY0ZiRlhwSFlUenk3eFJrVXhuaFd6VDFLVjI3Z2JnQ09Gc25GWXBPMFUxN0I0bDJnRDJPL0ZVY0ZoCjVnWjFXY1RrMGorcW51WVc5NlNmajRmTUhlTWJSSDJFbDJTcm5IT1VQekVvcS96bmgyTXNoK0RGUkhaemxvS1IKeFZpdCs5QndUNjVtcXVxTkQvajBOQWsway9sc2ErQU1KQjd2bHBiSGU2U1kva2d3QzJmS3ZiR0N6LzdrSytPVApEU0VyNkxEZ2d6alVPS3BrTEJ0azFQSklxdHVKSDcwUVVlaWZwaTgzRXY2Wm0rR0owUUtDQVFFQTNBLzZTOTRiCkxJQkdHa1pHanpDSHN3akpqaStzVms4TjdNY2VZL0pLaFc1dGxROTFpOW5QSFM2dkJtSWljZ0NqK0xmc0hBTlUKcXQ3a3B6Rm81RUNuKzVzV1JtZi9RUmtsWCs0eGdrSDgzRk9wUUZXMlhIdm4wekVONHoxRFZBRUFHZE54a1d1NQphWEJQQUJiNjJtdnY2RWRjN3FKNmNncEcrSjJwZTRUb2NuMGl3dWlTM25wRXhJb3AzdTlvZkhMa0NqdXp0Z05tCnY2TjZKVjc3SmtKaSs2V1VOVmNIblFETkl4WTZQM1RESmZwSTNvR1UxMHNuZDVMNW5IekFaZ3Y1T3FPVmVzdysKNE42ZTYxaFhKNGwwbDVpV1dqVXJsRlJ5MUFOQ0VxUW9pRnM3amo5OWlPTnVxTGhKNmp6QUJZVGFLTXR1alRHagprK0paQWdQNEk3RDl2UUtDQVFFQXhSRWxFdlhNaVBGWjFrQ25LWWpwVldGK0pCVzdYUGlsNVdUcWQrRUExOE8rClVFUU5BbXdkS1VlVmp0U3V5R3preDlPUlZOSWQvWFJZNjNkQytDNWZtV2xrU2Y5STJLUHNlRFRGTjZDTGxRbEUKaGdJRmlQanE3VG01SVRjOEFDV2h0RU9aVzZ5aDNtZ2pZWVRjL2psOHgwSm1zSUxiNmRJajQwYWlxRHZhd1g0OApLM3h3a0doRTEveU1aZTh1eCs3K1FHWmNMdUNTc3MrQjEvL2ZZTFBseFlxWHdISGhIUlNkdWZ5SmlLOUNRWCtBCjZveVJzS2VadGdpUmlQcG9aU1ZOSTQyN2lGOUpQRllCYzRhYkpRbmtxY2Nsb3RQcjhHZEw2clcwbHN3aUlDOXIKNHpoaTBtWTNuWm1Sby82bGs1TVNLZWZxNVhaVGJ4T2tObVV0TVBlbWJ3S0NBUUVBbktFcXI1RTRtaVg2djN6QgpmKzAweUxNNmNsOXdORVkwVFJFMGlrTEwwOGpUYWtNRWtUUEgvS2Rib0JsOUZ0TnA0ZjlXcUdHZWc3ZjkyeHFKCjQ5dzVOQllnZHFCU3g4elVFMEdBTDR3MEk4WFFNNkJPR2VsL0NZeGlyQmpRc1J1OUxLU1lzcjZQeTRKS1dIQUMKd0RBUEk2NlhjS3BvcUg2MElRUm11eW5RSUp6OGZ1bjJqeTlMZnZBditTTkJwVUpKR1JlQlE4Mkw5bDZ1ODZaMAo3N2hVcDNRazF3cEl1dkQxRVVJaTEwT09HcHZYL2JjUmIyTm9oZTB3aUxjcVlmZW52cHJzSytqdjRESFR2d0t6CmxjcDZGK1JkTDN1a3R3Q2NjbGFYVXBsUTJDekhhT203dWhEV2xIUUcrQmVoUE14Z1VxVG8zTW8xK1c0am1CR24KTGhBdHFRS0NBUUJ2bjlDQ050eU9UVGQzMXROWkFZWTVxS21Nd3dxSUZRZlRNaUJsbXZ1aE8rMFhCaFRnWkdENApPdVlLSzRwOVdNRVdOMDdBM3V2QjN4OW04UHpzRzBoblczMUZOT2NNWDMzWWNQdXRFTEUvMENvU1JoN2dnUzZCCkJRdEtOMEV6VEIyV1FCd2tBMXNFNGJQNHp3dG1yU1Z5c0xmK0Q0R0ZwbFJScm5jQUdEZEhGcnY2WGRoYzA5TDcKRE1CeUZOTkl5S1VYMFdNeDRsNzJEZmdjWTRFaUE0U0pLb0hlcmlLM0dEQm5ZeFo4WjlsOXhEMC9ualAvL2s5ZwovdVdBMDJaQ3RLaEhGVWMrYmNyTDlHT0tEcHJlbGdZQjdSM3laMXZqcG44b0NaaXJPaUd5WFdvZElKbEpXRUQwCmNrSSt4RVVCbkhMVGJkeHQ4V0k5THNEN1ZzQU1WcHJkQW9JQkFRRFFhaUg5T29QL3psZC90S2YzakRFZmYrV00KdG5ZeVppM0lpclJCUE1GK2RHOEhpMWNXQVo0YklJREpNSjhxQnpYTGdFU1p5VmJGR3hROUIrV0JhZ1ByZUNJVQo2YlJvQ0UwZXJrRDgvQ2hBUVV0V21JckFITFNET1VHaVdOK2lNSVp4M3BNTmtJcUk2UUMxNlVZY0hpMkh2Skk2Ck1tcHhQMllIRUF3NEI2V29IbTlTRHhjejFENnZJcUhLUVVFemNFT2h2UnpkN3ZZb0U2U0QvRk0vT1Rhbnh4cksKdnpGajF6dGpDZWR2U2VxV2tWYkdGb0l1dmwxWHc3QXJWSlIyM21HZHVwemFEQU5WVmI5bXFJMHVWWG4zN0xrYwpta3hVVVpWN0JvR0x2ZndmN3pKNXhqM0hLeUVTK0dBNHlrMHJ5eDA2Q0hWS1lkRFFzSmVhcUs5VjBOQlcKLS0tLS1FTkQgUlNBIFBSSVZBVEUgS0VZLS0tLS0K"
  public_key: "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQ0lqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FnOEFNSUlDQ2dLQ0FnRUFxV2NJaXB5TVZubzdFRFNWdmIrbQoydUpDSHV3S3pQRHI0ZU5vMHdsRU1qTGVZaHdQcHNNQW1VRnZ0aDhvK3NJUGVENUh2YUxJSWQ0Vm9KejBwNXEzCnFtUjE5eWI4cUlMYjJzWE95RDNCeXUvZmVsRnp6dDJYUWxkSkxEcXJCZGxmVEMvRUZkR01vT3d6eHN6V21ZWFUKV0dtRVRZenZpRFNaWXpYOWpvMnRWNVU4UEhEUjJqdkM0N1RNbUdqRmxBa0d2ZExlenczR09CMTZkTU9sU2grVgo3THBSL2krbmpnQk1IaTJHUFpDUXRFcTM4OTVmTzdzbU5xajdaMVI0dmVJcE8wTlEvcHNCcDlqQ2lkazBCZE1wCllQTkp5TVFRQ3EvZFdXckxOWjl0V1JOL2RubFVENEFxYnZsdFZoY01NOStMREFib1dDVG52ZThwbFEzalUySzEKZGREUEx5aU9DQ3REVStOeE5ZWVVqSlJLZTJKdmxRQ0Vja2R3SkFrcnZOaURITk5sVTdRV1NvWG9ndUF4WnhRVgpsNWdnajEyeW9MVHpocWx2RVI1ZW9SdUxJUEF2bzRXRTdrQVQzSkxKY3RvbktHLzZlNjVaTHgyTXUrcklSbk5IClNjOUNmZHkrYXg0K2RkQlliUTA2Z1lKVCtLNkgwQlo4RDIrRGs3K3Fra1Y4ZWwyOXFhQ2VuTGZPanArV1ZtVlYKcHVwLzNOMnBBbnlwRloyQVJNS3FyUXZUaForRXZRczlqNm11TG1aUkxaZVBmNGhRTVdEdHA1WFlZcDIxOHpWVgp3RUpXVGdTSHM1WVlhT21UeEJ6RHdBQWxvbDJPY2RkQ1hWSlF2RC8wQmNmTmhMenlOS0JvamRVZzZjbGxkZ29wCnpVZmVNU0s4ZFhFbzlhUnFUWCtoa3ZNQ0F3RUFBUT09Ci0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
  signing_method: RS512
  is_refresh_token: false
  refresh_token_expire: 216000
  long_token_expire: 216000000
  short_token_expire: 21600
  issuer: "simpson"
  validate_password: false
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
  zap_type: null
  log_file: true
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
		LogFile           bool   `yaml:"log_file" mapstructure:"log_file"`
		Encoding          string `yaml:"encoding" mapstructure:"encoding"`
		Level             string `yaml:"level" mapstructure:"level"`
		ZapType           string `yaml:"zap_type" mapstructure:"zap_type"`
	}

	Jwt struct {
		SigningMethod          string `yaml:"signing_method" mapstructure:"signing_method"`
		PrivateKey             string `yaml:"private_key" mapstructure:"private_key"`
		PublicKey              string `yaml:"public_key" mapstructure:"public_key"`
		Issuer                 string `yaml:"issuer" mapstructure:"issuer"`
		RefreshTokenExpireTime uint32 `yaml:"refresh_token_expire" mapstructure:"refresh_token_expire"`
		LongTokenExpireTime    uint32 `yaml:"long_token_expire" mapstructure:"long_token_expire"`
		ShortTokenExpireTime   uint32 `yaml:"short_token_expire" mapstructure:"short_token_expire"`
		IsRefreshToken         bool   `yaml:"is_refresh_token" mapstructure:"is_refresh_token"`
		ValidatePassword       bool   `yaml:"validate_password" mapstructure:"validate_password"`
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
		IsDebug     bool   `yaml:"is_debug" mapstructure:"is_debug"`
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
