package config

import "net/http"

type App struct {
	Cookie *Cookie `mapstructure:"cookie" yaml:"cookie"`
}

type Cookie struct {
	Secret      string `mapstructure:"secret" yaml:"secret"`
	http.Cookie `mapstructure:",squash"`
}
