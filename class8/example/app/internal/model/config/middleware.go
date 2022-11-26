package config

type Middleware struct {
	Cors *CORS `mapstructure:"cors" yaml:"cors"`
	Jwt  *Jwt  `mapstructure:"jwt" yaml:"jwt"`
}

type CORS struct {
	Mode      string          `mapstructure:"mode" yaml:"mode"`
	Whitelist []CORSWhitelist `mapstructure:"whitelist" yaml:"whitelist"`
}

type CORSWhitelist struct {
	AllowOrigin      string `mapstructure:"allowOrigin" yaml:"allowOrigin"`
	AllowMethods     string `mapstructure:"allowMethods" yaml:"allowMethods"`
	AllowHeaders     string `mapstructure:"allowHeaders" yaml:"allowHeaders"`
	ExposeHeaders    string `mapstructure:"exposeHeaders" yaml:"exposeHeaders"`
	AllowCredentials bool   `mapstructure:"allowCredentials" yaml:"allowCredentials"`
}

type Jwt struct {
	SecretKey   string `mapstructure:"secretKey" yaml:"secretKey"`
	ExpiresTime int64  `mapstructure:"expiresTime" yaml:"expiresTime"`
	BufferTime  int64  `mapstructure:"bufferTime" yaml:"bufferTime"`
	Issuer      string `mapstructure:"issuer" yaml:"issuer"`
}
