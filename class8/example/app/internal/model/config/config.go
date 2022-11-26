package config

type Config struct {
	Logger *Logger `mapstructure:"logger" yaml:"logger"`

	DataBase *Database `mapstructure:"database"  yaml:"database"`

	App *App `mapstructure:"app"  yaml:"app"`

	Server *Server `mapstructure:"server"  yaml:"server"`

	Middleware *Middleware `mapstructure:"middleware" yaml:"middleware"`
}
